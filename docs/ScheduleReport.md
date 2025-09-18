# ScheduleReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | Pointer to **time.Time** | The start time of the route formatted according to [RFC 3339](https://tools.ietf.org/html/rfc3339). Only present with **options[trafficMode]** _REALISTIC_ or when **options[startTime]** is specified. | [optional] 
**EndTime** | Pointer to **time.Time** | The end time of the route formatted according to [RFC 3339](https://tools.ietf.org/html/rfc3339). Only present with **options[trafficMode]** _REALISTIC_ or when **options[startTime]** is specified. | [optional] 
**DrivingTime** | **int32** | The total driving time of the route [s]. Time which is not service, waiting, break or rest is considered as driving. | 
**ServiceTime** | **int32** | The total service time of the route [s], equal to the sum of **schedule** events of type _SERVICE_. | 
**WaitingTime** | **int32** | The total waiting time of the route [s], equal to the sum of **schedule** events of type _WAITING_ which are not _SERVICE_ at the same time. | 
**BreakTime** | **int32** | The total break time of the route [s], equal to the sum of **schedule** events of type _BREAK_ which are not _SERVICE_ at the same time. | 
**RestTime** | **int32** | The total rest time of the route [s], equal to the sum of **schedule** events of type _DAILY_REST_ which are not _SERVICE_ at the same time. | 

## Methods

### NewScheduleReport

`func NewScheduleReport(drivingTime int32, serviceTime int32, waitingTime int32, breakTime int32, restTime int32, ) *ScheduleReport`

NewScheduleReport instantiates a new ScheduleReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleReportWithDefaults

`func NewScheduleReportWithDefaults() *ScheduleReport`

NewScheduleReportWithDefaults instantiates a new ScheduleReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *ScheduleReport) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ScheduleReport) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ScheduleReport) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ScheduleReport) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *ScheduleReport) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ScheduleReport) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ScheduleReport) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ScheduleReport) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetDrivingTime

`func (o *ScheduleReport) GetDrivingTime() int32`

GetDrivingTime returns the DrivingTime field if non-nil, zero value otherwise.

### GetDrivingTimeOk

`func (o *ScheduleReport) GetDrivingTimeOk() (*int32, bool)`

GetDrivingTimeOk returns a tuple with the DrivingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivingTime

`func (o *ScheduleReport) SetDrivingTime(v int32)`

SetDrivingTime sets DrivingTime field to given value.


### GetServiceTime

`func (o *ScheduleReport) GetServiceTime() int32`

GetServiceTime returns the ServiceTime field if non-nil, zero value otherwise.

### GetServiceTimeOk

`func (o *ScheduleReport) GetServiceTimeOk() (*int32, bool)`

GetServiceTimeOk returns a tuple with the ServiceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTime

`func (o *ScheduleReport) SetServiceTime(v int32)`

SetServiceTime sets ServiceTime field to given value.


### GetWaitingTime

`func (o *ScheduleReport) GetWaitingTime() int32`

GetWaitingTime returns the WaitingTime field if non-nil, zero value otherwise.

### GetWaitingTimeOk

`func (o *ScheduleReport) GetWaitingTimeOk() (*int32, bool)`

GetWaitingTimeOk returns a tuple with the WaitingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitingTime

`func (o *ScheduleReport) SetWaitingTime(v int32)`

SetWaitingTime sets WaitingTime field to given value.


### GetBreakTime

`func (o *ScheduleReport) GetBreakTime() int32`

GetBreakTime returns the BreakTime field if non-nil, zero value otherwise.

### GetBreakTimeOk

`func (o *ScheduleReport) GetBreakTimeOk() (*int32, bool)`

GetBreakTimeOk returns a tuple with the BreakTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakTime

`func (o *ScheduleReport) SetBreakTime(v int32)`

SetBreakTime sets BreakTime field to given value.


### GetRestTime

`func (o *ScheduleReport) GetRestTime() int32`

GetRestTime returns the RestTime field if non-nil, zero value otherwise.

### GetRestTimeOk

`func (o *ScheduleReport) GetRestTimeOk() (*int32, bool)`

GetRestTimeOk returns a tuple with the RestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestTime

`func (o *ScheduleReport) SetRestTime(v int32)`

SetRestTime sets RestTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


