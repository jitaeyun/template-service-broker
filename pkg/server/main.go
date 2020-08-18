package main

import (
	"fmt"
	"net/http"
	"template-service-broker/pkg/server/apis"

	"github.com/gorilla/mux"
	"github.com/operator-framework/operator-sdk/pkg/log/zap"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	port = 8081

	apiPathPrefix         = "/v2/"
	serviceCatalogPrefix  = "/catalog"
	serviceInstancePrefix = "/service_instances/{instanceId}"
)

var log = logf.Log.WithName("TSB-main")

func main() {
	logf.SetLogger(zap.Logger())
	log.Info("initializing server....")

	router := mux.NewRouter()
	apiRouter := router.PathPrefix(apiPathPrefix).Subrouter()

	//catalog
	apiRouter.HandleFunc(serviceCatalogPrefix, apis.GetCatalog).Methods("GET")

	//provision
	apiRouter.HandleFunc(serviceInstancePrefix, apis.ProvisionServiceInstance).Methods("PUT")
	apiRouter.HandleFunc(serviceInstancePrefix, apis.DeprovisionServiceInstance).Methods("DELETE")

	//binding

	http.Handle("/", router)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Error(err, "failed to initialize a server")
	}
}
