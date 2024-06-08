package managers

import (
	"context"
	"database/sql"
	"fmt"
	pb "public-service/proto-service/genprotos"
)

type PartyManager struct {
	conn *sql.DB
}

func NewPartyManager(db *sql.DB) *PartyManager {
	return &PartyManager{conn: db}
}

func (m *PartyManager) Create(context context.Context, party *pb.PartyCreate) error {
	query :=
		`
		INSERT INTO party (id, name , slogan, opened_date, description) 
		VALUES ($1, $2, $3, $4, $5)
		`
	_, err := m.conn.Exec(query, party.Id, party.Name, party.Slogan, party.OpenedDate, party.Description)
	if err != nil {
		return err
	}
	return nil
}

func (m *PartyManager) Update(context context.Context, party *pb.PartyUpdate) error {
	query :=
		`
		UPDATE party SET name = $1, slogan = $2, opened_date = $3, description = $4 WHERE id = $5
		`
	_, err := m.conn.Exec(query, party.Name, party.Slogan, party.OpenedDate, party.Description, party.Id)
	if err != nil {
		panic(err)
	}
	return nil
}

func (m *PartyManager) Delete(context context.Context, party *pb.PartyDelete) error {
	query :=
		`
		UPDATE party SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1
		`
	if _, err := m.conn.Exec(query, party.Id); err != nil {
		panic(err)
	}
	return nil
}

func (m *PartyManager) GetById(context context.Context, party *pb.PartyGetById) (*pb.Party, error) {
	query :=
		`
		SELECT id, name, slogan, opened_date, description, created_at, updated_at, deleted_at
		FROM party WHERE id = $1 AND deleted_at = 0
		`
	res := &pb.Party{}
	err := m.conn.QueryRow(
		query, party.Id,
	).Scan(
		&res.Id, &res.Name, &res.Slogan, &res.OpenedDate, &res.Description, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *PartyManager) GetAll(context context.Context, party *pb.PartyGetAllReq) (*pb.PartyGetAllResp, error) {
	query :=
		`
		SELECT id, name, slogan, opened_date, description, created_at, updated_at, deleted_at
		FROM party WHERE deleted_at = 0
		`
	res := &pb.PartyGetAllResp{}

	var args []interface{}
	paramIndex := 1
	if party.Name != "" {
		query += fmt.Sprintf(" AND name = $%d", paramIndex)
		args = append(args, party.Name)
		paramIndex++
	}
	if party.OpenedDate != "" {
		query += fmt.Sprintf(" AND opened_date = $%d", paramIndex)
		args = append(args, party.OpenedDate)
		paramIndex++
	}
	rows, err := m.conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		r := &pb.Party{}
		rows.Scan(
			&r.Id, &r.Name, &r.Slogan, &r.OpenedDate, &r.Description, &r.CreatedAt, &r.UpdatedAt, &r.DeletedAt,
		)
		res.Parties = append(res.Parties, r)
	}
	return res, nil
}
