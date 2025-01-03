package util

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
	"os"

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

	// same as featureMock
	_, errExec := db.Exec(ctx, `
		DROP TABLE IF EXISTS mock_a CASCADE;
		CREATE TABLE IF NOT EXISTS public.mock_a (
			id SERIAL PRIMARY KEY,
			geometry public.geometry(Point, 4326),
			prop_a text,
			prop_b int,
			prop_c text,
			prop_d int
		);
	`)
	if errExec != nil {
		CloseTestDb(db)
		log.Fatal(errExec)
	}

	feats := data.MakePointFeatures(data.Extent{Minx: -120, Miny: 40, Maxx: -74, Maxy: 50},
		3, 3)
	b := &pgx.Batch{}
	sqlStatement := `
		INSERT INTO public.mock_a (geometry, prop_a, prop_b, prop_c, prop_d)
		VALUES (ST_GeomFromGeoJSON($1), $2, $3, $4, $5)`

	for _, f := range feats {
		geomStr, _ := f.Geom.MarshalJSON()
		b.Queue(sqlStatement, geomStr, f.PropA, f.PropB, f.PropC, f.PropD)
	}
	res := db.SendBatch(ctx, b)
	if res == nil {
		CloseTestDb(db)
		log.Fatal("Injection failed")
	}
	res.Close()

	log.Debugf("Sample data injected")

	return db
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