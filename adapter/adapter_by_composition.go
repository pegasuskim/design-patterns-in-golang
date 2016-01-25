package adapter

type ComposiotionDecorateBanner struct {
	banner *Banner
}

func NewCompositionDecorateBanner(str string) *ComposiotionDecorateBanner {
	return &ComposiotionDecorateBanner{&Banner{str}}
}

func (self *ComposiotionDecorateBanner) Decorate() string {
	return self.banner.getString()
}
