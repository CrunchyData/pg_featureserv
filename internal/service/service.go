package service

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/CrunchyData/pg_featureserv/internal/api"
	"github.com/CrunchyData/pg_featureserv/internal/conf"
	"github.com/CrunchyData/pg_featureserv/internal/data"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

/*
 Copyright 2019 - 2025 Crunchy Data Solutions, Inc.
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

var catalogInstance data.Catalog
var router *mux.Router
var server *http.Server
var isTLSEnabled bool
var serverTLS *http.Server

// Initialize sets the service state from configuration
func Initialize() {
	initTransforms(conf.Configuration.Server.TransformFunctions)
}

func createServers() {
	confServ := conf.Configuration.Server

	bindAddress := fmt.Sprintf("%v:%v", confServ.HttpHost, confServ.HttpPort)
	bindAddressTLS := fmt.Sprintf("%v:%v", confServ.HttpHost, confServ.HttpsPort)
	// Use HTTPS only if server certificate and private key files specified
	isTLSEnabled = conf.Configuration.IsTLSEnabled()

	log.Infof("Serving HTTP  at %s", formatBaseURL("http://", bindAddress, confServ.BasePath))
	if isTLSEnabled {
		log.Infof("Serving HTTPS at %s", formatBaseURL("https://", bindAddressTLS, confServ.BasePath))
	}
	log.Infof("CORS Allowed Origins: %v\n", conf.Configuration.Server.CORSOrigins)

	router = initRouter(confServ.BasePath)

	// writeTimeout is slighlty longer than request timeout to allow writing error response
	timeoutSecRequest := conf.Configuration.Server.WriteTimeoutSec
	timeoutSecWrite := timeoutSecRequest + 1

	// ----  Handler chain  --------
	// set CORS handling according to config
	corsOpt := handlers.AllowedOrigins([]string{conf.Configuration.Server.CORSOrigins})
	corsHandler := handlers.CORS(corsOpt)(router)
	compressHandler := handlers.CompressHandler(corsHandler)

	// Use a TimeoutHandler to ensure a request does not run past the WriteTimeout duration.
	// This provides a context that allows cancellation to be propagated
	// down to the database driver.
	//(Unfortunately this does not propagate to the database itself.
	// That will require another mechanism such as session config statement_timeout)
	// If timeout expires, service returns 503 and a text message
	timeoutHandler := http.TimeoutHandler(compressHandler,
		time.Duration(timeoutSecRequest)*time.Second,
		api.ErrMsgRequestTimeout)

	// more "production friendly" timeouts
	// https://blog.simon-frey.eu/go-as-in-golang-standard-net-http-config-will-break-your-production/#You_should_at_least_do_this_The_easy_path
	server = &http.Server{
		ReadTimeout:  time.Duration(conf.Configuration.Server.ReadTimeoutSec) * time.Second,
		WriteTimeout: time.Duration(timeoutSecWrite) * time.Second,
		Addr:         bindAddress,
		Handler:      timeoutHandler,
	}

	if isTLSEnabled {
		serverTLS = &http.Server{
			ReadTimeout:  time.Duration(conf.Configuration.Server.ReadTimeoutSec) * time.Second,
			WriteTimeout: time.Duration(timeoutSecWrite) * time.Second,
			Addr:         bindAddressTLS,
			Handler:      timeoutHandler,
			TLSConfig: &tls.Config{
				MinVersion: tls.VersionTLS12, // Secure TLS versions only
			},
		}
	}
}

// Serve starts the web service
func Serve(catalog data.Catalog) {
	confServ := conf.Configuration.Server
	catalogInstance = catalog
	createServers()

	log.Infof("====  Service: %s  ====\n", conf.Configuration.Metadata.Title)

	// start http service
	go func() {
		// ListenAndServe returns http.ErrServerClosed when the server receives
		// a call to Shutdown(). Other errors are unexpected.
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// start https service
	if isTLSEnabled {
		go func() {
			// ListenAndServe returns http.ErrServerClosed when the server receives
			// a call to Shutdown(). Other errors are unexpected.
			if err := serverTLS.ListenAndServeTLS(confServ.TlsServerCertificateFile, confServ.TlsServerPrivateKeyFile); err != nil && err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}()
	}

	// wait here for interrupt signal (^C)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	// Interrupt signal received:  Start shutting down
	log.Infoln("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	errConn := server.Shutdown(ctx)
	if errConn != nil {
		log.Warnf("Server connection failed to shutdown: %v", errConn.Error())
	}
	if isTLSEnabled {
		errConnTls := serverTLS.Shutdown(ctx)
		if errConnTls != nil {
			log.Warnf("Server TLS connection failed to shutdown: %v", errConnTls.Error())
		}
	}

	// abort after waiting long enough for service to shutdown gracefully
	// this terminates long-running DB queries, which otherwise block shutdown
	abortTimeoutSec := conf.Configuration.Server.WriteTimeoutSec + 10
	chanCancelFatal := FatalAfter(abortTimeoutSec, "Timeout on shutdown - aborting.")

	log.Debugln("Closing DB connections")
	catalogInstance.Close()

	log.Infoln("Service stopped.")
	// cancel the abort since it is not needed
	close(chanCancelFatal)
}
