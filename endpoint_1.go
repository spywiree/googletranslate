package googletranslate

import (
	"encoding/json"
	"fmt"
	"io"
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

	url := "https://translate.googleapis.com/translate_a/single?client=gtx&dt=t"
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

	if r.StatusCode != http.StatusOK {
		return "", HttpError(r.StatusCode)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	var translatedTexts []string
	var result []any

	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	if len(result) > 0 {
		inner := result[0]
		if inner, ok := inner.([]any); ok {
			for _, slice := range inner {
				for _, translated := range slice.([]any) {
					translatedTexts = append(
						translatedTexts,
						fmt.Sprintf("%v", translated),
					)
					break
				}
			}
			combined := strings.Join(translatedTexts, "")

			return combined, nil
		}
	}

	return "", NoTranslatedDataErr
}
