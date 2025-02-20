# PositionAtWaypoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the waypoint. | 
**PerformedServiceTime** | Pointer to **int32** | The service time [s] which has already been performed. | [optional] [default to 0]

## Methods

### NewPositionAtWaypoint

`func NewPositionAtWaypoint(name string, ) *PositionAtWaypoint`

NewPositionAtWaypoint instantiates a new PositionAtWaypoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPositionAtWaypointWithDefaults

`func NewPositionAtWaypointWithDefaults() *PositionAtWaypoint`

NewPositionAtWaypointWithDefaults instantiates a new PositionAtWaypoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PositionAtWaypoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PositionAtWaypoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PositionAtWaypoint) SetName(v string)`

SetName sets Name field to given value.


### GetPerformedServiceTime

`func (o *PositionAtWaypoint) GetPerformedServiceTime() int32`

GetPerformedServiceTime returns the PerformedServiceTime field if non-nil, zero value otherwise.

### GetPerformedServiceTimeOk

`func (o *PositionAtWaypoint) GetPerformedServiceTimeOk() (*int32, bool)`

GetPerformedServiceTimeOk returns a tuple with the PerformedServiceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformedServiceTime

`func (o *PositionAtWaypoint) SetPerformedServiceTime(v int32)`

SetPerformedServiceTime sets PerformedServiceTime field to given value.

### HasPerformedServiceTime

`func (o *PositionAtWaypoint) HasPerformedServiceTime() bool`

HasPerformedServiceTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


