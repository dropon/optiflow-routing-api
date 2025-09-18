# ExchangeRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** | The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code of the currency from which a price has been converted. | 
**Rate** | **float64** | The exchange rate to convert a price from the base currency to this currency. | 

## Methods

### NewExchangeRate

`func NewExchangeRate(currency string, rate float64, ) *ExchangeRate`

NewExchangeRate instantiates a new ExchangeRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeRateWithDefaults

`func NewExchangeRateWithDefaults() *ExchangeRate`

NewExchangeRateWithDefaults instantiates a new ExchangeRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *ExchangeRate) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ExchangeRate) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ExchangeRate) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetRate

`func (o *ExchangeRate) GetRate() float64`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *ExchangeRate) GetRateOk() (*float64, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *ExchangeRate) SetRate(v float64)`

SetRate sets Rate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


