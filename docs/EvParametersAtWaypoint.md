# EvParametersAtWaypoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Weather** | Pointer to [**Weather**](Weather.md) |  | [optional] 
**ChargingStation** | Pointer to [**ChargingStation**](ChargingStation.md) |  | [optional] 
**PowerConsumptionDuringService** | Pointer to **float64** | The electricity consumed during service by electrical appliances (e.g. crane, cooling, tail lift) [kWh].  | [optional] [default to 0]

## Methods

### NewEvParametersAtWaypoint

`func NewEvParametersAtWaypoint() *EvParametersAtWaypoint`

NewEvParametersAtWaypoint instantiates a new EvParametersAtWaypoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvParametersAtWaypointWithDefaults

`func NewEvParametersAtWaypointWithDefaults() *EvParametersAtWaypoint`

NewEvParametersAtWaypointWithDefaults instantiates a new EvParametersAtWaypoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWeather

`func (o *EvParametersAtWaypoint) GetWeather() Weather`

GetWeather returns the Weather field if non-nil, zero value otherwise.

### GetWeatherOk

`func (o *EvParametersAtWaypoint) GetWeatherOk() (*Weather, bool)`

GetWeatherOk returns a tuple with the Weather field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeather

`func (o *EvParametersAtWaypoint) SetWeather(v Weather)`

SetWeather sets Weather field to given value.

### HasWeather

`func (o *EvParametersAtWaypoint) HasWeather() bool`

HasWeather returns a boolean if a field has been set.

### GetChargingStation

`func (o *EvParametersAtWaypoint) GetChargingStation() ChargingStation`

GetChargingStation returns the ChargingStation field if non-nil, zero value otherwise.

### GetChargingStationOk

`func (o *EvParametersAtWaypoint) GetChargingStationOk() (*ChargingStation, bool)`

GetChargingStationOk returns a tuple with the ChargingStation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargingStation

`func (o *EvParametersAtWaypoint) SetChargingStation(v ChargingStation)`

SetChargingStation sets ChargingStation field to given value.

### HasChargingStation

`func (o *EvParametersAtWaypoint) HasChargingStation() bool`

HasChargingStation returns a boolean if a field has been set.

### GetPowerConsumptionDuringService

`func (o *EvParametersAtWaypoint) GetPowerConsumptionDuringService() float64`

GetPowerConsumptionDuringService returns the PowerConsumptionDuringService field if non-nil, zero value otherwise.

### GetPowerConsumptionDuringServiceOk

`func (o *EvParametersAtWaypoint) GetPowerConsumptionDuringServiceOk() (*float64, bool)`

GetPowerConsumptionDuringServiceOk returns a tuple with the PowerConsumptionDuringService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerConsumptionDuringService

`func (o *EvParametersAtWaypoint) SetPowerConsumptionDuringService(v float64)`

SetPowerConsumptionDuringService sets PowerConsumptionDuringService field to given value.

### HasPowerConsumptionDuringService

`func (o *EvParametersAtWaypoint) HasPowerConsumptionDuringService() bool`

HasPowerConsumptionDuringService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


