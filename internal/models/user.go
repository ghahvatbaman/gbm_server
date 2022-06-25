package models

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

var _ bun.BeforeAppendModelHook = (*User)(nil)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID    uuid.UUID `bun:",pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email" bun:",unique"`
	Phone string    `json:"phone"`
	Role  string    `json:"-"`
	Posts []*Post   `bun:"rel:has-many,join:id=author_id"`

	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

func (u *User) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		u.CreatedAt = time.Now()
	case *bun.UpdateQuery:
		u.UpdatedAt = time.Now()
	}
	return nil
}
