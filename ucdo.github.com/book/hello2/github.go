package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const queryUrl = "https://api.github.com/search/issues"

type IssuesRes struct {
	TotalCount int `json:"count"`
	Issues     []*Issues
}

type Issues struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func IssuesQuery(params []string) (*IssuesRes, error) {
	q := url.QueryEscape(strings.Join(params, " "))
	fmt.Println(queryUrl + "?q=" + q)
	res, err := http.Get(queryUrl + "?q=" + q)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", res.Status)
	}

	var result IssuesRes
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func main() {
	result, err := IssuesQuery(os.Args[1:])
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Issues {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
