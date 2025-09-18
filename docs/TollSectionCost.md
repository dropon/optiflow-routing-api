# TollSectionCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | **float64** | The price in the specified currency. | 
**Currency** | **string** | The currency code according to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217). | 
**PaymentMethods** | Pointer to [**[]PaymentMethod**](PaymentMethod.md) | The payment methods for toll costs.    * &#x60;ELECTRONIC_TOLL_COLLECTION_SUBSCRIPTION&#x60; - Electronic toll collection system with a subscription required.    * &#x60;ELECTRONIC_TOLL_COLLECTION&#x60; - Electronic toll collection system with no subscription required.    * &#x60;CASH&#x60; - Cash payment at a toll booth.    * &#x60;CREDIT_CARD&#x60; - Credit card payment at a toll booth. | [optional] 
**EtcSubscriptions** | Pointer to **[]string** | The required electronic toll collection subscriptions for the payment method _ELECTRONIC_TOLL_COLLECTION_SUBSCRIPTION_. | [optional] 
**ConvertedPrice** | Pointer to [**TollPrice**](TollPrice.md) | The price of the section in the converted currency. | [optional] 

## Methods

### NewTollSectionCost

`func NewTollSectionCost(price float64, currency string, ) *TollSectionCost`

NewTollSectionCost instantiates a new TollSectionCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTollSectionCostWithDefaults

`func NewTollSectionCostWithDefaults() *TollSectionCost`

NewTollSectionCostWithDefaults instantiates a new TollSectionCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *TollSectionCost) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *TollSectionCost) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *TollSectionCost) SetPrice(v float64)`

SetPrice sets Price field to given value.


### GetCurrency

`func (o *TollSectionCost) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TollSectionCost) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TollSectionCost) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetPaymentMethods

`func (o *TollSectionCost) GetPaymentMethods() []PaymentMethod`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *TollSectionCost) GetPaymentMethodsOk() (*[]PaymentMethod, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *TollSectionCost) SetPaymentMethods(v []PaymentMethod)`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *TollSectionCost) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.

### GetEtcSubscriptions

`func (o *TollSectionCost) GetEtcSubscriptions() []string`

GetEtcSubscriptions returns the EtcSubscriptions field if non-nil, zero value otherwise.

### GetEtcSubscriptionsOk

`func (o *TollSectionCost) GetEtcSubscriptionsOk() (*[]string, bool)`

GetEtcSubscriptionsOk returns a tuple with the EtcSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtcSubscriptions

`func (o *TollSectionCost) SetEtcSubscriptions(v []string)`

SetEtcSubscriptions sets EtcSubscriptions field to given value.

### HasEtcSubscriptions

`func (o *TollSectionCost) HasEtcSubscriptions() bool`

HasEtcSubscriptions returns a boolean if a field has been set.

### GetConvertedPrice

`func (o *TollSectionCost) GetConvertedPrice() TollPrice`

GetConvertedPrice returns the ConvertedPrice field if non-nil, zero value otherwise.

### GetConvertedPriceOk

`func (o *TollSectionCost) GetConvertedPriceOk() (*TollPrice, bool)`

GetConvertedPriceOk returns a tuple with the ConvertedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertedPrice

`func (o *TollSectionCost) SetConvertedPrice(v TollPrice)`

SetConvertedPrice sets ConvertedPrice field to given value.

### HasConvertedPrice

`func (o *TollSectionCost) HasConvertedPrice() bool`

HasConvertedPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


