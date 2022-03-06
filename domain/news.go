package domain

import (
	"time"

	"github.com/uptrace/bun"
)

type News struct {
	bun.BaseModel `bun:"table:news,alias:n"`

	ID     int64     `bun:"id,pk,autoincrement"`
	Title  string    `bun:"title,notnull"`
	Body   string    `bun:"body,notnull"`
	Status string    `bun:"status,notnull"`
	Topics []*Topics `bun:"rel:has-many,join:id=topic_id"`

	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}
