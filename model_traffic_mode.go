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

// TrafficMode Defines how to consider traffic in a route calculation.  * `REALISTIC` - Returns the most realistic **travelTime** and **distance** for the selected vehicle and the  given **startTime** or **arrivalTime** (or the current time if none of them is specified).  Takes into account the live traffic situation such as traffic jams or road works  as well as the typical traffic situation at the time of day and the day of week of travel such as the rush-hour  on Monday morning or light traffic on Saturday evening.  * `AVERAGE` - Returns the average **travelTime** and **distance** for the selected vehicle.  If **startTime** or **arrivalTime** is specified, the typical traffic situation for that time of day and day of week will be considered such as the rush-hour  on Monday morning or light traffic on Saturday evening. Toll will be calculated according to that date and time.  If none of them is specified the typical traffic situation will not be considered, and **travelTime** and **distance** are an average independent of when to travel.  Toll will be calculated for an arbitrary Monday at noon.  See [here](./concepts/traffic-modes) for more information. This parameter will be ignored for non-motorized profiles such as _BICYCLE_ or _PEDESTRIAN_.
type TrafficMode string

// List of TrafficMode
const (
	TRAFFIC_MODE_REALISTIC TrafficMode = "REALISTIC"
	TRAFFIC_MODE_AVERAGE   TrafficMode = "AVERAGE"
)

// All allowed values of TrafficMode enum
var AllowedTrafficModeEnumValues = []TrafficMode{
	"REALISTIC",
	"AVERAGE",
}

func (v *TrafficMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TrafficMode(value)
	for _, existing := range AllowedTrafficModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TrafficMode", value)
}

// NewTrafficModeFromValue returns a pointer to a valid TrafficMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTrafficModeFromValue(v string) (*TrafficMode, error) {
	ev := TrafficMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TrafficMode: valid values are %v", v, AllowedTrafficModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TrafficMode) IsValid() bool {
	for _, existing := range AllowedTrafficModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TrafficMode value
func (v TrafficMode) Ptr() *TrafficMode {
	return &v
}

type NullableTrafficMode struct {
	value *TrafficMode
	isSet bool
}

func (v NullableTrafficMode) Get() *TrafficMode {
	return v.value
}

func (v *NullableTrafficMode) Set(val *TrafficMode) {
	v.value = val
	v.isSet = true
}

func (v NullableTrafficMode) IsSet() bool {
	return v.isSet
}

func (v *NullableTrafficMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrafficMode(val *TrafficMode) *NullableTrafficMode {
	return &NullableTrafficMode{value: val, isSet: true}
}

func (v NullableTrafficMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrafficMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
