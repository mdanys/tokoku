package cart

type Core struct {
	ID        uint
	UserID    uint
	Qty       uint
	ProductID uint
	Product   ProductCore
}

type ProductCore struct {
	ID     uint
	Name   string
	Price  uint
	Qty    uint
	Detail string
	Image  string
}

type Repository interface {
	Insert(data Core) (Core, error)
	Edit(data Core) (Core, error)
	Remove(id uint) error
	Get(id uint) ([]Core, error)
}

type Service interface {
	Create(data Core) (Core, error)
	Update(data Core) (Core, error)
	Delete(id uint) error
	Show(id uint) ([]Core, error)
}
