package managers

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	pb "voting-service/proto-service/genprotos"
)

type PublicVotesManager struct {
	Conn *sql.DB
}

func NewPublicVotesManager(db *sql.DB) *PublicVotesManager {
	return &PublicVotesManager{
		Conn: db,
	}
}
func (ps *PublicVotesManager) CreatePublicVotes(pv *pb.PublicVoteCreate, id *string, id2 *string) (*pb.Void, error) {
	tx, err := ps.Conn.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	query := `INSERT INTO public_vote(id, election_id, public_id) VALUES($1, $2, $3)`
	_, err = tx.Exec(query, id, pv.ElectionId, pv.PublicId)
	if err != nil {
		return nil, err
	}

	query2 := `INSERT INTO vote(id, election_id, candidate_id) VALUES($1, $2, $3)`
	_, err = tx.Exec(query2, id2, pv.ElectionId, pv.CandidateId)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}
func (ps *PublicVotesManager) DeletePublicVotes(pv *pb.ById) (*pb.Void, error) {
	query := `delete from public_vote where id = $1`
	_, err := ps.Conn.Exec(query, pv.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
func (ps *PublicVotesManager) UpdatePublicVotes(pv *pb.PublicVote) (*pb.Void, error) {
	query := `update public_vote set `
	var conditions []string
	var args []interface{}
	if pv.ElectionId != "" {
		conditions = append(conditions, fmt.Sprintf("election_id = $%d", len(args)+1))
		args = append(args, pv.ElectionId)
	}
	if pv.PublicId != "" {
		conditions = append(conditions, fmt.Sprintf("public_id = $%d", len(args)+1))
		args = append(args, pv.PublicId)
	}
	query += strings.Join(conditions, ", ")
	query += fmt.Sprintf("where id = $%d", len(args)+1)
	args = append(args, pv.Id)
	_, err := ps.Conn.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
func (ps *PublicVotesManager) GetByIdPublicVote(id *pb.ById) (*pb.PublicVote, error) {
	query := `select 
        id, 
        election_id, 
        public_id
    from 
        public_vote 
    where 
        id = $1`
	row := ps.Conn.QueryRow(query, id.Id)
	var pv pb.PublicVote
	err := row.Scan(
		&pv.Id,
		&pv.ElectionId,
		&pv.PublicId)
	if err != nil {
		return nil, err
	}
	return &pv, nil
}
func (ps *PublicVotesManager) GetAllPublucVotes(filter *pb.Filter) (*pb.GetAllPublicVote, error) {
	query := `select 
        id, 
        election_id, 
        public_id
    from 
        public_vote`
	var conditions []string
	var args []interface{}
	if filter.Election != "" {
		conditions = append(conditions, fmt.Sprintf("election_id = $%d", len(args)+1))
		args = append(args, filter.Election)
	}
	if len(conditions) > 0 {
		query = query + " WHERE " + strings.Join(conditions, " AND ")
	}
	rows, err := ps.Conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out pb.GetAllPublicVote
	for rows.Next() {
		var pv pb.PublicVote
		err := rows.Scan(
			&pv.Id,
			&pv.ElectionId,
			&pv.PublicId)
		if err != nil {
			return nil, err
		}
		out.PublicVotes = append(out.PublicVotes, &pv)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &out, nil
}
func (ps *PublicVotesManager) FindWinner(we *pb.WhichElection) (*pb.Winner, error) {
	query := `
		SELECT election_id, candidate_id, COUNT(*) as vote_count
		FROM votes
		WHERE election_id = $1
		GROUP BY election_id, candidate_id
		ORDER BY vote_count DESC
		LIMIT 1`

	row := ps.Conn.QueryRow(query, we.ElectionId)

	var winner pb.Winner
	err := row.Scan(&winner.ElectionId, &winner.CandidateId, &winner.Votes)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("no votes found for the given election")
		}
		return nil, err
	}
	return &winner, nil
}
