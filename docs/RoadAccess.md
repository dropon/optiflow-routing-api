# RoadAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Latitude** | **float64** | The latitude value in degrees (WGS84/EPSG:4326) from south to north. | 
**Longitude** | **float64** | The longitude value in degrees (WGS84/EPSG:4326) from west to east. | 

## Methods

### NewRoadAccess

`func NewRoadAccess(latitude float64, longitude float64, ) *RoadAccess`

NewRoadAccess instantiates a new RoadAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoadAccessWithDefaults

`func NewRoadAccessWithDefaults() *RoadAccess`

NewRoadAccessWithDefaults instantiates a new RoadAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLatitude

`func (o *RoadAccess) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *RoadAccess) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *RoadAccess) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.


### GetLongitude

`func (o *RoadAccess) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *RoadAccess) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *RoadAccess) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


