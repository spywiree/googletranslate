package googletranslate

import (
	"encoding/json"
	"net/http"
	u "net/url"
	"strings"

	langcodes "github.com/spywiree/langcodes"
)

// TranslateE1 translates the given text from the source language to the target language
// using the translate.googleapis.com endpoint
func TranslateE1(text string, source, target langcodes.LanguageCode) (string, error) {
	if source == target || text == "" {
		return text, nil
	}

	url := "https://translate.googleapis.com/translate_a/single?client=gtx&dt=t" +
		"&sl=" + string(source) +
		"&tl=" + string(target) +
		"&q=" + u.QueryEscape(text)

	_ = sem.Acquire(1)
	defer sem.Release(1)

	r, err := http.Get(url) //#nosec G107
	if err != nil {
		return "", err
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		return "", HttpError(r.StatusCode)
	}

	var result []any
	err = json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		return "", err
	}
	if len(result) == 0 {
		return "", NoTranslatedDataErr
	}

	result, ok := result[0].([]any)
	if !ok {
		return "", NoTranslatedDataErr
	}

	var sb strings.Builder
	for _, part := range result {
		part, ok := part.([]any)
		if !ok || len(part) == 0 {
			return "", NoTranslatedDataErr
		}

		s, ok := part[0].(string)
		if !ok {
			return "", NoTranslatedDataErr
		}

		sb.WriteString(s)
	}

	return sb.String(), nil
}
