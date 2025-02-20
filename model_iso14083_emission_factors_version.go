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

// Iso14083EmissionFactorsVersion Defines the version of the emission factors to use for emission calculation based on ISO 14083. Will be ignored for other calculation methods.
type Iso14083EmissionFactorsVersion string

// List of Iso14083EmissionFactorsVersion
const (
	INITIAL Iso14083EmissionFactorsVersion = "INITIAL"
	VERSION_2 Iso14083EmissionFactorsVersion = "VERSION_2"
	LATEST Iso14083EmissionFactorsVersion = "LATEST"
)

// All allowed values of Iso14083EmissionFactorsVersion enum
var AllowedIso14083EmissionFactorsVersionEnumValues = []Iso14083EmissionFactorsVersion{
	"INITIAL",
	"VERSION_2",
	"LATEST",
}

func (v *Iso14083EmissionFactorsVersion) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Iso14083EmissionFactorsVersion(value)
	for _, existing := range AllowedIso14083EmissionFactorsVersionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Iso14083EmissionFactorsVersion", value)
}

// NewIso14083EmissionFactorsVersionFromValue returns a pointer to a valid Iso14083EmissionFactorsVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIso14083EmissionFactorsVersionFromValue(v string) (*Iso14083EmissionFactorsVersion, error) {
	ev := Iso14083EmissionFactorsVersion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Iso14083EmissionFactorsVersion: valid values are %v", v, AllowedIso14083EmissionFactorsVersionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Iso14083EmissionFactorsVersion) IsValid() bool {
	for _, existing := range AllowedIso14083EmissionFactorsVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Iso14083EmissionFactorsVersion value
func (v Iso14083EmissionFactorsVersion) Ptr() *Iso14083EmissionFactorsVersion {
	return &v
}

type NullableIso14083EmissionFactorsVersion struct {
	value *Iso14083EmissionFactorsVersion
	isSet bool
}

func (v NullableIso14083EmissionFactorsVersion) Get() *Iso14083EmissionFactorsVersion {
	return v.value
}

func (v *NullableIso14083EmissionFactorsVersion) Set(val *Iso14083EmissionFactorsVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableIso14083EmissionFactorsVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableIso14083EmissionFactorsVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIso14083EmissionFactorsVersion(val *Iso14083EmissionFactorsVersion) *NullableIso14083EmissionFactorsVersion {
	return &NullableIso14083EmissionFactorsVersion{value: val, isSet: true}
}

func (v NullableIso14083EmissionFactorsVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIso14083EmissionFactorsVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

