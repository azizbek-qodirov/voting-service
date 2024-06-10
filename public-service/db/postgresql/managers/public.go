package managers

import (
	"context"
	"database/sql"
	"fmt"
	pb "public-service/proto-service/genprotos"
)

type PublicManager struct {
	conn *sql.DB
}

func NewPublicManager(conn *sql.DB) *PublicManager {
	return &PublicManager{conn: conn}
}

func (m *PublicManager) Create(context context.Context, public *pb.PublicCreate) error {
	query :=
		`
		INSERT INTO public (id, name, last_name, phone, email, birthday, gender, party_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		`
	_, err := m.conn.Exec(query, public.Id, public.Name, public.LastName, public.Phone, public.Email, public.Birthday, public.Gender, public.PartyId)
	if err != nil {
		return err
	}
	return nil
}

func (m *PublicManager) Update(context context.Context, public *pb.PublicUpdate) error {
	query :=
		`
		UPDATE public
		SET name = $1, last_name = $2, phone = $3, email = $4, birthday = $5, gender = $6, party_id = $7
		WHERE id = $8
		`
	_, err := m.conn.Exec(query, public.Name, public.LastName, public.Phone, public.Email, public.Birthday, public.Gender, public.PartyId, public.Id)
	if err != nil {
		return err
	}
	return nil
}

func (m *PublicManager) Delete(context context.Context, public *pb.PublicDelete) error {
	query :=
		`
		UPDATE public SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1
		`
	_, err := m.conn.Exec(query, public.Id)
	if err != nil {
		return err
	}
	return nil
}

func (m *PublicManager) GetById(context context.Context, public *pb.PublicGetById) (*pb.Public, error) {
	query :=
		`
		SELECT id, name, last_name, phone, email, birthday, gender, party_id, created_at, updated_at, deleted_at 
		FROM public WHERE id = $1 AND deleted_at = 0
		`
	res := &pb.Public{}
	err := m.conn.QueryRow(query, public.Id).Scan(&res.Id, &res.Name, &res.LastName, &res.Phone, &res.Email, &res.Birthday, &res.Gender, &res.PartyId, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *PublicManager) GetAll(context context.Context, public *pb.PublicGetAllReq) (*pb.PublicGetAllResp, error) {
	query :=
		`
		SELECT id, name, last_name, phone, email, birthday, gender, party_id, created_at, updated_at, deleted_at 
		FROM public WHERE deleted_at = 0
		`
	var args []interface{}
	paramIndex := 1

	if public.PartyId != "" {
		query += fmt.Sprintf(" AND party_id = $%d", paramIndex)
		args = append(args, public.PartyId)
		paramIndex++
	}
	if public.Name != "" {
		query += fmt.Sprintf(" AND name = $%d", paramIndex)
		args = append(args, public.Name)
		paramIndex++
	}
	if public.LastName != "" {
		query += fmt.Sprintf(" AND last_name = $%d", paramIndex)
		args = append(args, public.LastName)
		paramIndex++
	}
	if public.Phone != "" {
		query += fmt.Sprintf(" AND phone = $%d", paramIndex)
		args = append(args, public.Phone)
		paramIndex++
	}
	if public.Email != "" {
		query += fmt.Sprintf(" AND email = $%d", paramIndex)
		args = append(args, public.Email)
		paramIndex++
	}
	if public.Birthday != "" {
		query += fmt.Sprintf(" AND birthday = $%d", paramIndex)
		args = append(args, public.Birthday)
		paramIndex++
	}
	if public.Gender != "" {
		query += fmt.Sprintf(" AND gender = $%d", paramIndex)
		args = append(args, public.Gender)
		paramIndex++
	}
	rows, err := m.conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	res := &pb.PublicGetAllResp{}
	for rows.Next() {
		r := &pb.Public{}
		err := rows.Scan(&r.Id, &r.Name, &r.LastName, &r.Phone, &r.Email, &r.Birthday, &r.Gender, &r.PartyId, &r.CreatedAt, &r.UpdatedAt, &r.DeletedAt)
		if err != nil {
			return nil, err
		}
		res.Publics = append(res.Publics, r)
	}
	return res, nil
}
