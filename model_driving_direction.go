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

// DrivingDirection Specifies the driving direction, i.e. from start to destination or from destination to start..  * `OUTBOUND` - Indicates an outbound routing from start to destination, i.e. the area which can be reached from the location within the given horizon. Use this value to calculate which can be covered e.g. by an emergency service such as a fire department.  * `INBOUND` - Indicates an inbound routing from destination to start, i.e. from where the location can be reached within the given horizon. Use this value to calculate the catchment area, e.g. of a school or a hospital.
type DrivingDirection string

// List of DrivingDirection
const (
	OUTBOUND DrivingDirection = "OUTBOUND"
	INBOUND DrivingDirection = "INBOUND"
)

// All allowed values of DrivingDirection enum
var AllowedDrivingDirectionEnumValues = []DrivingDirection{
	"OUTBOUND",
	"INBOUND",
}

func (v *DrivingDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DrivingDirection(value)
	for _, existing := range AllowedDrivingDirectionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DrivingDirection", value)
}

// NewDrivingDirectionFromValue returns a pointer to a valid DrivingDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDrivingDirectionFromValue(v string) (*DrivingDirection, error) {
	ev := DrivingDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DrivingDirection: valid values are %v", v, AllowedDrivingDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DrivingDirection) IsValid() bool {
	for _, existing := range AllowedDrivingDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DrivingDirection value
func (v DrivingDirection) Ptr() *DrivingDirection {
	return &v
}

type NullableDrivingDirection struct {
	value *DrivingDirection
	isSet bool
}

func (v NullableDrivingDirection) Get() *DrivingDirection {
	return v.value
}

func (v *NullableDrivingDirection) Set(val *DrivingDirection) {
	v.value = val
	v.isSet = true
}

func (v NullableDrivingDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableDrivingDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDrivingDirection(val *DrivingDirection) *NullableDrivingDirection {
	return &NullableDrivingDirection{value: val, isSet: true}
}

func (v NullableDrivingDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDrivingDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

