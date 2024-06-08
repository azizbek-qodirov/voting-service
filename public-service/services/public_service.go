package service

import (
	"context"
	"public-service/db/postgresql"
	pb "public-service/proto-service/genprotos"
)

type PublicService struct {
	stg *postgresql.Storage
	pb.UnimplementedPublicServiceServer
}

func NewPublicService(stg *postgresql.Storage) *PublicService {
	return &PublicService{stg: stg}
}

func (s *PublicService) Create(context context.Context, party *pb.PublicCreate) (*pb.Void, error) {
	return nil, s.stg.PublicI.Create(context, party)
}

func (s *PublicService) Update(context context.Context, party *pb.PublicUpdate) (*pb.Void, error) {
	return nil, s.stg.PublicI.Update(context, party)
}

func (s *PublicService) Delete(context context.Context, party *pb.PublicDelete) (*pb.Void, error) {
	return nil, s.stg.PublicI.Delete(context, party)
}

func (s *PublicService) GetById(context context.Context, party *pb.PublicGetById) (*pb.Public, error) {
	return s.stg.PublicI.GetById(context, party)
}

func (s *PublicService) GetAll(context context.Context, party *pb.PublicGetAllReq) (*pb.PublicGetAllResp, error) {
	return s.stg.PublicI.GetAll(context, party)
}
