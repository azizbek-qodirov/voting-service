package main

import (
	"api-gateway/api"
	cf "api-gateway/config"
	"api-gateway/config/logger"
	"fmt"
	"path/filepath"
	"runtime"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func main() {
	config := cf.Load()
	logger := logger.NewLogger(basepath, config.LOG_PATH) // Don't forget to change your log path
	em := cf.NewErrorManager(logger)

	PublicConn, err := grpc.NewClient(fmt.Sprintf("localhost%s", config.PUBLIC_SERVICE_PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	em.CheckErr(err)
	defer PublicConn.Close()
	VotingConn, err := grpc.NewClient(fmt.Sprintf("localhost%s", config.VOTING_SERVICE_PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	em.CheckErr(err)
	defer VotingConn.Close()

	r := api.NewGin(PublicConn, VotingConn, logger)

	fmt.Printf("Server started on port %s\n", config.HTTPPort)
	logger.INFO.Println("Server started on port: " + config.HTTPPort)
	if r.Run(config.HTTPPort); err != nil {
		logger.ERROR.Panicln("Handling stopped due to error " + err.Error())
	}
}
