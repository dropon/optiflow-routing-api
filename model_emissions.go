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

// checks if the Emissions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Emissions{}

// Emissions Emissions such as the CO2-equivalent (CO2e) according to the selected standards.  The CO2e value is the unit for comparing the radiative forcing of a greenhouse gas to carbon dioxide according to [ISO 14064-1:2006](https://en.wikipedia.org/wiki/ISO_14064).
type Emissions struct {
	En162582012 *EmissionsEN162582012 `json:"en16258_2012,omitempty"`
	Iso140832022 *EmissionsISO140832022 `json:"iso14083_2022,omitempty"`
	Iso140832023 *EmissionsISO140832023 `json:"iso14083_2023,omitempty"`
	FrenchCO2eDecree2017639 *EmissionsFrenchCO2eDecree2017639 `json:"frenchCO2eDecree2017_639,omitempty"`
}

// NewEmissions instantiates a new Emissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmissions() *Emissions {
	this := Emissions{}
	return &this
}

// NewEmissionsWithDefaults instantiates a new Emissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmissionsWithDefaults() *Emissions {
	this := Emissions{}
	return &this
}

// GetEn162582012 returns the En162582012 field value if set, zero value otherwise.
func (o *Emissions) GetEn162582012() EmissionsEN162582012 {
	if o == nil || IsNil(o.En162582012) {
		var ret EmissionsEN162582012
		return ret
	}
	return *o.En162582012
}

// GetEn162582012Ok returns a tuple with the En162582012 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Emissions) GetEn162582012Ok() (*EmissionsEN162582012, bool) {
	if o == nil || IsNil(o.En162582012) {
		return nil, false
	}
	return o.En162582012, true
}

// HasEn162582012 returns a boolean if a field has been set.
func (o *Emissions) HasEn162582012() bool {
	if o != nil && !IsNil(o.En162582012) {
		return true
	}

	return false
}

// SetEn162582012 gets a reference to the given EmissionsEN162582012 and assigns it to the En162582012 field.
func (o *Emissions) SetEn162582012(v EmissionsEN162582012) {
	o.En162582012 = &v
}

// GetIso140832022 returns the Iso140832022 field value if set, zero value otherwise.
func (o *Emissions) GetIso140832022() EmissionsISO140832022 {
	if o == nil || IsNil(o.Iso140832022) {
		var ret EmissionsISO140832022
		return ret
	}
	return *o.Iso140832022
}

// GetIso140832022Ok returns a tuple with the Iso140832022 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Emissions) GetIso140832022Ok() (*EmissionsISO140832022, bool) {
	if o == nil || IsNil(o.Iso140832022) {
		return nil, false
	}
	return o.Iso140832022, true
}

// HasIso140832022 returns a boolean if a field has been set.
func (o *Emissions) HasIso140832022() bool {
	if o != nil && !IsNil(o.Iso140832022) {
		return true
	}

	return false
}

// SetIso140832022 gets a reference to the given EmissionsISO140832022 and assigns it to the Iso140832022 field.
func (o *Emissions) SetIso140832022(v EmissionsISO140832022) {
	o.Iso140832022 = &v
}

// GetIso140832023 returns the Iso140832023 field value if set, zero value otherwise.
func (o *Emissions) GetIso140832023() EmissionsISO140832023 {
	if o == nil || IsNil(o.Iso140832023) {
		var ret EmissionsISO140832023
		return ret
	}
	return *o.Iso140832023
}

// GetIso140832023Ok returns a tuple with the Iso140832023 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Emissions) GetIso140832023Ok() (*EmissionsISO140832023, bool) {
	if o == nil || IsNil(o.Iso140832023) {
		return nil, false
	}
	return o.Iso140832023, true
}

// HasIso140832023 returns a boolean if a field has been set.
func (o *Emissions) HasIso140832023() bool {
	if o != nil && !IsNil(o.Iso140832023) {
		return true
	}

	return false
}

// SetIso140832023 gets a reference to the given EmissionsISO140832023 and assigns it to the Iso140832023 field.
func (o *Emissions) SetIso140832023(v EmissionsISO140832023) {
	o.Iso140832023 = &v
}

// GetFrenchCO2eDecree2017639 returns the FrenchCO2eDecree2017639 field value if set, zero value otherwise.
func (o *Emissions) GetFrenchCO2eDecree2017639() EmissionsFrenchCO2eDecree2017639 {
	if o == nil || IsNil(o.FrenchCO2eDecree2017639) {
		var ret EmissionsFrenchCO2eDecree2017639
		return ret
	}
	return *o.FrenchCO2eDecree2017639
}

// GetFrenchCO2eDecree2017639Ok returns a tuple with the FrenchCO2eDecree2017639 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Emissions) GetFrenchCO2eDecree2017639Ok() (*EmissionsFrenchCO2eDecree2017639, bool) {
	if o == nil || IsNil(o.FrenchCO2eDecree2017639) {
		return nil, false
	}
	return o.FrenchCO2eDecree2017639, true
}

// HasFrenchCO2eDecree2017639 returns a boolean if a field has been set.
func (o *Emissions) HasFrenchCO2eDecree2017639() bool {
	if o != nil && !IsNil(o.FrenchCO2eDecree2017639) {
		return true
	}

	return false
}

// SetFrenchCO2eDecree2017639 gets a reference to the given EmissionsFrenchCO2eDecree2017639 and assigns it to the FrenchCO2eDecree2017639 field.
func (o *Emissions) SetFrenchCO2eDecree2017639(v EmissionsFrenchCO2eDecree2017639) {
	o.FrenchCO2eDecree2017639 = &v
}

func (o Emissions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Emissions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.En162582012) {
		toSerialize["en16258_2012"] = o.En162582012
	}
	if !IsNil(o.Iso140832022) {
		toSerialize["iso14083_2022"] = o.Iso140832022
	}
	if !IsNil(o.Iso140832023) {
		toSerialize["iso14083_2023"] = o.Iso140832023
	}
	if !IsNil(o.FrenchCO2eDecree2017639) {
		toSerialize["frenchCO2eDecree2017_639"] = o.FrenchCO2eDecree2017639
	}
	return toSerialize, nil
}

type NullableEmissions struct {
	value *Emissions
	isSet bool
}

func (v NullableEmissions) Get() *Emissions {
	return v.value
}

func (v *NullableEmissions) Set(val *Emissions) {
	v.value = val
	v.isSet = true
}

func (v NullableEmissions) IsSet() bool {
	return v.isSet
}

func (v *NullableEmissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmissions(val *Emissions) *NullableEmissions {
	return &NullableEmissions{value: val, isSet: true}
}

func (v NullableEmissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


