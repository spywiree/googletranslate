package googletranslate

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	u "net/url"

	"github.com/spywiree/languagecodes"
)

// TranslateE2 translates the given text from the source language to the target language
// using clients5.google.com/translate_a/t endpoint.
func TranslateE2(text string, source, target languagecodes.LanguageCode) (string, error) {
	if source == target || text == "" {
		return text, nil
	}

	url := "https://clients5.google.com/translate_a/t?client=dict-chrome-ex"
	url += "&sl=" + string(source)
	url += "&tl=" + string(target)
	url += "&q=" + u.QueryEscape(text)

	_ = sem.Acquire(1)
	defer sem.Release(1)

	r, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()

	if r.StatusCode != 200 {
		return "", errors.New(r.Status)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	var result []any
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	inner := result[0]
	switch inner := inner.(type) {
	case string:
		// Handle the case when source language is "auto".
		return inner, nil
	case []any:
		// Handle the case when a specific source language is provided.
		return inner[0].(string), nil
	}

	return "", NoTranslatedDataErr
}
