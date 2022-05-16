package main

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Cartoon struct {
	ID       string
	title    *string
	episodes *[]Episode
}

func (c *Cartoon) Title() (*string, error) {
	var err error
	if c.title == nil {
		err = c.Update()
	}
	return c.title, err
}

func (c *Cartoon) Episodes() (*[]Episode, error) {
	var err error
	if c.episodes == nil {
		err = c.Update()
	}
	return c.episodes, err
}

func (c *Cartoon) Update() error {
	resp, err := http.Get("https://kimcartoon.li/Cartoon/" + c.ID)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}
	tt := doc.Find("div.heading").First().Text()
	c.title = &tt
	var eps []Episode
	doc.Find("li").Each(func(_ int, sel *goquery.Selection) {
		sel = sel.Find("a")
		rel, ex := sel.Attr("rel")
		if !ex {
			return
		}
		if rel != "noreferrer noopener" {
			return
		}
		n := sel.Text()
		l, ex := sel.Attr("href")
		if !ex {
			return
		}
		epID := l[strings.LastIndex(l, "/")+1 : strings.LastIndex(l, "?")]
		eps = append(eps, Episode{epID, n, c, nil, nil})
	})
	c.episodes = &eps
	return nil
}

func SearchCartoons(keywords string) (*[]Cartoon, error) {
	val := url.Values{}
	val.Set("keyword", keywords)
	resp, err := http.Post(
		"https://kimcartoon.li/Search/Cartoon",
		"application/x-www-form-urlencoded",
		strings.NewReader(val.Encode()),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	var carts []Cartoon
	doc.Find("a").Each(func(_ int, sel *goquery.Selection) {
		t, ex := sel.Find("img").Attr("title")
		if !ex {
			return
		}
		l, ex := sel.Attr("href")
		if !ex {
			return
		}
		cartID := l[strings.LastIndex(l, "/")+1:]
		carts = append(carts, Cartoon{cartID, &t, nil})
	})
	return &carts, nil
}
