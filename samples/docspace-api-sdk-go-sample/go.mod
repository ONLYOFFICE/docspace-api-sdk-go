module github.com/ONLYOFFICE/docspace-api-sdk-go-sample

go 1.23.0

require github.com/ONLYOFFICE/docspace-api-sdk-go/v4 v4.0.0

require (
	golang.org/x/oauth2 v0.27.0 // indirect
	gopkg.in/validator.v2 v2.0.1 // indirect
)

replace github.com/ONLYOFFICE/docspace-api-sdk-go/v4 => ../..
