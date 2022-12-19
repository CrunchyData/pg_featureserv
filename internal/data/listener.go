package data

/*
 Copyright 2019 Crunchy Data Solutions, Inc.
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
      http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.

 Date     : October 2022
 Authors  : Amaury Zarzelli (amaury dot zarzelli at ign dot fr)
*/

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/CrunchyData/pg_featureserv/internal/api"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

// TODO: make the schema name configurable
const tempDBSchema = "pgfeatureserv"

// A listenerDB is associated to a catalogDB, and manages the operations required for listening
// the events occuring on the database. This includes creating the trigger function in the base,
// applying the trigger function to the tables included in pg_featureserv, and listening to
// events on those tables
type listenerDB struct {
	dbconn        *pgxpool.Pool      // connection to database
	tableIncludes map[string]string  // list of included tables
	tableExcludes map[string]string  // list of excluded tables
	cache         Cacher             // cache of the catalog
	stopListen    context.CancelFunc // channel used to stop the listen goroutine
}

// An eventNotification is a notification sent by the database after a INSERT, UPDATE or DELETE
// event on the databases included in pg_featureserv. It is populated using the return value of
// the pl/pgSQL procedure named `sqlNotifyFunction` defined in db_sql.go
type eventNotification struct {
	Schema   string                 // schema of the table triggering the event
	Table    string                 // name of the table triggering the event
	Action   string                 // action triggering the event (INSERT, UPDATE or DELETE)
	Old_xmin string                 // xmin of the previous version of the row (`nil` in case of INSERT)
	New_xmin string                 // xmin of the new version of the row (`nil` in case of INSERT)
	Data     map[string]interface{} // data contained in the row
}

// toString for eventNotification
func (e eventNotification) String() string {
	return fmt.Sprintf("eventNotification[table: '%v.%v', action: '%v', xmin: %v/%v, data: %v]", e.Schema, e.Table, e.Action, e.Old_xmin, e.New_xmin, e.Data)
}

// creates new db listener
func newListenerDB(conn *pgxpool.Pool, cache Cacher) *listenerDB {

	listener := &listenerDB{
		dbconn: conn,
		cache:  cache,
	}

	return listener
}

// Initialize the listener using include and exclude maps to:
//   - add temporary DB schema
//   - add trigger function temp schema
//   - add trigger functions to included tables
//   - start listening to database operations
func (listener *listenerDB) Initialize(tableIncludes map[string]string, tableExcludes map[string]string) {
	listener.tableIncludes = tableIncludes
	listener.tableExcludes = tableExcludes

	ctx := context.Background()
	ctxGoroutine, stopListen := context.WithCancel(ctx)
	listener.stopListen = stopListen

	listener.addTemporaryDBSchema()
	listener.addTriggerFunctionToDB()
	listener.addTriggerToTables()
	go listener.listen(ctxGoroutine)
}

// Listen for INSERT or UPDATE or DELETE using triggers
// pgxPool can't listen, code snippet from https://github.com/jackc/pgx/issues/1121
func (listener *listenerDB) listen(ctx context.Context) {
	for {
		listener.listenOneNotification(ctx)
		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}

func (listener *listenerDB) listenOneNotification(ctx context.Context) {
	listenConn, err := listener.dbconn.Acquire(ctx)
	if err != nil {
		if pgconn.Timeout(err) {
			return
		}
		log.Fatal(err)
	}
	defer listenConn.Release()

	_, err = listenConn.Exec(ctx, "LISTEN table_update")
	if err != nil {
		if pgconn.Timeout(err) {
			return
		}
		log.Fatal(err)
	}

	notification, err := listenConn.Conn().WaitForNotification(ctx)
	if err != nil {
		if pgconn.Timeout(err) {
			return
		}
		log.Fatal(err)
	}

	var notificationData eventNotification

	errUnMarsh := json.Unmarshal([]byte(notification.Payload), &notificationData)
	if errUnMarsh != nil {
		log.Fatal(errUnMarsh)
	}

	log.Debugf("Listener received notification: %v, cache: %v", notificationData, listener.cache)
	if notificationData.Action == "DELETE" || notificationData.Action == "UPDATE" {
		weakEtag := api.MakeWeakEtag("", "", notificationData.Old_xmin, "")
		_, err = listener.cache.RemoveWeakEtag(weakEtag.CacheKey())
		if err != nil {
			log.Warnf("Error removing weak Etag to cache: %v", err)
		}
	}
	if notificationData.Action == "INSERT" || notificationData.Action == "UPDATE" {
		collection := fmt.Sprintf(`%s.%s`, notificationData.Schema, notificationData.Table)
		// ==== retrieve tabe data
		table, errCat := CatDBInstance().TableByName(collection)
		if errCat != nil {
			log.Warnf("Listener received notification about unknown table '%v'. Error: %v", collection, errCat.Error())

		} else {
			// ==== retrieve the id
			var id string
			switch table.DbTypes[table.IDColumn].Type {
			case api.PGTypeText, api.PGTypeVarChar:
				id = notificationData.Data[table.IDColumn].(string)
			case api.PGTypeFloat8, api.PGTypeFloat4, api.PGTypeInt, api.PGTypeInt4, api.PGTypeInt8:
				id = fmt.Sprintf("%f", notificationData.Data[table.IDColumn].(float64))
			default:
				log.Warnf("Listener received notification about table '%v' with unhandled id '%v' of type '%v'.",
					collection, table.IDColumn, table.DbTypes[table.IDColumn].Type)
				id = ""
			}

			if id != "" {
				weakEtag := api.MakeWeakEtag(collection, id, notificationData.New_xmin, api.GetCurrentHttpDate())
				weakEtag.Data = notificationData.Data

				// ===== DOUBLE ADD!!
				_, err = listener.cache.AddWeakEtag(weakEtag.CacheKey(), weakEtag)
				if err != nil {
					log.Warnf("Error adding weak Etag to cache: %v", err)
				}
				_, err = listener.cache.AddWeakEtag(weakEtag.AlternateCacheKey(), weakEtag)
				if err != nil {
					log.Warnf("Error adding weak Etag to cache: %v", err)
				}
			}
		}
	}
}

func (listener *listenerDB) Close() {
	if listener.stopListen != nil {
		listener.stopListen()
	}
	listener.dropTriggers()
	listener.dropTemporaryDBSchema()
}

func (listener *listenerDB) addTemporaryDBSchema() {
	sqlStatement := "CREATE SCHEMA IF NOT EXISTS %s"
	_, errExec := listener.dbconn.Exec(context.Background(), fmt.Sprintf(sqlStatement, tempDBSchema))
	if errExec != nil {
		log.Fatal(errExec)
	}
}

func (listener *listenerDB) dropTemporaryDBSchema() {
	sqlStatement := "DROP SCHEMA IF EXISTS %s CASCADE"
	_, errExec := listener.dbconn.Exec(context.Background(), fmt.Sprintf(sqlStatement, tempDBSchema))
	if errExec != nil {
		log.Fatal(errExec)
	}
}

func (listener *listenerDB) addTriggerFunctionToDB() {
	_, errExec := listener.dbconn.Exec(context.Background(), fmt.Sprintf(sqlNotifyFunction, tempDBSchema))
	if errExec != nil {
		log.Fatal(errExec)
	}
}

func (listener *listenerDB) addTriggerToTables() {
	log.Debugf("Add trigger to tables:\n%v", sqlTables)
	rows, err := listener.dbconn.Query(context.Background(), sqlTables)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		tbl := scanTable(rows)
		if isIncluded(tbl, listener.tableIncludes, listener.tableExcludes) {
			listener.addTriggerToTable(tbl)
		}
	}
	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	rows.Close()
}

func (listener *listenerDB) addTriggerToTable(tbl *api.Table) {
	dropTriggerBytes := []byte(`
	DROP TRIGGER IF EXISTS %[1]s_notify_event ON %[2]s;
	`)
	triggerBytes := []byte(`
	CREATE TRIGGER %[1]s_notify_event
	AFTER INSERT OR UPDATE OR DELETE ON %[2]s
	FOR EACH ROW EXECUTE PROCEDURE %[3]s.notify_event();
	`)

	triggerName := tbl.Schema + "_" + tbl.Table
	tableName := tbl.Schema + "." + tbl.Table
	dropTriggerStatement := fmt.Sprintf(string(dropTriggerBytes), triggerName, tableName)
	triggerStatement := fmt.Sprintf(string(triggerBytes), triggerName, tableName, tempDBSchema)
	_, errDrop := listener.dbconn.Exec(context.Background(), dropTriggerStatement)
	if errDrop != nil {
		log.Fatal(errDrop)
	}
	_, err := listener.dbconn.Exec(context.Background(), triggerStatement)
	if err != nil {
		log.Fatal(err)
	}
}

func (listener *listenerDB) dropTriggers() {
	log.Debugf("Load table catalog:\n%v", sqlTables)
	rows, err := listener.dbconn.Query(context.Background(), sqlTables)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		tbl := scanTable(rows)
		if isIncluded(tbl, listener.tableIncludes, listener.tableExcludes) {
			listener.dropTrigger(tbl)
		}
	}
	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	rows.Close()
}

func (listener *listenerDB) dropTrigger(tbl *api.Table) {
	dropTriggerBytes := []byte(`
	DROP TRIGGER IF EXISTS %[1]s_notify_event ON %[2]s;
	`)

	triggerName := tbl.Schema + "_" + tbl.Table
	tableName := tbl.Schema + "." + tbl.Table
	dropTriggerStatement := fmt.Sprintf(string(dropTriggerBytes), triggerName, tableName)

	_, errDrop := listener.dbconn.Exec(context.Background(), dropTriggerStatement)
	if errDrop != nil {
		log.Fatal(errDrop)
	}
}
