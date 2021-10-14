package main

import (
	"fmt"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"
const APIURL = "https://api.github.com"

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"total_count"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type IssuesSearchResult struct {
	TotalCount int
	Items      []*Issue
}

func (i Issue) CacheURL() string {
	return fmt.Sprintf("/issues/%d", i.Number)
}
