/*
 * Power API v1.0
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Contact: lowell@lanl.gov
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package powerapi

import (
	"encoding/json"
	"fmt"
)

// PowerState Power state for a component
type PowerState string

// List of PowerState
const (
	POWERSTATE_ON PowerState = "On"
	POWERSTATE_OFF PowerState = "Off"
)

func (v *PowerState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PowerState(value)
	for _, existing := range []PowerState{ "On", "Off",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PowerState", value)
}

// Ptr returns reference to PowerState value
func (v PowerState) Ptr() *PowerState {
	return &v
}

type NullablePowerState struct {
	value *PowerState
	isSet bool
}

func (v NullablePowerState) Get() *PowerState {
	return v.value
}

func (v *NullablePowerState) Set(val *PowerState) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerState) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerState(val *PowerState) *NullablePowerState {
	return &NullablePowerState{value: val, isSet: true}
}

func (v NullablePowerState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

