package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manjdk/Carbon-Based-Life-Forms/config"
	"github.com/manjdk/Carbon-Based-Life-Forms/domain/usecase"
	"github.com/manjdk/Carbon-Based-Life-Forms/factory/controller"
	"github.com/manjdk/Carbon-Based-Life-Forms/log"
	"github.com/manjdk/Carbon-Based-Life-Forms/queue"
	"github.com/manjdk/Carbon-Based-Life-Forms/repository"
)

func main() {
	cfg, err := config.NewConfig("config")
	if err != nil {
		log.FatalZ(err).Msg("Failed to initiate factory config")
	}

	db, err := repository.NewDynamoDB(cfg)
	if err != nil {
		log.FatalZ(err).Msg("Failed to create database")
	}

	factorySQSClient, err := queue.NewSQSClient(
		cfg.QueueFactory.Name,
		cfg.QueueFactory.Host,
		cfg.AwsRegion,
	)
	if err != nil {
		log.FatalZ(err).Msg("Failed to create factory SQS")
	}

	createMineralUC := usecase.NewCreateMineralUC(db)
	getMineralByIDUC := usecase.NewGetMineralUC(db)
	deleteMineralUC := usecase.NewDeleteMineralUC(db)
	getAllMineralsUC := usecase.NewGetAllMineralUC(db, db)
	meltMineralUC := usecase.NewMeltMineralUC(db, db)
	condenseMineralUC := usecase.NewCondenseMineralUC(db, db)
	fractureMineralUC := usecase.NewFractureMineralUC(db, db)

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/minerals", controller.CreateMineralFactory(createMineralUC)).Methods(http.MethodPost)
	router.HandleFunc("/minerals", controller.GetMineralsFactory(getAllMineralsUC)).Methods(http.MethodGet)
	router.HandleFunc("/minerals/{mineralId}", controller.GetMineralFactory(getMineralByIDUC)).Methods(http.MethodGet)
	router.HandleFunc("/minerals/{mineralId}", controller.DeleteMineralFactory(deleteMineralUC)).Methods(http.MethodDelete)

	go factorySQSClient.RunFactoryConsumer(
		meltMineralUC,
		condenseMineralUC,
		fractureMineralUC,
	)

	log.InfoZ("NoTraceID").Msg("Factory is running")

	if err := http.ListenAndServe(":8383", router); err != nil {
		log.FatalZ(err).Msg("Failed to run factory client")
	}
}
