package services

import (
	"context"

	postgres "voting-service/db/postgresql"
	pb "voting-service/proto-service/genprotos"

	"github.com/google/uuid"
)

type CandidateService struct {
	stg *postgres.Storage
	pb.UnimplementedCandidateServiceServer
}

func NewCandidateService(stg *postgres.Storage) *CandidateService {
	return &CandidateService{stg: stg}
}
func (cs *CandidateService) CreateCandidate(ctx context.Context, c *pb.CandidateCreate) (*pb.Void, error) {
	id := uuid.NewString()
	return cs.stg.CandidateS.CreateCandidate(c, &id)
}

func (cs *CandidateService) DeleteCandidate(ctx context.Context, id *pb.ById) (*pb.Void, error) {
	return cs.stg.Candidate().DeleteCandidate(id)
}

func (cs *CandidateService) UpdateCandidate(ctx context.Context, c *pb.Candidate) (*pb.Void, error) {
	return cs.stg.Candidate().UpdateCandidate(c)
}

func (cs *CandidateService) GetByIdCandidate(ctx context.Context, id *pb.ById) (*pb.Candidate, error) {
	return cs.stg.Candidate().GetByIdCandidate(id)
}

func (cs *CandidateService) GetAllCandidates(ctx context.Context, f *pb.Filter) (*pb.GetAllCandidate, error) {
	return cs.stg.Candidate().GetAllCandidates(f)
}
