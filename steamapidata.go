// Package steamapidata provides functions for retrieving data from the Steam API
package steamapidata

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"net/http"
)

// GetSteamID converts a Steam username to a Steam ID
func GetSteamID(APIKey, username string) (string, int, error) {
	url := fmt.Sprintf("http://api.steampowered.com/ISteamUser/ResolveVanityURL/v0001/?key=%s&vanityurl=%s", APIKey, username)

	resp, err := http.Get(url)
	if err != nil {
		return "", resp.StatusCode, err
	}
	defer resp.Body.Close()

	// Check for HTTP status code errors
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusForbidden {
			return "", resp.StatusCode, fmt.Errorf("access forbidden: check your API key")
		} else if resp.StatusCode == http.StatusNotFound {
			return "", resp.StatusCode, fmt.Errorf("API endpoint not found")
		}

		return "", resp.StatusCode, fmt.Errorf("unexpected HTTP status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", resp.StatusCode, err
	}

	var response SteamAPIResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return "", resp.StatusCode, err
	}

	if response.Response.Success != 1 {
		fmt.Println("statuscode: ", resp.StatusCode)
		return "", resp.StatusCode, fmt.Errorf("failed to find Steam ID for username: %s", username)
	}

	return response.Response.SteamID, resp.StatusCode, nil
}

// SteamAppDetails retrieves details for a Steam game
func SteamAppDetails(appID int) (*SteamAppData, error) {
	var jsonData string
	appidstr := strconv.Itoa(appID)
	fileName := fmt.Sprintf("steamdata/%s.json", appidstr)

	// Check if file exists
	if _, err := os.Stat(fileName); err == nil {
		// File exists, read it
		byteData, err := os.ReadFile(fileName)
		if err != nil {
			return nil, err
		}
		jsonData = string(byteData)
	} else {
		// File does not exist, make HTTP call
		url := fmt.Sprintf("http://store.steampowered.com/api/appdetails?appids=%d&l=english", appID)
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		byteData, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		jsonData = string(byteData)

		// Save the data to a file
		err = os.WriteFile(fileName, byteData, 0644)
		if err != nil {
			return nil, err
		}
	}

	// Decode JSON data
	var detailsResponse SteamAppDetailsResponse
	err := json.Unmarshal([]byte(jsonData), &detailsResponse)
	if err != nil {
		return nil, err
	}

	result, ok := detailsResponse[appidstr]
	if !ok || !result.Success {
		return nil, fmt.Errorf("Failed to retrieve game details for app ID %d", appID)
	}

	return &result.Data, nil
}

// SteamUserAppsUsed retrieves the games played by a Steam user
func SteamUserAppsUsed(APIKey, steamID string) ([]AppsUsedInfo, error) {
	url := fmt.Sprintf("http://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=%s&steamid=%s&include_appinfo=true&include_played_free_games=true&format=json", APIKey, steamID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var appResponse UserAppsUsed
	err = json.NewDecoder(resp.Body).Decode(&appResponse)
	if err != nil {
		return nil, err
	}

	if appResponse.Response.AppCount == 0 {
		return nil, fmt.Errorf("no games found for Steam ID %s", steamID)
	}

	var apps []AppsUsedInfo
	for _, app := range appResponse.Response.Apps {
		gameInfo := AppsUsedInfo{
			Appid:           app.Appid,
			Name:            app.Name,
			PlaytimeForever: app.PlaytimeForever,
			RtimeLastPlayed: app.RtimeLastPlayed,
		}
		apps = append(apps, gameInfo)
	}

	return apps, nil
}

func getSteamAppList() (*steamAppList, error) {
	const url = "http://api.steampowered.com/ISteamApps/GetAppList/v2"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Applist steamAppList `json:"applist"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result.Applist, nil
}

func cleanString(s string) string {
	var cleaned strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			cleaned.WriteRune(r)
		}
	}
	return strings.ToLower(cleaned.String())
}

func findSteamGame(appList *steamAppList, input string) (string, error) {
	cleanedInput := cleanString(input)

	if appID, err := strconv.Atoi(input); err == nil {
		for _, app := range appList.Apps {
			if app.AppID == appID {
				return app.Name, nil
			}
		}
		return "", fmt.Errorf("AppID %d not found", appID)
	}

	for _, app := range appList.Apps {
		if cleanString(app.Name) == cleanedInput {
			return strconv.Itoa(app.AppID), nil
		}
	}
	return "", fmt.Errorf("Game %s not found", input)
}

// SteamSearchApp searches for a game in the Steam app list
func SteamSearchApp(input string) (string, error) {
	const fileName = "steamdata/steamgames.json"

	var appList *steamAppList

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		appList, err = getSteamAppList()
		if err != nil {
			return "", err
		}
		file, _ := json.MarshalIndent(appList, "", " ")
		_ = os.WriteFile(fileName, file, 0644)
	} else {
		file, _ := os.ReadFile(fileName)
		appList = &steamAppList{}
		_ = json.Unmarshal(file, appList)
	}

	return findSteamGame(appList, input)
}

// SortApps sorts a slice of GamesPlayedInfo based on the specified criterion.
func SortApps(apps []AppsUsedInfo, criterion string, count int) []AppsUsedInfo {
	switch criterion {
	case "playtime":
		sort.Slice(apps, func(i, j int) bool {
			return apps[i].PlaytimeForever > apps[j].PlaytimeForever
		})
	case "lastplayed":
		sort.Slice(apps, func(i, j int) bool {
			return apps[i].RtimeLastPlayed > apps[j].RtimeLastPlayed
		})
	}

	if count > 0 && count < len(apps) {
		return apps[:count]
	}
	return apps
}
