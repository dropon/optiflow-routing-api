# PositionOnRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Latitude** | **float64** | The latitude value in degrees (WGS84/EPSG:4326) from south to north. | 
**Longitude** | **float64** | The longitude value in degrees (WGS84/EPSG:4326) from west to east. | 
**NextWaypointName** | **string** | The name of the next waypoint on the route. | 
**Heading** | Pointer to **int32** | The current heading of the vehicle [deg]. It denotes the driving direction, North represents 0 degrees, East represents 90 degrees, South represents 180 degrees, West represents 270 degrees.  If specified and if the heading of the vehicle is not in the direction of the route near the given **position**, it is assumed that the vehicle has left the route so that a new route will be calculated. This new route will then be  calculated from the road closest to the position matching the vehicle heading. For more information see this [concept](./concepts/estimated-time-arrival). | [optional] 
**HeadingTolerance** | Pointer to **int32** | Denotes the tolerance between **heading** and the direction of a road, i.e. roads with a direction of **heading**±**headingTolerance** are taken into account. Applies only if **heading** is specified. | [optional] [default to 45]

## Methods

### NewPositionOnRoute

`func NewPositionOnRoute(latitude float64, longitude float64, nextWaypointName string, ) *PositionOnRoute`

NewPositionOnRoute instantiates a new PositionOnRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPositionOnRouteWithDefaults

`func NewPositionOnRouteWithDefaults() *PositionOnRoute`

NewPositionOnRouteWithDefaults instantiates a new PositionOnRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLatitude

`func (o *PositionOnRoute) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *PositionOnRoute) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *PositionOnRoute) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.


### GetLongitude

`func (o *PositionOnRoute) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *PositionOnRoute) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *PositionOnRoute) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.


### GetNextWaypointName

`func (o *PositionOnRoute) GetNextWaypointName() string`

GetNextWaypointName returns the NextWaypointName field if non-nil, zero value otherwise.

### GetNextWaypointNameOk

`func (o *PositionOnRoute) GetNextWaypointNameOk() (*string, bool)`

GetNextWaypointNameOk returns a tuple with the NextWaypointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextWaypointName

`func (o *PositionOnRoute) SetNextWaypointName(v string)`

SetNextWaypointName sets NextWaypointName field to given value.


### GetHeading

`func (o *PositionOnRoute) GetHeading() int32`

GetHeading returns the Heading field if non-nil, zero value otherwise.

### GetHeadingOk

`func (o *PositionOnRoute) GetHeadingOk() (*int32, bool)`

GetHeadingOk returns a tuple with the Heading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeading

`func (o *PositionOnRoute) SetHeading(v int32)`

SetHeading sets Heading field to given value.

### HasHeading

`func (o *PositionOnRoute) HasHeading() bool`

HasHeading returns a boolean if a field has been set.

### GetHeadingTolerance

`func (o *PositionOnRoute) GetHeadingTolerance() int32`

GetHeadingTolerance returns the HeadingTolerance field if non-nil, zero value otherwise.

### GetHeadingToleranceOk

`func (o *PositionOnRoute) GetHeadingToleranceOk() (*int32, bool)`

GetHeadingToleranceOk returns a tuple with the HeadingTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadingTolerance

`func (o *PositionOnRoute) SetHeadingTolerance(v int32)`

SetHeadingTolerance sets HeadingTolerance field to given value.

### HasHeadingTolerance

`func (o *PositionOnRoute) HasHeadingTolerance() bool`

HasHeadingTolerance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


