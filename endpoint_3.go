package googletranslate

import (
	"errors"
	"net/http"
	u "net/url"

	"golang.org/x/net/html"
)

// TranslateE3 translates the given text from the source language to the target language
// using translate.google.com/m endpoint.
func TranslateE3(text string, source, target string) (string, error) {
	if source == target || text == "" {
		return text, nil
	}

	if len(text) > 5000 {
		return "", MaxTextLengthExceededErr
	}

	url := "https://translate.google.com/m"
	url += "?sl=" + string(source)
	url += "&tl=" + string(target)
	url += "&q=" + u.QueryEscape(text)

	_ = sem.Acquire(1)
	defer sem.Release(1)

	r, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()

	if r.StatusCode == 429 {
		return "", TooManyRequestsErr
	} else if r.StatusCode != 200 {
		return "", errors.New(r.Status)
	}

	z := html.NewTokenizer(r.Body)

	var keyStr string
	var valueStr string
	var tagNameStr string

	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}

		if keyStr == "class" &&
			valueStr == "result-container" &&
			tagNameStr == "div" {
			return string(z.Text()), nil
		}

		key, value, _ := z.TagAttr()
		keyStr = string(key)
		valueStr = string(value)

		tagName, _ := z.TagName()
		tagNameStr = string(tagName)
	}

	return "", NoTranslatedDataErr
}
