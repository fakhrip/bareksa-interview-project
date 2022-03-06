package domain

import (
	"time"

	"github.com/uptrace/bun"
)

type Topics struct {
	bun.BaseModel `bun:"table:topics,alias:n"`

	ID     int64  `json:"id" bun:"id,pk,autoincrement"`
	Name   string `json:"name" bun:"name,notnull"`
	NewsID int64  `json:"newsId" bun:"news_id,notnull"`
	News   News   `json:"-" bun:"rel:belongs-to,join:news_id=id"`

	CreatedAt time.Time `json:"-" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `json:"-" bun:",nullzero,notnull,default:current_timestamp"`
}
