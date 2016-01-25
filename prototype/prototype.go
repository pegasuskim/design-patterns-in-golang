package prototype

type producter interface {
	clone() producter
	GetName() string
}

type Manager struct {
	product producter
}

func (self *Manager) Register(producter producter) {
	self.product = producter
}

func (self *Manager) Create(name string) producter {
	producter := self.product
	return producter.clone()
}

type Product struct {
	name string
}

func (self *Product) SetUp() {
	// something takes time...
}

func (self *Product) GetName() string {
	return self.name
}

func (self *Product) clone() producter {
	return &Product{self.name}
}
