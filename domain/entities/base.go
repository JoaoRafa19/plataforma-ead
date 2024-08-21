package entities

import (
	"time"

	"github.com/google/uuid"
)

type BaseEntity struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdateAt  time.Time
}
