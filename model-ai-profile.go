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
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AiProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiProfile{}

// AiProfile Complete AI provider + model configuration saved by the user. Profiles are the primary way users save and reuse provider configurations.
type AiProfile struct {
	// Unique profile identifier (UUID).
	Id string `json:"id"`
	// User-defined profile display name.
	Name string `json:"name"`
	// Provider type for this profile. Use `external` to delegate all HTTP transport to `PlatformAdapter.externalFetch` while reusing an existing provider's response parser — see `Profile.basedOn` for the format selector.
	ProviderType AiProviderType `json:"providerType"`
	// Selects the response-format parser used by the `external` provider. Ignored for any other `providerType`.  Supported values are `openai`, `anthropic`, `mistral` and `openrouter`. Remaining values (`genai`, `stabilityai`, …) are accepted by the type but not yet implemented; passing one raises an error at request time.
	BasedOn *AiBuiltinProviderType `json:"basedOn,omitempty"`
	// Base URL of the provider API.
	BaseUrl string `json:"baseUrl"`
	// API key or token. Optional for local providers.
	Key *string `json:"key,omitempty"`
	// Extra HTTP headers sent with every request to this provider. Merged into the SDK client's default headers; an explicit `Authorization` here wins over the one derived from `key`. Honoured by the OpenAI-family providers.
	Headers map[string]string `json:"headers,omitempty"`
	// Selected model ID within this provider.
	ModelId string `json:"modelId"`
	// Whether extended thinking is enabled for this profile's model.
	Reasoning *bool `json:"reasoning,omitempty"`
	// Extended-thinking capabilities of the selected model as reported by the provider's catalogue at save time (see `Model.reasoningSupport`). When present the composer's Effort row follows it exactly; when absent the provider's id-based table answers. Hosts persist it with the rest of the profile.
	ReasoningSupport *AiReasoningSupport `json:"reasoningSupport,omitempty"`
	// Bitmask of capabilities supported by the selected model.
	Capabilities *float32 `json:"capabilities,omitempty"`
	// Result of the live tool-capability probe performed at create time and on changes to `modelId` / `providerType` / `baseUrl`. `undefined` means the probe has never run for this profile (legacy record).
	CanUseTool *bool `json:"canUseTool,omitempty"`
	// Result of the live Responses-API probe (parallel to `canUseTool`). `true` means the model speaks `/v1/responses` and the OpenAI provider must route through `client.responses.create` — required for gpt-5+ reasoning models that reject `reasoning_effort` together with `tools` on `/v1/chat/completions`. Probed at create time and whenever `modelId` / `providerType` / `baseUrl` change. `undefined` means the probe never ran (legacy record) — readers treat that as `false`.
	UseResponsesApi *bool `json:"useResponsesApi,omitempty"`
	// Whether this profile uses a cloud-hosted provider (e.g. ONLYOFFICE DocSpace).
	IsCloudProvider *bool `json:"isCloudProvider,omitempty"`
	// Route every provider request through the host's `fetchProxy` instead of the global `fetch`. Useful when the host runs the widget in a sandbox without direct network access (CORS, custom auth, etc.). Has no effect when the `PlatformAdapter.fetchProxy` is not configured.
	UseProxy *bool `json:"useProxy,omitempty"`
	// Creation timestamp (ms since epoch). Used to sort the AI models list newest-first.
	CreatedAt *float32 `json:"createdAt,omitempty"`
}

type _AiProfile AiProfile

// NewAiProfile instantiates a new AiProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiProfile(id string, name string, providerType AiProviderType, baseUrl string, modelId string) *AiProfile {
	this := AiProfile{}
	this.Id = id
	this.Name = name
	this.ProviderType = providerType
	this.BaseUrl = baseUrl
	this.ModelId = modelId
	return &this
}

// NewAiProfileWithDefaults instantiates a new AiProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiProfileWithDefaults() *AiProfile {
	this := AiProfile{}
	return &this
}

// GetId returns the Id field value
func (o *AiProfile) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AiProfile) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AiProfile) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AiProfile) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiProfile) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AiProfile) SetName(v string) {
	o.Name = v
}

// GetProviderType returns the ProviderType field value
func (o *AiProfile) GetProviderType() AiProviderType {
	if o == nil {
		var ret AiProviderType
		return ret
	}

	return o.ProviderType
}

// GetProviderTypeOk returns a tuple with the ProviderType field value
// and a boolean to check if the value has been set.
func (o *AiProfile) GetProviderTypeOk() (*AiProviderType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderType, true
}

// SetProviderType sets field value
func (o *AiProfile) SetProviderType(v AiProviderType) {
	o.ProviderType = v
}

// GetBasedOn returns the BasedOn field value if set, zero value otherwise.
func (o *AiProfile) GetBasedOn() AiBuiltinProviderType {
	if o == nil || IsNil(o.BasedOn) {
		var ret AiBuiltinProviderType
		return ret
	}
	return *o.BasedOn
}

// GetBasedOnOk returns a tuple with the BasedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiProfile) GetBasedOnOk() (*AiBuiltinProviderType, bool) {
	if o == nil || IsNil(o.BasedOn) {
		return nil, false
	}
	return o.BasedOn, true
}

// HasBasedOn returns a boolean if a field has been set.
func (o *AiProfile) IsBasedOnSet() bool {
	if o != nil && !IsNil(o.BasedOn) {
		return true
	}

	return false
}

// SetBasedOn gets a reference to the given AiBuiltinProviderType and assigns it to the BasedOn field.
func (o *AiProfile) SetBasedOn(v AiBuiltinProviderType) {
	o.BasedOn = &v
}

// GetBaseUrl returns the BaseUrl field value
func (o *AiProfile) GetBaseUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value
// and a boolean to check if the value has been set.
func (o *AiProfile) GetBaseUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseUrl, true
}

// SetBaseUrl sets field value
func (o *AiProfile) SetBaseUrl(v string) {
	o.BaseUrl = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *AiProfile) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiProfile) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *AiProfile) IsKeySet() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *AiProfile) SetKey(v string) {
	o.Key = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *AiProfile) GetHeaders() map[string]string {
	if o == nil || IsNil(o.Headers) {
		var ret map[string]string
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiProfile) GetHeadersOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Headers) {
		return map[string]string{}, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *AiProfile) IsHeadersSet() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]string and assigns it to the Headers field.
func (o *AiProfile) SetHeaders(v map[string]string) {
	o.Headers = v
}

// GetModelId returns the ModelId field value
func (o *AiProfile) GetModelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value
// and a boolean to check if the value has been set.
func (o *AiProfile) GetModelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelId, true
}

// SetModelId sets field value
func (o *AiProfile) SetModelId(v string) {
	o.ModelId = v
}

// GetReasoning returns the Reasoning field value if set, zero value otherwise.
func (o *AiProfile) GetReasoning() bool {
	if o == nil || IsNil(o.Reasoning) {
		var ret bool
		return ret
	}
	return *o.Reasoning
}

// GetReasoningOk returns a tuple with the Reasoning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiProfile) GetReasoningOk() (*bool, bool) {
	if o == nil || IsNil(o.Reasoning) {
		return nil, false
	}
	return o.Reasoning, true
}

// HasReasoning returns a boolean if a field has been set.
func (o *AiProfile) IsReasoningSet() bool {
	if o != nil && !IsNil(o.Reasoning) {
		return true
	}

	return false
}

// SetReasoning gets a reference to the given bool and assigns it to the Reasoning field.
func (o *AiProfile) SetReasoning(v bool) {
	o.Reasoning = &v
}

// GetReasoningSupport returns the ReasoningSupport field value if set, zero value otherwise.
func (o *AiProfile) GetReasoningSupport() AiReasoningSupport {
	if o == nil || IsNil(o.ReasoningSupport) {
		var ret AiReasoningSupport
		return ret
	}
	return *o.ReasoningSupport
}

// GetReasoningSupportOk returns a tuple with the ReasoningSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiProfile) GetReasoningSupportOk() (*AiReasoningSupport, bool) {
	if o == nil || IsNil(o.ReasoningSupport) {
		return nil, false
	}
	return o.ReasoningSupport, true
}

// HasReasoningSupport returns a boolean if a field has been set.
func (o *AiProfile) IsReasoningSupportSet() bool {
	if o != nil && !IsNil(o.ReasoningSupport) {
		return true
	}

	return false
}

// SetReasoningSupport gets a reference to the given AiReasoningSupport and assigns it to the ReasoningSupport field.
func (o *AiProfile) SetReasoningSupport(v AiReasoningSupport) {
	o.ReasoningSupport = &v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *AiProfile) GetCapabilities() float32 {
	if o == nil || IsNil(o.Capabilities) {
		var ret float32
		return ret
	}
	return *o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiProfile) GetCapabilitiesOk() (*float32, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *AiProfile) IsCapabilitiesSet() bool {
	if o != nil && !IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given float32 and assigns it to the Capabilities field.
func (o *AiProfile) SetCapabilities(v float32) {
	o.Capabilities = &v
}

// GetCanUseTool returns the CanUseTool field value if set, zero value otherwise.
func (o *AiProfile) GetCanUseTool() bool {
	if o == nil || IsNil(o.CanUseTool) {
		var ret bool
		return ret
	}
	return *o.CanUseTool
}

// GetCanUseToolOk returns a tuple with the CanUseTool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiProfile) GetCanUseToolOk() (*bool, bool) {
	if o == nil || IsNil(o.CanUseTool) {
		return nil, false
	}
	return o.CanUseTool, true
}

// HasCanUseTool returns a boolean if a field has been set.
func (o *AiProfile) IsCanUseToolSet() bool {
	if o != nil && !IsNil(o.CanUseTool) {
		return true
	}

	return false
}

// SetCanUseTool gets a reference to the given bool and assigns it to the CanUseTool field.
func (o *AiProfile) SetCanUseTool(v bool) {
	o.CanUseTool = &v
}

// GetUseResponsesApi returns the UseResponsesApi field value if set, zero value otherwise.
func (o *AiProfile) GetUseResponsesApi() bool {
	if o == nil || IsNil(o.UseResponsesApi) {
		var ret bool
		return ret
	}
	return *o.UseResponsesApi
}

// GetUseResponsesApiOk returns a tuple with the UseResponsesApi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiProfile) GetUseResponsesApiOk() (*bool, bool) {
	if o == nil || IsNil(o.UseResponsesApi) {
		return nil, false
	}
	return o.UseResponsesApi, true
}

// HasUseResponsesApi returns a boolean if a field has been set.
func (o *AiProfile) IsUseResponsesApiSet() bool {
	if o != nil && !IsNil(o.UseResponsesApi) {
		return true
	}

	return false
}

// SetUseResponsesApi gets a reference to the given bool and assigns it to the UseResponsesApi field.
func (o *AiProfile) SetUseResponsesApi(v bool) {
	o.UseResponsesApi = &v
}

// GetIsCloudProvider returns the IsCloudProvider field value if set, zero value otherwise.
func (o *AiProfile) GetIsCloudProvider() bool {
	if o == nil || IsNil(o.IsCloudProvider) {
		var ret bool
		return ret
	}
	return *o.IsCloudProvider
}

// GetIsCloudProviderOk returns a tuple with the IsCloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiProfile) GetIsCloudProviderOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCloudProvider) {
		return nil, false
	}
	return o.IsCloudProvider, true
}

// HasIsCloudProvider returns a boolean if a field has been set.
func (o *AiProfile) IsIsCloudProviderSet() bool {
	if o != nil && !IsNil(o.IsCloudProvider) {
		return true
	}

	return false
}

// SetIsCloudProvider gets a reference to the given bool and assigns it to the IsCloudProvider field.
func (o *AiProfile) SetIsCloudProvider(v bool) {
	o.IsCloudProvider = &v
}

// GetUseProxy returns the UseProxy field value if set, zero value otherwise.
func (o *AiProfile) GetUseProxy() bool {
	if o == nil || IsNil(o.UseProxy) {
		var ret bool
		return ret
	}
	return *o.UseProxy
}

// GetUseProxyOk returns a tuple with the UseProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiProfile) GetUseProxyOk() (*bool, bool) {
	if o == nil || IsNil(o.UseProxy) {
		return nil, false
	}
	return o.UseProxy, true
}

// HasUseProxy returns a boolean if a field has been set.
func (o *AiProfile) IsUseProxySet() bool {
	if o != nil && !IsNil(o.UseProxy) {
		return true
	}

	return false
}

// SetUseProxy gets a reference to the given bool and assigns it to the UseProxy field.
func (o *AiProfile) SetUseProxy(v bool) {
	o.UseProxy = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AiProfile) GetCreatedAt() float32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret float32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiProfile) GetCreatedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AiProfile) IsCreatedAtSet() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given float32 and assigns it to the CreatedAt field.
func (o *AiProfile) SetCreatedAt(v float32) {
	o.CreatedAt = &v
}

func (o AiProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["providerType"] = o.ProviderType
	if !IsNil(o.BasedOn) {
		toSerialize["basedOn"] = o.BasedOn
	}
	toSerialize["baseUrl"] = o.BaseUrl
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	toSerialize["modelId"] = o.ModelId
	if !IsNil(o.Reasoning) {
		toSerialize["reasoning"] = o.Reasoning
	}
	if !IsNil(o.ReasoningSupport) {
		toSerialize["reasoningSupport"] = o.ReasoningSupport
	}
	if !IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	if !IsNil(o.CanUseTool) {
		toSerialize["canUseTool"] = o.CanUseTool
	}
	if !IsNil(o.UseResponsesApi) {
		toSerialize["useResponsesApi"] = o.UseResponsesApi
	}
	if !IsNil(o.IsCloudProvider) {
		toSerialize["isCloudProvider"] = o.IsCloudProvider
	}
	if !IsNil(o.UseProxy) {
		toSerialize["useProxy"] = o.UseProxy
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	return toSerialize, nil
}

func (o *AiProfile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"providerType",
		"baseUrl",
		"modelId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAiProfile := _AiProfile{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiProfile)

	if err != nil {
		return err
	}

	*o = AiProfile(varAiProfile)

	return err
}

type NullableAiProfile struct {
	value *AiProfile
	isSet bool
}

func (v NullableAiProfile) Get() *AiProfile {
	return v.value
}

func (v *NullableAiProfile) Set(val *AiProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableAiProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableAiProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiProfile(val *AiProfile) *NullableAiProfile {
	return &NullableAiProfile{value: val, isSet: true}
}

func (v NullableAiProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

