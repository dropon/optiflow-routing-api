# TollEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SectionIndex** | Pointer to **int32** | The index of the corresponding toll section.  The section indexes of a pair of ENTER and EXIT events define the range of sections between the two events. | [optional] 
**DisplayName** | Pointer to **string** | The display name of the toll event. Only present if a name is available. For example, this name can be a toll location name defined by the toll operator. In some cases the display names of the toll event and the toll section can differ. In case the toll section has been approximated, the display name of the affected toll event contains the hint \&quot;(approximated)\&quot;. | [optional] 
**AccessType** | [**AccessType**](AccessType.md) |  | 
**RelatedEventIndex** | Pointer to **int32** | For **accessType** _ENTER_ this index points to the corresponding event with **accessType** _EXIT_ and vice-versa. Not present otherwise. | [optional] 

## Methods

### NewTollEvent

`func NewTollEvent(accessType AccessType, ) *TollEvent`

NewTollEvent instantiates a new TollEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTollEventWithDefaults

`func NewTollEventWithDefaults() *TollEvent`

NewTollEventWithDefaults instantiates a new TollEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSectionIndex

`func (o *TollEvent) GetSectionIndex() int32`

GetSectionIndex returns the SectionIndex field if non-nil, zero value otherwise.

### GetSectionIndexOk

`func (o *TollEvent) GetSectionIndexOk() (*int32, bool)`

GetSectionIndexOk returns a tuple with the SectionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionIndex

`func (o *TollEvent) SetSectionIndex(v int32)`

SetSectionIndex sets SectionIndex field to given value.

### HasSectionIndex

`func (o *TollEvent) HasSectionIndex() bool`

HasSectionIndex returns a boolean if a field has been set.

### GetDisplayName

`func (o *TollEvent) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TollEvent) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TollEvent) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *TollEvent) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetAccessType

`func (o *TollEvent) GetAccessType() AccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *TollEvent) GetAccessTypeOk() (*AccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *TollEvent) SetAccessType(v AccessType)`

SetAccessType sets AccessType field to given value.


### GetRelatedEventIndex

`func (o *TollEvent) GetRelatedEventIndex() int32`

GetRelatedEventIndex returns the RelatedEventIndex field if non-nil, zero value otherwise.

### GetRelatedEventIndexOk

`func (o *TollEvent) GetRelatedEventIndexOk() (*int32, bool)`

GetRelatedEventIndexOk returns a tuple with the RelatedEventIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedEventIndex

`func (o *TollEvent) SetRelatedEventIndex(v int32)`

SetRelatedEventIndex sets RelatedEventIndex field to given value.

### HasRelatedEventIndex

`func (o *TollEvent) HasRelatedEventIndex() bool`

HasRelatedEventIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


