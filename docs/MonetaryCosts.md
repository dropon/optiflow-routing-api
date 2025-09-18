# MonetaryCosts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** | The currency code according to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217). | 
**TotalCost** | **float64** | The total monetary cost of the route. | 
**DistanceCost** | **float64** | The distance cost based on the **monetaryCostOptions[costPerKilometer]** and the driving distance. | 
**WorkingTimeCost** | **float64** | The working time cost based on **monetaryCostOptions[costPerWorkingHour]** and the driving time. Break, service and rest times are not considered. | 
**EnergyCost** | **float64** | The energy cost based on the vehicle&#39;s consumption and the corresponding cost parameters. | 
**TollCost** | **float64** | The toll cost based on the route and the vehicle.   Toll prices do not include VAT.  | 

## Methods

### NewMonetaryCosts

`func NewMonetaryCosts(currency string, totalCost float64, distanceCost float64, workingTimeCost float64, energyCost float64, tollCost float64, ) *MonetaryCosts`

NewMonetaryCosts instantiates a new MonetaryCosts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonetaryCostsWithDefaults

`func NewMonetaryCostsWithDefaults() *MonetaryCosts`

NewMonetaryCostsWithDefaults instantiates a new MonetaryCosts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *MonetaryCosts) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *MonetaryCosts) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *MonetaryCosts) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetTotalCost

`func (o *MonetaryCosts) GetTotalCost() float64`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *MonetaryCosts) GetTotalCostOk() (*float64, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *MonetaryCosts) SetTotalCost(v float64)`

SetTotalCost sets TotalCost field to given value.


### GetDistanceCost

`func (o *MonetaryCosts) GetDistanceCost() float64`

GetDistanceCost returns the DistanceCost field if non-nil, zero value otherwise.

### GetDistanceCostOk

`func (o *MonetaryCosts) GetDistanceCostOk() (*float64, bool)`

GetDistanceCostOk returns a tuple with the DistanceCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceCost

`func (o *MonetaryCosts) SetDistanceCost(v float64)`

SetDistanceCost sets DistanceCost field to given value.


### GetWorkingTimeCost

`func (o *MonetaryCosts) GetWorkingTimeCost() float64`

GetWorkingTimeCost returns the WorkingTimeCost field if non-nil, zero value otherwise.

### GetWorkingTimeCostOk

`func (o *MonetaryCosts) GetWorkingTimeCostOk() (*float64, bool)`

GetWorkingTimeCostOk returns a tuple with the WorkingTimeCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingTimeCost

`func (o *MonetaryCosts) SetWorkingTimeCost(v float64)`

SetWorkingTimeCost sets WorkingTimeCost field to given value.


### GetEnergyCost

`func (o *MonetaryCosts) GetEnergyCost() float64`

GetEnergyCost returns the EnergyCost field if non-nil, zero value otherwise.

### GetEnergyCostOk

`func (o *MonetaryCosts) GetEnergyCostOk() (*float64, bool)`

GetEnergyCostOk returns a tuple with the EnergyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnergyCost

`func (o *MonetaryCosts) SetEnergyCost(v float64)`

SetEnergyCost sets EnergyCost field to given value.


### GetTollCost

`func (o *MonetaryCosts) GetTollCost() float64`

GetTollCost returns the TollCost field if non-nil, zero value otherwise.

### GetTollCostOk

`func (o *MonetaryCosts) GetTollCostOk() (*float64, bool)`

GetTollCostOk returns a tuple with the TollCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTollCost

`func (o *MonetaryCosts) SetTollCost(v float64)`

SetTollCost sets TollCost field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


