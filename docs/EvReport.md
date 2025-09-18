# EvReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ElectricityConsumption** | **float64** | The electricity consumption since the start of the route [kWh]. | 
**BatteryStateOfCharge** | **float64** | The remaining state of charge at the end of the route [%]. May be below the defined minimum state of charge or even below zero, if the electricity consumption exceeds the available energy in the battery and charging is not possible before falling below zero. | 
**ChargingTime** | **int32** | The time spent for charging the battery [s]. The charging time is a proposal, currently for information only. It is not included in the travel time of the route and the start time of subsequent events is not offset by it. | 
**ElectricityCharged** | **float64** | The amount of electricity charged along the route [kWh]. | 
**PercentageCharged** | **int32** | The percentage of battery charged along the route [%]. | 
**Cost** | [**ElectricityPrice**](ElectricityPrice.md) |  | 

## Methods

### NewEvReport

`func NewEvReport(electricityConsumption float64, batteryStateOfCharge float64, chargingTime int32, electricityCharged float64, percentageCharged int32, cost ElectricityPrice, ) *EvReport`

NewEvReport instantiates a new EvReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvReportWithDefaults

`func NewEvReportWithDefaults() *EvReport`

NewEvReportWithDefaults instantiates a new EvReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElectricityConsumption

`func (o *EvReport) GetElectricityConsumption() float64`

GetElectricityConsumption returns the ElectricityConsumption field if non-nil, zero value otherwise.

### GetElectricityConsumptionOk

`func (o *EvReport) GetElectricityConsumptionOk() (*float64, bool)`

GetElectricityConsumptionOk returns a tuple with the ElectricityConsumption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElectricityConsumption

`func (o *EvReport) SetElectricityConsumption(v float64)`

SetElectricityConsumption sets ElectricityConsumption field to given value.


### GetBatteryStateOfCharge

`func (o *EvReport) GetBatteryStateOfCharge() float64`

GetBatteryStateOfCharge returns the BatteryStateOfCharge field if non-nil, zero value otherwise.

### GetBatteryStateOfChargeOk

`func (o *EvReport) GetBatteryStateOfChargeOk() (*float64, bool)`

GetBatteryStateOfChargeOk returns a tuple with the BatteryStateOfCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryStateOfCharge

`func (o *EvReport) SetBatteryStateOfCharge(v float64)`

SetBatteryStateOfCharge sets BatteryStateOfCharge field to given value.


### GetChargingTime

`func (o *EvReport) GetChargingTime() int32`

GetChargingTime returns the ChargingTime field if non-nil, zero value otherwise.

### GetChargingTimeOk

`func (o *EvReport) GetChargingTimeOk() (*int32, bool)`

GetChargingTimeOk returns a tuple with the ChargingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargingTime

`func (o *EvReport) SetChargingTime(v int32)`

SetChargingTime sets ChargingTime field to given value.


### GetElectricityCharged

`func (o *EvReport) GetElectricityCharged() float64`

GetElectricityCharged returns the ElectricityCharged field if non-nil, zero value otherwise.

### GetElectricityChargedOk

`func (o *EvReport) GetElectricityChargedOk() (*float64, bool)`

GetElectricityChargedOk returns a tuple with the ElectricityCharged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElectricityCharged

`func (o *EvReport) SetElectricityCharged(v float64)`

SetElectricityCharged sets ElectricityCharged field to given value.


### GetPercentageCharged

`func (o *EvReport) GetPercentageCharged() int32`

GetPercentageCharged returns the PercentageCharged field if non-nil, zero value otherwise.

### GetPercentageChargedOk

`func (o *EvReport) GetPercentageChargedOk() (*int32, bool)`

GetPercentageChargedOk returns a tuple with the PercentageCharged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageCharged

`func (o *EvReport) SetPercentageCharged(v int32)`

SetPercentageCharged sets PercentageCharged field to given value.


### GetCost

`func (o *EvReport) GetCost() ElectricityPrice`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *EvReport) GetCostOk() (*ElectricityPrice, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *EvReport) SetCost(v ElectricityPrice)`

SetCost sets Cost field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


