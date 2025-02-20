# ElectricityPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | **float64** | The electricity price in the specified currency. | 
**Currency** | **string** | The currency code according to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217). | 

## Methods

### NewElectricityPrice

`func NewElectricityPrice(price float64, currency string, ) *ElectricityPrice`

NewElectricityPrice instantiates a new ElectricityPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewElectricityPriceWithDefaults

`func NewElectricityPriceWithDefaults() *ElectricityPrice`

NewElectricityPriceWithDefaults instantiates a new ElectricityPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *ElectricityPrice) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ElectricityPrice) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ElectricityPrice) SetPrice(v float64)`

SetPrice sets Price field to given value.


### GetCurrency

`func (o *ElectricityPrice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ElectricityPrice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ElectricityPrice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


