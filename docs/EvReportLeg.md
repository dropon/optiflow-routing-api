# EvReportLeg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ElectricityConsumption** | **float64** | The electricity consumption since the start of the leg [kWh]. | 
**BatteryStateOfCharge** | **float64** | The remaining state of charge at the end of the leg [%]. May be below the defined minimum state of charge or even below zero, if the electricity consumption exceeds the available energy in the battery and charging is not possible before falling below zero. | 
**WeatherAtStart** | [**WeatherResponse**](WeatherResponse.md) |  | 
**WeatherAtEnd** | [**WeatherResponse**](WeatherResponse.md) |  | 

## Methods

### NewEvReportLeg

`func NewEvReportLeg(electricityConsumption float64, batteryStateOfCharge float64, weatherAtStart WeatherResponse, weatherAtEnd WeatherResponse, ) *EvReportLeg`

NewEvReportLeg instantiates a new EvReportLeg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvReportLegWithDefaults

`func NewEvReportLegWithDefaults() *EvReportLeg`

NewEvReportLegWithDefaults instantiates a new EvReportLeg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElectricityConsumption

`func (o *EvReportLeg) GetElectricityConsumption() float64`

GetElectricityConsumption returns the ElectricityConsumption field if non-nil, zero value otherwise.

### GetElectricityConsumptionOk

`func (o *EvReportLeg) GetElectricityConsumptionOk() (*float64, bool)`

GetElectricityConsumptionOk returns a tuple with the ElectricityConsumption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElectricityConsumption

`func (o *EvReportLeg) SetElectricityConsumption(v float64)`

SetElectricityConsumption sets ElectricityConsumption field to given value.


### GetBatteryStateOfCharge

`func (o *EvReportLeg) GetBatteryStateOfCharge() float64`

GetBatteryStateOfCharge returns the BatteryStateOfCharge field if non-nil, zero value otherwise.

### GetBatteryStateOfChargeOk

`func (o *EvReportLeg) GetBatteryStateOfChargeOk() (*float64, bool)`

GetBatteryStateOfChargeOk returns a tuple with the BatteryStateOfCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryStateOfCharge

`func (o *EvReportLeg) SetBatteryStateOfCharge(v float64)`

SetBatteryStateOfCharge sets BatteryStateOfCharge field to given value.


### GetWeatherAtStart

`func (o *EvReportLeg) GetWeatherAtStart() WeatherResponse`

GetWeatherAtStart returns the WeatherAtStart field if non-nil, zero value otherwise.

### GetWeatherAtStartOk

`func (o *EvReportLeg) GetWeatherAtStartOk() (*WeatherResponse, bool)`

GetWeatherAtStartOk returns a tuple with the WeatherAtStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeatherAtStart

`func (o *EvReportLeg) SetWeatherAtStart(v WeatherResponse)`

SetWeatherAtStart sets WeatherAtStart field to given value.


### GetWeatherAtEnd

`func (o *EvReportLeg) GetWeatherAtEnd() WeatherResponse`

GetWeatherAtEnd returns the WeatherAtEnd field if non-nil, zero value otherwise.

### GetWeatherAtEndOk

`func (o *EvReportLeg) GetWeatherAtEndOk() (*WeatherResponse, bool)`

GetWeatherAtEndOk returns a tuple with the WeatherAtEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeatherAtEnd

`func (o *EvReportLeg) SetWeatherAtEnd(v WeatherResponse)`

SetWeatherAtEnd sets WeatherAtEnd field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


