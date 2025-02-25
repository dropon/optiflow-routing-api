/*
Routing

With the Routing service you can calculate routes from A to B taking into account vehicle-specific restrictions, traffic situations, toll, emissions, drivers' working hours, service times and opening intervals.

API version: 1.33
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package optiflow_routing

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the EvReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EvReport{}

// EvReport The consumption and charging summary for the specified vehicle model. Only present if _EV_REPORT_ is requested. 
type EvReport struct {
	// The electricity consumption since the start of the route [kWh].
	ElectricityConsumption float64 `json:"electricityConsumption"`
	// The remaining state of charge at the end of the route [%]. May be below the defined minimum state of charge or even below zero, if the electricity consumption exceeds the available energy in the battery and charging is not possible before falling below zero.
	BatteryStateOfCharge float64 `json:"batteryStateOfCharge"`
	// The time spent for charging the battery [s]. The charging time is a proposal, currently for information only. It is not included in the travel time of the route and the start time of subsequent events is not offset by it.
	ChargingTime int32 `json:"chargingTime"`
	// The amount of electricity charged along the route [kWh].
	ElectricityCharged float64 `json:"electricityCharged"`
	// The percentage of battery charged along the route [%].
	PercentageCharged int32 `json:"percentageCharged"`
	Cost ElectricityPrice `json:"cost"`
}

type _EvReport EvReport

// NewEvReport instantiates a new EvReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvReport(electricityConsumption float64, batteryStateOfCharge float64, chargingTime int32, electricityCharged float64, percentageCharged int32, cost ElectricityPrice) *EvReport {
	this := EvReport{}
	this.ElectricityConsumption = electricityConsumption
	this.BatteryStateOfCharge = batteryStateOfCharge
	this.ChargingTime = chargingTime
	this.ElectricityCharged = electricityCharged
	this.PercentageCharged = percentageCharged
	this.Cost = cost
	return &this
}

// NewEvReportWithDefaults instantiates a new EvReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEvReportWithDefaults() *EvReport {
	this := EvReport{}
	return &this
}

// GetElectricityConsumption returns the ElectricityConsumption field value
func (o *EvReport) GetElectricityConsumption() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.ElectricityConsumption
}

// GetElectricityConsumptionOk returns a tuple with the ElectricityConsumption field value
// and a boolean to check if the value has been set.
func (o *EvReport) GetElectricityConsumptionOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ElectricityConsumption, true
}

// SetElectricityConsumption sets field value
func (o *EvReport) SetElectricityConsumption(v float64) {
	o.ElectricityConsumption = v
}

// GetBatteryStateOfCharge returns the BatteryStateOfCharge field value
func (o *EvReport) GetBatteryStateOfCharge() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.BatteryStateOfCharge
}

// GetBatteryStateOfChargeOk returns a tuple with the BatteryStateOfCharge field value
// and a boolean to check if the value has been set.
func (o *EvReport) GetBatteryStateOfChargeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BatteryStateOfCharge, true
}

// SetBatteryStateOfCharge sets field value
func (o *EvReport) SetBatteryStateOfCharge(v float64) {
	o.BatteryStateOfCharge = v
}

// GetChargingTime returns the ChargingTime field value
func (o *EvReport) GetChargingTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ChargingTime
}

// GetChargingTimeOk returns a tuple with the ChargingTime field value
// and a boolean to check if the value has been set.
func (o *EvReport) GetChargingTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChargingTime, true
}

// SetChargingTime sets field value
func (o *EvReport) SetChargingTime(v int32) {
	o.ChargingTime = v
}

// GetElectricityCharged returns the ElectricityCharged field value
func (o *EvReport) GetElectricityCharged() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.ElectricityCharged
}

// GetElectricityChargedOk returns a tuple with the ElectricityCharged field value
// and a boolean to check if the value has been set.
func (o *EvReport) GetElectricityChargedOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ElectricityCharged, true
}

// SetElectricityCharged sets field value
func (o *EvReport) SetElectricityCharged(v float64) {
	o.ElectricityCharged = v
}

// GetPercentageCharged returns the PercentageCharged field value
func (o *EvReport) GetPercentageCharged() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PercentageCharged
}

// GetPercentageChargedOk returns a tuple with the PercentageCharged field value
// and a boolean to check if the value has been set.
func (o *EvReport) GetPercentageChargedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PercentageCharged, true
}

// SetPercentageCharged sets field value
func (o *EvReport) SetPercentageCharged(v int32) {
	o.PercentageCharged = v
}

// GetCost returns the Cost field value
func (o *EvReport) GetCost() ElectricityPrice {
	if o == nil {
		var ret ElectricityPrice
		return ret
	}

	return o.Cost
}

// GetCostOk returns a tuple with the Cost field value
// and a boolean to check if the value has been set.
func (o *EvReport) GetCostOk() (*ElectricityPrice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cost, true
}

// SetCost sets field value
func (o *EvReport) SetCost(v ElectricityPrice) {
	o.Cost = v
}

func (o EvReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EvReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["electricityConsumption"] = o.ElectricityConsumption
	toSerialize["batteryStateOfCharge"] = o.BatteryStateOfCharge
	toSerialize["chargingTime"] = o.ChargingTime
	toSerialize["electricityCharged"] = o.ElectricityCharged
	toSerialize["percentageCharged"] = o.PercentageCharged
	toSerialize["cost"] = o.Cost
	return toSerialize, nil
}

func (o *EvReport) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"electricityConsumption",
		"batteryStateOfCharge",
		"chargingTime",
		"electricityCharged",
		"percentageCharged",
		"cost",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varEvReport := _EvReport{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEvReport)

	if err != nil {
		return err
	}

	*o = EvReport(varEvReport)

	return err
}

type NullableEvReport struct {
	value *EvReport
	isSet bool
}

func (v NullableEvReport) Get() *EvReport {
	return v.value
}

func (v *NullableEvReport) Set(val *EvReport) {
	v.value = val
	v.isSet = true
}

func (v NullableEvReport) IsSet() bool {
	return v.isSet
}

func (v *NullableEvReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvReport(val *EvReport) *NullableEvReport {
	return &NullableEvReport{value: val, isSet: true}
}

func (v NullableEvReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


