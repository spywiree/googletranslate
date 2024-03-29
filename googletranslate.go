package googletranslate

import (
	"github.com/spywiree/googletranslate/semaphore"
)

var sem = semaphore.NewSemaphore(256) // Default value is 256

// SetMaxConnections sets the maximum number of concurrent connections for Google Translate API.
// Specify -1 to disable the connection limit.
func SetMaxConnections(maxConcurrent int) {
	sem.Resize(int64(maxConcurrent))
}

// Translate translates the given text from the source language to the target language using Google Translate API.
func Translate(text, source, target string) (string, error) {
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
