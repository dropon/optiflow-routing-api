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

// checks if the CombinedTransport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CombinedTransport{}

// CombinedTransport Influences the route course, so that the route uses a ferry or railway connection between the given locations. Both locations will be matched to the nearest ports looking for a direct connection. If no connection can be found, this waypoint will be ignored, and the warning _ROUTING_COMBINED_TRANSPORT_WAYPOINT_IGNORED_ will be returned. If more than one connection is found, the best one will be used, and the alternative connections will be returned in the response in a warning _ROUTING_COMBINED_TRANSPORT_WAYPOINT_AMBIGUOUS_. This waypoint will not appear as a waypoint event in the response and may not be used as start or destination. We will refer to this type of waypoint as a _combined-transport waypoint_.
type CombinedTransport struct {
	Start CombinedTransportLocation `json:"start"`
	Destination CombinedTransportLocation `json:"destination"`
}

type _CombinedTransport CombinedTransport

// NewCombinedTransport instantiates a new CombinedTransport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCombinedTransport(start CombinedTransportLocation, destination CombinedTransportLocation) *CombinedTransport {
	this := CombinedTransport{}
	this.Start = start
	this.Destination = destination
	return &this
}

// NewCombinedTransportWithDefaults instantiates a new CombinedTransport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCombinedTransportWithDefaults() *CombinedTransport {
	this := CombinedTransport{}
	return &this
}

// GetStart returns the Start field value
func (o *CombinedTransport) GetStart() CombinedTransportLocation {
	if o == nil {
		var ret CombinedTransportLocation
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *CombinedTransport) GetStartOk() (*CombinedTransportLocation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *CombinedTransport) SetStart(v CombinedTransportLocation) {
	o.Start = v
}

// GetDestination returns the Destination field value
func (o *CombinedTransport) GetDestination() CombinedTransportLocation {
	if o == nil {
		var ret CombinedTransportLocation
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *CombinedTransport) GetDestinationOk() (*CombinedTransportLocation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *CombinedTransport) SetDestination(v CombinedTransportLocation) {
	o.Destination = v
}

func (o CombinedTransport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CombinedTransport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["start"] = o.Start
	toSerialize["destination"] = o.Destination
	return toSerialize, nil
}

func (o *CombinedTransport) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"start",
		"destination",
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

	varCombinedTransport := _CombinedTransport{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCombinedTransport)

	if err != nil {
		return err
	}

	*o = CombinedTransport(varCombinedTransport)

	return err
}

type NullableCombinedTransport struct {
	value *CombinedTransport
	isSet bool
}

func (v NullableCombinedTransport) Get() *CombinedTransport {
	return v.value
}

func (v *NullableCombinedTransport) Set(val *CombinedTransport) {
	v.value = val
	v.isSet = true
}

func (v NullableCombinedTransport) IsSet() bool {
	return v.isSet
}

func (v *NullableCombinedTransport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCombinedTransport(val *CombinedTransport) *NullableCombinedTransport {
	return &NullableCombinedTransport{value: val, isSet: true}
}

func (v NullableCombinedTransport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCombinedTransport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


