package models

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

var _ bun.BeforeAppendModelHook = (*Post)(nil)

type Post struct {
	bun.BaseModel `bun:"table:posts,alias:p"`

	ID       int64  `bun:",pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Content  string `json:"content"`
	AuthorId uuid.UUID

	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

func (p *Post) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		p.CreatedAt = time.Now()
	case *bun.UpdateQuery:
		p.UpdatedAt = time.Now()
	}
	return nil
}
