package statuspage

import (
	"encoding/json"
	"log"
	"net/http"
)

type Status struct {
	Page struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"page"`
	Components []struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Status string `json:"status"`
	} `json:"components"`
}

func Download(statusPage *StatusPage) (Status, error) {
	resp, err := http.Get(GetStatusPageComponentURL(statusPage))

	if err != nil {
		log.Fatal("error occurred, please try again")
	}
	defer resp.Body.Close()
	var cResp Status
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		log.Fatal("error occurred, please try again")
	}
	return cResp, nil
}
