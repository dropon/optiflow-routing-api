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

// checks if the ManipulateRouteWaypoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManipulateRouteWaypoint{}

// ManipulateRouteWaypoint A _manipulate-route waypoint_ will not actually be reached but it influences the route course, so that the route passes an area defined by the given radius. This waypoint will not appear as a waypoint event in the response and may not be used as start and destination.
type ManipulateRouteWaypoint struct {
	// The latitude value in degrees (WGS84/EPSG:4326) from south to north.
	Latitude float64 `json:"latitude"`
	// The longitude value in degrees (WGS84/EPSG:4326) from west to east.
	Longitude float64 `json:"longitude"`
	// The radius [m] at which the waypoint has to be passed.
	Radius int32 `json:"radius"`
}

type _ManipulateRouteWaypoint ManipulateRouteWaypoint

// NewManipulateRouteWaypoint instantiates a new ManipulateRouteWaypoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManipulateRouteWaypoint(latitude float64, longitude float64, radius int32) *ManipulateRouteWaypoint {
	this := ManipulateRouteWaypoint{}
	this.Latitude = latitude
	this.Longitude = longitude
	this.Radius = radius
	return &this
}

// NewManipulateRouteWaypointWithDefaults instantiates a new ManipulateRouteWaypoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManipulateRouteWaypointWithDefaults() *ManipulateRouteWaypoint {
	this := ManipulateRouteWaypoint{}
	return &this
}

// GetLatitude returns the Latitude field value
func (o *ManipulateRouteWaypoint) GetLatitude() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value
// and a boolean to check if the value has been set.
func (o *ManipulateRouteWaypoint) GetLatitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Latitude, true
}

// SetLatitude sets field value
func (o *ManipulateRouteWaypoint) SetLatitude(v float64) {
	o.Latitude = v
}

// GetLongitude returns the Longitude field value
func (o *ManipulateRouteWaypoint) GetLongitude() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value
// and a boolean to check if the value has been set.
func (o *ManipulateRouteWaypoint) GetLongitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Longitude, true
}

// SetLongitude sets field value
func (o *ManipulateRouteWaypoint) SetLongitude(v float64) {
	o.Longitude = v
}

// GetRadius returns the Radius field value
func (o *ManipulateRouteWaypoint) GetRadius() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Radius
}

// GetRadiusOk returns a tuple with the Radius field value
// and a boolean to check if the value has been set.
func (o *ManipulateRouteWaypoint) GetRadiusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Radius, true
}

// SetRadius sets field value
func (o *ManipulateRouteWaypoint) SetRadius(v int32) {
	o.Radius = v
}

func (o ManipulateRouteWaypoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManipulateRouteWaypoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["latitude"] = o.Latitude
	toSerialize["longitude"] = o.Longitude
	toSerialize["radius"] = o.Radius
	return toSerialize, nil
}

func (o *ManipulateRouteWaypoint) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"latitude",
		"longitude",
		"radius",
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

	varManipulateRouteWaypoint := _ManipulateRouteWaypoint{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varManipulateRouteWaypoint)

	if err != nil {
		return err
	}

	*o = ManipulateRouteWaypoint(varManipulateRouteWaypoint)

	return err
}

type NullableManipulateRouteWaypoint struct {
	value *ManipulateRouteWaypoint
	isSet bool
}

func (v NullableManipulateRouteWaypoint) Get() *ManipulateRouteWaypoint {
	return v.value
}

func (v *NullableManipulateRouteWaypoint) Set(val *ManipulateRouteWaypoint) {
	v.value = val
	v.isSet = true
}

func (v NullableManipulateRouteWaypoint) IsSet() bool {
	return v.isSet
}

func (v *NullableManipulateRouteWaypoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManipulateRouteWaypoint(val *ManipulateRouteWaypoint) *NullableManipulateRouteWaypoint {
	return &NullableManipulateRouteWaypoint{value: val, isSet: true}
}

func (v NullableManipulateRouteWaypoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManipulateRouteWaypoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


