// (c) Copyright Ascensio System SIA 2026
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package docspace_api_sdk

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// DocsAPIService DocsAPI service
type DocsAPIService service

type ApiConvertFileRequest struct {
	ctx context.Context
	ApiService *DocsAPIService
	convertFileRequestDto *ConvertFileRequestDto
}

func (r ApiConvertFileRequest) ConvertFileRequestDto(convertFileRequestDto ConvertFileRequestDto) ApiConvertFileRequest {	r.convertFileRequestDto = &convertFileRequestDto
	return r
}

func (r ApiConvertFileRequest) Execute() (*FileWrapper, *http.Response, error) {
	return r.ApiService.ConvertFileExecute(r)
}

// ConvertFile Convert a file
//
// Converts a portal file into another format and saves the result as a new file, answering with that file. The  caller needs read access to a source that may be copied and the right to create files in `folderId`; without  `folderId` the result goes to the My documents section of the caller, so a guest, who has none, has to pass  it. A file in a private room or in a room with a watermark is not converted, as it would leave the room without  its protection, and no result is saved into such a room. `outputtype` is the only parameter a plain  conversion needs; the other fields are passed to the document service under its own names, while the source  format, name and cache key are taken from the file. The call is mutating and not idempotent: the source is kept,  and each call adds another file under a free name. It answers once the conversion is done, so a large file keeps  the request open. Only files stored in the portal itself are accepted. To convert while copying into another  folder, a third-party one included, use `POST api/2.0/files/file/{fileId}/copyas`; to convert into the portal's  own editable format beside the source, use `PUT api/2.0/files/file/{fileId}/checkconversion`. A missing or  unsupported `outputtype` is refused with 400, and so is a source the document service rejects: a wrong or  missing `password`, or a file too large to convert; retrying the same request will not help. Missing access or a  protected room is refused with 403 and an unknown file or folder with 404, both before any conversion starts. A  result that does not fit the storage quota of the portal, the room or the caller is refused with 402. A  conversion the document service fails is answered with 502: the service may be down or the file damaged; a retry  helps only in the first case.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/convert-file/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiConvertFileRequest
func (a *DocsAPIService) ConvertFile(ctx context.Context) ApiConvertFileRequest {
	return ApiConvertFileRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FileWrapper
func (a *DocsAPIService) ConvertFileExecute(r ApiConvertFileRequest) (*FileWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DocsAPIService.ConvertFile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/docs/converter"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.convertFileRequestDto
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyBearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ApiKeyBearer"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiRunBuilderScriptRequest struct {
	ctx context.Context
	ApiService *DocsAPIService
	runBuilderScriptRequestDto *RunBuilderScriptRequestDto
}

func (r ApiRunBuilderScriptRequest) RunBuilderScriptRequestDto(runBuilderScriptRequestDto RunBuilderScriptRequestDto) ApiRunBuilderScriptRequest {	r.runBuilderScriptRequestDto = &runBuilderScriptRequestDto
	return r
}

func (r ApiRunBuilderScriptRequest) Execute() (*FileOperationWrapper, *http.Response, error) {
	return r.ApiService.RunBuilderScriptExecute(r)
}

// RunBuilderScript Run a document builder script
//
// Queues a background run of a document builder script and answers with the operation just started. The script  names portal files by identifier - `builder.OpenFile(1234)` - where the document builder documentation writes  an address; only portal files can be opened, and none from a private room or a room with a watermark, nor is a  result saved into such a room. A script may call OpenFile and SaveFile at most 20 times each, counting every  call written in it. The caller needs edit access to every file the script opens, which may also be copied, and  to each file `outputs` replaces, which must not be locked or open in an editor, and the right to create files in  each folder a result is saved into. Each file the script saves goes where `outputs` says, keyed by the name  given to SaveFile: `fileId` stores it as a new version of that file, `folderId` as a new file there; an unlisted  file goes to `folderId` or to the folder of the opened file. Each entry of `outputs` names exactly one of  `fileId` and `folderId`; `title` goes only with `folderId`, since a replaced file keeps its title, and a result  stored over `fileId` has to be in that file's format. The call is not idempotent: each call is a new run. Poll  `GET api/2.0/files/fileops` with the returned `id` until the operation reports `finished`: `files` then lists  the saved files, a replaced one with its new version, including those saved before a failure, and `error` the  reason a failed build gave. A build that has not finished after 10 minutes is given up and reported as failed. A  caller runs one script at a time: a new request while an earlier one is queued or running is refused with 403,  so wait for `finished` before starting the next. For a plain format change use `POST api/2.0/docs/converter`.  What can be refused in advance fails at once with 400, 403 or 404.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/run-builder-script/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiRunBuilderScriptRequest
func (a *DocsAPIService) RunBuilderScript(ctx context.Context) ApiRunBuilderScriptRequest {
	return ApiRunBuilderScriptRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FileOperationWrapper
func (a *DocsAPIService) RunBuilderScriptExecute(r ApiRunBuilderScriptRequest) (*FileOperationWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileOperationWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DocsAPIService.RunBuilderScript")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/docs/builder"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.runBuilderScriptRequestDto
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyBearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ApiKeyBearer"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
