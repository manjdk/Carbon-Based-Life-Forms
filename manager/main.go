package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manjdk/Carbon-Based-Life-Forms/config"
	"github.com/manjdk/Carbon-Based-Life-Forms/custom_http"
	"github.com/manjdk/Carbon-Based-Life-Forms/log"
	"github.com/manjdk/Carbon-Based-Life-Forms/manager/controller"
	"github.com/manjdk/Carbon-Based-Life-Forms/queue"
)

func main() {
	cfg, err := config.NewConfig("config")
	if err != nil {
		log.FatalZ(err).Msg("Failed to initiate manager config")
	}

	factorySQSClient, err := queue.NewSQSClient(
		cfg.QueueFactory.Name,
		cfg.QueueFactory.Host,
		cfg.AwsRegion,
	)
	if err != nil {
		log.FatalZ(err).Msg("Failed to create factory SQS client in manager")
	}

	managerSQSClient, err := queue.NewSQSClient(
		cfg.QueueManager.Name,
		cfg.QueueManager.Host,
		cfg.AwsRegion,
	)
	if err != nil {
		log.FatalZ(err).Msg("Failed to create manager SQS client in manager")
	}

	httpClient := custom_http.NewHttpClient(http.DefaultClient)

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/minerals", controller.CreateMineralManager(httpClient, cfg.FactoryURL)).Methods(http.MethodPost)
	router.HandleFunc("/minerals", controller.GetMineralsManager(httpClient, cfg.FactoryURL)).Methods(http.MethodGet)
	router.HandleFunc("/minerals/{mineralId}", controller.GetMineralManager(httpClient, cfg.FactoryURL)).Methods(http.MethodGet)
	router.HandleFunc("/minerals/{mineralId}", controller.DeleteMineralManager(httpClient, cfg.FactoryURL)).Methods(http.MethodDelete)

	go managerSQSClient.RunManagerConsumer(factorySQSClient)

	log.InfoZ("NoTraceID").Msg("Manager is running")

	if err := http.ListenAndServe(":8282", router); err != nil {
		log.FatalZ(err).Msg("Failed to run manager client")
	}
}
