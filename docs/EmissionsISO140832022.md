# EmissionsISO140832022

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FuelConsumption** | **float64** | The total conventional fuel consumption [kg]. | 
**ElectricityConsumption** | **float64** | The total electric power consumption of the vehicle if the **engineType** is _ELECTRIC_ or _HYBRID_ [kWh]. | 
**Co2eTankToWheel** | **float64** | The amount of emitted CO2e from tank to wheel [kg]. | 
**Co2eWellToWheel** | **float64** | The amount of emitted CO2e from well to wheel [kg]. | 
**EnergyUseTankToWheel** | **float64** | The tank-to-wheel energy use [MJ]. | 
**EnergyUseWellToWheel** | **float64** | The well-to-wheel energy use [MJ]. | 

## Methods

### NewEmissionsISO140832022

`func NewEmissionsISO140832022(fuelConsumption float64, electricityConsumption float64, co2eTankToWheel float64, co2eWellToWheel float64, energyUseTankToWheel float64, energyUseWellToWheel float64, ) *EmissionsISO140832022`

NewEmissionsISO140832022 instantiates a new EmissionsISO140832022 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmissionsISO140832022WithDefaults

`func NewEmissionsISO140832022WithDefaults() *EmissionsISO140832022`

NewEmissionsISO140832022WithDefaults instantiates a new EmissionsISO140832022 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFuelConsumption

`func (o *EmissionsISO140832022) GetFuelConsumption() float64`

GetFuelConsumption returns the FuelConsumption field if non-nil, zero value otherwise.

### GetFuelConsumptionOk

`func (o *EmissionsISO140832022) GetFuelConsumptionOk() (*float64, bool)`

GetFuelConsumptionOk returns a tuple with the FuelConsumption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelConsumption

`func (o *EmissionsISO140832022) SetFuelConsumption(v float64)`

SetFuelConsumption sets FuelConsumption field to given value.


### GetElectricityConsumption

`func (o *EmissionsISO140832022) GetElectricityConsumption() float64`

GetElectricityConsumption returns the ElectricityConsumption field if non-nil, zero value otherwise.

### GetElectricityConsumptionOk

`func (o *EmissionsISO140832022) GetElectricityConsumptionOk() (*float64, bool)`

GetElectricityConsumptionOk returns a tuple with the ElectricityConsumption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElectricityConsumption

`func (o *EmissionsISO140832022) SetElectricityConsumption(v float64)`

SetElectricityConsumption sets ElectricityConsumption field to given value.


### GetCo2eTankToWheel

`func (o *EmissionsISO140832022) GetCo2eTankToWheel() float64`

GetCo2eTankToWheel returns the Co2eTankToWheel field if non-nil, zero value otherwise.

### GetCo2eTankToWheelOk

`func (o *EmissionsISO140832022) GetCo2eTankToWheelOk() (*float64, bool)`

GetCo2eTankToWheelOk returns a tuple with the Co2eTankToWheel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCo2eTankToWheel

`func (o *EmissionsISO140832022) SetCo2eTankToWheel(v float64)`

SetCo2eTankToWheel sets Co2eTankToWheel field to given value.


### GetCo2eWellToWheel

`func (o *EmissionsISO140832022) GetCo2eWellToWheel() float64`

GetCo2eWellToWheel returns the Co2eWellToWheel field if non-nil, zero value otherwise.

### GetCo2eWellToWheelOk

`func (o *EmissionsISO140832022) GetCo2eWellToWheelOk() (*float64, bool)`

GetCo2eWellToWheelOk returns a tuple with the Co2eWellToWheel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCo2eWellToWheel

`func (o *EmissionsISO140832022) SetCo2eWellToWheel(v float64)`

SetCo2eWellToWheel sets Co2eWellToWheel field to given value.


### GetEnergyUseTankToWheel

`func (o *EmissionsISO140832022) GetEnergyUseTankToWheel() float64`

GetEnergyUseTankToWheel returns the EnergyUseTankToWheel field if non-nil, zero value otherwise.

### GetEnergyUseTankToWheelOk

`func (o *EmissionsISO140832022) GetEnergyUseTankToWheelOk() (*float64, bool)`

GetEnergyUseTankToWheelOk returns a tuple with the EnergyUseTankToWheel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnergyUseTankToWheel

`func (o *EmissionsISO140832022) SetEnergyUseTankToWheel(v float64)`

SetEnergyUseTankToWheel sets EnergyUseTankToWheel field to given value.


### GetEnergyUseWellToWheel

`func (o *EmissionsISO140832022) GetEnergyUseWellToWheel() float64`

GetEnergyUseWellToWheel returns the EnergyUseWellToWheel field if non-nil, zero value otherwise.

### GetEnergyUseWellToWheelOk

`func (o *EmissionsISO140832022) GetEnergyUseWellToWheelOk() (*float64, bool)`

GetEnergyUseWellToWheelOk returns a tuple with the EnergyUseWellToWheel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnergyUseWellToWheel

`func (o *EmissionsISO140832022) SetEnergyUseWellToWheel(v float64)`

SetEnergyUseWellToWheel sets EnergyUseWellToWheel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


