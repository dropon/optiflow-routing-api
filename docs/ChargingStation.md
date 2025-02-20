# ChargingStation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Power** | **int32** | The power delivered by the charging station [kW]. | 
**CurrentType** | [**CurrentType**](CurrentType.md) |  | [default to ALTERNATING]
**SetupTime** | Pointer to **int32** | The time to setup the connection [s]. | [optional] [default to 0]
**KWhPrice** | Pointer to **float64** | The price of one kWh in the currency defined in **options[currency]**. | [optional] [default to 0]
**UseServiceTimeForCharging** | Pointer to **bool** | If true, the service time can be used for charging. Will be ignored, if **serviceTime** is 0. This can have  an influence on charging optimization because the additional time needed for charging is minimized. | [optional] [default to false]

## Methods

### NewChargingStation

`func NewChargingStation(power int32, currentType CurrentType, ) *ChargingStation`

NewChargingStation instantiates a new ChargingStation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargingStationWithDefaults

`func NewChargingStationWithDefaults() *ChargingStation`

NewChargingStationWithDefaults instantiates a new ChargingStation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPower

`func (o *ChargingStation) GetPower() int32`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *ChargingStation) GetPowerOk() (*int32, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *ChargingStation) SetPower(v int32)`

SetPower sets Power field to given value.


### GetCurrentType

`func (o *ChargingStation) GetCurrentType() CurrentType`

GetCurrentType returns the CurrentType field if non-nil, zero value otherwise.

### GetCurrentTypeOk

`func (o *ChargingStation) GetCurrentTypeOk() (*CurrentType, bool)`

GetCurrentTypeOk returns a tuple with the CurrentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentType

`func (o *ChargingStation) SetCurrentType(v CurrentType)`

SetCurrentType sets CurrentType field to given value.


### GetSetupTime

`func (o *ChargingStation) GetSetupTime() int32`

GetSetupTime returns the SetupTime field if non-nil, zero value otherwise.

### GetSetupTimeOk

`func (o *ChargingStation) GetSetupTimeOk() (*int32, bool)`

GetSetupTimeOk returns a tuple with the SetupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupTime

`func (o *ChargingStation) SetSetupTime(v int32)`

SetSetupTime sets SetupTime field to given value.

### HasSetupTime

`func (o *ChargingStation) HasSetupTime() bool`

HasSetupTime returns a boolean if a field has been set.

### GetKWhPrice

`func (o *ChargingStation) GetKWhPrice() float64`

GetKWhPrice returns the KWhPrice field if non-nil, zero value otherwise.

### GetKWhPriceOk

`func (o *ChargingStation) GetKWhPriceOk() (*float64, bool)`

GetKWhPriceOk returns a tuple with the KWhPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKWhPrice

`func (o *ChargingStation) SetKWhPrice(v float64)`

SetKWhPrice sets KWhPrice field to given value.

### HasKWhPrice

`func (o *ChargingStation) HasKWhPrice() bool`

HasKWhPrice returns a boolean if a field has been set.

### GetUseServiceTimeForCharging

`func (o *ChargingStation) GetUseServiceTimeForCharging() bool`

GetUseServiceTimeForCharging returns the UseServiceTimeForCharging field if non-nil, zero value otherwise.

### GetUseServiceTimeForChargingOk

`func (o *ChargingStation) GetUseServiceTimeForChargingOk() (*bool, bool)`

GetUseServiceTimeForChargingOk returns a tuple with the UseServiceTimeForCharging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseServiceTimeForCharging

`func (o *ChargingStation) SetUseServiceTimeForCharging(v bool)`

SetUseServiceTimeForCharging sets UseServiceTimeForCharging field to given value.

### HasUseServiceTimeForCharging

`func (o *ChargingStation) HasUseServiceTimeForCharging() bool`

HasUseServiceTimeForCharging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


