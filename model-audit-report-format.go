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
	"fmt"
)

// AuditReportFormat []
type AuditReportFormat int32

// List of AuditReportFormat
const (
	AUDITREPORTFORMAT_Xlsx AuditReportFormat = 0
	AUDITREPORTFORMAT_Csv AuditReportFormat = 1
)

// All allowed values of AuditReportFormat enum
var AllowedAuditReportFormatEnumValues = []AuditReportFormat{
	0,
	1,
}

func (v *AuditReportFormat) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AuditReportFormat(value)
	for _, existing := range AllowedAuditReportFormatEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AuditReportFormat", value)
}

// NewAuditReportFormatFromValue returns a pointer to a valid AuditReportFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAuditReportFormatFromValue(v int32) (*AuditReportFormat, error) {
	ev := AuditReportFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AuditReportFormat: valid values are %v", v, AllowedAuditReportFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AuditReportFormat) IsValid() bool {
	for _, existing := range AllowedAuditReportFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AuditReportFormat value
func (v AuditReportFormat) Ptr() *AuditReportFormat {
	return &v
}

type NullableAuditReportFormat struct {
	value *AuditReportFormat
	isSet bool
}

func (v NullableAuditReportFormat) Get() *AuditReportFormat {
	return v.value
}

func (v *NullableAuditReportFormat) Set(val *AuditReportFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditReportFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditReportFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditReportFormat(val *AuditReportFormat) *NullableAuditReportFormat {
	return &NullableAuditReportFormat{value: val, isSet: true}
}

func (v NullableAuditReportFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditReportFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

