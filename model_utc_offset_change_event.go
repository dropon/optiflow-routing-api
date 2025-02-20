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

// checks if the UTCOffsetChangeEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UTCOffsetChangeEvent{}

// UTCOffsetChangeEvent Issued when the offset to UTC changes, mostly when traveling into a different time zone. Requires _UTC_OFFSET_CHANGE_EVENTS_ to be requested.    Changing the UTC offset does not necessarily mean to change the time zone. Vice-versa changing the time zone does not necessarily mean to change the UTC offset. There are some special cases to consider.  *  The UTC offset may change even within a time zone when the route takes place exactly when the daylight-saving time changes.  *  The UTC offset may not change when changing the time zone. In Canada, for example, there are regions which do not use DST   so that the neighboring time zone has the same UTC offset in summer.
type UTCOffsetChangeEvent struct {
	// The new UTC offset [min].
	UtcOffset int32 `json:"utcOffset"`
}

type _UTCOffsetChangeEvent UTCOffsetChangeEvent

// NewUTCOffsetChangeEvent instantiates a new UTCOffsetChangeEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUTCOffsetChangeEvent(utcOffset int32) *UTCOffsetChangeEvent {
	this := UTCOffsetChangeEvent{}
	this.UtcOffset = utcOffset
	return &this
}

// NewUTCOffsetChangeEventWithDefaults instantiates a new UTCOffsetChangeEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUTCOffsetChangeEventWithDefaults() *UTCOffsetChangeEvent {
	this := UTCOffsetChangeEvent{}
	return &this
}

// GetUtcOffset returns the UtcOffset field value
func (o *UTCOffsetChangeEvent) GetUtcOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UtcOffset
}

// GetUtcOffsetOk returns a tuple with the UtcOffset field value
// and a boolean to check if the value has been set.
func (o *UTCOffsetChangeEvent) GetUtcOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UtcOffset, true
}

// SetUtcOffset sets field value
func (o *UTCOffsetChangeEvent) SetUtcOffset(v int32) {
	o.UtcOffset = v
}

func (o UTCOffsetChangeEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UTCOffsetChangeEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["utcOffset"] = o.UtcOffset
	return toSerialize, nil
}

func (o *UTCOffsetChangeEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"utcOffset",
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

	varUTCOffsetChangeEvent := _UTCOffsetChangeEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUTCOffsetChangeEvent)

	if err != nil {
		return err
	}

	*o = UTCOffsetChangeEvent(varUTCOffsetChangeEvent)

	return err
}

type NullableUTCOffsetChangeEvent struct {
	value *UTCOffsetChangeEvent
	isSet bool
}

func (v NullableUTCOffsetChangeEvent) Get() *UTCOffsetChangeEvent {
	return v.value
}

func (v *NullableUTCOffsetChangeEvent) Set(val *UTCOffsetChangeEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableUTCOffsetChangeEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableUTCOffsetChangeEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUTCOffsetChangeEvent(val *UTCOffsetChangeEvent) *NullableUTCOffsetChangeEvent {
	return &NullableUTCOffsetChangeEvent{value: val, isSet: true}
}

func (v NullableUTCOffsetChangeEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUTCOffsetChangeEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


