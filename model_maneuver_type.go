/*
Routing

With the Routing service you can calculate routes from A to B taking into account vehicle-specific restrictions, traffic situations, toll, emissions, drivers' working hours, service times and opening intervals.

API version: 1.33
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package optiflow_routing

import (
	"encoding/json"
	"fmt"
)

// ManeuverType Describes the type of maneuver to execute.     * `START` - Departure at an on-road waypoint.     * `START_LEFT` - Departure at an off-road waypoint to begin the route heading left.     * `START_RIGHT` - Departure at an off-road waypoint to begin the route heading right.     * `ARRIVE` - Arrival at an on-road waypoint.     * `ARRIVE_LEFT` - Arrival at an off-road waypoint if the waypoint is on the left.     * `ARRIVE_RIGHT` - Arrival at an off-road waypoint if the waypoint is on the right.     * `CONTINUE` - Follow the current road, usually when the road type changes although the road goes straight.     * `KEEP_STRAIGHT` - Stay on the straight lane, usually in fork-shaped intersections with more than two spikes.     * `KEEP_LEFT` - Keep left or to take the left lane, usually in Y-shaped intersections.     * `KEEP_RIGHT` - Keep right or to take the right lane, usually  in Y-shaped intersections.     * `TURN_HALF_LEFT` - Turn half left at a crossing where at least one additional trailing road exists,  usually at T- or X-shaped crossings on urban or country roads. Even if the name of the turning road does not change a maneuver will be created.     * `TURN_LEFT` - Turn left at a crossing where at least one additional trailing road exists, usually  at T- or X-shaped crossings on urban or country roads. Even if the name of the turning road does not change a maneuver will be created.     * `TURN_SHARP_LEFT` - Turn sharp left at a crossing where at least one additional trailing road exists, usually   at T- or X-shaped crossings on urban or country roads. Even if the name of the turning road does not change a maneuver will be created.     * `TURN_HALF_RIGHT` - Turn half right at a crossing where at least one additional trailing road exists, usually   at T- or X-shaped crossings on urban or country roads. Even if the name of the turning road does not change a maneuver will be created.     * `TURN_RIGHT` - Turn right at a crossing where at least one additional trailing road exists, usually  at T- or X-shaped crossings on urban or country roads. Even if the name of the turning road does not change a maneuver will be created.     * `TURN_SHARP_RIGHT` - Turn sharp right at a crossing where at least one additional trailing road exists, usually   at T- or X-shaped crossings on urban or country roads. Even if the name of the turning road does not change a maneuver will be created.     * `MAKE_U_TURN` - Make a U-turn, either at the same road or at structurally separated roads. The **relativeDirection**   indicates whether to turn left or right.     * `TAKE_ROUNDABOUT_LEFT` - Enter a roundabout to the left and leave it at the given **roundaboutExit**.   Only drivable roads of the roundabout are counted.     * `TAKE_ROUNDABOUT_RIGHT` - Enter a roundabout to the right and leave it at the given **roundaboutExit**.   Only drivable roads of the roundabout are counted.     * `TAKE_COMBINED_TRANSPORT` - Take boat (ferry) or rail (train or rail shuttle) connection.  * `ENTER` - Enter a highway/freeway/major road straight.     * `ENTER_LEFT` - Enter a highway/freeway/major road to the left.     * `ENTER_RIGHT` - Enter a highway/freeway/major road to the right.     * `EXIT` - Leave a highway/freeway/major road straight.     * `EXIT_LEFT` - Leave a highway/freeway/major road to the left.     * `EXIT_RIGHT` - Leave a highway/freeway/major road to the right.     * `CHANGE` - Change straight to another highway/freeway at an interchange.     * `CHANGE_LEFT` - Change to the left to another highway/freeway at an interchange.     * `CHANGE_RIGHT` - Change to the right to another highway/freeway at an interchange.
type ManeuverType string

// List of ManeuverType
const (
	START ManeuverType = "START"
	START_LEFT ManeuverType = "START_LEFT"
	START_RIGHT ManeuverType = "START_RIGHT"
	ARRIVE ManeuverType = "ARRIVE"
	ARRIVE_LEFT ManeuverType = "ARRIVE_LEFT"
	ARRIVE_RIGHT ManeuverType = "ARRIVE_RIGHT"
	CONTINUE ManeuverType = "CONTINUE"
	KEEP_STRAIGHT ManeuverType = "KEEP_STRAIGHT"
	KEEP_LEFT ManeuverType = "KEEP_LEFT"
	KEEP_RIGHT ManeuverType = "KEEP_RIGHT"
	TURN_HALF_LEFT ManeuverType = "TURN_HALF_LEFT"
	TURN_LEFT ManeuverType = "TURN_LEFT"
	TURN_SHARP_LEFT ManeuverType = "TURN_SHARP_LEFT"
	TURN_HALF_RIGHT ManeuverType = "TURN_HALF_RIGHT"
	TURN_RIGHT ManeuverType = "TURN_RIGHT"
	TURN_SHARP_RIGHT ManeuverType = "TURN_SHARP_RIGHT"
	MAKE_U_TURN ManeuverType = "MAKE_U_TURN"
	TAKE_ROUNDABOUT_LEFT ManeuverType = "TAKE_ROUNDABOUT_LEFT"
	TAKE_ROUNDABOUT_RIGHT ManeuverType = "TAKE_ROUNDABOUT_RIGHT"
	TAKE_COMBINED_TRANSPORT ManeuverType = "TAKE_COMBINED_TRANSPORT"
	ENTER ManeuverType = "ENTER"
	ENTER_LEFT ManeuverType = "ENTER_LEFT"
	ENTER_RIGHT ManeuverType = "ENTER_RIGHT"
	EXIT ManeuverType = "EXIT"
	EXIT_LEFT ManeuverType = "EXIT_LEFT"
	EXIT_RIGHT ManeuverType = "EXIT_RIGHT"
	CHANGE ManeuverType = "CHANGE"
	CHANGE_LEFT ManeuverType = "CHANGE_LEFT"
	CHANGE_RIGHT ManeuverType = "CHANGE_RIGHT"
)

// All allowed values of ManeuverType enum
var AllowedManeuverTypeEnumValues = []ManeuverType{
	"START",
	"START_LEFT",
	"START_RIGHT",
	"ARRIVE",
	"ARRIVE_LEFT",
	"ARRIVE_RIGHT",
	"CONTINUE",
	"KEEP_STRAIGHT",
	"KEEP_LEFT",
	"KEEP_RIGHT",
	"TURN_HALF_LEFT",
	"TURN_LEFT",
	"TURN_SHARP_LEFT",
	"TURN_HALF_RIGHT",
	"TURN_RIGHT",
	"TURN_SHARP_RIGHT",
	"MAKE_U_TURN",
	"TAKE_ROUNDABOUT_LEFT",
	"TAKE_ROUNDABOUT_RIGHT",
	"TAKE_COMBINED_TRANSPORT",
	"ENTER",
	"ENTER_LEFT",
	"ENTER_RIGHT",
	"EXIT",
	"EXIT_LEFT",
	"EXIT_RIGHT",
	"CHANGE",
	"CHANGE_LEFT",
	"CHANGE_RIGHT",
}

func (v *ManeuverType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ManeuverType(value)
	for _, existing := range AllowedManeuverTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ManeuverType", value)
}

// NewManeuverTypeFromValue returns a pointer to a valid ManeuverType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewManeuverTypeFromValue(v string) (*ManeuverType, error) {
	ev := ManeuverType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ManeuverType: valid values are %v", v, AllowedManeuverTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ManeuverType) IsValid() bool {
	for _, existing := range AllowedManeuverTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ManeuverType value
func (v ManeuverType) Ptr() *ManeuverType {
	return &v
}

type NullableManeuverType struct {
	value *ManeuverType
	isSet bool
}

func (v NullableManeuverType) Get() *ManeuverType {
	return v.value
}

func (v *NullableManeuverType) Set(val *ManeuverType) {
	v.value = val
	v.isSet = true
}

func (v NullableManeuverType) IsSet() bool {
	return v.isSet
}

func (v *NullableManeuverType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManeuverType(val *ManeuverType) *NullableManeuverType {
	return &NullableManeuverType{value: val, isSet: true}
}

func (v NullableManeuverType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManeuverType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

