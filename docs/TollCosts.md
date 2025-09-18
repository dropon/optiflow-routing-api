# TollCosts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prices** | [**[]TollPrice**](TollPrice.md) | The toll prices payable in the given currencies. The price may be 0 if no toll applies for a currency in a country passed by the route. | 
**ConvertedPrice** | Pointer to [**TollPrice**](TollPrice.md) | The cumulated toll price in the converted currency. | [optional] 
**Countries** | [**[]TollCostsByCountry**](TollCostsByCountry.md) | The toll prices by country or subdivision. The price may be 0 if no toll applies in a country passed by the route. | 
**ContainsApproximatedSections** | Pointer to **bool** | True, if the start or destination waypoint is located inside a toll section. In such cases the exact toll price cannot be calculated and the closest toll location after the waypoint is used to approximate the toll price. When toll sections are requested, the affected section is marked as well. | [optional] 

## Methods

### NewTollCosts

`func NewTollCosts(prices []TollPrice, countries []TollCostsByCountry, ) *TollCosts`

NewTollCosts instantiates a new TollCosts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTollCostsWithDefaults

`func NewTollCostsWithDefaults() *TollCosts`

NewTollCostsWithDefaults instantiates a new TollCosts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrices

`func (o *TollCosts) GetPrices() []TollPrice`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *TollCosts) GetPricesOk() (*[]TollPrice, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *TollCosts) SetPrices(v []TollPrice)`

SetPrices sets Prices field to given value.


### GetConvertedPrice

`func (o *TollCosts) GetConvertedPrice() TollPrice`

GetConvertedPrice returns the ConvertedPrice field if non-nil, zero value otherwise.

### GetConvertedPriceOk

`func (o *TollCosts) GetConvertedPriceOk() (*TollPrice, bool)`

GetConvertedPriceOk returns a tuple with the ConvertedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertedPrice

`func (o *TollCosts) SetConvertedPrice(v TollPrice)`

SetConvertedPrice sets ConvertedPrice field to given value.

### HasConvertedPrice

`func (o *TollCosts) HasConvertedPrice() bool`

HasConvertedPrice returns a boolean if a field has been set.

### GetCountries

`func (o *TollCosts) GetCountries() []TollCostsByCountry`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *TollCosts) GetCountriesOk() (*[]TollCostsByCountry, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *TollCosts) SetCountries(v []TollCostsByCountry)`

SetCountries sets Countries field to given value.


### GetContainsApproximatedSections

`func (o *TollCosts) GetContainsApproximatedSections() bool`

GetContainsApproximatedSections returns the ContainsApproximatedSections field if non-nil, zero value otherwise.

### GetContainsApproximatedSectionsOk

`func (o *TollCosts) GetContainsApproximatedSectionsOk() (*bool, bool)`

GetContainsApproximatedSectionsOk returns a tuple with the ContainsApproximatedSections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsApproximatedSections

`func (o *TollCosts) SetContainsApproximatedSections(v bool)`

SetContainsApproximatedSections sets ContainsApproximatedSections field to given value.

### HasContainsApproximatedSections

`func (o *TollCosts) HasContainsApproximatedSections() bool`

HasContainsApproximatedSections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


