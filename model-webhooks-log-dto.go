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

// WebhooksLogDto One delivery of a webhook event, updated in place by its automatic retries: what was sent where, and what came  back the last time. A manual retry is a delivery record of its own for the same event.
type WebhooksLogDto struct {
	// The identifier of this delivery record, which is what `PUT api/2.0/settings/webhook/{id}/retry` re-sends.  The automatic retries of the delivery update this same record, while a manual retry gets a record and an  identifier of its own.
	Id int32 `json:"id"`
	// The identifier of the event the delivery is for, which the receiver sees as `event.id` and in the  `x-docspace-event-id` header. It is the `id` of the event's first delivery record, so it equals `id` there and  stays the same on every manual retry, which makes it the value to group the records of one event by.
	EventId *int32 `json:"eventId,omitempty"`
	// The member who asked for this delivery with a manual retry. It is empty on the first delivery of an event,  which the portal sends by itself.
	RetriedBy *EmployeeDto `json:"retriedBy,omitempty"`
	// The name of the subscription the delivery belongs to. It is the name as it stands now, so it follows a  later rename of the subscription rather than recording what it was called at the time.
	ConfigName NullableString `json:"configName,omitempty"`
	// The event that caused the delivery, as a single bit rather than a mask - a delivery is always for one  event, even though a subscription covers several.
	Trigger *WebhookTrigger `json:"trigger,omitempty"`
	// When this delivery record was queued, as a UTC instant ending in `Z`. For the first delivery that is when the  event happened; for a manual retry it is the  moment of the retry. Records come back newest first by this moment.
	CreationTime *time.Time `json:"creationTime,omitempty"`
	// The HTTP method the delivery was sent with, which is `POST` for every webhook the portal sends.
	Method NullableString `json:"method,omitempty"`
	// The address the delivery is sent to: the subscription's URL as it stands now. The journal does not keep the  address of each attempt, so after the subscription's URL changes, an older record names the new one too.
	Route NullableString `json:"route,omitempty"`
	// The headers the portal sent, serialised as one string, including the signature header a receiver verifies  the payload with.
	RequestHeaders NullableString `json:"requestHeaders,omitempty"`
	// The body the portal sent, which is the event payload as JSON text. It is stored as it was sent, so it  still describes the entity as it looked at the time of the event.
	RequestPayload NullableString `json:"requestPayload,omitempty"`
	// The headers the target answered with on the latest attempt, serialised the same way as `requestHeaders`. It  is empty until an attempt is answered, and when the target was not reached.
	ResponseHeaders NullableString `json:"responseHeaders,omitempty"`
	// The body the target answered with on the latest attempt, truncated for storage. Empty under the same  conditions as `responseHeaders`, and also for a target that answers with no body at all. A record the portal  stopped without sending holds the portal's own note here instead - a text that starts with `Not sent:` and  gives the reason - which is not an answer of the target: on a record that never went out `status` is then `0`  and `responseHeaders` is empty, while on one stopped after failed attempts they still describe the last of them.
	ResponsePayload NullableString `json:"responsePayload,omitempty"`
	// The HTTP status code the target answered on the latest attempt. It is `0` while the first attempt is still  on its way and on one that never reached the target, so `0` is not a failure code - it is the absence of an  answer. While `nextAttemptOn` is set it is the status of the failed attempt that caused the retry.
	Status *int32 `json:"status,omitempty"`
	// When the successful answer came back, as a UTC instant ending in `Z` like `creationTime`. It stays empty until an  attempt succeeds, and for good once the portal has given up retrying.
	Delivery NullableTime `json:"delivery,omitempty"`
	// How many times the portal has tried to deliver this record so far. A failed attempt is retried with  growing pauses - on the default schedule about 1 minute, 5 minutes, 30 minutes, 2, 5 and 10 hours, though  an installation may configure its own - all of them under the same record and the same  `x-docspace-event-id` header, so a receiver can drop a delivery it has already processed. The one exception  is a portal moved to another region while retries are pending: the moved records get new IDs, so the attempts  made there carry a new `x-docspace-event-id`.
	Attempts *int32 `json:"attempts,omitempty"`
	// When the next attempt is due, as a UTC instant ending in `Z` like `creationTime`. While an attempt is being sent  right now it holds the moment that attempt's lease expires, a few minutes ahead, rather than a scheduled retry;  it changes again once the attempt ends. Empty means no attempt will follow: the record is delivered, its retries  ran out, a manual retry of the same event took over from it, or the portal stopped it - because the subscription  was disabled or removed or its address rejected, the subscription's creator may no longer receive the event,  the stored body cannot be read or is in the payload format of a previous version, the portal is removed, the  record waited longer than the retry schedule allows, the portal was restored from a backup taken before the  record was answered, or the portal was moved to another region while the record was being sent. A stopped  record that was never attempted says why in `responsePayload`; one that was keeps the target's last answer there.
	NextAttemptOn NullableTime `json:"nextAttemptOn,omitempty"`
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

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *WebhooksLogDto) GetEventId() int32 {
	if o == nil || IsNil(o.EventId) {
		var ret int32
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksLogDto) GetEventIdOk() (*int32, bool) {
	if o == nil || IsNil(o.EventId) {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *WebhooksLogDto) IsEventIdSet() bool {
	if o != nil && !IsNil(o.EventId) {
		return true
	}

	return false
}

// SetEventId gets a reference to the given int32 and assigns it to the EventId field.
func (o *WebhooksLogDto) SetEventId(v int32) {
	o.EventId = &v
}

// GetRetriedBy returns the RetriedBy field value if set, zero value otherwise.
func (o *WebhooksLogDto) GetRetriedBy() EmployeeDto {
	if o == nil || IsNil(o.RetriedBy) {
		var ret EmployeeDto
		return ret
	}
	return *o.RetriedBy
}

// GetRetriedByOk returns a tuple with the RetriedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksLogDto) GetRetriedByOk() (*EmployeeDto, bool) {
	if o == nil || IsNil(o.RetriedBy) {
		return nil, false
	}
	return o.RetriedBy, true
}

// HasRetriedBy returns a boolean if a field has been set.
func (o *WebhooksLogDto) IsRetriedBySet() bool {
	if o != nil && !IsNil(o.RetriedBy) {
		return true
	}

	return false
}

// SetRetriedBy gets a reference to the given EmployeeDto and assigns it to the RetriedBy field.
func (o *WebhooksLogDto) SetRetriedBy(v EmployeeDto) {
	o.RetriedBy = &v
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

// GetAttempts returns the Attempts field value if set, zero value otherwise.
func (o *WebhooksLogDto) GetAttempts() int32 {
	if o == nil || IsNil(o.Attempts) {
		var ret int32
		return ret
	}
	return *o.Attempts
}

// GetAttemptsOk returns a tuple with the Attempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksLogDto) GetAttemptsOk() (*int32, bool) {
	if o == nil || IsNil(o.Attempts) {
		return nil, false
	}
	return o.Attempts, true
}

// HasAttempts returns a boolean if a field has been set.
func (o *WebhooksLogDto) IsAttemptsSet() bool {
	if o != nil && !IsNil(o.Attempts) {
		return true
	}

	return false
}

// SetAttempts gets a reference to the given int32 and assigns it to the Attempts field.
func (o *WebhooksLogDto) SetAttempts(v int32) {
	o.Attempts = &v
}

// GetNextAttemptOn returns the NextAttemptOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhooksLogDto) GetNextAttemptOn() time.Time {
	if o == nil || IsNil(o.NextAttemptOn.Get()) {
		var ret time.Time
		return ret
	}
	return *o.NextAttemptOn.Get()
}

// GetNextAttemptOnOk returns a tuple with the NextAttemptOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhooksLogDto) GetNextAttemptOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextAttemptOn.Get(), o.NextAttemptOn.IsSet()
}

// HasNextAttemptOn returns a boolean if a field has been set.
func (o *WebhooksLogDto) IsNextAttemptOnSet() bool {
	if o != nil && o.NextAttemptOn.IsSet() {
		return true
	}

	return false
}

// SetNextAttemptOn gets a reference to the given NullableTime and assigns it to the NextAttemptOn field.
func (o *WebhooksLogDto) SetNextAttemptOn(v time.Time) {
	o.NextAttemptOn.Set(&v)
}
// SetNextAttemptOnNil sets the value for NextAttemptOn to be an explicit nil
func (o *WebhooksLogDto) SetNextAttemptOnNil() {
	o.NextAttemptOn.Set(nil)
}

// UnsetNextAttemptOn ensures that no value is present for NextAttemptOn, not even an explicit nil
func (o *WebhooksLogDto) UnsetNextAttemptOn() {
	o.NextAttemptOn.Unset()
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
	if !IsNil(o.EventId) {
		toSerialize["eventId"] = o.EventId
	}
	if !IsNil(o.RetriedBy) {
		toSerialize["retriedBy"] = o.RetriedBy
	}
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
	if !IsNil(o.Attempts) {
		toSerialize["attempts"] = o.Attempts
	}
	if o.NextAttemptOn.IsSet() {
		toSerialize["nextAttemptOn"] = o.NextAttemptOn.Get()
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

