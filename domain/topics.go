package domain

import (
	"time"

	"github.com/uptrace/bun"
)

type Topics struct {
	bun.BaseModel `bun:"table:topics,alias:n"`

	ID     int64  `bun:"id,pk,autoincrement"`
	Name   string `bun:"name,notnull"`
	NewsID int64  `bun:"news_id,notnull"`
	News   *News  `bun:"rel:belongs-to,join=news_id=id"`

	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}
