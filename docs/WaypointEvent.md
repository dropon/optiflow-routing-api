# WaypointEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** | The index of the waypoint in the request. | 
**Name** | Pointer to **string** | The name of the waypoint as given in the request. This parameter is only present if it is not empty. | [optional] 

## Methods

### NewWaypointEvent

`func NewWaypointEvent(index int32, ) *WaypointEvent`

NewWaypointEvent instantiates a new WaypointEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWaypointEventWithDefaults

`func NewWaypointEventWithDefaults() *WaypointEvent`

NewWaypointEventWithDefaults instantiates a new WaypointEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *WaypointEvent) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *WaypointEvent) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *WaypointEvent) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetName

`func (o *WaypointEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WaypointEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WaypointEvent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WaypointEvent) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


