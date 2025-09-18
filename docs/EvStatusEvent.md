# EvStatusEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatteryStateOfCharge** | **float64** | The remaining state of charge at the location of this event [%]. May be below the defined minimum state of charge or even below zero, if the electricity consumption exceeds the available energy in the battery and charging is not possible before falling below zero. | 
**ElectricityConsumption** | **float64** | The electricity consumption since the previous event containing **evStatus** [kWh]. | 
**Polyline** | Pointer to **string** | The route polyline snippet since the previous event containing **evStatus** in the format specified by **options[polylineFormat]**. The polyline may be null for events without distance to the previous **evStatus** event, for example events generated at a waypoint  with power consumption during service. Requires _EV_STATUS_EVENTS_POLYLINE_ to be requested. | [optional] 

## Methods

### NewEvStatusEvent

`func NewEvStatusEvent(batteryStateOfCharge float64, electricityConsumption float64, ) *EvStatusEvent`

NewEvStatusEvent instantiates a new EvStatusEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvStatusEventWithDefaults

`func NewEvStatusEventWithDefaults() *EvStatusEvent`

NewEvStatusEventWithDefaults instantiates a new EvStatusEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatteryStateOfCharge

`func (o *EvStatusEvent) GetBatteryStateOfCharge() float64`

GetBatteryStateOfCharge returns the BatteryStateOfCharge field if non-nil, zero value otherwise.

### GetBatteryStateOfChargeOk

`func (o *EvStatusEvent) GetBatteryStateOfChargeOk() (*float64, bool)`

GetBatteryStateOfChargeOk returns a tuple with the BatteryStateOfCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryStateOfCharge

`func (o *EvStatusEvent) SetBatteryStateOfCharge(v float64)`

SetBatteryStateOfCharge sets BatteryStateOfCharge field to given value.


### GetElectricityConsumption

`func (o *EvStatusEvent) GetElectricityConsumption() float64`

GetElectricityConsumption returns the ElectricityConsumption field if non-nil, zero value otherwise.

### GetElectricityConsumptionOk

`func (o *EvStatusEvent) GetElectricityConsumptionOk() (*float64, bool)`

GetElectricityConsumptionOk returns a tuple with the ElectricityConsumption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElectricityConsumption

`func (o *EvStatusEvent) SetElectricityConsumption(v float64)`

SetElectricityConsumption sets ElectricityConsumption field to given value.


### GetPolyline

`func (o *EvStatusEvent) GetPolyline() string`

GetPolyline returns the Polyline field if non-nil, zero value otherwise.

### GetPolylineOk

`func (o *EvStatusEvent) GetPolylineOk() (*string, bool)`

GetPolylineOk returns a tuple with the Polyline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolyline

`func (o *EvStatusEvent) SetPolyline(v string)`

SetPolyline sets Polyline field to given value.

### HasPolyline

`func (o *EvStatusEvent) HasPolyline() bool`

HasPolyline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


