package googletranslate

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	u "net/url"
	"strings"
)

// TranslateApiV1 translates the given text from the source language to the target language
// using the translate.googleapis.com endpoint
func TranslateApiV1(text, source, target string) (string, error) {
	text = strings.TrimSpace(text)

	// Return early if source and target languages are the same, or if the text is empty.
	if source == target || text == "" {
		return text, nil
	}

	var translatedTexts []string
	var result []interface{}

	// Build the URL for the Google Translate API.
	url := "https://translate.googleapis.com/translate_a/single?client=gtx&dt=t"
	url += "&sl=" + source
	url += "&tl=" + target
	url += "&q=" + u.QueryEscape(text)

	// Acquire a lock to prevent concurrent API calls.
	sem.Acquire(1)
	r, err := http.Get(url)
	sem.Release(1)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()

	// Read the response body.
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	// Check if the response indicates a Bad Request (400) error.
	isBadRequest := strings.Contains(string(body), `<title>Error 400 (Bad Request)`)
	if isBadRequest {
		return "", errors.New("error 400 (Bad Request)")
	}

	// Unmarshal the JSON response.
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	// Process the translated data in the response.
	if len(result) > 0 {
		inner := result[0]
		switch inner := inner.(type) {
		case []interface{}:
			// Iterate through the response data and extract translated text.
			for _, slice := range inner {
				for _, translatedText := range slice.([]interface{}) {
					translatedTexts = append(translatedTexts, fmt.Sprintf("%v", translatedText))
					break
				}
			}
			combinedText := strings.Join(translatedTexts, "")

			return combinedText, nil
		}
	}

	// If no translated data is found in the response, return an error.
	return "", errors.New("no translated data in response")
}
