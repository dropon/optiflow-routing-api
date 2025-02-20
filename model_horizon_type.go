/*
Routing

With the Routing service you can calculate routes from A to B taking into account vehicle-specific restrictions, traffic situations, toll, emissions, drivers' working hours, service times and opening intervals.

API version: 1.33
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package optiflow_routing

import (
	"encoding/json"
	"fmt"
)

// HorizonType Specifies the geographical horizon.  * `DISTANCE` - Represents a geographical horizon that is described by a distance. Every point or road segment that is reachable from the source within the specified distance is included in the horizon.  * `TRAVEL_TIME` - Represents a geographical horizon that is described by a travel time. Every point or road segment that is reachable from the source within the specified travel time is included in the horizon.
type HorizonType string

// List of HorizonType
const (
	DISTANCE HorizonType = "DISTANCE"
	TRAVEL_TIME HorizonType = "TRAVEL_TIME"
)

// All allowed values of HorizonType enum
var AllowedHorizonTypeEnumValues = []HorizonType{
	"DISTANCE",
	"TRAVEL_TIME",
}

func (v *HorizonType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HorizonType(value)
	for _, existing := range AllowedHorizonTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HorizonType", value)
}

// NewHorizonTypeFromValue returns a pointer to a valid HorizonType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHorizonTypeFromValue(v string) (*HorizonType, error) {
	ev := HorizonType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HorizonType: valid values are %v", v, AllowedHorizonTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HorizonType) IsValid() bool {
	for _, existing := range AllowedHorizonTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HorizonType value
func (v HorizonType) Ptr() *HorizonType {
	return &v
}

type NullableHorizonType struct {
	value *HorizonType
	isSet bool
}

func (v NullableHorizonType) Get() *HorizonType {
	return v.value
}

func (v *NullableHorizonType) Set(val *HorizonType) {
	v.value = val
	v.isSet = true
}

func (v NullableHorizonType) IsSet() bool {
	return v.isSet
}

func (v *NullableHorizonType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHorizonType(val *HorizonType) *NullableHorizonType {
	return &NullableHorizonType{value: val, isSet: true}
}

func (v NullableHorizonType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHorizonType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

