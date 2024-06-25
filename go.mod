module github.com/tiwarishubham635/go-calculator/v2

go 1.21.4

require github.com/stretchr/testify v1.9.0

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract (
    v2.0.1 // Published accidentally.
    v2.0.2 // Contains retractions only.
)
