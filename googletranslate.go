package googletranslate

import (
	"github.com/spywiree/googletranslate/semaphore"
	"github.com/spywiree/languagecodes"
)

var sem = semaphore.NewSemaphore(256) // Default value is 256

// SetMaxConnections sets the maximum number of concurrent connections for Google Translate API.
// Specify -1 to disable the connection limit.
func SetMaxConnections(maxConcurrent int) {
	_ = sem.Resize(int64(maxConcurrent))
}

// Translate translates the given text from the source language to the target language using Google Translate API.
func Translate(text string, source, target languagecodes.LanguageCode) (string, error) {
	translated, err := TranslateE1(text, source, target)
	if err == nil {
		return translated, nil
	}

	translated, err = TranslateE2(text, source, target)
	if err == nil {
		return translated, nil
	}

	translated, err = TranslateE3(text, source, target)
	if err == nil {
		return translated, nil
	}

	return translated, err
}
