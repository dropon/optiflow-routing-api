# ReachableLocations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reachable** | Pointer to [**[]ReachableLocation**](ReachableLocation.md) | Reachable locations ordered by distance or travel time (depending on the request **horizonType**). | [optional] 
**Unreachable** | Pointer to **[]int32** | Indexes of the unreachable locations as given in the request locations. | [optional] 
**Warnings** | Pointer to [**[]Warning**](Warning.md) | A list of warnings concerning the validity of the result. | [optional] 

## Methods

### NewReachableLocations

`func NewReachableLocations() *ReachableLocations`

NewReachableLocations instantiates a new ReachableLocations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReachableLocationsWithDefaults

`func NewReachableLocationsWithDefaults() *ReachableLocations`

NewReachableLocationsWithDefaults instantiates a new ReachableLocations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReachable

`func (o *ReachableLocations) GetReachable() []ReachableLocation`

GetReachable returns the Reachable field if non-nil, zero value otherwise.

### GetReachableOk

`func (o *ReachableLocations) GetReachableOk() (*[]ReachableLocation, bool)`

GetReachableOk returns a tuple with the Reachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachable

`func (o *ReachableLocations) SetReachable(v []ReachableLocation)`

SetReachable sets Reachable field to given value.

### HasReachable

`func (o *ReachableLocations) HasReachable() bool`

HasReachable returns a boolean if a field has been set.

### GetUnreachable

`func (o *ReachableLocations) GetUnreachable() []int32`

GetUnreachable returns the Unreachable field if non-nil, zero value otherwise.

### GetUnreachableOk

`func (o *ReachableLocations) GetUnreachableOk() (*[]int32, bool)`

GetUnreachableOk returns a tuple with the Unreachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnreachable

`func (o *ReachableLocations) SetUnreachable(v []int32)`

SetUnreachable sets Unreachable field to given value.

### HasUnreachable

`func (o *ReachableLocations) HasUnreachable() bool`

HasUnreachable returns a boolean if a field has been set.

### GetWarnings

`func (o *ReachableLocations) GetWarnings() []Warning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *ReachableLocations) GetWarningsOk() (*[]Warning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *ReachableLocations) SetWarnings(v []Warning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *ReachableLocations) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


