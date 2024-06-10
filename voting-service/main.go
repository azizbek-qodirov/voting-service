package main

import (
	"log"
	"net"
	"path/filepath"
	"runtime"

	cf "voting-service/config"
	"voting-service/config/logger"
	"voting-service/db/postgresql"
	pb "voting-service/proto-service/genprotos"
	service "voting-service/services"

	"google.golang.org/grpc"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func main() {
	config := cf.Load()
	logger := logger.NewLogger(basepath, config.LOG_PATH) // Don't forget to change your log path
	em := cf.NewErrorManager(logger)
	db, err := postgresql.ConnectDB(&config)
	em.CheckErr(err)
	defer db.DB.Close()

	listener, err := net.Listen("tcp", config.TCP_PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCandidateServiceServer(s, service.NewCandidateService(db))
	pb.RegisterElectionServiceServer(s, service.NewElectionService(db))
	pb.RegisterPublicVoteServiceServer(s, service.NewPublicVoteService(db))

	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
