package entity

import (
	"tanyaustadz/pkg/common"
	"time"
)

type DataObject struct {
	Id        common.ID
	Path      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
