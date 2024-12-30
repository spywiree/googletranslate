package googletranslate_test

import (
	_ "embed"
	"testing"

	"github.com/mitchellh/go-wordwrap"
	"github.com/spywiree/googletranslate/v2"
	langcodes "github.com/spywiree/langcodes"
)

//go:embed sample.txt
var sampleText string

// Shorten to two paragraphs
var shortSampleText = func() string {
	IndexN := func(s, substr string, n int) int {
		count := 0
		for i := 0; i <= len(s)-len(substr); i++ {
			if s[i:i+len(substr)] == substr {
				count++
				if count == n {
					return i
				}
			}
		}
		return -1
	}

	return sampleText[:IndexN(sampleText, "\n\n", 2)]
}()

func TestEndpoint1A(t *testing.T) {
	translated, err := googletranslate.TranslateE1(sampleText, langcodes.DETECT_LANGUAGE, langcodes.ENGLISH)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("\n\n" + wordwrap.WrapString(translated, 80))
}

func TestEndpoint1B(t *testing.T) {
	translated, err := googletranslate.TranslateE1(sampleText, langcodes.POLISH, langcodes.ENGLISH)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("\n\n" + wordwrap.WrapString(translated, 80))
}

func TestEndpoint2A(t *testing.T) {
	translated, err := googletranslate.TranslateE2(sampleText, langcodes.DETECT_LANGUAGE, langcodes.ENGLISH)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("\n\n" + wordwrap.WrapString(translated, 80))
}

func TestEndpoint2B(t *testing.T) {
	translated, err := googletranslate.TranslateE2(sampleText, langcodes.POLISH, langcodes.ENGLISH)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("\n\n" + wordwrap.WrapString(translated, 80))
}

func TestEndpoint3A(t *testing.T) {
	translated, err := googletranslate.TranslateE3(shortSampleText, langcodes.DETECT_LANGUAGE, langcodes.ENGLISH)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("\n\n" + wordwrap.WrapString(translated, 80))
}

func TestEndpoint3B(t *testing.T) {
	translated, err := googletranslate.TranslateE3(shortSampleText, langcodes.POLISH, langcodes.ENGLISH)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("\n\n" + wordwrap.WrapString(translated, 80))
}
