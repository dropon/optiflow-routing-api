# DeliveryOnlyEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessType** | [**AccessType**](AccessType.md) |  | 
**RelatedEventIndex** | Pointer to **int32** | For **accessType** _ENTER_ this index points to the corresponding event with **accessType** _EXIT_ and vice-versa. Not present otherwise. | [optional] 
**Polyline** | Pointer to **string** | The polyline of the delivery-only event in the format specified by **options[polylineFormat]**. Only present for **accessType** _ENTER_. Requires _DELIVERY_ONLY_EVENTS_POLYLINE_ to be requested. | [optional] 

## Methods

### NewDeliveryOnlyEvent

`func NewDeliveryOnlyEvent(accessType AccessType, ) *DeliveryOnlyEvent`

NewDeliveryOnlyEvent instantiates a new DeliveryOnlyEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryOnlyEventWithDefaults

`func NewDeliveryOnlyEventWithDefaults() *DeliveryOnlyEvent`

NewDeliveryOnlyEventWithDefaults instantiates a new DeliveryOnlyEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessType

`func (o *DeliveryOnlyEvent) GetAccessType() AccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *DeliveryOnlyEvent) GetAccessTypeOk() (*AccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *DeliveryOnlyEvent) SetAccessType(v AccessType)`

SetAccessType sets AccessType field to given value.


### GetRelatedEventIndex

`func (o *DeliveryOnlyEvent) GetRelatedEventIndex() int32`

GetRelatedEventIndex returns the RelatedEventIndex field if non-nil, zero value otherwise.

### GetRelatedEventIndexOk

`func (o *DeliveryOnlyEvent) GetRelatedEventIndexOk() (*int32, bool)`

GetRelatedEventIndexOk returns a tuple with the RelatedEventIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedEventIndex

`func (o *DeliveryOnlyEvent) SetRelatedEventIndex(v int32)`

SetRelatedEventIndex sets RelatedEventIndex field to given value.

### HasRelatedEventIndex

`func (o *DeliveryOnlyEvent) HasRelatedEventIndex() bool`

HasRelatedEventIndex returns a boolean if a field has been set.

### GetPolyline

`func (o *DeliveryOnlyEvent) GetPolyline() string`

GetPolyline returns the Polyline field if non-nil, zero value otherwise.

### GetPolylineOk

`func (o *DeliveryOnlyEvent) GetPolylineOk() (*string, bool)`

GetPolylineOk returns a tuple with the Polyline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolyline

`func (o *DeliveryOnlyEvent) SetPolyline(v string)`

SetPolyline sets Polyline field to given value.

### HasPolyline

`func (o *DeliveryOnlyEvent) HasPolyline() bool`

HasPolyline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


