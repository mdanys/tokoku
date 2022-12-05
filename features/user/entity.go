package user

import "time"

type Core struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Repository interface {
	Insert(data Core) (Core, error)
	Edit(data Core, id uint) (Core, error)
	Remove(id uint) error
	Login(data Core) (Core, error)
}

type Service interface {
	Create(data Core) (Core, error)
	Update(data Core, id uint) (Core, error)
	Delete(id uint) error
	Login(data Core) (Core, error)
}
