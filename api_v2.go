package googletranslate

import (
	"errors"
	"net/http"
	u "net/url"
	"strings"

	"golang.org/x/net/html"
)

// TranslateApiV2 translates the given text from the source language to the target language
// using translate.google.com/m endpoint.
func TranslateApiV2(text, source, target string) (string, error) {
	text = strings.TrimSpace(text)

	// Return early if source and target languages are the same, or if the text is empty.
	if source == target || text == "" {
		return text, nil
	}

	// Check if the text length exceeds the maximum allowed (5000 characters).
	if len(text) <= 5000 {
		// Set the source language to "auto" if not provided.
		if source == "" {
			source = "auto"
		}

		// Build the URL for the Google Translate API.
		url := "https://translate.google.com/m"
		url += "?sl=" + source
		url += "&tl=" + target
		url += "&q=" + u.QueryEscape(text)

		// Acquire a lock to prevent concurrent API calls.
		sem.Acquire(1)
		r, err := http.Get(url)
		sem.Release(1)

		if err != nil {
			return "", err
		}

		// Check for rate limiting (429 Too Many Requests).
		if r.StatusCode == 429 {
			return "", errors.New("too many requests. You can make up to 5 requests per second and up to 200k per day")
		}

		// Use HTML tokenizer to parse the response.
		z := html.NewTokenizer(r.Body)

		var keyStr string
		var valueStr string
		var tagNameStr string

		// Iterate through HTML tokens to find the translated text.
		for {
			tt := z.Next()
			if tt == html.ErrorToken {
				break
			}

			// Check if the current token corresponds to the result container.
			if keyStr == "class" &&
				valueStr == "result-container" &&
				tagNameStr == "div" {
				return string(z.Text()), nil
			}

			// Extract attributes and tag name from the token.
			key, value, _ := z.TagAttr()
			keyStr = string(key)
			valueStr = string(value)

			tagName, _ := z.TagName()
			tagNameStr = string(tagName)
		}

		// If no translated data is found in the response, return an error.
		return "", errors.New("no translated data in response")
	} else {
		// Return an error if the maximum text length has been exceeded.
		return "", errors.New("the maximum text length of 5000 characters has been exceeded")
	}
}
