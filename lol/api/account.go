package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Account struct {
	PUUID    string `json:"puuid,omitempty"`
	GameName string `json:"gameName,omitempty"`
	TagLine  string `json:"tagLine,omitempty"`
}

func AccountFromRiotId(name, tag string) (*Account, error) {
	resp, err := http.Get(fmt.Sprintf("%s/account/v1/accounts/by-riot-id/%s/%s?api_key=%s", riotApiBase, name, tag, apiKey))
	if err != nil {
		return nil, err
	}

	return processAccountResponse(resp)
}

func AccountFromPUUID(id string) (*Account, error) {
	resp, err := http.Get(fmt.Sprintf("%s/account/v1/accounts/by-puuid/%s?api_key=%s", riotApiBase, id, apiKey))
	if err != nil {
		return nil, err
	}

	return processAccountResponse(resp)
}

func processAccountResponse(resp *http.Response) (*Account, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, statusCodeToError(resp.StatusCode, string(body))
	}
	acc := Account{}
	err = json.Unmarshal(body, &acc)
	if err != nil {
		return nil, fmt.Errorf("error while marshaling account data: %v", err)
	}
	return &acc, nil
}
