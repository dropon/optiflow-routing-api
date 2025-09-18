# ManeuverEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ManeuverType**](ManeuverType.md) |  | 
**RelativeDirection** | Pointer to **int32** | The direction of the outgoing road relative to continuing in the same direction as the incoming road (clockwise). | [optional] 
**AbsoluteDirection** | Pointer to **int32** | The absolute direction of the outgoing road (clockwise). North represents 0 degrees. | [optional] 
**Description** | **string** | A descriptive text for the current maneuver. The language can be specified by the parameter **options[language]**. A warning with **warningCode** _ROUTING_MANEUVERS_IN_DIFFERENT_LANGUAGE_ and the actual language is returned when the requested language is not available. Geographical names such as town and road names are always given in the local language. | 
**RoadAhead** | Pointer to [**RoadAhead**](RoadAhead.md) |  | [optional] 
**DirectionSignText** | Pointer to **string** | The city names and road numbers on a signpost at the current location to follow for the current maneuver. Empty if no signpost is present or the data is not available. | [optional] 
**ExitNumber** | Pointer to **string** | The number of an exit or interchange of a highway or a freeway-like road. Only present if the maneuver type is _CHANGE_ or _EXIT_. Empty if the data does not contain an exit number. | [optional] 
**ExitName** | Pointer to **string** | The name of an exit or interchange of a highway or a freeway-like road. Only present if the maneuver type is _CHANGE_ or _EXIT_. Empty if the data does not contain an exit name. | [optional] 
**RoundaboutExit** | Pointer to **int32** | The exit number at a roundabout. Only drivable roads are counted. Only present if the maneuver type is _TAKE\\_ROUNDABOUT_. | [optional] 
**CombinedTransportName** | Pointer to **string** | The name of the combined transport to take a the current location. Only present if the maneuver type is _TAKE\\_COMBINED\\_TRANSPORT_. | [optional] 
**CombinedTransportType** | Pointer to [**CombinedTransportType**](CombinedTransportType.md) |  | [optional] 
**CrossingRoadName** | Pointer to **string** | The name of the crossing road at which a U-turn has to be made. Only present if the maneuver type is _MAKE\\_U\\_TURN_ and if the U-turn takes place at a crossing. | [optional] 

## Methods

### NewManeuverEvent

`func NewManeuverEvent(type_ ManeuverType, description string, ) *ManeuverEvent`

NewManeuverEvent instantiates a new ManeuverEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManeuverEventWithDefaults

`func NewManeuverEventWithDefaults() *ManeuverEvent`

NewManeuverEventWithDefaults instantiates a new ManeuverEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ManeuverEvent) GetType() ManeuverType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManeuverEvent) GetTypeOk() (*ManeuverType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManeuverEvent) SetType(v ManeuverType)`

SetType sets Type field to given value.


### GetRelativeDirection

`func (o *ManeuverEvent) GetRelativeDirection() int32`

GetRelativeDirection returns the RelativeDirection field if non-nil, zero value otherwise.

### GetRelativeDirectionOk

`func (o *ManeuverEvent) GetRelativeDirectionOk() (*int32, bool)`

GetRelativeDirectionOk returns a tuple with the RelativeDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeDirection

`func (o *ManeuverEvent) SetRelativeDirection(v int32)`

SetRelativeDirection sets RelativeDirection field to given value.

### HasRelativeDirection

`func (o *ManeuverEvent) HasRelativeDirection() bool`

HasRelativeDirection returns a boolean if a field has been set.

### GetAbsoluteDirection

`func (o *ManeuverEvent) GetAbsoluteDirection() int32`

GetAbsoluteDirection returns the AbsoluteDirection field if non-nil, zero value otherwise.

### GetAbsoluteDirectionOk

`func (o *ManeuverEvent) GetAbsoluteDirectionOk() (*int32, bool)`

GetAbsoluteDirectionOk returns a tuple with the AbsoluteDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteDirection

`func (o *ManeuverEvent) SetAbsoluteDirection(v int32)`

SetAbsoluteDirection sets AbsoluteDirection field to given value.

### HasAbsoluteDirection

`func (o *ManeuverEvent) HasAbsoluteDirection() bool`

HasAbsoluteDirection returns a boolean if a field has been set.

### GetDescription

`func (o *ManeuverEvent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManeuverEvent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManeuverEvent) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetRoadAhead

`func (o *ManeuverEvent) GetRoadAhead() RoadAhead`

GetRoadAhead returns the RoadAhead field if non-nil, zero value otherwise.

### GetRoadAheadOk

`func (o *ManeuverEvent) GetRoadAheadOk() (*RoadAhead, bool)`

GetRoadAheadOk returns a tuple with the RoadAhead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoadAhead

`func (o *ManeuverEvent) SetRoadAhead(v RoadAhead)`

SetRoadAhead sets RoadAhead field to given value.

### HasRoadAhead

`func (o *ManeuverEvent) HasRoadAhead() bool`

HasRoadAhead returns a boolean if a field has been set.

### GetDirectionSignText

`func (o *ManeuverEvent) GetDirectionSignText() string`

GetDirectionSignText returns the DirectionSignText field if non-nil, zero value otherwise.

### GetDirectionSignTextOk

`func (o *ManeuverEvent) GetDirectionSignTextOk() (*string, bool)`

GetDirectionSignTextOk returns a tuple with the DirectionSignText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectionSignText

`func (o *ManeuverEvent) SetDirectionSignText(v string)`

SetDirectionSignText sets DirectionSignText field to given value.

### HasDirectionSignText

`func (o *ManeuverEvent) HasDirectionSignText() bool`

HasDirectionSignText returns a boolean if a field has been set.

### GetExitNumber

`func (o *ManeuverEvent) GetExitNumber() string`

GetExitNumber returns the ExitNumber field if non-nil, zero value otherwise.

### GetExitNumberOk

`func (o *ManeuverEvent) GetExitNumberOk() (*string, bool)`

GetExitNumberOk returns a tuple with the ExitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitNumber

`func (o *ManeuverEvent) SetExitNumber(v string)`

SetExitNumber sets ExitNumber field to given value.

### HasExitNumber

`func (o *ManeuverEvent) HasExitNumber() bool`

HasExitNumber returns a boolean if a field has been set.

### GetExitName

`func (o *ManeuverEvent) GetExitName() string`

GetExitName returns the ExitName field if non-nil, zero value otherwise.

### GetExitNameOk

`func (o *ManeuverEvent) GetExitNameOk() (*string, bool)`

GetExitNameOk returns a tuple with the ExitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitName

`func (o *ManeuverEvent) SetExitName(v string)`

SetExitName sets ExitName field to given value.

### HasExitName

`func (o *ManeuverEvent) HasExitName() bool`

HasExitName returns a boolean if a field has been set.

### GetRoundaboutExit

`func (o *ManeuverEvent) GetRoundaboutExit() int32`

GetRoundaboutExit returns the RoundaboutExit field if non-nil, zero value otherwise.

### GetRoundaboutExitOk

`func (o *ManeuverEvent) GetRoundaboutExitOk() (*int32, bool)`

GetRoundaboutExitOk returns a tuple with the RoundaboutExit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundaboutExit

`func (o *ManeuverEvent) SetRoundaboutExit(v int32)`

SetRoundaboutExit sets RoundaboutExit field to given value.

### HasRoundaboutExit

`func (o *ManeuverEvent) HasRoundaboutExit() bool`

HasRoundaboutExit returns a boolean if a field has been set.

### GetCombinedTransportName

`func (o *ManeuverEvent) GetCombinedTransportName() string`

GetCombinedTransportName returns the CombinedTransportName field if non-nil, zero value otherwise.

### GetCombinedTransportNameOk

`func (o *ManeuverEvent) GetCombinedTransportNameOk() (*string, bool)`

GetCombinedTransportNameOk returns a tuple with the CombinedTransportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedTransportName

`func (o *ManeuverEvent) SetCombinedTransportName(v string)`

SetCombinedTransportName sets CombinedTransportName field to given value.

### HasCombinedTransportName

`func (o *ManeuverEvent) HasCombinedTransportName() bool`

HasCombinedTransportName returns a boolean if a field has been set.

### GetCombinedTransportType

`func (o *ManeuverEvent) GetCombinedTransportType() CombinedTransportType`

GetCombinedTransportType returns the CombinedTransportType field if non-nil, zero value otherwise.

### GetCombinedTransportTypeOk

`func (o *ManeuverEvent) GetCombinedTransportTypeOk() (*CombinedTransportType, bool)`

GetCombinedTransportTypeOk returns a tuple with the CombinedTransportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedTransportType

`func (o *ManeuverEvent) SetCombinedTransportType(v CombinedTransportType)`

SetCombinedTransportType sets CombinedTransportType field to given value.

### HasCombinedTransportType

`func (o *ManeuverEvent) HasCombinedTransportType() bool`

HasCombinedTransportType returns a boolean if a field has been set.

### GetCrossingRoadName

`func (o *ManeuverEvent) GetCrossingRoadName() string`

GetCrossingRoadName returns the CrossingRoadName field if non-nil, zero value otherwise.

### GetCrossingRoadNameOk

`func (o *ManeuverEvent) GetCrossingRoadNameOk() (*string, bool)`

GetCrossingRoadNameOk returns a tuple with the CrossingRoadName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossingRoadName

`func (o *ManeuverEvent) SetCrossingRoadName(v string)`

SetCrossingRoadName sets CrossingRoadName field to given value.

### HasCrossingRoadName

`func (o *ManeuverEvent) HasCrossingRoadName() bool`

HasCrossingRoadName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


