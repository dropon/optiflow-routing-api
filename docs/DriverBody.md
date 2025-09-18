# DriverBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkingHoursPreset** | [**WorkingHoursPreset**](WorkingHoursPreset.md) |  | 
**WorkLogbook** | Pointer to [**WorkLogbook**](WorkLogbook.md) |  | [optional] 
**UseTimeOnCombinedTransportForRecreation** | Pointer to **bool** | If true, the time on a combined transport, e.g. on a ferry, can be used for recreation.  That means that a break or a rest can be scheduled when traveling on a ferry or by rail. As breaks and daily rests may be split to match the time of the combined transport, we recommend to additionally request _COMBINED_TRANSPORT_EVENTS_ when _SCHEDULE_EVENTS_ or _SCHEDULE_EVENTS_WITH_DRIVING_ are requested. | [optional] [default to false]

## Methods

### NewDriverBody

`func NewDriverBody(workingHoursPreset WorkingHoursPreset, ) *DriverBody`

NewDriverBody instantiates a new DriverBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriverBodyWithDefaults

`func NewDriverBodyWithDefaults() *DriverBody`

NewDriverBodyWithDefaults instantiates a new DriverBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkingHoursPreset

`func (o *DriverBody) GetWorkingHoursPreset() WorkingHoursPreset`

GetWorkingHoursPreset returns the WorkingHoursPreset field if non-nil, zero value otherwise.

### GetWorkingHoursPresetOk

`func (o *DriverBody) GetWorkingHoursPresetOk() (*WorkingHoursPreset, bool)`

GetWorkingHoursPresetOk returns a tuple with the WorkingHoursPreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingHoursPreset

`func (o *DriverBody) SetWorkingHoursPreset(v WorkingHoursPreset)`

SetWorkingHoursPreset sets WorkingHoursPreset field to given value.


### GetWorkLogbook

`func (o *DriverBody) GetWorkLogbook() WorkLogbook`

GetWorkLogbook returns the WorkLogbook field if non-nil, zero value otherwise.

### GetWorkLogbookOk

`func (o *DriverBody) GetWorkLogbookOk() (*WorkLogbook, bool)`

GetWorkLogbookOk returns a tuple with the WorkLogbook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkLogbook

`func (o *DriverBody) SetWorkLogbook(v WorkLogbook)`

SetWorkLogbook sets WorkLogbook field to given value.

### HasWorkLogbook

`func (o *DriverBody) HasWorkLogbook() bool`

HasWorkLogbook returns a boolean if a field has been set.

### GetUseTimeOnCombinedTransportForRecreation

`func (o *DriverBody) GetUseTimeOnCombinedTransportForRecreation() bool`

GetUseTimeOnCombinedTransportForRecreation returns the UseTimeOnCombinedTransportForRecreation field if non-nil, zero value otherwise.

### GetUseTimeOnCombinedTransportForRecreationOk

`func (o *DriverBody) GetUseTimeOnCombinedTransportForRecreationOk() (*bool, bool)`

GetUseTimeOnCombinedTransportForRecreationOk returns a tuple with the UseTimeOnCombinedTransportForRecreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTimeOnCombinedTransportForRecreation

`func (o *DriverBody) SetUseTimeOnCombinedTransportForRecreation(v bool)`

SetUseTimeOnCombinedTransportForRecreation sets UseTimeOnCombinedTransportForRecreation field to given value.

### HasUseTimeOnCombinedTransportForRecreation

`func (o *DriverBody) HasUseTimeOnCombinedTransportForRecreation() bool`

HasUseTimeOnCombinedTransportForRecreation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


