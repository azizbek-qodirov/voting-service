package service

import (
	"context"
	"public-service/db/postgresql"
	pb "public-service/proto-service/genprotos"
)

type PartyService struct {
	stg *postgresql.Storage
	pb.UnimplementedPartyServiceServer
}

func NewPartyService(stg *postgresql.Storage) *PartyService {
	return &PartyService{stg: stg}
}

func (s *PartyService) Create(context context.Context, party *pb.PartyCreate) (*pb.Void, error) {
	return nil, s.stg.PartyI.Create(context, party)
}

func (s *PartyService) Update(context context.Context, party *pb.PartyUpdate) (*pb.Void, error) {
	return nil, s.stg.PartyI.Update(context, party)
}

func (s *PartyService) Delete(context context.Context, party *pb.PartyDelete) (*pb.Void, error) {
	return nil, s.stg.PartyI.Delete(context, party)
}

func (s *PartyService) GetById(context context.Context, party *pb.PartyGetById) (*pb.Party, error) {
	return s.stg.PartyI.GetById(context, party)
}

func (s *PartyService) GetAll(context context.Context, party *pb.PartyGetAllReq) (*pb.PartyGetAllResp, error) {
	return s.stg.PartyI.GetAll(context, party)
}
