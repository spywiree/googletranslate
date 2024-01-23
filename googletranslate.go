package googletranslate

import (
	"github.com/spywiree/googletranslate/multimutex"
)

var mu = multimutex.NewMultiMutex(256) // Default value is 256
var called = false

// SetMaxConnections sets the maximum number of concurrent connections for Google Translate API.
// This must be done before any API call. Specify -1 to disable the connection limit.
func SetMaxConnections(maxConcurrent int) {
	if !called {
		mu = multimutex.NewMultiMutex(maxConcurrent)
	}
}

// Translate translates the given text from the source language to the target language using Google Translate API.
func Translate(text, source, target string) (string, error) {
	called = true

	translated, err := TranslateApiV1(text, source, target)
	if err == nil {
		return translated, nil
	}

	translated, err = TranslateApiV3(text, source, target)
	if err == nil {
		return translated, nil
	}

	translated, err = TranslateApiV2(text, source, target)
	if err == nil {
		return translated, nil
	}

	return translated, err
}
