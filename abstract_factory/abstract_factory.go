package abstract_factory

type item interface {
	toString() string
}

type link interface {
	item
}

type tray interface {
	item
	AddToTray(item item)
}

type baseTray struct {
	tray []item
}

func (self *baseTray) AddToTray(item item) {
	self.tray = append(self.tray, item)
}

type page interface {
	AddToContent(item item)
	Output() string
}

type basePage struct {
	content []item
}

func (self *basePage) AddToContent(item item) {
	self.content = append(self.content, item)
}

type factory interface {
	CreateLink(caption, url string) link
	CreateTray(caption string) tray
	CreatePage(title, author string) page
}

type mdLink struct {
	caption, url string
}

func (self *mdLink) toString() string {
	return "[" + self.caption + "](" + self.url + ")"
}

type mdTray struct {
	baseTray
	caption string
}

func (self *mdTray) toString() string {
	tray := "- " + self.caption + "\n"
	for _, item := range self.tray {
		tray += item.toString() + "\n"
	}
	return tray
}

type mdPage struct {
	basePage
	title, author string
}

func (self *mdPage) Output() string {
	content := "title: " + self.title + "\n"
	content += "author: " + self.author + "\n"
	for _, item := range self.content {
		content += item.toString() + "\n"
	}
	return content
}

type MdFactory struct {
}

func (self *MdFactory) CreateLink(caption, url string) link {
	return &mdLink{caption, url}
}
func (self *MdFactory) CreateTray(caption string) tray {
	return &mdTray{caption: caption}
}
func (self *MdFactory) CreatePage(title, author string) page {
	return &mdPage{title: title, author: author}
}
