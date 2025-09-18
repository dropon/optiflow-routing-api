# UTCOffsetChangeEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UtcOffset** | **int32** | The new UTC offset [min]. | 

## Methods

### NewUTCOffsetChangeEvent

`func NewUTCOffsetChangeEvent(utcOffset int32, ) *UTCOffsetChangeEvent`

NewUTCOffsetChangeEvent instantiates a new UTCOffsetChangeEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUTCOffsetChangeEventWithDefaults

`func NewUTCOffsetChangeEventWithDefaults() *UTCOffsetChangeEvent`

NewUTCOffsetChangeEventWithDefaults instantiates a new UTCOffsetChangeEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUtcOffset

`func (o *UTCOffsetChangeEvent) GetUtcOffset() int32`

GetUtcOffset returns the UtcOffset field if non-nil, zero value otherwise.

### GetUtcOffsetOk

`func (o *UTCOffsetChangeEvent) GetUtcOffsetOk() (*int32, bool)`

GetUtcOffsetOk returns a tuple with the UtcOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtcOffset

`func (o *UTCOffsetChangeEvent) SetUtcOffset(v int32)`

SetUtcOffset sets UtcOffset field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


