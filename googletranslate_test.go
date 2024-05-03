package googletranslate_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/spywiree/googletranslate/v2"
	"github.com/spywiree/languagecodes"
	"github.com/stretchr/testify/assert"
)

//go:embed text.txt
var text string

func TestEndpoint1A(t *testing.T) {
	result, err := googletranslate.TranslateE1(text, languagecodes.DETECT_LANGUAGE, languagecodes.ENGLISH)
	assert.NotEqual(t, "", result)
	assert.Equal(t, nil, err)
}

func TestEndpoint1B(t *testing.T) {
	result, err := googletranslate.TranslateE1(text, languagecodes.POLISH, languagecodes.ENGLISH)
	assert.NotEqual(t, "", result)
	assert.Equal(t, nil, err)
}

func TestEndpoint2A(t *testing.T) {
	result, err := googletranslate.TranslateE2(text, languagecodes.DETECT_LANGUAGE, languagecodes.ENGLISH)
	assert.NotEqual(t, "", result)
	assert.Equal(t, nil, err)
}

func TestEndpoint2B(t *testing.T) {
	result, err := googletranslate.TranslateE2(text, languagecodes.POLISH, languagecodes.ENGLISH)
	assert.NotEqual(t, "", result)
	assert.Equal(t, nil, err)
}

// Shorten to two paragraphs
var shortText = strings.Join(strings.Split(text, "\n")[:4], "\n")

func TestEndpoint3A(t *testing.T) {
	result, err := googletranslate.TranslateE3(shortText, languagecodes.DETECT_LANGUAGE, languagecodes.ENGLISH)
	assert.NotEqual(t, "", result)
	assert.Equal(t, nil, err)
}

func TestEndpoint3B(t *testing.T) {
	result, err := googletranslate.TranslateE3(shortText, languagecodes.POLISH, languagecodes.ENGLISH)
	assert.NotEqual(t, "", result)
	assert.Equal(t, nil, err)
}
