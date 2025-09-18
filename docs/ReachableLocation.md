# ReachableLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Distance** | Pointer to **int32** | The distance from the input waypoint to this location or vice versa. | [optional] 
**TravelTime** | Pointer to **int32** | The travel time from the input waypoint to this location or vice versa. | [optional] 
**Index** | Pointer to **int32** | The index of the reached input location. | [optional] 

## Methods

### NewReachableLocation

`func NewReachableLocation() *ReachableLocation`

NewReachableLocation instantiates a new ReachableLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReachableLocationWithDefaults

`func NewReachableLocationWithDefaults() *ReachableLocation`

NewReachableLocationWithDefaults instantiates a new ReachableLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistance

`func (o *ReachableLocation) GetDistance() int32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *ReachableLocation) GetDistanceOk() (*int32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *ReachableLocation) SetDistance(v int32)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *ReachableLocation) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### GetTravelTime

`func (o *ReachableLocation) GetTravelTime() int32`

GetTravelTime returns the TravelTime field if non-nil, zero value otherwise.

### GetTravelTimeOk

`func (o *ReachableLocation) GetTravelTimeOk() (*int32, bool)`

GetTravelTimeOk returns a tuple with the TravelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelTime

`func (o *ReachableLocation) SetTravelTime(v int32)`

SetTravelTime sets TravelTime field to given value.

### HasTravelTime

`func (o *ReachableLocation) HasTravelTime() bool`

HasTravelTime returns a boolean if a field has been set.

### GetIndex

`func (o *ReachableLocation) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ReachableLocation) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ReachableLocation) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *ReachableLocation) HasIndex() bool`

HasIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


