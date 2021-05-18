package repository

import (
	"log"
	"os"
	"testing"

	"github.com/manjdk/Carbon-Based-Life-Forms/config"
)

var testDynamoDB *DynamoDB

func TestMain(m *testing.M) {
	testConfig, err := config.NewConfig("../config")
	if err != nil {
		log.Panicf("Failed to load test config: %s", err.Error())
	}

	if testDynamoDB, err = NewDynamoDB(testConfig); err != nil {
		panic(err)
	}

	retCode := m.Run()
	os.Exit(retCode)
}
