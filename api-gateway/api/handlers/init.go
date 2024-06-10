package handlers

import (
	"api-gateway/config/logger"
	pb "api-gateway/proto-service/genprotos"

	"google.golang.org/grpc"
)

type HTTPHandler struct {
	Public     pb.PublicServiceClient
	Party      pb.PartyServiceClient
	Candidate  pb.CandidateServiceClient
	Election   pb.ElectionServiceClient
	PublicVote pb.PublicVoteServiceClient
	Logger     logger.Logger
}

func NewHandler(connP, connV *grpc.ClientConn, logger logger.Logger) *HTTPHandler {
	return &HTTPHandler{
		Public:     pb.NewPublicServiceClient(connP),
		Party:      pb.NewPartyServiceClient(connP),
		Candidate:  pb.NewCandidateServiceClient(connV),
		Election:   pb.NewElectionServiceClient(connV),
		PublicVote: pb.NewPublicVoteServiceClient(connV),
		Logger:     logger,
	}
}
