package googletranslate

import (
	"net/http"
	u "net/url"

	langcodes "github.com/spywiree/langcodes"
	"golang.org/x/net/html"
)

// TranslateE3 translates the given text from the source language to the target language
// using translate.google.com/m endpoint.
func TranslateE3(text string, source, target langcodes.LanguageCode) (string, error) {
	if source == target || text == "" {
		return text, nil
	}

	if len(text) > 5000 {
		return "", MaxTextLengthExceededErr
	}

	url := "https://translate.google.com/m" +
		"?sl=" + string(source) +
		"&tl=" + string(target) +
		"&q=" + u.QueryEscape(text)

	_ = sem.Acquire(1)
	defer sem.Release(1)

	r, err := http.Get(url) //#nosec G107
	if err != nil {
		return "", err
	}
	defer r.Body.Close()

	if r.StatusCode == http.StatusTooManyRequests {
		return "", TooManyRequestsErr
	} else if r.StatusCode != http.StatusOK {
		return "", HttpError(r.StatusCode)
	}

	z := html.NewTokenizer(r.Body)
	var key, value, tagName []byte
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}

		if string(key) == "class" &&
			string(value) == "result-container" &&
			string(tagName) == "div" {
			return string(z.Text()), nil
		}

		key, value, _ = z.TagAttr()
		tagName, _ = z.TagName()
	}

	return "", NoTranslatedDataErr
}
