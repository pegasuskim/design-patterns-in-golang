package abstract_factory

import (
	"testing"
)

func TestAbstractFactory(t *testing.T) {

	factory := MdFactory{}

	usYahoo := factory.CreateLink("Yahoo!", "http://www.yahoo.com")
	jaYahoo := factory.CreateLink("Yahoo! korea", "http://www.yahoo.co.kr")

	tray := factory.CreateTray("Yahoo!")
	tray.AddToTray(usYahoo)
	tray.AddToTray(jaYahoo)

	page := factory.CreatePage("Title", "Author")
	page.AddToContent(tray)

	output := page.Output()

	expect := "title: Title\nauthor: Author\n- Yahoo!\n[Yahoo!](http://www.yahoo.com)\n[Yahoo!](http://www.yahoo.co.kr)\n\n"

	if output != expect {
		t.Errorf("Expect output to %s, but %s\n", expect, output)
	}

}
