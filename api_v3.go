package googletranslate

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	u "net/url"
	"strings"
)

// TranslateApiV3 translates the given text from the source language to the target language
// using clients5.google.com/translate_a/t endpoint.
func TranslateApiV3(text, source, target string) (string, error) {
	// Trim leading and trailing whitespaces from the input text.
	text = strings.TrimSpace(text)

	// Return early if source and target languages are the same, or if the text is empty.
	if source == target || text == "" {
		return text, nil
	}

	// Set the source language to "auto" if not provided.
	if source == "" {
		source = "auto"
	}

	// Build the URL for the Google Translate API.
	url := "https://clients5.google.com/translate_a/t?client=dict-chrome-ex"
	url += "&sl=" + source
	url += "&tl=" + target
	url += "&q=" + u.QueryEscape(text)

	// Acquire a lock to prevent concurrent API calls.
	_ = sem.Acquire(1)
	r, err := http.Get(url)
	sem.Release(1)

	if err != nil {
		return "", err
	}

	// Read the response body.
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	// Unmarshal the JSON response.
	var result []interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	// Process the translated data in the response.
	if len(result) > 0 {
		inner := result[0]
		switch inner := inner.(type) {
		case []string:
			// Handle the case when source language is "auto".
			translated := inner[0]
			return translated, nil
		case string:
			// Handle the case when a specific source language is provided.
			return inner, nil
		}
	}

	// If no translated data is found in the response, return an error.
	return "", errors.New("no translated data in response")
}
