package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Episode struct {
	ID      string
	Name    string
	Cartoon *Cartoon
	videos  *[]Video
	videoID *string
}

func (e *Episode) Videos() (*[]Video, error) {
	if e.videos != nil {
		return e.videos, nil
	}
	vidID, err := e.VideoID()
	if err != nil {
		return nil, err
	}
	resp, err := http.Post("https://www.luxubu.review/api/source/"+*vidID, "application/x-www-form-urlencoded", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var jb struct {
		Data []Video `json:"data"`
	}
	if err := json.Unmarshal(body, &jb); err != nil {
		return nil, err
	}
	var vids []Video
	for _, v := range jb.Data {
		v.Episode = e
		vids = append(vids, v)
	}
	e.videos = &vids
	return e.videos, nil
}

func (e *Episode) VideoID() (*string, error) {
	if e.videoID != nil {
		return e.videoID, nil
	}
	resp, err := http.Get("https://kimcartoon.li/Cartoon/" + e.Cartoon.ID + "/" + e.ID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	src, ex := doc.Find("iframe#mVideo").Attr("src")
	if !ex {
		return nil, errors.New("cannot find source")
	}
	vidi := src[strings.LastIndex(src, "/")+1:]
	e.videoID = &vidi
	return e.videoID, nil
}
