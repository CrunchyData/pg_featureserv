package utiltest

/*
 Copyright 2022 Crunchy Data Solutions, Inc.
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
      http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

import (
	"context"
	"fmt"
	"os"

	"github.com/CrunchyData/pg_featureserv/internal/api"
	"github.com/CrunchyData/pg_featureserv/internal/conf"
	"github.com/CrunchyData/pg_featureserv/internal/data"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

func CreateTestDb() *pgxpool.Pool {
	dbURL := os.Getenv(conf.AppConfig.EnvDBURL)
	if dbURL == "" {
		dbURL = "postgresql://postgres@localhost/pg_featureserv"
		log.Warnf("No env var '%s' defined, using default value: %s", conf.AppConfig.EnvDBURL, dbURL)
	}
	conf.Configuration.Database.DbConnection = dbURL
	conf.Configuration.Database.DbPoolMaxConnLifeTime = "1h"

	ctx := context.Background()
	dbconfig, errConf := pgxpool.ParseConfig(conf.Configuration.Database.DbConnection)
	if errConf != nil {
		log.Fatal(errConf)
	}
	db, errConn := pgxpool.ConnectConfig(ctx, dbconfig)
	if errConn != nil {
		log.Fatal(errConn)
	}

	dbName := dbconfig.ConnConfig.Config.Database
	dbUser := dbconfig.ConnConfig.Config.User
	dbHost := dbconfig.ConnConfig.Config.Host
	log.Debugf("Connected as %s to %s @ %s", dbUser, dbName, dbHost)

	insertSimpleDataset(db)

	log.Debugf("Sample data injected")

	return db
}

func insertSimpleDataset(db *pgxpool.Pool) {
	ctx := context.Background()
	// collections tables
	// tables := []string{"mock_a", "mock_b", "mock_c"}
	type tableContent struct {
		extent api.Extent
		nx     int
		ny     int
	}
	tablesAndExtents := map[string]tableContent{
		"mock_a": {api.Extent{Minx: -120, Miny: 40, Maxx: -74, Maxy: 50}, 3, 3},
		"mock_b": {api.Extent{Minx: -75, Miny: 45, Maxx: -74, Maxy: 46}, 10, 10},
		"mock_c": {api.Extent{Minx: -120, Miny: 40, Maxx: -74, Maxy: 60}, 100, 100},
	}

	createBytes := []byte(`
		DROP TABLE IF EXISTS %s CASCADE;
		CREATE TABLE IF NOT EXISTS public.%s (
			id SERIAL PRIMARY KEY,
			geometry public.geometry(Point, 4326),
			prop_a text,
			prop_b int,
			prop_c text,
			prop_d int
		);
	`)
	for s := range tablesAndExtents {

		createStatement := fmt.Sprintf(string(createBytes), s, s)

		_, errExec := db.Exec(ctx, createStatement)
		if errExec != nil {
			CloseTestDb(db)
			log.Fatal(errExec)
		}
	}

	// collections features/table records
	b := &pgx.Batch{}

	insertBytes := []byte(`
		INSERT INTO public.%s (geometry, prop_a, prop_b, prop_c, prop_d)
		VALUES (ST_GeomFromGeoJSON($1), $2, $3, $4, $5)
	`)
	for tableName, tableElements := range tablesAndExtents {
		insertStatement := fmt.Sprintf(string(insertBytes), tableName)
		featuresMock := data.MakePointFeatures(tableElements.extent, tableElements.nx, tableElements.ny)

		for _, f := range featuresMock {
			geomStr, _ := f.Geom.MarshalJSON()
			b.Queue(insertStatement, geomStr, f.Props["prop_a"], f.Props["prop_b"], f.Props["prop_c"], f.Props["prop_d"])
		}
		res := db.SendBatch(ctx, b)
		if res == nil {
			CloseTestDb(db)
			log.Fatal("Injection failed")
		}
		resClose := res.Close()
		if resClose != nil {
			CloseTestDb(db)
			log.Fatal(fmt.Sprintf("Injection failed: %v", resClose.Error()))
		}
	}

	log.Debugf("Sample data injected")
}

func CloseTestDb(db *pgxpool.Pool) {
	log.Debugf("Sample db will be cleared...")
	_, errExec := db.Exec(context.Background(), "DROP TABLE IF EXISTS public.mock_a CASCADE;")
	if errExec != nil {
		log.Warnf("Failed to drop sample db! ")
		log.Warnf(errExec.Error())
	}
	db.Close()
	log.Debugf("Sample db cleared!")
}
