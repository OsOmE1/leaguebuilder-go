package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Match struct {
	Metadata MatchMetadata `json:"metadata"`
	Info     MatchInfo     `json:"info"`
}

type MatchTimeline struct {
	Metadata MatchMetadata     `json:"metadata"`
	Info     MatchTimelineInfo `json:"info"`
}

func GetMatchFromId(id string) (*Match, error) {
	resp, err := http.Get(fmt.Sprintf("%s/lol/match/v5/matches/%s?api_key=%s", riotLolApiBase, id, apiKey))
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, statusCodeToError(resp.StatusCode, string(body))
	}
	match := Match{}
	err = json.Unmarshal(body, &match)
	if err != nil {
		return nil, fmt.Errorf("error while marshaling match data: %v", err)
	}
	return &match, nil
}
