# WeatherResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Temperature** | Pointer to **float64** | The average temperature at 2 meters above ground [°C]. Used to estimate the state of health of the  battery and consumption due to air conditioning. | [optional] 
**WindSpeed** | Pointer to **int32** | The average wind speed at 10 meters above ground [km/h]. | [optional] 
**WindDirection** | Pointer to **int32** | The wind direction (clockwise) at 10 meters above ground. North represents 0 degrees. | [optional] 

## Methods

### NewWeatherResponse

`func NewWeatherResponse() *WeatherResponse`

NewWeatherResponse instantiates a new WeatherResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWeatherResponseWithDefaults

`func NewWeatherResponseWithDefaults() *WeatherResponse`

NewWeatherResponseWithDefaults instantiates a new WeatherResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemperature

`func (o *WeatherResponse) GetTemperature() float64`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *WeatherResponse) GetTemperatureOk() (*float64, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *WeatherResponse) SetTemperature(v float64)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *WeatherResponse) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetWindSpeed

`func (o *WeatherResponse) GetWindSpeed() int32`

GetWindSpeed returns the WindSpeed field if non-nil, zero value otherwise.

### GetWindSpeedOk

`func (o *WeatherResponse) GetWindSpeedOk() (*int32, bool)`

GetWindSpeedOk returns a tuple with the WindSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindSpeed

`func (o *WeatherResponse) SetWindSpeed(v int32)`

SetWindSpeed sets WindSpeed field to given value.

### HasWindSpeed

`func (o *WeatherResponse) HasWindSpeed() bool`

HasWindSpeed returns a boolean if a field has been set.

### GetWindDirection

`func (o *WeatherResponse) GetWindDirection() int32`

GetWindDirection returns the WindDirection field if non-nil, zero value otherwise.

### GetWindDirectionOk

`func (o *WeatherResponse) GetWindDirectionOk() (*int32, bool)`

GetWindDirectionOk returns a tuple with the WindDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindDirection

`func (o *WeatherResponse) SetWindDirection(v int32)`

SetWindDirection sets WindDirection field to given value.

### HasWindDirection

`func (o *WeatherResponse) HasWindDirection() bool`

HasWindDirection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


