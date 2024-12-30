module github.com/spywiree/googletranslate/v2

go 1.23.4

require golang.org/x/sync v0.10.0

require (
	github.com/spywiree/langcodes v1.1.0
	golang.org/x/net v0.33.0
)

require github.com/mitchellh/go-wordwrap v1.0.1

retract v2.0.0 // Incorrect go.mod.
