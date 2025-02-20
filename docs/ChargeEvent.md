# ChargeEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChargingTime** | **int32** | The time spent for charging the battery [s]. The charging time is a proposal, currently for information only. It is not included in the travel time of the route and the start time of subsequent events is not offset by it. | 
**ElectricityCharged** | **float64** | The amount of electricity charged [kWh]. | 
**PercentageCharged** | **int32** | The percentage of battery charged [%]. | 
**Cost** | [**ElectricityPrice**](ElectricityPrice.md) |  | 

## Methods

### NewChargeEvent

`func NewChargeEvent(chargingTime int32, electricityCharged float64, percentageCharged int32, cost ElectricityPrice, ) *ChargeEvent`

NewChargeEvent instantiates a new ChargeEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeEventWithDefaults

`func NewChargeEventWithDefaults() *ChargeEvent`

NewChargeEventWithDefaults instantiates a new ChargeEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChargingTime

`func (o *ChargeEvent) GetChargingTime() int32`

GetChargingTime returns the ChargingTime field if non-nil, zero value otherwise.

### GetChargingTimeOk

`func (o *ChargeEvent) GetChargingTimeOk() (*int32, bool)`

GetChargingTimeOk returns a tuple with the ChargingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargingTime

`func (o *ChargeEvent) SetChargingTime(v int32)`

SetChargingTime sets ChargingTime field to given value.


### GetElectricityCharged

`func (o *ChargeEvent) GetElectricityCharged() float64`

GetElectricityCharged returns the ElectricityCharged field if non-nil, zero value otherwise.

### GetElectricityChargedOk

`func (o *ChargeEvent) GetElectricityChargedOk() (*float64, bool)`

GetElectricityChargedOk returns a tuple with the ElectricityCharged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElectricityCharged

`func (o *ChargeEvent) SetElectricityCharged(v float64)`

SetElectricityCharged sets ElectricityCharged field to given value.


### GetPercentageCharged

`func (o *ChargeEvent) GetPercentageCharged() int32`

GetPercentageCharged returns the PercentageCharged field if non-nil, zero value otherwise.

### GetPercentageChargedOk

`func (o *ChargeEvent) GetPercentageChargedOk() (*int32, bool)`

GetPercentageChargedOk returns a tuple with the PercentageCharged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageCharged

`func (o *ChargeEvent) SetPercentageCharged(v int32)`

SetPercentageCharged sets PercentageCharged field to given value.


### GetCost

`func (o *ChargeEvent) GetCost() ElectricityPrice`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ChargeEvent) GetCostOk() (*ElectricityPrice, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ChargeEvent) SetCost(v ElectricityPrice)`

SetCost sets Cost field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


