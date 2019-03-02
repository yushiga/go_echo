package form

import "time"

type User struct {
	ID        uint       `validate:"required"`
	UserID    string     `validate:"-"`
	Name      string     `validate:"-"`
	Remark    string     `validate:"-"`
	CreatedAt time.Time  `validate:"-"`
	UpdatedAt time.Time  `validate:"-"`
	DeletedAt *time.Time `validate:"-"`
}
