package factory_method

type creater interface {
	createProduct(owner string) User
	registerProduct(User)
}

type User interface {
	Use() string
}

type Factory struct {
}

// func (self *Factory) Create(factoryMethod func() User) {
//   self.product = factoryMethod()
//   return self.product
// }
func (self *Factory) Create(factory creater, owner string) User {
	user := factory.createProduct(owner)
	factory.registerProduct(user)
	return user
}

// type Product struct {
// }

type IDCard struct {
	owner string
}

func (self *IDCard) Use() string {
	return self.owner
}

type IDCardFactory struct {
	*Factory
	owners []*string
}

func (self *IDCardFactory) createProduct(owner string) User {
	return &IDCard{owner}
}

func (self *IDCardFactory) registerProduct(product User) {
	owner := product.Use()
	self.owners = append(self.owners, &owner)
}
