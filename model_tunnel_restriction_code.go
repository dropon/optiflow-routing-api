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

// TunnelRestrictionCode The tunnel restriction code according to ADR (European Agreement Concerning the International Carriage of Dangerous Goods by Road) depending on the load of the vehicle.  Relevant for `routing`.
type TunnelRestrictionCode string

// List of TunnelRestrictionCode
const (
	TUNNEL_RESTRICTION_NONE TunnelRestrictionCode = "NONE"
	B                       TunnelRestrictionCode = "B"
	C                       TunnelRestrictionCode = "C"
	D                       TunnelRestrictionCode = "D"
	E                       TunnelRestrictionCode = "E"
)

// All allowed values of TunnelRestrictionCode enum
var AllowedTunnelRestrictionCodeEnumValues = []TunnelRestrictionCode{
	"NONE",
	"B",
	"C",
	"D",
	"E",
}

func (v *TunnelRestrictionCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TunnelRestrictionCode(value)
	for _, existing := range AllowedTunnelRestrictionCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TunnelRestrictionCode", value)
}

// NewTunnelRestrictionCodeFromValue returns a pointer to a valid TunnelRestrictionCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTunnelRestrictionCodeFromValue(v string) (*TunnelRestrictionCode, error) {
	ev := TunnelRestrictionCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TunnelRestrictionCode: valid values are %v", v, AllowedTunnelRestrictionCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TunnelRestrictionCode) IsValid() bool {
	for _, existing := range AllowedTunnelRestrictionCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TunnelRestrictionCode value
func (v TunnelRestrictionCode) Ptr() *TunnelRestrictionCode {
	return &v
}

type NullableTunnelRestrictionCode struct {
	value *TunnelRestrictionCode
	isSet bool
}

func (v NullableTunnelRestrictionCode) Get() *TunnelRestrictionCode {
	return v.value
}

func (v *NullableTunnelRestrictionCode) Set(val *TunnelRestrictionCode) {
	v.value = val
	v.isSet = true
}

func (v NullableTunnelRestrictionCode) IsSet() bool {
	return v.isSet
}

func (v *NullableTunnelRestrictionCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTunnelRestrictionCode(val *TunnelRestrictionCode) *NullableTunnelRestrictionCode {
	return &NullableTunnelRestrictionCode{value: val, isSet: true}
}

func (v NullableTunnelRestrictionCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTunnelRestrictionCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
