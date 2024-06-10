package postgresql

import (
	pb "voting-service/proto-service/genprotos"
)

type StorageI interface {
	Election() ElectionI
	Candidate() CandidateI
	PublicVote() PublicVoteI
}

type ElectionI interface {
	CreateElection(el *pb.Election) (*pb.Void, error)
	DeleteElection(id *pb.ById) (*pb.Void, error)
	UpdateElection(el *pb.Election) (*pb.Void, error)
	GetByIdElection(id *pb.ById) (*pb.Election, error)
	GetAllElections(filter *pb.Filter) (*pb.GetAllElection, error)
}

type CandidateI interface {
	CreateCandidate(c *pb.CandidateCreate, id *string) (*pb.Void, error)
	DeleteCandidate(c *pb.ById) (*pb.Void, error)
	UpdateCandidate(c *pb.Candidate) (*pb.Void, error)
	GetByIdCandidate(c *pb.ById) (*pb.Candidate, error)
	GetAllCandidates(filter *pb.Filter) (*pb.GetAllCandidate, error)
}

type PublicVoteI interface {
	CreatePublicVotes(pv *pb.PublicVoteCreate, id *string, id2 *string) (*pb.Void, error)
	DeletePublicVotes(pv *pb.ById) (*pb.Void, error)
	UpdatePublicVotes(pv *pb.PublicVote) (*pb.Void, error)
	GetByIdPublicVote(id *pb.ById) (*pb.PublicVote, error)
	GetAllPublucVotes(filter *pb.Filter) (*pb.GetAllPublicVote, error)
	FindWinner(we *pb.WhichElection) (*pb.Winner, error)
}
