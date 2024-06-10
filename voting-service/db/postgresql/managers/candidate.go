package managers

import (
	"database/sql"
	"fmt"
	"strings"

	pb "voting-service/proto-service/genprotos"
)

type CandidateManager struct {
	Conn *sql.DB
}

func NewCandidateManager(db *sql.DB) *CandidateManager {
	return &CandidateManager{Conn: db}
}
func (cs *CandidateManager) CreateCandidate(c *pb.CandidateCreate, id *string) (*pb.Void, error) {
	query := `insert into candidate(
		id,
		election_id,
		public_id,
		party_id) values($1, $2, $3, $4)`
	_, err := cs.Conn.Exec(query, id, c.ElectionId, c.PublicId, c.PartyId)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (cs *CandidateManager) DeleteCandidate(c *pb.ById) (*pb.Void, error) {
	query := `delete from candidate where id = $1`
	_, err := cs.Conn.Exec(query, c.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (cs *CandidateManager) UpdateCandidate(c *pb.Candidate) (*pb.Void, error) {
	query := `update candidate set election_id = $1, party_id = $2, public_id = $3 where id = $4`
	_, err := cs.Conn.Exec(query, c.Election, c.Party, c.Public, c.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (cs *CandidateManager) GetByIdCandidate(c *pb.ById) (*pb.Candidate, error) {
	query := `select id, election_id, public_id, party_id from candidate where id = $1`
	row := cs.Conn.QueryRow(query, c.Id)
	var candidate pb.Candidate
	err := row.Scan(&candidate.Id, &candidate.Election, &candidate.Party, &candidate.Public)
	if err != nil {
		return nil, err
	}
	return &candidate, nil
}

func (cs *CandidateManager) GetAllCandidates(filter *pb.Filter) (*pb.GetAllCandidate, error) {
	query := `select 
        id, 
        election_id, 
        public_id, 
        party_id
    from 
        candidate`
	var conditions []string
	var args []interface{}
	if filter.Party != "" {
		conditions = append(conditions, fmt.Sprintf("party_id = $%d", len(args)+1))
		args = append(args, filter.Party)
	}
	if len(conditions) > 0 {
		query = query + " WHERE " + strings.Join(conditions, " AND ")
	}
	rows, err := cs.Conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out pb.GetAllCandidate
	for rows.Next() {
		var candidate pb.Candidate
		err := rows.Scan(
			&candidate.Id,
			&candidate.Election,
			&candidate.Party,
			&candidate.Public)
		if err != nil {
			return nil, err
		}
		out.Candidates = append(out.Candidates, &candidate)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &out, nil
}
