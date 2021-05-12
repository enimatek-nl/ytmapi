package webdriver

import "fmt"

type Element struct {
	wd *WebDriver
	id string
}

func (e *Element) GetText() (txt string, err error) {
	text, err := e.wd.gecko.Get(fmt.Sprintf("/session/%v/element/%v/text", e.wd.sessionId, e.id))
	if err == nil {
		txt = text.S("value")
	}
	return
}

func (e *Element) GetAttribute(query string) (txt string, err error) {
	text, err := e.wd.gecko.Get(fmt.Sprintf("/session/%v/element/%v/attribute/%v", e.wd.sessionId, e.id, query))
	if err == nil {
		txt = text.S("value")
	}
	return
}

func (e *Element) FindElements(query string) (elements []*Element, err error) {
	data := []byte(fmt.Sprintf(`{"using": "css selector", "value": "%v"}`, query))
	res, err := e.wd.gecko.Post(fmt.Sprintf("/session/%v/element/%v/elements", e.wd.sessionId, e.id), data)
	if err == nil {
		elements = make([]*Element, 0, len(res.A("value")))
		for _, v := range res.A("value") {
			for _, v := range v.(map[string]interface{}) {
				elements = append(elements, &Element{wd: e.wd, id: v.(string)})
			}
		}
	}
	return
}
