package webdriver

import (
	"fmt"
	"ytmapi/geckodriver"
)

type WebDriver struct {
	sessionId string
	gecko     *geckodriver.GeckoDriver
}

func NewWebDriver(gecko *geckodriver.GeckoDriver) (d *WebDriver) {
	d = &WebDriver{}
	d.gecko = gecko
	return
}

func (d *WebDriver) StartSession() error {
	capabilities := []byte(`{"capabilities":{"firstMatch":[{"acceptInsecureCerts":true,"browserName": "firefox","moz:firefoxOptions":{"args":["-headless"],"log":{"level":"trace"}}}]}}`)
	session, err := d.gecko.Post("/session", capabilities)
	if err == nil {
		d.sessionId = session.M("value").S("sessionId")
	} else {
		return err
	}
	return nil
}

func (d *WebDriver) StopSession() error {
	err := d.gecko.Delete(fmt.Sprintf("/session/%v", d.sessionId))
	return err
}

func (d *WebDriver) Navigate(url string) error {
	data := []byte(fmt.Sprintf(`{"url": "%s"}`, url))
	_, err := d.gecko.Post(fmt.Sprintf("/session/%v/url", d.sessionId), data)
	return err
}

func (d *WebDriver) FindElement(query string) (elem *Element, err error) {
	data := []byte(fmt.Sprintf(`{"using": "css selector", "value": "%v"}`, query))
	res, err := d.gecko.Post(fmt.Sprintf("/session/%v/element", d.sessionId), data)
	if err == nil {
		keys := make([]string, 0, 1)
		for k := range res.M("value") {
			keys = append(keys, k)
		}
		elem = &Element{wd: d, id: res.M("value").S(keys[0])}
	}
	return
}

func (d *WebDriver) FindElements(query string) (elements []*Element, err error) {
	data := []byte(fmt.Sprintf(`{"using": "css selector", "value": "%v"}`, query))
	res, err := d.gecko.Post(fmt.Sprintf("/session/%v/elements", d.sessionId), data)
	if err == nil {
		elements = make([]*Element, 0, len(res.A("value")))
		for _, v := range res.A("value") {
			for _, v := range v.(map[string]interface{}) {
				elements = append(elements, &Element{wd: d, id: v.(string)})
			}
		}
	}
	return
}
