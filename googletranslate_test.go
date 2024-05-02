package googletranslate_test

import (
	_ "embed"
	"fmt"
	"strings"
	"testing"

	"github.com/spywiree/googletranslate"
	"github.com/stretchr/testify/assert"
)

//go:embed text.txt
var text string

func TestEndpoint1A(t *testing.T) {
	result, err := googletranslate.TranslateE1(text, "auto", "en")
	assert.NotEqual(t, "", result)
	assert.Equal(t, nil, err)
}

func TestEndpoint1B(t *testing.T) {
	result, err := googletranslate.TranslateE1(text, "pl", "en")
	assert.NotEqual(t, "", result)
	assert.Equal(t, nil, err)
}

func TestEndpoint2A(t *testing.T) {
	result, err := googletranslate.TranslateE2(text, "auto", "en")
	// assert.Equal(t, translated, result)
	assert.NotEqual(t, "", result)
	assert.Equal(t, nil, err)
}

func TestEndpoint2B(t *testing.T) {
	result, err := googletranslate.TranslateE2(text, "pl", "en")
	assert.NotEqual(t, "", result)
	assert.Equal(t, nil, err)
}

// Shorten to two paragraphs
var shortText = strings.Join(strings.Split(text, "\n")[:4], "\n")

func TestEndpoint3A(t *testing.T) {
	result, err := googletranslate.TranslateE3(shortText, "auto", "en")
	assert.NotEqual(t, "", result)
	assert.Equal(t, nil, err)
}

func TestEndpoint3B(t *testing.T) {
	fmt.Println(shortText)
	result, err := googletranslate.TranslateE3(shortText, "pl", "en")
	assert.NotEqual(t, "", result)
	assert.Equal(t, nil, err)
}
