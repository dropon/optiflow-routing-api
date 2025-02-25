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

// checks if the ManeuverEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManeuverEvent{}

// ManeuverEvent Issued when the driver has to perform a maneuver, e.g. to turn or to enter a roundabout. Requires _MANEUVER_EVENTS_ to be requested.
type ManeuverEvent struct {
	Type ManeuverType `json:"type"`
	// The direction of the outgoing road relative to continuing in the same direction as the incoming road (clockwise).
	RelativeDirection *int32 `json:"relativeDirection,omitempty"`
	// The absolute direction of the outgoing road (clockwise). North represents 0 degrees.
	AbsoluteDirection *int32 `json:"absoluteDirection,omitempty"`
	// A descriptive text for the current maneuver. The language can be specified by the parameter **options[language]**. A warning with **warningCode** _ROUTING_MANEUVERS_IN_DIFFERENT_LANGUAGE_ and the actual language is returned when the requested language is not available. Geographical names such as town and road names are always given in the local language.
	Description string `json:"description"`
	RoadAhead *RoadAhead `json:"roadAhead,omitempty"`
	// The city names and road numbers on a signpost at the current location to follow for the current maneuver. Empty if no signpost is present or the data is not available.
	DirectionSignText *string `json:"directionSignText,omitempty"`
	// The number of an exit or interchange of a highway or a freeway-like road. Only present if the maneuver type is _CHANGE_ or _EXIT_. Empty if the data does not contain an exit number.
	ExitNumber *string `json:"exitNumber,omitempty"`
	// The name of an exit or interchange of a highway or a freeway-like road. Only present if the maneuver type is _CHANGE_ or _EXIT_. Empty if the data does not contain an exit name.
	ExitName *string `json:"exitName,omitempty"`
	// The exit number at a roundabout. Only drivable roads are counted. Only present if the maneuver type is _TAKE\\_ROUNDABOUT_.
	RoundaboutExit *int32 `json:"roundaboutExit,omitempty"`
	// The name of the combined transport to take a the current location. Only present if the maneuver type is _TAKE\\_COMBINED\\_TRANSPORT_.
	CombinedTransportName *string `json:"combinedTransportName,omitempty"`
	CombinedTransportType *CombinedTransportType `json:"combinedTransportType,omitempty"`
	// The name of the crossing road at which a U-turn has to be made. Only present if the maneuver type is _MAKE\\_U\\_TURN_ and if the U-turn takes place at a crossing.
	CrossingRoadName *string `json:"crossingRoadName,omitempty"`
}

type _ManeuverEvent ManeuverEvent

// NewManeuverEvent instantiates a new ManeuverEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManeuverEvent(type_ ManeuverType, description string) *ManeuverEvent {
	this := ManeuverEvent{}
	this.Type = type_
	this.Description = description
	return &this
}

// NewManeuverEventWithDefaults instantiates a new ManeuverEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManeuverEventWithDefaults() *ManeuverEvent {
	this := ManeuverEvent{}
	return &this
}

// GetType returns the Type field value
func (o *ManeuverEvent) GetType() ManeuverType {
	if o == nil {
		var ret ManeuverType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ManeuverEvent) GetTypeOk() (*ManeuverType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ManeuverEvent) SetType(v ManeuverType) {
	o.Type = v
}

// GetRelativeDirection returns the RelativeDirection field value if set, zero value otherwise.
func (o *ManeuverEvent) GetRelativeDirection() int32 {
	if o == nil || IsNil(o.RelativeDirection) {
		var ret int32
		return ret
	}
	return *o.RelativeDirection
}

// GetRelativeDirectionOk returns a tuple with the RelativeDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManeuverEvent) GetRelativeDirectionOk() (*int32, bool) {
	if o == nil || IsNil(o.RelativeDirection) {
		return nil, false
	}
	return o.RelativeDirection, true
}

// HasRelativeDirection returns a boolean if a field has been set.
func (o *ManeuverEvent) HasRelativeDirection() bool {
	if o != nil && !IsNil(o.RelativeDirection) {
		return true
	}

	return false
}

// SetRelativeDirection gets a reference to the given int32 and assigns it to the RelativeDirection field.
func (o *ManeuverEvent) SetRelativeDirection(v int32) {
	o.RelativeDirection = &v
}

// GetAbsoluteDirection returns the AbsoluteDirection field value if set, zero value otherwise.
func (o *ManeuverEvent) GetAbsoluteDirection() int32 {
	if o == nil || IsNil(o.AbsoluteDirection) {
		var ret int32
		return ret
	}
	return *o.AbsoluteDirection
}

// GetAbsoluteDirectionOk returns a tuple with the AbsoluteDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManeuverEvent) GetAbsoluteDirectionOk() (*int32, bool) {
	if o == nil || IsNil(o.AbsoluteDirection) {
		return nil, false
	}
	return o.AbsoluteDirection, true
}

// HasAbsoluteDirection returns a boolean if a field has been set.
func (o *ManeuverEvent) HasAbsoluteDirection() bool {
	if o != nil && !IsNil(o.AbsoluteDirection) {
		return true
	}

	return false
}

// SetAbsoluteDirection gets a reference to the given int32 and assigns it to the AbsoluteDirection field.
func (o *ManeuverEvent) SetAbsoluteDirection(v int32) {
	o.AbsoluteDirection = &v
}

// GetDescription returns the Description field value
func (o *ManeuverEvent) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ManeuverEvent) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ManeuverEvent) SetDescription(v string) {
	o.Description = v
}

// GetRoadAhead returns the RoadAhead field value if set, zero value otherwise.
func (o *ManeuverEvent) GetRoadAhead() RoadAhead {
	if o == nil || IsNil(o.RoadAhead) {
		var ret RoadAhead
		return ret
	}
	return *o.RoadAhead
}

// GetRoadAheadOk returns a tuple with the RoadAhead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManeuverEvent) GetRoadAheadOk() (*RoadAhead, bool) {
	if o == nil || IsNil(o.RoadAhead) {
		return nil, false
	}
	return o.RoadAhead, true
}

// HasRoadAhead returns a boolean if a field has been set.
func (o *ManeuverEvent) HasRoadAhead() bool {
	if o != nil && !IsNil(o.RoadAhead) {
		return true
	}

	return false
}

// SetRoadAhead gets a reference to the given RoadAhead and assigns it to the RoadAhead field.
func (o *ManeuverEvent) SetRoadAhead(v RoadAhead) {
	o.RoadAhead = &v
}

// GetDirectionSignText returns the DirectionSignText field value if set, zero value otherwise.
func (o *ManeuverEvent) GetDirectionSignText() string {
	if o == nil || IsNil(o.DirectionSignText) {
		var ret string
		return ret
	}
	return *o.DirectionSignText
}

// GetDirectionSignTextOk returns a tuple with the DirectionSignText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManeuverEvent) GetDirectionSignTextOk() (*string, bool) {
	if o == nil || IsNil(o.DirectionSignText) {
		return nil, false
	}
	return o.DirectionSignText, true
}

// HasDirectionSignText returns a boolean if a field has been set.
func (o *ManeuverEvent) HasDirectionSignText() bool {
	if o != nil && !IsNil(o.DirectionSignText) {
		return true
	}

	return false
}

// SetDirectionSignText gets a reference to the given string and assigns it to the DirectionSignText field.
func (o *ManeuverEvent) SetDirectionSignText(v string) {
	o.DirectionSignText = &v
}

// GetExitNumber returns the ExitNumber field value if set, zero value otherwise.
func (o *ManeuverEvent) GetExitNumber() string {
	if o == nil || IsNil(o.ExitNumber) {
		var ret string
		return ret
	}
	return *o.ExitNumber
}

// GetExitNumberOk returns a tuple with the ExitNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManeuverEvent) GetExitNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ExitNumber) {
		return nil, false
	}
	return o.ExitNumber, true
}

// HasExitNumber returns a boolean if a field has been set.
func (o *ManeuverEvent) HasExitNumber() bool {
	if o != nil && !IsNil(o.ExitNumber) {
		return true
	}

	return false
}

// SetExitNumber gets a reference to the given string and assigns it to the ExitNumber field.
func (o *ManeuverEvent) SetExitNumber(v string) {
	o.ExitNumber = &v
}

// GetExitName returns the ExitName field value if set, zero value otherwise.
func (o *ManeuverEvent) GetExitName() string {
	if o == nil || IsNil(o.ExitName) {
		var ret string
		return ret
	}
	return *o.ExitName
}

// GetExitNameOk returns a tuple with the ExitName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManeuverEvent) GetExitNameOk() (*string, bool) {
	if o == nil || IsNil(o.ExitName) {
		return nil, false
	}
	return o.ExitName, true
}

// HasExitName returns a boolean if a field has been set.
func (o *ManeuverEvent) HasExitName() bool {
	if o != nil && !IsNil(o.ExitName) {
		return true
	}

	return false
}

// SetExitName gets a reference to the given string and assigns it to the ExitName field.
func (o *ManeuverEvent) SetExitName(v string) {
	o.ExitName = &v
}

// GetRoundaboutExit returns the RoundaboutExit field value if set, zero value otherwise.
func (o *ManeuverEvent) GetRoundaboutExit() int32 {
	if o == nil || IsNil(o.RoundaboutExit) {
		var ret int32
		return ret
	}
	return *o.RoundaboutExit
}

// GetRoundaboutExitOk returns a tuple with the RoundaboutExit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManeuverEvent) GetRoundaboutExitOk() (*int32, bool) {
	if o == nil || IsNil(o.RoundaboutExit) {
		return nil, false
	}
	return o.RoundaboutExit, true
}

// HasRoundaboutExit returns a boolean if a field has been set.
func (o *ManeuverEvent) HasRoundaboutExit() bool {
	if o != nil && !IsNil(o.RoundaboutExit) {
		return true
	}

	return false
}

// SetRoundaboutExit gets a reference to the given int32 and assigns it to the RoundaboutExit field.
func (o *ManeuverEvent) SetRoundaboutExit(v int32) {
	o.RoundaboutExit = &v
}

// GetCombinedTransportName returns the CombinedTransportName field value if set, zero value otherwise.
func (o *ManeuverEvent) GetCombinedTransportName() string {
	if o == nil || IsNil(o.CombinedTransportName) {
		var ret string
		return ret
	}
	return *o.CombinedTransportName
}

// GetCombinedTransportNameOk returns a tuple with the CombinedTransportName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManeuverEvent) GetCombinedTransportNameOk() (*string, bool) {
	if o == nil || IsNil(o.CombinedTransportName) {
		return nil, false
	}
	return o.CombinedTransportName, true
}

// HasCombinedTransportName returns a boolean if a field has been set.
func (o *ManeuverEvent) HasCombinedTransportName() bool {
	if o != nil && !IsNil(o.CombinedTransportName) {
		return true
	}

	return false
}

// SetCombinedTransportName gets a reference to the given string and assigns it to the CombinedTransportName field.
func (o *ManeuverEvent) SetCombinedTransportName(v string) {
	o.CombinedTransportName = &v
}

// GetCombinedTransportType returns the CombinedTransportType field value if set, zero value otherwise.
func (o *ManeuverEvent) GetCombinedTransportType() CombinedTransportType {
	if o == nil || IsNil(o.CombinedTransportType) {
		var ret CombinedTransportType
		return ret
	}
	return *o.CombinedTransportType
}

// GetCombinedTransportTypeOk returns a tuple with the CombinedTransportType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManeuverEvent) GetCombinedTransportTypeOk() (*CombinedTransportType, bool) {
	if o == nil || IsNil(o.CombinedTransportType) {
		return nil, false
	}
	return o.CombinedTransportType, true
}

// HasCombinedTransportType returns a boolean if a field has been set.
func (o *ManeuverEvent) HasCombinedTransportType() bool {
	if o != nil && !IsNil(o.CombinedTransportType) {
		return true
	}

	return false
}

// SetCombinedTransportType gets a reference to the given CombinedTransportType and assigns it to the CombinedTransportType field.
func (o *ManeuverEvent) SetCombinedTransportType(v CombinedTransportType) {
	o.CombinedTransportType = &v
}

// GetCrossingRoadName returns the CrossingRoadName field value if set, zero value otherwise.
func (o *ManeuverEvent) GetCrossingRoadName() string {
	if o == nil || IsNil(o.CrossingRoadName) {
		var ret string
		return ret
	}
	return *o.CrossingRoadName
}

// GetCrossingRoadNameOk returns a tuple with the CrossingRoadName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManeuverEvent) GetCrossingRoadNameOk() (*string, bool) {
	if o == nil || IsNil(o.CrossingRoadName) {
		return nil, false
	}
	return o.CrossingRoadName, true
}

// HasCrossingRoadName returns a boolean if a field has been set.
func (o *ManeuverEvent) HasCrossingRoadName() bool {
	if o != nil && !IsNil(o.CrossingRoadName) {
		return true
	}

	return false
}

// SetCrossingRoadName gets a reference to the given string and assigns it to the CrossingRoadName field.
func (o *ManeuverEvent) SetCrossingRoadName(v string) {
	o.CrossingRoadName = &v
}

func (o ManeuverEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManeuverEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.RelativeDirection) {
		toSerialize["relativeDirection"] = o.RelativeDirection
	}
	if !IsNil(o.AbsoluteDirection) {
		toSerialize["absoluteDirection"] = o.AbsoluteDirection
	}
	toSerialize["description"] = o.Description
	if !IsNil(o.RoadAhead) {
		toSerialize["roadAhead"] = o.RoadAhead
	}
	if !IsNil(o.DirectionSignText) {
		toSerialize["directionSignText"] = o.DirectionSignText
	}
	if !IsNil(o.ExitNumber) {
		toSerialize["exitNumber"] = o.ExitNumber
	}
	if !IsNil(o.ExitName) {
		toSerialize["exitName"] = o.ExitName
	}
	if !IsNil(o.RoundaboutExit) {
		toSerialize["roundaboutExit"] = o.RoundaboutExit
	}
	if !IsNil(o.CombinedTransportName) {
		toSerialize["combinedTransportName"] = o.CombinedTransportName
	}
	if !IsNil(o.CombinedTransportType) {
		toSerialize["combinedTransportType"] = o.CombinedTransportType
	}
	if !IsNil(o.CrossingRoadName) {
		toSerialize["crossingRoadName"] = o.CrossingRoadName
	}
	return toSerialize, nil
}

func (o *ManeuverEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"description",
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

	varManeuverEvent := _ManeuverEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varManeuverEvent)

	if err != nil {
		return err
	}

	*o = ManeuverEvent(varManeuverEvent)

	return err
}

type NullableManeuverEvent struct {
	value *ManeuverEvent
	isSet bool
}

func (v NullableManeuverEvent) Get() *ManeuverEvent {
	return v.value
}

func (v *NullableManeuverEvent) Set(val *ManeuverEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableManeuverEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableManeuverEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManeuverEvent(val *ManeuverEvent) *NullableManeuverEvent {
	return &NullableManeuverEvent{value: val, isSet: true}
}

func (v NullableManeuverEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManeuverEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


