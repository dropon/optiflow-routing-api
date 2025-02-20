# CombinedTransportEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the combined transport. | 
**Type** | [**CombinedTransportType**](CombinedTransportType.md) |  | 
**AccessType** | [**AccessType**](AccessType.md) |  | 
**RelatedEventIndex** | Pointer to **int32** | For **accessType** _ENTER_ this index points to the corresponding event with **accessType** _EXIT_ and vice-versa. Not present otherwise. | [optional] 

## Methods

### NewCombinedTransportEvent

`func NewCombinedTransportEvent(name string, type_ CombinedTransportType, accessType AccessType, ) *CombinedTransportEvent`

NewCombinedTransportEvent instantiates a new CombinedTransportEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCombinedTransportEventWithDefaults

`func NewCombinedTransportEventWithDefaults() *CombinedTransportEvent`

NewCombinedTransportEventWithDefaults instantiates a new CombinedTransportEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CombinedTransportEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CombinedTransportEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CombinedTransportEvent) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CombinedTransportEvent) GetType() CombinedTransportType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CombinedTransportEvent) GetTypeOk() (*CombinedTransportType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CombinedTransportEvent) SetType(v CombinedTransportType)`

SetType sets Type field to given value.


### GetAccessType

`func (o *CombinedTransportEvent) GetAccessType() AccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *CombinedTransportEvent) GetAccessTypeOk() (*AccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *CombinedTransportEvent) SetAccessType(v AccessType)`

SetAccessType sets AccessType field to given value.


### GetRelatedEventIndex

`func (o *CombinedTransportEvent) GetRelatedEventIndex() int32`

GetRelatedEventIndex returns the RelatedEventIndex field if non-nil, zero value otherwise.

### GetRelatedEventIndexOk

`func (o *CombinedTransportEvent) GetRelatedEventIndexOk() (*int32, bool)`

GetRelatedEventIndexOk returns a tuple with the RelatedEventIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedEventIndex

`func (o *CombinedTransportEvent) SetRelatedEventIndex(v int32)`

SetRelatedEventIndex sets RelatedEventIndex field to given value.

### HasRelatedEventIndex

`func (o *CombinedTransportEvent) HasRelatedEventIndex() bool`

HasRelatedEventIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


