# ReachableAreasResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**CalculationStatus**](CalculationStatus.md) |  | [optional] 
**ReachableAreas** | Pointer to [**ReachableAreas**](ReachableAreas.md) |  | [optional] 
**Error** | Pointer to [**ErrorResponse**](ErrorResponse.md) |  | [optional] 

## Methods

### NewReachableAreasResponse

`func NewReachableAreasResponse() *ReachableAreasResponse`

NewReachableAreasResponse instantiates a new ReachableAreasResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReachableAreasResponseWithDefaults

`func NewReachableAreasResponseWithDefaults() *ReachableAreasResponse`

NewReachableAreasResponseWithDefaults instantiates a new ReachableAreasResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ReachableAreasResponse) GetStatus() CalculationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReachableAreasResponse) GetStatusOk() (*CalculationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReachableAreasResponse) SetStatus(v CalculationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReachableAreasResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReachableAreas

`func (o *ReachableAreasResponse) GetReachableAreas() ReachableAreas`

GetReachableAreas returns the ReachableAreas field if non-nil, zero value otherwise.

### GetReachableAreasOk

`func (o *ReachableAreasResponse) GetReachableAreasOk() (*ReachableAreas, bool)`

GetReachableAreasOk returns a tuple with the ReachableAreas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachableAreas

`func (o *ReachableAreasResponse) SetReachableAreas(v ReachableAreas)`

SetReachableAreas sets ReachableAreas field to given value.

### HasReachableAreas

`func (o *ReachableAreasResponse) HasReachableAreas() bool`

HasReachableAreas returns a boolean if a field has been set.

### GetError

`func (o *ReachableAreasResponse) GetError() ErrorResponse`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ReachableAreasResponse) GetErrorOk() (*ErrorResponse, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ReachableAreasResponse) SetError(v ErrorResponse)`

SetError sets Error field to given value.

### HasError

`func (o *ReachableAreasResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


