package domain

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type UserIT struct {
	ID        string    `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	NIP       int       `db:"nip"`
	Name      string    `db:"name"`
	Password  string    `db:"password"`
	Role      string    `db:"role"`
}

type RegisterUserIT struct {
	NIP      *int   `json:"nip" validate:"required,nip=it"`
	Name     string `json:"name" validate:"required,min=5,max=50"`
	Password string `json:"password" validate:"required,min=5,max=33"`
}

type LoginUserIT struct {
	NIP      *int   `json:"nip" validate:"required,nip=it"`
	Password string `json:"password" validate:"required,min=5,max=33"`
}

type GetUserIT struct {
	ID        string    `json:"userId"`
	NIP       int       `json:"nip"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

func (u RegisterUserIT) NewUserITFromDTO() UserIT {
	id := ulid.Make()
	rawCreatedAt := time.Now().Format(time.RFC3339)
	createdAt, _ := time.Parse(time.RFC3339, rawCreatedAt)

	return UserIT{
		ID:        id.String(),
		CreatedAt: createdAt,
		NIP:       *u.NIP,
		Name:      u.Name,
		Password:  u.Password,
		Role:      "it",
	}
}
