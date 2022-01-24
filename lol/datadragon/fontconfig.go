package datadragon

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type ApiFontConfig struct {
	Entries map[string]string `json:"entries"`
}

type FontConfig struct {
	entries map[string]string
}

func GetFontConfig(lang string) (*FontConfig, error) {
	resp, err := http.Get(fmt.Sprintf("%s/game/data/menu/fontconfig_%s.txt.json", cDragonUrl, strings.ToLower(lang)))
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected code: %d", resp.StatusCode)
	}

	apiresp := ApiFontConfig{}

	err = json.Unmarshal(body, &apiresp)
	if err != nil {
		return nil, fmt.Errorf("error while marshaling champion data: %v", err)
	}

	return &FontConfig{entries: apiresp.Entries}, nil
}

func (f *FontConfig) Get(key string) (string, error) {
	if val, ok := f.entries[key]; ok {
		return val, nil
	}
	if val, ok := f.entries[fmt.Sprintf("{%s}", GetXXH64(key))]; ok {
		return val, nil
	}
	return "", fmt.Errorf("could not find entry: %s({%s})", key, GetXXH64(key))
}
