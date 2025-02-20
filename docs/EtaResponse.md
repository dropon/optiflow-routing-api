# EtaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemainingWaypoints** | [**[]RemainingWaypoint**](RemainingWaypoint.md) | The ETA and the remaining distance and travel time to the waypoints not yet reached by the route. | 
**RouteId** | Pointer to **string** | If the route has been recalculated, this is the ID of the new route. That will happen when it is likely that the vehicle has left the route, so the vehicle position is too far away from the route, or when the estimated time of arrival at the last waypoint is delayed in a way that another route may result in arriving earlier. | [optional] 
**Warnings** | Pointer to [**[]Warning**](Warning.md) | A list of warnings concerning the validity of the result. | [optional] 

## Methods

### NewEtaResponse

`func NewEtaResponse(remainingWaypoints []RemainingWaypoint, ) *EtaResponse`

NewEtaResponse instantiates a new EtaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEtaResponseWithDefaults

`func NewEtaResponseWithDefaults() *EtaResponse`

NewEtaResponseWithDefaults instantiates a new EtaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemainingWaypoints

`func (o *EtaResponse) GetRemainingWaypoints() []RemainingWaypoint`

GetRemainingWaypoints returns the RemainingWaypoints field if non-nil, zero value otherwise.

### GetRemainingWaypointsOk

`func (o *EtaResponse) GetRemainingWaypointsOk() (*[]RemainingWaypoint, bool)`

GetRemainingWaypointsOk returns a tuple with the RemainingWaypoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingWaypoints

`func (o *EtaResponse) SetRemainingWaypoints(v []RemainingWaypoint)`

SetRemainingWaypoints sets RemainingWaypoints field to given value.


### GetRouteId

`func (o *EtaResponse) GetRouteId() string`

GetRouteId returns the RouteId field if non-nil, zero value otherwise.

### GetRouteIdOk

`func (o *EtaResponse) GetRouteIdOk() (*string, bool)`

GetRouteIdOk returns a tuple with the RouteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteId

`func (o *EtaResponse) SetRouteId(v string)`

SetRouteId sets RouteId field to given value.

### HasRouteId

`func (o *EtaResponse) HasRouteId() bool`

HasRouteId returns a boolean if a field has been set.

### GetWarnings

`func (o *EtaResponse) GetWarnings() []Warning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *EtaResponse) GetWarningsOk() (*[]Warning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *EtaResponse) SetWarnings(v []Warning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *EtaResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


