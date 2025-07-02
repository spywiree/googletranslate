module github.com/spywiree/googletranslate/v2

go 1.24.4

require golang.org/x/sync v0.15.0

require (
	github.com/spywiree/langcodes v1.1.0
	golang.org/x/net v0.41.0
)

require github.com/mitchellh/go-wordwrap v1.0.1

retract v2.0.0 // Incorrect go.mod.
