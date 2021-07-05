package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
	"ytmapi/geckodriver"
	"ytmapi/webdriver"
)

type API struct {
	port    string
	driver  string
	verbose bool
}

type YTContainer struct {
	Id       string
	Artist   string
	Title    string
	Duration string
}

type BasicReply struct {
	Status int
	Error  string
	Results []*YTContainer
}

func NewAPI(port string, driver string, verbose bool) *API {
	api := &API{
		port:    port,
		driver:  driver,
		verbose: verbose,
	}
	return api
}

func (a *API) Start() {
	http.HandleFunc("/", a.handleSearch)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", a.port), nil))
}

func (a *API) handleSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	q, p := query["q"]
	if !p {
		w.WriteHeader(404)
		r := &BasicReply{
			Status: 404,
			Error:  "search query in the get parameter 'q' is missing",
		}
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(r)
	} else {
		containers, err := a.search(strings.Join(q, " "))
		if err != nil {
			r := &BasicReply{
				Status: 500,
				Error:  err.Error(),
			}
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(r)
		} else {
			r := &BasicReply{
				Status: 200,
				Results: containers,
			}
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(r)
		}
	}
}

func (a *API) search(q string) (c []*YTContainer, err error) {
	gecko := geckodriver.NewGeckoDriver(a.driver, a.verbose)

	if err := gecko.Start(); err != nil {
		return nil, err
	}
	defer gecko.Stop()

	driver := webdriver.NewWebDriver(gecko)

	if err := driver.StartSession(); err != nil {
		return nil, err
	}
	defer driver.StopSession()

	if err := driver.Navigate(fmt.Sprintf("https://music.youtube.com/search?q=%v", url.QueryEscape(q))); err != nil {
		return nil, err
	}

	var element *webdriver.Element = nil
	var elements []*webdriver.Element = nil

	start := time.Now().Add(time.Second * 10)
	for element == nil {
		if time.Now().After(start) {
			return nil, errors.New("timeout")
		}
		element, err = driver.FindElement(`ytmusic-section-list-renderer`)
	}

	start = time.Now().Add(time.Second * 10)
	for elements == nil {
		if time.Now().After(start) {
			return nil, errors.New("timeout")
		}
		elements, _ = element.FindElements(`ytmusic-shelf-renderer`)
	}

	//log.Println(len(elements))
	for _, e := range elements {
		header, _ := e.FindElements("h2")
		if len(header) > 0 {
			head, _ := header[0].GetText()
			if strings.HasPrefix(head, "Song") {
				songs, _ := e.FindElements("ytmusic-responsive-list-item-renderer")
				for _, s := range songs {
					container := &YTContainer{}
					links, _ := s.FindElements("a")
					if len(links) > 0 {
						uri, _ := links[0].GetAttribute("href")
						s := strings.Split(uri, "?v=")
						if strings.Contains(s[1], "&") {
							s = strings.Split(s[1], "&")
							container.Id = s[0]
						} else {
							container.Id = s[1]
						}
					}
					texts, _ := s.FindElements("yt-formatted-string")
					if len(texts) >= 1 {
						title, _ := texts[1].GetAttribute("title")
						s := strings.Split(title, " • ")
						container.Artist = s[1]
						container.Title = s[2]
						if len(s) >= 3 {
							container.Duration = s[3]
						}
						c = append(c, container)
					}
				}
			}
		}

	}

	return
}
