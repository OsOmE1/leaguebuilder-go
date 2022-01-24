package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Summoner struct {
	Id            string `json:"id,omitempty"`
	AccountId     string `json:"accountId,omitempty"`
	PUUID         string `json:"puuid,omitempty"`
	Name          string `json:"name,omitempty"`
	ProfileIconId int    `json:"profileIconId,omitempty"`
	RevisionDate  int64  `json:"revisionDate,omitempty"`
	SummonerLevel int    `json:"summonerLevel,omitempty"`
}

func SummonerFromAccount(acc *Account) (*Summoner, error) {
	resp, err := http.Get(fmt.Sprintf("%s/summoner/v4/summoners/by-puuid/%s?api_key=%s", lolApiBase, acc.PUUID, apiKey))
	if err != nil {
		return nil, err
	}

	return processSummonerResponse(resp)
}

func SummonerFromName(name string) (*Summoner, error) {
	resp, err := http.Get(fmt.Sprintf("%s/summoner/v4/summoners/by-name/%s?api_key=%s", lolApiBase, name, apiKey))
	if err != nil {
		return nil, err
	}

	return processSummonerResponse(resp)
}

func SummonerFromPUUID(puuid string) (*Summoner, error) {
	resp, err := http.Get(fmt.Sprintf("%s/summoner/v4/summoners/by-puuid/%s?api_key=%s", lolApiBase, puuid, apiKey))
	if err != nil {
		return nil, err
	}

	return processSummonerResponse(resp)
}

func SummonerFromId(id string) (*Summoner, error) {
	resp, err := http.Get(fmt.Sprintf("%s/summoner/v4/summoners/%s?api_key=%s", lolApiBase, id, apiKey))
	if err != nil {
		return nil, err
	}

	return processSummonerResponse(resp)
}

func SummonerFromAccountId(id string) (*Summoner, error) {
	resp, err := http.Get(fmt.Sprintf("%s/summoner/v4/summoners/by-account/%s?api_key=%s", lolApiBase, id, apiKey))
	if err != nil {
		return nil, err
	}

	return processSummonerResponse(resp)
}

func processSummonerResponse(resp *http.Response) (*Summoner, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, statusCodeToError(resp.StatusCode, string(body))
	}
	summoner := Summoner{}
	err = json.Unmarshal(body, &summoner)
	if err != nil {
		return nil, fmt.Errorf("error while marshaling summoner data: %v", err)
	}
	return &summoner, nil
}
