package api

import "fmt"

const (
	apiKey         = "RGAPI-ea3be1b7-5d9f-4b17-a8af-95bdb1f8cdc7"
	lolApiBase     = "https://euw1.api.riotgames.com/lol"
	riotApiBase    = "https://europe.api.riotgames.com/riot"
	riotLolApiBase = "https://europe.api.riotgames.com/lol"
	dataDragonUrl  = "http://ddragon.leagueoflegends.com/cdn/12.1.1/data/en_US"
)

func statusCodeToError(code int, content string) error {
	switch code {
	case 400:
		return fmt.Errorf("bad request: %s", content)
	case 401:
		return fmt.Errorf("unauthorized: %s", content)
	case 403:
		return fmt.Errorf("forbidden: %s", content)
	case 404:
		return fmt.Errorf("data not found: %s", content)
	case 405:
		return fmt.Errorf("method not allowed: %s", content)
	case 415:
		return fmt.Errorf("unsupported media type: %s", content)
	case 429:
		return fmt.Errorf("rate limit exceeded: %s", content)
	case 500:
		return fmt.Errorf("internal server error: %s", content)
	case 502:
		return fmt.Errorf("bad gateway: %s", content)
	case 503:
		return fmt.Errorf("service unavailable: %s", content)
	case 504:
		return fmt.Errorf("gateway timeout: %s", content)
	default:
		return fmt.Errorf("%d: %s", code, content)
	}
}
