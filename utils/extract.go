package utils

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/go-shiori/go-readability"
)

type ArticleData struct {
	Title   string `json:"title"`
	Byline  string `json:"byline"`
	Content string `json:"content"`
	Image   string `json:"image"`
}

func ExtractData(pageURL string) (*ArticleData, error) {
	parsedURL, err := url.Parse(pageURL)
	if err != nil {
		return nil, fmt.Errorf("invalid url: %w", err)
	}

	client := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   15 * time.Second,
				DualStack: true,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 15 * time.Second,
			ExpectContinueTimeout: 5 * time.Second,
		},
	}

	req, err := http.NewRequest("GET", pageURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; BookmarkApp/1.0)")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch url: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	article, err := readability.FromReader(resp.Body, parsedURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse article: %w", err)
	}

	return &ArticleData{
		Title:   article.Title,
		Byline:  article.Byline,
		Content: article.Content,
		Image:   article.Image,
	}, nil
}
