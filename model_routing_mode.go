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

// RoutingMode Specifies which objective function should be used for the route calculation.   * _FAST_ is the default which returns a route considering a tradeoff between distance and travel time. All vehicle-specific restrictions are properly considered and violated only, if no other route can be found. * _SHORT_ returns a route which is probably shorter but accepting longer travel times. The resulting route is still sensible and can be driven with the given vehicle, but it may disregard restrictions like driving through residential areas. * _MONETARY_ assigns monetary costs to each road segment according the provided cost parameters and the vehicles properties like its consumption. Furthermore, toll costs are integrated as well. See [here](./concepts/monetary-costs) for more information.
type RoutingMode string

// List of RoutingMode
const (
	FAST RoutingMode = "FAST"
	SHORT RoutingMode = "SHORT"
	MONETARY RoutingMode = "MONETARY"
)

// All allowed values of RoutingMode enum
var AllowedRoutingModeEnumValues = []RoutingMode{
	"FAST",
	"SHORT",
	"MONETARY",
}

func (v *RoutingMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RoutingMode(value)
	for _, existing := range AllowedRoutingModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RoutingMode", value)
}

// NewRoutingModeFromValue returns a pointer to a valid RoutingMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRoutingModeFromValue(v string) (*RoutingMode, error) {
	ev := RoutingMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RoutingMode: valid values are %v", v, AllowedRoutingModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RoutingMode) IsValid() bool {
	for _, existing := range AllowedRoutingModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RoutingMode value
func (v RoutingMode) Ptr() *RoutingMode {
	return &v
}

type NullableRoutingMode struct {
	value *RoutingMode
	isSet bool
}

func (v NullableRoutingMode) Get() *RoutingMode {
	return v.value
}

func (v *NullableRoutingMode) Set(val *RoutingMode) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingMode) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingMode(val *RoutingMode) *NullableRoutingMode {
	return &NullableRoutingMode{value: val, isSet: true}
}

func (v NullableRoutingMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

