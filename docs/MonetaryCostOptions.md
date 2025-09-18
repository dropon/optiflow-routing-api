# MonetaryCostOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CostPerKilometer** | Pointer to **float64** | Specifies the cost per kilometer. This value can contain the cost from the energy consumption, but it also possible to define the energy costs using **monetaryCostOptions[costPerKwh]** and **monetaryCostOptions[costPerFuelUnit]**. If it is not specified the default value depends on the used **profile**. The default value of the profile can change at any time. | [optional] 
**WorkingCostPerHour** | Pointer to **float64** | Specifies the cost per hour. It is not applied to service, break or rest periods. If it is not specified the default value depends on the used **profile**. The default value of the profile can change at any time. | [optional] 
**CostPerKwh** | Pointer to **float64** | Specifies the cost per kilowatt hour. Only relevant for **vehicle[engineType]** _ELECTRIC_ and _HYBRID_. | [optional] 
**CostPerFuelUnit** | Pointer to **float64** | Specifies the cost per fuel unit (per liter Diesel or per kg CNG). Only used for **vehicle[engineType]** _COMBUSTION_ and _HYBRID_. | [optional] 

## Methods

### NewMonetaryCostOptions

`func NewMonetaryCostOptions() *MonetaryCostOptions`

NewMonetaryCostOptions instantiates a new MonetaryCostOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonetaryCostOptionsWithDefaults

`func NewMonetaryCostOptionsWithDefaults() *MonetaryCostOptions`

NewMonetaryCostOptionsWithDefaults instantiates a new MonetaryCostOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCostPerKilometer

`func (o *MonetaryCostOptions) GetCostPerKilometer() float64`

GetCostPerKilometer returns the CostPerKilometer field if non-nil, zero value otherwise.

### GetCostPerKilometerOk

`func (o *MonetaryCostOptions) GetCostPerKilometerOk() (*float64, bool)`

GetCostPerKilometerOk returns a tuple with the CostPerKilometer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerKilometer

`func (o *MonetaryCostOptions) SetCostPerKilometer(v float64)`

SetCostPerKilometer sets CostPerKilometer field to given value.

### HasCostPerKilometer

`func (o *MonetaryCostOptions) HasCostPerKilometer() bool`

HasCostPerKilometer returns a boolean if a field has been set.

### GetWorkingCostPerHour

`func (o *MonetaryCostOptions) GetWorkingCostPerHour() float64`

GetWorkingCostPerHour returns the WorkingCostPerHour field if non-nil, zero value otherwise.

### GetWorkingCostPerHourOk

`func (o *MonetaryCostOptions) GetWorkingCostPerHourOk() (*float64, bool)`

GetWorkingCostPerHourOk returns a tuple with the WorkingCostPerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingCostPerHour

`func (o *MonetaryCostOptions) SetWorkingCostPerHour(v float64)`

SetWorkingCostPerHour sets WorkingCostPerHour field to given value.

### HasWorkingCostPerHour

`func (o *MonetaryCostOptions) HasWorkingCostPerHour() bool`

HasWorkingCostPerHour returns a boolean if a field has been set.

### GetCostPerKwh

`func (o *MonetaryCostOptions) GetCostPerKwh() float64`

GetCostPerKwh returns the CostPerKwh field if non-nil, zero value otherwise.

### GetCostPerKwhOk

`func (o *MonetaryCostOptions) GetCostPerKwhOk() (*float64, bool)`

GetCostPerKwhOk returns a tuple with the CostPerKwh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerKwh

`func (o *MonetaryCostOptions) SetCostPerKwh(v float64)`

SetCostPerKwh sets CostPerKwh field to given value.

### HasCostPerKwh

`func (o *MonetaryCostOptions) HasCostPerKwh() bool`

HasCostPerKwh returns a boolean if a field has been set.

### GetCostPerFuelUnit

`func (o *MonetaryCostOptions) GetCostPerFuelUnit() float64`

GetCostPerFuelUnit returns the CostPerFuelUnit field if non-nil, zero value otherwise.

### GetCostPerFuelUnitOk

`func (o *MonetaryCostOptions) GetCostPerFuelUnitOk() (*float64, bool)`

GetCostPerFuelUnitOk returns a tuple with the CostPerFuelUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerFuelUnit

`func (o *MonetaryCostOptions) SetCostPerFuelUnit(v float64)`

SetCostPerFuelUnit sets CostPerFuelUnit field to given value.

### HasCostPerFuelUnit

`func (o *MonetaryCostOptions) HasCostPerFuelUnit() bool`

HasCostPerFuelUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


