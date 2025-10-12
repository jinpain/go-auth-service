package user

import (
	"time"

	"github.com/google/uuid"
)

type Model struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"creation_ts"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	Verified  bool      `json:"verified"`
	Blocked   bool      `json:"blocked"`
}
