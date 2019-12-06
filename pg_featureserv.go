package main

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
*/

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/CrunchyData/pg_featureserv/api"
	"github.com/CrunchyData/pg_featureserv/config"
	"github.com/CrunchyData/pg_featureserv/data"
	"github.com/gorilla/mux"
)

// CatalogInstance mock
var catalogInstance data.Catalog

func init() {
	catalogInstance = data.InstanceCatMock()
}

func serveURLBase(r *http.Request) string {
	// Preferred host:port
	php := r.Host
	php = strings.TrimRight(php, "/")

	// Preferred scheme
	ps := "http"

	// Preferred base path
	pbp := "/"

	// Preferred scheme / host / port / base
	pshpb := fmt.Sprintf("%v://%v%v", ps, php, pbp)
	return pshpb
}

func serveRequests() {
	confServ := config.Configuration.Server
	bindAddress := fmt.Sprintf("%v:%v", confServ.BindHost, confServ.BindPort)
	fmt.Printf("Serving at: %v\n", bindAddress)

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", handleRootJSON)
	myRouter.HandleFunc("/home{.fmt}", handleHome)

	myRouter.HandleFunc("/conformance", handleConformance)
	myRouter.HandleFunc("/conformance.{fmt}", handleConformance)

	myRouter.HandleFunc("/collections", handleCollections)
	myRouter.HandleFunc("/collections.{fmt}", handleCollections)

	myRouter.HandleFunc("/collections/{cid}", handleCollection)
	myRouter.HandleFunc("/collections/{cid}.{fmt}", handleCollection)

	myRouter.HandleFunc("/collections/{cid}/items", handleCollectionItems)
	myRouter.HandleFunc("/collections/{cid}/items.{fmt}", handleCollectionItems)

	myRouter.HandleFunc("/collections/{cid}/items/{fid}", handleItem)
	myRouter.HandleFunc("/collections/{cid}/items/{fid}.{fmt}", handleItem)

	log.Fatal(http.ListenAndServe(bindAddress, myRouter))
}

func getRequestVar(varname string, r *http.Request) string {
	vars := mux.Vars(r)
	nameFull := vars[varname]
	name := api.PathStripFormat(nameFull)
	return name
}

func main() {
	serveRequests()
}
