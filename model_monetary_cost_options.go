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

// checks if the MonetaryCostOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonetaryCostOptions{}

// MonetaryCostOptions struct for MonetaryCostOptions
type MonetaryCostOptions struct {
	// Specifies the cost per kilometer. This value can contain the cost from the energy consumption, but it also possible to define the energy costs using **monetaryCostOptions[costPerKwh]** and **monetaryCostOptions[costPerFuelUnit]**. If it is not specified the default value depends on the used **profile**. The default value of the profile can change at any time.
	CostPerKilometer *float64 `json:"costPerKilometer,omitempty"`
	// Specifies the cost per hour. It is not applied to service, break or rest periods. If it is not specified the default value depends on the used **profile**. The default value of the profile can change at any time.
	WorkingCostPerHour *float64 `json:"workingCostPerHour,omitempty"`
	// Specifies the cost per kilowatt hour. Only relevant for **vehicle[engineType]** _ELECTRIC_ and _HYBRID_.
	CostPerKwh *float64 `json:"costPerKwh,omitempty"`
	// Specifies the cost per fuel unit (per liter Diesel or per kg CNG). Only used for **vehicle[engineType]** _COMBUSTION_ and _HYBRID_.
	CostPerFuelUnit *float64 `json:"costPerFuelUnit,omitempty"`
}

// NewMonetaryCostOptions instantiates a new MonetaryCostOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonetaryCostOptions() *MonetaryCostOptions {
	this := MonetaryCostOptions{}
	return &this
}

// NewMonetaryCostOptionsWithDefaults instantiates a new MonetaryCostOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonetaryCostOptionsWithDefaults() *MonetaryCostOptions {
	this := MonetaryCostOptions{}
	return &this
}

// GetCostPerKilometer returns the CostPerKilometer field value if set, zero value otherwise.
func (o *MonetaryCostOptions) GetCostPerKilometer() float64 {
	if o == nil || IsNil(o.CostPerKilometer) {
		var ret float64
		return ret
	}
	return *o.CostPerKilometer
}

// GetCostPerKilometerOk returns a tuple with the CostPerKilometer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonetaryCostOptions) GetCostPerKilometerOk() (*float64, bool) {
	if o == nil || IsNil(o.CostPerKilometer) {
		return nil, false
	}
	return o.CostPerKilometer, true
}

// HasCostPerKilometer returns a boolean if a field has been set.
func (o *MonetaryCostOptions) HasCostPerKilometer() bool {
	if o != nil && !IsNil(o.CostPerKilometer) {
		return true
	}

	return false
}

// SetCostPerKilometer gets a reference to the given float64 and assigns it to the CostPerKilometer field.
func (o *MonetaryCostOptions) SetCostPerKilometer(v float64) {
	o.CostPerKilometer = &v
}

// GetWorkingCostPerHour returns the WorkingCostPerHour field value if set, zero value otherwise.
func (o *MonetaryCostOptions) GetWorkingCostPerHour() float64 {
	if o == nil || IsNil(o.WorkingCostPerHour) {
		var ret float64
		return ret
	}
	return *o.WorkingCostPerHour
}

// GetWorkingCostPerHourOk returns a tuple with the WorkingCostPerHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonetaryCostOptions) GetWorkingCostPerHourOk() (*float64, bool) {
	if o == nil || IsNil(o.WorkingCostPerHour) {
		return nil, false
	}
	return o.WorkingCostPerHour, true
}

// HasWorkingCostPerHour returns a boolean if a field has been set.
func (o *MonetaryCostOptions) HasWorkingCostPerHour() bool {
	if o != nil && !IsNil(o.WorkingCostPerHour) {
		return true
	}

	return false
}

// SetWorkingCostPerHour gets a reference to the given float64 and assigns it to the WorkingCostPerHour field.
func (o *MonetaryCostOptions) SetWorkingCostPerHour(v float64) {
	o.WorkingCostPerHour = &v
}

// GetCostPerKwh returns the CostPerKwh field value if set, zero value otherwise.
func (o *MonetaryCostOptions) GetCostPerKwh() float64 {
	if o == nil || IsNil(o.CostPerKwh) {
		var ret float64
		return ret
	}
	return *o.CostPerKwh
}

// GetCostPerKwhOk returns a tuple with the CostPerKwh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonetaryCostOptions) GetCostPerKwhOk() (*float64, bool) {
	if o == nil || IsNil(o.CostPerKwh) {
		return nil, false
	}
	return o.CostPerKwh, true
}

// HasCostPerKwh returns a boolean if a field has been set.
func (o *MonetaryCostOptions) HasCostPerKwh() bool {
	if o != nil && !IsNil(o.CostPerKwh) {
		return true
	}

	return false
}

// SetCostPerKwh gets a reference to the given float64 and assigns it to the CostPerKwh field.
func (o *MonetaryCostOptions) SetCostPerKwh(v float64) {
	o.CostPerKwh = &v
}

// GetCostPerFuelUnit returns the CostPerFuelUnit field value if set, zero value otherwise.
func (o *MonetaryCostOptions) GetCostPerFuelUnit() float64 {
	if o == nil || IsNil(o.CostPerFuelUnit) {
		var ret float64
		return ret
	}
	return *o.CostPerFuelUnit
}

// GetCostPerFuelUnitOk returns a tuple with the CostPerFuelUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonetaryCostOptions) GetCostPerFuelUnitOk() (*float64, bool) {
	if o == nil || IsNil(o.CostPerFuelUnit) {
		return nil, false
	}
	return o.CostPerFuelUnit, true
}

// HasCostPerFuelUnit returns a boolean if a field has been set.
func (o *MonetaryCostOptions) HasCostPerFuelUnit() bool {
	if o != nil && !IsNil(o.CostPerFuelUnit) {
		return true
	}

	return false
}

// SetCostPerFuelUnit gets a reference to the given float64 and assigns it to the CostPerFuelUnit field.
func (o *MonetaryCostOptions) SetCostPerFuelUnit(v float64) {
	o.CostPerFuelUnit = &v
}

func (o MonetaryCostOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonetaryCostOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CostPerKilometer) {
		toSerialize["costPerKilometer"] = o.CostPerKilometer
	}
	if !IsNil(o.WorkingCostPerHour) {
		toSerialize["workingCostPerHour"] = o.WorkingCostPerHour
	}
	if !IsNil(o.CostPerKwh) {
		toSerialize["costPerKwh"] = o.CostPerKwh
	}
	if !IsNil(o.CostPerFuelUnit) {
		toSerialize["costPerFuelUnit"] = o.CostPerFuelUnit
	}
	return toSerialize, nil
}

type NullableMonetaryCostOptions struct {
	value *MonetaryCostOptions
	isSet bool
}

func (v NullableMonetaryCostOptions) Get() *MonetaryCostOptions {
	return v.value
}

func (v *NullableMonetaryCostOptions) Set(val *MonetaryCostOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableMonetaryCostOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableMonetaryCostOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonetaryCostOptions(val *MonetaryCostOptions) *NullableMonetaryCostOptions {
	return &NullableMonetaryCostOptions{value: val, isSet: true}
}

func (v NullableMonetaryCostOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonetaryCostOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


