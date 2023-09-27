module github.com/test/testpackage

go 1.18

require (
	github.com/becheran/smock v0.0.0-20230927163458-39f3c48af310
	github.com/stretchr/testify v1.8.4
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/becheran/smock => ../../
