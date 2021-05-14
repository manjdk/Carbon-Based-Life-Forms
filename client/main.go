package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manjdk/Carbon-Based-Life-Forms/client/controller"
	"github.com/manjdk/Carbon-Based-Life-Forms/config"
	"github.com/manjdk/Carbon-Based-Life-Forms/custom_http"
	"github.com/manjdk/Carbon-Based-Life-Forms/log"
	"github.com/manjdk/Carbon-Based-Life-Forms/queue"
)

func main() {
	cfg, err := config.NewConfig("config")
	if err != nil {
		log.FatalZ(err).Msg("Failed to initiate app config")
	}

	factorySQSClient, err := queue.NewSQSClient(
		cfg.QueueFactory.Name,
		cfg.QueueFactory.Host,
		cfg.AwsRegion,
	)
	if err != nil {
		log.FatalZ(err).Msg("Failed to create factory SQS")
	}

	httpClient := custom_http.NewHttpClient(http.DefaultClient)

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/minerals", controller.CreateMineral(httpClient, cfg.ManagerURL)).Methods(http.MethodPost)
	router.HandleFunc("/minerals", controller.GetMinerals(httpClient, cfg.ManagerURL)).Methods(http.MethodGet)
	router.HandleFunc("/minerals/{mineralId}", controller.GetMineral(httpClient, cfg.ManagerURL)).Methods(http.MethodGet)
	router.HandleFunc("/minerals", controller.UpdateMineral(factorySQSClient)).Methods(http.MethodPut)

	log.InfoZ("noTraceID").Msg("App is running")

	if err := http.ListenAndServe(":8181", router); err != nil {
		log.FatalZ(err).Msg("Failed to run app client")
	}
}
