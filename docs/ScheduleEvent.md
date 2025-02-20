# ScheduleEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int32** | The duration [s]. | 
**ScheduleTypes** | [**[]ScheduleType**](ScheduleType.md) | Tells what happens at this position of the route.  | 

## Methods

### NewScheduleEvent

`func NewScheduleEvent(duration int32, scheduleTypes []ScheduleType, ) *ScheduleEvent`

NewScheduleEvent instantiates a new ScheduleEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleEventWithDefaults

`func NewScheduleEventWithDefaults() *ScheduleEvent`

NewScheduleEventWithDefaults instantiates a new ScheduleEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *ScheduleEvent) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ScheduleEvent) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ScheduleEvent) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetScheduleTypes

`func (o *ScheduleEvent) GetScheduleTypes() []ScheduleType`

GetScheduleTypes returns the ScheduleTypes field if non-nil, zero value otherwise.

### GetScheduleTypesOk

`func (o *ScheduleEvent) GetScheduleTypesOk() (*[]ScheduleType, bool)`

GetScheduleTypesOk returns a tuple with the ScheduleTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleTypes

`func (o *ScheduleEvent) SetScheduleTypes(v []ScheduleType)`

SetScheduleTypes sets ScheduleTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


