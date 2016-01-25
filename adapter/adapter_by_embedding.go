package adapter

type Decorator interface {
	Decorate() string
}

type Banner struct {
	str string
}

func (self *Banner) getString() string {
	return "*" + self.str + "*"
}

type EmbeddedDecorateBanner struct {
	*Banner
}

func NewEmbeddedDecorateBanner(str string) *EmbeddedDecorateBanner {
	return &EmbeddedDecorateBanner{&Banner{str}}
}

func (self *EmbeddedDecorateBanner) Decorate() string {
	return self.getString()
}
