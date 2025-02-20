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

// checks if the Locations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Locations{}

// Locations struct for Locations
type Locations struct {
	Locations []Location `json:"locations"`
}

type _Locations Locations

// NewLocations instantiates a new Locations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocations(locations []Location) *Locations {
	this := Locations{}
	this.Locations = locations
	return &this
}

// NewLocationsWithDefaults instantiates a new Locations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationsWithDefaults() *Locations {
	this := Locations{}
	return &this
}

// GetLocations returns the Locations field value
func (o *Locations) GetLocations() []Location {
	if o == nil {
		var ret []Location
		return ret
	}

	return o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value
// and a boolean to check if the value has been set.
func (o *Locations) GetLocationsOk() ([]Location, bool) {
	if o == nil {
		return nil, false
	}
	return o.Locations, true
}

// SetLocations sets field value
func (o *Locations) SetLocations(v []Location) {
	o.Locations = v
}

func (o Locations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Locations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["locations"] = o.Locations
	return toSerialize, nil
}

func (o *Locations) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"locations",
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

	varLocations := _Locations{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLocations)

	if err != nil {
		return err
	}

	*o = Locations(varLocations)

	return err
}

type NullableLocations struct {
	value *Locations
	isSet bool
}

func (v NullableLocations) Get() *Locations {
	return v.value
}

func (v *NullableLocations) Set(val *Locations) {
	v.value = val
	v.isSet = true
}

func (v NullableLocations) IsSet() bool {
	return v.isSet
}

func (v *NullableLocations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocations(val *Locations) *NullableLocations {
	return &NullableLocations{value: val, isSet: true}
}

func (v NullableLocations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


