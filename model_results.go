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

// Results the model 'Results'
type Results string

// List of Results
const (
	ROUTE_ID Results = "ROUTE_ID"
	LEGS Results = "LEGS"
	LEGS_POLYLINE Results = "LEGS_POLYLINE"
	TOLL_COSTS Results = "TOLL_COSTS"
	TOLL_SECTIONS Results = "TOLL_SECTIONS"
	TOLL_SYSTEMS Results = "TOLL_SYSTEMS"
	TOLL_EVENTS Results = "TOLL_EVENTS"
	POLYLINE Results = "POLYLINE"
	MANEUVER_EVENTS Results = "MANEUVER_EVENTS"
	BORDER_EVENTS Results = "BORDER_EVENTS"
	VIOLATION_EVENTS Results = "VIOLATION_EVENTS"
	VIOLATION_EVENTS_POLYLINE Results = "VIOLATION_EVENTS_POLYLINE"
	WAYPOINT_EVENTS Results = "WAYPOINT_EVENTS"
	UTC_OFFSET_CHANGE_EVENTS Results = "UTC_OFFSET_CHANGE_EVENTS"
	COMBINED_TRANSPORT_EVENTS Results = "COMBINED_TRANSPORT_EVENTS"
	TRAFFIC_EVENTS Results = "TRAFFIC_EVENTS"
	TRAFFIC_EVENTS_POLYLINE Results = "TRAFFIC_EVENTS_POLYLINE"
	LOW_EMISSION_ZONE_EVENTS Results = "LOW_EMISSION_ZONE_EVENTS"
	DELIVERY_ONLY_EVENTS Results = "DELIVERY_ONLY_EVENTS"
	DELIVERY_ONLY_EVENTS_POLYLINE Results = "DELIVERY_ONLY_EVENTS_POLYLINE"
	SCHEDULE_EVENTS Results = "SCHEDULE_EVENTS"
	EMISSIONS_EN16258_2012 Results = "EMISSIONS_EN16258_2012"
	EMISSIONS_EN16258_2012_HBEFA Results = "EMISSIONS_EN16258_2012_HBEFA"
	EMISSIONS_ISO14083_2022 Results = "EMISSIONS_ISO14083_2022"
	EMISSIONS_ISO14083_2022_DEFAULT_CONSUMPTION Results = "EMISSIONS_ISO14083_2022_DEFAULT_CONSUMPTION"
	EMISSIONS_ISO14083_2023 Results = "EMISSIONS_ISO14083_2023"
	EMISSIONS_ISO14083_2023_DEFAULT_CONSUMPTION Results = "EMISSIONS_ISO14083_2023_DEFAULT_CONSUMPTION"
	EMISSIONS_FRENCH_CO2E_DECREE_2017_639 Results = "EMISSIONS_FRENCH_CO2E_DECREE_2017_639"
	ALTERNATIVE_ROUTES Results = "ALTERNATIVE_ROUTES"
	SCHEDULE_REPORT Results = "SCHEDULE_REPORT"
	GUIDED_NAVIGATION Results = "GUIDED_NAVIGATION"
	MONETARY_COSTS Results = "MONETARY_COSTS"
	EV_REPORT Results = "EV_REPORT"
	EV_STATUS_EVENTS Results = "EV_STATUS_EVENTS"
	EV_STATUS_EVENTS_POLYLINE Results = "EV_STATUS_EVENTS_POLYLINE"
	EV_CHARGE_EVENTS Results = "EV_CHARGE_EVENTS"
)

// All allowed values of Results enum
var AllowedResultsEnumValues = []Results{
	"ROUTE_ID",
	"LEGS",
	"LEGS_POLYLINE",
	"TOLL_COSTS",
	"TOLL_SECTIONS",
	"TOLL_SYSTEMS",
	"TOLL_EVENTS",
	"POLYLINE",
	"MANEUVER_EVENTS",
	"BORDER_EVENTS",
	"VIOLATION_EVENTS",
	"VIOLATION_EVENTS_POLYLINE",
	"WAYPOINT_EVENTS",
	"UTC_OFFSET_CHANGE_EVENTS",
	"COMBINED_TRANSPORT_EVENTS",
	"TRAFFIC_EVENTS",
	"TRAFFIC_EVENTS_POLYLINE",
	"LOW_EMISSION_ZONE_EVENTS",
	"DELIVERY_ONLY_EVENTS",
	"DELIVERY_ONLY_EVENTS_POLYLINE",
	"SCHEDULE_EVENTS",
	"EMISSIONS_EN16258_2012",
	"EMISSIONS_EN16258_2012_HBEFA",
	"EMISSIONS_ISO14083_2022",
	"EMISSIONS_ISO14083_2022_DEFAULT_CONSUMPTION",
	"EMISSIONS_ISO14083_2023",
	"EMISSIONS_ISO14083_2023_DEFAULT_CONSUMPTION",
	"EMISSIONS_FRENCH_CO2E_DECREE_2017_639",
	"ALTERNATIVE_ROUTES",
	"SCHEDULE_REPORT",
	"GUIDED_NAVIGATION",
	"MONETARY_COSTS",
	"EV_REPORT",
	"EV_STATUS_EVENTS",
	"EV_STATUS_EVENTS_POLYLINE",
	"EV_CHARGE_EVENTS",
}

func (v *Results) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Results(value)
	for _, existing := range AllowedResultsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Results", value)
}

// NewResultsFromValue returns a pointer to a valid Results
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResultsFromValue(v string) (*Results, error) {
	ev := Results(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Results: valid values are %v", v, AllowedResultsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Results) IsValid() bool {
	for _, existing := range AllowedResultsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Results value
func (v Results) Ptr() *Results {
	return &v
}

type NullableResults struct {
	value *Results
	isSet bool
}

func (v NullableResults) Get() *Results {
	return v.value
}

func (v *NullableResults) Set(val *Results) {
	v.value = val
	v.isSet = true
}

func (v NullableResults) IsSet() bool {
	return v.isSet
}

func (v *NullableResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResults(val *Results) *NullableResults {
	return &NullableResults{value: val, isSet: true}
}

func (v NullableResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

