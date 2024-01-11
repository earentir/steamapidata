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
	Type                string `json:"type"`
	Name                string `json:"name"`
	SteamAppid          int    `json:"steam_appid"`
	IsFree              bool   `json:"is_free"`
	Dlc                 []int  `json:"dlc"`
	DetailedDescription string `json:"detailed_description"`
	AboutTheGame        string `json:"about_the_game"`
	ShortDescription    string `json:"short_description"`
	SupportedLanguages  string `json:"supported_languages"`
	HeaderImage         string `json:"header_image"`
	CapsuleImage        string `json:"capsule_image"`
	CapsuleImagev5      string `json:"capsule_imagev5"`
	PcRequirements      struct {
		Minimum string `json:"minimum"`
	} `json:"pc_requirements"`
	MacRequirements struct {
		Minimum string `json:"minimum"`
	} `json:"mac_requirements"`
	LinuxRequirements struct {
		Minimum string `json:"minimum"`
	} `json:"linux_requirements"`
	Developers []string `json:"developers"`
	Publishers []string `json:"publishers"`
	Platforms  struct {
		Windows bool `json:"windows"`
		Mac     bool `json:"mac"`
		Linux   bool `json:"linux"`
	} `json:"platforms"`
	Categories []struct {
		ID          int    `json:"id"`
		Description string `json:"description"`
	} `json:"categories"`
	Genres []struct {
		ID          string `json:"id"`
		Description string `json:"description"`
	} `json:"genres"`
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
	Background         string `json:"background"`
	BackgroundRaw      string `json:"background_raw"`
	ContentDescriptors struct {
		Ids   []int  `json:"ids"`
		Notes string `json:"notes"`
	}
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
