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

// checks if the ReachableAreas type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReachableAreas{}

// ReachableAreas The result of the reachable areas calculation.
type ReachableAreas struct {
	// The list of polygons calculated for the specified horizons in GeoJson format. For each horizon there is a separate polygon at the same index.
	Polygons []string `json:"polygons"`
	// A list of warnings concerning the validity of the result.
	Warnings []Warning `json:"warnings,omitempty"`
}

type _ReachableAreas ReachableAreas

// NewReachableAreas instantiates a new ReachableAreas object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReachableAreas(polygons []string) *ReachableAreas {
	this := ReachableAreas{}
	this.Polygons = polygons
	return &this
}

// NewReachableAreasWithDefaults instantiates a new ReachableAreas object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReachableAreasWithDefaults() *ReachableAreas {
	this := ReachableAreas{}
	return &this
}

// GetPolygons returns the Polygons field value
func (o *ReachableAreas) GetPolygons() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Polygons
}

// GetPolygonsOk returns a tuple with the Polygons field value
// and a boolean to check if the value has been set.
func (o *ReachableAreas) GetPolygonsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Polygons, true
}

// SetPolygons sets field value
func (o *ReachableAreas) SetPolygons(v []string) {
	o.Polygons = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ReachableAreas) GetWarnings() []Warning {
	if o == nil || IsNil(o.Warnings) {
		var ret []Warning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReachableAreas) GetWarningsOk() ([]Warning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ReachableAreas) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []Warning and assigns it to the Warnings field.
func (o *ReachableAreas) SetWarnings(v []Warning) {
	o.Warnings = v
}

func (o ReachableAreas) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReachableAreas) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["polygons"] = o.Polygons
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *ReachableAreas) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"polygons",
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

	varReachableAreas := _ReachableAreas{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReachableAreas)

	if err != nil {
		return err
	}

	*o = ReachableAreas(varReachableAreas)

	return err
}

type NullableReachableAreas struct {
	value *ReachableAreas
	isSet bool
}

func (v NullableReachableAreas) Get() *ReachableAreas {
	return v.value
}

func (v *NullableReachableAreas) Set(val *ReachableAreas) {
	v.value = val
	v.isSet = true
}

func (v NullableReachableAreas) IsSet() bool {
	return v.isSet
}

func (v *NullableReachableAreas) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReachableAreas(val *ReachableAreas) *NullableReachableAreas {
	return &NullableReachableAreas{value: val, isSet: true}
}

func (v NullableReachableAreas) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReachableAreas) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


