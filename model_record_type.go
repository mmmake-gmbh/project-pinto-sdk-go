/*
 * Pinto - OpenApi Gateway - DNS
 *
 * <h1>Additional info</h1> <h2><a href=\"/nexus-info\">Nexus Info (DNS - CloudProviders)</h2></a>  <h2><a href=\"/markdown/OPENAPIGATEWAY-USAGE.MD\">OPENAPIGATEWAY-USAGE.MD</h2></a> <h2><a href=\"/markdown/CHANGELOG.md\">CHANGELOG.md</h2></a> <h2><a href=\"/markdown/version.md\">version.md</h2></a>
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gopinto

import (
	"encoding/json"
	"fmt"
)

// RecordType Resource record types  as defined in  <see href=\"https://tools.ietf.org/html/rfc1035#section-3.2.2\">rfc1035</see>
type RecordType string

// List of RecordType
const (
	A     RecordType = "A"
	NS    RecordType = "NS"
	CNAME RecordType = "CNAME"
	SOA   RecordType = "SOA"
	PTR   RecordType = "PTR"
	MX    RecordType = "MX"
	TXT   RecordType = "TXT"
	SRV   RecordType = "SRV"
	AAAA  RecordType = "AAAA"
	SPF   RecordType = "SPF"
)

func (v *RecordType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RecordType(value)
	for _, existing := range []RecordType{"A", "NS", "CNAME", "SOA", "PTR", "MX", "TXT", "SRV", "AAAA", "SPF"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RecordType", value)
}

// Ptr returns reference to RecordType value
func (v RecordType) Ptr() *RecordType {
	return &v
}

type NullableRecordType struct {
	value *RecordType
	isSet bool
}

func (v NullableRecordType) Get() *RecordType {
	return v.value
}

func (v *NullableRecordType) Set(val *RecordType) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordType) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordType(val *RecordType) *NullableRecordType {
	return &NullableRecordType{value: val, isSet: true}
}

func (v NullableRecordType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
