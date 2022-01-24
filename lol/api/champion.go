package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetChampionFromName(name string) (*Champion, error) {
	resp, err := http.Get(fmt.Sprintf("%s/champion/%s.json", dataDragonUrl, name))
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

	var res map[string]interface{}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, fmt.Errorf("error while marshaling champion data: %v", err)
	}

	if _, ok := res["data"]; !ok {
		return nil, fmt.Errorf("error data not found in response")
	}
	champData := res["data"].(map[string]interface{})
	if _, ok := champData[name]; !ok {
		return nil, fmt.Errorf("error champion not found in response")
	}
	buf, _ := json.Marshal(champData[name])

	champion := Champion{}
	err = json.Unmarshal(buf, &champion)
	if err != nil {
		return nil, fmt.Errorf("error while marshaling champion data: %v", err)
	}
	return &champion, nil
}
