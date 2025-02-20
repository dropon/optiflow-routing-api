# AlternativeRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Distance** | **int32** | The distance of the alternative route [m]. | 
**TravelTime** | **int32** | The travel time for the alternative route [s]. | 
**TrafficDelay** | Pointer to **int32** | The total delay due to live traffic on this alternative route [s].  This value contains the sum of all traffic events on this alternative route and  will be non-zero only if **options[trafficMode]&#x3D;REALISTIC**. See [here](./concepts/traffic-modes) for more information. | [optional] 
**Violated** | **bool** | If the alternative route cannot be calculated for the given vehicle the resulting alternative route is marked as violated. | 
**Polyline** | Pointer to **string** | The polyline of the alternative route in the format specified by **options[polylineFormat]**. | [optional] 
**RouteId** | Pointer to **string** | The ID of the alternative route. | [optional] 

## Methods

### NewAlternativeRoute

`func NewAlternativeRoute(distance int32, travelTime int32, violated bool, ) *AlternativeRoute`

NewAlternativeRoute instantiates a new AlternativeRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlternativeRouteWithDefaults

`func NewAlternativeRouteWithDefaults() *AlternativeRoute`

NewAlternativeRouteWithDefaults instantiates a new AlternativeRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistance

`func (o *AlternativeRoute) GetDistance() int32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *AlternativeRoute) GetDistanceOk() (*int32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *AlternativeRoute) SetDistance(v int32)`

SetDistance sets Distance field to given value.


### GetTravelTime

`func (o *AlternativeRoute) GetTravelTime() int32`

GetTravelTime returns the TravelTime field if non-nil, zero value otherwise.

### GetTravelTimeOk

`func (o *AlternativeRoute) GetTravelTimeOk() (*int32, bool)`

GetTravelTimeOk returns a tuple with the TravelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelTime

`func (o *AlternativeRoute) SetTravelTime(v int32)`

SetTravelTime sets TravelTime field to given value.


### GetTrafficDelay

`func (o *AlternativeRoute) GetTrafficDelay() int32`

GetTrafficDelay returns the TrafficDelay field if non-nil, zero value otherwise.

### GetTrafficDelayOk

`func (o *AlternativeRoute) GetTrafficDelayOk() (*int32, bool)`

GetTrafficDelayOk returns a tuple with the TrafficDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficDelay

`func (o *AlternativeRoute) SetTrafficDelay(v int32)`

SetTrafficDelay sets TrafficDelay field to given value.

### HasTrafficDelay

`func (o *AlternativeRoute) HasTrafficDelay() bool`

HasTrafficDelay returns a boolean if a field has been set.

### GetViolated

`func (o *AlternativeRoute) GetViolated() bool`

GetViolated returns the Violated field if non-nil, zero value otherwise.

### GetViolatedOk

`func (o *AlternativeRoute) GetViolatedOk() (*bool, bool)`

GetViolatedOk returns a tuple with the Violated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolated

`func (o *AlternativeRoute) SetViolated(v bool)`

SetViolated sets Violated field to given value.


### GetPolyline

`func (o *AlternativeRoute) GetPolyline() string`

GetPolyline returns the Polyline field if non-nil, zero value otherwise.

### GetPolylineOk

`func (o *AlternativeRoute) GetPolylineOk() (*string, bool)`

GetPolylineOk returns a tuple with the Polyline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolyline

`func (o *AlternativeRoute) SetPolyline(v string)`

SetPolyline sets Polyline field to given value.

### HasPolyline

`func (o *AlternativeRoute) HasPolyline() bool`

HasPolyline returns a boolean if a field has been set.

### GetRouteId

`func (o *AlternativeRoute) GetRouteId() string`

GetRouteId returns the RouteId field if non-nil, zero value otherwise.

### GetRouteIdOk

`func (o *AlternativeRoute) GetRouteIdOk() (*string, bool)`

GetRouteIdOk returns a tuple with the RouteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteId

`func (o *AlternativeRoute) SetRouteId(v string)`

SetRouteId sets RouteId field to given value.

### HasRouteId

`func (o *AlternativeRoute) HasRouteId() bool`

HasRouteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


