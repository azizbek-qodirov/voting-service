package postgresql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"voting-service/config"
	"voting-service/db/postgresql/managers"
)

type Storage struct {
	DB          *sql.DB
	CandidateS  CandidateI
	ElectionS   ElectionI
	PublicVoteS PublicVoteI
}

func (s *Storage) Election() ElectionI {
	if s.ElectionS == nil {
		s.ElectionS = &managers.ElectionManager{Conn: s.DB}
	}
	return s.ElectionS
}
func (s *Storage) Candidate() CandidateI {
	if s.CandidateS == nil {
		s.CandidateS = &managers.CandidateManager{Conn: s.DB}
	}
	return s.CandidateS
}
func (s *Storage) PublicVote() PublicVoteI {
	if s.PublicVoteS == nil {
		s.PublicVoteS = &managers.PublicVotesManager{Conn: s.DB}
	}
	return s.PublicVoteS
}

func ConnectDB(cf *config.Config) (*Storage, error) {
	dbConn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cf.DB_USER, cf.DB_PASSWORD, cf.DB_HOST, cf.DB_PORT, cf.DB_NAME)
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	e_repo := managers.NewElectionManager(db)
	c_repo := managers.NewCandidateManager(db)
	pv_repo := managers.NewPublicVotesManager(db)
	return &Storage{
		DB:          db,
		CandidateS:  c_repo,
		ElectionS:   e_repo,
		PublicVoteS: pv_repo,
	}, nil
}
