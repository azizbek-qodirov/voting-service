package postgresql

import (
	"context"
	pb "public-service/proto-service/genprotos"
)

type StorageI interface {
	Party() PartyI
	Public() PublicI
}
type PartyI interface {
	Create(ctx context.Context, partyReq *pb.PartyCreate) error
	Update(ctx context.Context, partyReq *pb.PartyUpdate) error
	Delete(ctx context.Context, partyReq *pb.PartyDelete) error
	GetById(ctx context.Context, partyReq *pb.PartyGetById) (*pb.Party, error)
	GetAll(ctx context.Context, partyReq *pb.PartyGetAllReq) (*pb.PartyGetAllResp, error)
}

type PublicI interface {
	Create(ctx context.Context, publicReq *pb.PublicCreate) error
	Update(ctx context.Context, publicReq *pb.PublicUpdate) error
	Delete(ctx context.Context, publicReq *pb.PublicDelete) error
	GetById(ctx context.Context, publicReq *pb.PublicGetById) (*pb.Public, error)
	GetAll(ctx context.Context, publicReq *pb.PublicGetAllReq) (*pb.PublicGetAllResp, error)
}
