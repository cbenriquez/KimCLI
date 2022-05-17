package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type LuxubuAPISourceResponse struct {
	Success bool                `json:"success"`
	Player  Player              `json:"player"`
	Data    []map[string]string `json:"data"`
}

type Player struct {
	PosterFile           string `json:"poster_file"`
	LogoFile             string `json:"logo_file"`
	LogoPosition         string `json:"logo_position"`
	LogoLink             string `json:"logo_link"`
	LogoMargin           int    `json:"logo_margin"`
	AspectRatio          string `json:"aspectratio"`
	PoweredText          string `json:"powered_text"`
	PoweredURL           string `json:"powered_url"`
	CSSBackground        string `json:"css_background"`
	CSSText              string `json:"css_text"`
	CSSMenu              string `json:"css_menu"`
	CSSMenuText          string `json:"css_mntext"`
	CSSCaption           string `json:"css_caption"`
	CSSCaptionText       string `json:"css_cttext"`
	CSSCaptionSize       string `json:"css_ctsize"`
	CSSCaptionOpacity    string `json:"css_ctopacity"`
	CSSIcon              string `json:"css_icon"`
	CSSIconHover         string `json:"css_ichover"`
	CSSTimestampProgress string `json:"css_tsprogress"`
	CSSTimestampRail     string `json:"css_tsrail"`
	CSSButton            string `json:"css_button"`
	CSSButtonText        string `json:"css_bttext"`
	OptionAutostart      bool   `json:"opt_autostart"`
	OptionTitle          bool   `json:"opt_title"`
	OptionQuality        bool   `json:"opt_quality"`
	OptionCaption        bool   `json:"opt_caption"`
	OptionDownload       bool   `json:"opt_download"`
	OptionSharing        bool   `json:"opt_sharing"`
	OptionPlayRate       bool   `json:"opt_playrate"`
	OptionMute           bool   `json:"opt_mute"`
	OptionLoop           bool   `json:"opt_loop"`
	OptionVR             bool   `json:"opt_vr"`
	OptionCast           Cast   `json:"opt_cast"`
	OptionNoDefault      bool   `json:"opt_nodefault"`
	OptionForcePoster    bool   `json:"opt_forceposter"`
	OptionParameter      bool   `json:"opt_parameter"`
	RestrictDomain       string `json:"restrict_domain"`
	RestrictAction       string `json:"restrict_action"`
	RestrictTarget       string `json:"restrict_target"`
	AdblockEnable        bool   `json:"adb_enable"`
	AdblockOffset        string `json:"adb_offset"`
	AdblockText          string `json:"adb_text"`
	AdsAdult             bool   `json:"ads_adult"`
	AdsPop               bool   `json:"ads_pop"`
	AdsVast              bool   `json:"ads_vast"`
	AdsFree              int    `json:"ads_free"`
	TrackingID           string `json:"trackingId"`
	ViewID               string `json:"viewId"`
	Income               bool   `json:"income"`
	IncomePop            bool   `json:"incomePop"`
	ResumeText           string `json:"resume_text"`
	ResumeYes            string `json:"resume_yes"`
	ResumeNo             string `json:"resume_no"`
	ResumeEnable         bool   `json:"resume_enable"`
	CSSCTEdge            string `json:"css_ctedge"`
	Logger               string `json:"logger"`
	Revenue              string `json:"revenue"`
	RevenueFallback      string `json:"revenue_fallback"`
	RevenueTrack         string `json:"revenue_track"`
}

type Cast struct {
	AppID string `json:"appid"`
}

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
	var jb LuxubuAPISourceResponse
	if err := json.Unmarshal(body, &jb); err != nil {
		return nil, err
	}
	var vids []Video
	for _, v := range jb.Data {
		vids = append(vids, Video{v["file"], v["label"], v["type"], e})
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
