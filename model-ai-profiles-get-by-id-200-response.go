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

// checks if the AiProfilesGetById200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiProfilesGetById200Response{}

// AiProfilesGetById200Response struct for AiProfilesGetById200Response
type AiProfilesGetById200Response struct {
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
	// Creation timestamp (ms since epoch). Used to sort the AI models list newest-first.
	CreatedAt *float32 `json:"createdAt,omitempty"`
}

type _AiProfilesGetById200Response AiProfilesGetById200Response

// NewAiProfilesGetById200Response instantiates a new AiProfilesGetById200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiProfilesGetById200Response(id string, name string, providerType AiProviderType, baseUrl string, modelId string) *AiProfilesGetById200Response {
	this := AiProfilesGetById200Response{}
	this.Id = id
	this.Name = name
	this.ProviderType = providerType
	this.BaseUrl = baseUrl
	this.ModelId = modelId
	return &this
}

// NewAiProfilesGetById200ResponseWithDefaults instantiates a new AiProfilesGetById200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiProfilesGetById200ResponseWithDefaults() *AiProfilesGetById200Response {
	this := AiProfilesGetById200Response{}
	return &this
}

// GetId returns the Id field value
func (o *AiProfilesGetById200Response) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AiProfilesGetById200Response) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AiProfilesGetById200Response) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AiProfilesGetById200Response) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiProfilesGetById200Response) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AiProfilesGetById200Response) SetName(v string) {
	o.Name = v
}

// GetProviderType returns the ProviderType field value
func (o *AiProfilesGetById200Response) GetProviderType() AiProviderType {
	if o == nil {
		var ret AiProviderType
		return ret
	}

	return o.ProviderType
}

// GetProviderTypeOk returns a tuple with the ProviderType field value
// and a boolean to check if the value has been set.
func (o *AiProfilesGetById200Response) GetProviderTypeOk() (*AiProviderType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderType, true
}

// SetProviderType sets field value
func (o *AiProfilesGetById200Response) SetProviderType(v AiProviderType) {
	o.ProviderType = v
}

// GetBasedOn returns the BasedOn field value if set, zero value otherwise.
func (o *AiProfilesGetById200Response) GetBasedOn() AiBuiltinProviderType {
	if o == nil || IsNil(o.BasedOn) {
		var ret AiBuiltinProviderType
		return ret
	}
	return *o.BasedOn
}

// GetBasedOnOk returns a tuple with the BasedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiProfilesGetById200Response) GetBasedOnOk() (*AiBuiltinProviderType, bool) {
	if o == nil || IsNil(o.BasedOn) {
		return nil, false
	}
	return o.BasedOn, true
}

// HasBasedOn returns a boolean if a field has been set.
func (o *AiProfilesGetById200Response) IsBasedOnSet() bool {
	if o != nil && !IsNil(o.BasedOn) {
		return true
	}

	return false
}

// SetBasedOn gets a reference to the given AiBuiltinProviderType and assigns it to the BasedOn field.
func (o *AiProfilesGetById200Response) SetBasedOn(v AiBuiltinProviderType) {
	o.BasedOn = &v
}

// GetBaseUrl returns the BaseUrl field value
func (o *AiProfilesGetById200Response) GetBaseUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value
// and a boolean to check if the value has been set.
func (o *AiProfilesGetById200Response) GetBaseUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseUrl, true
}

// SetBaseUrl sets field value
func (o *AiProfilesGetById200Response) SetBaseUrl(v string) {
	o.BaseUrl = v
}

// GetModelId returns the ModelId field value
func (o *AiProfilesGetById200Response) GetModelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value
// and a boolean to check if the value has been set.
func (o *AiProfilesGetById200Response) GetModelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelId, true
}

// SetModelId sets field value
func (o *AiProfilesGetById200Response) SetModelId(v string) {
	o.ModelId = v
}

// GetReasoning returns the Reasoning field value if set, zero value otherwise.
func (o *AiProfilesGetById200Response) GetReasoning() bool {
	if o == nil || IsNil(o.Reasoning) {
		var ret bool
		return ret
	}
	return *o.Reasoning
}

// GetReasoningOk returns a tuple with the Reasoning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiProfilesGetById200Response) GetReasoningOk() (*bool, bool) {
	if o == nil || IsNil(o.Reasoning) {
		return nil, false
	}
	return o.Reasoning, true
}

// HasReasoning returns a boolean if a field has been set.
func (o *AiProfilesGetById200Response) IsReasoningSet() bool {
	if o != nil && !IsNil(o.Reasoning) {
		return true
	}

	return false
}

// SetReasoning gets a reference to the given bool and assigns it to the Reasoning field.
func (o *AiProfilesGetById200Response) SetReasoning(v bool) {
	o.Reasoning = &v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *AiProfilesGetById200Response) GetCapabilities() float32 {
	if o == nil || IsNil(o.Capabilities) {
		var ret float32
		return ret
	}
	return *o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiProfilesGetById200Response) GetCapabilitiesOk() (*float32, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *AiProfilesGetById200Response) IsCapabilitiesSet() bool {
	if o != nil && !IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given float32 and assigns it to the Capabilities field.
func (o *AiProfilesGetById200Response) SetCapabilities(v float32) {
	o.Capabilities = &v
}

// GetCanUseTool returns the CanUseTool field value if set, zero value otherwise.
func (o *AiProfilesGetById200Response) GetCanUseTool() bool {
	if o == nil || IsNil(o.CanUseTool) {
		var ret bool
		return ret
	}
	return *o.CanUseTool
}

// GetCanUseToolOk returns a tuple with the CanUseTool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiProfilesGetById200Response) GetCanUseToolOk() (*bool, bool) {
	if o == nil || IsNil(o.CanUseTool) {
		return nil, false
	}
	return o.CanUseTool, true
}

// HasCanUseTool returns a boolean if a field has been set.
func (o *AiProfilesGetById200Response) IsCanUseToolSet() bool {
	if o != nil && !IsNil(o.CanUseTool) {
		return true
	}

	return false
}

// SetCanUseTool gets a reference to the given bool and assigns it to the CanUseTool field.
func (o *AiProfilesGetById200Response) SetCanUseTool(v bool) {
	o.CanUseTool = &v
}

// GetUseResponsesApi returns the UseResponsesApi field value if set, zero value otherwise.
func (o *AiProfilesGetById200Response) GetUseResponsesApi() bool {
	if o == nil || IsNil(o.UseResponsesApi) {
		var ret bool
		return ret
	}
	return *o.UseResponsesApi
}

// GetUseResponsesApiOk returns a tuple with the UseResponsesApi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiProfilesGetById200Response) GetUseResponsesApiOk() (*bool, bool) {
	if o == nil || IsNil(o.UseResponsesApi) {
		return nil, false
	}
	return o.UseResponsesApi, true
}

// HasUseResponsesApi returns a boolean if a field has been set.
func (o *AiProfilesGetById200Response) IsUseResponsesApiSet() bool {
	if o != nil && !IsNil(o.UseResponsesApi) {
		return true
	}

	return false
}

// SetUseResponsesApi gets a reference to the given bool and assigns it to the UseResponsesApi field.
func (o *AiProfilesGetById200Response) SetUseResponsesApi(v bool) {
	o.UseResponsesApi = &v
}

// GetIsCloudProvider returns the IsCloudProvider field value if set, zero value otherwise.
func (o *AiProfilesGetById200Response) GetIsCloudProvider() bool {
	if o == nil || IsNil(o.IsCloudProvider) {
		var ret bool
		return ret
	}
	return *o.IsCloudProvider
}

// GetIsCloudProviderOk returns a tuple with the IsCloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiProfilesGetById200Response) GetIsCloudProviderOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCloudProvider) {
		return nil, false
	}
	return o.IsCloudProvider, true
}

// HasIsCloudProvider returns a boolean if a field has been set.
func (o *AiProfilesGetById200Response) IsIsCloudProviderSet() bool {
	if o != nil && !IsNil(o.IsCloudProvider) {
		return true
	}

	return false
}

// SetIsCloudProvider gets a reference to the given bool and assigns it to the IsCloudProvider field.
func (o *AiProfilesGetById200Response) SetIsCloudProvider(v bool) {
	o.IsCloudProvider = &v
}

// GetUseProxy returns the UseProxy field value if set, zero value otherwise.
func (o *AiProfilesGetById200Response) GetUseProxy() bool {
	if o == nil || IsNil(o.UseProxy) {
		var ret bool
		return ret
	}
	return *o.UseProxy
}

// GetUseProxyOk returns a tuple with the UseProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiProfilesGetById200Response) GetUseProxyOk() (*bool, bool) {
	if o == nil || IsNil(o.UseProxy) {
		return nil, false
	}
	return o.UseProxy, true
}

// HasUseProxy returns a boolean if a field has been set.
func (o *AiProfilesGetById200Response) IsUseProxySet() bool {
	if o != nil && !IsNil(o.UseProxy) {
		return true
	}

	return false
}

// SetUseProxy gets a reference to the given bool and assigns it to the UseProxy field.
func (o *AiProfilesGetById200Response) SetUseProxy(v bool) {
	o.UseProxy = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AiProfilesGetById200Response) GetCreatedAt() float32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret float32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiProfilesGetById200Response) GetCreatedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AiProfilesGetById200Response) IsCreatedAtSet() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given float32 and assigns it to the CreatedAt field.
func (o *AiProfilesGetById200Response) SetCreatedAt(v float32) {
	o.CreatedAt = &v
}

func (o AiProfilesGetById200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiProfilesGetById200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["providerType"] = o.ProviderType
	if !IsNil(o.BasedOn) {
		toSerialize["basedOn"] = o.BasedOn
	}
	toSerialize["baseUrl"] = o.BaseUrl
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
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	return toSerialize, nil
}

func (o *AiProfilesGetById200Response) UnmarshalJSON(data []byte) (err error) {
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

	varAiProfilesGetById200Response := _AiProfilesGetById200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiProfilesGetById200Response)

	if err != nil {
		return err
	}

	*o = AiProfilesGetById200Response(varAiProfilesGetById200Response)

	return err
}

type NullableAiProfilesGetById200Response struct {
	value *AiProfilesGetById200Response
	isSet bool
}

func (v NullableAiProfilesGetById200Response) Get() *AiProfilesGetById200Response {
	return v.value
}

func (v *NullableAiProfilesGetById200Response) Set(val *AiProfilesGetById200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAiProfilesGetById200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAiProfilesGetById200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiProfilesGetById200Response(val *AiProfilesGetById200Response) *NullableAiProfilesGetById200Response {
	return &NullableAiProfilesGetById200Response{value: val, isSet: true}
}

func (v NullableAiProfilesGetById200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiProfilesGetById200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

