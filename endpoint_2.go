package googletranslate

import (
	"encoding/json"
	"net/http"
	u "net/url"

	langcodes "github.com/spywiree/langcodes"
)

// TranslateE2 translates the given text from the source language to the target language
// using clients5.google.com/translate_a/t endpoint.
func TranslateE2(text string, source, target langcodes.LanguageCode) (string, error) {
	if source == target || text == "" {
		return text, nil
	}

	url := "https://clients5.google.com/translate_a/t?client=dict-chrome-ex" +
		"&sl=" + string(source) +
		"&tl=" + string(target) +
		"&q=" + u.QueryEscape(text)

	_ = sem.Acquire(1)
	defer sem.Release(1)

	r, err := http.Get(url) //#nosec G107
	if err != nil {
		return "", err
	}
	defer r.Body.Close() //nolint:errcheck

	if r.StatusCode != http.StatusOK {
		return "", HttpError(r.StatusCode)
	}

	dec := json.NewDecoder(r.Body)
	if source != langcodes.DETECT_LANGUAGE {
		var result []string
		err = dec.Decode(&result)
		if err != nil {
			return "", err
		}

		return result[0], nil
	} else {
		var result [][]string
		err = dec.Decode(&result)
		if err != nil {
			return "", err
		}

		return result[0][0], nil
	}
}
