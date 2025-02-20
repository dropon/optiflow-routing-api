# TrafficEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delay** | **int32** | The delay due to this incident [s]. | 
**Description** | Pointer to **string** | The description of the incident in the language specified by the parameter **options[language]**.  | [optional] 
**Language** | Pointer to **string** | The language of the description, if the given language is not supported for this incident. Not present otherwise. | [optional] 
**AccessType** | [**AccessType**](AccessType.md) |  | 
**RelatedEventIndex** | Pointer to **int32** | For **accessType** _ENTER_ this index points to the corresponding event with **accessType** _EXIT_ and vice-versa. Not present otherwise. | [optional] 
**Polyline** | Pointer to **string** | The polyline of the traffic event in the format specified by **options[polylineFormat]**. Only present for **accessType** _ENTER_. Requires _TRAFFIC_EVENTS_POLYLINE_ to be requested. | [optional] 

## Methods

### NewTrafficEvent

`func NewTrafficEvent(delay int32, accessType AccessType, ) *TrafficEvent`

NewTrafficEvent instantiates a new TrafficEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficEventWithDefaults

`func NewTrafficEventWithDefaults() *TrafficEvent`

NewTrafficEventWithDefaults instantiates a new TrafficEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelay

`func (o *TrafficEvent) GetDelay() int32`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *TrafficEvent) GetDelayOk() (*int32, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *TrafficEvent) SetDelay(v int32)`

SetDelay sets Delay field to given value.


### GetDescription

`func (o *TrafficEvent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TrafficEvent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TrafficEvent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TrafficEvent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLanguage

`func (o *TrafficEvent) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *TrafficEvent) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *TrafficEvent) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *TrafficEvent) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetAccessType

`func (o *TrafficEvent) GetAccessType() AccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *TrafficEvent) GetAccessTypeOk() (*AccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *TrafficEvent) SetAccessType(v AccessType)`

SetAccessType sets AccessType field to given value.


### GetRelatedEventIndex

`func (o *TrafficEvent) GetRelatedEventIndex() int32`

GetRelatedEventIndex returns the RelatedEventIndex field if non-nil, zero value otherwise.

### GetRelatedEventIndexOk

`func (o *TrafficEvent) GetRelatedEventIndexOk() (*int32, bool)`

GetRelatedEventIndexOk returns a tuple with the RelatedEventIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedEventIndex

`func (o *TrafficEvent) SetRelatedEventIndex(v int32)`

SetRelatedEventIndex sets RelatedEventIndex field to given value.

### HasRelatedEventIndex

`func (o *TrafficEvent) HasRelatedEventIndex() bool`

HasRelatedEventIndex returns a boolean if a field has been set.

### GetPolyline

`func (o *TrafficEvent) GetPolyline() string`

GetPolyline returns the Polyline field if non-nil, zero value otherwise.

### GetPolylineOk

`func (o *TrafficEvent) GetPolylineOk() (*string, bool)`

GetPolylineOk returns a tuple with the Polyline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolyline

`func (o *TrafficEvent) SetPolyline(v string)`

SetPolyline sets Polyline field to given value.

### HasPolyline

`func (o *TrafficEvent) HasPolyline() bool`

HasPolyline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


