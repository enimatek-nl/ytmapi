package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"runtime"
	"time"
	"ytmapi/model"
)

type APIReply struct {
	Status  int
	Error   string
	Results []*YTMusicItem
}

type YTMusicItem struct {
	Id       string
	Artist   string
	Title    string
	Album    string
	Duration string
}

func main() {

	var port string
	flag.StringVar(&port, "p", "8700", "Specify the listening port.")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		r := APIReply{Status: 500}
		if err := req.ParseForm(); err == nil {
			if r.Results, err = doSearch(req.Form.Get("q")); err == nil {
				if r.Results != nil {
					r.Status = 200
				}
			} else {
				r.Error = err.Error()
			}
		} else {
			r.Error = err.Error()
		}
		json.NewEncoder(w).Encode(r)
	})

	log.Printf("OS: %s, Architecture: %s", runtime.GOOS, runtime.GOARCH)
	log.Printf("API listening on port: %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func doSearch(q string) (list []*YTMusicItem, err error) {
	if ytm, err := doGet(q); err == nil {
		for _, tab := range ytm.Contents.TabbedSearchResultsRenderer.Tabs {
			for _, content := range tab.TabRenderer.Content.SectionListRenderer.Contents {
				for _, shelf := range content.MusicShelfRenderer.Contents {
					kind := shelf.MusicResponsiveListItemRenderer.FlexColumns[1].MusicResponsiveListItemFlexColumnRenderer.Text.Runs[0].Text
					if kind == "Song" {
						item := &YTMusicItem{
							Id:       shelf.MusicResponsiveListItemRenderer.Overlay.MusicItemThumbnailOverlayRenderer.Content.MusicPlayButtonRenderer.PlayNavigationEndpoint.WatchEndpoint.VideoId,
							Title:    shelf.MusicResponsiveListItemRenderer.FlexColumns[0].MusicResponsiveListItemFlexColumnRenderer.Text.Runs[0].Text,
							Artist:   shelf.MusicResponsiveListItemRenderer.FlexColumns[1].MusicResponsiveListItemFlexColumnRenderer.Text.Runs[2].Text,
							Album:    shelf.MusicResponsiveListItemRenderer.FlexColumns[1].MusicResponsiveListItemFlexColumnRenderer.Text.Runs[4].Text,
							Duration: shelf.MusicResponsiveListItemRenderer.FlexColumns[1].MusicResponsiveListItemFlexColumnRenderer.Text.Runs[6].Text,
						}
						list = append(list, item)
					}
				}
			}
		}
	}
	log.Printf("Searched for `%s` and found %d results.", q, len(list))
	return
}

func doGet(q string) (r model.YTSearch, err error) {
	body := &model.YTContext{}
	body.Context.Client.ClientName = "WEB_REMIX"
	body.Context.Client.ClientVersion = "1.20211201.00.01"
	body.Query = url.QueryEscape(q)

	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	b, err := json.Marshal(body)
	if err != nil {
		return
	}
	req, err := http.NewRequest("POST", "https://music.youtube.com/youtubei/v1/search?key=AIzaSyC9XL3ZjWddXya6X74dJoCTL-WEYFDNX30", bytes.NewBuffer(b))
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.87 Safari/537.36")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Referer", "https://music.youtube.com/")

	if resp, err := client.Do(req); err == nil {
		defer resp.Body.Close()
		json.NewDecoder(resp.Body).Decode(&r)
	}

	return
}
