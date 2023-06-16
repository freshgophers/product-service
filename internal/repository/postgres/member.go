package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	"product-service/internal/domain/member"
	"product-service/pkg/store"
)

type MemberRepository struct {
	db *sqlx.DB
}

func NewMemberRepository(db *sqlx.DB) *MemberRepository {
	return &MemberRepository{
		db: db,
	}
}

func (s *MemberRepository) Select(ctx context.Context) (dest []member.Entity, err error) {
	query := `
		SELECT id, full_name, books
		FROM members
		ORDER BY id`

	dest = make([]member.Entity, 0)
	err = s.db.SelectContext(ctx, &dest, query)

	return
}

func (s *MemberRepository) Create(ctx context.Context, data member.Entity) (id string, err error) {
	query := `
		INSERT INTO members (full_name, books)
		VALUES ($1, $2)
		RETURNING id`

	args := []any{data.FullName, data.Books.String()}

	err = s.db.QueryRowContext(ctx, query, args...).Scan(&id)

	return
}

func (s *MemberRepository) Get(ctx context.Context, id string) (dest member.Entity, err error) {
	query := `
		SELECT id, full_name, books
		FROM members
		WHERE id=$1`

	args := []any{id}

	if err = s.db.GetContext(ctx, &dest, query, args...); err != nil && err != sql.ErrNoRows {
		return
	}

	if err == sql.ErrNoRows {
		err = store.ErrorNotFound
	}

	return
}

func (s *MemberRepository) Update(ctx context.Context, id string, data member.Entity) (err error) {
	sets, args := s.prepareArgs(data)
	if len(args) > 0 {

		args = append(args, id)
		sets = append(sets, "updated_at=CURRENT_TIMESTAMP")

		query := fmt.Sprintf("UPDATE members SET %s WHERE id=$%d", strings.Join(sets, ", "), len(args))
		_, err = s.db.ExecContext(ctx, query, args...)
	}

	return
}

func (s *MemberRepository) prepareArgs(data member.Entity) (sets []string, args []any) {
	if data.FullName != nil {
		args = append(args, data.FullName)
		sets = append(sets, fmt.Sprintf("full_name=$%d", len(args)))
	}

	if len(data.Books) > 0 {
		args = append(args, data.Books.String())
		sets = append(sets, fmt.Sprintf("books=$%d", len(args)))
	}

	return
}

func (s *MemberRepository) Delete(ctx context.Context, id string) (err error) {
	query := `
		DELETE 
		FROM members
		WHERE id=$1`

	args := []any{id}

	_, err = s.db.ExecContext(ctx, query, args...)

	return
}
