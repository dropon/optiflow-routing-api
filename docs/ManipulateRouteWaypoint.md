# ManipulateRouteWaypoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Latitude** | **float64** | The latitude value in degrees (WGS84/EPSG:4326) from south to north. | 
**Longitude** | **float64** | The longitude value in degrees (WGS84/EPSG:4326) from west to east. | 
**Radius** | **int32** | The radius [m] at which the waypoint has to be passed. | 

## Methods

### NewManipulateRouteWaypoint

`func NewManipulateRouteWaypoint(latitude float64, longitude float64, radius int32, ) *ManipulateRouteWaypoint`

NewManipulateRouteWaypoint instantiates a new ManipulateRouteWaypoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManipulateRouteWaypointWithDefaults

`func NewManipulateRouteWaypointWithDefaults() *ManipulateRouteWaypoint`

NewManipulateRouteWaypointWithDefaults instantiates a new ManipulateRouteWaypoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLatitude

`func (o *ManipulateRouteWaypoint) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *ManipulateRouteWaypoint) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *ManipulateRouteWaypoint) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.


### GetLongitude

`func (o *ManipulateRouteWaypoint) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *ManipulateRouteWaypoint) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *ManipulateRouteWaypoint) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.


### GetRadius

`func (o *ManipulateRouteWaypoint) GetRadius() int32`

GetRadius returns the Radius field if non-nil, zero value otherwise.

### GetRadiusOk

`func (o *ManipulateRouteWaypoint) GetRadiusOk() (*int32, bool)`

GetRadiusOk returns a tuple with the Radius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadius

`func (o *ManipulateRouteWaypoint) SetRadius(v int32)`

SetRadius sets Radius field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


