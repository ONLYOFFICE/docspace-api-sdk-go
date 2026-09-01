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

// checks if the FirebaseDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirebaseDto{}

// FirebaseDto The Firebase parameters.
type FirebaseDto struct {
	// The Firebase API key.
	ApiKey NullableString `json:"apiKey"`
	// The Firebase authentication domain.
	AuthDomain NullableString `json:"authDomain"`
	// The Firebase project ID.
	ProjectId NullableString `json:"projectId"`
	// The Firebase storage bucket.
	StorageBucket NullableString `json:"storageBucket"`
	// The Firebase messaging sender ID.
	MessagingSenderId NullableString `json:"messagingSenderId"`
	// The Firebase application ID.
	AppId NullableString `json:"appId"`
	// The Firebase measurement ID.
	MeasurementId NullableString `json:"measurementId"`
	// The Firebase database URL.
	DatabaseURL NullableString `json:"databaseURL"`
}

type _FirebaseDto FirebaseDto

// NewFirebaseDto instantiates a new FirebaseDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirebaseDto(apiKey NullableString, authDomain NullableString, projectId NullableString, storageBucket NullableString, messagingSenderId NullableString, appId NullableString, measurementId NullableString, databaseURL NullableString) *FirebaseDto {
	this := FirebaseDto{}
	this.ApiKey = apiKey
	this.AuthDomain = authDomain
	this.ProjectId = projectId
	this.StorageBucket = storageBucket
	this.MessagingSenderId = messagingSenderId
	this.AppId = appId
	this.MeasurementId = measurementId
	this.DatabaseURL = databaseURL
	return &this
}

// NewFirebaseDtoWithDefaults instantiates a new FirebaseDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirebaseDtoWithDefaults() *FirebaseDto {
	this := FirebaseDto{}
	return &this
}

// GetApiKey returns the ApiKey field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FirebaseDto) GetApiKey() string {
	if o == nil || o.ApiKey.Get() == nil {
		var ret string
		return ret
	}

	return *o.ApiKey.Get()
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirebaseDto) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiKey.Get(), o.ApiKey.IsSet()
}

// SetApiKey sets field value
func (o *FirebaseDto) SetApiKey(v string) {
	o.ApiKey.Set(&v)
}

// GetAuthDomain returns the AuthDomain field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FirebaseDto) GetAuthDomain() string {
	if o == nil || o.AuthDomain.Get() == nil {
		var ret string
		return ret
	}

	return *o.AuthDomain.Get()
}

// GetAuthDomainOk returns a tuple with the AuthDomain field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirebaseDto) GetAuthDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthDomain.Get(), o.AuthDomain.IsSet()
}

// SetAuthDomain sets field value
func (o *FirebaseDto) SetAuthDomain(v string) {
	o.AuthDomain.Set(&v)
}

// GetProjectId returns the ProjectId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FirebaseDto) GetProjectId() string {
	if o == nil || o.ProjectId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ProjectId.Get()
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirebaseDto) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectId.Get(), o.ProjectId.IsSet()
}

// SetProjectId sets field value
func (o *FirebaseDto) SetProjectId(v string) {
	o.ProjectId.Set(&v)
}

// GetStorageBucket returns the StorageBucket field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FirebaseDto) GetStorageBucket() string {
	if o == nil || o.StorageBucket.Get() == nil {
		var ret string
		return ret
	}

	return *o.StorageBucket.Get()
}

// GetStorageBucketOk returns a tuple with the StorageBucket field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirebaseDto) GetStorageBucketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageBucket.Get(), o.StorageBucket.IsSet()
}

// SetStorageBucket sets field value
func (o *FirebaseDto) SetStorageBucket(v string) {
	o.StorageBucket.Set(&v)
}

// GetMessagingSenderId returns the MessagingSenderId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FirebaseDto) GetMessagingSenderId() string {
	if o == nil || o.MessagingSenderId.Get() == nil {
		var ret string
		return ret
	}

	return *o.MessagingSenderId.Get()
}

// GetMessagingSenderIdOk returns a tuple with the MessagingSenderId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirebaseDto) GetMessagingSenderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MessagingSenderId.Get(), o.MessagingSenderId.IsSet()
}

// SetMessagingSenderId sets field value
func (o *FirebaseDto) SetMessagingSenderId(v string) {
	o.MessagingSenderId.Set(&v)
}

// GetAppId returns the AppId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FirebaseDto) GetAppId() string {
	if o == nil || o.AppId.Get() == nil {
		var ret string
		return ret
	}

	return *o.AppId.Get()
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirebaseDto) GetAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppId.Get(), o.AppId.IsSet()
}

// SetAppId sets field value
func (o *FirebaseDto) SetAppId(v string) {
	o.AppId.Set(&v)
}

// GetMeasurementId returns the MeasurementId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FirebaseDto) GetMeasurementId() string {
	if o == nil || o.MeasurementId.Get() == nil {
		var ret string
		return ret
	}

	return *o.MeasurementId.Get()
}

// GetMeasurementIdOk returns a tuple with the MeasurementId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirebaseDto) GetMeasurementIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MeasurementId.Get(), o.MeasurementId.IsSet()
}

// SetMeasurementId sets field value
func (o *FirebaseDto) SetMeasurementId(v string) {
	o.MeasurementId.Set(&v)
}

// GetDatabaseURL returns the DatabaseURL field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FirebaseDto) GetDatabaseURL() string {
	if o == nil || o.DatabaseURL.Get() == nil {
		var ret string
		return ret
	}

	return *o.DatabaseURL.Get()
}

// GetDatabaseURLOk returns a tuple with the DatabaseURL field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirebaseDto) GetDatabaseURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatabaseURL.Get(), o.DatabaseURL.IsSet()
}

// SetDatabaseURL sets field value
func (o *FirebaseDto) SetDatabaseURL(v string) {
	o.DatabaseURL.Set(&v)
}

func (o FirebaseDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirebaseDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiKey"] = o.ApiKey.Get()
	toSerialize["authDomain"] = o.AuthDomain.Get()
	toSerialize["projectId"] = o.ProjectId.Get()
	toSerialize["storageBucket"] = o.StorageBucket.Get()
	toSerialize["messagingSenderId"] = o.MessagingSenderId.Get()
	toSerialize["appId"] = o.AppId.Get()
	toSerialize["measurementId"] = o.MeasurementId.Get()
	toSerialize["databaseURL"] = o.DatabaseURL.Get()
	return toSerialize, nil
}

func (o *FirebaseDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"apiKey",
		"authDomain",
		"projectId",
		"storageBucket",
		"messagingSenderId",
		"appId",
		"measurementId",
		"databaseURL",
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

	varFirebaseDto := _FirebaseDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varFirebaseDto)

	if err != nil {
		return err
	}

	*o = FirebaseDto(varFirebaseDto)

	return err
}

type NullableFirebaseDto struct {
	value *FirebaseDto
	isSet bool
}

func (v NullableFirebaseDto) Get() *FirebaseDto {
	return v.value
}

func (v *NullableFirebaseDto) Set(val *FirebaseDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFirebaseDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFirebaseDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirebaseDto(val *FirebaseDto) *NullableFirebaseDto {
	return &NullableFirebaseDto{value: val, isSet: true}
}

func (v NullableFirebaseDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirebaseDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

