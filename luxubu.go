package main

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
