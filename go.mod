module github.com/mickahell/Gotismadex

go 1.23.8

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gorilla/mux v1.8.0
	github.com/stretchr/testify v1.8.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-openapi/analysis v0.21.4 // indirect
	github.com/go-openapi/errors v0.20.3 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/loads v0.21.2 // indirect
	github.com/go-openapi/runtime v0.24.1
	github.com/go-openapi/spec v0.20.7 // indirect
	github.com/go-openapi/strfmt v0.21.3 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/go-openapi/validate v0.22.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.10.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// Packages system

replace github.com/mickahell/Gotismadex/router => ../router

replace github.com/mickahell/Gotismadex/database => ../database

replace github.com/mickahell/Gotismadex/helpers => ../helpers

replace github.com/mickahell/Gotismadex/logger => ../logger

replace github.com/mickahell/Gotismadex/docs => ../docs

replace github.com/mickahell/Gotismadex/modules/launcher => ../launcher

replace github.com/mickahell/Gotismadex/modules/health => ../health
