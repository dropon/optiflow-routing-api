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

// checks if the RouteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RouteRequest{}

// RouteRequest struct for RouteRequest
type RouteRequest struct {
	// The list of waypoints the route will be calculated for. At least two waypoints are necessary, a maximum number may apply according to your subscription. The first waypoint is the start and the last is the destination of the route. Additional intermediate waypoints are possible.  Each waypoint must either have latitude and longitude or one of the representations combinedTransport, address or place.
	Waypoints []Waypoint `json:"waypoints,omitempty"`
	// Instead of the list of waypoints, a **routeId** from a previously calculated route or a matched track can be entered. See [here](./concepts/waypoints) for more information.
	RouteId *string `json:"routeId,omitempty"`
	Driver *DriverBody `json:"driver,omitempty"`
}

// NewRouteRequest instantiates a new RouteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouteRequest() *RouteRequest {
	this := RouteRequest{}
	return &this
}

// NewRouteRequestWithDefaults instantiates a new RouteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteRequestWithDefaults() *RouteRequest {
	this := RouteRequest{}
	return &this
}

// GetWaypoints returns the Waypoints field value if set, zero value otherwise.
func (o *RouteRequest) GetWaypoints() []Waypoint {
	if o == nil || IsNil(o.Waypoints) {
		var ret []Waypoint
		return ret
	}
	return o.Waypoints
}

// GetWaypointsOk returns a tuple with the Waypoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteRequest) GetWaypointsOk() ([]Waypoint, bool) {
	if o == nil || IsNil(o.Waypoints) {
		return nil, false
	}
	return o.Waypoints, true
}

// HasWaypoints returns a boolean if a field has been set.
func (o *RouteRequest) HasWaypoints() bool {
	if o != nil && !IsNil(o.Waypoints) {
		return true
	}

	return false
}

// SetWaypoints gets a reference to the given []Waypoint and assigns it to the Waypoints field.
func (o *RouteRequest) SetWaypoints(v []Waypoint) {
	o.Waypoints = v
}

// GetRouteId returns the RouteId field value if set, zero value otherwise.
func (o *RouteRequest) GetRouteId() string {
	if o == nil || IsNil(o.RouteId) {
		var ret string
		return ret
	}
	return *o.RouteId
}

// GetRouteIdOk returns a tuple with the RouteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteRequest) GetRouteIdOk() (*string, bool) {
	if o == nil || IsNil(o.RouteId) {
		return nil, false
	}
	return o.RouteId, true
}

// HasRouteId returns a boolean if a field has been set.
func (o *RouteRequest) HasRouteId() bool {
	if o != nil && !IsNil(o.RouteId) {
		return true
	}

	return false
}

// SetRouteId gets a reference to the given string and assigns it to the RouteId field.
func (o *RouteRequest) SetRouteId(v string) {
	o.RouteId = &v
}

// GetDriver returns the Driver field value if set, zero value otherwise.
func (o *RouteRequest) GetDriver() DriverBody {
	if o == nil || IsNil(o.Driver) {
		var ret DriverBody
		return ret
	}
	return *o.Driver
}

// GetDriverOk returns a tuple with the Driver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteRequest) GetDriverOk() (*DriverBody, bool) {
	if o == nil || IsNil(o.Driver) {
		return nil, false
	}
	return o.Driver, true
}

// HasDriver returns a boolean if a field has been set.
func (o *RouteRequest) HasDriver() bool {
	if o != nil && !IsNil(o.Driver) {
		return true
	}

	return false
}

// SetDriver gets a reference to the given DriverBody and assigns it to the Driver field.
func (o *RouteRequest) SetDriver(v DriverBody) {
	o.Driver = &v
}

func (o RouteRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RouteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Waypoints) {
		toSerialize["waypoints"] = o.Waypoints
	}
	if !IsNil(o.RouteId) {
		toSerialize["routeId"] = o.RouteId
	}
	if !IsNil(o.Driver) {
		toSerialize["driver"] = o.Driver
	}
	return toSerialize, nil
}

type NullableRouteRequest struct {
	value *RouteRequest
	isSet bool
}

func (v NullableRouteRequest) Get() *RouteRequest {
	return v.value
}

func (v *NullableRouteRequest) Set(val *RouteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteRequest(val *RouteRequest) *NullableRouteRequest {
	return &NullableRouteRequest{value: val, isSet: true}
}

func (v NullableRouteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


