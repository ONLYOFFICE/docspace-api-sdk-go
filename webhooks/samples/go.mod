module docspace-webhooks-sdk-go-samples

go 1.26

require github.com/ONLYOFFICE/docspace-api-sdk-go/webhooks v1.0.0

require gopkg.in/validator.v2 v2.0.1 // indirect

// The SDK is the module one directory up, not a published release.
replace github.com/ONLYOFFICE/docspace-api-sdk-go/webhooks => ..
