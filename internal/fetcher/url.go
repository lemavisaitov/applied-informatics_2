package fetcher

import (
	"io/ioutil"
	"log"
	"net/http"
)

type URLFetcher struct{}

func (u *URLFetcher) Fetch(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching URL %s: %v", url, err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
