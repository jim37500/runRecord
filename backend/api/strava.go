package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"runRecord/data"
	"runRecord/database"
	"runRecord/model"
)

var urlMap = map[string]string{
	"getLoggedInAthleteActivities": "https://www.strava.com/api/v3/athlete/activities",
	"getLapsByActivityId":          "https://www.strava.com/api/v3/activities/{id}/laps",
	"refreshToken":                 "https://www.strava.com/api/v3/oauth/token?", // client_id={client_id}&client_secret={client_secret}&grant_type=refresh_token&refresh_token={refresh_token}
}

func FetchStravaApi[T any](method string, url string, accessToken string, result T) (T, error) {
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return result, err
	}

	request.Header.Add("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("Error making request:", err)
		return result, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return result, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return result, err
	}

	return result, nil
}

func FetchNewAccessToken(myAthlete model.Athlete) string {
	refreshTokenUrl := fmt.Sprintf("%sclient_id=%d&client_secret=%s&grant_type=refresh_token&refresh_token=%s", urlMap["refreshToken"], myAthlete.ClientID, myAthlete.ClientSecret, myAthlete.RefreshToken)

	tokens, err := FetchStravaApi("POST", refreshTokenUrl, "", data.Token{})
	if err != nil {
		return ""
	}

	_ = database.UpdateAthleteToken(myAthlete.ID, tokens.AccessToken, tokens.RefreshToken)

	return tokens.AccessToken
}

func FetchStravaActivities(myAthlete model.Athlete) error {
	url := urlMap["getLoggedInAthleteActivities"]

	lastRunActivity := database.GetLatestRunActivity(uint64(myAthlete.ID)) // 取得運動員最近一次的跑步活動
	// 若 有最後一次跑步活動 則 爬取這之後的活動
	if lastRunActivity.ID != 0 {
		url += fmt.Sprintf("?after=%d", lastRunActivity.Date.Unix())
	}

	result, err := FetchStravaApi("GET", url, myAthlete.AccessToken, []data.RunActivities{})
	if err != nil {
		accessToken := FetchNewAccessToken(myAthlete)

		if accessToken == "" {
			return errors.New("Invalid AccessToken")
		}

		result, err = FetchStravaApi("GET", url, accessToken, []data.RunActivities{})
		if err != nil {
			return errors.New("Error fetching strava api")
		}
	}

	if !database.AddRunActivity(result) {
		return errors.New("Error Saving run activities")
	}

	return nil
}

func FetchStravaLaps(myAthlete model.Athlete, activityID string) error {
	url := strings.Replace(urlMap["getLapsByActivityId"], "{id}", activityID, 1)

	result, err := FetchStravaApi("GET", url, myAthlete.AccessToken, []data.RunLap{})
	if err != nil {
		accessToken := FetchNewAccessToken(myAthlete)
		if accessToken == "" {
			return errors.New("Invalid AccessToken")
		}

		result, err = FetchStravaApi("GET", url, accessToken, []data.RunLap{})
		if err != nil {
			return errors.New("Error fetching strava api")
		}
	}

	if !database.AddRunLaps(result) {
		return errors.New("Error Saving run laps error")
	}

	return nil
}
