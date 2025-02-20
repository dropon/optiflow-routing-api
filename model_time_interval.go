/*
Routing

With the Routing service you can calculate routes from A to B taking into account vehicle-specific restrictions, traffic situations, toll, emissions, drivers' working hours, service times and opening intervals.

API version: 1.33
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package optiflow_routing

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the TimeInterval type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeInterval{}

// TimeInterval A time interval specified by two points in time - the beginning and the end of the interval.
type TimeInterval struct {
	// The beginning of the time interval formatted according to [RFC 3339](https://tools.ietf.org/html/rfc3339). The date must not be before 1970-01-01T00:00:00+00:00 nor after 2037-12-31T23:59:59+00:00.
	Start time.Time `json:"start"`
	// The end of the time interval formatted according to [RFC 3339](https://tools.ietf.org/html/rfc3339). The date must not be before 1970-01-01T00:00:00+00:00 nor after 2037-12-31T23:59:59+00:00.
	End time.Time `json:"end"`
}

type _TimeInterval TimeInterval

// NewTimeInterval instantiates a new TimeInterval object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeInterval(start time.Time, end time.Time) *TimeInterval {
	this := TimeInterval{}
	this.Start = start
	this.End = end
	return &this
}

// NewTimeIntervalWithDefaults instantiates a new TimeInterval object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeIntervalWithDefaults() *TimeInterval {
	this := TimeInterval{}
	return &this
}

// GetStart returns the Start field value
func (o *TimeInterval) GetStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *TimeInterval) GetStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *TimeInterval) SetStart(v time.Time) {
	o.Start = v
}

// GetEnd returns the End field value
func (o *TimeInterval) GetEnd() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *TimeInterval) GetEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *TimeInterval) SetEnd(v time.Time) {
	o.End = v
}

func (o TimeInterval) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeInterval) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["start"] = o.Start
	toSerialize["end"] = o.End
	return toSerialize, nil
}

func (o *TimeInterval) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"start",
		"end",
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

	varTimeInterval := _TimeInterval{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTimeInterval)

	if err != nil {
		return err
	}

	*o = TimeInterval(varTimeInterval)

	return err
}

type NullableTimeInterval struct {
	value *TimeInterval
	isSet bool
}

func (v NullableTimeInterval) Get() *TimeInterval {
	return v.value
}

func (v *NullableTimeInterval) Set(val *TimeInterval) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeInterval) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeInterval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeInterval(val *TimeInterval) *NullableTimeInterval {
	return &NullableTimeInterval{value: val, isSet: true}
}

func (v NullableTimeInterval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeInterval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


