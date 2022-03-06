package domain

import (
	"time"

	"github.com/uptrace/bun"
)

type News struct {
	bun.BaseModel `bun:"table:news,alias:n"`

	ID     int64     `json:"-" bun:"id,pk,autoincrement"`
	Title  string    `json:"title" bun:"title,notnull,unique"`
	Body   string    `json:"body" bun:"body,notnull"`
	Status string    `json:"status" bun:"status,notnull"`
	Topics []*Topics `json:"-" bun:"rel:has-many,join:id=news_id"`

	CreatedAt time.Time `json:"-" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `json:"-" bun:",nullzero,notnull,default:current_timestamp"`
}
