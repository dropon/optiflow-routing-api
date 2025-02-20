# WorkLogbook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastTimeTheDriverWorked** | **time.Time** | The last time the driver worked formatted according to [RFC 3339](https://tools.ietf.org/html/rfc3339).    The date must not be before 1970-01-01T00:00:00+00:00 nor after 2037-12-31T23:59:59+00:00.  The date must provide an offset to UTC.   See [here](./concepts/date-and-time) for more information on the relevance of date and time. | 
**AccumulatedDrivingTimeSinceLastBreak** | Pointer to **int32** | Accumulated driving time since end of last break [s], this includes all time behind the wheel.   Values higher than the maximum driving time between breaks of the **workingHoursPreset** are capped and effectively equal to the value given by the preset.  | [optional] [default to 0]
**AccumulatedWorkingTimeSinceLastBreak** | Pointer to **int32** | Accumulated working time since end of last break [s], this includes driving time, service time at depot and customers, and idle time if lower than working time threshold.     Values higher than the maximum working time between breaks of the **workingHoursPreset** are capped and effectively equal to the value given by the preset. May not be smaller than **accumulatedDrivingTimeSinceLastBreak** if the value is specified. If no other value is specified, the default value is 0.  | [optional] 
**AccumulatedDrivingTimeSinceLastDailyRest** | Pointer to **int32** | Accumulated driving time since end of last daily rest [s], this includes all time behind the wheel.   Values higher than the maximum driving time between daily rests of the **workingHoursPreset** are capped and effectively equal to the value given by preset. May not be smaller than **accumulatedDrivingTimeSinceLastBreak** if the value is specified. If no other value is specified, the default value is 0.  | [optional] 
**AccumulatedTravelTimeSinceLastDailyRest** | Pointer to **int32** | Accumulated travel time since end of last daily rest [s], this includes all time since the last daily rest.   Values higher than maximum travel time between daily rests of the **workingHoursPreset** are capped and effectively equal to the value given by the preset. May not be smaller than **accumulatedWorkingTimeSinceLastBreak** or **accumulatedDrivingTimeSinceLastDailyRest** if the values are specified. If no other value is specified, the default value is 0.  | [optional] 

## Methods

### NewWorkLogbook

`func NewWorkLogbook(lastTimeTheDriverWorked time.Time, ) *WorkLogbook`

NewWorkLogbook instantiates a new WorkLogbook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkLogbookWithDefaults

`func NewWorkLogbookWithDefaults() *WorkLogbook`

NewWorkLogbookWithDefaults instantiates a new WorkLogbook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastTimeTheDriverWorked

`func (o *WorkLogbook) GetLastTimeTheDriverWorked() time.Time`

GetLastTimeTheDriverWorked returns the LastTimeTheDriverWorked field if non-nil, zero value otherwise.

### GetLastTimeTheDriverWorkedOk

`func (o *WorkLogbook) GetLastTimeTheDriverWorkedOk() (*time.Time, bool)`

GetLastTimeTheDriverWorkedOk returns a tuple with the LastTimeTheDriverWorked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTimeTheDriverWorked

`func (o *WorkLogbook) SetLastTimeTheDriverWorked(v time.Time)`

SetLastTimeTheDriverWorked sets LastTimeTheDriverWorked field to given value.


### GetAccumulatedDrivingTimeSinceLastBreak

`func (o *WorkLogbook) GetAccumulatedDrivingTimeSinceLastBreak() int32`

GetAccumulatedDrivingTimeSinceLastBreak returns the AccumulatedDrivingTimeSinceLastBreak field if non-nil, zero value otherwise.

### GetAccumulatedDrivingTimeSinceLastBreakOk

`func (o *WorkLogbook) GetAccumulatedDrivingTimeSinceLastBreakOk() (*int32, bool)`

GetAccumulatedDrivingTimeSinceLastBreakOk returns a tuple with the AccumulatedDrivingTimeSinceLastBreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccumulatedDrivingTimeSinceLastBreak

`func (o *WorkLogbook) SetAccumulatedDrivingTimeSinceLastBreak(v int32)`

SetAccumulatedDrivingTimeSinceLastBreak sets AccumulatedDrivingTimeSinceLastBreak field to given value.

### HasAccumulatedDrivingTimeSinceLastBreak

`func (o *WorkLogbook) HasAccumulatedDrivingTimeSinceLastBreak() bool`

HasAccumulatedDrivingTimeSinceLastBreak returns a boolean if a field has been set.

### GetAccumulatedWorkingTimeSinceLastBreak

`func (o *WorkLogbook) GetAccumulatedWorkingTimeSinceLastBreak() int32`

GetAccumulatedWorkingTimeSinceLastBreak returns the AccumulatedWorkingTimeSinceLastBreak field if non-nil, zero value otherwise.

### GetAccumulatedWorkingTimeSinceLastBreakOk

`func (o *WorkLogbook) GetAccumulatedWorkingTimeSinceLastBreakOk() (*int32, bool)`

GetAccumulatedWorkingTimeSinceLastBreakOk returns a tuple with the AccumulatedWorkingTimeSinceLastBreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccumulatedWorkingTimeSinceLastBreak

`func (o *WorkLogbook) SetAccumulatedWorkingTimeSinceLastBreak(v int32)`

SetAccumulatedWorkingTimeSinceLastBreak sets AccumulatedWorkingTimeSinceLastBreak field to given value.

### HasAccumulatedWorkingTimeSinceLastBreak

`func (o *WorkLogbook) HasAccumulatedWorkingTimeSinceLastBreak() bool`

HasAccumulatedWorkingTimeSinceLastBreak returns a boolean if a field has been set.

### GetAccumulatedDrivingTimeSinceLastDailyRest

`func (o *WorkLogbook) GetAccumulatedDrivingTimeSinceLastDailyRest() int32`

GetAccumulatedDrivingTimeSinceLastDailyRest returns the AccumulatedDrivingTimeSinceLastDailyRest field if non-nil, zero value otherwise.

### GetAccumulatedDrivingTimeSinceLastDailyRestOk

`func (o *WorkLogbook) GetAccumulatedDrivingTimeSinceLastDailyRestOk() (*int32, bool)`

GetAccumulatedDrivingTimeSinceLastDailyRestOk returns a tuple with the AccumulatedDrivingTimeSinceLastDailyRest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccumulatedDrivingTimeSinceLastDailyRest

`func (o *WorkLogbook) SetAccumulatedDrivingTimeSinceLastDailyRest(v int32)`

SetAccumulatedDrivingTimeSinceLastDailyRest sets AccumulatedDrivingTimeSinceLastDailyRest field to given value.

### HasAccumulatedDrivingTimeSinceLastDailyRest

`func (o *WorkLogbook) HasAccumulatedDrivingTimeSinceLastDailyRest() bool`

HasAccumulatedDrivingTimeSinceLastDailyRest returns a boolean if a field has been set.

### GetAccumulatedTravelTimeSinceLastDailyRest

`func (o *WorkLogbook) GetAccumulatedTravelTimeSinceLastDailyRest() int32`

GetAccumulatedTravelTimeSinceLastDailyRest returns the AccumulatedTravelTimeSinceLastDailyRest field if non-nil, zero value otherwise.

### GetAccumulatedTravelTimeSinceLastDailyRestOk

`func (o *WorkLogbook) GetAccumulatedTravelTimeSinceLastDailyRestOk() (*int32, bool)`

GetAccumulatedTravelTimeSinceLastDailyRestOk returns a tuple with the AccumulatedTravelTimeSinceLastDailyRest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccumulatedTravelTimeSinceLastDailyRest

`func (o *WorkLogbook) SetAccumulatedTravelTimeSinceLastDailyRest(v int32)`

SetAccumulatedTravelTimeSinceLastDailyRest sets AccumulatedTravelTimeSinceLastDailyRest field to given value.

### HasAccumulatedTravelTimeSinceLastDailyRest

`func (o *WorkLogbook) HasAccumulatedTravelTimeSinceLastDailyRest() bool`

HasAccumulatedTravelTimeSinceLastDailyRest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


