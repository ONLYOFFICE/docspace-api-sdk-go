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
	"time"
	"bytes"
	"fmt"
)

// checks if the WebhooksLogDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhooksLogDto{}

// WebhooksLogDto The webhook log parameters.
type WebhooksLogDto struct {
	// The webhook log ID.
	Id int32 `json:"id"`
	// The webhook configuration name.
	ConfigName NullableString `json:"configName,omitempty"`
	Trigger *WebhookTrigger `json:"trigger,omitempty"`
	// The webhook creation time.
	CreationTime *time.Time `json:"creationTime,omitempty"`
	// The webhook method.
	Method NullableString `json:"method,omitempty"`
	// The webhook route.
	Route NullableString `json:"route,omitempty"`
	// The webhook request headers.
	RequestHeaders NullableString `json:"requestHeaders,omitempty"`
	// The webhook request payload.
	RequestPayload NullableString `json:"requestPayload,omitempty"`
	// The webhook response headers.
	ResponseHeaders NullableString `json:"responseHeaders,omitempty"`
	// The webhook response payload.
	ResponsePayload NullableString `json:"responsePayload,omitempty"`
	// The webhook status.
	Status *int32 `json:"status,omitempty"`
	// The webhook delivery time.
	Delivery NullableTime `json:"delivery,omitempty"`
}

type _WebhooksLogDto WebhooksLogDto

// NewWebhooksLogDto instantiates a new WebhooksLogDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhooksLogDto(id int32) *WebhooksLogDto {
	this := WebhooksLogDto{}
	this.Id = id
	return &this
}

// NewWebhooksLogDtoWithDefaults instantiates a new WebhooksLogDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhooksLogDtoWithDefaults() *WebhooksLogDto {
	this := WebhooksLogDto{}
	return &this
}

// GetId returns the Id field value
func (o *WebhooksLogDto) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WebhooksLogDto) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WebhooksLogDto) SetId(v int32) {
	o.Id = v
}

// GetConfigName returns the ConfigName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhooksLogDto) GetConfigName() string {
	if o == nil || IsNil(o.ConfigName.Get()) {
		var ret string
		return ret
	}
	return *o.ConfigName.Get()
}

// GetConfigNameOk returns a tuple with the ConfigName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhooksLogDto) GetConfigNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigName.Get(), o.ConfigName.IsSet()
}

// HasConfigName returns a boolean if a field has been set.
func (o *WebhooksLogDto) IsConfigNameSet() bool {
	if o != nil && o.ConfigName.IsSet() {
		return true
	}

	return false
}

// SetConfigName gets a reference to the given NullableString and assigns it to the ConfigName field.
func (o *WebhooksLogDto) SetConfigName(v string) {
	o.ConfigName.Set(&v)
}
// SetConfigNameNil sets the value for ConfigName to be an explicit nil
func (o *WebhooksLogDto) SetConfigNameNil() {
	o.ConfigName.Set(nil)
}

// UnsetConfigName ensures that no value is present for ConfigName, not even an explicit nil
func (o *WebhooksLogDto) UnsetConfigName() {
	o.ConfigName.Unset()
}

// GetTrigger returns the Trigger field value if set, zero value otherwise.
func (o *WebhooksLogDto) GetTrigger() WebhookTrigger {
	if o == nil || IsNil(o.Trigger) {
		var ret WebhookTrigger
		return ret
	}
	return *o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksLogDto) GetTriggerOk() (*WebhookTrigger, bool) {
	if o == nil || IsNil(o.Trigger) {
		return nil, false
	}
	return o.Trigger, true
}

// HasTrigger returns a boolean if a field has been set.
func (o *WebhooksLogDto) IsTriggerSet() bool {
	if o != nil && !IsNil(o.Trigger) {
		return true
	}

	return false
}

// SetTrigger gets a reference to the given WebhookTrigger and assigns it to the Trigger field.
func (o *WebhooksLogDto) SetTrigger(v WebhookTrigger) {
	o.Trigger = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *WebhooksLogDto) GetCreationTime() time.Time {
	if o == nil || IsNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksLogDto) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *WebhooksLogDto) IsCreationTimeSet() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *WebhooksLogDto) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetMethod returns the Method field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhooksLogDto) GetMethod() string {
	if o == nil || IsNil(o.Method.Get()) {
		var ret string
		return ret
	}
	return *o.Method.Get()
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhooksLogDto) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Method.Get(), o.Method.IsSet()
}

// HasMethod returns a boolean if a field has been set.
func (o *WebhooksLogDto) IsMethodSet() bool {
	if o != nil && o.Method.IsSet() {
		return true
	}

	return false
}

// SetMethod gets a reference to the given NullableString and assigns it to the Method field.
func (o *WebhooksLogDto) SetMethod(v string) {
	o.Method.Set(&v)
}
// SetMethodNil sets the value for Method to be an explicit nil
func (o *WebhooksLogDto) SetMethodNil() {
	o.Method.Set(nil)
}

// UnsetMethod ensures that no value is present for Method, not even an explicit nil
func (o *WebhooksLogDto) UnsetMethod() {
	o.Method.Unset()
}

// GetRoute returns the Route field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhooksLogDto) GetRoute() string {
	if o == nil || IsNil(o.Route.Get()) {
		var ret string
		return ret
	}
	return *o.Route.Get()
}

// GetRouteOk returns a tuple with the Route field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhooksLogDto) GetRouteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Route.Get(), o.Route.IsSet()
}

// HasRoute returns a boolean if a field has been set.
func (o *WebhooksLogDto) IsRouteSet() bool {
	if o != nil && o.Route.IsSet() {
		return true
	}

	return false
}

// SetRoute gets a reference to the given NullableString and assigns it to the Route field.
func (o *WebhooksLogDto) SetRoute(v string) {
	o.Route.Set(&v)
}
// SetRouteNil sets the value for Route to be an explicit nil
func (o *WebhooksLogDto) SetRouteNil() {
	o.Route.Set(nil)
}

// UnsetRoute ensures that no value is present for Route, not even an explicit nil
func (o *WebhooksLogDto) UnsetRoute() {
	o.Route.Unset()
}

// GetRequestHeaders returns the RequestHeaders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhooksLogDto) GetRequestHeaders() string {
	if o == nil || IsNil(o.RequestHeaders.Get()) {
		var ret string
		return ret
	}
	return *o.RequestHeaders.Get()
}

// GetRequestHeadersOk returns a tuple with the RequestHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhooksLogDto) GetRequestHeadersOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestHeaders.Get(), o.RequestHeaders.IsSet()
}

// HasRequestHeaders returns a boolean if a field has been set.
func (o *WebhooksLogDto) IsRequestHeadersSet() bool {
	if o != nil && o.RequestHeaders.IsSet() {
		return true
	}

	return false
}

// SetRequestHeaders gets a reference to the given NullableString and assigns it to the RequestHeaders field.
func (o *WebhooksLogDto) SetRequestHeaders(v string) {
	o.RequestHeaders.Set(&v)
}
// SetRequestHeadersNil sets the value for RequestHeaders to be an explicit nil
func (o *WebhooksLogDto) SetRequestHeadersNil() {
	o.RequestHeaders.Set(nil)
}

// UnsetRequestHeaders ensures that no value is present for RequestHeaders, not even an explicit nil
func (o *WebhooksLogDto) UnsetRequestHeaders() {
	o.RequestHeaders.Unset()
}

// GetRequestPayload returns the RequestPayload field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhooksLogDto) GetRequestPayload() string {
	if o == nil || IsNil(o.RequestPayload.Get()) {
		var ret string
		return ret
	}
	return *o.RequestPayload.Get()
}

// GetRequestPayloadOk returns a tuple with the RequestPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhooksLogDto) GetRequestPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestPayload.Get(), o.RequestPayload.IsSet()
}

// HasRequestPayload returns a boolean if a field has been set.
func (o *WebhooksLogDto) IsRequestPayloadSet() bool {
	if o != nil && o.RequestPayload.IsSet() {
		return true
	}

	return false
}

// SetRequestPayload gets a reference to the given NullableString and assigns it to the RequestPayload field.
func (o *WebhooksLogDto) SetRequestPayload(v string) {
	o.RequestPayload.Set(&v)
}
// SetRequestPayloadNil sets the value for RequestPayload to be an explicit nil
func (o *WebhooksLogDto) SetRequestPayloadNil() {
	o.RequestPayload.Set(nil)
}

// UnsetRequestPayload ensures that no value is present for RequestPayload, not even an explicit nil
func (o *WebhooksLogDto) UnsetRequestPayload() {
	o.RequestPayload.Unset()
}

// GetResponseHeaders returns the ResponseHeaders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhooksLogDto) GetResponseHeaders() string {
	if o == nil || IsNil(o.ResponseHeaders.Get()) {
		var ret string
		return ret
	}
	return *o.ResponseHeaders.Get()
}

// GetResponseHeadersOk returns a tuple with the ResponseHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhooksLogDto) GetResponseHeadersOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponseHeaders.Get(), o.ResponseHeaders.IsSet()
}

// HasResponseHeaders returns a boolean if a field has been set.
func (o *WebhooksLogDto) IsResponseHeadersSet() bool {
	if o != nil && o.ResponseHeaders.IsSet() {
		return true
	}

	return false
}

// SetResponseHeaders gets a reference to the given NullableString and assigns it to the ResponseHeaders field.
func (o *WebhooksLogDto) SetResponseHeaders(v string) {
	o.ResponseHeaders.Set(&v)
}
// SetResponseHeadersNil sets the value for ResponseHeaders to be an explicit nil
func (o *WebhooksLogDto) SetResponseHeadersNil() {
	o.ResponseHeaders.Set(nil)
}

// UnsetResponseHeaders ensures that no value is present for ResponseHeaders, not even an explicit nil
func (o *WebhooksLogDto) UnsetResponseHeaders() {
	o.ResponseHeaders.Unset()
}

// GetResponsePayload returns the ResponsePayload field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhooksLogDto) GetResponsePayload() string {
	if o == nil || IsNil(o.ResponsePayload.Get()) {
		var ret string
		return ret
	}
	return *o.ResponsePayload.Get()
}

// GetResponsePayloadOk returns a tuple with the ResponsePayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhooksLogDto) GetResponsePayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponsePayload.Get(), o.ResponsePayload.IsSet()
}

// HasResponsePayload returns a boolean if a field has been set.
func (o *WebhooksLogDto) IsResponsePayloadSet() bool {
	if o != nil && o.ResponsePayload.IsSet() {
		return true
	}

	return false
}

// SetResponsePayload gets a reference to the given NullableString and assigns it to the ResponsePayload field.
func (o *WebhooksLogDto) SetResponsePayload(v string) {
	o.ResponsePayload.Set(&v)
}
// SetResponsePayloadNil sets the value for ResponsePayload to be an explicit nil
func (o *WebhooksLogDto) SetResponsePayloadNil() {
	o.ResponsePayload.Set(nil)
}

// UnsetResponsePayload ensures that no value is present for ResponsePayload, not even an explicit nil
func (o *WebhooksLogDto) UnsetResponsePayload() {
	o.ResponsePayload.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WebhooksLogDto) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksLogDto) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WebhooksLogDto) IsStatusSet() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *WebhooksLogDto) SetStatus(v int32) {
	o.Status = &v
}

// GetDelivery returns the Delivery field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhooksLogDto) GetDelivery() time.Time {
	if o == nil || IsNil(o.Delivery.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Delivery.Get()
}

// GetDeliveryOk returns a tuple with the Delivery field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhooksLogDto) GetDeliveryOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Delivery.Get(), o.Delivery.IsSet()
}

// HasDelivery returns a boolean if a field has been set.
func (o *WebhooksLogDto) IsDeliverySet() bool {
	if o != nil && o.Delivery.IsSet() {
		return true
	}

	return false
}

// SetDelivery gets a reference to the given NullableTime and assigns it to the Delivery field.
func (o *WebhooksLogDto) SetDelivery(v time.Time) {
	o.Delivery.Set(&v)
}
// SetDeliveryNil sets the value for Delivery to be an explicit nil
func (o *WebhooksLogDto) SetDeliveryNil() {
	o.Delivery.Set(nil)
}

// UnsetDelivery ensures that no value is present for Delivery, not even an explicit nil
func (o *WebhooksLogDto) UnsetDelivery() {
	o.Delivery.Unset()
}

func (o WebhooksLogDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhooksLogDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.ConfigName.IsSet() {
		toSerialize["configName"] = o.ConfigName.Get()
	}
	if !IsNil(o.Trigger) {
		toSerialize["trigger"] = o.Trigger
	}
	if !IsNil(o.CreationTime) {
		toSerialize["creationTime"] = o.CreationTime
	}
	if o.Method.IsSet() {
		toSerialize["method"] = o.Method.Get()
	}
	if o.Route.IsSet() {
		toSerialize["route"] = o.Route.Get()
	}
	if o.RequestHeaders.IsSet() {
		toSerialize["requestHeaders"] = o.RequestHeaders.Get()
	}
	if o.RequestPayload.IsSet() {
		toSerialize["requestPayload"] = o.RequestPayload.Get()
	}
	if o.ResponseHeaders.IsSet() {
		toSerialize["responseHeaders"] = o.ResponseHeaders.Get()
	}
	if o.ResponsePayload.IsSet() {
		toSerialize["responsePayload"] = o.ResponsePayload.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Delivery.IsSet() {
		toSerialize["delivery"] = o.Delivery.Get()
	}
	return toSerialize, nil
}

func (o *WebhooksLogDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varWebhooksLogDto := _WebhooksLogDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varWebhooksLogDto)

	if err != nil {
		return err
	}

	*o = WebhooksLogDto(varWebhooksLogDto)

	return err
}

type NullableWebhooksLogDto struct {
	value *WebhooksLogDto
	isSet bool
}

func (v NullableWebhooksLogDto) Get() *WebhooksLogDto {
	return v.value
}

func (v *NullableWebhooksLogDto) Set(val *WebhooksLogDto) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhooksLogDto) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhooksLogDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhooksLogDto(val *WebhooksLogDto) *NullableWebhooksLogDto {
	return &NullableWebhooksLogDto{value: val, isSet: true}
}

func (v NullableWebhooksLogDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhooksLogDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

