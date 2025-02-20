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

// checks if the ReachableLocationsId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReachableLocationsId{}

// ReachableLocationsId struct for ReachableLocationsId
type ReachableLocationsId struct {
	// The ID of the calculated reachable locations.
	Id string `json:"id"`
}

type _ReachableLocationsId ReachableLocationsId

// NewReachableLocationsId instantiates a new ReachableLocationsId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReachableLocationsId(id string) *ReachableLocationsId {
	this := ReachableLocationsId{}
	this.Id = id
	return &this
}

// NewReachableLocationsIdWithDefaults instantiates a new ReachableLocationsId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReachableLocationsIdWithDefaults() *ReachableLocationsId {
	this := ReachableLocationsId{}
	return &this
}

// GetId returns the Id field value
func (o *ReachableLocationsId) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReachableLocationsId) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReachableLocationsId) SetId(v string) {
	o.Id = v
}

func (o ReachableLocationsId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReachableLocationsId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *ReachableLocationsId) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varReachableLocationsId := _ReachableLocationsId{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReachableLocationsId)

	if err != nil {
		return err
	}

	*o = ReachableLocationsId(varReachableLocationsId)

	return err
}

type NullableReachableLocationsId struct {
	value *ReachableLocationsId
	isSet bool
}

func (v NullableReachableLocationsId) Get() *ReachableLocationsId {
	return v.value
}

func (v *NullableReachableLocationsId) Set(val *ReachableLocationsId) {
	v.value = val
	v.isSet = true
}

func (v NullableReachableLocationsId) IsSet() bool {
	return v.isSet
}

func (v *NullableReachableLocationsId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReachableLocationsId(val *ReachableLocationsId) *NullableReachableLocationsId {
	return &NullableReachableLocationsId{value: val, isSet: true}
}

func (v NullableReachableLocationsId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReachableLocationsId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


