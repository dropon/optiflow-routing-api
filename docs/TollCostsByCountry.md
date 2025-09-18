# TollCostsByCountry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | **string** | Countries are represented according to their [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) or [ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2) if referring to a subdivision. | 
**Price** | [**TollPrice**](TollPrice.md) |  | 
**ConvertedPrice** | Pointer to [**TollPrice**](TollPrice.md) |  | [optional] 

## Methods

### NewTollCostsByCountry

`func NewTollCostsByCountry(countryCode string, price TollPrice, ) *TollCostsByCountry`

NewTollCostsByCountry instantiates a new TollCostsByCountry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTollCostsByCountryWithDefaults

`func NewTollCostsByCountryWithDefaults() *TollCostsByCountry`

NewTollCostsByCountryWithDefaults instantiates a new TollCostsByCountry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *TollCostsByCountry) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *TollCostsByCountry) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *TollCostsByCountry) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetPrice

`func (o *TollCostsByCountry) GetPrice() TollPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *TollCostsByCountry) GetPriceOk() (*TollPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *TollCostsByCountry) SetPrice(v TollPrice)`

SetPrice sets Price field to given value.


### GetConvertedPrice

`func (o *TollCostsByCountry) GetConvertedPrice() TollPrice`

GetConvertedPrice returns the ConvertedPrice field if non-nil, zero value otherwise.

### GetConvertedPriceOk

`func (o *TollCostsByCountry) GetConvertedPriceOk() (*TollPrice, bool)`

GetConvertedPriceOk returns a tuple with the ConvertedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertedPrice

`func (o *TollCostsByCountry) SetConvertedPrice(v TollPrice)`

SetConvertedPrice sets ConvertedPrice field to given value.

### HasConvertedPrice

`func (o *TollCostsByCountry) HasConvertedPrice() bool`

HasConvertedPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


