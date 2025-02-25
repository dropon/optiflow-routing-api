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

// checks if the ViolationEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ViolationEvent{}

// ViolationEvent Issued when the route passes a road which is prohibited for the given vehicle but passed nevertheless as there is no other valid route or when schedule restrictions are violated. Requires _VIOLATION_EVENTS_ to be requested. See [here](./concepts/violations) for more information.
type ViolationEvent struct {
	Type ViolationType `json:"type"`
	// Contains the list of violated schedule restrictions at this event. Only present if **type=SCHEDULE**. 
	ScheduleViolationTypes []ScheduleViolationType `json:"scheduleViolationTypes,omitempty"`
	// Contains the list of violated vehicle properties at this event. Only present if **type=VEHICLE_PROPERTY**. 
	ViolatedVehicleProperties []ViolatedVehicleProperty `json:"violatedVehicleProperties,omitempty"`
	// States whether the cause of the violation is only temporary, e.g. a road construction site. In such cases  see the corresponding traffic event for more information. 
	Temporary *bool `json:"temporary,omitempty"`
	// Returns the time validity of the violated restriction in the GDF time domain format if the restriction is only active at certain times. See [here](./concepts/date-and-time#gdftimedomains) for more information.
	TimeDomain *string `json:"timeDomain,omitempty"`
	AccessType *AccessType `json:"accessType,omitempty"`
	// For **accessType** _ENTER_ this index points to the corresponding event with **accessType** _EXIT_ and vice-versa. Not present otherwise.
	RelatedEventIndex *int32 `json:"relatedEventIndex,omitempty"`
	// The polyline of the violation in the format specified by **options[polylineFormat]**. Only present for **accessType** _ENTER_. Requires _VIOLATION_EVENTS_POLYLINE_ to be requested.
	Polyline *string `json:"polyline,omitempty"`
}

type _ViolationEvent ViolationEvent

// NewViolationEvent instantiates a new ViolationEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViolationEvent(type_ ViolationType) *ViolationEvent {
	this := ViolationEvent{}
	this.Type = type_
	return &this
}

// NewViolationEventWithDefaults instantiates a new ViolationEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViolationEventWithDefaults() *ViolationEvent {
	this := ViolationEvent{}
	return &this
}

// GetType returns the Type field value
func (o *ViolationEvent) GetType() ViolationType {
	if o == nil {
		var ret ViolationType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ViolationEvent) GetTypeOk() (*ViolationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ViolationEvent) SetType(v ViolationType) {
	o.Type = v
}

// GetScheduleViolationTypes returns the ScheduleViolationTypes field value if set, zero value otherwise.
func (o *ViolationEvent) GetScheduleViolationTypes() []ScheduleViolationType {
	if o == nil || IsNil(o.ScheduleViolationTypes) {
		var ret []ScheduleViolationType
		return ret
	}
	return o.ScheduleViolationTypes
}

// GetScheduleViolationTypesOk returns a tuple with the ScheduleViolationTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolationEvent) GetScheduleViolationTypesOk() ([]ScheduleViolationType, bool) {
	if o == nil || IsNil(o.ScheduleViolationTypes) {
		return nil, false
	}
	return o.ScheduleViolationTypes, true
}

// HasScheduleViolationTypes returns a boolean if a field has been set.
func (o *ViolationEvent) HasScheduleViolationTypes() bool {
	if o != nil && !IsNil(o.ScheduleViolationTypes) {
		return true
	}

	return false
}

// SetScheduleViolationTypes gets a reference to the given []ScheduleViolationType and assigns it to the ScheduleViolationTypes field.
func (o *ViolationEvent) SetScheduleViolationTypes(v []ScheduleViolationType) {
	o.ScheduleViolationTypes = v
}

// GetViolatedVehicleProperties returns the ViolatedVehicleProperties field value if set, zero value otherwise.
func (o *ViolationEvent) GetViolatedVehicleProperties() []ViolatedVehicleProperty {
	if o == nil || IsNil(o.ViolatedVehicleProperties) {
		var ret []ViolatedVehicleProperty
		return ret
	}
	return o.ViolatedVehicleProperties
}

// GetViolatedVehiclePropertiesOk returns a tuple with the ViolatedVehicleProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolationEvent) GetViolatedVehiclePropertiesOk() ([]ViolatedVehicleProperty, bool) {
	if o == nil || IsNil(o.ViolatedVehicleProperties) {
		return nil, false
	}
	return o.ViolatedVehicleProperties, true
}

// HasViolatedVehicleProperties returns a boolean if a field has been set.
func (o *ViolationEvent) HasViolatedVehicleProperties() bool {
	if o != nil && !IsNil(o.ViolatedVehicleProperties) {
		return true
	}

	return false
}

// SetViolatedVehicleProperties gets a reference to the given []ViolatedVehicleProperty and assigns it to the ViolatedVehicleProperties field.
func (o *ViolationEvent) SetViolatedVehicleProperties(v []ViolatedVehicleProperty) {
	o.ViolatedVehicleProperties = v
}

// GetTemporary returns the Temporary field value if set, zero value otherwise.
func (o *ViolationEvent) GetTemporary() bool {
	if o == nil || IsNil(o.Temporary) {
		var ret bool
		return ret
	}
	return *o.Temporary
}

// GetTemporaryOk returns a tuple with the Temporary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolationEvent) GetTemporaryOk() (*bool, bool) {
	if o == nil || IsNil(o.Temporary) {
		return nil, false
	}
	return o.Temporary, true
}

// HasTemporary returns a boolean if a field has been set.
func (o *ViolationEvent) HasTemporary() bool {
	if o != nil && !IsNil(o.Temporary) {
		return true
	}

	return false
}

// SetTemporary gets a reference to the given bool and assigns it to the Temporary field.
func (o *ViolationEvent) SetTemporary(v bool) {
	o.Temporary = &v
}

// GetTimeDomain returns the TimeDomain field value if set, zero value otherwise.
func (o *ViolationEvent) GetTimeDomain() string {
	if o == nil || IsNil(o.TimeDomain) {
		var ret string
		return ret
	}
	return *o.TimeDomain
}

// GetTimeDomainOk returns a tuple with the TimeDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolationEvent) GetTimeDomainOk() (*string, bool) {
	if o == nil || IsNil(o.TimeDomain) {
		return nil, false
	}
	return o.TimeDomain, true
}

// HasTimeDomain returns a boolean if a field has been set.
func (o *ViolationEvent) HasTimeDomain() bool {
	if o != nil && !IsNil(o.TimeDomain) {
		return true
	}

	return false
}

// SetTimeDomain gets a reference to the given string and assigns it to the TimeDomain field.
func (o *ViolationEvent) SetTimeDomain(v string) {
	o.TimeDomain = &v
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *ViolationEvent) GetAccessType() AccessType {
	if o == nil || IsNil(o.AccessType) {
		var ret AccessType
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolationEvent) GetAccessTypeOk() (*AccessType, bool) {
	if o == nil || IsNil(o.AccessType) {
		return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *ViolationEvent) HasAccessType() bool {
	if o != nil && !IsNil(o.AccessType) {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given AccessType and assigns it to the AccessType field.
func (o *ViolationEvent) SetAccessType(v AccessType) {
	o.AccessType = &v
}

// GetRelatedEventIndex returns the RelatedEventIndex field value if set, zero value otherwise.
func (o *ViolationEvent) GetRelatedEventIndex() int32 {
	if o == nil || IsNil(o.RelatedEventIndex) {
		var ret int32
		return ret
	}
	return *o.RelatedEventIndex
}

// GetRelatedEventIndexOk returns a tuple with the RelatedEventIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolationEvent) GetRelatedEventIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.RelatedEventIndex) {
		return nil, false
	}
	return o.RelatedEventIndex, true
}

// HasRelatedEventIndex returns a boolean if a field has been set.
func (o *ViolationEvent) HasRelatedEventIndex() bool {
	if o != nil && !IsNil(o.RelatedEventIndex) {
		return true
	}

	return false
}

// SetRelatedEventIndex gets a reference to the given int32 and assigns it to the RelatedEventIndex field.
func (o *ViolationEvent) SetRelatedEventIndex(v int32) {
	o.RelatedEventIndex = &v
}

// GetPolyline returns the Polyline field value if set, zero value otherwise.
func (o *ViolationEvent) GetPolyline() string {
	if o == nil || IsNil(o.Polyline) {
		var ret string
		return ret
	}
	return *o.Polyline
}

// GetPolylineOk returns a tuple with the Polyline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolationEvent) GetPolylineOk() (*string, bool) {
	if o == nil || IsNil(o.Polyline) {
		return nil, false
	}
	return o.Polyline, true
}

// HasPolyline returns a boolean if a field has been set.
func (o *ViolationEvent) HasPolyline() bool {
	if o != nil && !IsNil(o.Polyline) {
		return true
	}

	return false
}

// SetPolyline gets a reference to the given string and assigns it to the Polyline field.
func (o *ViolationEvent) SetPolyline(v string) {
	o.Polyline = &v
}

func (o ViolationEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ViolationEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.ScheduleViolationTypes) {
		toSerialize["scheduleViolationTypes"] = o.ScheduleViolationTypes
	}
	if !IsNil(o.ViolatedVehicleProperties) {
		toSerialize["violatedVehicleProperties"] = o.ViolatedVehicleProperties
	}
	if !IsNil(o.Temporary) {
		toSerialize["temporary"] = o.Temporary
	}
	if !IsNil(o.TimeDomain) {
		toSerialize["timeDomain"] = o.TimeDomain
	}
	if !IsNil(o.AccessType) {
		toSerialize["accessType"] = o.AccessType
	}
	if !IsNil(o.RelatedEventIndex) {
		toSerialize["relatedEventIndex"] = o.RelatedEventIndex
	}
	if !IsNil(o.Polyline) {
		toSerialize["polyline"] = o.Polyline
	}
	return toSerialize, nil
}

func (o *ViolationEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varViolationEvent := _ViolationEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varViolationEvent)

	if err != nil {
		return err
	}

	*o = ViolationEvent(varViolationEvent)

	return err
}

type NullableViolationEvent struct {
	value *ViolationEvent
	isSet bool
}

func (v NullableViolationEvent) Get() *ViolationEvent {
	return v.value
}

func (v *NullableViolationEvent) Set(val *ViolationEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableViolationEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableViolationEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViolationEvent(val *ViolationEvent) *NullableViolationEvent {
	return &NullableViolationEvent{value: val, isSet: true}
}

func (v NullableViolationEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViolationEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


