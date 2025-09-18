# Waypoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | An identifier to reference this waypoint in the response. | [optional] 
**OnRoad** | Pointer to [**OnRoadWaypoint**](OnRoadWaypoint.md) |  | [optional] 
**OffRoad** | Pointer to [**OffRoadWaypoint**](OffRoadWaypoint.md) |  | [optional] 
**Manipulate** | Pointer to [**ManipulateRouteWaypoint**](ManipulateRouteWaypoint.md) |  | [optional] 
**CombinedTransport** | Pointer to [**CombinedTransport**](CombinedTransport.md) |  | [optional] 

## Methods

### NewWaypoint

`func NewWaypoint() *Waypoint`

NewWaypoint instantiates a new Waypoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWaypointWithDefaults

`func NewWaypointWithDefaults() *Waypoint`

NewWaypointWithDefaults instantiates a new Waypoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Waypoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Waypoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Waypoint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Waypoint) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnRoad

`func (o *Waypoint) GetOnRoad() OnRoadWaypoint`

GetOnRoad returns the OnRoad field if non-nil, zero value otherwise.

### GetOnRoadOk

`func (o *Waypoint) GetOnRoadOk() (*OnRoadWaypoint, bool)`

GetOnRoadOk returns a tuple with the OnRoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnRoad

`func (o *Waypoint) SetOnRoad(v OnRoadWaypoint)`

SetOnRoad sets OnRoad field to given value.

### HasOnRoad

`func (o *Waypoint) HasOnRoad() bool`

HasOnRoad returns a boolean if a field has been set.

### GetOffRoad

`func (o *Waypoint) GetOffRoad() OffRoadWaypoint`

GetOffRoad returns the OffRoad field if non-nil, zero value otherwise.

### GetOffRoadOk

`func (o *Waypoint) GetOffRoadOk() (*OffRoadWaypoint, bool)`

GetOffRoadOk returns a tuple with the OffRoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffRoad

`func (o *Waypoint) SetOffRoad(v OffRoadWaypoint)`

SetOffRoad sets OffRoad field to given value.

### HasOffRoad

`func (o *Waypoint) HasOffRoad() bool`

HasOffRoad returns a boolean if a field has been set.

### GetManipulate

`func (o *Waypoint) GetManipulate() ManipulateRouteWaypoint`

GetManipulate returns the Manipulate field if non-nil, zero value otherwise.

### GetManipulateOk

`func (o *Waypoint) GetManipulateOk() (*ManipulateRouteWaypoint, bool)`

GetManipulateOk returns a tuple with the Manipulate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManipulate

`func (o *Waypoint) SetManipulate(v ManipulateRouteWaypoint)`

SetManipulate sets Manipulate field to given value.

### HasManipulate

`func (o *Waypoint) HasManipulate() bool`

HasManipulate returns a boolean if a field has been set.

### GetCombinedTransport

`func (o *Waypoint) GetCombinedTransport() CombinedTransport`

GetCombinedTransport returns the CombinedTransport field if non-nil, zero value otherwise.

### GetCombinedTransportOk

`func (o *Waypoint) GetCombinedTransportOk() (*CombinedTransport, bool)`

GetCombinedTransportOk returns a tuple with the CombinedTransport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedTransport

`func (o *Waypoint) SetCombinedTransport(v CombinedTransport)`

SetCombinedTransport sets CombinedTransport field to given value.

### HasCombinedTransport

`func (o *Waypoint) HasCombinedTransport() bool`

HasCombinedTransport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


