package util

import (
	"encoding/json"
)

type Item struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	URL         string `json:"html_url"`
	Description string `json:"description"`
}

type GithubResponse struct {
	TotalCount        int    `json:"total_count"`
	IncompleteResults bool   `json:"incomplete_results"`
	Items             []Item `json:"items"`
}

func ParseResponse(resp []byte) (*GithubResponse, error) {
	jsonResp := &GithubResponse{}
	if err := json.Unmarshal(resp, jsonResp); err != nil {
		return &GithubResponse{}, err
	}

	return jsonResp, nil
}
