# RemainingWaypoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of this waypoint. | 
**EstimatedTimeOfArrival** | **time.Time** | The estimated time of arrival at this waypoint formatted according to [RFC 3339](https://tools.ietf.org/html/rfc3339). Does not include service, break, rest and waiting at this waypoint. | 
**Distance** | **int32** | The distance from the current position to this waypoint [m]. | 
**TravelTime** | **int32** | The travel time from the current position to this waypoint [s]. Does not include service, break, rest and waiting at this waypoint. | 
**TrafficDelay** | **int32** | The delay due to live traffic from the current position to this waypoint [s]. | 
**WaitingTime** | **int32** | The waiting time at this waypoint [s]. | 
**ScheduleViolations** | [**[]ScheduleViolationType**](ScheduleViolationType.md) | Contains the list of violated schedule restrictions at this waypoint. | 
**Violated** | **bool** | If there is no valid route for the given vehicle between the vehicle position and this waypoint, but the resulting route can be calculated using roads actually prohibited, the route is marked as violated. | 

## Methods

### NewRemainingWaypoint

`func NewRemainingWaypoint(name string, estimatedTimeOfArrival time.Time, distance int32, travelTime int32, trafficDelay int32, waitingTime int32, scheduleViolations []ScheduleViolationType, violated bool, ) *RemainingWaypoint`

NewRemainingWaypoint instantiates a new RemainingWaypoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemainingWaypointWithDefaults

`func NewRemainingWaypointWithDefaults() *RemainingWaypoint`

NewRemainingWaypointWithDefaults instantiates a new RemainingWaypoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RemainingWaypoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RemainingWaypoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RemainingWaypoint) SetName(v string)`

SetName sets Name field to given value.


### GetEstimatedTimeOfArrival

`func (o *RemainingWaypoint) GetEstimatedTimeOfArrival() time.Time`

GetEstimatedTimeOfArrival returns the EstimatedTimeOfArrival field if non-nil, zero value otherwise.

### GetEstimatedTimeOfArrivalOk

`func (o *RemainingWaypoint) GetEstimatedTimeOfArrivalOk() (*time.Time, bool)`

GetEstimatedTimeOfArrivalOk returns a tuple with the EstimatedTimeOfArrival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTimeOfArrival

`func (o *RemainingWaypoint) SetEstimatedTimeOfArrival(v time.Time)`

SetEstimatedTimeOfArrival sets EstimatedTimeOfArrival field to given value.


### GetDistance

`func (o *RemainingWaypoint) GetDistance() int32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *RemainingWaypoint) GetDistanceOk() (*int32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *RemainingWaypoint) SetDistance(v int32)`

SetDistance sets Distance field to given value.


### GetTravelTime

`func (o *RemainingWaypoint) GetTravelTime() int32`

GetTravelTime returns the TravelTime field if non-nil, zero value otherwise.

### GetTravelTimeOk

`func (o *RemainingWaypoint) GetTravelTimeOk() (*int32, bool)`

GetTravelTimeOk returns a tuple with the TravelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelTime

`func (o *RemainingWaypoint) SetTravelTime(v int32)`

SetTravelTime sets TravelTime field to given value.


### GetTrafficDelay

`func (o *RemainingWaypoint) GetTrafficDelay() int32`

GetTrafficDelay returns the TrafficDelay field if non-nil, zero value otherwise.

### GetTrafficDelayOk

`func (o *RemainingWaypoint) GetTrafficDelayOk() (*int32, bool)`

GetTrafficDelayOk returns a tuple with the TrafficDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficDelay

`func (o *RemainingWaypoint) SetTrafficDelay(v int32)`

SetTrafficDelay sets TrafficDelay field to given value.


### GetWaitingTime

`func (o *RemainingWaypoint) GetWaitingTime() int32`

GetWaitingTime returns the WaitingTime field if non-nil, zero value otherwise.

### GetWaitingTimeOk

`func (o *RemainingWaypoint) GetWaitingTimeOk() (*int32, bool)`

GetWaitingTimeOk returns a tuple with the WaitingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitingTime

`func (o *RemainingWaypoint) SetWaitingTime(v int32)`

SetWaitingTime sets WaitingTime field to given value.


### GetScheduleViolations

`func (o *RemainingWaypoint) GetScheduleViolations() []ScheduleViolationType`

GetScheduleViolations returns the ScheduleViolations field if non-nil, zero value otherwise.

### GetScheduleViolationsOk

`func (o *RemainingWaypoint) GetScheduleViolationsOk() (*[]ScheduleViolationType, bool)`

GetScheduleViolationsOk returns a tuple with the ScheduleViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleViolations

`func (o *RemainingWaypoint) SetScheduleViolations(v []ScheduleViolationType)`

SetScheduleViolations sets ScheduleViolations field to given value.


### GetViolated

`func (o *RemainingWaypoint) GetViolated() bool`

GetViolated returns the Violated field if non-nil, zero value otherwise.

### GetViolatedOk

`func (o *RemainingWaypoint) GetViolatedOk() (*bool, bool)`

GetViolatedOk returns a tuple with the Violated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolated

`func (o *RemainingWaypoint) SetViolated(v bool)`

SetViolated sets Violated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


