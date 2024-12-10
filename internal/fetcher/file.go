package fetcher

import (
	"io/ioutil"
)

type FileFetcher struct{}

func (f *FileFetcher) Fetch(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
