module github.com/spywiree/googletranslate/v2

go 1.22.2

require golang.org/x/sync v0.8.0

require (
	github.com/spywiree/langcodes v1.1.0
	github.com/stretchr/testify v1.9.0
	golang.org/x/net v0.29.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract v2.0.0 // Incorrect go.mod.
