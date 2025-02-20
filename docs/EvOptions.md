# EvOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InitialStateOfCharge** | Pointer to **float64** | The available battery capacity at the start of the route [%]. | [optional] [default to 100]
**MinimumStateOfCharge** | Pointer to **float64** | The minimum wanted remaining battery capacity at the end of the route and at the end of each leg [%]. | [optional] [default to 10]
**EnergyEfficientRoute** | Pointer to **bool** | Specifies if an energy efficient route should be calculated. | [optional] [default to false]

## Methods

### NewEvOptions

`func NewEvOptions() *EvOptions`

NewEvOptions instantiates a new EvOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvOptionsWithDefaults

`func NewEvOptionsWithDefaults() *EvOptions`

NewEvOptionsWithDefaults instantiates a new EvOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInitialStateOfCharge

`func (o *EvOptions) GetInitialStateOfCharge() float64`

GetInitialStateOfCharge returns the InitialStateOfCharge field if non-nil, zero value otherwise.

### GetInitialStateOfChargeOk

`func (o *EvOptions) GetInitialStateOfChargeOk() (*float64, bool)`

GetInitialStateOfChargeOk returns a tuple with the InitialStateOfCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialStateOfCharge

`func (o *EvOptions) SetInitialStateOfCharge(v float64)`

SetInitialStateOfCharge sets InitialStateOfCharge field to given value.

### HasInitialStateOfCharge

`func (o *EvOptions) HasInitialStateOfCharge() bool`

HasInitialStateOfCharge returns a boolean if a field has been set.

### GetMinimumStateOfCharge

`func (o *EvOptions) GetMinimumStateOfCharge() float64`

GetMinimumStateOfCharge returns the MinimumStateOfCharge field if non-nil, zero value otherwise.

### GetMinimumStateOfChargeOk

`func (o *EvOptions) GetMinimumStateOfChargeOk() (*float64, bool)`

GetMinimumStateOfChargeOk returns a tuple with the MinimumStateOfCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumStateOfCharge

`func (o *EvOptions) SetMinimumStateOfCharge(v float64)`

SetMinimumStateOfCharge sets MinimumStateOfCharge field to given value.

### HasMinimumStateOfCharge

`func (o *EvOptions) HasMinimumStateOfCharge() bool`

HasMinimumStateOfCharge returns a boolean if a field has been set.

### GetEnergyEfficientRoute

`func (o *EvOptions) GetEnergyEfficientRoute() bool`

GetEnergyEfficientRoute returns the EnergyEfficientRoute field if non-nil, zero value otherwise.

### GetEnergyEfficientRouteOk

`func (o *EvOptions) GetEnergyEfficientRouteOk() (*bool, bool)`

GetEnergyEfficientRouteOk returns a tuple with the EnergyEfficientRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnergyEfficientRoute

`func (o *EvOptions) SetEnergyEfficientRoute(v bool)`

SetEnergyEfficientRoute sets EnergyEfficientRoute field to given value.

### HasEnergyEfficientRoute

`func (o *EvOptions) HasEnergyEfficientRoute() bool`

HasEnergyEfficientRoute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


