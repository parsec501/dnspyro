package dns

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WhoisResponse struct {
	Data struct {
		IrrRecords [][]struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"irr_records"`
	} `json:"data"`
}

const (
	API_URL = "https://stat.ripe.net/data/whois/data.json?resource="
)

func ResolveHostAPI(ip string) (string, error) {
	resp, err := http.Get(API_URL + ip)
	if err != nil {
		return "", err
	}

	var whoisResponse WhoisResponse
	err = json.NewDecoder(resp.Body).Decode(&whoisResponse)
	if err != nil {
		return "", err
	}

	for _, irrRecords := range whoisResponse.Data.IrrRecords {
		for _, irrRecord := range irrRecords {
			if irrRecord.Key == "descr" {
				return irrRecord.Value, nil
			}
		}

	}
	return "", fmt.Errorf("key 'descr' not found in irrRecords")
}
