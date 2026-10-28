# \DocsAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConvertFile**](DocsAPI.md#ConvertFile) | **Post** /api/2.0/docs/converter | Convert a file
[**RunBuilderScript**](DocsAPI.md#RunBuilderScript) | **Post** /api/2.0/docs/builder | Run a document builder script



## ConvertFile

> FileWrapper ConvertFile(ctx).ConvertFileRequestDto(convertFileRequestDto).Execute()

Convert a file



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/convert-file/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	convertFileRequestDto := *openapiclient.NewConvertFileRequestDto(int32(1234), "pdf") // ConvertFileRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocsAPI.ConvertFile(context.Background()).ConvertFileRequestDto(convertFileRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocsAPI.ConvertFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConvertFile`: FileWrapper
	fmt.Fprintf(os.Stdout, "Response from `DocsAPI.ConvertFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConvertFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **convertFileRequestDto** | [**ConvertFileRequestDto**](ConvertFileRequestDto.md) |  | 

### Return type

[**FileWrapper**](FileWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunBuilderScript

> FileOperationWrapper RunBuilderScript(ctx).RunBuilderScriptRequestDto(runBuilderScriptRequestDto).Execute()

Run a document builder script



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/run-builder-script/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	runBuilderScriptRequestDto := *openapiclient.NewRunBuilderScriptRequestDto("builder.OpenFile(\"1234\"); Api.GetDocument().GetElement(0).AddText(Argument.title); builder.SaveFile(\"docx\", \"result.docx\"); builder.SaveFile(\"pdf\", \"result.pdf\"); builder.CloseFile();") // RunBuilderScriptRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocsAPI.RunBuilderScript(context.Background()).RunBuilderScriptRequestDto(runBuilderScriptRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocsAPI.RunBuilderScript``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunBuilderScript`: FileOperationWrapper
	fmt.Fprintf(os.Stdout, "Response from `DocsAPI.RunBuilderScript`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunBuilderScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runBuilderScriptRequestDto** | [**RunBuilderScriptRequestDto**](RunBuilderScriptRequestDto.md) |  | 

### Return type

[**FileOperationWrapper**](FileOperationWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

