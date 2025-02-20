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

// PolylineMapType Defines whether polylines should match the Raster Maps API or the Vector Maps API.  Not only the polyline itself, but also other results and parameters that use the polyline of the route may change slightly, e.g. emission related results and **matchSideOfStreet**.
type PolylineMapType string

// List of PolylineMapType
const (
	RASTER PolylineMapType = "RASTER"
	VECTOR PolylineMapType = "VECTOR"
)

// All allowed values of PolylineMapType enum
var AllowedPolylineMapTypeEnumValues = []PolylineMapType{
	"RASTER",
	"VECTOR",
}

func (v *PolylineMapType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PolylineMapType(value)
	for _, existing := range AllowedPolylineMapTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PolylineMapType", value)
}

// NewPolylineMapTypeFromValue returns a pointer to a valid PolylineMapType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPolylineMapTypeFromValue(v string) (*PolylineMapType, error) {
	ev := PolylineMapType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PolylineMapType: valid values are %v", v, AllowedPolylineMapTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PolylineMapType) IsValid() bool {
	for _, existing := range AllowedPolylineMapTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PolylineMapType value
func (v PolylineMapType) Ptr() *PolylineMapType {
	return &v
}

type NullablePolylineMapType struct {
	value *PolylineMapType
	isSet bool
}

func (v NullablePolylineMapType) Get() *PolylineMapType {
	return v.value
}

func (v *NullablePolylineMapType) Set(val *PolylineMapType) {
	v.value = val
	v.isSet = true
}

func (v NullablePolylineMapType) IsSet() bool {
	return v.isSet
}

func (v *NullablePolylineMapType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolylineMapType(val *PolylineMapType) *NullablePolylineMapType {
	return &NullablePolylineMapType{value: val, isSet: true}
}

func (v NullablePolylineMapType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolylineMapType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

