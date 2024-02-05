package steamapidata

type steamAppList struct {
	Apps []steamApp `json:"apps"`
}

type steamApp struct {
	AppID int    `json:"appid"`
	Name  string `json:"name"`
}

// SteamAppDetailsResponse is the response from the Steam API
type SteamAppDetailsResponse map[string]SteamAppDetailResult

// SteamAppDetailResult is the response from the Steam API
type SteamAppDetailResult struct {
	Success bool         `json:"success"`
	Data    SteamAppData `json:"data"`
}

// SteamAppData is the data returned from the Steam API
type SteamAppData struct {
	Type                string      `json:"type"`
	Name                string      `json:"name"`
	SteamAppid          int         `json:"steam_appid"`
	RequiredAge         interface{} `json:"required_age"`
	IsFree              bool        `json:"is_free"`
	Dlc                 []int       `json:"dlc"`
	DetailedDescription string      `json:"detailed_description"`
	AboutTheGame        string      `json:"about_the_game"`
	ShortDescription    string      `json:"short_description"`
	SupportedLanguages  string      `json:"supported_languages"`
	Reviews             string      `json:"reviews"`
	HeaderImage         string      `json:"header_image"`
	CapsuleImage        string      `json:"capsule_image"`
	CapsuleImagev5      string      `json:"capsule_imagev5"`
	Website             string      `json:"website"`
	PcRequirements      struct {
		Minimum     string `json:"minimum"`
		Recommended string `json:"recommended"`
	} `json:"pc_requirements"`
	MacRequirements struct {
		Minimum     string `json:"minimum"`
		Recommended string `json:"recommended"`
	} `json:"mac_requirements"`
	LinuxRequirements struct {
		Minimum     string `json:"minimum"`
		Recommended string `json:"recommended"`
	} `json:"linux_requirements"`
	Developers []string `json:"developers"`
	Publishers []string `json:"publishers"`
	Demos      []struct {
		Appid       int    `json:"appid"`
		Description string `json:"description"`
	} `json:"demos"`
	PriceOverview struct {
		Currency         string `json:"currency"`
		Initial          int    `json:"initial"`
		Final            int    `json:"final"`
		DiscountPercent  int    `json:"discount_percent"`
		InitialFormatted string `json:"initial_formatted"`
		FinalFormatted   string `json:"final_formatted"`
	} `json:"price_overview"`
	Packages      []int `json:"packages"`
	PackageGroups []struct {
		Name                    string `json:"name"`
		Title                   string `json:"title"`
		Description             string `json:"description"`
		SelectionText           string `json:"selection_text"`
		SaveText                string `json:"save_text"`
		DisplayType             int    `json:"display_type"`
		IsRecurringSubscription string `json:"is_recurring_subscription"`
		Subs                    []struct {
			Packageid                int    `json:"packageid"`
			PercentSavingsText       string `json:"percent_savings_text"`
			PercentSavings           int    `json:"percent_savings"`
			OptionText               string `json:"option_text"`
			OptionDescription        string `json:"option_description"`
			CanGetFreeLicense        string `json:"can_get_free_license"`
			IsFreeLicense            bool   `json:"is_free_license"`
			PriceInCentsWithDiscount int    `json:"price_in_cents_with_discount"`
		} `json:"subs"`
	} `json:"package_groups"`
	Platforms struct {
		Windows bool `json:"windows"`
		Mac     bool `json:"mac"`
		Linux   bool `json:"linux"`
	} `json:"platforms"`
	Metacritic struct {
		Score int    `json:"score"`
		URL   string `json:"url"`
	} `json:"metacritic"`
	Categories []struct {
		ID          int    `json:"id"`
		Description string `json:"description"`
	} `json:"categories"`
	Genres []struct {
		ID          string `json:"id"`
		Description string `json:"description"`
	} `json:"genres"`
	Screenshots []struct {
		ID            int    `json:"id"`
		PathThumbnail string `json:"path_thumbnail"`
		PathFull      string `json:"path_full"`
	} `json:"screenshots"`
	Movies []struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Thumbnail string `json:"thumbnail"`
		Webm      struct {
			Num480 string `json:"480"`
			Max    string `json:"max"`
		} `json:"webm"`
		Mp4 struct {
			Num480 string `json:"480"`
			Max    string `json:"max"`
		} `json:"mp4"`
		Highlight bool `json:"highlight"`
	} `json:"movies"`
	Recommendations struct {
		Total int `json:"total"`
	} `json:"recommendations"`
	Achievements struct {
		Total       int `json:"total"`
		Highlighted []struct {
			Name string `json:"name"`
			Path string `json:"path"`
		} `json:"highlighted"`
	} `json:"achievements"`
	ReleaseDate struct {
		ComingSoon bool   `json:"coming_soon"`
		Date       string `json:"date"`
	} `json:"release_date"`
	SupportInfo struct {
		URL   string `json:"url"`
		Email string `json:"email"`
	} `json:"support_info"`
	Background         string `json:"background"`
	BackgroundRaw      string `json:"background_raw"`
	ContentDescriptors struct {
		Ids   []interface{} `json:"ids"`
		Notes interface{}   `json:"notes"`
	} `json:"content_descriptors"`
}

// SteamAPIResponse is the response from the Steam API
type SteamAPIResponse struct {
	Response struct {
		Success int    `json:"success"`
		SteamID string `json:"steamid"`
	} `json:"response"`
}

// AppsUsedInfo is the data returned from the Steam API
type AppsUsedInfo struct {
	Appid           int    `json:"appid"`
	Name            string `json:"name"`
	PlaytimeForever int    `json:"playtime_forever"`
	RtimeLastPlayed int    `json:"rtime_last_played"`
}

// UserAppsUsed is the data returned from the Steam API
type UserAppsUsed struct {
	Response struct {
		AppCount int `json:"game_count"`
		Apps     []struct {
			Appid                    int    `json:"appid"`
			Name                     string `json:"name"`
			PlaytimeForever          int    `json:"playtime_forever"`
			ImgIconURL               string `json:"img_icon_url"`
			HasCommunityVisibleStats bool   `json:"has_community_visible_stats,omitempty"`
			PlaytimeWindowsForever   int    `json:"playtime_windows_forever"`
			PlaytimeMacForever       int    `json:"playtime_mac_forever"`
			PlaytimeLinuxForever     int    `json:"playtime_linux_forever"`
			RtimeLastPlayed          int    `json:"rtime_last_played"`
			PlaytimeDisconnected     int    `json:"playtime_disconnected"`
			ContentDescriptorids     []int  `json:"content_descriptorids,omitempty"`
			HasLeaderboards          bool   `json:"has_leaderboards,omitempty"`
			Playtime2Weeks           int    `json:"playtime_2weeks,omitempty"`
		} `json:"games"`
	} `json:"response"`
}
