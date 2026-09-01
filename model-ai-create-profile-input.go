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

// checks if the AiCreateProfileInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiCreateProfileInput{}

// AiCreateProfileInput Input for creating a new profile — the same shape as `Profile` without the engine-generated fields (`id`, `createdAt`).
type AiCreateProfileInput struct {
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
}

type _AiCreateProfileInput AiCreateProfileInput

// NewAiCreateProfileInput instantiates a new AiCreateProfileInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiCreateProfileInput(name string, providerType AiProviderType, baseUrl string, modelId string) *AiCreateProfileInput {
	this := AiCreateProfileInput{}
	this.Name = name
	this.ProviderType = providerType
	this.BaseUrl = baseUrl
	this.ModelId = modelId
	return &this
}

// NewAiCreateProfileInputWithDefaults instantiates a new AiCreateProfileInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiCreateProfileInputWithDefaults() *AiCreateProfileInput {
	this := AiCreateProfileInput{}
	return &this
}

// GetName returns the Name field value
func (o *AiCreateProfileInput) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiCreateProfileInput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AiCreateProfileInput) SetName(v string) {
	o.Name = v
}

// GetProviderType returns the ProviderType field value
func (o *AiCreateProfileInput) GetProviderType() AiProviderType {
	if o == nil {
		var ret AiProviderType
		return ret
	}

	return o.ProviderType
}

// GetProviderTypeOk returns a tuple with the ProviderType field value
// and a boolean to check if the value has been set.
func (o *AiCreateProfileInput) GetProviderTypeOk() (*AiProviderType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderType, true
}

// SetProviderType sets field value
func (o *AiCreateProfileInput) SetProviderType(v AiProviderType) {
	o.ProviderType = v
}

// GetBasedOn returns the BasedOn field value if set, zero value otherwise.
func (o *AiCreateProfileInput) GetBasedOn() AiBuiltinProviderType {
	if o == nil || IsNil(o.BasedOn) {
		var ret AiBuiltinProviderType
		return ret
	}
	return *o.BasedOn
}

// GetBasedOnOk returns a tuple with the BasedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiCreateProfileInput) GetBasedOnOk() (*AiBuiltinProviderType, bool) {
	if o == nil || IsNil(o.BasedOn) {
		return nil, false
	}
	return o.BasedOn, true
}

// HasBasedOn returns a boolean if a field has been set.
func (o *AiCreateProfileInput) IsBasedOnSet() bool {
	if o != nil && !IsNil(o.BasedOn) {
		return true
	}

	return false
}

// SetBasedOn gets a reference to the given AiBuiltinProviderType and assigns it to the BasedOn field.
func (o *AiCreateProfileInput) SetBasedOn(v AiBuiltinProviderType) {
	o.BasedOn = &v
}

// GetBaseUrl returns the BaseUrl field value
func (o *AiCreateProfileInput) GetBaseUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value
// and a boolean to check if the value has been set.
func (o *AiCreateProfileInput) GetBaseUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseUrl, true
}

// SetBaseUrl sets field value
func (o *AiCreateProfileInput) SetBaseUrl(v string) {
	o.BaseUrl = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *AiCreateProfileInput) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiCreateProfileInput) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *AiCreateProfileInput) IsKeySet() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *AiCreateProfileInput) SetKey(v string) {
	o.Key = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *AiCreateProfileInput) GetHeaders() map[string]string {
	if o == nil || IsNil(o.Headers) {
		var ret map[string]string
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiCreateProfileInput) GetHeadersOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Headers) {
		return map[string]string{}, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *AiCreateProfileInput) IsHeadersSet() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]string and assigns it to the Headers field.
func (o *AiCreateProfileInput) SetHeaders(v map[string]string) {
	o.Headers = v
}

// GetModelId returns the ModelId field value
func (o *AiCreateProfileInput) GetModelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value
// and a boolean to check if the value has been set.
func (o *AiCreateProfileInput) GetModelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelId, true
}

// SetModelId sets field value
func (o *AiCreateProfileInput) SetModelId(v string) {
	o.ModelId = v
}

// GetReasoning returns the Reasoning field value if set, zero value otherwise.
func (o *AiCreateProfileInput) GetReasoning() bool {
	if o == nil || IsNil(o.Reasoning) {
		var ret bool
		return ret
	}
	return *o.Reasoning
}

// GetReasoningOk returns a tuple with the Reasoning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiCreateProfileInput) GetReasoningOk() (*bool, bool) {
	if o == nil || IsNil(o.Reasoning) {
		return nil, false
	}
	return o.Reasoning, true
}

// HasReasoning returns a boolean if a field has been set.
func (o *AiCreateProfileInput) IsReasoningSet() bool {
	if o != nil && !IsNil(o.Reasoning) {
		return true
	}

	return false
}

// SetReasoning gets a reference to the given bool and assigns it to the Reasoning field.
func (o *AiCreateProfileInput) SetReasoning(v bool) {
	o.Reasoning = &v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *AiCreateProfileInput) GetCapabilities() float32 {
	if o == nil || IsNil(o.Capabilities) {
		var ret float32
		return ret
	}
	return *o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiCreateProfileInput) GetCapabilitiesOk() (*float32, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *AiCreateProfileInput) IsCapabilitiesSet() bool {
	if o != nil && !IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given float32 and assigns it to the Capabilities field.
func (o *AiCreateProfileInput) SetCapabilities(v float32) {
	o.Capabilities = &v
}

// GetCanUseTool returns the CanUseTool field value if set, zero value otherwise.
func (o *AiCreateProfileInput) GetCanUseTool() bool {
	if o == nil || IsNil(o.CanUseTool) {
		var ret bool
		return ret
	}
	return *o.CanUseTool
}

// GetCanUseToolOk returns a tuple with the CanUseTool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiCreateProfileInput) GetCanUseToolOk() (*bool, bool) {
	if o == nil || IsNil(o.CanUseTool) {
		return nil, false
	}
	return o.CanUseTool, true
}

// HasCanUseTool returns a boolean if a field has been set.
func (o *AiCreateProfileInput) IsCanUseToolSet() bool {
	if o != nil && !IsNil(o.CanUseTool) {
		return true
	}

	return false
}

// SetCanUseTool gets a reference to the given bool and assigns it to the CanUseTool field.
func (o *AiCreateProfileInput) SetCanUseTool(v bool) {
	o.CanUseTool = &v
}

// GetUseResponsesApi returns the UseResponsesApi field value if set, zero value otherwise.
func (o *AiCreateProfileInput) GetUseResponsesApi() bool {
	if o == nil || IsNil(o.UseResponsesApi) {
		var ret bool
		return ret
	}
	return *o.UseResponsesApi
}

// GetUseResponsesApiOk returns a tuple with the UseResponsesApi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiCreateProfileInput) GetUseResponsesApiOk() (*bool, bool) {
	if o == nil || IsNil(o.UseResponsesApi) {
		return nil, false
	}
	return o.UseResponsesApi, true
}

// HasUseResponsesApi returns a boolean if a field has been set.
func (o *AiCreateProfileInput) IsUseResponsesApiSet() bool {
	if o != nil && !IsNil(o.UseResponsesApi) {
		return true
	}

	return false
}

// SetUseResponsesApi gets a reference to the given bool and assigns it to the UseResponsesApi field.
func (o *AiCreateProfileInput) SetUseResponsesApi(v bool) {
	o.UseResponsesApi = &v
}

// GetIsCloudProvider returns the IsCloudProvider field value if set, zero value otherwise.
func (o *AiCreateProfileInput) GetIsCloudProvider() bool {
	if o == nil || IsNil(o.IsCloudProvider) {
		var ret bool
		return ret
	}
	return *o.IsCloudProvider
}

// GetIsCloudProviderOk returns a tuple with the IsCloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiCreateProfileInput) GetIsCloudProviderOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCloudProvider) {
		return nil, false
	}
	return o.IsCloudProvider, true
}

// HasIsCloudProvider returns a boolean if a field has been set.
func (o *AiCreateProfileInput) IsIsCloudProviderSet() bool {
	if o != nil && !IsNil(o.IsCloudProvider) {
		return true
	}

	return false
}

// SetIsCloudProvider gets a reference to the given bool and assigns it to the IsCloudProvider field.
func (o *AiCreateProfileInput) SetIsCloudProvider(v bool) {
	o.IsCloudProvider = &v
}

// GetUseProxy returns the UseProxy field value if set, zero value otherwise.
func (o *AiCreateProfileInput) GetUseProxy() bool {
	if o == nil || IsNil(o.UseProxy) {
		var ret bool
		return ret
	}
	return *o.UseProxy
}

// GetUseProxyOk returns a tuple with the UseProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiCreateProfileInput) GetUseProxyOk() (*bool, bool) {
	if o == nil || IsNil(o.UseProxy) {
		return nil, false
	}
	return o.UseProxy, true
}

// HasUseProxy returns a boolean if a field has been set.
func (o *AiCreateProfileInput) IsUseProxySet() bool {
	if o != nil && !IsNil(o.UseProxy) {
		return true
	}

	return false
}

// SetUseProxy gets a reference to the given bool and assigns it to the UseProxy field.
func (o *AiCreateProfileInput) SetUseProxy(v bool) {
	o.UseProxy = &v
}

func (o AiCreateProfileInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiCreateProfileInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	return toSerialize, nil
}

func (o *AiCreateProfileInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varAiCreateProfileInput := _AiCreateProfileInput{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiCreateProfileInput)

	if err != nil {
		return err
	}

	*o = AiCreateProfileInput(varAiCreateProfileInput)

	return err
}

type NullableAiCreateProfileInput struct {
	value *AiCreateProfileInput
	isSet bool
}

func (v NullableAiCreateProfileInput) Get() *AiCreateProfileInput {
	return v.value
}

func (v *NullableAiCreateProfileInput) Set(val *AiCreateProfileInput) {
	v.value = val
	v.isSet = true
}

func (v NullableAiCreateProfileInput) IsSet() bool {
	return v.isSet
}

func (v *NullableAiCreateProfileInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiCreateProfileInput(val *AiCreateProfileInput) *NullableAiCreateProfileInput {
	return &NullableAiCreateProfileInput{value: val, isSet: true}
}

func (v NullableAiCreateProfileInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiCreateProfileInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

