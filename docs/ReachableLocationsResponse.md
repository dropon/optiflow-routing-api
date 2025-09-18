# ReachableLocationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**CalculationStatus**](CalculationStatus.md) |  | [optional] 
**ReachableLocations** | Pointer to [**ReachableLocations**](ReachableLocations.md) |  | [optional] 
**Error** | Pointer to [**ErrorResponse**](ErrorResponse.md) |  | [optional] 

## Methods

### NewReachableLocationsResponse

`func NewReachableLocationsResponse() *ReachableLocationsResponse`

NewReachableLocationsResponse instantiates a new ReachableLocationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReachableLocationsResponseWithDefaults

`func NewReachableLocationsResponseWithDefaults() *ReachableLocationsResponse`

NewReachableLocationsResponseWithDefaults instantiates a new ReachableLocationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ReachableLocationsResponse) GetStatus() CalculationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReachableLocationsResponse) GetStatusOk() (*CalculationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReachableLocationsResponse) SetStatus(v CalculationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReachableLocationsResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReachableLocations

`func (o *ReachableLocationsResponse) GetReachableLocations() ReachableLocations`

GetReachableLocations returns the ReachableLocations field if non-nil, zero value otherwise.

### GetReachableLocationsOk

`func (o *ReachableLocationsResponse) GetReachableLocationsOk() (*ReachableLocations, bool)`

GetReachableLocationsOk returns a tuple with the ReachableLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachableLocations

`func (o *ReachableLocationsResponse) SetReachableLocations(v ReachableLocations)`

SetReachableLocations sets ReachableLocations field to given value.

### HasReachableLocations

`func (o *ReachableLocationsResponse) HasReachableLocations() bool`

HasReachableLocations returns a boolean if a field has been set.

### GetError

`func (o *ReachableLocationsResponse) GetError() ErrorResponse`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ReachableLocationsResponse) GetErrorOk() (*ErrorResponse, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ReachableLocationsResponse) SetError(v ErrorResponse)`

SetError sets Error field to given value.

### HasError

`func (o *ReachableLocationsResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


