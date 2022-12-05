package user

type Core struct {
	ID       uint
	Name     string
	Phone    string
	Email    string
	Password string
	Product  []ProductCore
}

type ProductCore struct {
	ID     uint
	UserID uint
	Name   string
	Image  string
	Qty    uint
	Price  uint
}

type Repository interface {
	Insert(data Core) (Core, error)
	Edit(data Core, id uint) (Core, error)
	Remove(id uint) error
	GetMyProfile(token uint) (Core, error)
	GetByID(id uint) (Core, error)
}

type Service interface {
	Create(data Core) (Core, error)
	Update(data Core, id uint) (Core, error)
	Delete(id uint) error
	ShowMyProfile(token uint) (Core, error)
	GetByID(id uint) (Core, error)
}
