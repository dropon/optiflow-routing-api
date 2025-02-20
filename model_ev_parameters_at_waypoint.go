/*
Routing

With the Routing service you can calculate routes from A to B taking into account vehicle-specific restrictions, traffic situations, toll, emissions, drivers' working hours, service times and opening intervals.

API version: 1.33
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package optiflow_routing

import (
	"encoding/json"
)

// checks if the EvParametersAtWaypoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EvParametersAtWaypoint{}

// EvParametersAtWaypoint The ev parameters that are specific to a waypoint.  This parameter is in a preview state, the API is stable, feature changes could be introduced in future. 
type EvParametersAtWaypoint struct {
	Weather *Weather `json:"weather,omitempty"`
	ChargingStation *ChargingStation `json:"chargingStation,omitempty"`
	// The electricity consumed during service by electrical appliances (e.g. crane, cooling, tail lift) [kWh]. 
	PowerConsumptionDuringService *float64 `json:"powerConsumptionDuringService,omitempty"`
}

// NewEvParametersAtWaypoint instantiates a new EvParametersAtWaypoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvParametersAtWaypoint() *EvParametersAtWaypoint {
	this := EvParametersAtWaypoint{}
	var powerConsumptionDuringService float64 = 0
	this.PowerConsumptionDuringService = &powerConsumptionDuringService
	return &this
}

// NewEvParametersAtWaypointWithDefaults instantiates a new EvParametersAtWaypoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEvParametersAtWaypointWithDefaults() *EvParametersAtWaypoint {
	this := EvParametersAtWaypoint{}
	var powerConsumptionDuringService float64 = 0
	this.PowerConsumptionDuringService = &powerConsumptionDuringService
	return &this
}

// GetWeather returns the Weather field value if set, zero value otherwise.
func (o *EvParametersAtWaypoint) GetWeather() Weather {
	if o == nil || IsNil(o.Weather) {
		var ret Weather
		return ret
	}
	return *o.Weather
}

// GetWeatherOk returns a tuple with the Weather field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvParametersAtWaypoint) GetWeatherOk() (*Weather, bool) {
	if o == nil || IsNil(o.Weather) {
		return nil, false
	}
	return o.Weather, true
}

// HasWeather returns a boolean if a field has been set.
func (o *EvParametersAtWaypoint) HasWeather() bool {
	if o != nil && !IsNil(o.Weather) {
		return true
	}

	return false
}

// SetWeather gets a reference to the given Weather and assigns it to the Weather field.
func (o *EvParametersAtWaypoint) SetWeather(v Weather) {
	o.Weather = &v
}

// GetChargingStation returns the ChargingStation field value if set, zero value otherwise.
func (o *EvParametersAtWaypoint) GetChargingStation() ChargingStation {
	if o == nil || IsNil(o.ChargingStation) {
		var ret ChargingStation
		return ret
	}
	return *o.ChargingStation
}

// GetChargingStationOk returns a tuple with the ChargingStation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvParametersAtWaypoint) GetChargingStationOk() (*ChargingStation, bool) {
	if o == nil || IsNil(o.ChargingStation) {
		return nil, false
	}
	return o.ChargingStation, true
}

// HasChargingStation returns a boolean if a field has been set.
func (o *EvParametersAtWaypoint) HasChargingStation() bool {
	if o != nil && !IsNil(o.ChargingStation) {
		return true
	}

	return false
}

// SetChargingStation gets a reference to the given ChargingStation and assigns it to the ChargingStation field.
func (o *EvParametersAtWaypoint) SetChargingStation(v ChargingStation) {
	o.ChargingStation = &v
}

// GetPowerConsumptionDuringService returns the PowerConsumptionDuringService field value if set, zero value otherwise.
func (o *EvParametersAtWaypoint) GetPowerConsumptionDuringService() float64 {
	if o == nil || IsNil(o.PowerConsumptionDuringService) {
		var ret float64
		return ret
	}
	return *o.PowerConsumptionDuringService
}

// GetPowerConsumptionDuringServiceOk returns a tuple with the PowerConsumptionDuringService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvParametersAtWaypoint) GetPowerConsumptionDuringServiceOk() (*float64, bool) {
	if o == nil || IsNil(o.PowerConsumptionDuringService) {
		return nil, false
	}
	return o.PowerConsumptionDuringService, true
}

// HasPowerConsumptionDuringService returns a boolean if a field has been set.
func (o *EvParametersAtWaypoint) HasPowerConsumptionDuringService() bool {
	if o != nil && !IsNil(o.PowerConsumptionDuringService) {
		return true
	}

	return false
}

// SetPowerConsumptionDuringService gets a reference to the given float64 and assigns it to the PowerConsumptionDuringService field.
func (o *EvParametersAtWaypoint) SetPowerConsumptionDuringService(v float64) {
	o.PowerConsumptionDuringService = &v
}

func (o EvParametersAtWaypoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EvParametersAtWaypoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Weather) {
		toSerialize["weather"] = o.Weather
	}
	if !IsNil(o.ChargingStation) {
		toSerialize["chargingStation"] = o.ChargingStation
	}
	if !IsNil(o.PowerConsumptionDuringService) {
		toSerialize["powerConsumptionDuringService"] = o.PowerConsumptionDuringService
	}
	return toSerialize, nil
}

type NullableEvParametersAtWaypoint struct {
	value *EvParametersAtWaypoint
	isSet bool
}

func (v NullableEvParametersAtWaypoint) Get() *EvParametersAtWaypoint {
	return v.value
}

func (v *NullableEvParametersAtWaypoint) Set(val *EvParametersAtWaypoint) {
	v.value = val
	v.isSet = true
}

func (v NullableEvParametersAtWaypoint) IsSet() bool {
	return v.isSet
}

func (v *NullableEvParametersAtWaypoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvParametersAtWaypoint(val *EvParametersAtWaypoint) *NullableEvParametersAtWaypoint {
	return &NullableEvParametersAtWaypoint{value: val, isSet: true}
}

func (v NullableEvParametersAtWaypoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvParametersAtWaypoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


