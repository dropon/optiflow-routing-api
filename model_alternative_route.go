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

// checks if the AlternativeRoute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlternativeRoute{}

// AlternativeRoute struct for AlternativeRoute
type AlternativeRoute struct {
	// The distance of the alternative route [m].
	Distance int32 `json:"distance"`
	// The travel time for the alternative route [s].
	TravelTime int32 `json:"travelTime"`
	// The total delay due to live traffic on this alternative route [s].  This value contains the sum of all traffic events on this alternative route and  will be non-zero only if **options[trafficMode]=REALISTIC**. See [here](./concepts/traffic-modes) for more information.
	TrafficDelay *int32 `json:"trafficDelay,omitempty"`
	// If the alternative route cannot be calculated for the given vehicle the resulting alternative route is marked as violated.
	Violated bool `json:"violated"`
	// The polyline of the alternative route in the format specified by **options[polylineFormat]**.
	Polyline *string `json:"polyline,omitempty"`
	// The ID of the alternative route.
	RouteId *string `json:"routeId,omitempty"`
}

type _AlternativeRoute AlternativeRoute

// NewAlternativeRoute instantiates a new AlternativeRoute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlternativeRoute(distance int32, travelTime int32, violated bool) *AlternativeRoute {
	this := AlternativeRoute{}
	this.Distance = distance
	this.TravelTime = travelTime
	this.Violated = violated
	return &this
}

// NewAlternativeRouteWithDefaults instantiates a new AlternativeRoute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlternativeRouteWithDefaults() *AlternativeRoute {
	this := AlternativeRoute{}
	return &this
}

// GetDistance returns the Distance field value
func (o *AlternativeRoute) GetDistance() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Distance
}

// GetDistanceOk returns a tuple with the Distance field value
// and a boolean to check if the value has been set.
func (o *AlternativeRoute) GetDistanceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Distance, true
}

// SetDistance sets field value
func (o *AlternativeRoute) SetDistance(v int32) {
	o.Distance = v
}

// GetTravelTime returns the TravelTime field value
func (o *AlternativeRoute) GetTravelTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TravelTime
}

// GetTravelTimeOk returns a tuple with the TravelTime field value
// and a boolean to check if the value has been set.
func (o *AlternativeRoute) GetTravelTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TravelTime, true
}

// SetTravelTime sets field value
func (o *AlternativeRoute) SetTravelTime(v int32) {
	o.TravelTime = v
}

// GetTrafficDelay returns the TrafficDelay field value if set, zero value otherwise.
func (o *AlternativeRoute) GetTrafficDelay() int32 {
	if o == nil || IsNil(o.TrafficDelay) {
		var ret int32
		return ret
	}
	return *o.TrafficDelay
}

// GetTrafficDelayOk returns a tuple with the TrafficDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlternativeRoute) GetTrafficDelayOk() (*int32, bool) {
	if o == nil || IsNil(o.TrafficDelay) {
		return nil, false
	}
	return o.TrafficDelay, true
}

// HasTrafficDelay returns a boolean if a field has been set.
func (o *AlternativeRoute) HasTrafficDelay() bool {
	if o != nil && !IsNil(o.TrafficDelay) {
		return true
	}

	return false
}

// SetTrafficDelay gets a reference to the given int32 and assigns it to the TrafficDelay field.
func (o *AlternativeRoute) SetTrafficDelay(v int32) {
	o.TrafficDelay = &v
}

// GetViolated returns the Violated field value
func (o *AlternativeRoute) GetViolated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Violated
}

// GetViolatedOk returns a tuple with the Violated field value
// and a boolean to check if the value has been set.
func (o *AlternativeRoute) GetViolatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Violated, true
}

// SetViolated sets field value
func (o *AlternativeRoute) SetViolated(v bool) {
	o.Violated = v
}

// GetPolyline returns the Polyline field value if set, zero value otherwise.
func (o *AlternativeRoute) GetPolyline() string {
	if o == nil || IsNil(o.Polyline) {
		var ret string
		return ret
	}
	return *o.Polyline
}

// GetPolylineOk returns a tuple with the Polyline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlternativeRoute) GetPolylineOk() (*string, bool) {
	if o == nil || IsNil(o.Polyline) {
		return nil, false
	}
	return o.Polyline, true
}

// HasPolyline returns a boolean if a field has been set.
func (o *AlternativeRoute) HasPolyline() bool {
	if o != nil && !IsNil(o.Polyline) {
		return true
	}

	return false
}

// SetPolyline gets a reference to the given string and assigns it to the Polyline field.
func (o *AlternativeRoute) SetPolyline(v string) {
	o.Polyline = &v
}

// GetRouteId returns the RouteId field value if set, zero value otherwise.
func (o *AlternativeRoute) GetRouteId() string {
	if o == nil || IsNil(o.RouteId) {
		var ret string
		return ret
	}
	return *o.RouteId
}

// GetRouteIdOk returns a tuple with the RouteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlternativeRoute) GetRouteIdOk() (*string, bool) {
	if o == nil || IsNil(o.RouteId) {
		return nil, false
	}
	return o.RouteId, true
}

// HasRouteId returns a boolean if a field has been set.
func (o *AlternativeRoute) HasRouteId() bool {
	if o != nil && !IsNil(o.RouteId) {
		return true
	}

	return false
}

// SetRouteId gets a reference to the given string and assigns it to the RouteId field.
func (o *AlternativeRoute) SetRouteId(v string) {
	o.RouteId = &v
}

func (o AlternativeRoute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlternativeRoute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["distance"] = o.Distance
	toSerialize["travelTime"] = o.TravelTime
	if !IsNil(o.TrafficDelay) {
		toSerialize["trafficDelay"] = o.TrafficDelay
	}
	toSerialize["violated"] = o.Violated
	if !IsNil(o.Polyline) {
		toSerialize["polyline"] = o.Polyline
	}
	if !IsNil(o.RouteId) {
		toSerialize["routeId"] = o.RouteId
	}
	return toSerialize, nil
}

func (o *AlternativeRoute) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"distance",
		"travelTime",
		"violated",
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

	varAlternativeRoute := _AlternativeRoute{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAlternativeRoute)

	if err != nil {
		return err
	}

	*o = AlternativeRoute(varAlternativeRoute)

	return err
}

type NullableAlternativeRoute struct {
	value *AlternativeRoute
	isSet bool
}

func (v NullableAlternativeRoute) Get() *AlternativeRoute {
	return v.value
}

func (v *NullableAlternativeRoute) Set(val *AlternativeRoute) {
	v.value = val
	v.isSet = true
}

func (v NullableAlternativeRoute) IsSet() bool {
	return v.isSet
}

func (v *NullableAlternativeRoute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlternativeRoute(val *AlternativeRoute) *NullableAlternativeRoute {
	return &NullableAlternativeRoute{value: val, isSet: true}
}

func (v NullableAlternativeRoute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlternativeRoute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


