// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// SnippetResult - struct for SnippetResult
type SnippetResult struct {
	SnippetResultOption        *SnippetResultOption
	ArrayOfSnippetResultOption *[]SnippetResultOption
}

// SnippetResultOptionAsSnippetResult is a convenience function that returns SnippetResultOption wrapped in SnippetResult
func SnippetResultOptionAsSnippetResult(v *SnippetResultOption) SnippetResult {
	return SnippetResult{
		SnippetResultOption: v,
	}
}

// []SnippetResultOptionAsSnippetResult is a convenience function that returns []SnippetResultOption wrapped in SnippetResult
func ArrayOfSnippetResultOptionAsSnippetResult(v *[]SnippetResultOption) SnippetResult {
	return SnippetResult{
		ArrayOfSnippetResultOption: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SnippetResult) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal data into SnippetResultOption
	err = newStrictDecoder(data).Decode(&dst.SnippetResultOption)
	if err == nil && validateStruct(dst.SnippetResultOption) == nil {
		jsonSnippetResultOption, _ := json.Marshal(dst.SnippetResultOption)
		if string(jsonSnippetResultOption) == "{}" { // empty struct
			dst.SnippetResultOption = nil
		} else {
			return nil
		}
	} else {
		dst.SnippetResultOption = nil
	}

	// try to unmarshal data into ArrayOfSnippetResultOption
	err = newStrictDecoder(data).Decode(&dst.ArrayOfSnippetResultOption)
	if err == nil && validateStruct(dst.ArrayOfSnippetResultOption) == nil {
		jsonArrayOfSnippetResultOption, _ := json.Marshal(dst.ArrayOfSnippetResultOption)
		if string(jsonArrayOfSnippetResultOption) == "{}" { // empty struct
			dst.ArrayOfSnippetResultOption = nil
		} else {
			return nil
		}
	} else {
		dst.ArrayOfSnippetResultOption = nil
	}

	return fmt.Errorf("Data failed to match schemas in oneOf(SnippetResult)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SnippetResult) MarshalJSON() ([]byte, error) {
	if src.SnippetResultOption != nil {
		return json.Marshal(&src.SnippetResultOption)
	}

	if src.ArrayOfSnippetResultOption != nil {
		return json.Marshal(&src.ArrayOfSnippetResultOption)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SnippetResult) GetActualInstance() any {
	if obj == nil {
		return nil
	}
	if obj.SnippetResultOption != nil {
		return obj.SnippetResultOption
	}

	if obj.ArrayOfSnippetResultOption != nil {
		return obj.ArrayOfSnippetResultOption
	}

	// all schemas are nil
	return nil
}

type NullableSnippetResult struct {
	value *SnippetResult
	isSet bool
}

func (v NullableSnippetResult) Get() *SnippetResult {
	return v.value
}

func (v *NullableSnippetResult) Set(val *SnippetResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSnippetResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSnippetResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnippetResult(val *SnippetResult) *NullableSnippetResult {
	return &NullableSnippetResult{value: val, isSet: true}
}

func (v NullableSnippetResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnippetResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
