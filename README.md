# docspace_api_sdk

The ONLYOFFICE DocSpace SDK for Golang is a library that provides tools for integrating and managing DocSpace features within your applications. It simplifies interaction with the DocSpace API by offering ready-to-use methods and models.

For more information, please visit [https://helpdesk.onlyoffice.com/hc/en-us](https://helpdesk.onlyoffice.com/hc/en-us)

## Installation

Import the package in a Go file in your project and run `go mod tidy`:

```go
import (
    "context"
    "fmt"

    docspace_api_sdk "github.com/ONLYOFFICE/docspace-api-sdk-go/v4"
)
```

## Getting Started

Please follow the [installation](#installation) instruction and execute the following Go code:

```go
package main

import (
    "context"
    "fmt"

    docspace_api_sdk "github.com/ONLYOFFICE/docspace-api-sdk-go/v4"
)

func main() {
    cfg := docspace_api_sdk.NewConfiguration()
    cfg.Servers = docspace_api_sdk.ServerConfigurations{
        {
            URL: "https://your-docspace.onlyoffice.com",
            Description: "Default server",
        },
    }

    client := docspace_api_sdk.NewAPIClient(cfg)
    ctx := context.Background()

    ctx = context.WithValue(ctx, docspace_api_sdk.ContextAccessToken, "YOUR_ACCESS_TOKEN")
    ctx = context.WithValue(ctx, docspace_api_sdk.ContextAccessToken, "YOUR_ACCESS_TOKEN")

    
    aiAiApproveToolCallRequest :=  // 

    request := client.AIAIAPI.AiAiApproveToolCall(ctx, aiAiApproveToolCallRequest)

    response, httpResponse, err := request.Execute()
    if err != nil {
        fmt.Printf("Error when calling AIAIAPI.AiAiApproveToolCall: %v\n", err)
        if httpResponse != nil {
            fmt.Printf("HTTP response: %s\n", httpResponse.Status)
        }
        return
    }

    fmt.Printf("Response from `AiAiApproveToolCall`: %v\n", response)
    
}
```

## Documentation For Authorization


Authentication schemes defined for the API:
<a id="asc_auth_key"></a>
### asc_auth_key

- **Type**: API key
- **API key parameter name**: asc_auth_key
- **Location**: Cookie

Note, each API key must be added to a map of `map[string]APIKey` where the key is: asc_auth_key and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
    context.Background(),
    docspace_api_sdk.ContextAPIKeys,
    map[string]docspace_api_sdk.APIKey{
        "asc_auth_key": {Key: "API_KEY_STRING"},
    },
)
request := client.AIAIAPI.AiAiApproveToolCall(auth, /* aiAiApproveToolCallRequest */)
_, _, err := request.Execute()
```

<a id="Basic"></a>
### Basic

- **Type**: HTTP basic authentication

Example

```go
auth := context.WithValue(context.Background(), docspace_api_sdk.ContextBasicAuth, docspace_api_sdk.BasicAuth{
    UserName: "username",
    Password: "password",
})
request := client.AIAIAPI.AiAiApproveToolCall(auth, /* aiAiApproveToolCallRequest */)
_, _, err := request.Execute()
```

<a id="Bearer"></a>
### Bearer

- **Type**: Bearer authentication (JWT)

Example

```go
auth := context.WithValue(context.Background(), docspace_api_sdk.ContextAccessToken, "BEARER_TOKEN_STRING")
request := client.AIAIAPI.AiAiApproveToolCall(auth, /* aiAiApproveToolCallRequest */)
_, _, err := request.Execute()
```

<a id="ApiKeyBearer"></a>
### ApiKeyBearer

- **Type**: API key
- **API key parameter name**: ApiKeyBearer
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: ApiKeyBearer and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
    context.Background(),
    docspace_api_sdk.ContextAPIKeys,
    map[string]docspace_api_sdk.APIKey{
        "ApiKeyBearer": {Key: "API_KEY_STRING"},
    },
)
request := client.AIAIAPI.AiAiApproveToolCall(auth, /* aiAiApproveToolCallRequest */)
_, _, err := request.Execute()
```

<a id="OAuth2"></a>
### OAuth2

- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: {{authBaseUrl}}/oauth2/authorize
- **Token URL**: {{authBaseUrl}}/oauth2/token
- **Scopes**: 
 - **read**: Read access to protected resources
 - **write**: Write access to protected resources

Example

```go
auth := context.WithValue(context.Background(), docspace_api_sdk.ContextAccessToken, "ACCESSTOKENSTRING")
request := client.AIAIAPI.AiAiApproveToolCall(auth, /* aiAiApproveToolCallRequest */)
_, _, err := request.Execute()
```

<a id="OpenId"></a>
### OpenId

- **Type**: OpenId Connect
- **OpenId Connect URL**: {{authBaseUrl}}/.well-known/openid-configuration

<a id="cookieAuth"></a>
### cookieAuth

- **Type**: API key
- **API key parameter name**: asc_auth_key
- **Location**: Cookie

Note, each API key must be added to a map of `map[string]APIKey` where the key is: cookieAuth and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
    context.Background(),
    docspace_api_sdk.ContextAPIKeys,
    map[string]docspace_api_sdk.APIKey{
        "cookieAuth": {Key: "API_KEY_STRING"},
    },
)
request := client.AIAIAPI.AiAiApproveToolCall(auth, /* aiAiApproveToolCallRequest */)
_, _, err := request.Execute()
```

<a id="bearerAuth"></a>
### bearerAuth

- **Type**: Bearer authentication

Example

```go
auth := context.WithValue(context.Background(), docspace_api_sdk.ContextAccessToken, "BEARER_TOKEN_STRING")
request := client.AIAIAPI.AiAiApproveToolCall(auth, /* aiAiApproveToolCallRequest */)
_, _, err := request.Execute()
```

<a id="x-signature"></a>
### x-signature

- **Type**: API key
- **API key parameter name**: x-signature
- **Location**: Cookie

Note, each API key must be added to a map of `map[string]APIKey` where the key is: x-signature and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
    context.Background(),
    docspace_api_sdk.ContextAPIKeys,
    map[string]docspace_api_sdk.APIKey{
        "x-signature": {Key: "API_KEY_STRING"},
    },
)
request := client.AIAIAPI.AiAiApproveToolCall(auth, /* aiAiApproveToolCallRequest */)
_, _, err := request.Execute()
```


## Rate Limiting

All API responses may include the following rate limiting headers:

| Header | Description |
|--------|-------------|
| `X-RateLimit-Limit` | Sliding window rate limit: 1500 requests per minute per user/IP. |
| `X-RateLimit-Remaining` | Number of requests remaining in the current sliding window (1500 req/min). Concurrent limits also apply: 50 parallel GET requests, 15 parallel POST/PUT requests. |
| `X-RateLimit-Reset` | Unix timestamp (seconds) when the current sliding window rate limit resets. |
| `Retry-After` | Seconds to wait before retrying. Up to 60s for the sliding window (1500 req/min), up to 86400s for the daily POST/PUT limit (10000/day). |

### Documentation for API Endpoints

All URIs are relative to *https://your-docspace.onlyoffice.com*

### API endpoints tables

<details>
  <summary>AI</summary>

  <table>
    <tbody>
      <tr>
        <th>Method</th>
        <th>HTTP request</th>
        <th>Description</th>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>AIAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/AIAIAPI.md#aiaiapprovetoolcall"><strong>AiAiApproveToolCall</strong></a></td>
        <td><strong>Post</strong> /api/2.0/ai/ai/approve-tool-call</td>
        <td>Approve tool call</td>
      </tr>
      <tr>
        <td><a href="docs/AIAIAPI.md#aiaidenytoolcall"><strong>AiAiDenyToolCall</strong></a></td>
        <td><strong>Post</strong> /api/2.0/ai/ai/deny-tool-call</td>
        <td>Deny tool call</td>
      </tr>
      <tr>
        <td><a href="docs/AIAIAPI.md#aiairegeneratestream"><strong>AiAiRegenerateStream</strong></a></td>
        <td><strong>Post</strong> /api/2.0/ai/ai/regenerate-stream</td>
        <td>Regenerate stream</td>
      </tr>
      <tr>
        <td><a href="docs/AIAIAPI.md#aiaisend"><strong>AiAiSend</strong></a></td>
        <td><strong>Post</strong> /api/2.0/ai/ai/send</td>
        <td>Send</td>
      </tr>
      <tr>
        <td><a href="docs/AIAIAPI.md#aiaisendcustom"><strong>AiAiSendCustom</strong></a></td>
        <td><strong>Post</strong> /api/2.0/ai/ai/send-custom</td>
        <td>Send custom</td>
      </tr>
      <tr>
        <td><a href="docs/AIAIAPI.md#aiaisendwithstream"><strong>AiAiSendWithStream</strong></a></td>
        <td><strong>Post</strong> /api/2.0/ai/ai/send-with-stream</td>
        <td>Send with stream</td>
      </tr>
      <tr>
        <td><a href="docs/AIAIAPI.md#aiaisendwithstreamopenai"><strong>AiAiSendWithStreamOpenAI</strong></a></td>
        <td><strong>Post</strong> /api/2.0/ai/ai/send-with-stream-openai</td>
        <td>Send with stream open ai</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>AgentsAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/AIAgentsAPI.md#aiagentscreate"><strong>AiAgentsCreate</strong></a></td>
        <td><strong>Post</strong> /api/2.0/ai/agents</td>
        <td>Create an agent</td>
      </tr>
      <tr>
        <td><a href="docs/AIAgentsAPI.md#aiagentsdelete"><strong>AiAgentsDelete</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/ai/agents/{id}</td>
        <td>Delete an agent</td>
      </tr>
      <tr>
        <td><a href="docs/AIAgentsAPI.md#aiagentsget"><strong>AiAgentsGet</strong></a></td>
        <td><strong>Get</strong> /api/2.0/ai/agents/{id}</td>
        <td>Get an agent</td>
      </tr>
      <tr>
        <td><a href="docs/AIAgentsAPI.md#aiagentslist"><strong>AiAgentsList</strong></a></td>
        <td><strong>Get</strong> /api/2.0/ai/agents</td>
        <td>List agents</td>
      </tr>
      <tr>
        <td><a href="docs/AIAgentsAPI.md#aiagentsnews"><strong>AiAgentsNews</strong></a></td>
        <td><strong>Get</strong> /api/2.0/ai/agents/news</td>
        <td>List agent news items</td>
      </tr>
      <tr>
        <td><a href="docs/AIAgentsAPI.md#aiagentsresetquota"><strong>AiAgentsResetQuota</strong></a></td>
        <td><strong>Put</strong> /api/2.0/ai/agents/resetquota</td>
        <td>Reset agents' quota</td>
      </tr>
      <tr>
        <td><a href="docs/AIAgentsAPI.md#aiagentsupdate"><strong>AiAgentsUpdate</strong></a></td>
        <td><strong>Put</strong> /api/2.0/ai/agents/{id}</td>
        <td>Update an agent</td>
      </tr>
      <tr>
        <td><a href="docs/AIAgentsAPI.md#aiagentsupdatequota"><strong>AiAgentsUpdateQuota</strong></a></td>
        <td><strong>Put</strong> /api/2.0/ai/agents/agentquota</td>
        <td>Update agents' quota</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>AssignmentsAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/AIAssignmentsAPI.md#aiassignmentsassign"><strong>AiAssignmentsAssign</strong></a></td>
        <td><strong>Put</strong> /api/2.0/ai/assignments/assign</td>
        <td>Assign</td>
      </tr>
      <tr>
        <td><a href="docs/AIAssignmentsAPI.md#aiassignmentsbulkassign"><strong>AiAssignmentsBulkAssign</strong></a></td>
        <td><strong>Put</strong> /api/2.0/ai/assignments/bulk-assign</td>
        <td>Bulk assign</td>
      </tr>
      <tr>
        <td><a href="docs/AIAssignmentsAPI.md#aiassignmentscascadeprofiledelete"><strong>AiAssignmentsCascadeProfileDelete</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/ai/assignments/cascade-profile-delete</td>
        <td>Cascade profile delete</td>
      </tr>
      <tr>
        <td><a href="docs/AIAssignmentsAPI.md#aiassignmentsgetallassignments"><strong>AiAssignmentsGetAllAssignments</strong></a></td>
        <td><strong>Get</strong> /api/2.0/ai/assignments/get-all-assignments</td>
        <td>Get all assignments</td>
      </tr>
      <tr>
        <td><a href="docs/AIAssignmentsAPI.md#aiassignmentsgetassignment"><strong>AiAssignmentsGetAssignment</strong></a></td>
        <td><strong>Get</strong> /api/2.0/ai/assignments/get-assignment</td>
        <td>Get assignment</td>
      </tr>
      <tr>
        <td><a href="docs/AIAssignmentsAPI.md#aiassignmentsresolveforaction"><strong>AiAssignmentsResolveForAction</strong></a></td>
        <td><strong>Get</strong> /api/2.0/ai/assignments/resolve-for-action</td>
        <td>Resolve for action</td>
      </tr>
      <tr>
        <td><a href="docs/AIAssignmentsAPI.md#aiassignmentstryresolveforaction"><strong>AiAssignmentsTryResolveForAction</strong></a></td>
        <td><strong>Get</strong> /api/2.0/ai/assignments/try-resolve-for-action</td>
        <td>Try resolve for action</td>
      </tr>
      <tr>
        <td><a href="docs/AIAssignmentsAPI.md#aiassignmentsunassign"><strong>AiAssignmentsUnassign</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/ai/assignments/unassign</td>
        <td>Unassign</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>AttachmentsAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/AIAttachmentsAPI.md#aiattachmentsdelete"><strong>AiAttachmentsDelete</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/ai/attachments/delete</td>
        <td>Delete</td>
      </tr>
      <tr>
        <td><a href="docs/AIAttachmentsAPI.md#aiattachmentsdeletemany"><strong>AiAttachmentsDeleteMany</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/ai/attachments/delete-many</td>
        <td>Delete many</td>
      </tr>
      <tr>
        <td><a href="docs/AIAttachmentsAPI.md#aiattachmentsget"><strong>AiAttachmentsGet</strong></a></td>
        <td><strong>Post</strong> /api/2.0/ai/attachments/get</td>
        <td>Get</td>
      </tr>
      <tr>
        <td><a href="docs/AIAttachmentsAPI.md#aiattachmentsgetmany"><strong>AiAttachmentsGetMany</strong></a></td>
        <td><strong>Post</strong> /api/2.0/ai/attachments/get-many</td>
        <td>Get many</td>
      </tr>
      <tr>
        <td><a href="docs/AIAttachmentsAPI.md#aiattachmentslinktomessage"><strong>AiAttachmentsLinkToMessage</strong></a></td>
        <td><strong>Post</strong> /api/2.0/ai/attachments/link-to-message</td>
        <td>Link to message</td>
      </tr>
      <tr>
        <td><a href="docs/AIAttachmentsAPI.md#aiattachmentssavefile"><strong>AiAttachmentsSaveFile</strong></a></td>
        <td><strong>Post</strong> /api/2.0/ai/attachments/save-file</td>
        <td>Save file</td>
      </tr>
      <tr>
        <td><a href="docs/AIAttachmentsAPI.md#aiattachmentssavefilesmany"><strong>AiAttachmentsSaveFilesMany</strong></a></td>
        <td><strong>Post</strong> /api/2.0/ai/attachments/save-files-many</td>
        <td>Save files many</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>EditorToolsAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/AIEditorToolsAPI.md#aieditortoolscall"><strong>AiEditorToolsCall</strong></a></td>
        <td><strong>Post</strong> /api/2.0/ai/editor-tools/call</td>
        <td>Execute a DocSpace tool on behalf of the editor AI plugin</td>
      </tr>
      <tr>
        <td><a href="docs/AIEditorToolsAPI.md#aieditortoolslist"><strong>AiEditorToolsList</strong></a></td>
        <td><strong>Get</strong> /api/2.0/ai/editor-tools/list</td>
        <td>Sanitized DocSpace tool catalog for the editor AI plugin</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>ExportAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/AIExportAPI.md#aiexporttexttodocx"><strong>AiExportTextToDocx</strong></a></td>
        <td><strong>Post</strong> /api/2.0/ai/text-to-docx</td>
        <td>Start markdown → docx export</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>OpenAIPassthroughAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/AIOpenAIPassthroughAPI.md#aiopenaichatcompletions"><strong>AiOpenaiChatCompletions</strong></a></td>
        <td><strong>Post</strong> /api/2.0/ai/openai/{profileId}/v1/chat/completions</td>
        <td>OpenAI-compatible chat completions proxied to the profile's provider</td>
      </tr>
      <tr>
        <td><a href="docs/AIOpenAIPassthroughAPI.md#aiopenaiimagesgenerations"><strong>AiOpenaiImagesGenerations</strong></a></td>
        <td><strong>Post</strong> /api/2.0/ai/openai/{profileId}/v1/images/generations</td>
        <td>OpenAI-compatible image generation proxied to the profile's provider</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>PreferencesAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/AIPreferencesAPI.md#aipreferencescleardeepmode"><strong>AiPreferencesClearDeepMode</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/ai/preferences/clear-deep-mode</td>
        <td>Clear deep mode</td>
      </tr>
      <tr>
        <td><a href="docs/AIPreferencesAPI.md#aipreferencesgetdeepmode"><strong>AiPreferencesGetDeepMode</strong></a></td>
        <td><strong>Get</strong> /api/2.0/ai/preferences/get-deep-mode</td>
        <td>Get deep mode</td>
      </tr>
      <tr>
        <td><a href="docs/AIPreferencesAPI.md#aipreferencesisdeepmodeset"><strong>AiPreferencesIsDeepModeSet</strong></a></td>
        <td><strong>Get</strong> /api/2.0/ai/preferences/is-deep-mode-set</td>
        <td>Is deep mode set</td>
      </tr>
      <tr>
        <td><a href="docs/AIPreferencesAPI.md#aipreferencessetdeepmode"><strong>AiPreferencesSetDeepMode</strong></a></td>
        <td><strong>Put</strong> /api/2.0/ai/preferences/set-deep-mode</td>
        <td>Set deep mode</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>ProfilesAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/AIProfilesAPI.md#aiprofilescreate"><strong>AiProfilesCreate</strong></a></td>
        <td><strong>Post</strong> /api/2.0/ai/profiles/create</td>
        <td>Create</td>
      </tr>
      <tr>
        <td><a href="docs/AIProfilesAPI.md#aiprofilesdelete"><strong>AiProfilesDelete</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/ai/profiles/delete</td>
        <td>Delete</td>
      </tr>
      <tr>
        <td><a href="docs/AIProfilesAPI.md#aiprofilesgetbyid"><strong>AiProfilesGetById</strong></a></td>
        <td><strong>Get</strong> /api/2.0/ai/profiles/get-by-id</td>
        <td>Get by id</td>
      </tr>
      <tr>
        <td><a href="docs/AIProfilesAPI.md#aiprofileslist"><strong>AiProfilesList</strong></a></td>
        <td><strong>Get</strong> /api/2.0/ai/profiles/list</td>
        <td>List</td>
      </tr>
      <tr>
        <td><a href="docs/AIProfilesAPI.md#aiprofileslistmodels"><strong>AiProfilesListModels</strong></a></td>
        <td><strong>Get</strong> /api/2.0/ai/profiles/list-models</td>
        <td>List models</td>
      </tr>
      <tr>
        <td><a href="docs/AIProfilesAPI.md#aiprofileslistprovidermodels"><strong>AiProfilesListProviderModels</strong></a></td>
        <td><strong>Post</strong> /api/2.0/ai/profiles/list-provider-models</td>
        <td>List provider models</td>
      </tr>
      <tr>
        <td><a href="docs/AIProfilesAPI.md#aiprofilestestconnection"><strong>AiProfilesTestConnection</strong></a></td>
        <td><strong>Post</strong> /api/2.0/ai/profiles/test-connection</td>
        <td>Test connection</td>
      </tr>
      <tr>
        <td><a href="docs/AIProfilesAPI.md#aiprofilesupdate"><strong>AiProfilesUpdate</strong></a></td>
        <td><strong>Put</strong> /api/2.0/ai/profiles/update</td>
        <td>Update</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>PromptsAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/AIPromptsAPI.md#aipromptscreate"><strong>AiPromptsCreate</strong></a></td>
        <td><strong>Post</strong> /api/2.0/ai/prompts/create</td>
        <td>Create</td>
      </tr>
      <tr>
        <td><a href="docs/AIPromptsAPI.md#aipromptscreatefolder"><strong>AiPromptsCreateFolder</strong></a></td>
        <td><strong>Post</strong> /api/2.0/ai/prompts/create-folder</td>
        <td>Create folder</td>
      </tr>
      <tr>
        <td><a href="docs/AIPromptsAPI.md#aipromptsdelete"><strong>AiPromptsDelete</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/ai/prompts/delete</td>
        <td>Delete</td>
      </tr>
      <tr>
        <td><a href="docs/AIPromptsAPI.md#aipromptsdeletefolder"><strong>AiPromptsDeleteFolder</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/ai/prompts/delete-folder</td>
        <td>Delete folder</td>
      </tr>
      <tr>
        <td><a href="docs/AIPromptsAPI.md#aipromptsexport"><strong>AiPromptsExport</strong></a></td>
        <td><strong>Get</strong> /api/2.0/ai/prompts/export</td>
        <td>Export</td>
      </tr>
      <tr>
        <td><a href="docs/AIPromptsAPI.md#aipromptsgetbyid"><strong>AiPromptsGetById</strong></a></td>
        <td><strong>Get</strong> /api/2.0/ai/prompts/get-by-id</td>
        <td>Get by id</td>
      </tr>
      <tr>
        <td><a href="docs/AIPromptsAPI.md#aipromptsgetfolderbyid"><strong>AiPromptsGetFolderById</strong></a></td>
        <td><strong>Get</strong> /api/2.0/ai/prompts/get-folder-by-id</td>
        <td>Get folder by id</td>
      </tr>
      <tr>
        <td><a href="docs/AIPromptsAPI.md#aipromptsimportbundle"><strong>AiPromptsImportBundle</strong></a></td>
        <td><strong>Post</strong> /api/2.0/ai/prompts/import-bundle</td>
        <td>Import bundle</td>
      </tr>
      <tr>
        <td><a href="docs/AIPromptsAPI.md#aipromptslist"><strong>AiPromptsList</strong></a></td>
        <td><strong>Get</strong> /api/2.0/ai/prompts/list</td>
        <td>List</td>
      </tr>
      <tr>
        <td><a href="docs/AIPromptsAPI.md#aipromptslistfolders"><strong>AiPromptsListFolders</strong></a></td>
        <td><strong>Get</strong> /api/2.0/ai/prompts/list-folders</td>
        <td>List folders</td>
      </tr>
      <tr>
        <td><a href="docs/AIPromptsAPI.md#aipromptsmove"><strong>AiPromptsMove</strong></a></td>
        <td><strong>Put</strong> /api/2.0/ai/prompts/move</td>
        <td>Move</td>
      </tr>
      <tr>
        <td><a href="docs/AIPromptsAPI.md#aipromptsrenamefolder"><strong>AiPromptsRenameFolder</strong></a></td>
        <td><strong>Put</strong> /api/2.0/ai/prompts/rename-folder</td>
        <td>Rename folder</td>
      </tr>
      <tr>
        <td><a href="docs/AIPromptsAPI.md#aipromptsupdate"><strong>AiPromptsUpdate</strong></a></td>
        <td><strong>Put</strong> /api/2.0/ai/prompts/update</td>
        <td>Update</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>AISettingsAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/AISettingsAPI.md#aisettingsget"><strong>AiSettingsGet</strong></a></td>
        <td><strong>Get</strong> /api/2.0/ai/config</td>
        <td>Get AI settings</td>
      </tr>
      <tr>
        <td><a href="docs/AISettingsAPI.md#aisettingsgetuser"><strong>AiSettingsGetUser</strong></a></td>
        <td><strong>Get</strong> /api/2.0/ai/config/user</td>
        <td>Get user AI settings</td>
      </tr>
      <tr>
        <td><a href="docs/AISettingsAPI.md#aisettingsgetvectorization"><strong>AiSettingsGetVectorization</strong></a></td>
        <td><strong>Get</strong> /api/2.0/ai/config/vectorization</td>
        <td>Get vectorization settings</td>
      </tr>
      <tr>
        <td><a href="docs/AISettingsAPI.md#aisettingssetuser"><strong>AiSettingsSetUser</strong></a></td>
        <td><strong>Put</strong> /api/2.0/ai/config/user</td>
        <td>Update user AI settings</td>
      </tr>
      <tr>
        <td><a href="docs/AISettingsAPI.md#aisettingssetvectorization"><strong>AiSettingsSetVectorization</strong></a></td>
        <td><strong>Put</strong> /api/2.0/ai/config/vectorization</td>
        <td>Update vectorization settings</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>ThreadsAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/AIThreadsAPI.md#aithreadsappendusermessage"><strong>AiThreadsAppendUserMessage</strong></a></td>
        <td><strong>Post</strong> /api/2.0/ai/threads/append-user-message</td>
        <td>Append user message</td>
      </tr>
      <tr>
        <td><a href="docs/AIThreadsAPI.md#aithreadsclearmessages"><strong>AiThreadsClearMessages</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/ai/threads/clear-messages</td>
        <td>Clear messages</td>
      </tr>
      <tr>
        <td><a href="docs/AIThreadsAPI.md#aithreadscreate"><strong>AiThreadsCreate</strong></a></td>
        <td><strong>Post</strong> /api/2.0/ai/threads/create</td>
        <td>Create</td>
      </tr>
      <tr>
        <td><a href="docs/AIThreadsAPI.md#aithreadsdelete"><strong>AiThreadsDelete</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/ai/threads/delete</td>
        <td>Delete</td>
      </tr>
      <tr>
        <td><a href="docs/AIThreadsAPI.md#aithreadsdeletemessage"><strong>AiThreadsDeleteMessage</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/ai/threads/delete-message</td>
        <td>Delete message</td>
      </tr>
      <tr>
        <td><a href="docs/AIThreadsAPI.md#aithreadsgetbyid"><strong>AiThreadsGetById</strong></a></td>
        <td><strong>Get</strong> /api/2.0/ai/threads/get-by-id</td>
        <td>Get by id</td>
      </tr>
      <tr>
        <td><a href="docs/AIThreadsAPI.md#aithreadsgetmessagebyid"><strong>AiThreadsGetMessageById</strong></a></td>
        <td><strong>Get</strong> /api/2.0/ai/threads/get-message-by-id</td>
        <td>Get message by id</td>
      </tr>
      <tr>
        <td><a href="docs/AIThreadsAPI.md#aithreadslist"><strong>AiThreadsList</strong></a></td>
        <td><strong>Get</strong> /api/2.0/ai/threads/list</td>
        <td>List</td>
      </tr>
      <tr>
        <td><a href="docs/AIThreadsAPI.md#aithreadsopenorcreate"><strong>AiThreadsOpenOrCreate</strong></a></td>
        <td><strong>Post</strong> /api/2.0/ai/threads/open-or-create</td>
        <td>Open or create</td>
      </tr>
      <tr>
        <td><a href="docs/AIThreadsAPI.md#aithreadsreadmessages"><strong>AiThreadsReadMessages</strong></a></td>
        <td><strong>Get</strong> /api/2.0/ai/threads/read-messages</td>
        <td>Read messages</td>
      </tr>
      <tr>
        <td><a href="docs/AIThreadsAPI.md#aithreadsregeneratetitle"><strong>AiThreadsRegenerateTitle</strong></a></td>
        <td><strong>Post</strong> /api/2.0/ai/threads/regenerate-title</td>
        <td>Regenerate title</td>
      </tr>
      <tr>
        <td><a href="docs/AIThreadsAPI.md#aithreadsrename"><strong>AiThreadsRename</strong></a></td>
        <td><strong>Put</strong> /api/2.0/ai/threads/rename</td>
        <td>Rename</td>
      </tr>
      <tr>
        <td><a href="docs/AIThreadsAPI.md#aithreadstouch"><strong>AiThreadsTouch</strong></a></td>
        <td><strong>Post</strong> /api/2.0/ai/threads/touch</td>
        <td>Touch</td>
      </tr>
      <tr>
        <td><a href="docs/AIThreadsAPI.md#aithreadsupdatemessage"><strong>AiThreadsUpdateMessage</strong></a></td>
        <td><strong>Put</strong> /api/2.0/ai/threads/update-message</td>
        <td>Update message</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>ToolsAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/AIToolsAPI.md#aitoolsaddcustomserver"><strong>AiToolsAddCustomServer</strong></a></td>
        <td><strong>Post</strong> /api/2.0/ai/tools/add-custom-server</td>
        <td>Add custom server</td>
      </tr>
      <tr>
        <td><a href="docs/AIToolsAPI.md#aitoolsgetallowalways"><strong>AiToolsGetAllowAlways</strong></a></td>
        <td><strong>Get</strong> /api/2.0/ai/tools/get-allow-always</td>
        <td>Get allow always</td>
      </tr>
      <tr>
        <td><a href="docs/AIToolsAPI.md#aitoolsgetcustomserver"><strong>AiToolsGetCustomServer</strong></a></td>
        <td><strong>Get</strong> /api/2.0/ai/tools/get-custom-server</td>
        <td>Get custom server</td>
      </tr>
      <tr>
        <td><a href="docs/AIToolsAPI.md#aitoolsgetdisabled"><strong>AiToolsGetDisabled</strong></a></td>
        <td><strong>Get</strong> /api/2.0/ai/tools/get-disabled</td>
        <td>Get disabled</td>
      </tr>
      <tr>
        <td><a href="docs/AIToolsAPI.md#aitoolsisallowalways"><strong>AiToolsIsAllowAlways</strong></a></td>
        <td><strong>Get</strong> /api/2.0/ai/tools/is-allow-always</td>
        <td>Is allow always</td>
      </tr>
      <tr>
        <td><a href="docs/AIToolsAPI.md#aitoolsistooldisabled"><strong>AiToolsIsToolDisabled</strong></a></td>
        <td><strong>Get</strong> /api/2.0/ai/tools/is-tool-disabled</td>
        <td>Is tool disabled</td>
      </tr>
      <tr>
        <td><a href="docs/AIToolsAPI.md#aitoolslistcustomservers"><strong>AiToolsListCustomServers</strong></a></td>
        <td><strong>Get</strong> /api/2.0/ai/tools/list-custom-servers</td>
        <td>List custom servers</td>
      </tr>
      <tr>
        <td><a href="docs/AIToolsAPI.md#aitoolslistsystemtools"><strong>AiToolsListSystemTools</strong></a></td>
        <td><strong>Get</strong> /api/2.0/ai/tools/list-system-tools</td>
        <td>List system tools</td>
      </tr>
      <tr>
        <td><a href="docs/AIToolsAPI.md#aitoolsremovecustomserver"><strong>AiToolsRemoveCustomServer</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/ai/tools/remove-custom-server</td>
        <td>Remove custom server</td>
      </tr>
      <tr>
        <td><a href="docs/AIToolsAPI.md#aitoolsreplaceallcustomservers"><strong>AiToolsReplaceAllCustomServers</strong></a></td>
        <td><strong>Put</strong> /api/2.0/ai/tools/replace-all-custom-servers</td>
        <td>Replace all custom servers</td>
      </tr>
      <tr>
        <td><a href="docs/AIToolsAPI.md#aitoolssetallowalways"><strong>AiToolsSetAllowAlways</strong></a></td>
        <td><strong>Put</strong> /api/2.0/ai/tools/set-allow-always</td>
        <td>Set allow always</td>
      </tr>
      <tr>
        <td><a href="docs/AIToolsAPI.md#aitoolssetdisabled"><strong>AiToolsSetDisabled</strong></a></td>
        <td><strong>Put</strong> /api/2.0/ai/tools/set-disabled</td>
        <td>Set disabled</td>
      </tr>
      <tr>
        <td><a href="docs/AIToolsAPI.md#aitoolsupdatecustomserver"><strong>AiToolsUpdateCustomServer</strong></a></td>
        <td><strong>Put</strong> /api/2.0/ai/tools/update-custom-server</td>
        <td>Update custom server</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>VectorizationAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/AIVectorizationAPI.md#aivectorizationstarttask"><strong>AiVectorizationStartTask</strong></a></td>
        <td><strong>Post</strong> /api/2.0/ai/vectorization/tasks</td>
        <td>Start a vectorization task</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>WebSearchAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/AIWebSearchAPI.md#aiwebsearchclear"><strong>AiWebSearchClear</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/ai/web-search/clear</td>
        <td>Clear</td>
      </tr>
      <tr>
        <td><a href="docs/AIWebSearchAPI.md#aiwebsearchconfigure"><strong>AiWebSearchConfigure</strong></a></td>
        <td><strong>Put</strong> /api/2.0/ai/web-search/configure</td>
        <td>Configure</td>
      </tr>
      <tr>
        <td><a href="docs/AIWebSearchAPI.md#aiwebsearchgetactiveconfig"><strong>AiWebSearchGetActiveConfig</strong></a></td>
        <td><strong>Get</strong> /api/2.0/ai/web-search/get-active-config</td>
        <td>Get active config</td>
      </tr>
      <tr>
        <td><a href="docs/AIWebSearchAPI.md#aiwebsearchisconfigured"><strong>AiWebSearchIsConfigured</strong></a></td>
        <td><strong>Get</strong> /api/2.0/ai/web-search/is-configured</td>
        <td>Is configured</td>
      </tr>
      <tr>
        <td><a href="docs/AIWebSearchAPI.md#aiwebsearchpassthroughcontents"><strong>AiWebSearchPassthroughContents</strong></a></td>
        <td><strong>Post</strong> /api/2.0/ai/websearch/v1/contents</td>
        <td>Web page contents proxied to the portal's active web-search provider</td>
      </tr>
      <tr>
        <td><a href="docs/AIWebSearchAPI.md#aiwebsearchpassthroughsearch"><strong>AiWebSearchPassthroughSearch</strong></a></td>
        <td><strong>Post</strong> /api/2.0/ai/websearch/v1/search</td>
        <td>Web search proxied to the portal's active web-search provider</td>
      </tr>
      <tr>
        <td><a href="docs/AIWebSearchAPI.md#aiwebsearchsetactiveconfig"><strong>AiWebSearchSetActiveConfig</strong></a></td>
        <td><strong>Put</strong> /api/2.0/ai/web-search/set-active-config</td>
        <td>Set active config</td>
      </tr>
      <tr>
        <td><a href="docs/AIWebSearchAPI.md#aiwebsearchtestconnection"><strong>AiWebSearchTestConnection</strong></a></td>
        <td><strong>Post</strong> /api/2.0/ai/web-search/test-connection</td>
        <td>Test connection</td>
      </tr>
    </tbody>
  </table>

</details>
<details>
  <summary>ApiKeys</summary>

  <table>
    <tbody>
      <tr>
        <th>Method</th>
        <th>HTTP request</th>
        <th>Description</th>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>ApiKeysAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/ApiKeysAPI.md#createapikey"><strong>CreateApiKey</strong></a></td>
        <td><strong>Post</strong> /api/2.0/keys</td>
        <td>Create a user API key</td>
      </tr>
      <tr>
        <td><a href="docs/ApiKeysAPI.md#deleteapikey"><strong>DeleteApiKey</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/keys/{keyId}</td>
        <td>Delete a user API key</td>
      </tr>
      <tr>
        <td><a href="docs/ApiKeysAPI.md#getallpermissions"><strong>GetAllPermissions</strong></a></td>
        <td><strong>Get</strong> /api/2.0/keys/permissions</td>
        <td>Get API key permissions</td>
      </tr>
      <tr>
        <td><a href="docs/ApiKeysAPI.md#getapikey"><strong>GetApiKey</strong></a></td>
        <td><strong>Get</strong> /api/2.0/keys/@self</td>
        <td>Get current user's API key</td>
      </tr>
      <tr>
        <td><a href="docs/ApiKeysAPI.md#getapikeys"><strong>GetApiKeys</strong></a></td>
        <td><strong>Get</strong> /api/2.0/keys</td>
        <td>Get current user's API keys</td>
      </tr>
      <tr>
        <td><a href="docs/ApiKeysAPI.md#updateapikey"><strong>UpdateApiKey</strong></a></td>
        <td><strong>Put</strong> /api/2.0/keys/{keyId}</td>
        <td>Update an API key</td>
      </tr>
    </tbody>
  </table>

</details>
<details>
  <summary>Apps</summary>

  <table>
    <tbody>
      <tr>
        <th>Method</th>
        <th>HTTP request</th>
        <th>Description</th>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>AppsAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/AppsAPI.md#get"><strong>Get</strong></a></td>
        <td><strong>Get</strong> /api/2.0/apps/{id}</td>
        <td>Get a single app</td>
      </tr>
      <tr>
        <td><a href="docs/AppsAPI.md#getall"><strong>GetAll</strong></a></td>
        <td><strong>Get</strong> /api/2.0/apps</td>
        <td>Get all apps</td>
      </tr>
      <tr>
        <td><a href="docs/AppsAPI.md#getsettings"><strong>GetSettings</strong></a></td>
        <td><strong>Get</strong> /api/2.0/apps/{id}/settings</td>
        <td>Get app settings</td>
      </tr>
      <tr>
        <td><a href="docs/AppsAPI.md#setenabled"><strong>SetEnabled</strong></a></td>
        <td><strong>Put</strong> /api/2.0/apps/{id}/enabled</td>
        <td>Enable or disable an app</td>
      </tr>
      <tr>
        <td><a href="docs/AppsAPI.md#setsettings"><strong>SetSettings</strong></a></td>
        <td><strong>Put</strong> /api/2.0/apps/{id}/settings</td>
        <td>Save app settings</td>
      </tr>
    </tbody>
  </table>

</details>
<details>
  <summary>Authentication</summary>

  <table>
    <tbody>
      <tr>
        <th>Method</th>
        <th>HTTP request</th>
        <th>Description</th>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>AuthenticationAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/AuthenticationAPI.md#authenticateme"><strong>AuthenticateMe</strong></a></td>
        <td><strong>Post</strong> /api/2.0/authentication</td>
        <td>Authenticate a user</td>
      </tr>
      <tr>
        <td><a href="docs/AuthenticationAPI.md#authenticatemefrombodywithcode"><strong>AuthenticateMeFromBodyWithCode</strong></a></td>
        <td><strong>Post</strong> /api/2.0/authentication/{code}</td>
        <td>Authenticate a user by code</td>
      </tr>
      <tr>
        <td><a href="docs/AuthenticationAPI.md#checkconfirm"><strong>CheckConfirm</strong></a></td>
        <td><strong>Post</strong> /api/2.0/authentication/confirm</td>
        <td>Open confirmation email URL</td>
      </tr>
      <tr>
        <td><a href="docs/AuthenticationAPI.md#getisauthentificated"><strong>GetIsAuthentificated</strong></a></td>
        <td><strong>Get</strong> /api/2.0/authentication</td>
        <td>Check authentication</td>
      </tr>
      <tr>
        <td><a href="docs/AuthenticationAPI.md#logout"><strong>Logout</strong></a></td>
        <td><strong>Post</strong> /api/2.0/authentication/logout</td>
        <td>Log out</td>
      </tr>
      <tr>
        <td><a href="docs/AuthenticationAPI.md#savemobilephone"><strong>SaveMobilePhone</strong></a></td>
        <td><strong>Post</strong> /api/2.0/authentication/setphone</td>
        <td>Set a mobile phone</td>
      </tr>
      <tr>
        <td><a href="docs/AuthenticationAPI.md#sendsmscode"><strong>SendSmsCode</strong></a></td>
        <td><strong>Post</strong> /api/2.0/authentication/sendsms</td>
        <td>Send SMS code</td>
      </tr>
    </tbody>
  </table>

</details>
<details>
  <summary>Backup</summary>

  <table>
    <tbody>
      <tr>
        <th>Method</th>
        <th>HTTP request</th>
        <th>Description</th>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>BackupAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/BackupAPI.md#cancelbackup"><strong>CancelBackup</strong></a></td>
        <td><strong>Post</strong> /api/2.0/backup/cancelbackup</td>
        <td>Cancel current backup</td>
      </tr>
      <tr>
        <td><a href="docs/BackupAPI.md#createbackupschedule"><strong>CreateBackupSchedule</strong></a></td>
        <td><strong>Post</strong> /api/2.0/backup/createbackupschedule</td>
        <td>Create the backup schedule</td>
      </tr>
      <tr>
        <td><a href="docs/BackupAPI.md#deletebackup"><strong>DeleteBackup</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/backup/deletebackup/{id}</td>
        <td>Delete the backup</td>
      </tr>
      <tr>
        <td><a href="docs/BackupAPI.md#deletebackuphistory"><strong>DeleteBackupHistory</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/backup/deletebackuphistory</td>
        <td>Delete the backup history</td>
      </tr>
      <tr>
        <td><a href="docs/BackupAPI.md#deletebackupschedule"><strong>DeleteBackupSchedule</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/backup/deletebackupschedule</td>
        <td>Delete the backup schedule</td>
      </tr>
      <tr>
        <td><a href="docs/BackupAPI.md#getbackuphistory"><strong>GetBackupHistory</strong></a></td>
        <td><strong>Get</strong> /api/2.0/backup/getbackuphistory</td>
        <td>Get the backup history</td>
      </tr>
      <tr>
        <td><a href="docs/BackupAPI.md#getbackupprogress"><strong>GetBackupProgress</strong></a></td>
        <td><strong>Get</strong> /api/2.0/backup/getbackupprogress</td>
        <td>Get the backup progress</td>
      </tr>
      <tr>
        <td><a href="docs/BackupAPI.md#getbackupschedule"><strong>GetBackupSchedule</strong></a></td>
        <td><strong>Get</strong> /api/2.0/backup/getbackupschedule</td>
        <td>Get the backup schedule</td>
      </tr>
      <tr>
        <td><a href="docs/BackupAPI.md#getbackupscount"><strong>GetBackupsCount</strong></a></td>
        <td><strong>Get</strong> /api/2.0/backup/getbackupscount</td>
        <td>Get the number of backups</td>
      </tr>
      <tr>
        <td><a href="docs/BackupAPI.md#getbackupscounts"><strong>GetBackupsCounts</strong></a></td>
        <td><strong>Get</strong> /api/2.0/backup/getbackupscountbypaid</td>
        <td>Get the number of free and paid backups</td>
      </tr>
      <tr>
        <td><a href="docs/BackupAPI.md#getbackupsservicestate"><strong>GetBackupsServiceState</strong></a></td>
        <td><strong>Get</strong> /api/2.0/backup/getservicestate</td>
        <td>Get the backup service state</td>
      </tr>
      <tr>
        <td><a href="docs/BackupAPI.md#getrestoreprogress"><strong>GetRestoreProgress</strong></a></td>
        <td><strong>Get</strong> /api/2.0/backup/getrestoreprogress</td>
        <td>Get the restoring progress</td>
      </tr>
      <tr>
        <td><a href="docs/BackupAPI.md#startbackup"><strong>StartBackup</strong></a></td>
        <td><strong>Post</strong> /api/2.0/backup/startbackup</td>
        <td>Start the backup</td>
      </tr>
      <tr>
        <td><a href="docs/BackupAPI.md#startbackuprestore"><strong>StartBackupRestore</strong></a></td>
        <td><strong>Post</strong> /api/2.0/backup/startrestore</td>
        <td>Start the restoring process</td>
      </tr>
    </tbody>
  </table>

</details>
<details>
  <summary>Capabilities</summary>

  <table>
    <tbody>
      <tr>
        <th>Method</th>
        <th>HTTP request</th>
        <th>Description</th>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>CapabilitiesAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/CapabilitiesAPI.md#getportalcapabilities"><strong>GetPortalCapabilities</strong></a></td>
        <td><strong>Get</strong> /api/2.0/capabilities</td>
        <td>Get portal capabilities</td>
      </tr>
    </tbody>
  </table>

</details>
<details>
  <summary>Files</summary>

  <table>
    <tbody>
      <tr>
        <th>Method</th>
        <th>HTTP request</th>
        <th>Description</th>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>FilesAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#addfiletorecent"><strong>AddFileToRecent</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/file/{fileId}/recent</td>
        <td>Add a file to the Recent section</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#addtemplates"><strong>AddTemplates</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/templates</td>
        <td>Add template files</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#changeversionhistory"><strong>ChangeVersionHistory</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/file/{fileId}/history</td>
        <td>Change version history</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#checkfillformdraft"><strong>CheckFillFormDraft</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/masterform/{fileId}/checkfillformdraft</td>
        <td>Check the form draft filling</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#copyfileas"><strong>CopyFileAs</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/file/{fileId}/copyas</td>
        <td>Copy a file</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#createeditsession"><strong>CreateEditSession</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/file/{fileId}/edit_session</td>
        <td>Create the editing session</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#createfile"><strong>CreateFile</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/{folderId}/file</td>
        <td>Create a file</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#createfileinmydocuments"><strong>CreateFileInMyDocuments</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/@my/file</td>
        <td>Create a file in the My documents section</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#createfileprimaryexternallink"><strong>CreateFilePrimaryExternalLink</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/file/{id}/link</td>
        <td>Create primary external link</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#createhtmlfile"><strong>CreateHtmlFile</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/{folderId}/html</td>
        <td>Create an HTML file</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#createhtmlfileinmydocuments"><strong>CreateHtmlFileInMyDocuments</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/@my/html</td>
        <td>Create an HTML file in the My documents section</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#createtextfile"><strong>CreateTextFile</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/{folderId}/text</td>
        <td>Create a text file</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#createtextfileinmydocuments"><strong>CreateTextFileInMyDocuments</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/@my/text</td>
        <td>Create a text file in the My documents section</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#createthumbnails"><strong>CreateThumbnails</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/thumbnails</td>
        <td>Create file thumbnails</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#deletefile"><strong>DeleteFile</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/files/file/{fileId}</td>
        <td>Delete a file</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#deleterecent"><strong>DeleteRecent</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/files/recent</td>
        <td>Delete recent files</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#deletetemplates"><strong>DeleteTemplates</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/files/templates</td>
        <td>Delete template files</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#generatexlsx"><strong>GenerateXlsx</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/file/{fileId}/xlsx</td>
        <td>Generate XLSX report</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#getallformroles"><strong>GetAllFormRoles</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/file/{fileId}/formroles</td>
        <td>Get form roles</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#geteditdiffurl"><strong>GetEditDiffUrl</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/file/{fileId}/edit/diff</td>
        <td>Get changes URL</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#getedithistory"><strong>GetEditHistory</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/file/{fileId}/edit/history</td>
        <td>Get version history</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#getencryptioninfo"><strong>GetEncryptionInfo</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/{fileId}/access</td>
        <td>Get file encryption information</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#getfilehistory"><strong>GetFileHistory</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/file/{fileId}/log</td>
        <td>Get file history</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#getfileinfo"><strong>GetFileInfo</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/file/{fileId}</td>
        <td>Get file information</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#getfilelinks"><strong>GetFileLinks</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/file/{id}/links</td>
        <td>Get file external links</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#getfileprimaryexternallink"><strong>GetFilePrimaryExternalLink</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/file/{id}/link</td>
        <td>Get primary external link</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#getfileversioninfo"><strong>GetFileVersionInfo</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/file/{fileId}/history</td>
        <td>Get file versions</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#getfillresult"><strong>GetFillResult</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/file/fillresult</td>
        <td>Get form-filling result</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#getformsubmissions"><strong>GetFormSubmissions</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/file/{fileId}/submissions</td>
        <td>Get form submission results</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#getpresignedfileuri"><strong>GetPresignedFileUri</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/file/{fileId}/presigned</td>
        <td>Get file download link asynchronously</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#getpresigneduri"><strong>GetPresignedUri</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/file/{fileId}/presigneduri</td>
        <td>Get file download link</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#getprotectedfileusers"><strong>GetProtectedFileUsers</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/file/{fileId}/protectusers</td>
        <td>Get users access rights to the protected file</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#getreferencedata"><strong>GetReferenceData</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/file/referencedata</td>
        <td>Get reference data</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#getxlsx"><strong>GetXlsx</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/file/{fileId}/xlsx</td>
        <td>Get XLSX report generation status</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#isformpdf"><strong>IsFormPDF</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/file/{fileId}/isformpdf</td>
        <td>Check the PDF file</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#lockfile"><strong>LockFile</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/file/{fileId}/lock</td>
        <td>Lock a file</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#manageformfilling"><strong>ManageFormFilling</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/file/{fileId}/manageformfilling</td>
        <td>Perform form filling action</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#openeditfile"><strong>OpenEditFile</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/file/{fileId}/openedit</td>
        <td>Open a file configuration</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#restorefileversion"><strong>RestoreFileVersion</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/file/{fileId}/restoreversion</td>
        <td>Restore a file version</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#saveeditingfilefromform"><strong>SaveEditingFileFromForm</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/file/{fileId}/saveediting</td>
        <td>Save file edits</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#savefileaspdf"><strong>SaveFileAsPdf</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/file/{id}/saveaspdf</td>
        <td>Save a file as PDF</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#saveformrolemapping"><strong>SaveFormRoleMapping</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/file/{fileId}/formrolemapping</td>
        <td>Save form role mapping</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#setcustomfiltertag"><strong>SetCustomFilterTag</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/file/{fileId}/customfilter</td>
        <td>Set the Custom Filter editing mode</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#setencryptioninfo"><strong>SetEncryptionInfo</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/{fileId}/access</td>
        <td>Set file encryption information</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#setfileexternallink"><strong>SetFileExternalLink</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/file/{id}/links</td>
        <td>Set an external link</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#setfileorder"><strong>SetFileOrder</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/{fileId}/order</td>
        <td>Set file order</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#setfilesorder"><strong>SetFilesOrder</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/order</td>
        <td>Set order of files</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#starteditfile"><strong>StartEditFile</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/file/{fileId}/startedit</td>
        <td>Start file editing</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#startfillingfile"><strong>StartFillingFile</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/file/{fileId}/startfilling</td>
        <td>Start file filling</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#togglefilefavorite"><strong>ToggleFileFavorite</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/favorites/{fileId}</td>
        <td>Change the file favorite status</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#trackeditfile"><strong>TrackEditFile</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/file/{fileId}/trackeditfile</td>
        <td>Track file editing</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFilesAPI.md#updatefile"><strong>UpdateFile</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/file/{fileId}</td>
        <td>Update a file</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>FoldersAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/FilesFoldersAPI.md#checkupload"><strong>CheckUpload</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/{folderId}/upload/check</td>
        <td>Check file uploads</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFoldersAPI.md#createfolder"><strong>CreateFolder</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/folder/{folderId}</td>
        <td>Create a folder</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFoldersAPI.md#createfolderprimaryexternallink"><strong>CreateFolderPrimaryExternalLink</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/folder/{id}/link</td>
        <td>Create primary external link</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFoldersAPI.md#createreportfolderhistory"><strong>CreateReportFolderHistory</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/folder/{folderId}/log/report</td>
        <td>Start the folder history report generation</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFoldersAPI.md#deletefolder"><strong>DeleteFolder</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/files/folder/{folderId}</td>
        <td>Delete a folder</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFoldersAPI.md#generatexlsxbyfolder"><strong>GenerateXlsxByFolder</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/folder/{folderId}/xlsx</td>
        <td>Generate XLSX report by folder</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFoldersAPI.md#getfavoritesfolder"><strong>GetFavoritesFolder</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/@favorites</td>
        <td>Get the Favorites section</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFoldersAPI.md#getfilesusedspace"><strong>GetFilesUsedSpace</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/filesusedspace</td>
        <td>Get used space of files</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFoldersAPI.md#getfolder"><strong>GetFolder</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/{folderId}/formfilter</td>
        <td>Get folder form filter</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFoldersAPI.md#getfolderbyfolderid"><strong>GetFolderByFolderId</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/{folderId}</td>
        <td>Get a folder by ID</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFoldersAPI.md#getfolderhistory"><strong>GetFolderHistory</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/folder/{folderId}/log</td>
        <td>Get folder history</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFoldersAPI.md#getfolderinfo"><strong>GetFolderInfo</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/folder/{folderId}</td>
        <td>Get folder information</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFoldersAPI.md#getfolderlinks"><strong>GetFolderLinks</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/folder/{id}/links</td>
        <td>Get the folder links</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFoldersAPI.md#getfolderpath"><strong>GetFolderPath</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/folder/{folderId}/path</td>
        <td>Get the folder path</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFoldersAPI.md#getfolderprimaryexternallink"><strong>GetFolderPrimaryExternalLink</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/folder/{id}/link</td>
        <td>Get primary external link</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFoldersAPI.md#getfolders"><strong>GetFolders</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/{folderId}/subfolders</td>
        <td>Get subfolders</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFoldersAPI.md#getformsfolder"><strong>GetFormsFolder</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/@forms</td>
        <td>Get the Forms section</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFoldersAPI.md#getmyfolder"><strong>GetMyFolder</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/@my</td>
        <td>Get the My documents section</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFoldersAPI.md#getnewfolderitems"><strong>GetNewFolderItems</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/{folderId}/news</td>
        <td>Get new folder items</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFoldersAPI.md#getrecentfolder"><strong>GetRecentFolder</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/recent</td>
        <td>Get the Recent section</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFoldersAPI.md#getreportfolderhistory"><strong>GetReportFolderHistory</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/folder/{folderId}/log/report</td>
        <td>Get the folder history report generation status</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFoldersAPI.md#getrootfolders"><strong>GetRootFolders</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/@root</td>
        <td>Get filtered sections</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFoldersAPI.md#gettrashfolder"><strong>GetTrashFolder</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/@trash</td>
        <td>Get the Trash section</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFoldersAPI.md#insertfile"><strong>InsertFile</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/{folderId}/insert</td>
        <td>Insert a file</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFoldersAPI.md#insertfiletomyfrombody"><strong>InsertFileToMyFromBody</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/@my/insert</td>
        <td>Insert a file to the My documents section</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFoldersAPI.md#renamefolder"><strong>RenameFolder</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/folder/{folderId}</td>
        <td>Rename a folder</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFoldersAPI.md#setfolderorder"><strong>SetFolderOrder</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/folder/{folderId}/order</td>
        <td>Set folder order</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFoldersAPI.md#setfolderprimaryexternallink"><strong>SetFolderPrimaryExternalLink</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/folder/{id}/links</td>
        <td>Set the folder external link</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFoldersAPI.md#terminatereportfolderhistory"><strong>TerminateReportFolderHistory</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/files/folder/{folderId}/log/report</td>
        <td>Terminate the folder history report generation</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFoldersAPI.md#uploadfile"><strong>UploadFile</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/{folderId}/upload</td>
        <td>Upload a file</td>
      </tr>
      <tr>
        <td><a href="docs/FilesFoldersAPI.md#uploadfiletomy"><strong>UploadFileToMy</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/@my/upload</td>
        <td>Upload a file to the My documents section</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>OperationsAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/FilesOperationsAPI.md#abortuploadsession"><strong>AbortUploadSession</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/files/{folderId}/session/{sessionId}</td>
        <td>Aborts an in-progress file upload session.</td>
      </tr>
      <tr>
        <td><a href="docs/FilesOperationsAPI.md#addfavorites"><strong>AddFavorites</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/favorites</td>
        <td>Add favorite files and folders</td>
      </tr>
      <tr>
        <td><a href="docs/FilesOperationsAPI.md#bulkdownload"><strong>BulkDownload</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/fileops/bulkdownload</td>
        <td>Bulk download</td>
      </tr>
      <tr>
        <td><a href="docs/FilesOperationsAPI.md#checkconversionstatus"><strong>CheckConversionStatus</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/file/{fileId}/checkconversion</td>
        <td>Get conversion status</td>
      </tr>
      <tr>
        <td><a href="docs/FilesOperationsAPI.md#checkmoveorcopybatchitems"><strong>CheckMoveOrCopyBatchItems</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/fileops/move</td>
        <td>Move or copy files to a folder</td>
      </tr>
      <tr>
        <td><a href="docs/FilesOperationsAPI.md#checkmoveorcopydestfolder"><strong>CheckMoveOrCopyDestFolder</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/fileops/checkdestfolder</td>
        <td>Check for moving or copying files to a folder</td>
      </tr>
      <tr>
        <td><a href="docs/FilesOperationsAPI.md#copybatchitems"><strong>CopyBatchItems</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/fileops/copy</td>
        <td>Copy to the folder</td>
      </tr>
      <tr>
        <td><a href="docs/FilesOperationsAPI.md#createuploadsession"><strong>CreateUploadSession</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/{folderId}/upload/create_session</td>
        <td>Chunked upload</td>
      </tr>
      <tr>
        <td><a href="docs/FilesOperationsAPI.md#createuploadsessioninfolder"><strong>CreateUploadSessionInFolder</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/{folderId}/session</td>
        <td>Creates a session for uploading a file to a specific folder in chunks.</td>
      </tr>
      <tr>
        <td><a href="docs/FilesOperationsAPI.md#deletebatchitems"><strong>DeleteBatchItems</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/fileops/delete</td>
        <td>Delete files and folders</td>
      </tr>
      <tr>
        <td><a href="docs/FilesOperationsAPI.md#deletefavoritesfrombody"><strong>DeleteFavoritesFromBody</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/files/favorites</td>
        <td>Delete favorite files and folders (using body parameters)</td>
      </tr>
      <tr>
        <td><a href="docs/FilesOperationsAPI.md#deletefileversions"><strong>DeleteFileVersions</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/fileops/deleteversion</td>
        <td>Delete file versions</td>
      </tr>
      <tr>
        <td><a href="docs/FilesOperationsAPI.md#duplicatebatchitems"><strong>DuplicateBatchItems</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/fileops/duplicate</td>
        <td>Duplicate files and folders</td>
      </tr>
      <tr>
        <td><a href="docs/FilesOperationsAPI.md#emptytrash"><strong>EmptyTrash</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/fileops/emptytrash</td>
        <td>Empty the Trash folder</td>
      </tr>
      <tr>
        <td><a href="docs/FilesOperationsAPI.md#finalizesession"><strong>FinalizeSession</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/{folderId}/session/{sessionId}/finalize</td>
        <td>Finalize an upload session</td>
      </tr>
      <tr>
        <td><a href="docs/FilesOperationsAPI.md#getoperationstatuses"><strong>GetOperationStatuses</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/fileops</td>
        <td>Get active file operations</td>
      </tr>
      <tr>
        <td><a href="docs/FilesOperationsAPI.md#getoperationstatusesbytype"><strong>GetOperationStatusesByType</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/fileops/{operationType}</td>
        <td>Get file operation statuses</td>
      </tr>
      <tr>
        <td><a href="docs/FilesOperationsAPI.md#markasread"><strong>MarkAsRead</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/fileops/markasread</td>
        <td>Mark as read</td>
      </tr>
      <tr>
        <td><a href="docs/FilesOperationsAPI.md#movebatchitems"><strong>MoveBatchItems</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/fileops/move</td>
        <td>Move or copy to a folder</td>
      </tr>
      <tr>
        <td><a href="docs/FilesOperationsAPI.md#startfileconversion"><strong>StartFileConversion</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/file/{fileId}/checkconversion</td>
        <td>Start file conversion</td>
      </tr>
      <tr>
        <td><a href="docs/FilesOperationsAPI.md#terminatetasks"><strong>TerminateTasks</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/fileops/terminate/{id}</td>
        <td>Finish active operations</td>
      </tr>
      <tr>
        <td><a href="docs/FilesOperationsAPI.md#updatefilecomment"><strong>UpdateFileComment</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/file/{fileId}/comment</td>
        <td>Update a comment</td>
      </tr>
      <tr>
        <td><a href="docs/FilesOperationsAPI.md#uploadasyncsession"><strong>UploadAsyncSession</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/{folderId}/session/{sessionId}/upload</td>
        <td>Handles the upload of a chunk for an existing upload session.</td>
      </tr>
      <tr>
        <td><a href="docs/FilesOperationsAPI.md#uploadsession"><strong>UploadSession</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/{folderId}/session/{sessionId}</td>
        <td>Resumes an ongoing file upload session for uploading additional chunks of data.</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>QuotaAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/FilesQuotaAPI.md#resetroomquota"><strong>ResetRoomQuota</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/rooms/resetquota</td>
        <td>Reset the room quota limit</td>
      </tr>
      <tr>
        <td><a href="docs/FilesQuotaAPI.md#updateroomsquota"><strong>UpdateRoomsQuota</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/rooms/roomquota</td>
        <td>Change the room quota limit</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>SettingsAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/FilesSettingsAPI.md#changeaccesstothirdparty"><strong>ChangeAccessToThirdparty</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/thirdparty</td>
        <td>Change the third-party settings access</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSettingsAPI.md#changeautomaticallycleanup"><strong>ChangeAutomaticallyCleanUp</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/settings/autocleanup</td>
        <td>Update the trash bin auto-clearing setting</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSettingsAPI.md#changedefaultaccessrights"><strong>ChangeDefaultAccessRights</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/settings/dafaultaccessrights</td>
        <td>Change the default access rights</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSettingsAPI.md#changedeleteconfirm"><strong>ChangeDeleteConfirm</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/changedeleteconfrim</td>
        <td>Confirm the file deletion</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSettingsAPI.md#changedownloadzip"><strong>ChangeDownloadZip</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/settings/downloadtargz</td>
        <td>Change the archive format (using body parameters)</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSettingsAPI.md#changeexternalsharingsettings"><strong>ChangeExternalSharingSettings</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/settings/externalsharingsettings</td>
        <td>Change the Access Control external sharing settings</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSettingsAPI.md#checkdocserviceurl"><strong>CheckDocServiceUrl</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/docservice</td>
        <td>Check the document service URL</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSettingsAPI.md#displayfileextension"><strong>DisplayFileExtension</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/displayfileextension</td>
        <td>Display a file extension</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSettingsAPI.md#displayrecent"><strong>DisplayRecent</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/displayrecent</td>
        <td>Display the Recent folder</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSettingsAPI.md#externalshare"><strong>ExternalShare</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/settings/external</td>
        <td>Change the external sharing ability</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSettingsAPI.md#externalsharesocialmedia"><strong>ExternalShareSocialMedia</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/settings/externalsocialmedia</td>
        <td>Change the external sharing ability on social networks</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSettingsAPI.md#forcesave"><strong>Forcesave</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/forcesave</td>
        <td>Change the forcesaving ability</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSettingsAPI.md#getautomaticallycleanup"><strong>GetAutomaticallyCleanUp</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/settings/autocleanup</td>
        <td>Get the trash bin auto-clearing setting</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSettingsAPI.md#getdefaulttemplates"><strong>GetDefaultTemplates</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/settings/defaulttemplate</td>
        <td>Get the default template setting</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSettingsAPI.md#getdocserviceurl"><strong>GetDocServiceUrl</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/docservice</td>
        <td>Get the document service URL</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSettingsAPI.md#getfilesmodule"><strong>GetFilesModule</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/info</td>
        <td>Get the Documents information</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSettingsAPI.md#getfilessettings"><strong>GetFilesSettings</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/settings</td>
        <td>Get file settings</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSettingsAPI.md#hideconfirmcanceloperation"><strong>HideConfirmCancelOperation</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/hideconfirmcanceloperation</td>
        <td>Hide confirmation dialog when canceling operations</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSettingsAPI.md#hideconfirmconvert"><strong>HideConfirmConvert</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/hideconfirmconvert</td>
        <td>Hide the confirmation dialog when converting</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSettingsAPI.md#hideconfirmroomlifetime"><strong>HideConfirmRoomLifetime</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/hideconfirmroomlifetime</td>
        <td>Hide confirmation dialog when changing room lifetime settings</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSettingsAPI.md#keepnewfilename"><strong>KeepNewFileName</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/keepnewfilename</td>
        <td>Ask a new file name</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSettingsAPI.md#resetdefaulttemplate"><strong>ResetDefaultTemplate</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/files/settings/defaulttemplate</td>
        <td>Reset the default template setting</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSettingsAPI.md#setdefaulttemplate"><strong>SetDefaultTemplate</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/settings/defaulttemplate</td>
        <td>Change the default template setting</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSettingsAPI.md#setopeneditorinsametab"><strong>SetOpenEditorInSameTab</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/settings/openeditorinsametab</td>
        <td>Open document in the same browser tab</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSettingsAPI.md#setorganizeroomsgrouping"><strong>SetOrganizeRoomsGrouping</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/settings/organizegrouping</td>
        <td>Organize rooms grouping</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSettingsAPI.md#storeforcesave"><strong>StoreForcesave</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/storeforcesave</td>
        <td>Change the ability to store the forcesaved files</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSettingsAPI.md#storeoriginal"><strong>StoreOriginal</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/storeoriginal</td>
        <td>Change the ability to upload original formats</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSettingsAPI.md#updatefileifexist"><strong>UpdateFileIfExist</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/updateifexist</td>
        <td>Update a file version if it exists</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSettingsAPI.md#uploaddefaulttemplate"><strong>UploadDefaultTemplate</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/settings/defaulttemplate</td>
        <td>Upload a file as the default template setting</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>SharingAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/FilesSharingAPI.md#applyexternalsharepassword"><strong>ApplyExternalSharePassword</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/share/{key}/password</td>
        <td>Apply external data password</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSharingAPI.md#changefileowner"><strong>ChangeFileOwner</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/owner</td>
        <td>Change the file owner</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSharingAPI.md#getencryptionaccess"><strong>GetEncryptionAccess</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/file/{fileId}/publickeys</td>
        <td>Get file encryption keys</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSharingAPI.md#getexternalsharedata"><strong>GetExternalShareData</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/share/{key}</td>
        <td>Get the external data</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSharingAPI.md#getfilesecurityinfo"><strong>GetFileSecurityInfo</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/file/{id}/share</td>
        <td>Get the shared file information</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSharingAPI.md#getfoldersecurityinfo"><strong>GetFolderSecurityInfo</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/folder/{id}/share</td>
        <td>Get the shared folder information</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSharingAPI.md#getgroupsmemberswithfilesecurity"><strong>GetGroupsMembersWithFileSecurity</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/file/{fileId}/group/{groupId}/share</td>
        <td>Get file group members with security information</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSharingAPI.md#getgroupsmemberswithfoldersecurity"><strong>GetGroupsMembersWithFolderSecurity</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/folder/{folderId}/group/{groupId}/share</td>
        <td>Get folder group members with security information</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSharingAPI.md#getsecurityinfo"><strong>GetSecurityInfo</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/share</td>
        <td>Get the sharing rights</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSharingAPI.md#getsharedusers"><strong>GetSharedUsers</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/file/{fileId}/sharedusers</td>
        <td>Get user access rights by file ID</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSharingAPI.md#removesecurityinfo"><strong>RemoveSecurityInfo</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/files/share</td>
        <td>Remove the sharing rights</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSharingAPI.md#sendeditornotify"><strong>SendEditorNotify</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/file/{fileId}/sendeditornotify</td>
        <td>Send the mention message</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSharingAPI.md#setfilesecurityinfo"><strong>SetFileSecurityInfo</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/file/{id}/share</td>
        <td>Share a file</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSharingAPI.md#setfoldersecurityinfo"><strong>SetFolderSecurityInfo</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/folder/{id}/share</td>
        <td>Share a folder</td>
      </tr>
      <tr>
        <td><a href="docs/FilesSharingAPI.md#setsecurityinfo"><strong>SetSecurityInfo</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/share</td>
        <td>Set the sharing rights</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>ThirdPartyIntegrationAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/FilesThirdPartyIntegrationAPI.md#deletethirdparty"><strong>DeleteThirdParty</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/files/thirdparty/{providerId}</td>
        <td>Remove a third-party account</td>
      </tr>
      <tr>
        <td><a href="docs/FilesThirdPartyIntegrationAPI.md#getallproviders"><strong>GetAllProviders</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/thirdparty/providers</td>
        <td>Get all providers</td>
      </tr>
      <tr>
        <td><a href="docs/FilesThirdPartyIntegrationAPI.md#getbackupthirdpartyaccount"><strong>GetBackupThirdPartyAccount</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/thirdparty/backup</td>
        <td>Get a third-party account backup</td>
      </tr>
      <tr>
        <td><a href="docs/FilesThirdPartyIntegrationAPI.md#getcapabilities"><strong>GetCapabilities</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/thirdparty/capabilities</td>
        <td>Get providers</td>
      </tr>
      <tr>
        <td><a href="docs/FilesThirdPartyIntegrationAPI.md#getcommonthirdpartyfolders"><strong>GetCommonThirdPartyFolders</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/thirdparty/common</td>
        <td>Get the common third-party services</td>
      </tr>
      <tr>
        <td><a href="docs/FilesThirdPartyIntegrationAPI.md#getthirdpartyaccounts"><strong>GetThirdPartyAccounts</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/thirdparty</td>
        <td>Get the third-party accounts</td>
      </tr>
      <tr>
        <td><a href="docs/FilesThirdPartyIntegrationAPI.md#savethirdparty"><strong>SaveThirdParty</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/thirdparty</td>
        <td>Save a third-party account</td>
      </tr>
      <tr>
        <td><a href="docs/FilesThirdPartyIntegrationAPI.md#savethirdpartybackup"><strong>SaveThirdPartyBackup</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/thirdparty/backup</td>
        <td>Save a third-party account backup</td>
      </tr>
    </tbody>
  </table>

</details>
<details>
  <summary>Group</summary>

  <table>
    <tbody>
      <tr>
        <th>Method</th>
        <th>HTTP request</th>
        <th>Description</th>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>GroupAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/GroupAPI.md#addgroup"><strong>AddGroup</strong></a></td>
        <td><strong>Post</strong> /api/2.0/group</td>
        <td>Add a new group</td>
      </tr>
      <tr>
        <td><a href="docs/GroupAPI.md#addmembersto"><strong>AddMembersTo</strong></a></td>
        <td><strong>Put</strong> /api/2.0/group/{id}/members</td>
        <td>Add group members</td>
      </tr>
      <tr>
        <td><a href="docs/GroupAPI.md#deletegroup"><strong>DeleteGroup</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/group/{id}</td>
        <td>Delete a group</td>
      </tr>
      <tr>
        <td><a href="docs/GroupAPI.md#getgroup"><strong>GetGroup</strong></a></td>
        <td><strong>Get</strong> /api/2.0/group/{id}</td>
        <td>Get a group</td>
      </tr>
      <tr>
        <td><a href="docs/GroupAPI.md#getgroupbyuserid"><strong>GetGroupByUserId</strong></a></td>
        <td><strong>Get</strong> /api/2.0/group/user/{userid}</td>
        <td>Get user groups</td>
      </tr>
      <tr>
        <td><a href="docs/GroupAPI.md#getgroups"><strong>GetGroups</strong></a></td>
        <td><strong>Get</strong> /api/2.0/group</td>
        <td>Get groups</td>
      </tr>
      <tr>
        <td><a href="docs/GroupAPI.md#movemembersto"><strong>MoveMembersTo</strong></a></td>
        <td><strong>Put</strong> /api/2.0/group/{fromId}/members/{toId}</td>
        <td>Move group members</td>
      </tr>
      <tr>
        <td><a href="docs/GroupAPI.md#removemembersfrom"><strong>RemoveMembersFrom</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/group/{id}/members</td>
        <td>Remove group members</td>
      </tr>
      <tr>
        <td><a href="docs/GroupAPI.md#setgroupmanager"><strong>SetGroupManager</strong></a></td>
        <td><strong>Put</strong> /api/2.0/group/{id}/manager</td>
        <td>Set a group manager</td>
      </tr>
      <tr>
        <td><a href="docs/GroupAPI.md#setmembersto"><strong>SetMembersTo</strong></a></td>
        <td><strong>Post</strong> /api/2.0/group/{id}/members</td>
        <td>Replace group members</td>
      </tr>
      <tr>
        <td><a href="docs/GroupAPI.md#updategroup"><strong>UpdateGroup</strong></a></td>
        <td><strong>Put</strong> /api/2.0/group/{id}</td>
        <td>Update a group</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>SearchAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/GroupSearchAPI.md#getgroupswithfilesshared"><strong>GetGroupsWithFilesShared</strong></a></td>
        <td><strong>Get</strong> /api/2.0/group/file/{id}</td>
        <td>Get groups with file sharing settings</td>
      </tr>
      <tr>
        <td><a href="docs/GroupSearchAPI.md#getgroupswithfoldersshared"><strong>GetGroupsWithFoldersShared</strong></a></td>
        <td><strong>Get</strong> /api/2.0/group/folder/{id}</td>
        <td>Get groups with folder sharing settings</td>
      </tr>
      <tr>
        <td><a href="docs/GroupSearchAPI.md#getgroupswithroomsshared"><strong>GetGroupsWithRoomsShared</strong></a></td>
        <td><strong>Get</strong> /api/2.0/group/room/{id}</td>
        <td>Get groups with room sharing settings</td>
      </tr>
    </tbody>
  </table>

</details>
<details>
  <summary>Migration</summary>

  <table>
    <tbody>
      <tr>
        <th>Method</th>
        <th>HTTP request</th>
        <th>Description</th>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>MigrationAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/MigrationAPI.md#cancelmigration"><strong>CancelMigration</strong></a></td>
        <td><strong>Post</strong> /api/2.0/migration/cancel</td>
        <td>Cancel migration</td>
      </tr>
      <tr>
        <td><a href="docs/MigrationAPI.md#clearmigration"><strong>ClearMigration</strong></a></td>
        <td><strong>Post</strong> /api/2.0/migration/clear</td>
        <td>Clear migration</td>
      </tr>
      <tr>
        <td><a href="docs/MigrationAPI.md#finishmigration"><strong>FinishMigration</strong></a></td>
        <td><strong>Post</strong> /api/2.0/migration/finish</td>
        <td>Finish migration</td>
      </tr>
      <tr>
        <td><a href="docs/MigrationAPI.md#getmigrationlogs"><strong>GetMigrationLogs</strong></a></td>
        <td><strong>Get</strong> /api/2.0/migration/logs</td>
        <td>Get migration logs</td>
      </tr>
      <tr>
        <td><a href="docs/MigrationAPI.md#getmigrationstatus"><strong>GetMigrationStatus</strong></a></td>
        <td><strong>Get</strong> /api/2.0/migration/status</td>
        <td>Get migration status</td>
      </tr>
      <tr>
        <td><a href="docs/MigrationAPI.md#listmigrations"><strong>ListMigrations</strong></a></td>
        <td><strong>Get</strong> /api/2.0/migration/list</td>
        <td>Get migrations</td>
      </tr>
      <tr>
        <td><a href="docs/MigrationAPI.md#startmigration"><strong>StartMigration</strong></a></td>
        <td><strong>Post</strong> /api/2.0/migration/migrate</td>
        <td>Start migration</td>
      </tr>
      <tr>
        <td><a href="docs/MigrationAPI.md#uploadandinitializemigration"><strong>UploadAndInitializeMigration</strong></a></td>
        <td><strong>Post</strong> /api/2.0/migration/init/{migratorName}</td>
        <td>Upload and initialize migration</td>
      </tr>
    </tbody>
  </table>

</details>
<details>
  <summary>OAuth20</summary>

  <table>
    <tbody>
      <tr>
        <th>Method</th>
        <th>HTTP request</th>
        <th>Description</th>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>AuthorizationAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/OAuth20AuthorizationAPI.md#authorizeoauth"><strong>AuthorizeOAuth</strong></a></td>
        <td><strong>Get</strong> /oauth2/authorize</td>
        <td>OAuth2 Authorization Endpoint</td>
      </tr>
      <tr>
        <td><a href="docs/OAuth20AuthorizationAPI.md#exchangetoken"><strong>ExchangeToken</strong></a></td>
        <td><strong>Post</strong> /oauth2/token</td>
        <td>OAuth2 Token Endpoint</td>
      </tr>
      <tr>
        <td><a href="docs/OAuth20AuthorizationAPI.md#submitconsent"><strong>SubmitConsent</strong></a></td>
        <td><strong>Post</strong> /oauth2/authorize</td>
        <td>OAuth2 consent endpoint</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>ClientManagementAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/OAuth20ClientManagementAPI.md#changeactivation"><strong>ChangeActivation</strong></a></td>
        <td><strong>Patch</strong> /api/2.0/clients/{clientId}/activation</td>
        <td>Change client activation status</td>
      </tr>
      <tr>
        <td><a href="docs/OAuth20ClientManagementAPI.md#createclient"><strong>CreateClient</strong></a></td>
        <td><strong>Post</strong> /api/2.0/clients</td>
        <td>Create a new OAuth2 client</td>
      </tr>
      <tr>
        <td><a href="docs/OAuth20ClientManagementAPI.md#deleteclient"><strong>DeleteClient</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/clients/{clientId}</td>
        <td>Delete an OAuth2 client</td>
      </tr>
      <tr>
        <td><a href="docs/OAuth20ClientManagementAPI.md#deletetenantclients"><strong>DeleteTenantClients</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/clients/tenant</td>
        <td>Delete all tenant OAuth2 clients</td>
      </tr>
      <tr>
        <td><a href="docs/OAuth20ClientManagementAPI.md#deleteuserclients"><strong>DeleteUserClients</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/clients</td>
        <td>Delete all user OAuth2 clients</td>
      </tr>
      <tr>
        <td><a href="docs/OAuth20ClientManagementAPI.md#regeneratesecret"><strong>RegenerateSecret</strong></a></td>
        <td><strong>Patch</strong> /api/2.0/clients/{clientId}/regenerate</td>
        <td>Regenerate client secret</td>
      </tr>
      <tr>
        <td><a href="docs/OAuth20ClientManagementAPI.md#revokeuserclient"><strong>RevokeUserClient</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/clients/{clientId}/revoke</td>
        <td>Revoke client consent</td>
      </tr>
      <tr>
        <td><a href="docs/OAuth20ClientManagementAPI.md#updateclient"><strong>UpdateClient</strong></a></td>
        <td><strong>Put</strong> /api/2.0/clients/{clientId}</td>
        <td>Update an existing OAuth2 client</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>ClientQueryingAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/OAuth20ClientQueryingAPI.md#getclient"><strong>GetClient</strong></a></td>
        <td><strong>Get</strong> /api/2.0/clients/{clientId}</td>
        <td>Get client details</td>
      </tr>
      <tr>
        <td><a href="docs/OAuth20ClientQueryingAPI.md#getclientinfo"><strong>GetClientInfo</strong></a></td>
        <td><strong>Get</strong> /api/2.0/clients/{clientId}/info</td>
        <td>Retrieves detailed information for a specific client</td>
      </tr>
      <tr>
        <td><a href="docs/OAuth20ClientQueryingAPI.md#getclients"><strong>GetClients</strong></a></td>
        <td><strong>Get</strong> /api/2.0/clients</td>
        <td>List clients</td>
      </tr>
      <tr>
        <td><a href="docs/OAuth20ClientQueryingAPI.md#getclientsinfo"><strong>GetClientsInfo</strong></a></td>
        <td><strong>Get</strong> /api/2.0/clients/info</td>
        <td>Retrieves a pageable list of client information</td>
      </tr>
      <tr>
        <td><a href="docs/OAuth20ClientQueryingAPI.md#getconsents"><strong>GetConsents</strong></a></td>
        <td><strong>Get</strong> /api/2.0/clients/consents</td>
        <td>Retrieves a pageable list of consents</td>
      </tr>
      <tr>
        <td><a href="docs/OAuth20ClientQueryingAPI.md#getpublicclientinfo"><strong>GetPublicClientInfo</strong></a></td>
        <td><strong>Get</strong> /api/2.0/clients/{clientId}/public/info</td>
        <td>Handles the GET request for public client information</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>DiscoveryAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/OAuth20DiscoveryAPI.md#handleoptions"><strong>HandleOptions</strong></a></td>
        <td><strong>Options</strong> /.well-known/oauth-authorization-server</td>
        <td></td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>ScopeManagementAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/OAuth20ScopeManagementAPI.md#getscopes"><strong>GetScopes</strong></a></td>
        <td><strong>Get</strong> /api/2.0/scopes</td>
        <td>List available OAuth2 scopes</td>
      </tr>
    </tbody>
  </table>

</details>
<details>
  <summary>People</summary>

  <table>
    <tbody>
      <tr>
        <th>Method</th>
        <th>HTTP request</th>
        <th>Description</th>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>EmailAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/PeopleEmailAPI.md#changeuseremail"><strong>ChangeUserEmail</strong></a></td>
        <td><strong>Put</strong> /api/2.0/people/{userid}/email</td>
        <td>Change a user email</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleEmailAPI.md#sendemailchangeinstructions"><strong>SendEmailChangeInstructions</strong></a></td>
        <td><strong>Post</strong> /api/2.0/people/email</td>
        <td>Send instructions to change email</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>GuestsAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/PeopleGuestsAPI.md#approveguestsharelink"><strong>ApproveGuestShareLink</strong></a></td>
        <td><strong>Post</strong> /api/2.0/people/guests/share/approve</td>
        <td>Approve a guest sharing link</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleGuestsAPI.md#deleteguests"><strong>DeleteGuests</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/people/guests</td>
        <td>Delete guests</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>PasswordAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/PeoplePasswordAPI.md#changeuserpassword"><strong>ChangeUserPassword</strong></a></td>
        <td><strong>Put</strong> /api/2.0/people/{userid}/password</td>
        <td>Change a user password</td>
      </tr>
      <tr>
        <td><a href="docs/PeoplePasswordAPI.md#senduserpassword"><strong>SendUserPassword</strong></a></td>
        <td><strong>Post</strong> /api/2.0/people/password</td>
        <td>Remind a user password</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>PhotosAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/PeoplePhotosAPI.md#creatememberphotothumbnails"><strong>CreateMemberPhotoThumbnails</strong></a></td>
        <td><strong>Post</strong> /api/2.0/people/{userid}/photo/thumbnails</td>
        <td>Create photo thumbnails</td>
      </tr>
      <tr>
        <td><a href="docs/PeoplePhotosAPI.md#deletememberphoto"><strong>DeleteMemberPhoto</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/people/{userid}/photo</td>
        <td>Delete a user photo</td>
      </tr>
      <tr>
        <td><a href="docs/PeoplePhotosAPI.md#getmemberphoto"><strong>GetMemberPhoto</strong></a></td>
        <td><strong>Get</strong> /api/2.0/people/{userid}/photo</td>
        <td>Get a user photo</td>
      </tr>
      <tr>
        <td><a href="docs/PeoplePhotosAPI.md#updatememberphoto"><strong>UpdateMemberPhoto</strong></a></td>
        <td><strong>Put</strong> /api/2.0/people/{userid}/photo</td>
        <td>Update a user photo</td>
      </tr>
      <tr>
        <td><a href="docs/PeoplePhotosAPI.md#uploadmemberphoto"><strong>UploadMemberPhoto</strong></a></td>
        <td><strong>Post</strong> /api/2.0/people/{userid}/photo</td>
        <td>Upload a user photo</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>PeopleProfilesAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/PeopleProfilesAPI.md#addmember"><strong>AddMember</strong></a></td>
        <td><strong>Post</strong> /api/2.0/people</td>
        <td>Add a user</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleProfilesAPI.md#checkuserexistsbyemail"><strong>CheckUserExistsByEmail</strong></a></td>
        <td><strong>Get</strong> /api/2.0/people/exists</td>
        <td>Check if a user exists by email</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleProfilesAPI.md#deletemember"><strong>DeleteMember</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/people/{userid}</td>
        <td>Delete a user</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleProfilesAPI.md#deleteprofile"><strong>DeleteProfile</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/people/@self</td>
        <td>Delete my profile</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleProfilesAPI.md#getallprofiles"><strong>GetAllProfiles</strong></a></td>
        <td><strong>Get</strong> /api/2.0/people</td>
        <td>Get profiles</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleProfilesAPI.md#getclaims"><strong>GetClaims</strong></a></td>
        <td><strong>Get</strong> /api/2.0/people/tokendiagnostics</td>
        <td>Get user claims</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleProfilesAPI.md#getprofilebyemail"><strong>GetProfileByEmail</strong></a></td>
        <td><strong>Get</strong> /api/2.0/people/email</td>
        <td>Get a profile by user email</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleProfilesAPI.md#getprofilebyuserid"><strong>GetProfileByUserId</strong></a></td>
        <td><strong>Get</strong> /api/2.0/people/{userid}</td>
        <td>Get a profile by user ID</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleProfilesAPI.md#getselfprofile"><strong>GetSelfProfile</strong></a></td>
        <td><strong>Get</strong> /api/2.0/people/@self</td>
        <td>Get my profile</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleProfilesAPI.md#inviteusers"><strong>InviteUsers</strong></a></td>
        <td><strong>Post</strong> /api/2.0/people/invite</td>
        <td>Invite users</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleProfilesAPI.md#removeusers"><strong>RemoveUsers</strong></a></td>
        <td><strong>Put</strong> /api/2.0/people/delete</td>
        <td>Delete users</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleProfilesAPI.md#resenduserinvites"><strong>ResendUserInvites</strong></a></td>
        <td><strong>Put</strong> /api/2.0/people/invite</td>
        <td>Resend activation emails</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleProfilesAPI.md#updatemember"><strong>UpdateMember</strong></a></td>
        <td><strong>Put</strong> /api/2.0/people/{userid}</td>
        <td>Update a user</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleProfilesAPI.md#updatememberculture"><strong>UpdateMemberCulture</strong></a></td>
        <td><strong>Put</strong> /api/2.0/people/{userid}/culture</td>
        <td>Update a user culture</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>PeopleQuotaAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/PeopleQuotaAPI.md#resetusersquota"><strong>ResetUsersQuota</strong></a></td>
        <td><strong>Put</strong> /api/2.0/people/resetquota</td>
        <td>Reset a user quota limit</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleQuotaAPI.md#updateuserquota"><strong>UpdateUserQuota</strong></a></td>
        <td><strong>Put</strong> /api/2.0/people/userquota</td>
        <td>Change a user quota limit</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>PeopleSearchAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/PeopleSearchAPI.md#getaccountsentrieswithfilesshared"><strong>GetAccountsEntriesWithFilesShared</strong></a></td>
        <td><strong>Get</strong> /api/2.0/accounts/file/{id}/search</td>
        <td>Get account entries with file sharing settings</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleSearchAPI.md#getaccountsentrieswithfoldersshared"><strong>GetAccountsEntriesWithFoldersShared</strong></a></td>
        <td><strong>Get</strong> /api/2.0/accounts/folder/{id}/search</td>
        <td>Get account entries with folder sharing settings</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleSearchAPI.md#getaccountsentrieswithroomsshared"><strong>GetAccountsEntriesWithRoomsShared</strong></a></td>
        <td><strong>Get</strong> /api/2.0/accounts/room/{id}/search</td>
        <td>Get account entries</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleSearchAPI.md#getsearch"><strong>GetSearch</strong></a></td>
        <td><strong>Get</strong> /api/2.0/people/@search/{query}</td>
        <td>Search users</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleSearchAPI.md#getsimplebyfilter"><strong>GetSimpleByFilter</strong></a></td>
        <td><strong>Get</strong> /api/2.0/people/simple/filter</td>
        <td>Search users by extended filter</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleSearchAPI.md#getuserswithfilesshared"><strong>GetUsersWithFilesShared</strong></a></td>
        <td><strong>Get</strong> /api/2.0/people/file/{id}</td>
        <td>Get users with file sharing settings</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleSearchAPI.md#getuserswithfoldersshared"><strong>GetUsersWithFoldersShared</strong></a></td>
        <td><strong>Get</strong> /api/2.0/people/folder/{id}</td>
        <td>Get users with folder sharing settings</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleSearchAPI.md#getuserswithroomshared"><strong>GetUsersWithRoomShared</strong></a></td>
        <td><strong>Get</strong> /api/2.0/people/room/{id}</td>
        <td>Get users with room sharing settings</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleSearchAPI.md#searchusersbyextendedfilter"><strong>SearchUsersByExtendedFilter</strong></a></td>
        <td><strong>Get</strong> /api/2.0/people/filter</td>
        <td>Search users with detailed information by extended filter</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleSearchAPI.md#searchusersbyquery"><strong>SearchUsersByQuery</strong></a></td>
        <td><strong>Get</strong> /api/2.0/people/search</td>
        <td>Search users (using query parameters)</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleSearchAPI.md#searchusersbystatus"><strong>SearchUsersByStatus</strong></a></td>
        <td><strong>Get</strong> /api/2.0/people/status/{status}/search</td>
        <td>Search users by status filter</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>ThemeAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/PeopleThemeAPI.md#changeportaltheme"><strong>ChangePortalTheme</strong></a></td>
        <td><strong>Put</strong> /api/2.0/people/theme</td>
        <td>Change the portal theme</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleThemeAPI.md#getportaltheme"><strong>GetPortalTheme</strong></a></td>
        <td><strong>Get</strong> /api/2.0/people/theme</td>
        <td>Get the portal theme</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>ThirdPartyAccountsAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/PeopleThirdPartyAccountsAPI.md#getthirdpartyauthproviders"><strong>GetThirdPartyAuthProviders</strong></a></td>
        <td><strong>Get</strong> /api/2.0/people/thirdparty/providers</td>
        <td>Get third-party accounts</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleThirdPartyAccountsAPI.md#linkthirdpartyaccount"><strong>LinkThirdPartyAccount</strong></a></td>
        <td><strong>Put</strong> /api/2.0/people/thirdparty/linkaccount</td>
        <td>Link a third-pary account</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleThirdPartyAccountsAPI.md#signupthirdpartyaccount"><strong>SignupThirdPartyAccount</strong></a></td>
        <td><strong>Post</strong> /api/2.0/people/thirdparty/signup</td>
        <td>Create a third-pary account</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleThirdPartyAccountsAPI.md#unlinkthirdpartyaccount"><strong>UnlinkThirdPartyAccount</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/people/thirdparty/unlinkaccount</td>
        <td>Unlink a third-pary account</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>UserDataAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/PeopleUserDataAPI.md#getdeletepersonalfolderprogress"><strong>GetDeletePersonalFolderProgress</strong></a></td>
        <td><strong>Get</strong> /api/2.0/people/delete/personal/progress</td>
        <td>Get the progress of deleting the personal folder</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleUserDataAPI.md#getreassignprogress"><strong>GetReassignProgress</strong></a></td>
        <td><strong>Get</strong> /api/2.0/people/reassign/progress/{userid}</td>
        <td>Get the reassignment progress</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleUserDataAPI.md#getremoveprogress"><strong>GetRemoveProgress</strong></a></td>
        <td><strong>Get</strong> /api/2.0/people/remove/progress/{userid}</td>
        <td>Get the deletion progress</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleUserDataAPI.md#necessaryreassign"><strong>NecessaryReassign</strong></a></td>
        <td><strong>Get</strong> /api/2.0/people/reassign/necessary</td>
        <td>Check data for reassignment need</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleUserDataAPI.md#sendinstructionstodelete"><strong>SendInstructionsToDelete</strong></a></td>
        <td><strong>Put</strong> /api/2.0/people/self/delete</td>
        <td>Send the deletion instructions</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleUserDataAPI.md#startdeletepersonalfolder"><strong>StartDeletePersonalFolder</strong></a></td>
        <td><strong>Post</strong> /api/2.0/people/delete/personal/start</td>
        <td>Delete the personal folder</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleUserDataAPI.md#startreassign"><strong>StartReassign</strong></a></td>
        <td><strong>Post</strong> /api/2.0/people/reassign/start</td>
        <td>Start the data reassignment</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleUserDataAPI.md#startremove"><strong>StartRemove</strong></a></td>
        <td><strong>Post</strong> /api/2.0/people/remove/start</td>
        <td>Start the data deletion</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleUserDataAPI.md#terminatereassign"><strong>TerminateReassign</strong></a></td>
        <td><strong>Put</strong> /api/2.0/people/reassign/terminate</td>
        <td>Terminate the data reassignment</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleUserDataAPI.md#terminateremove"><strong>TerminateRemove</strong></a></td>
        <td><strong>Put</strong> /api/2.0/people/remove/terminate</td>
        <td>Terminate the data deletion</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>UserStatusAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/PeopleUserStatusAPI.md#getbystatus"><strong>GetByStatus</strong></a></td>
        <td><strong>Get</strong> /api/2.0/people/status/{status}</td>
        <td>Get profiles by status</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleUserStatusAPI.md#updateuseractivationstatus"><strong>UpdateUserActivationStatus</strong></a></td>
        <td><strong>Put</strong> /api/2.0/people/activationstatus/{activationstatus}</td>
        <td>Set an activation status to the users</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleUserStatusAPI.md#updateuserstatus"><strong>UpdateUserStatus</strong></a></td>
        <td><strong>Put</strong> /api/2.0/people/status/{status}</td>
        <td>Change a user status</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>UserTypeAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/PeopleUserTypeAPI.md#getusertypeupdateprogress"><strong>GetUserTypeUpdateProgress</strong></a></td>
        <td><strong>Get</strong> /api/2.0/people/type/progress/{userid}</td>
        <td>Get the progress of updating user type</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleUserTypeAPI.md#startusertypeupdate"><strong>StartUserTypeUpdate</strong></a></td>
        <td><strong>Post</strong> /api/2.0/people/type</td>
        <td>Start updating user type</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleUserTypeAPI.md#terminateusertypeupdate"><strong>TerminateUserTypeUpdate</strong></a></td>
        <td><strong>Put</strong> /api/2.0/people/type/terminate</td>
        <td>Terminate updating user type</td>
      </tr>
      <tr>
        <td><a href="docs/PeopleUserTypeAPI.md#updateusertype"><strong>UpdateUserType</strong></a></td>
        <td><strong>Put</strong> /api/2.0/people/type/{type}</td>
        <td>Change a user type</td>
      </tr>
    </tbody>
  </table>

</details>
<details>
  <summary>Portal</summary>

  <table>
    <tbody>
      <tr>
        <th>Method</th>
        <th>HTTP request</th>
        <th>Description</th>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>PortalGuestsAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/PortalGuestsAPI.md#getguestsharinglink"><strong>GetGuestSharingLink</strong></a></td>
        <td><strong>Get</strong> /api/2.0/people/guests/{userid}/share</td>
        <td>Get a guest sharing link</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>PaymentAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#calculatewalletpayment"><strong>CalculateWalletPayment</strong></a></td>
        <td><strong>Put</strong> /api/2.0/portal/payment/calculatewallet</td>
        <td>Calculate the wallet payment amount</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#changetenantwalletservicestate"><strong>ChangeTenantWalletServiceState</strong></a></td>
        <td><strong>Post</strong> /api/2.0/portal/payment/servicestate</td>
        <td>Change tenant wallet service state</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#createcustomermonthlyusagereport"><strong>CreateCustomerMonthlyUsageReport</strong></a></td>
        <td><strong>Post</strong> /api/2.0/portal/payment/customer/usage/monthly/report</td>
        <td>Start the customer monthly usage report generation</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#createcustomeroperationsreport"><strong>CreateCustomerOperationsReport</strong></a></td>
        <td><strong>Post</strong> /api/2.0/portal/payment/customer/operationsreport</td>
        <td>Start the customer operations report generation</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#createcustomerserviceusagereport"><strong>CreateCustomerServiceUsageReport</strong></a></td>
        <td><strong>Post</strong> /api/2.0/portal/payment/customer/usage/report</td>
        <td>Start the customer service usage report generation</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#getactiveservices"><strong>GetActiveServices</strong></a></td>
        <td><strong>Get</strong> /api/2.0/portal/payment/activeservices</td>
        <td>Get the active wallet services</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#getaiprices"><strong>GetAiPrices</strong></a></td>
        <td><strong>Get</strong> /api/2.0/portal/payment/ai-prices</td>
        <td>Get AI model prices</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#getcheckoutsetupurl"><strong>GetCheckoutSetupUrl</strong></a></td>
        <td><strong>Get</strong> /api/2.0/portal/payment/checkoutsetupurl</td>
        <td>Get the checkout setup page URL</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#getcustomerbalance"><strong>GetCustomerBalance</strong></a></td>
        <td><strong>Get</strong> /api/2.0/portal/payment/customer/balance</td>
        <td>Get the customer balance</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#getcustomerinfo"><strong>GetCustomerInfo</strong></a></td>
        <td><strong>Get</strong> /api/2.0/portal/payment/customerinfo</td>
        <td>Get the customer information</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#getcustomermonthlyusage"><strong>GetCustomerMonthlyUsage</strong></a></td>
        <td><strong>Get</strong> /api/2.0/portal/payment/customer/usage/monthly</td>
        <td>Get the customer monthly usage</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#getcustomermonthlyusagereport"><strong>GetCustomerMonthlyUsageReport</strong></a></td>
        <td><strong>Get</strong> /api/2.0/portal/payment/customer/usage/monthly/report</td>
        <td>Get the status of the customer monthly usage report generation</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#getcustomeroperations"><strong>GetCustomerOperations</strong></a></td>
        <td><strong>Get</strong> /api/2.0/portal/payment/customer/operations</td>
        <td>Get the customer operations</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#getcustomeroperationsreport"><strong>GetCustomerOperationsReport</strong></a></td>
        <td><strong>Get</strong> /api/2.0/portal/payment/customer/operationsreport</td>
        <td>Get the status of the customer operations report generation</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#getcustomerserviceusage"><strong>GetCustomerServiceUsage</strong></a></td>
        <td><strong>Get</strong> /api/2.0/portal/payment/customer/usage</td>
        <td>Get the customer service usage</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#getcustomerserviceusagereport"><strong>GetCustomerServiceUsageReport</strong></a></td>
        <td><strong>Get</strong> /api/2.0/portal/payment/customer/usage/report</td>
        <td>Get the status of the customer service usage report generation</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#getpaymentaccount"><strong>GetPaymentAccount</strong></a></td>
        <td><strong>Get</strong> /api/2.0/portal/payment/account</td>
        <td>Get the payment account</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#getpaymentcurrencies"><strong>GetPaymentCurrencies</strong></a></td>
        <td><strong>Get</strong> /api/2.0/portal/payment/currencies</td>
        <td>Get currencies</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#getpaymentquotas"><strong>GetPaymentQuotas</strong></a></td>
        <td><strong>Get</strong> /api/2.0/portal/payment/quotas</td>
        <td>Get quotas</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#getpaymenturl"><strong>GetPaymentUrl</strong></a></td>
        <td><strong>Put</strong> /api/2.0/portal/payment/url</td>
        <td>Get the payment page URL</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#getportalprices"><strong>GetPortalPrices</strong></a></td>
        <td><strong>Get</strong> /api/2.0/portal/payment/prices</td>
        <td>Get prices</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#getquotapaymentinformation"><strong>GetQuotaPaymentInformation</strong></a></td>
        <td><strong>Get</strong> /api/2.0/portal/payment/quota</td>
        <td>Get quota payment information</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#getrestrictedaimodels"><strong>GetRestrictedAiModels</strong></a></td>
        <td><strong>Get</strong> /api/2.0/portal/payment/ai-model/restrictions</td>
        <td>Get restricted AI models</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#getsubscriptionbalanceinfo"><strong>GetSubscriptionBalanceInfo</strong></a></td>
        <td><strong>Get</strong> /api/2.0/portal/payment/subscription/balance</td>
        <td>Get the subscription balance information</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#gettenantwalletservicesettings"><strong>GetTenantWalletServiceSettings</strong></a></td>
        <td><strong>Get</strong> /api/2.0/portal/payment/servicessettings</td>
        <td>Gets the wallet service settings for the tenant.</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#gettenantwalletsettings"><strong>GetTenantWalletSettings</strong></a></td>
        <td><strong>Get</strong> /api/2.0/portal/payment/topupsettings</td>
        <td>Gets the tenant wallet auto top up settings</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#getwalletservice"><strong>GetWalletService</strong></a></td>
        <td><strong>Get</strong> /api/2.0/portal/payment/walletservice</td>
        <td>Get wallet service</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#getwalletservices"><strong>GetWalletServices</strong></a></td>
        <td><strong>Get</strong> /api/2.0/portal/payment/walletservices</td>
        <td>Get wallet services</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#movesubscriptiontowallet"><strong>MoveSubscriptionToWallet</strong></a></td>
        <td><strong>Post</strong> /api/2.0/portal/payment/subscription/movetowallet</td>
        <td>Move the subscription balance to the wallet and purchase admins</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#sendpaymentrequest"><strong>SendPaymentRequest</strong></a></td>
        <td><strong>Post</strong> /api/2.0/portal/payment/request</td>
        <td>Send a payment request</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#setrestrictedaimodels"><strong>SetRestrictedAiModels</strong></a></td>
        <td><strong>Put</strong> /api/2.0/portal/payment/ai-model/restrictions</td>
        <td>Set restricted AI models</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#settenantwalletsettings"><strong>SetTenantWalletSettings</strong></a></td>
        <td><strong>Post</strong> /api/2.0/portal/payment/topupsettings</td>
        <td>Set the wallet auto top up settings</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#terminatecustomermonthlyusagereport"><strong>TerminateCustomerMonthlyUsageReport</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/portal/payment/customer/usage/monthly/report</td>
        <td>Terminate the customer monthly usage report generation</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#terminatecustomeroperationsreport"><strong>TerminateCustomerOperationsReport</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/portal/payment/customer/operationsreport</td>
        <td>Terminate the customer operations report generation</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#terminatecustomerserviceusagereport"><strong>TerminateCustomerServiceUsageReport</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/portal/payment/customer/usage/report</td>
        <td>Terminate the customer service usage report generation</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#topupdeposit"><strong>TopUpDeposit</strong></a></td>
        <td><strong>Post</strong> /api/2.0/portal/payment/deposit</td>
        <td>Put money on deposit</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#updatepayment"><strong>UpdatePayment</strong></a></td>
        <td><strong>Put</strong> /api/2.0/portal/payment/update</td>
        <td>Update the payment quantity</td>
      </tr>
      <tr>
        <td><a href="docs/PortalPaymentAPI.md#updatewalletpayment"><strong>UpdateWalletPayment</strong></a></td>
        <td><strong>Put</strong> /api/2.0/portal/payment/updatewallet</td>
        <td>Update the wallet payment quantity</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>PortalQuotaAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/PortalQuotaAPI.md#getportalquota"><strong>GetPortalQuota</strong></a></td>
        <td><strong>Get</strong> /api/2.0/portal/quota</td>
        <td>Get a portal quota</td>
      </tr>
      <tr>
        <td><a href="docs/PortalQuotaAPI.md#getportaltariff"><strong>GetPortalTariff</strong></a></td>
        <td><strong>Get</strong> /api/2.0/portal/tariff</td>
        <td>Get a portal tariff</td>
      </tr>
      <tr>
        <td><a href="docs/PortalQuotaAPI.md#getportalusedspace"><strong>GetPortalUsedSpace</strong></a></td>
        <td><strong>Get</strong> /api/2.0/portal/usedspace</td>
        <td>Get the portal used space</td>
      </tr>
      <tr>
        <td><a href="docs/PortalQuotaAPI.md#getrightquota"><strong>GetRightQuota</strong></a></td>
        <td><strong>Get</strong> /api/2.0/portal/quota/right</td>
        <td>Get the recommended quota</td>
      </tr>
      <tr>
        <td><a href="docs/PortalQuotaAPI.md#getupcomingpayments"><strong>GetUpcomingPayments</strong></a></td>
        <td><strong>Get</strong> /api/2.0/portal/tariff/upcoming</td>
        <td>Get upcoming payments</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>PortalSettingsAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/PortalSettingsAPI.md#continueportal"><strong>ContinuePortal</strong></a></td>
        <td><strong>Put</strong> /api/2.0/portal/continue</td>
        <td>Restore a portal</td>
      </tr>
      <tr>
        <td><a href="docs/PortalSettingsAPI.md#deleteportal"><strong>DeletePortal</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/portal/delete</td>
        <td>Delete a portal</td>
      </tr>
      <tr>
        <td><a href="docs/PortalSettingsAPI.md#getportalinformation"><strong>GetPortalInformation</strong></a></td>
        <td><strong>Get</strong> /api/2.0/portal</td>
        <td>Get a portal</td>
      </tr>
      <tr>
        <td><a href="docs/PortalSettingsAPI.md#getportalpath"><strong>GetPortalPath</strong></a></td>
        <td><strong>Get</strong> /api/2.0/portal/path</td>
        <td>Get a path to the portal</td>
      </tr>
      <tr>
        <td><a href="docs/PortalSettingsAPI.md#senddeleteinstructions"><strong>SendDeleteInstructions</strong></a></td>
        <td><strong>Post</strong> /api/2.0/portal/delete</td>
        <td>Send removal instructions</td>
      </tr>
      <tr>
        <td><a href="docs/PortalSettingsAPI.md#sendsuspendinstructions"><strong>SendSuspendInstructions</strong></a></td>
        <td><strong>Post</strong> /api/2.0/portal/suspend</td>
        <td>Send suspension instructions</td>
      </tr>
      <tr>
        <td><a href="docs/PortalSettingsAPI.md#suspendportal"><strong>SuspendPortal</strong></a></td>
        <td><strong>Put</strong> /api/2.0/portal/suspend</td>
        <td>Deactivate a portal</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>UsersAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/PortalUsersAPI.md#createinvitationlink"><strong>CreateInvitationLink</strong></a></td>
        <td><strong>Post</strong> /api/2.0/portal/users/invitationlink</td>
        <td>Create an invitation link</td>
      </tr>
      <tr>
        <td><a href="docs/PortalUsersAPI.md#deleteinvitationlink"><strong>DeleteInvitationLink</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/portal/users/invitationlink</td>
        <td>Deletes an invitation link.</td>
      </tr>
      <tr>
        <td><a href="docs/PortalUsersAPI.md#getinvitationlink"><strong>GetInvitationLink</strong></a></td>
        <td><strong>Get</strong> /api/2.0/portal/users/invite/{employeeType}</td>
        <td>Get an invitation link</td>
      </tr>
      <tr>
        <td><a href="docs/PortalUsersAPI.md#getinvitationlinkbyemployeetype"><strong>GetInvitationLinkByEmployeeType</strong></a></td>
        <td><strong>Get</strong> /api/2.0/portal/users/invitationlink/{employeeType}</td>
        <td>Get an invitation link</td>
      </tr>
      <tr>
        <td><a href="docs/PortalUsersAPI.md#getportaluserscount"><strong>GetPortalUsersCount</strong></a></td>
        <td><strong>Get</strong> /api/2.0/portal/userscount</td>
        <td>Get a number of portal users</td>
      </tr>
      <tr>
        <td><a href="docs/PortalUsersAPI.md#getuserbyid"><strong>GetUserById</strong></a></td>
        <td><strong>Get</strong> /api/2.0/portal/users/{userID}</td>
        <td>Get a user by ID</td>
      </tr>
      <tr>
        <td><a href="docs/PortalUsersAPI.md#markgiftmessageasread"><strong>MarkGiftMessageAsRead</strong></a></td>
        <td><strong>Post</strong> /api/2.0/portal/present/mark</td>
        <td>Mark a gift message as read</td>
      </tr>
      <tr>
        <td><a href="docs/PortalUsersAPI.md#sendcongratulations"><strong>SendCongratulations</strong></a></td>
        <td><strong>Post</strong> /api/2.0/portal/sendcongratulations</td>
        <td>Send congratulations</td>
      </tr>
      <tr>
        <td><a href="docs/PortalUsersAPI.md#updateinvitationlink"><strong>UpdateInvitationLink</strong></a></td>
        <td><strong>Put</strong> /api/2.0/portal/users/invitationlink</td>
        <td>Update an invitation link</td>
      </tr>
    </tbody>
  </table>

</details>
<details>
  <summary>Rooms</summary>

  <table>
    <tbody>
      <tr>
        <th>Method</th>
        <th>HTTP request</th>
        <th>Description</th>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>RoomsAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#addroomtags"><strong>AddRoomTags</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/rooms/{id}/tags</td>
        <td>Add the room tags</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#archiveroom"><strong>ArchiveRoom</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/rooms/{id}/archive</td>
        <td>Archive a room</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#changeroomcover"><strong>ChangeRoomCover</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/rooms/{id}/cover</td>
        <td>Change the room cover</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#createroom"><strong>CreateRoom</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/rooms</td>
        <td>Create a room</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#createroomfromtemplate"><strong>CreateRoomFromTemplate</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/rooms/fromtemplate</td>
        <td>Create a room from the template</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#createroomlogo"><strong>CreateRoomLogo</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/rooms/{id}/logo</td>
        <td>Create a room logo</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#createroomtag"><strong>CreateRoomTag</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/tags</td>
        <td>Create a room tag</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#createroomtemplate"><strong>CreateRoomTemplate</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/roomtemplate</td>
        <td>Start creating room template</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#createroomthirdparty"><strong>CreateRoomThirdParty</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/rooms/thirdparty/{id}</td>
        <td>Create a third-party room</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#deletecustomtags"><strong>DeleteCustomTags</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/files/tags</td>
        <td>Delete the custom room tags</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#deleteroom"><strong>DeleteRoom</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/files/rooms/{id}</td>
        <td>Remove a room</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#deleteroomlogo"><strong>DeleteRoomLogo</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/files/rooms/{id}/logo</td>
        <td>Remove a room logo</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#deleteroomtags"><strong>DeleteRoomTags</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/files/rooms/{id}/tags</td>
        <td>Remove the room tags</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#getexternaldbsyncstatus"><strong>GetExternalDbSyncStatus</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/rooms/{id}/externaldbsync</td>
        <td>Get external DB sync status</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#getnewroomitems"><strong>GetNewRoomItems</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/rooms/{id}/news</td>
        <td>Get the new room items</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#getpublicsettings"><strong>GetPublicSettings</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/roomtemplate/{id}/public</td>
        <td>Get public settings</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#getroomcovers"><strong>GetRoomCovers</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/rooms/covers</td>
        <td>Get covers</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#getroomcreatingstatus"><strong>GetRoomCreatingStatus</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/rooms/fromtemplate/status</td>
        <td>Get the room creation progress</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#getroomindexexport"><strong>GetRoomIndexExport</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/rooms/indexexport</td>
        <td>Get the room index export</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#getroominfo"><strong>GetRoomInfo</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/rooms/{id}</td>
        <td>Get room information</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#getroomlinks"><strong>GetRoomLinks</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/rooms/{id}/links</td>
        <td>Get the room links</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#getroomsecurityinfo"><strong>GetRoomSecurityInfo</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/rooms/{id}/share</td>
        <td>Get the room access rights</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#getroomtagsinfo"><strong>GetRoomTagsInfo</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/tags</td>
        <td>Get the room tags</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#getroomtemplatecreatingstatus"><strong>GetRoomTemplateCreatingStatus</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/roomtemplate/status</td>
        <td>Get status of room template creation</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#getroomsfolder"><strong>GetRoomsFolder</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/rooms</td>
        <td>Get rooms</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#getroomsnewitems"><strong>GetRoomsNewItems</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/rooms/news</td>
        <td>Get the room new items</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#getroomsprimaryexternallink"><strong>GetRoomsPrimaryExternalLink</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/rooms/{id}/link</td>
        <td>Get the room primary external link</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#hastaglinks"><strong>HasTagLinks</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/tags/{tagName}/haslinks</td>
        <td>Has tag links</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#pinroom"><strong>PinRoom</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/rooms/{id}/pin</td>
        <td>Pin a room</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#reorderroom"><strong>ReorderRoom</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/rooms/{id}/reorder</td>
        <td>Reorder the room</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#resendemailinvitations"><strong>ResendEmailInvitations</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/rooms/{id}/resend</td>
        <td>Resend the room invitations</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#setpublicsettings"><strong>SetPublicSettings</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/roomtemplate/public</td>
        <td>Set public settings</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#setroomlink"><strong>SetRoomLink</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/rooms/{id}/links</td>
        <td>Set the room external or invitation link</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#setroomsecurity"><strong>SetRoomSecurity</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/rooms/{id}/share</td>
        <td>Set the room access rights</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#startexternaldbsync"><strong>StartExternalDbSync</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/rooms/{id}/externaldbsync</td>
        <td>Start external DB sync</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#startroomindexexport"><strong>StartRoomIndexExport</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/rooms/{id}/indexexport</td>
        <td>Start the room index export</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#terminateroomindexexport"><strong>TerminateRoomIndexExport</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/files/rooms/indexexport</td>
        <td>Terminate the room index export</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#unarchiveroom"><strong>UnarchiveRoom</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/rooms/{id}/unarchive</td>
        <td>Unarchive a room</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#unpinroom"><strong>UnpinRoom</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/rooms/{id}/unpin</td>
        <td>Unpin a room</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#updateroom"><strong>UpdateRoom</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/rooms/{id}</td>
        <td>Update a room</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#updateroomtag"><strong>UpdateRoomTag</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/tags</td>
        <td>Update tag</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsAPI.md#uploadroomlogo"><strong>UploadRoomLogo</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/logos</td>
        <td>Upload a room logo image</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>GroupsAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/RoomsGroupsAPI.md#addroomgroup"><strong>AddRoomGroup</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/group</td>
        <td>Add a new room group</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsGroupsAPI.md#changeroomgroupicon"><strong>ChangeRoomGroupIcon</strong></a></td>
        <td><strong>Post</strong> /api/2.0/files/group/{id}/icon</td>
        <td>Change group icon</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsGroupsAPI.md#deleteroomgroup"><strong>DeleteRoomGroup</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/files/group/{id}</td>
        <td>Delete group</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsGroupsAPI.md#getroomgroupinfo"><strong>GetRoomGroupInfo</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/group/{id}</td>
        <td>Get room group info</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsGroupsAPI.md#getroomgroups"><strong>GetRoomGroups</strong></a></td>
        <td><strong>Get</strong> /api/2.0/files/group</td>
        <td>List room groups</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsGroupsAPI.md#updateroomgroup"><strong>UpdateRoomGroup</strong></a></td>
        <td><strong>Put</strong> /api/2.0/files/group/{id}</td>
        <td>Update room group</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>PrivacyRoomAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/RoomsPrivacyRoomAPI.md#deletekeys"><strong>DeleteKeys</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/privacyroom/keys/{id}</td>
        <td>Deletes an encryption key and removes it from the system.</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsPrivacyRoomAPI.md#getuserkeys"><strong>GetUserKeys</strong></a></td>
        <td><strong>Get</strong> /api/2.0/privacyroom/keys</td>
        <td>Retrieves encryption keys associated with the current user.</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsPrivacyRoomAPI.md#getuserkeysforroom"><strong>GetUserKeysForRoom</strong></a></td>
        <td><strong>Get</strong> /api/2.0/privacyroom/{roomId}/access</td>
        <td>Retrieves the encryption keys associated with a specific privacy room.</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsPrivacyRoomAPI.md#replacekey"><strong>ReplaceKey</strong></a></td>
        <td><strong>Put</strong> /api/2.0/privacyroom/keys</td>
        <td>Replaces an existing encryption key with a new one for the user.</td>
      </tr>
      <tr>
        <td><a href="docs/RoomsPrivacyRoomAPI.md#setkeys"><strong>SetKeys</strong></a></td>
        <td><strong>Post</strong> /api/2.0/privacyroom/keys</td>
        <td>Creates and sets encryption keys for the user.</td>
      </tr>
    </tbody>
  </table>

</details>
<details>
  <summary>Security</summary>

  <table>
    <tbody>
      <tr>
        <th>Method</th>
        <th>HTTP request</th>
        <th>Description</th>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>SecurityAccessToDevToolsAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/SecurityAccessToDevToolsAPI.md#settenantdevtoolsaccesssettings"><strong>SetTenantDevToolsAccessSettings</strong></a></td>
        <td><strong>Post</strong> /api/2.0/settings/devtoolsaccess</td>
        <td>Set the Developer Tools access settings</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>ActiveConnectionsAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/SecurityActiveConnectionsAPI.md#getallactiveconnections"><strong>GetAllActiveConnections</strong></a></td>
        <td><strong>Get</strong> /api/2.0/security/activeconnections</td>
        <td>Get active connections</td>
      </tr>
      <tr>
        <td><a href="docs/SecurityActiveConnectionsAPI.md#logoutactiveconnection"><strong>LogOutActiveConnection</strong></a></td>
        <td><strong>Put</strong> /api/2.0/security/activeconnections/logout/{loginEventId}</td>
        <td>Log out from the connection</td>
      </tr>
      <tr>
        <td><a href="docs/SecurityActiveConnectionsAPI.md#logoutallactiveconnectionschangepassword"><strong>LogOutAllActiveConnectionsChangePassword</strong></a></td>
        <td><strong>Put</strong> /api/2.0/security/activeconnections/logoutallchangepassword</td>
        <td>Log out and change password</td>
      </tr>
      <tr>
        <td><a href="docs/SecurityActiveConnectionsAPI.md#logoutallactiveconnectionsforuser"><strong>LogOutAllActiveConnectionsForUser</strong></a></td>
        <td><strong>Put</strong> /api/2.0/security/activeconnections/logoutall/{userId}</td>
        <td>Log out for the user by ID</td>
      </tr>
      <tr>
        <td><a href="docs/SecurityActiveConnectionsAPI.md#logoutallexceptthisconnection"><strong>LogOutAllExceptThisConnection</strong></a></td>
        <td><strong>Put</strong> /api/2.0/security/activeconnections/logoutallexceptthis</td>
        <td>Log out from all connections except the current one</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>AuditTrailDataAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/SecurityAuditTrailDataAPI.md#createaudittrailreport"><strong>CreateAuditTrailReport</strong></a></td>
        <td><strong>Post</strong> /api/2.0/security/audit/events/report</td>
        <td>Start the audit trail report generation</td>
      </tr>
      <tr>
        <td><a href="docs/SecurityAuditTrailDataAPI.md#getauditeventsbyfilter"><strong>GetAuditEventsByFilter</strong></a></td>
        <td><strong>Get</strong> /api/2.0/security/audit/events/filter</td>
        <td>Get filtered audit trail data</td>
      </tr>
      <tr>
        <td><a href="docs/SecurityAuditTrailDataAPI.md#getauditsettings"><strong>GetAuditSettings</strong></a></td>
        <td><strong>Get</strong> /api/2.0/security/audit/settings/lifetime</td>
        <td>Get the audit trail settings</td>
      </tr>
      <tr>
        <td><a href="docs/SecurityAuditTrailDataAPI.md#getaudittrailmappers"><strong>GetAuditTrailMappers</strong></a></td>
        <td><strong>Get</strong> /api/2.0/security/audit/mappers</td>
        <td>Get audit trail mappers</td>
      </tr>
      <tr>
        <td><a href="docs/SecurityAuditTrailDataAPI.md#getaudittrailreport"><strong>GetAuditTrailReport</strong></a></td>
        <td><strong>Get</strong> /api/2.0/security/audit/events/report</td>
        <td>Get the audit trail report generation status</td>
      </tr>
      <tr>
        <td><a href="docs/SecurityAuditTrailDataAPI.md#getaudittrailtypes"><strong>GetAuditTrailTypes</strong></a></td>
        <td><strong>Get</strong> /api/2.0/security/audit/types</td>
        <td>Get audit trail types</td>
      </tr>
      <tr>
        <td><a href="docs/SecurityAuditTrailDataAPI.md#getlastauditevents"><strong>GetLastAuditEvents</strong></a></td>
        <td><strong>Get</strong> /api/2.0/security/audit/events/last</td>
        <td>Get audit trail data</td>
      </tr>
      <tr>
        <td><a href="docs/SecurityAuditTrailDataAPI.md#setauditsettings"><strong>SetAuditSettings</strong></a></td>
        <td><strong>Post</strong> /api/2.0/security/audit/settings/lifetime</td>
        <td>Set the audit trail settings</td>
      </tr>
      <tr>
        <td><a href="docs/SecurityAuditTrailDataAPI.md#terminateaudittrailreport"><strong>TerminateAuditTrailReport</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/security/audit/events/report</td>
        <td>Terminate the audit trail report generation</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>SecurityBannersVisibilityAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/SecurityBannersVisibilityAPI.md#settenantbannersettings"><strong>SetTenantBannerSettings</strong></a></td>
        <td><strong>Post</strong> /api/2.0/settings/banner</td>
        <td>Set the banners visibility</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>CSPAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/SecurityCSPAPI.md#configurecsp"><strong>ConfigureCsp</strong></a></td>
        <td><strong>Post</strong> /api/2.0/security/csp</td>
        <td>Configure CSP settings</td>
      </tr>
      <tr>
        <td><a href="docs/SecurityCSPAPI.md#getcspsettings"><strong>GetCspSettings</strong></a></td>
        <td><strong>Get</strong> /api/2.0/security/csp</td>
        <td>Get CSP settings</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>FirebaseAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/SecurityFirebaseAPI.md#docregisterpusnnotificationdevice"><strong>DocRegisterPusnNotificationDevice</strong></a></td>
        <td><strong>Post</strong> /api/2.0/settings/push/docregisterdevice</td>
        <td>Save the Documents Firebase device token</td>
      </tr>
      <tr>
        <td><a href="docs/SecurityFirebaseAPI.md#subscribedocumentspushnotification"><strong>SubscribeDocumentsPushNotification</strong></a></td>
        <td><strong>Put</strong> /api/2.0/settings/push/docsubscribe</td>
        <td>Subscribe to Documents push notification</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>LoginHistoryAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/SecurityLoginHistoryAPI.md#createloginhistoryreport"><strong>CreateLoginHistoryReport</strong></a></td>
        <td><strong>Post</strong> /api/2.0/security/audit/login/report</td>
        <td>Start the login history report generation</td>
      </tr>
      <tr>
        <td><a href="docs/SecurityLoginHistoryAPI.md#getlastloginevents"><strong>GetLastLoginEvents</strong></a></td>
        <td><strong>Get</strong> /api/2.0/security/audit/login/last</td>
        <td>Get login history</td>
      </tr>
      <tr>
        <td><a href="docs/SecurityLoginHistoryAPI.md#getlogineventsbyfilter"><strong>GetLoginEventsByFilter</strong></a></td>
        <td><strong>Get</strong> /api/2.0/security/audit/login/filter</td>
        <td>Get filtered login events</td>
      </tr>
      <tr>
        <td><a href="docs/SecurityLoginHistoryAPI.md#getloginhistoryreport"><strong>GetLoginHistoryReport</strong></a></td>
        <td><strong>Get</strong> /api/2.0/security/audit/login/report</td>
        <td>Get the login history report generation status</td>
      </tr>
      <tr>
        <td><a href="docs/SecurityLoginHistoryAPI.md#terminateloginhistoryreport"><strong>TerminateLoginHistoryReport</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/security/audit/login/report</td>
        <td>Terminate the login history report generation</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>OAuth2API</strong></td>
      </tr>
      <tr>
        <td><a href="docs/SecurityOAuth2API.md#generatejwttoken"><strong>GenerateJwtToken</strong></a></td>
        <td><strong>Get</strong> /api/2.0/security/oauth2/token</td>
        <td>Generate JWT token</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>SMTPSettingsAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/SecuritySMTPSettingsAPI.md#getsmtpoperationstatus"><strong>GetSmtpOperationStatus</strong></a></td>
        <td><strong>Get</strong> /api/2.0/smtpsettings/smtp/test/status</td>
        <td>Get the SMTP testing process status</td>
      </tr>
      <tr>
        <td><a href="docs/SecuritySMTPSettingsAPI.md#getsmtpsettings"><strong>GetSmtpSettings</strong></a></td>
        <td><strong>Get</strong> /api/2.0/smtpsettings/smtp</td>
        <td>Get the SMTP settings</td>
      </tr>
      <tr>
        <td><a href="docs/SecuritySMTPSettingsAPI.md#resetsmtpsettings"><strong>ResetSmtpSettings</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/smtpsettings/smtp</td>
        <td>Reset the SMTP settings</td>
      </tr>
      <tr>
        <td><a href="docs/SecuritySMTPSettingsAPI.md#savesmtpsettings"><strong>SaveSmtpSettings</strong></a></td>
        <td><strong>Post</strong> /api/2.0/smtpsettings/smtp</td>
        <td>Save the SMTP settings</td>
      </tr>
      <tr>
        <td><a href="docs/SecuritySMTPSettingsAPI.md#testsmtpsettings"><strong>TestSmtpSettings</strong></a></td>
        <td><strong>Get</strong> /api/2.0/smtpsettings/smtp/test</td>
        <td>Test the SMTP settings</td>
      </tr>
    </tbody>
  </table>

</details>
<details>
  <summary>Settings</summary>

  <table>
    <tbody>
      <tr>
        <th>Method</th>
        <th>HTTP request</th>
        <th>Description</th>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>AccessToDevToolsAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/SettingsAccessToDevToolsAPI.md#gettenantaccessdevtoolssettings"><strong>GetTenantAccessDevToolsSettings</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/devtoolsaccess</td>
        <td>Get the Developer Tools access settings</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>SettingsAuthorizationAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/SettingsAuthorizationAPI.md#getauthservices"><strong>GetAuthServices</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/authservice</td>
        <td>Get the authorization services</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsAuthorizationAPI.md#saveauthkeys"><strong>SaveAuthKeys</strong></a></td>
        <td><strong>Post</strong> /api/2.0/settings/authservice</td>
        <td>Save the authorization keys</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsAuthorizationAPI.md#testexternaldatabaseconnection"><strong>TestExternalDatabaseConnection</strong></a></td>
        <td><strong>Post</strong> /api/2.0/settings/authservice/externaldb/test</td>
        <td>Test external database connection</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>BannersVisibilityAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/SettingsBannersVisibilityAPI.md#gettenantbannersettings"><strong>GetTenantBannerSettings</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/banner</td>
        <td>Get the banners visibility</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>CommonSettingsAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/SettingsCommonSettingsAPI.md#closeadminhelper"><strong>CloseAdminHelper</strong></a></td>
        <td><strong>Put</strong> /api/2.0/settings/closeadminhelper</td>
        <td>Close the admin helper</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsCommonSettingsAPI.md#completewizard"><strong>CompleteWizard</strong></a></td>
        <td><strong>Put</strong> /api/2.0/settings/wizard/complete</td>
        <td>Complete the Wizard settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsCommonSettingsAPI.md#configuredeeplink"><strong>ConfigureDeepLink</strong></a></td>
        <td><strong>Post</strong> /api/2.0/settings/deeplink</td>
        <td>Configure the deep link settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsCommonSettingsAPI.md#deleteportalcolortheme"><strong>DeletePortalColorTheme</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/settings/colortheme</td>
        <td>Delete a color theme</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsCommonSettingsAPI.md#getdeeplinksettings"><strong>GetDeepLinkSettings</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/deeplink</td>
        <td>Get the deep link settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsCommonSettingsAPI.md#getpaymentsettings"><strong>GetPaymentSettings</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/payment</td>
        <td>Get the payment settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsCommonSettingsAPI.md#getportalcolortheme"><strong>GetPortalColorTheme</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/colortheme</td>
        <td>Get a color theme</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsCommonSettingsAPI.md#getportalhostname"><strong>GetPortalHostname</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/machine</td>
        <td>Get hostname</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsCommonSettingsAPI.md#getportallogo"><strong>GetPortalLogo</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/logo</td>
        <td>Get a portal logo</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsCommonSettingsAPI.md#getportalsettings"><strong>GetPortalSettings</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings</td>
        <td>Get the portal settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsCommonSettingsAPI.md#getsocketsettings"><strong>GetSocketSettings</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/socket</td>
        <td>Get the socket settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsCommonSettingsAPI.md#getsupportedcultures"><strong>GetSupportedCultures</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/cultures</td>
        <td>Get supported languages</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsCommonSettingsAPI.md#gettenantaiaccesssettings"><strong>GetTenantAiAccessSettings</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/ai-access</td>
        <td>Get the AI access settings for the portal</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsCommonSettingsAPI.md#gettenantuserinvitationsettings"><strong>GetTenantUserInvitationSettings</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/invitationsettings</td>
        <td>Get the user invitation settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsCommonSettingsAPI.md#gettimezones"><strong>GetTimeZones</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/timezones</td>
        <td>Get time zones</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsCommonSettingsAPI.md#savedefaultfolder"><strong>SaveDefaultFolder</strong></a></td>
        <td><strong>Put</strong> /api/2.0/settings/defaultfolder</td>
        <td>Set the default folder</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsCommonSettingsAPI.md#savednssettings"><strong>SaveDnsSettings</strong></a></td>
        <td><strong>Put</strong> /api/2.0/settings/dns</td>
        <td>Save the DNS settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsCommonSettingsAPI.md#savemaildomainsettings"><strong>SaveMailDomainSettings</strong></a></td>
        <td><strong>Post</strong> /api/2.0/settings/maildomainsettings</td>
        <td>Save the mail domain settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsCommonSettingsAPI.md#saveportalcolortheme"><strong>SavePortalColorTheme</strong></a></td>
        <td><strong>Put</strong> /api/2.0/settings/colortheme</td>
        <td>Save a color theme</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsCommonSettingsAPI.md#settenantaiaccesssettings"><strong>SetTenantAiAccessSettings</strong></a></td>
        <td><strong>Post</strong> /api/2.0/settings/ai-access</td>
        <td>Set the AI access for the portal</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsCommonSettingsAPI.md#updateemailactivationsettings"><strong>UpdateEmailActivationSettings</strong></a></td>
        <td><strong>Put</strong> /api/2.0/settings/emailactivation</td>
        <td>Update the email activation settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsCommonSettingsAPI.md#updateinvitationsettings"><strong>UpdateInvitationSettings</strong></a></td>
        <td><strong>Put</strong> /api/2.0/settings/invitationsettings</td>
        <td>Update user invitation settings</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>CookiesAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/SettingsCookiesAPI.md#getcookiesettings"><strong>GetCookieSettings</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/cookiesettings</td>
        <td>Get cookies lifetime</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsCookiesAPI.md#updatecookiesettings"><strong>UpdateCookieSettings</strong></a></td>
        <td><strong>Put</strong> /api/2.0/settings/cookiesettings</td>
        <td>Update cookies lifetime</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>DocsCloudAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/SettingsDocsCloudAPI.md#calculatedevpack"><strong>CalculateDevPack</strong></a></td>
        <td><strong>Post</strong> /api/2.0/settings/docscloud/calculatedevpack</td>
        <td>Calculate the DocsCloud subscription switch cost</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsDocsCloudAPI.md#createtenantquotareport"><strong>CreateTenantQuotaReport</strong></a></td>
        <td><strong>Post</strong> /api/2.0/settings/docscloud/tenant/quota/report</td>
        <td>Start the DocsCloud tenant quota report generation</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsDocsCloudAPI.md#gettenant"><strong>GetTenant</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/docscloud/tenant</td>
        <td>Get the DocsCloud tenant</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsDocsCloudAPI.md#gettenantconfig"><strong>GetTenantConfig</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/docscloud/tenant/config</td>
        <td>Get the DocsCloud tenant configuration</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsDocsCloudAPI.md#gettenantinfo"><strong>GetTenantInfo</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/docscloud/tenant/info</td>
        <td>Get the DocsCloud tenant information</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsDocsCloudAPI.md#gettenantquota"><strong>GetTenantQuota</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/docscloud/tenant/quota</td>
        <td>Get the DocsCloud tenant quota</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsDocsCloudAPI.md#gettenantquotareport"><strong>GetTenantQuotaReport</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/docscloud/tenant/quota/report</td>
        <td>Get the status of the DocsCloud tenant quota report generation</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsDocsCloudAPI.md#gettenantusage"><strong>GetTenantUsage</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/docscloud/tenant/usage</td>
        <td>Get the DocsCloud tenant usage</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsDocsCloudAPI.md#startdocscloudtrial"><strong>StartDocsCloudTrial</strong></a></td>
        <td><strong>Post</strong> /api/2.0/settings/docscloud/trial</td>
        <td>Start the DocsCloud trial</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsDocsCloudAPI.md#switchtodevpack"><strong>SwitchToDevPack</strong></a></td>
        <td><strong>Post</strong> /api/2.0/settings/docscloud/switchtodevpack</td>
        <td>Switch the DocsCloud subscription to DocsCloudDevPack</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsDocsCloudAPI.md#terminatetenantquotareport"><strong>TerminateTenantQuotaReport</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/settings/docscloud/tenant/quota/report</td>
        <td>Terminate the DocsCloud tenant quota report generation</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsDocsCloudAPI.md#updatetenantconfig"><strong>UpdateTenantConfig</strong></a></td>
        <td><strong>Put</strong> /api/2.0/settings/docscloud/tenant/config</td>
        <td>Update the DocsCloud tenant configuration</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>EncryptionAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/SettingsEncryptionAPI.md#getstorageencryptionprogress"><strong>GetStorageEncryptionProgress</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/encryption/progress</td>
        <td>Get the storage encryption progress</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsEncryptionAPI.md#getstorageencryptionsettings"><strong>GetStorageEncryptionSettings</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/encryption/settings</td>
        <td>Get the storage encryption settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsEncryptionAPI.md#startstorageencryption"><strong>StartStorageEncryption</strong></a></td>
        <td><strong>Post</strong> /api/2.0/settings/encryption/start</td>
        <td>Start the storage encryption process</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>GreetingSettingsAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/SettingsGreetingSettingsAPI.md#getgreetingsettings"><strong>GetGreetingSettings</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/greetingsettings</td>
        <td>Get greeting settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsGreetingSettingsAPI.md#getisdefaultgreetingsettings"><strong>GetIsDefaultGreetingSettings</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/greetingsettings/isdefault</td>
        <td>Check the default greeting settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsGreetingSettingsAPI.md#restoregreetingsettings"><strong>RestoreGreetingSettings</strong></a></td>
        <td><strong>Post</strong> /api/2.0/settings/greetingsettings/restore</td>
        <td>Restore the greeting settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsGreetingSettingsAPI.md#savegreetingsettings"><strong>SaveGreetingSettings</strong></a></td>
        <td><strong>Post</strong> /api/2.0/settings/greetingsettings</td>
        <td>Save the greeting settings</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>IPRestrictionsAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/SettingsIPRestrictionsAPI.md#getiprestrictions"><strong>GetIpRestrictions</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/iprestrictions</td>
        <td>Get the IP portal restrictions</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsIPRestrictionsAPI.md#readiprestrictionssettings"><strong>ReadIpRestrictionsSettings</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/iprestrictions/settings</td>
        <td>Get the IP restriction settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsIPRestrictionsAPI.md#saveiprestrictions"><strong>SaveIpRestrictions</strong></a></td>
        <td><strong>Put</strong> /api/2.0/settings/iprestrictions</td>
        <td>Update the IP restrictions</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsIPRestrictionsAPI.md#updateiprestrictionssettings"><strong>UpdateIpRestrictionsSettings</strong></a></td>
        <td><strong>Put</strong> /api/2.0/settings/iprestrictions/settings</td>
        <td>Update the IP restriction settings</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>LicenseAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/SettingsLicenseAPI.md#acceptlicense"><strong>AcceptLicense</strong></a></td>
        <td><strong>Post</strong> /api/2.0/settings/license/accept</td>
        <td>Activate a license</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsLicenseAPI.md#getislicenserequired"><strong>GetIsLicenseRequired</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/license/required</td>
        <td>Request a license</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsLicenseAPI.md#refreshlicense"><strong>RefreshLicense</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/license/refresh</td>
        <td>Refresh the license</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsLicenseAPI.md#uploadlicense"><strong>UploadLicense</strong></a></td>
        <td><strong>Post</strong> /api/2.0/settings/license</td>
        <td>Upload a license</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>LoginSettingsAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/SettingsLoginSettingsAPI.md#getloginsettings"><strong>GetLoginSettings</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/security/loginsettings</td>
        <td>Get the login settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsLoginSettingsAPI.md#setdefaultloginsettings"><strong>SetDefaultLoginSettings</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/settings/security/loginsettings</td>
        <td>Reset the login settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsLoginSettingsAPI.md#updateloginsettings"><strong>UpdateLoginSettings</strong></a></td>
        <td><strong>Put</strong> /api/2.0/settings/security/loginsettings</td>
        <td>Update the login settings</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>MessagesAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/SettingsMessagesAPI.md#enableadminmessagesettings"><strong>EnableAdminMessageSettings</strong></a></td>
        <td><strong>Post</strong> /api/2.0/settings/messagesettings</td>
        <td>Enable the administrator message settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsMessagesAPI.md#sendadminmail"><strong>SendAdminMail</strong></a></td>
        <td><strong>Post</strong> /api/2.0/settings/sendadmmail</td>
        <td>Send a message to the administrator</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsMessagesAPI.md#sendjoininvitemail"><strong>SendJoinInviteMail</strong></a></td>
        <td><strong>Post</strong> /api/2.0/settings/sendjoininvite</td>
        <td>Sends an invitation email</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>NotificationsAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/SettingsNotificationsAPI.md#getnotificationchannels"><strong>GetNotificationChannels</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/notification/channels</td>
        <td>Get notification channels</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsNotificationsAPI.md#getnotificationsettings"><strong>GetNotificationSettings</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/notification/{type}</td>
        <td>Check notification availability</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsNotificationsAPI.md#getroomsnotificationsettings"><strong>GetRoomsNotificationSettings</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/notification/rooms</td>
        <td>Get room notification settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsNotificationsAPI.md#setnotificationsettings"><strong>SetNotificationSettings</strong></a></td>
        <td><strong>Post</strong> /api/2.0/settings/notification</td>
        <td>Enable notifications</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsNotificationsAPI.md#setroomsnotificationstatus"><strong>SetRoomsNotificationStatus</strong></a></td>
        <td><strong>Post</strong> /api/2.0/settings/notification/rooms</td>
        <td>Set room notification status</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>OwnerAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/SettingsOwnerAPI.md#sendownerchangeinstructions"><strong>SendOwnerChangeInstructions</strong></a></td>
        <td><strong>Post</strong> /api/2.0/settings/owner</td>
        <td>Send the owner change instructions</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsOwnerAPI.md#updateportalowner"><strong>UpdatePortalOwner</strong></a></td>
        <td><strong>Put</strong> /api/2.0/settings/owner</td>
        <td>Update the portal owner</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>SettingsQuotaAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/SettingsQuotaAPI.md#getuserquotasettings"><strong>GetUserQuotaSettings</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/userquotasettings</td>
        <td>Get the user quota settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsQuotaAPI.md#saveaiagentquotasettings"><strong>SaveAiAgentQuotaSettings</strong></a></td>
        <td><strong>Post</strong> /api/2.0/settings/aiagentquotasettings</td>
        <td>Save the AI Agent quota settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsQuotaAPI.md#saveroomquotasettings"><strong>SaveRoomQuotaSettings</strong></a></td>
        <td><strong>Post</strong> /api/2.0/settings/roomquotasettings</td>
        <td>Save the room quota settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsQuotaAPI.md#settenantquotasettings"><strong>SetTenantQuotaSettings</strong></a></td>
        <td><strong>Put</strong> /api/2.0/settings/tenantquotasettings</td>
        <td>Save the tenant quota settings</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>RebrandingAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/SettingsRebrandingAPI.md#deleteadditionalwhitelabelsettings"><strong>DeleteAdditionalWhiteLabelSettings</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/settings/rebranding/additional</td>
        <td>Delete the additional white label settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsRebrandingAPI.md#deletecompanywhitelabelsettings"><strong>DeleteCompanyWhiteLabelSettings</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/settings/rebranding/company</td>
        <td>Delete the company white label settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsRebrandingAPI.md#getadditionalwhitelabelsettings"><strong>GetAdditionalWhiteLabelSettings</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/rebranding/additional</td>
        <td>Get the additional white label settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsRebrandingAPI.md#getcompanywhitelabelsettings"><strong>GetCompanyWhiteLabelSettings</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/rebranding/company</td>
        <td>Get the company white label settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsRebrandingAPI.md#getenablewhitelabel"><strong>GetEnableWhitelabel</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/enablewhitelabel</td>
        <td>Check the white label availability</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsRebrandingAPI.md#getisdefaultwhitelabellogotext"><strong>GetIsDefaultWhiteLabelLogoText</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/whitelabel/logotext/isdefault</td>
        <td>Check the default white label logo text</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsRebrandingAPI.md#getisdefaultwhitelabellogos"><strong>GetIsDefaultWhiteLabelLogos</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/whitelabel/logos/isdefault</td>
        <td>Check the default white label logos</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsRebrandingAPI.md#getlicensordata"><strong>GetLicensorData</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/companywhitelabel</td>
        <td>Get the licensor data</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsRebrandingAPI.md#getwhitelabellogotext"><strong>GetWhiteLabelLogoText</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/whitelabel/logotext</td>
        <td>Get the white label logo text</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsRebrandingAPI.md#getwhitelabellogos"><strong>GetWhiteLabelLogos</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/whitelabel/logos</td>
        <td>Get the white label logos</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsRebrandingAPI.md#restorewhitelabellogotext"><strong>RestoreWhiteLabelLogoText</strong></a></td>
        <td><strong>Put</strong> /api/2.0/settings/whitelabel/logotext/restore</td>
        <td>Restore the white label logo text</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsRebrandingAPI.md#restorewhitelabellogos"><strong>RestoreWhiteLabelLogos</strong></a></td>
        <td><strong>Put</strong> /api/2.0/settings/whitelabel/logos/restore</td>
        <td>Restore the white label logos</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsRebrandingAPI.md#saveadditionalwhitelabelsettings"><strong>SaveAdditionalWhiteLabelSettings</strong></a></td>
        <td><strong>Post</strong> /api/2.0/settings/rebranding/additional</td>
        <td>Save the additional white label settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsRebrandingAPI.md#savecompanywhitelabelsettings"><strong>SaveCompanyWhiteLabelSettings</strong></a></td>
        <td><strong>Post</strong> /api/2.0/settings/rebranding/company</td>
        <td>Save the company white label settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsRebrandingAPI.md#savewhitelabellogotext"><strong>SaveWhiteLabelLogoText</strong></a></td>
        <td><strong>Post</strong> /api/2.0/settings/whitelabel/logotext/save</td>
        <td>Save the white label logo text settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsRebrandingAPI.md#savewhitelabelsettings"><strong>SaveWhiteLabelSettings</strong></a></td>
        <td><strong>Post</strong> /api/2.0/settings/whitelabel/logos/save</td>
        <td>Save the white label logos</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsRebrandingAPI.md#savewhitelabelsettingsfromfiles"><strong>SaveWhiteLabelSettingsFromFiles</strong></a></td>
        <td><strong>Post</strong> /api/2.0/settings/whitelabel/logos/savefromfiles</td>
        <td>Save the white label logos from files</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>SSOAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/SettingsSSOAPI.md#getdefaultssosettingsv2"><strong>GetDefaultSsoSettingsV2</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/ssov2/default</td>
        <td>Get the default SSO settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsSSOAPI.md#getssosettingsv2"><strong>GetSsoSettingsV2</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/ssov2</td>
        <td>Get the SSO settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsSSOAPI.md#getssosettingsv2constants"><strong>GetSsoSettingsV2Constants</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/ssov2/constants</td>
        <td>Get the SSO settings constants</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsSSOAPI.md#resetssosettingsv2"><strong>ResetSsoSettingsV2</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/settings/ssov2</td>
        <td>Reset the SSO settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsSSOAPI.md#savessosettingsv2"><strong>SaveSsoSettingsV2</strong></a></td>
        <td><strong>Post</strong> /api/2.0/settings/ssov2</td>
        <td>Save the SSO settings</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>SecurityAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/SettingsSecurityAPI.md#getenabledmodules"><strong>GetEnabledModules</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/security/modules</td>
        <td>Get the enabled modules</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsSecurityAPI.md#getisproductadministrator"><strong>GetIsProductAdministrator</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/security/administrator</td>
        <td>Check a product administrator</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsSecurityAPI.md#getpasswordsettings"><strong>GetPasswordSettings</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/security/password</td>
        <td>Get the password settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsSecurityAPI.md#getproductadministrators"><strong>GetProductAdministrators</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/security/administrator/{productid}</td>
        <td>Get the product administrators</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsSecurityAPI.md#getwebitemsecurityinfo"><strong>GetWebItemSecurityInfo</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/security/{id}</td>
        <td>Get the module availability</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsSecurityAPI.md#getwebitemsettingssecurityinfo"><strong>GetWebItemSettingsSecurityInfo</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/security</td>
        <td>Get the security settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsSecurityAPI.md#setaccesstowebitems"><strong>SetAccessToWebItems</strong></a></td>
        <td><strong>Put</strong> /api/2.0/settings/security/access</td>
        <td>Set the security settings to modules</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsSecurityAPI.md#setproductadministrator"><strong>SetProductAdministrator</strong></a></td>
        <td><strong>Put</strong> /api/2.0/settings/security/administrator</td>
        <td>Set a product administrator</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsSecurityAPI.md#setwebitemsecurity"><strong>SetWebItemSecurity</strong></a></td>
        <td><strong>Put</strong> /api/2.0/settings/security</td>
        <td>Set the module security settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsSecurityAPI.md#updatepasswordsettings"><strong>UpdatePasswordSettings</strong></a></td>
        <td><strong>Put</strong> /api/2.0/settings/security/password</td>
        <td>Set the password settings</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>StatisticsAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/SettingsStatisticsAPI.md#getspaceusagestatistics"><strong>GetSpaceUsageStatistics</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/statistics/spaceusage/{id}</td>
        <td>Get the space usage statistics</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>StorageAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/SettingsStorageAPI.md#getallbackupstorages"><strong>GetAllBackupStorages</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/storage/backup</td>
        <td>Get the backup storages</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsStorageAPI.md#getallcdnstorages"><strong>GetAllCdnStorages</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/storage/cdn</td>
        <td>Get the CDN storages</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsStorageAPI.md#getallstorages"><strong>GetAllStorages</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/storage</td>
        <td>Get storages</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsStorageAPI.md#getamazons3regions"><strong>GetAmazonS3Regions</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/storage/s3/regions</td>
        <td>Get Amazon regions</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsStorageAPI.md#getstorageprogress"><strong>GetStorageProgress</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/storage/progress</td>
        <td>Get the storage progress</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsStorageAPI.md#resetcdntodefault"><strong>ResetCdnToDefault</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/settings/storage/cdn</td>
        <td>Reset the CDN storage settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsStorageAPI.md#resetstoragetodefault"><strong>ResetStorageToDefault</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/settings/storage</td>
        <td>Reset the storage settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsStorageAPI.md#updatecdnstorage"><strong>UpdateCdnStorage</strong></a></td>
        <td><strong>Put</strong> /api/2.0/settings/storage/cdn</td>
        <td>Update the CDN storage</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsStorageAPI.md#updatestorage"><strong>UpdateStorage</strong></a></td>
        <td><strong>Put</strong> /api/2.0/settings/storage</td>
        <td>Update a storage</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>TFASettingsAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/SettingsTFASettingsAPI.md#gettfaappcodes"><strong>GetTfaAppCodes</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/tfaappcodes</td>
        <td>Get the TFA codes</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsTFASettingsAPI.md#gettfaconfirmdata"><strong>GetTfaConfirmData</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/tfaapp/confirm</td>
        <td>Get TFA confirmation data</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsTFASettingsAPI.md#gettfasettings"><strong>GetTfaSettings</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/tfaapp</td>
        <td>Get the TFA settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsTFASettingsAPI.md#tfaappgeneratesetupcode"><strong>TfaAppGenerateSetupCode</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/tfaapp/setup</td>
        <td>Generate setup code</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsTFASettingsAPI.md#tfavalidateauthcode"><strong>TfaValidateAuthCode</strong></a></td>
        <td><strong>Post</strong> /api/2.0/settings/tfaapp/validate</td>
        <td>Validate the TFA code</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsTFASettingsAPI.md#unlinktfaapp"><strong>UnlinkTfaApp</strong></a></td>
        <td><strong>Put</strong> /api/2.0/settings/tfaappnewapp</td>
        <td>Unlink the TFA application</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsTFASettingsAPI.md#updatetfaappcodes"><strong>UpdateTfaAppCodes</strong></a></td>
        <td><strong>Put</strong> /api/2.0/settings/tfaappnewcodes</td>
        <td>Update the TFA codes</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsTFASettingsAPI.md#updatetfasettings"><strong>UpdateTfaSettings</strong></a></td>
        <td><strong>Put</strong> /api/2.0/settings/tfaapp</td>
        <td>Update the TFA settings</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsTFASettingsAPI.md#updatetfasettingslink"><strong>UpdateTfaSettingsLink</strong></a></td>
        <td><strong>Put</strong> /api/2.0/settings/tfaappwithlink</td>
        <td>Updates TFA settings</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>TelegramAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/SettingsTelegramAPI.md#checktelegram"><strong>CheckTelegram</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/telegram/check</td>
        <td>Check the Telegram connection</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsTelegramAPI.md#linktelegram"><strong>LinkTelegram</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/telegram/link</td>
        <td>Get the Telegram link</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsTelegramAPI.md#unlinktelegram"><strong>UnlinkTelegram</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/settings/telegram/link</td>
        <td>Unlink Telegram</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>WebhooksAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/SettingsWebhooksAPI.md#createwebhook"><strong>CreateWebhook</strong></a></td>
        <td><strong>Post</strong> /api/2.0/settings/webhook</td>
        <td>Create a webhook</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsWebhooksAPI.md#enablewebhook"><strong>EnableWebhook</strong></a></td>
        <td><strong>Put</strong> /api/2.0/settings/webhook/enable</td>
        <td>Enable a webhook</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsWebhooksAPI.md#gettenantwebhooks"><strong>GetTenantWebhooks</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/webhook</td>
        <td>Get webhooks</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsWebhooksAPI.md#getwebhooktriggers"><strong>GetWebhookTriggers</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/webhook/triggers</td>
        <td>Get webhook triggers</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsWebhooksAPI.md#getwebhookslogs"><strong>GetWebhooksLogs</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/webhooks/log</td>
        <td>Get webhook logs</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsWebhooksAPI.md#removewebhook"><strong>RemoveWebhook</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/settings/webhook/{id}</td>
        <td>Remove a webhook</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsWebhooksAPI.md#retrywebhook"><strong>RetryWebhook</strong></a></td>
        <td><strong>Put</strong> /api/2.0/settings/webhook/{id}/retry</td>
        <td>Retry a webhook</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsWebhooksAPI.md#retrywebhooks"><strong>RetryWebhooks</strong></a></td>
        <td><strong>Put</strong> /api/2.0/settings/webhook/retry</td>
        <td>Retry webhooks</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsWebhooksAPI.md#updatewebhook"><strong>UpdateWebhook</strong></a></td>
        <td><strong>Put</strong> /api/2.0/settings/webhook</td>
        <td>Update a webhook</td>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>WebpluginsAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/SettingsWebpluginsAPI.md#addwebpluginfromfile"><strong>AddWebPluginFromFile</strong></a></td>
        <td><strong>Post</strong> /api/2.0/settings/webplugins</td>
        <td>Add a web plugin</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsWebpluginsAPI.md#deletewebplugin"><strong>DeleteWebPlugin</strong></a></td>
        <td><strong>Delete</strong> /api/2.0/settings/webplugins/{name}</td>
        <td>Delete a web plugin</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsWebpluginsAPI.md#getwebplugin"><strong>GetWebPlugin</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/webplugins/{name}</td>
        <td>Get a web plugin by name</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsWebpluginsAPI.md#getwebplugins"><strong>GetWebPlugins</strong></a></td>
        <td><strong>Get</strong> /api/2.0/settings/webplugins</td>
        <td>Get web plugins</td>
      </tr>
      <tr>
        <td><a href="docs/SettingsWebpluginsAPI.md#updatewebplugin"><strong>UpdateWebPlugin</strong></a></td>
        <td><strong>Put</strong> /api/2.0/settings/webplugins/{name}</td>
        <td>Update a web plugin</td>
      </tr>
    </tbody>
  </table>

</details>
<details>
  <summary>ThirdParty</summary>

  <table>
    <tbody>
      <tr>
        <th>Method</th>
        <th>HTTP request</th>
        <th>Description</th>
      </tr>
      <tr>
        <td colspan="3" style="text-align: center;"><strong>ThirdPartyAPI</strong></td>
      </tr>
      <tr>
        <td><a href="docs/ThirdPartyAPI.md#getthirdpartycode"><strong>GetThirdPartyCode</strong></a></td>
        <td><strong>Get</strong> /api/2.0/thirdparty/{provider}</td>
        <td>Get the code request</td>
      </tr>
    </tbody>
  </table>

</details>

### Documentation For Models

<details><summary>Models list</summary>

 - [AccessRequestKeyDto](docs/AccessRequestKeyDto.md)
 - [AccountInfoArrayWrapper](docs/AccountInfoArrayWrapper.md)
 - [AccountInfoDto](docs/AccountInfoDto.md)
 - [AccountLoginType](docs/AccountLoginType.md)
 - [AceShortWrapper](docs/AceShortWrapper.md)
 - [AceShortWrapperArrayWrapper](docs/AceShortWrapperArrayWrapper.md)
 - [ActionConfig](docs/ActionConfig.md)
 - [ActionLinkConfig](docs/ActionLinkConfig.md)
 - [ActionType](docs/ActionType.md)
 - [ActiveConnectionsDto](docs/ActiveConnectionsDto.md)
 - [ActiveConnectionsItemDto](docs/ActiveConnectionsItemDto.md)
 - [ActiveConnectionsWrapper](docs/ActiveConnectionsWrapper.md)
 - [ActiveServiceArrayWrapper](docs/ActiveServiceArrayWrapper.md)
 - [ActiveServiceDto](docs/ActiveServiceDto.md)
 - [AdditionalWhiteLabelSettings](docs/AdditionalWhiteLabelSettings.md)
 - [AdditionalWhiteLabelSettingsDto](docs/AdditionalWhiteLabelSettingsDto.md)
 - [AdditionalWhiteLabelSettingsDtoWrapper](docs/AdditionalWhiteLabelSettingsDtoWrapper.md)
 - [AdditionalWhiteLabelSettingsResponseWrapper](docs/AdditionalWhiteLabelSettingsResponseWrapper.md)
 - [AdditionalWhiteLabelSettingsWrapper](docs/AdditionalWhiteLabelSettingsWrapper.md)
 - [AdminMessageBaseSettingsRequestsDto](docs/AdminMessageBaseSettingsRequestsDto.md)
 - [AdminMessageSettingsRequestsDto](docs/AdminMessageSettingsRequestsDto.md)
 - [AiActionType](docs/AiActionType.md)
 - [AiAgentNewItemsDto](docs/AiAgentNewItemsDto.md)
 - [AiAgentsCreateRequest](docs/AiAgentsCreateRequest.md)
 - [AiAgentsDeleteRequest](docs/AiAgentsDeleteRequest.md)
 - [AiAgentsResetQuotaRequest](docs/AiAgentsResetQuotaRequest.md)
 - [AiAgentsUpdateQuotaRequest](docs/AiAgentsUpdateQuotaRequest.md)
 - [AiAgentsUpdateQuotaRequestRoomIdsInner](docs/AiAgentsUpdateQuotaRequestRoomIdsInner.md)
 - [AiAgentsUpdateRequest](docs/AiAgentsUpdateRequest.md)
 - [AiAiActionArgs](docs/AiAiActionArgs.md)
 - [AiAiActionArgsPrompt](docs/AiAiActionArgsPrompt.md)
 - [AiAiApproveToolCallRequest](docs/AiAiApproveToolCallRequest.md)
 - [AiAiRegenerateStreamRequest](docs/AiAiRegenerateStreamRequest.md)
 - [AiAiSendCustomRequest](docs/AiAiSendCustomRequest.md)
 - [AiAiSendRequest](docs/AiAiSendRequest.md)
 - [AiAiSendStreamBody](docs/AiAiSendStreamBody.md)
 - [AiAiSettingsDto](docs/AiAiSettingsDto.md)
 - [AiAiSettingsWrapper](docs/AiAiSettingsWrapper.md)
 - [AiAiToolCallData](docs/AiAiToolCallData.md)
 - [AiAiUserSettingsDto](docs/AiAiUserSettingsDto.md)
 - [AiAiUserSettingsWrapper](docs/AiAiUserSettingsWrapper.md)
 - [AiAssignmentMutationResult](docs/AiAssignmentMutationResult.md)
 - [AiAssignmentsAssignRequest](docs/AiAssignmentsAssignRequest.md)
 - [AiAttachment](docs/AiAttachment.md)
 - [AiAttachmentFormKeysInner](docs/AiAttachmentFormKeysInner.md)
 - [AiAttachmentsLinkToMessageRequest](docs/AiAttachmentsLinkToMessageRequest.md)
 - [AiAttachmentsSaveFileRequest](docs/AiAttachmentsSaveFileRequest.md)
 - [AiAttachmentsSaveFileRequestInput](docs/AiAttachmentsSaveFileRequestInput.md)
 - [AiAttachmentsSaveFilesManyRequest](docs/AiAttachmentsSaveFilesManyRequest.md)
 - [AiBuiltinProviderType](docs/AiBuiltinProviderType.md)
 - [AiBulkAssignmentResult](docs/AiBulkAssignmentResult.md)
 - [AiBulkAssignmentResultErrorsInner](docs/AiBulkAssignmentResultErrorsInner.md)
 - [AiChatEvent](docs/AiChatEvent.md)
 - [AiChatModelPricing](docs/AiChatModelPricing.md)
 - [AiChatPrice](docs/AiChatPrice.md)
 - [AiChatSettingsDto](docs/AiChatSettingsDto.md)
 - [AiCreateProfileInput](docs/AiCreateProfileInput.md)
 - [AiCreatePromptInput](docs/AiCreatePromptInput.md)
 - [AiDistributedTaskStatus](docs/AiDistributedTaskStatus.md)
 - [AiEmbeddingModelPricing](docs/AiEmbeddingModelPricing.md)
 - [AiEmbeddingPrice](docs/AiEmbeddingPrice.md)
 - [AiEmbeddingProviderType](docs/AiEmbeddingProviderType.md)
 - [AiEmployeeDto](docs/AiEmployeeDto.md)
 - [AiErrorResponse](docs/AiErrorResponse.md)
 - [AiExportTextToDocx200Response](docs/AiExportTextToDocx200Response.md)
 - [AiExportTextToDocxRequest](docs/AiExportTextToDocxRequest.md)
 - [AiExportTextToDocxRequestFolderId](docs/AiExportTextToDocxRequestFolderId.md)
 - [AiFileEntryBaseDto](docs/AiFileEntryBaseDto.md)
 - [AiFileEntryDtoInteger](docs/AiFileEntryDtoInteger.md)
 - [AiFileEntryType](docs/AiFileEntryType.md)
 - [AiFileOperationDto](docs/AiFileOperationDto.md)
 - [AiFileOperationType](docs/AiFileOperationType.md)
 - [AiFileOperationWrapper](docs/AiFileOperationWrapper.md)
 - [AiFileShare](docs/AiFileShare.md)
 - [AiFolderContentDtoInteger](docs/AiFolderContentDtoInteger.md)
 - [AiFolderContentIntegerWrapper](docs/AiFolderContentIntegerWrapper.md)
 - [AiFolderDtoInteger](docs/AiFolderDtoInteger.md)
 - [AiFolderIntegerArrayWrapper](docs/AiFolderIntegerArrayWrapper.md)
 - [AiFolderIntegerWrapper](docs/AiFolderIntegerWrapper.md)
 - [AiFolderMutationResult](docs/AiFolderMutationResult.md)
 - [AiFolderType](docs/AiFolderType.md)
 - [AiImageModelPricing](docs/AiImageModelPricing.md)
 - [AiImagePrice](docs/AiImagePrice.md)
 - [AiImportError](docs/AiImportError.md)
 - [AiImportMode](docs/AiImportMode.md)
 - [AiImportResult](docs/AiImportResult.md)
 - [AiImportResultImported](docs/AiImportResultImported.md)
 - [AiLogo](docs/AiLogo.md)
 - [AiLogoCover](docs/AiLogoCover.md)
 - [AiModel](docs/AiModel.md)
 - [AiNewItemsAgentNewItemsArrayWrapper](docs/AiNewItemsAgentNewItemsArrayWrapper.md)
 - [AiNewItemsDtoAgentNewItemsDto](docs/AiNewItemsDtoAgentNewItemsDto.md)
 - [AiOpenAIChatCompletionChunk](docs/AiOpenAIChatCompletionChunk.md)
 - [AiOpenAIChoiceDelta](docs/AiOpenAIChoiceDelta.md)
 - [AiOpenAIChunkChoice](docs/AiOpenAIChunkChoice.md)
 - [AiOpenAIFinishReason](docs/AiOpenAIFinishReason.md)
 - [AiOpenAIStreamChunk](docs/AiOpenAIStreamChunk.md)
 - [AiOpenAIStreamError](docs/AiOpenAIStreamError.md)
 - [AiOpenAIStreamErrorError](docs/AiOpenAIStreamErrorError.md)
 - [AiOpenAIToolCallDelta](docs/AiOpenAIToolCallDelta.md)
 - [AiOpenAIToolCallDeltaFunction](docs/AiOpenAIToolCallDeltaFunction.md)
 - [AiOpenOrCreateResult](docs/AiOpenOrCreateResult.md)
 - [AiPreferencesSetDeepModeRequest](docs/AiPreferencesSetDeepModeRequest.md)
 - [AiPricesResponse](docs/AiPricesResponse.md)
 - [AiPricesResponseWrapper](docs/AiPricesResponseWrapper.md)
 - [AiProfile](docs/AiProfile.md)
 - [AiProfileMutationResult](docs/AiProfileMutationResult.md)
 - [AiProfilesGetById200Response](docs/AiProfilesGetById200Response.md)
 - [AiProfilesListProviderModelsRequest](docs/AiProfilesListProviderModelsRequest.md)
 - [AiProfilesTestConnection200Response](docs/AiProfilesTestConnection200Response.md)
 - [AiProfilesTestConnection200ResponseAnyOf](docs/AiProfilesTestConnection200ResponseAnyOf.md)
 - [AiPrompt](docs/AiPrompt.md)
 - [AiPromptBundle](docs/AiPromptBundle.md)
 - [AiPromptFolder](docs/AiPromptFolder.md)
 - [AiPromptMutationResult](docs/AiPromptMutationResult.md)
 - [AiPromptsImportBundleRequest](docs/AiPromptsImportBundleRequest.md)
 - [AiPromptsImportBundleRequestOptions](docs/AiPromptsImportBundleRequestOptions.md)
 - [AiPromptsMoveRequest](docs/AiPromptsMoveRequest.md)
 - [AiPromptsRenameFolderRequest](docs/AiPromptsRenameFolderRequest.md)
 - [AiPromptsUpdateRequest](docs/AiPromptsUpdateRequest.md)
 - [AiPromptsUpdateRequestUpdates](docs/AiPromptsUpdateRequestUpdates.md)
 - [AiProviderType](docs/AiProviderType.md)
 - [AiResolvedAssignment](docs/AiResolvedAssignment.md)
 - [AiRoomDataLifetimeDto](docs/AiRoomDataLifetimeDto.md)
 - [AiRoomDataLifetimePeriod](docs/AiRoomDataLifetimePeriod.md)
 - [AiRoomType](docs/AiRoomType.md)
 - [AiSuccessResponse](docs/AiSuccessResponse.md)
 - [AiTErrorData](docs/AiTErrorData.md)
 - [AiTMCPItem](docs/AiTMCPItem.md)
 - [AiTProvider](docs/AiTProvider.md)
 - [AiThread](docs/AiThread.md)
 - [AiThreadMessageLike](docs/AiThreadMessageLike.md)
 - [AiThreadMessageLikeContent](docs/AiThreadMessageLikeContent.md)
 - [AiThreadMessageLikeContentAnyOfInner](docs/AiThreadMessageLikeContentAnyOfInner.md)
 - [AiThreadMessageLikeStatus](docs/AiThreadMessageLikeStatus.md)
 - [AiThreadsAppendUserMessageRequest](docs/AiThreadsAppendUserMessageRequest.md)
 - [AiThreadsCreateRequest](docs/AiThreadsCreateRequest.md)
 - [AiThreadsOpenOrCreateRequest](docs/AiThreadsOpenOrCreateRequest.md)
 - [AiThreadsOpenOrCreateRequestEntityMeta](docs/AiThreadsOpenOrCreateRequestEntityMeta.md)
 - [AiThreadsRegenerateTitleRequest](docs/AiThreadsRegenerateTitleRequest.md)
 - [AiThreadsRenameRequest](docs/AiThreadsRenameRequest.md)
 - [AiThreadsTouchRequest](docs/AiThreadsTouchRequest.md)
 - [AiThreadsUpdateMessageRequest](docs/AiThreadsUpdateMessageRequest.md)
 - [AiToolsAddCustomServerRequest](docs/AiToolsAddCustomServerRequest.md)
 - [AiToolsBulkResult](docs/AiToolsBulkResult.md)
 - [AiToolsBulkResultErrorsInner](docs/AiToolsBulkResultErrorsInner.md)
 - [AiToolsMutationResult](docs/AiToolsMutationResult.md)
 - [AiToolsRemoveCustomServerRequest](docs/AiToolsRemoveCustomServerRequest.md)
 - [AiToolsReplaceAllCustomServersRequest](docs/AiToolsReplaceAllCustomServersRequest.md)
 - [AiToolsSetAllowAlwaysRequest](docs/AiToolsSetAllowAlwaysRequest.md)
 - [AiToolsSetDisabledRequest](docs/AiToolsSetDisabledRequest.md)
 - [AiToolsUpdateCustomServerRequest](docs/AiToolsUpdateCustomServerRequest.md)
 - [AiVectorizationSettingsDto](docs/AiVectorizationSettingsDto.md)
 - [AiVectorizationSettingsWrapper](docs/AiVectorizationSettingsWrapper.md)
 - [AiWatermarkAdditions](docs/AiWatermarkAdditions.md)
 - [AiWatermarkDto](docs/AiWatermarkDto.md)
 - [AiWebSearchConfig](docs/AiWebSearchConfig.md)
 - [AiWebSearchConfigureRequest](docs/AiWebSearchConfigureRequest.md)
 - [AiWebSearchMutationResult](docs/AiWebSearchMutationResult.md)
 - [AiWebSearchPricing](docs/AiWebSearchPricing.md)
 - [AnonymousConfigDto](docs/AnonymousConfigDto.md)
 - [ApiKeyResponseArrayWrapper](docs/ApiKeyResponseArrayWrapper.md)
 - [ApiKeyResponseDto](docs/ApiKeyResponseDto.md)
 - [ApiKeyResponseWrapper](docs/ApiKeyResponseWrapper.md)
 - [AppArrayWrapper](docs/AppArrayWrapper.md)
 - [AppDto](docs/AppDto.md)
 - [AppDtoSettings](docs/AppDtoSettings.md)
 - [AppWrapper](docs/AppWrapper.md)
 - [ApplyFilterOption](docs/ApplyFilterOption.md)
 - [ArchiveRoomRequest](docs/ArchiveRoomRequest.md)
 - [Area](docs/Area.md)
 - [ArrayArrayWrapper](docs/ArrayArrayWrapper.md)
 - [AuditEventArrayWrapper](docs/AuditEventArrayWrapper.md)
 - [AuditEventDto](docs/AuditEventDto.md)
 - [AuditReportFormat](docs/AuditReportFormat.md)
 - [AuthData](docs/AuthData.md)
 - [AuthKey](docs/AuthKey.md)
 - [AuthRequestsDto](docs/AuthRequestsDto.md)
 - [AuthServiceRequestsArrayWrapper](docs/AuthServiceRequestsArrayWrapper.md)
 - [AuthServiceRequestsDto](docs/AuthServiceRequestsDto.md)
 - [AuthWithCodeRequestsDto](docs/AuthWithCodeRequestsDto.md)
 - [AuthenticationTokenDto](docs/AuthenticationTokenDto.md)
 - [AuthenticationTokenWrapper](docs/AuthenticationTokenWrapper.md)
 - [AutoCleanUpData](docs/AutoCleanUpData.md)
 - [AutoCleanUpDataWrapper](docs/AutoCleanUpDataWrapper.md)
 - [AutoCleanupRequestDto](docs/AutoCleanupRequestDto.md)
 - [BackupDto](docs/BackupDto.md)
 - [BackupHistoryRecord](docs/BackupHistoryRecord.md)
 - [BackupHistoryRecordArrayWrapper](docs/BackupHistoryRecordArrayWrapper.md)
 - [BackupPeriod](docs/BackupPeriod.md)
 - [BackupProgress](docs/BackupProgress.md)
 - [BackupProgressEnum](docs/BackupProgressEnum.md)
 - [BackupProgressWrapper](docs/BackupProgressWrapper.md)
 - [BackupRestoreDto](docs/BackupRestoreDto.md)
 - [BackupScheduleDto](docs/BackupScheduleDto.md)
 - [BackupServiceStateDto](docs/BackupServiceStateDto.md)
 - [BackupServiceStateWrapper](docs/BackupServiceStateWrapper.md)
 - [BackupStorageType](docs/BackupStorageType.md)
 - [BackupsCountResultDto](docs/BackupsCountResultDto.md)
 - [BackupsCountResultWrapper](docs/BackupsCountResultWrapper.md)
 - [Balance](docs/Balance.md)
 - [BalanceWrapper](docs/BalanceWrapper.md)
 - [BaseBatchRequestDto](docs/BaseBatchRequestDto.md)
 - [BaseBatchRequestDtoAllOfFileIds](docs/BaseBatchRequestDtoAllOfFileIds.md)
 - [BaseBatchRequestDtoAllOfFolderIds](docs/BaseBatchRequestDtoAllOfFolderIds.md)
 - [BatchRequestDto](docs/BatchRequestDto.md)
 - [BatchRequestDtoAllOfDestFolderId](docs/BatchRequestDtoAllOfDestFolderId.md)
 - [BatchRequestDtoAllOfFileIds](docs/BatchRequestDtoAllOfFileIds.md)
 - [BatchRequestDtoAllOfFolderIds](docs/BatchRequestDtoAllOfFolderIds.md)
 - [BatchTagsRequestDto](docs/BatchTagsRequestDto.md)
 - [BooleanWrapper](docs/BooleanWrapper.md)
 - [CapabilitiesDto](docs/CapabilitiesDto.md)
 - [CapabilitiesWrapper](docs/CapabilitiesWrapper.md)
 - [CdnStorageSettings](docs/CdnStorageSettings.md)
 - [CdnStorageSettingsWrapper](docs/CdnStorageSettingsWrapper.md)
 - [ChangeClientActivationRequest](docs/ChangeClientActivationRequest.md)
 - [ChangeEmailRequest](docs/ChangeEmailRequest.md)
 - [ChangeHistory](docs/ChangeHistory.md)
 - [ChangeOwnerRequestDto](docs/ChangeOwnerRequestDto.md)
 - [ChangePasswordRequest](docs/ChangePasswordRequest.md)
 - [ChangeWalletServiceStateRequestDto](docs/ChangeWalletServiceStateRequestDto.md)
 - [ChatSettings](docs/ChatSettings.md)
 - [ChatSettingsDto](docs/ChatSettingsDto.md)
 - [CheckConversionRequestDtoInteger](docs/CheckConversionRequestDtoInteger.md)
 - [CheckDestFolderDto](docs/CheckDestFolderDto.md)
 - [CheckDestFolderResult](docs/CheckDestFolderResult.md)
 - [CheckDestFolderWrapper](docs/CheckDestFolderWrapper.md)
 - [CheckDocServiceUrlRequestDto](docs/CheckDocServiceUrlRequestDto.md)
 - [CheckFillFormDraft](docs/CheckFillFormDraft.md)
 - [CheckUploadRequest](docs/CheckUploadRequest.md)
 - [ChunkedUploadSessionResponseInteger](docs/ChunkedUploadSessionResponseInteger.md)
 - [ChunkedUploadSessionResponseIntegerWrapper](docs/ChunkedUploadSessionResponseIntegerWrapper.md)
 - [ChunkedUploadSessionResponseWrapperInteger](docs/ChunkedUploadSessionResponseWrapperInteger.md)
 - [ChunkedUploadSessionResponseWrapperIntegerWrapper](docs/ChunkedUploadSessionResponseWrapperIntegerWrapper.md)
 - [ClientInfoResponse](docs/ClientInfoResponse.md)
 - [ClientResponse](docs/ClientResponse.md)
 - [ClientSecretResponse](docs/ClientSecretResponse.md)
 - [CoEditingConfig](docs/CoEditingConfig.md)
 - [CoEditingConfigMode](docs/CoEditingConfigMode.md)
 - [CompanyWhiteLabelSettings](docs/CompanyWhiteLabelSettings.md)
 - [CompanyWhiteLabelSettingsArrayWrapper](docs/CompanyWhiteLabelSettingsArrayWrapper.md)
 - [CompanyWhiteLabelSettingsDto](docs/CompanyWhiteLabelSettingsDto.md)
 - [CompanyWhiteLabelSettingsDtoWrapper](docs/CompanyWhiteLabelSettingsDtoWrapper.md)
 - [CompanyWhiteLabelSettingsResponseWrapper](docs/CompanyWhiteLabelSettingsResponseWrapper.md)
 - [CompanyWhiteLabelSettingsWrapper](docs/CompanyWhiteLabelSettingsWrapper.md)
 - [ConfigurationDtoInteger](docs/ConfigurationDtoInteger.md)
 - [ConfigurationIntegerWrapper](docs/ConfigurationIntegerWrapper.md)
 - [ConfirmData](docs/ConfirmData.md)
 - [ConfirmDto](docs/ConfirmDto.md)
 - [ConfirmType](docs/ConfirmType.md)
 - [ConfirmWrapper](docs/ConfirmWrapper.md)
 - [ConnectionTestResult](docs/ConnectionTestResult.md)
 - [ConnectionTestResultWrapper](docs/ConnectionTestResultWrapper.md)
 - [Contact](docs/Contact.md)
 - [ConversationResultArrayWrapper](docs/ConversationResultArrayWrapper.md)
 - [ConversationResultDto](docs/ConversationResultDto.md)
 - [CookieSettingsDto](docs/CookieSettingsDto.md)
 - [CookieSettingsRequestsDto](docs/CookieSettingsRequestsDto.md)
 - [CookieSettingsWrapper](docs/CookieSettingsWrapper.md)
 - [CopyAsJsonElement](docs/CopyAsJsonElement.md)
 - [CopyAsJsonElementDestFolderId](docs/CopyAsJsonElementDestFolderId.md)
 - [CoverRequestDto](docs/CoverRequestDto.md)
 - [CoversResultArrayWrapper](docs/CoversResultArrayWrapper.md)
 - [CoversResultDto](docs/CoversResultDto.md)
 - [CreateApiKeyRequestDto](docs/CreateApiKeyRequestDto.md)
 - [CreateClientRequest](docs/CreateClientRequest.md)
 - [CreateFileJsonElement](docs/CreateFileJsonElement.md)
 - [CreateFileJsonElementTemplateId](docs/CreateFileJsonElementTemplateId.md)
 - [CreateFolder](docs/CreateFolder.md)
 - [CreateRoomFromTemplateDto](docs/CreateRoomFromTemplateDto.md)
 - [CreateRoomRequestDto](docs/CreateRoomRequestDto.md)
 - [CreateTagRequestDto](docs/CreateTagRequestDto.md)
 - [CreateTextOrHtmlFile](docs/CreateTextOrHtmlFile.md)
 - [CreateThirdPartyRoom](docs/CreateThirdPartyRoom.md)
 - [CreateWebhooksConfigRequestsDto](docs/CreateWebhooksConfigRequestsDto.md)
 - [Cron](docs/Cron.md)
 - [CronParams](docs/CronParams.md)
 - [CspDto](docs/CspDto.md)
 - [CspRequestsDto](docs/CspRequestsDto.md)
 - [CspWrapper](docs/CspWrapper.md)
 - [Culture](docs/Culture.md)
 - [CultureSpecificExternalResource](docs/CultureSpecificExternalResource.md)
 - [CultureSpecificExternalResources](docs/CultureSpecificExternalResources.md)
 - [CurrenciesArrayWrapper](docs/CurrenciesArrayWrapper.md)
 - [CurrenciesDto](docs/CurrenciesDto.md)
 - [CurrencyAmount](docs/CurrencyAmount.md)
 - [CurrencyCode](docs/CurrencyCode.md)
 - [CurrencyInfo](docs/CurrencyInfo.md)
 - [CurrentLicenseInfo](docs/CurrentLicenseInfo.md)
 - [CustomColorThemesSettingsColorItem](docs/CustomColorThemesSettingsColorItem.md)
 - [CustomColorThemesSettingsDto](docs/CustomColorThemesSettingsDto.md)
 - [CustomColorThemesSettingsItem](docs/CustomColorThemesSettingsItem.md)
 - [CustomColorThemesSettingsRequestsDto](docs/CustomColorThemesSettingsRequestsDto.md)
 - [CustomColorThemesSettingsWrapper](docs/CustomColorThemesSettingsWrapper.md)
 - [CustomFilterParameters](docs/CustomFilterParameters.md)
 - [CustomerConfigDto](docs/CustomerConfigDto.md)
 - [CustomerInfoDto](docs/CustomerInfoDto.md)
 - [CustomerInfoWrapper](docs/CustomerInfoWrapper.md)
 - [CustomerMonthlyUsageArrayWrapper](docs/CustomerMonthlyUsageArrayWrapper.md)
 - [CustomerMonthlyUsageDto](docs/CustomerMonthlyUsageDto.md)
 - [CustomerMonthlyUsageReportRequestDto](docs/CustomerMonthlyUsageReportRequestDto.md)
 - [CustomerOperationsReportRequestDto](docs/CustomerOperationsReportRequestDto.md)
 - [CustomerServiceUsageDto](docs/CustomerServiceUsageDto.md)
 - [CustomerServiceUsageReportDto](docs/CustomerServiceUsageReportDto.md)
 - [CustomerServiceUsageReportRequestDto](docs/CustomerServiceUsageReportRequestDto.md)
 - [CustomerServiceUsageReportWrapper](docs/CustomerServiceUsageReportWrapper.md)
 - [CustomizationConfigDto](docs/CustomizationConfigDto.md)
 - [DarkThemeSettings](docs/DarkThemeSettings.md)
 - [DarkThemeSettingsRequestDto](docs/DarkThemeSettingsRequestDto.md)
 - [DarkThemeSettingsType](docs/DarkThemeSettingsType.md)
 - [DarkThemeSettingsWrapper](docs/DarkThemeSettingsWrapper.md)
 - [DateToAutoCleanUp](docs/DateToAutoCleanUp.md)
 - [DbTenant](docs/DbTenant.md)
 - [DbTenantPartner](docs/DbTenantPartner.md)
 - [DeepLinkConfigurationRequestsDto](docs/DeepLinkConfigurationRequestsDto.md)
 - [DeepLinkDto](docs/DeepLinkDto.md)
 - [DeepLinkHandlingMode](docs/DeepLinkHandlingMode.md)
 - [DefaultProductRequestDto](docs/DefaultProductRequestDto.md)
 - [DefaultTemplateItemDto](docs/DefaultTemplateItemDto.md)
 - [DefaultTemplateSettingsDto](docs/DefaultTemplateSettingsDto.md)
 - [DefaultTemplateSettingsRequestDto](docs/DefaultTemplateSettingsRequestDto.md)
 - [DefaultTemplateSettingsRequestDtoSelectedFile](docs/DefaultTemplateSettingsRequestDtoSelectedFile.md)
 - [DefaultTemplateSettingsResetRequestDto](docs/DefaultTemplateSettingsResetRequestDto.md)
 - [DefaultTemplateSettingsWrapper](docs/DefaultTemplateSettingsWrapper.md)
 - [Delete](docs/Delete.md)
 - [DeleteBatchRequestDto](docs/DeleteBatchRequestDto.md)
 - [DeleteBatchRequestDtoAllOfFileIds](docs/DeleteBatchRequestDtoAllOfFileIds.md)
 - [DeleteBatchRequestDtoAllOfFolderIds](docs/DeleteBatchRequestDtoAllOfFolderIds.md)
 - [DeleteFolder](docs/DeleteFolder.md)
 - [DeleteRoomRequest](docs/DeleteRoomRequest.md)
 - [DeleteVersionBatchRequestDto](docs/DeleteVersionBatchRequestDto.md)
 - [DisplayRequestDto](docs/DisplayRequestDto.md)
 - [DistributedTaskStatus](docs/DistributedTaskStatus.md)
 - [DnsSettingsRequestsDto](docs/DnsSettingsRequestsDto.md)
 - [DocServiceUrlDto](docs/DocServiceUrlDto.md)
 - [DocServiceUrlWrapper](docs/DocServiceUrlWrapper.md)
 - [DocsCloudConfig](docs/DocsCloudConfig.md)
 - [DocsCloudConfigWrapper](docs/DocsCloudConfigWrapper.md)
 - [DocsCloudDevPackRequestDto](docs/DocsCloudDevPackRequestDto.md)
 - [DocsCloudIpFilterConfig](docs/DocsCloudIpFilterConfig.md)
 - [DocsCloudIpFilterRule](docs/DocsCloudIpFilterRule.md)
 - [DocsCloudLicenseInfo](docs/DocsCloudLicenseInfo.md)
 - [DocsCloudPayment](docs/DocsCloudPayment.md)
 - [DocsCloudQuota](docs/DocsCloudQuota.md)
 - [DocsCloudQuotaUser](docs/DocsCloudQuotaUser.md)
 - [DocsCloudQuotaWrapper](docs/DocsCloudQuotaWrapper.md)
 - [DocsCloudSecurityConfig](docs/DocsCloudSecurityConfig.md)
 - [DocsCloudServerConfig](docs/DocsCloudServerConfig.md)
 - [DocsCloudServerInfo](docs/DocsCloudServerInfo.md)
 - [DocsCloudStats](docs/DocsCloudStats.md)
 - [DocsCloudTenant](docs/DocsCloudTenant.md)
 - [DocsCloudTenantInfo](docs/DocsCloudTenantInfo.md)
 - [DocsCloudTenantInfoWrapper](docs/DocsCloudTenantInfoWrapper.md)
 - [DocsCloudTenantWrapper](docs/DocsCloudTenantWrapper.md)
 - [DocsCloudUsage](docs/DocsCloudUsage.md)
 - [DocsCloudUsageWrapper](docs/DocsCloudUsageWrapper.md)
 - [DocsCloudUserStats](docs/DocsCloudUserStats.md)
 - [DocsCloudUsersLimit](docs/DocsCloudUsersLimit.md)
 - [DocsCloudWopiConfig](docs/DocsCloudWopiConfig.md)
 - [DocumentBuilderTaskDto](docs/DocumentBuilderTaskDto.md)
 - [DocumentBuilderTaskWrapper](docs/DocumentBuilderTaskWrapper.md)
 - [DocumentConfigDto](docs/DocumentConfigDto.md)
 - [DoubleNullableWrapper](docs/DoubleNullableWrapper.md)
 - [DoubleWrapper](docs/DoubleWrapper.md)
 - [DownloadRequestDto](docs/DownloadRequestDto.md)
 - [DownloadRequestDtoAllOfFileIds](docs/DownloadRequestDtoAllOfFileIds.md)
 - [DownloadRequestDtoAllOfFolderIds](docs/DownloadRequestDtoAllOfFolderIds.md)
 - [DownloadRequestItemDto](docs/DownloadRequestItemDto.md)
 - [DownloadRequestItemDtoKey](docs/DownloadRequestItemDtoKey.md)
 - [DraftLocationInteger](docs/DraftLocationInteger.md)
 - [DuplicateRequestDto](docs/DuplicateRequestDto.md)
 - [DuplicateRequestDtoAllOfFileIds](docs/DuplicateRequestDtoAllOfFileIds.md)
 - [DuplicateRequestDtoAllOfFolderIds](docs/DuplicateRequestDtoAllOfFolderIds.md)
 - [EditHistoryArrayWrapper](docs/EditHistoryArrayWrapper.md)
 - [EditHistoryAuthor](docs/EditHistoryAuthor.md)
 - [EditHistoryChangesWrapper](docs/EditHistoryChangesWrapper.md)
 - [EditHistoryDataDto](docs/EditHistoryDataDto.md)
 - [EditHistoryDataWrapper](docs/EditHistoryDataWrapper.md)
 - [EditHistoryDto](docs/EditHistoryDto.md)
 - [EditHistoryUrl](docs/EditHistoryUrl.md)
 - [EditorConfigurationDto](docs/EditorConfigurationDto.md)
 - [EditorToolCallStateDto](docs/EditorToolCallStateDto.md)
 - [EditorType](docs/EditorType.md)
 - [EmailActivationSettings](docs/EmailActivationSettings.md)
 - [EmailActivationSettingsWrapper](docs/EmailActivationSettingsWrapper.md)
 - [EmailInvitationDto](docs/EmailInvitationDto.md)
 - [EmailMemberRequestDto](docs/EmailMemberRequestDto.md)
 - [EmailValidationKeyModel](docs/EmailValidationKeyModel.md)
 - [EmbeddedConfig](docs/EmbeddedConfig.md)
 - [EmployeeActivationStatus](docs/EmployeeActivationStatus.md)
 - [EmployeeArrayWrapper](docs/EmployeeArrayWrapper.md)
 - [EmployeeDto](docs/EmployeeDto.md)
 - [EmployeeFullArrayWrapper](docs/EmployeeFullArrayWrapper.md)
 - [EmployeeFullDto](docs/EmployeeFullDto.md)
 - [EmployeeFullWrapper](docs/EmployeeFullWrapper.md)
 - [EmployeeStatus](docs/EmployeeStatus.md)
 - [EmployeeType](docs/EmployeeType.md)
 - [EmployeeWrapper](docs/EmployeeWrapper.md)
 - [EncryprtionStatus](docs/EncryprtionStatus.md)
 - [EncryptionKeyArrayWrapper](docs/EncryptionKeyArrayWrapper.md)
 - [EncryptionKeyDto](docs/EncryptionKeyDto.md)
 - [EncryptionKeyRequestDto](docs/EncryptionKeyRequestDto.md)
 - [EncryptionSettings](docs/EncryptionSettings.md)
 - [EncryptionSettingsWrapper](docs/EncryptionSettingsWrapper.md)
 - [EntryType](docs/EntryType.md)
 - [ErrorApiResponse](docs/ErrorApiResponse.md)
 - [ErrorApiResponseError](docs/ErrorApiResponseError.md)
 - [ExchangeToken200Response](docs/ExchangeToken200Response.md)
 - [ExternalDatabaseSettings](docs/ExternalDatabaseSettings.md)
 - [ExternalDatabaseType](docs/ExternalDatabaseType.md)
 - [ExternalDbSyncFormResultDto](docs/ExternalDbSyncFormResultDto.md)
 - [ExternalDbSyncTaskDto](docs/ExternalDbSyncTaskDto.md)
 - [ExternalDbSyncTaskWrapper](docs/ExternalDbSyncTaskWrapper.md)
 - [ExternalShareDto](docs/ExternalShareDto.md)
 - [ExternalShareRequestParam](docs/ExternalShareRequestParam.md)
 - [ExternalShareWrapper](docs/ExternalShareWrapper.md)
 - [ExternalSharingSettingsDto](docs/ExternalSharingSettingsDto.md)
 - [ExternalSharingSettingsRequestDto](docs/ExternalSharingSettingsRequestDto.md)
 - [ExternalSharingSettingsWrapper](docs/ExternalSharingSettingsWrapper.md)
 - [FeatureUsedDto](docs/FeatureUsedDto.md)
 - [FeedbackConfig](docs/FeedbackConfig.md)
 - [FileConflictResolveType](docs/FileConflictResolveType.md)
 - [FileDtoInteger](docs/FileDtoInteger.md)
 - [FileDtoIntegerAllOfViewAccessibility](docs/FileDtoIntegerAllOfViewAccessibility.md)
 - [FileEncryptionInfoDto](docs/FileEncryptionInfoDto.md)
 - [FileEncryptionInfoWrapper](docs/FileEncryptionInfoWrapper.md)
 - [FileEntryBaseArrayWrapper](docs/FileEntryBaseArrayWrapper.md)
 - [FileEntryBaseDto](docs/FileEntryBaseDto.md)
 - [FileEntryBaseWrapper](docs/FileEntryBaseWrapper.md)
 - [FileEntryDtoInteger](docs/FileEntryDtoInteger.md)
 - [FileEntryDtoIntegerAllOfAvailableShareRights](docs/FileEntryDtoIntegerAllOfAvailableShareRights.md)
 - [FileEntryDtoIntegerAllOfSecurity](docs/FileEntryDtoIntegerAllOfSecurity.md)
 - [FileEntryDtoIntegerAllOfShareSettings](docs/FileEntryDtoIntegerAllOfShareSettings.md)
 - [FileEntryDtoString](docs/FileEntryDtoString.md)
 - [FileEntryIntegerArrayWrapper](docs/FileEntryIntegerArrayWrapper.md)
 - [FileEntryType](docs/FileEntryType.md)
 - [FileIntegerArrayWrapper](docs/FileIntegerArrayWrapper.md)
 - [FileIntegerWrapper](docs/FileIntegerWrapper.md)
 - [FileKeys](docs/FileKeys.md)
 - [FileLink](docs/FileLink.md)
 - [FileLinkRequest](docs/FileLinkRequest.md)
 - [FileLinkWrapper](docs/FileLinkWrapper.md)
 - [FileOperationArrayWrapper](docs/FileOperationArrayWrapper.md)
 - [FileOperationDto](docs/FileOperationDto.md)
 - [FileOperationRequestBaseDto](docs/FileOperationRequestBaseDto.md)
 - [FileOperationType](docs/FileOperationType.md)
 - [FileOperationWrapper](docs/FileOperationWrapper.md)
 - [FileReference](docs/FileReference.md)
 - [FileReferenceData](docs/FileReferenceData.md)
 - [FileReferenceWrapper](docs/FileReferenceWrapper.md)
 - [FileShare](docs/FileShare.md)
 - [FileShareArrayWrapper](docs/FileShareArrayWrapper.md)
 - [FileShareDto](docs/FileShareDto.md)
 - [FileShareLink](docs/FileShareLink.md)
 - [FileShareParams](docs/FileShareParams.md)
 - [FileShareResponseArrayWrapper](docs/FileShareResponseArrayWrapper.md)
 - [FileShareWrapper](docs/FileShareWrapper.md)
 - [FileStatus](docs/FileStatus.md)
 - [FileType](docs/FileType.md)
 - [FileUploadResultDto](docs/FileUploadResultDto.md)
 - [FileUploadResultWrapper](docs/FileUploadResultWrapper.md)
 - [FilesSettingsDto](docs/FilesSettingsDto.md)
 - [FilesSettingsDtoInternalFormats](docs/FilesSettingsDtoInternalFormats.md)
 - [FilesSettingsWrapper](docs/FilesSettingsWrapper.md)
 - [FilesStatisticsFolder](docs/FilesStatisticsFolder.md)
 - [FilesStatisticsResultDto](docs/FilesStatisticsResultDto.md)
 - [FilesStatisticsResultWrapper](docs/FilesStatisticsResultWrapper.md)
 - [FillingFormResultDtoInteger](docs/FillingFormResultDtoInteger.md)
 - [FillingFormResultIntegerWrapper](docs/FillingFormResultIntegerWrapper.md)
 - [FilterType](docs/FilterType.md)
 - [FinishDto](docs/FinishDto.md)
 - [FireBaseUser](docs/FireBaseUser.md)
 - [FireBaseUserWrapper](docs/FireBaseUserWrapper.md)
 - [FirebaseDto](docs/FirebaseDto.md)
 - [FirebaseRequestsDto](docs/FirebaseRequestsDto.md)
 - [FolderContentDtoInteger](docs/FolderContentDtoInteger.md)
 - [FolderContentIntegerArrayWrapper](docs/FolderContentIntegerArrayWrapper.md)
 - [FolderContentIntegerWrapper](docs/FolderContentIntegerWrapper.md)
 - [FolderDtoInteger](docs/FolderDtoInteger.md)
 - [FolderDtoString](docs/FolderDtoString.md)
 - [FolderIntegerArrayWrapper](docs/FolderIntegerArrayWrapper.md)
 - [FolderIntegerWrapper](docs/FolderIntegerWrapper.md)
 - [FolderLinkRequest](docs/FolderLinkRequest.md)
 - [FolderStringArrayWrapper](docs/FolderStringArrayWrapper.md)
 - [FolderStringWrapper](docs/FolderStringWrapper.md)
 - [FolderType](docs/FolderType.md)
 - [FormFillingManageAction](docs/FormFillingManageAction.md)
 - [FormFillingStatus](docs/FormFillingStatus.md)
 - [FormGalleryDto](docs/FormGalleryDto.md)
 - [FormMetadata](docs/FormMetadata.md)
 - [FormResultsDto](docs/FormResultsDto.md)
 - [FormRole](docs/FormRole.md)
 - [FormRoleArrayWrapper](docs/FormRoleArrayWrapper.md)
 - [FormRoleDto](docs/FormRoleDto.md)
 - [FormSubmissionsDto](docs/FormSubmissionsDto.md)
 - [FormSubmissionsWrapper](docs/FormSubmissionsWrapper.md)
 - [FormsItemArrayWrapper](docs/FormsItemArrayWrapper.md)
 - [FormsItemData](docs/FormsItemData.md)
 - [FormsItemDto](docs/FormsItemDto.md)
 - [GetPortalPrices200Response](docs/GetPortalPrices200Response.md)
 - [GetPortalPrices200ResponseLinksInner](docs/GetPortalPrices200ResponseLinksInner.md)
 - [GetReferenceDataDtoInteger](docs/GetReferenceDataDtoInteger.md)
 - [GobackConfig](docs/GobackConfig.md)
 - [GreetingSettingsRequestsDto](docs/GreetingSettingsRequestsDto.md)
 - [GroupArrayWrapper](docs/GroupArrayWrapper.md)
 - [GroupDto](docs/GroupDto.md)
 - [GroupMemberSecurityRequestArrayWrapper](docs/GroupMemberSecurityRequestArrayWrapper.md)
 - [GroupMemberSecurityRequestDto](docs/GroupMemberSecurityRequestDto.md)
 - [GroupRequestDto](docs/GroupRequestDto.md)
 - [GroupSummaryArrayWrapper](docs/GroupSummaryArrayWrapper.md)
 - [GroupSummaryDto](docs/GroupSummaryDto.md)
 - [GroupWrapper](docs/GroupWrapper.md)
 - [HideConfirmConvertRequestDto](docs/HideConfirmConvertRequestDto.md)
 - [HistoryAction](docs/HistoryAction.md)
 - [HistoryArrayWrapper](docs/HistoryArrayWrapper.md)
 - [HistoryData](docs/HistoryData.md)
 - [HistoryDto](docs/HistoryDto.md)
 - [ICompressWrapper](docs/ICompressWrapper.md)
 - [IPRestriction](docs/IPRestriction.md)
 - [IPRestrictionArrayWrapper](docs/IPRestrictionArrayWrapper.md)
 - [IPRestrictionsSettings](docs/IPRestrictionsSettings.md)
 - [IPRestrictionsSettingsWrapper](docs/IPRestrictionsSettingsWrapper.md)
 - [IconRequest](docs/IconRequest.md)
 - [ImportableApiEntity](docs/ImportableApiEntity.md)
 - [InfoConfigDto](docs/InfoConfigDto.md)
 - [Int32Wrapper](docs/Int32Wrapper.md)
 - [Int64Wrapper](docs/Int64Wrapper.md)
 - [InvitationLinkCreateRequestDto](docs/InvitationLinkCreateRequestDto.md)
 - [InvitationLinkDeleteRequestDto](docs/InvitationLinkDeleteRequestDto.md)
 - [InvitationLinkDto](docs/InvitationLinkDto.md)
 - [InvitationLinkUpdateRequestDto](docs/InvitationLinkUpdateRequestDto.md)
 - [InvitationLinkWrapper](docs/InvitationLinkWrapper.md)
 - [InviteUsersRequestDto](docs/InviteUsersRequestDto.md)
 - [IpRestrictionBase](docs/IpRestrictionBase.md)
 - [IpRestrictionsDto](docs/IpRestrictionsDto.md)
 - [IpRestrictionsWrapper](docs/IpRestrictionsWrapper.md)
 - [IsDefaultWhiteLabelLogosArrayWrapper](docs/IsDefaultWhiteLabelLogosArrayWrapper.md)
 - [IsDefaultWhiteLabelLogosDto](docs/IsDefaultWhiteLabelLogosDto.md)
 - [IsDefaultWhiteLabelLogosWrapper](docs/IsDefaultWhiteLabelLogosWrapper.md)
 - [ItemKeyValuePairBooleanString](docs/ItemKeyValuePairBooleanString.md)
 - [ItemKeyValuePairBooleanStringWrapper](docs/ItemKeyValuePairBooleanStringWrapper.md)
 - [ItemKeyValuePairObjectObject](docs/ItemKeyValuePairObjectObject.md)
 - [ItemKeyValuePairStringBoolean](docs/ItemKeyValuePairStringBoolean.md)
 - [ItemKeyValuePairStringLogoRequestsDto](docs/ItemKeyValuePairStringLogoRequestsDto.md)
 - [ItemKeyValuePairStringString](docs/ItemKeyValuePairStringString.md)
 - [LinkAccountRequestDto](docs/LinkAccountRequestDto.md)
 - [LinkType](docs/LinkType.md)
 - [Location](docs/Location.md)
 - [LocationType](docs/LocationType.md)
 - [LockFileParameters](docs/LockFileParameters.md)
 - [LoginEventArrayWrapper](docs/LoginEventArrayWrapper.md)
 - [LoginEventDto](docs/LoginEventDto.md)
 - [LoginProvider](docs/LoginProvider.md)
 - [LoginSettingsDto](docs/LoginSettingsDto.md)
 - [LoginSettingsRequestDto](docs/LoginSettingsRequestDto.md)
 - [LoginSettingsWrapper](docs/LoginSettingsWrapper.md)
 - [Logo](docs/Logo.md)
 - [LogoConfigDto](docs/LogoConfigDto.md)
 - [LogoCover](docs/LogoCover.md)
 - [LogoRequest](docs/LogoRequest.md)
 - [LogoRequestsDto](docs/LogoRequestsDto.md)
 - [MailDomainSettingsRequestsDto](docs/MailDomainSettingsRequestsDto.md)
 - [ManageFormFillingDtoInteger](docs/ManageFormFillingDtoInteger.md)
 - [MemberRequestDto](docs/MemberRequestDto.md)
 - [MembersRequest](docs/MembersRequest.md)
 - [MentionMessageWrapper](docs/MentionMessageWrapper.md)
 - [MentionWrapper](docs/MentionWrapper.md)
 - [MentionWrapperArrayWrapper](docs/MentionWrapperArrayWrapper.md)
 - [MessageAction](docs/MessageAction.md)
 - [MigratingApiFiles](docs/MigratingApiFiles.md)
 - [MigratingApiGroup](docs/MigratingApiGroup.md)
 - [MigratingApiUser](docs/MigratingApiUser.md)
 - [MigrationApiInfo](docs/MigrationApiInfo.md)
 - [MigrationStatusDto](docs/MigrationStatusDto.md)
 - [MigrationStatusWrapper](docs/MigrationStatusWrapper.md)
 - [MobilePhoneActivationStatus](docs/MobilePhoneActivationStatus.md)
 - [MobileRequestsDto](docs/MobileRequestsDto.md)
 - [Module](docs/Module.md)
 - [ModuleWrapper](docs/ModuleWrapper.md)
 - [MultiSizeLogoCover](docs/MultiSizeLogoCover.md)
 - [NewItemsDtoFileEntryBaseDto](docs/NewItemsDtoFileEntryBaseDto.md)
 - [NewItemsDtoRoomNewItemsDto](docs/NewItemsDtoRoomNewItemsDto.md)
 - [NewItemsFileEntryBaseArrayWrapper](docs/NewItemsFileEntryBaseArrayWrapper.md)
 - [NewItemsRoomNewItemsArrayWrapper](docs/NewItemsRoomNewItemsArrayWrapper.md)
 - [NotificationChannelDto](docs/NotificationChannelDto.md)
 - [NotificationChannelStatusDto](docs/NotificationChannelStatusDto.md)
 - [NotificationChannelStatusWrapper](docs/NotificationChannelStatusWrapper.md)
 - [NotificationSettingsDto](docs/NotificationSettingsDto.md)
 - [NotificationSettingsRequestsDto](docs/NotificationSettingsRequestsDto.md)
 - [NotificationSettingsWrapper](docs/NotificationSettingsWrapper.md)
 - [NotificationType](docs/NotificationType.md)
 - [OAuth20Token](docs/OAuth20Token.md)
 - [ObjectArrayWrapper](docs/ObjectArrayWrapper.md)
 - [ObjectWrapper](docs/ObjectWrapper.md)
 - [OperationDto](docs/OperationDto.md)
 - [OperationOrderType](docs/OperationOrderType.md)
 - [OperationStatus](docs/OperationStatus.md)
 - [OperationType](docs/OperationType.md)
 - [Options](docs/Options.md)
 - [OrderBy](docs/OrderBy.md)
 - [OrderRequestDto](docs/OrderRequestDto.md)
 - [OrdersItemRequestDtoInteger](docs/OrdersItemRequestDtoInteger.md)
 - [OrdersRequestDtoInteger](docs/OrdersRequestDtoInteger.md)
 - [OwnerChangeInstructionsDto](docs/OwnerChangeInstructionsDto.md)
 - [OwnerChangeInstructionsWrapper](docs/OwnerChangeInstructionsWrapper.md)
 - [OwnerIdSettingsRequestDto](docs/OwnerIdSettingsRequestDto.md)
 - [PageableModificationResponse](docs/PageableModificationResponse.md)
 - [PageableResponse](docs/PageableResponse.md)
 - [PageableResponseClientInfoResponse](docs/PageableResponseClientInfoResponse.md)
 - [Paragraph](docs/Paragraph.md)
 - [PasswordHasher](docs/PasswordHasher.md)
 - [PasswordSettingsDto](docs/PasswordSettingsDto.md)
 - [PasswordSettingsRequestsDto](docs/PasswordSettingsRequestsDto.md)
 - [PasswordSettingsWrapper](docs/PasswordSettingsWrapper.md)
 - [PaymentCalculation](docs/PaymentCalculation.md)
 - [PaymentCalculationWrapper](docs/PaymentCalculationWrapper.md)
 - [PaymentMethodStatus](docs/PaymentMethodStatus.md)
 - [PaymentSettingsDto](docs/PaymentSettingsDto.md)
 - [PaymentSettingsWrapper](docs/PaymentSettingsWrapper.md)
 - [PaymentUrlRequestDto](docs/PaymentUrlRequestDto.md)
 - [Payments](docs/Payments.md)
 - [PermissionsConfig](docs/PermissionsConfig.md)
 - [PluginsConfig](docs/PluginsConfig.md)
 - [PluginsDto](docs/PluginsDto.md)
 - [PriceDto](docs/PriceDto.md)
 - [ProblemDetail](docs/ProblemDetail.md)
 - [ProductAdministratorDto](docs/ProductAdministratorDto.md)
 - [ProductAdministratorWrapper](docs/ProductAdministratorWrapper.md)
 - [ProductQuantityType](docs/ProductQuantityType.md)
 - [ProductType](docs/ProductType.md)
 - [ProviderArrayWrapper](docs/ProviderArrayWrapper.md)
 - [ProviderDto](docs/ProviderDto.md)
 - [ProviderFilter](docs/ProviderFilter.md)
 - [QuantityRequestDto](docs/QuantityRequestDto.md)
 - [Quota](docs/Quota.md)
 - [QuotaArrayWrapper](docs/QuotaArrayWrapper.md)
 - [QuotaDto](docs/QuotaDto.md)
 - [QuotaFilter](docs/QuotaFilter.md)
 - [QuotaScope](docs/QuotaScope.md)
 - [QuotaSettingsRequestsDto](docs/QuotaSettingsRequestsDto.md)
 - [QuotaSettingsRequestsDtoDefaultQuota](docs/QuotaSettingsRequestsDtoDefaultQuota.md)
 - [QuotaState](docs/QuotaState.md)
 - [QuotaWrapper](docs/QuotaWrapper.md)
 - [RecaptchaType](docs/RecaptchaType.md)
 - [RecentConfig](docs/RecentConfig.md)
 - [RegStatus](docs/RegStatus.md)
 - [ReportDto](docs/ReportDto.md)
 - [ReportWrapper](docs/ReportWrapper.md)
 - [RestrictedModelsResponse](docs/RestrictedModelsResponse.md)
 - [RestrictedModelsResponseWrapper](docs/RestrictedModelsResponseWrapper.md)
 - [ReviewConfig](docs/ReviewConfig.md)
 - [RoomDataLifetimeDto](docs/RoomDataLifetimeDto.md)
 - [RoomDataLifetimePeriod](docs/RoomDataLifetimePeriod.md)
 - [RoomFromTemplateStatusDto](docs/RoomFromTemplateStatusDto.md)
 - [RoomFromTemplateStatusWrapper](docs/RoomFromTemplateStatusWrapper.md)
 - [RoomGroupArrayWrapper](docs/RoomGroupArrayWrapper.md)
 - [RoomGroupDto](docs/RoomGroupDto.md)
 - [RoomGroupRequestDto](docs/RoomGroupRequestDto.md)
 - [RoomGroupWrapper](docs/RoomGroupWrapper.md)
 - [RoomInvitation](docs/RoomInvitation.md)
 - [RoomInvitationRequest](docs/RoomInvitationRequest.md)
 - [RoomLinkRequest](docs/RoomLinkRequest.md)
 - [RoomNewItemsDto](docs/RoomNewItemsDto.md)
 - [RoomPrivacyFilter](docs/RoomPrivacyFilter.md)
 - [RoomSecurityDto](docs/RoomSecurityDto.md)
 - [RoomSecurityError](docs/RoomSecurityError.md)
 - [RoomSecurityWrapper](docs/RoomSecurityWrapper.md)
 - [RoomTemplateDto](docs/RoomTemplateDto.md)
 - [RoomTemplateStatusDto](docs/RoomTemplateStatusDto.md)
 - [RoomTemplateStatusWrapper](docs/RoomTemplateStatusWrapper.md)
 - [RoomType](docs/RoomType.md)
 - [RoomsNotificationSettingsDto](docs/RoomsNotificationSettingsDto.md)
 - [RoomsNotificationSettingsWrapper](docs/RoomsNotificationSettingsWrapper.md)
 - [RoomsNotificationsSettingsRequestDto](docs/RoomsNotificationsSettingsRequestDto.md)
 - [Run](docs/Run.md)
 - [STRINGArrayWrapper](docs/STRINGArrayWrapper.md)
 - [SalesRequestsDto](docs/SalesRequestsDto.md)
 - [SaveAsPdfInteger](docs/SaveAsPdfInteger.md)
 - [SaveFormRoleMappingDtoInteger](docs/SaveFormRoleMappingDtoInteger.md)
 - [ScheduleDto](docs/ScheduleDto.md)
 - [ScheduleWrapper](docs/ScheduleWrapper.md)
 - [ScopeResponse](docs/ScopeResponse.md)
 - [SearchArea](docs/SearchArea.md)
 - [SecurityArrayWrapper](docs/SecurityArrayWrapper.md)
 - [SecurityDto](docs/SecurityDto.md)
 - [SecurityInfoRequestDto](docs/SecurityInfoRequestDto.md)
 - [SecurityInfoSimpleRequestDto](docs/SecurityInfoSimpleRequestDto.md)
 - [SecurityRequestsDto](docs/SecurityRequestsDto.md)
 - [SessionRequest](docs/SessionRequest.md)
 - [SetAppEnabledBody](docs/SetAppEnabledBody.md)
 - [SetAppSettingsBody](docs/SetAppSettingsBody.md)
 - [SetAppSettingsBodySettings](docs/SetAppSettingsBodySettings.md)
 - [SetManagerRequest](docs/SetManagerRequest.md)
 - [SetPublicDto](docs/SetPublicDto.md)
 - [SetRestrictedAiModelsRequestDto](docs/SetRestrictedAiModelsRequestDto.md)
 - [SettingsDto](docs/SettingsDto.md)
 - [SettingsRequestDto](docs/SettingsRequestDto.md)
 - [SettingsWrapper](docs/SettingsWrapper.md)
 - [ShareFilterType](docs/ShareFilterType.md)
 - [SignupAccountRequestDto](docs/SignupAccountRequestDto.md)
 - [Size](docs/Size.md)
 - [SmtpOperationStatusRequestsDto](docs/SmtpOperationStatusRequestsDto.md)
 - [SmtpOperationStatusRequestsWrapper](docs/SmtpOperationStatusRequestsWrapper.md)
 - [SmtpSettingsDto](docs/SmtpSettingsDto.md)
 - [SmtpSettingsWrapper](docs/SmtpSettingsWrapper.md)
 - [SortOrder](docs/SortOrder.md)
 - [SortedByType](docs/SortedByType.md)
 - [SsoCertificate](docs/SsoCertificate.md)
 - [SsoFieldMapping](docs/SsoFieldMapping.md)
 - [SsoIdpCertificateAdvanced](docs/SsoIdpCertificateAdvanced.md)
 - [SsoIdpSettings](docs/SsoIdpSettings.md)
 - [SsoSettingsRequestsDto](docs/SsoSettingsRequestsDto.md)
 - [SsoSettingsV2](docs/SsoSettingsV2.md)
 - [SsoSettingsV2Wrapper](docs/SsoSettingsV2Wrapper.md)
 - [SsoSpCertificateAdvanced](docs/SsoSpCertificateAdvanced.md)
 - [StartEdit](docs/StartEdit.md)
 - [StartFillingForm](docs/StartFillingForm.md)
 - [StartFillingMode](docs/StartFillingMode.md)
 - [StartReassignRequestDto](docs/StartReassignRequestDto.md)
 - [StartUpdateUserTypeDto](docs/StartUpdateUserTypeDto.md)
 - [Status](docs/Status.md)
 - [StorageArrayWrapper](docs/StorageArrayWrapper.md)
 - [StorageDto](docs/StorageDto.md)
 - [StorageEncryptionRequestsDto](docs/StorageEncryptionRequestsDto.md)
 - [StorageFilter](docs/StorageFilter.md)
 - [StorageRequestsDto](docs/StorageRequestsDto.md)
 - [StorageSettings](docs/StorageSettings.md)
 - [StorageSettingsWrapper](docs/StorageSettingsWrapper.md)
 - [StringWrapper](docs/StringWrapper.md)
 - [StudioDefaultPageSettings](docs/StudioDefaultPageSettings.md)
 - [StudioDefaultPageSettingsWrapper](docs/StudioDefaultPageSettingsWrapper.md)
 - [SubAccount](docs/SubAccount.md)
 - [SubjectType](docs/SubjectType.md)
 - [SubmitForm](docs/SubmitForm.md)
 - [SubscriptionBalanceInfo](docs/SubscriptionBalanceInfo.md)
 - [SubscriptionBalanceInfoWrapper](docs/SubscriptionBalanceInfoWrapper.md)
 - [Tariff](docs/Tariff.md)
 - [TariffState](docs/TariffState.md)
 - [TariffWrapper](docs/TariffWrapper.md)
 - [TaskProgressResponseDto](docs/TaskProgressResponseDto.md)
 - [TaskProgressResponseWrapper](docs/TaskProgressResponseWrapper.md)
 - [TelegramStatusDto](docs/TelegramStatusDto.md)
 - [TelegramStatusWrapper](docs/TelegramStatusWrapper.md)
 - [TemplatesConfig](docs/TemplatesConfig.md)
 - [TemplatesRequestDto](docs/TemplatesRequestDto.md)
 - [TenantAiAccessSettings](docs/TenantAiAccessSettings.md)
 - [TenantAiAccessSettingsDto](docs/TenantAiAccessSettingsDto.md)
 - [TenantAiAccessSettingsWrapper](docs/TenantAiAccessSettingsWrapper.md)
 - [TenantAiAgentQuotaSettings](docs/TenantAiAgentQuotaSettings.md)
 - [TenantAiAgentQuotaSettingsWrapper](docs/TenantAiAgentQuotaSettingsWrapper.md)
 - [TenantAuditSettings](docs/TenantAuditSettings.md)
 - [TenantAuditSettingsResponseWrapper](docs/TenantAuditSettingsResponseWrapper.md)
 - [TenantAuditSettingsWrapper](docs/TenantAuditSettingsWrapper.md)
 - [TenantBannerSettings](docs/TenantBannerSettings.md)
 - [TenantBannerSettingsDto](docs/TenantBannerSettingsDto.md)
 - [TenantBannerSettingsWrapper](docs/TenantBannerSettingsWrapper.md)
 - [TenantDeepLinkSettings](docs/TenantDeepLinkSettings.md)
 - [TenantDeepLinkSettingsWrapper](docs/TenantDeepLinkSettingsWrapper.md)
 - [TenantDevToolsAccessSettings](docs/TenantDevToolsAccessSettings.md)
 - [TenantDevToolsAccessSettingsDto](docs/TenantDevToolsAccessSettingsDto.md)
 - [TenantDevToolsAccessSettingsWrapper](docs/TenantDevToolsAccessSettingsWrapper.md)
 - [TenantDomainValidator](docs/TenantDomainValidator.md)
 - [TenantDto](docs/TenantDto.md)
 - [TenantEntityQuotaSettings](docs/TenantEntityQuotaSettings.md)
 - [TenantIndustry](docs/TenantIndustry.md)
 - [TenantQuota](docs/TenantQuota.md)
 - [TenantQuotaFeatureDto](docs/TenantQuotaFeatureDto.md)
 - [TenantQuotaSettings](docs/TenantQuotaSettings.md)
 - [TenantQuotaSettingsRequestsDto](docs/TenantQuotaSettingsRequestsDto.md)
 - [TenantQuotaSettingsWrapper](docs/TenantQuotaSettingsWrapper.md)
 - [TenantQuotaWrapper](docs/TenantQuotaWrapper.md)
 - [TenantRoomQuotaSettings](docs/TenantRoomQuotaSettings.md)
 - [TenantRoomQuotaSettingsWrapper](docs/TenantRoomQuotaSettingsWrapper.md)
 - [TenantStatus](docs/TenantStatus.md)
 - [TenantTrustedDomainsType](docs/TenantTrustedDomainsType.md)
 - [TenantUserInvitationSettingsDto](docs/TenantUserInvitationSettingsDto.md)
 - [TenantUserInvitationSettingsRequestDto](docs/TenantUserInvitationSettingsRequestDto.md)
 - [TenantUserInvitationSettingsWrapper](docs/TenantUserInvitationSettingsWrapper.md)
 - [TenantUserQuotaSettings](docs/TenantUserQuotaSettings.md)
 - [TenantUserQuotaSettingsWrapper](docs/TenantUserQuotaSettingsWrapper.md)
 - [TenantWalletService](docs/TenantWalletService.md)
 - [TenantWalletServiceSettings](docs/TenantWalletServiceSettings.md)
 - [TenantWalletServiceSettingsWrapper](docs/TenantWalletServiceSettingsWrapper.md)
 - [TenantWalletSettings](docs/TenantWalletSettings.md)
 - [TenantWalletSettingsResponseWrapper](docs/TenantWalletSettingsResponseWrapper.md)
 - [TenantWalletSettingsWrapper](docs/TenantWalletSettingsWrapper.md)
 - [TenantWrapper](docs/TenantWrapper.md)
 - [TerminateRequestDto](docs/TerminateRequestDto.md)
 - [TfaAppCodeArrayWrapper](docs/TfaAppCodeArrayWrapper.md)
 - [TfaAppCodeDto](docs/TfaAppCodeDto.md)
 - [TfaConfirmDataDto](docs/TfaConfirmDataDto.md)
 - [TfaConfirmDataWrapper](docs/TfaConfirmDataWrapper.md)
 - [TfaRequestsDto](docs/TfaRequestsDto.md)
 - [TfaRequestsDtoType](docs/TfaRequestsDtoType.md)
 - [TfaSettingsArrayWrapper](docs/TfaSettingsArrayWrapper.md)
 - [TfaSettingsDto](docs/TfaSettingsDto.md)
 - [TfaSetupCodeDto](docs/TfaSetupCodeDto.md)
 - [TfaSetupCodeWrapper](docs/TfaSetupCodeWrapper.md)
 - [TfaValidateRequestsDto](docs/TfaValidateRequestsDto.md)
 - [ThirdPartyBackupRequestDto](docs/ThirdPartyBackupRequestDto.md)
 - [ThirdPartyParams](docs/ThirdPartyParams.md)
 - [ThirdPartyParamsArrayWrapper](docs/ThirdPartyParamsArrayWrapper.md)
 - [ThirdPartyRequestDto](docs/ThirdPartyRequestDto.md)
 - [Thumbnail](docs/Thumbnail.md)
 - [ThumbnailsDataDto](docs/ThumbnailsDataDto.md)
 - [ThumbnailsDataWrapper](docs/ThumbnailsDataWrapper.md)
 - [ThumbnailsRequest](docs/ThumbnailsRequest.md)
 - [TimezonesRequestsArrayWrapper](docs/TimezonesRequestsArrayWrapper.md)
 - [TimezonesRequestsDto](docs/TimezonesRequestsDto.md)
 - [TopUpDepositRequestDto](docs/TopUpDepositRequestDto.md)
 - [TransactionInfo](docs/TransactionInfo.md)
 - [TurnOnAdminMessageSettingsRequestDto](docs/TurnOnAdminMessageSettingsRequestDto.md)
 - [UpcomingPaymentArrayWrapper](docs/UpcomingPaymentArrayWrapper.md)
 - [UpcomingPaymentDto](docs/UpcomingPaymentDto.md)
 - [UpdateApiKeyRequest](docs/UpdateApiKeyRequest.md)
 - [UpdateClientRequest](docs/UpdateClientRequest.md)
 - [UpdateComment](docs/UpdateComment.md)
 - [UpdateFile](docs/UpdateFile.md)
 - [UpdateGroupRequest](docs/UpdateGroupRequest.md)
 - [UpdateMemberRequestDto](docs/UpdateMemberRequestDto.md)
 - [UpdateMembersQuotaRequestDto](docs/UpdateMembersQuotaRequestDto.md)
 - [UpdateMembersQuotaRequestDtoQuota](docs/UpdateMembersQuotaRequestDtoQuota.md)
 - [UpdateMembersRequestDto](docs/UpdateMembersRequestDto.md)
 - [UpdatePhotoMemberRequest](docs/UpdatePhotoMemberRequest.md)
 - [UpdateRoomGroupRequest](docs/UpdateRoomGroupRequest.md)
 - [UpdateRoomRequest](docs/UpdateRoomRequest.md)
 - [UpdateRoomsQuotaRequestDtoInteger](docs/UpdateRoomsQuotaRequestDtoInteger.md)
 - [UpdateRoomsRoomIdsRequestDtoInteger](docs/UpdateRoomsRoomIdsRequestDtoInteger.md)
 - [UpdateTagRequestDto](docs/UpdateTagRequestDto.md)
 - [UpdateWebhooksConfigRequestsDto](docs/UpdateWebhooksConfigRequestsDto.md)
 - [UploadResultDto](docs/UploadResultDto.md)
 - [UploadResultWrapper](docs/UploadResultWrapper.md)
 - [UploadSessionResponseDtoInteger](docs/UploadSessionResponseDtoInteger.md)
 - [UploadSessionResponseIntegerWrapper](docs/UploadSessionResponseIntegerWrapper.md)
 - [UsageSpaceStatItemArrayWrapper](docs/UsageSpaceStatItemArrayWrapper.md)
 - [UsageSpaceStatItemDto](docs/UsageSpaceStatItemDto.md)
 - [UserConfig](docs/UserConfig.md)
 - [UserExistsResponseDto](docs/UserExistsResponseDto.md)
 - [UserExistsResponseWrapper](docs/UserExistsResponseWrapper.md)
 - [UserInfo](docs/UserInfo.md)
 - [UserInfoWrapper](docs/UserInfoWrapper.md)
 - [UserInvitation](docs/UserInvitation.md)
 - [UserInvitationRequestDto](docs/UserInvitationRequestDto.md)
 - [ValidationResult](docs/ValidationResult.md)
 - [VectorizationStatus](docs/VectorizationStatus.md)
 - [WalletQuantityRequestDto](docs/WalletQuantityRequestDto.md)
 - [WalletServiceArrayWrapper](docs/WalletServiceArrayWrapper.md)
 - [WalletServiceDto](docs/WalletServiceDto.md)
 - [WalletServiceWrapper](docs/WalletServiceWrapper.md)
 - [WatermarkAdditions](docs/WatermarkAdditions.md)
 - [WatermarkDto](docs/WatermarkDto.md)
 - [WatermarkOnDraw](docs/WatermarkOnDraw.md)
 - [WatermarkRequestDto](docs/WatermarkRequestDto.md)
 - [WebItemSecurityRequestsDto](docs/WebItemSecurityRequestsDto.md)
 - [WebItemsSecurityRequestsDto](docs/WebItemsSecurityRequestsDto.md)
 - [WebPluginArrayWrapper](docs/WebPluginArrayWrapper.md)
 - [WebPluginDto](docs/WebPluginDto.md)
 - [WebPluginRequests](docs/WebPluginRequests.md)
 - [WebPluginWrapper](docs/WebPluginWrapper.md)
 - [WebhookGroupStatus](docs/WebhookGroupStatus.md)
 - [WebhookRetryRequestsDto](docs/WebhookRetryRequestsDto.md)
 - [WebhookTrigger](docs/WebhookTrigger.md)
 - [WebhookTriggerArrayWrapper](docs/WebhookTriggerArrayWrapper.md)
 - [WebhookTriggerDto](docs/WebhookTriggerDto.md)
 - [WebhooksConfigDto](docs/WebhooksConfigDto.md)
 - [WebhooksConfigWithStatusArrayWrapper](docs/WebhooksConfigWithStatusArrayWrapper.md)
 - [WebhooksConfigWithStatusDto](docs/WebhooksConfigWithStatusDto.md)
 - [WebhooksConfigWrapper](docs/WebhooksConfigWrapper.md)
 - [WebhooksLogArrayWrapper](docs/WebhooksLogArrayWrapper.md)
 - [WebhooksLogDto](docs/WebhooksLogDto.md)
 - [WebhooksLogWrapper](docs/WebhooksLogWrapper.md)
 - [WhiteLabelItemArrayWrapper](docs/WhiteLabelItemArrayWrapper.md)
 - [WhiteLabelItemDto](docs/WhiteLabelItemDto.md)
 - [WhiteLabelItemPathDto](docs/WhiteLabelItemPathDto.md)
 - [WhiteLabelItemSizeDto](docs/WhiteLabelItemSizeDto.md)
 - [WhiteLabelLogoType](docs/WhiteLabelLogoType.md)
 - [WhiteLabelRequestsDto](docs/WhiteLabelRequestsDto.md)
 - [WizardRequestsDto](docs/WizardRequestsDto.md)
 - [WizardSettings](docs/WizardSettings.md)
 - [WizardSettingsWrapper](docs/WizardSettingsWrapper.md)
 - [XlsxReportResponseDto](docs/XlsxReportResponseDto.md)
 - [XlsxReportResponseWrapper](docs/XlsxReportResponseWrapper.md)

</details>

## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@onlyoffice.com

