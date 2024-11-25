package linkutil

import (
	"io"
	"net/http"
	"regexp"
)

func LinksFromUrl(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return LinksFromReader(resp.Body)
}

func LinksFromReader(reader io.Reader) (links []string, err error) {
	var bytesData []byte
	if bytesData, err = io.ReadAll(reader); err != nil {
		return links, err
	}

	pattern := regexp.MustCompile("href=\"(.*?)\">")
	matches := pattern.FindAllSubmatch(bytesData, -1)
	links = make([]string, 0, len(matches))
	seen := make(map[string]bool)
	for _, match := range matches {
		link := string(match[1])
		if _, ok := seen[link]; !ok && link != "" {
			links = append(links, link)
			seen[link] = true
		}
	}

	return links, nil
}
