# TollSystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the toll system. | 
**OperatorName** | Pointer to **string** | The name of the toll operator. | [optional] 
**TariffVersion** | Pointer to **string** | The tariff version that has been used. | [optional] 
**TariffVersionValidFrom** | Pointer to **time.Time** | The starting date of the tariff version validity formatted according to [RFC 3339](https://tools.ietf.org/html/rfc3339). | [optional] 

## Methods

### NewTollSystem

`func NewTollSystem(name string, ) *TollSystem`

NewTollSystem instantiates a new TollSystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTollSystemWithDefaults

`func NewTollSystemWithDefaults() *TollSystem`

NewTollSystemWithDefaults instantiates a new TollSystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TollSystem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TollSystem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TollSystem) SetName(v string)`

SetName sets Name field to given value.


### GetOperatorName

`func (o *TollSystem) GetOperatorName() string`

GetOperatorName returns the OperatorName field if non-nil, zero value otherwise.

### GetOperatorNameOk

`func (o *TollSystem) GetOperatorNameOk() (*string, bool)`

GetOperatorNameOk returns a tuple with the OperatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorName

`func (o *TollSystem) SetOperatorName(v string)`

SetOperatorName sets OperatorName field to given value.

### HasOperatorName

`func (o *TollSystem) HasOperatorName() bool`

HasOperatorName returns a boolean if a field has been set.

### GetTariffVersion

`func (o *TollSystem) GetTariffVersion() string`

GetTariffVersion returns the TariffVersion field if non-nil, zero value otherwise.

### GetTariffVersionOk

`func (o *TollSystem) GetTariffVersionOk() (*string, bool)`

GetTariffVersionOk returns a tuple with the TariffVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTariffVersion

`func (o *TollSystem) SetTariffVersion(v string)`

SetTariffVersion sets TariffVersion field to given value.

### HasTariffVersion

`func (o *TollSystem) HasTariffVersion() bool`

HasTariffVersion returns a boolean if a field has been set.

### GetTariffVersionValidFrom

`func (o *TollSystem) GetTariffVersionValidFrom() time.Time`

GetTariffVersionValidFrom returns the TariffVersionValidFrom field if non-nil, zero value otherwise.

### GetTariffVersionValidFromOk

`func (o *TollSystem) GetTariffVersionValidFromOk() (*time.Time, bool)`

GetTariffVersionValidFromOk returns a tuple with the TariffVersionValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTariffVersionValidFrom

`func (o *TollSystem) SetTariffVersionValidFrom(v time.Time)`

SetTariffVersionValidFrom sets TariffVersionValidFrom field to given value.

### HasTariffVersionValidFrom

`func (o *TollSystem) HasTariffVersionValidFrom() bool`

HasTariffVersionValidFrom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


