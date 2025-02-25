/*
Routing

With the Routing service you can calculate routes from A to B taking into account vehicle-specific restrictions, traffic situations, toll, emissions, drivers' working hours, service times and opening intervals.

API version: 1.33
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package optiflow_routing

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CausingError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CausingError{}

// CausingError struct for CausingError
type CausingError struct {
	// A human readable message that describes the error.
	Description string `json:"description"`
	// A constant string that can be used to identify this error class programmatically. An errorCode can have **details** to provide information in additional properties which are described with the code they apply to. They are of type string unless otherwise specified. Note that additional errorCodes as well as the **details** of existing errorCodes may be added at any time. Furthermore, the **description** may change at any time.    **Error codes for** `GENERAL_VALIDATION_ERROR`  * `GENERAL_INVALID_VALUE` - A parameter is set to an invalid value.   * `value` - The invalid value. * `GENERAL_UNRECOGNIZED_PARAMETER` - A parameter is unknown. * `GENERAL_MISSING_PARAMETER` - A required parameter is missing. * `GENERAL_MINIMUM_LENGTH_VIOLATED` - The minimum length is violated.   * `minimumLength` - The minimum length (integer). * `GENERAL_MAXIMUM_LENGTH_VIOLATED` - The maximum length is violated.   * `maximumLength` - The maximum length (integer). * `GENERAL_MINIMUM_VALUE_VIOLATED` - The minimum value restriction is violated.   * `minimumValue` - The minimum value (integer or double). * `GENERAL_MAXIMUM_VALUE_VIOLATED` - The maximum value restriction is violated.   * `maximumValue` - The maximum value (integer or double). * `GENERAL_DUPLICATE_PARAMETER` - A parameter is duplicated. * `GENERAL_INVALID_LIST` - A list has an invalid format such as duplicate commas.   * `value` - The invalid list. * `GENERAL_INVALID_INTERVAL` - A time interval is invalid, i.e. start is greater than end. * `ROUTING_INVALID_WAYPOINT_ATTRIBUTE` - A waypoint attribute is set to an invalid value.   * `attribute` - The invalid waypoint attribute. * `ROUTING_UNRECOGNIZED_WAYPOINT_ATTRIBUTE` - A waypoint attribute is unknown.   * `attribute` - The invalid waypoint key. * `ROUTING_DUPLICATE_WAYPOINT_ATTRIBUTE` - A waypoint attribute is duplicated.   * `attribute` - The duplicated waypoint key. * `ROUTING_WAYPOINT_ATTRIBUTE_CONFLICT` - Two waypoint attributes are in conflict with each other.   * `attribute` - The first conflicting attribute.   * `conflictingAttribute` - The second conflicting attribute. * `ROUTING_INVALID_MANIPULATION_WAYPOINT_ORDER` - The manipulation waypoint is not valid for start or destination. * `ROUTING_INVALID_COMBINED_TRANSPORT_WAYPOINT_ORDER` - The combinedTransport waypoint is not valid for start or destination. * `ROUTING_INVALID_WAYPOINT_LIST_FOR_ALTERNATIVE_ROUTES` - Alternative routes are supported only for two on-road or off-road waypoints. * `ROUTING_INVALID_WAYPOINT` - A waypoint contains multiple types or none of them, but exactly one must be specified. * `ROUTING_MUST_HAVE_WAYPOINTS_OR_ROUTE_ID` - The request must have either at least two **waypoints** or a **routeId**. * `ROUTING_EMISSIONS_MUTUALLY_EXCLUSIVE` - All emissions _EN16258_2012_ results and _ISO14083_2022_ or _ISO14083_2023_ results are mutually exclusive.   * `attribute` - The first conflicting emissions standard.   * `conflictingAttributes` - The list of other conflicting emissions standards. * `ROUTING_START_AND_ARRIVAL_TIME_MUTUALLY_EXCLUSIVE` - **options[startTime]** and **options[arrivalTime]** are mutually exclusive. - _The **parameter** remains empty._ * `ROUTING_ESTIMATED_DISTANCE_TOO_LONG` - The distance of the route (estimated by air-line) for the selected vehicle is too long. - _The **parameter** remains empty._   * `distance` - The estimated distance (integer).   * `limit` - The maximum allowable distance (integer). * `ROUTING_PARAMETER_CONFLICT` - Two parameters are in conflict with each other.   * `conflictingParameter` - The conflicting parameter.   * `message` - The error message. * `ROUTING_NO_VALID_COUNTRY_ALLOWED` - The list of allowed countries does not contain any of the available countries so that the effective list of countries allowed for routing is empty.   * `allowedCountries` - The list of allowed countries. * `ROUTING_ALL_VALID_COUNTRIES_PROHIBITED` - The list of prohibited countries contains all available countries so that the effective list of countries allowed for routing is empty.   * `prohibitedCountries` - The list of prohibited countries. * `ROUTING_MAXIMUM_HORIZON_VALUE_VIOLATED` - The maximum value of horizon is violated.   * `limit` - The maximum allowable horizon (integer). * `ROUTING_MUST_HAVE_ONE_WAYPOINT_OR_ROUTE_ID` - The request must have either a **waypoint** or a **routeId**. * `ROUTING_HORIZONS_EQUAL_OR_NOT_ASCENDING` - The horizons have equal values or are not ascending.   * `value` - The invalid horizon. * `ROUTING_ROUTE_TOO_LONG_FOR_REACHABILITY` - The route is too long to be used with reachable areas or locations.   * `length` - The actual route length (integer).   * `limit` - The maximum allowable route length (integer). * `ROUTING_ALLOWED_AND_PROHIBITED_COUNTRIES_IN_CONFLICT_WITH_ROUTE_ID` - The lists of allowed and prohibited countries are in conflict with the **routeId** which passes an effectively prohibited country.   * `value` - The value in conflict. * `ROUTING_ROUTE_ID_NOT_FOUND` - The **routeId** cannot be found.   * `value` - The routeId. * `ROUTING_ROUTE_ID_CANNOT_BE_USED` - The **routeId** cannot be used for this operation as it was created by a service other than routing and lacks a routing context.   * `value` - The routeId. * `ROUTING_PROFILE_NOT_FOUND` - The requested **profile** could not be found.   * `value` - The profile name. * `ROUTING_UNSUPPORTED_CURRENCY` - The specified currency is not supported.   * `currency` - The unsupported currency. * `ROUTING_PARAMETER_ONLY_SUPPORTED_BY_POST` - The requested parameter is not supported by this GET operation, it is only supported by the corresponding POST operation.   * `value` - The invalid parameter value. * `ROUTING_OPENING_INTERVALS_REQUIRE_TIME` - When using opening intervals with a waypoint and with **options[trafficMode]** _AVERAGE_ either a start or an arrival time has to be specified. * `ROUTING_ARRIVAL_TIME_WITH_SCHEDULE` - **options[arrivalTime]** cannot be used with the **results** _SCHEDULE_REPORT_ and _SCHEDULE_EVENT_ nor when **openingIntervals**, **serviceTime** or **workingHoursPreset** are specified.   * `value` - The invalid parameter value. * `ROUTING_INVALID_NUMBER_OF_COORDINATES` - The polyline cannot be parsed because the number of coordinates is not even or less than 4.   * `value` - The invalid parameter value.   * `polylineIndex` - The index denoting the polyline in which the error was found (integer). * `ROUTING_INVALID_COORDINATE` - The provided coordinate is not in the valid range or cannot be parsed.   * `value` - The invalid parameter value.   * `polylineIndex` - The index denoting the polyline in which the error was found (integer).   * `coordinateIndex` - The index denoting the erroneous coordinate within the polyline (integer). * `ROUTING_FEATURE_NOT_SUPPORTED_WITH_MONETARY_COSTS` - The requested feature is not supported when **options[routingMode]** is _MONETARY_.   * `value` - The invalid parameter value. * `ROUTING_MUST_HAVE_MONETARY_COST_VALUE` - Both values **monetaryCostOptions[costPerKilometer]** and **monetaryCostOptions[workingCostPerHour]** are zero. Use a value greater zero for at least one of this **monetaryCostOptions** parameters. * `ROUTING_CUSTOM_ROAD_ATTRIBUTE_SCENARIO_NOT_FOUND` - At least one of the requested **options[customRoadAttributeScenarios]** could not be found.   * `scenarios` - The scenarios which could not be found (comma-separated list). * `ROUTING_CUSTOM_ROAD_ATTRIBUTE_SCENARIOS_TOO_LARGE` - The scenarios given in **options[customRoadAttributeScenarios]** are too large and can not all be considered at the same time. * `ROUTING_POSITION_AND_WAYPOINT_MUTUALLY_EXCLUSIVE` - **position** and **waypoint** are mutually exclusive. - _The **parameter** remains empty._ * `ROUTING_VEHICLE_POSITION_MISSING` - The position of the vehicle must be specified by either **position** or **waypoint**. - _The **parameter** remains empty._ * `ROUTING_UNSUPPORTED_WAYPOINT_TYPE_FOR_ETA_CALCULATION` - The ETA calculation does not support route-manipulation waypoints, combined-transport waypoints or vehicle parameters at waypoints.   * `waypointIndex` - The index of the waypoint (integer). * `ROUTING_MISSING_WAYPOINT_NAME` - The requested route for the **routeId** contains a waypoint which does not have a name.   * `waypointIndex` - The index of the waypoint (integer). * `ROUTING_DUPLICATE_WAYPOINT_NAME` - The requested route for the **routeId** contains waypoints with a duplicate name.   * `waypointIndexes` - The indexes of the waypoints with the duplicated name (comma-separated list).   * `name` - The duplicate waypoint name. * `ROUTING_WAYPOINT_NAME_NOT_FOUND` - The waypoint name could not be found in the requested route for the **routeId**.   * `name` - The invalid waypoint name. * `ROUTING_VEHICLE_POSITION_BEFORE_FIRST_WAYPOINT` - The position of the vehicle cannot be before the first waypoint. * `ROUTING_ROUTE_ID_REQUIRES_WORKLOGBOOK` - The route associated with the given **routeId** has been calculated with a **driver** and requires the **workLogbook** to be specified in the ETA request. * `ROUTING_ROUTE_ID_DOES_NOT_SUPPORT_WORKLOGBOOK` - The route associated with the given **routeId** has been calculated without a **driver** and does not support the **workLogbook** specified in the ETA request. * `ROUTING_CONFLICTING_LOW_EMISSION_ZONE_TYPES` - Some requested low-emission zone types are in conflict with each other, specify only one of them. * `ROUTING_SERVICE_CANNOT_BE_SCHEDULED` - The **serviceTime** at a waypoint cannot be scheduled as it exceeds the maximum time between breaks permitted by the selected **workingHoursPreset**. - _The **parameter** remains empty._  **Error codes for** `ROUTING_ERROR`  * `ROUTING_WAYPOINT_CANNOT_BE_MATCHED` - The waypoint cannot be matched to the nearest possible road. * `ROUTING_ROUTE_NOT_FOUND` - A route between at least two waypoints could not be found for the current configuration and profile. The **parameter** contains the waypoint where the problematic part that could not be routed starts, i.e., the problematic part of the route is between this waypoint and the next one. Note that only the first problematic part that was encountered is reported. * `ROUTING_TIMEOUT` - The route calculation has timed out. * `ROUTING_UTC_OFFSET_CANNOT_BE_DETERMINED` - The UTC offset of the start waypoint cannot be determined. * `ROUTING_BLOCK_INTERSECTING_ROADS_TOO_MANY_SEGMENTS` - The maximum number of road segments intersecting one polyline must not exceed 5000.  **Error codes for** `ROUTING_RESTRICTION_EXCEEDED`  * `ROUTING_TOO_MANY_WAYPOINTS` - The request contains too many waypoints.   * `limit`- The maximum allowed number of waypoints for a single request (integer).  **Error codes for** `GENERAL_RESOURCE_NOT_FOUND`  * `GENERAL_INVALID_ID` - The ID does not exist.   * `value` - The invalid ID.
	ErrorCode string `json:"errorCode"`
	// The name of the affected query or path parameter or a JSONPath to the affected property of the request.
	Parameter *string `json:"parameter,omitempty"`
	// Additional properties specific to this error class.
	Details map[string]interface{} `json:"details,omitempty"`
}

type _CausingError CausingError

// NewCausingError instantiates a new CausingError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCausingError(description string, errorCode string) *CausingError {
	this := CausingError{}
	this.Description = description
	this.ErrorCode = errorCode
	return &this
}

// NewCausingErrorWithDefaults instantiates a new CausingError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCausingErrorWithDefaults() *CausingError {
	this := CausingError{}
	return &this
}

// GetDescription returns the Description field value
func (o *CausingError) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CausingError) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CausingError) SetDescription(v string) {
	o.Description = v
}

// GetErrorCode returns the ErrorCode field value
func (o *CausingError) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *CausingError) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *CausingError) SetErrorCode(v string) {
	o.ErrorCode = v
}

// GetParameter returns the Parameter field value if set, zero value otherwise.
func (o *CausingError) GetParameter() string {
	if o == nil || IsNil(o.Parameter) {
		var ret string
		return ret
	}
	return *o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CausingError) GetParameterOk() (*string, bool) {
	if o == nil || IsNil(o.Parameter) {
		return nil, false
	}
	return o.Parameter, true
}

// HasParameter returns a boolean if a field has been set.
func (o *CausingError) HasParameter() bool {
	if o != nil && !IsNil(o.Parameter) {
		return true
	}

	return false
}

// SetParameter gets a reference to the given string and assigns it to the Parameter field.
func (o *CausingError) SetParameter(v string) {
	o.Parameter = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *CausingError) GetDetails() map[string]interface{} {
	if o == nil || IsNil(o.Details) {
		var ret map[string]interface{}
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CausingError) GetDetailsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Details) {
		return map[string]interface{}{}, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *CausingError) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given map[string]interface{} and assigns it to the Details field.
func (o *CausingError) SetDetails(v map[string]interface{}) {
	o.Details = v
}

func (o CausingError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CausingError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["errorCode"] = o.ErrorCode
	if !IsNil(o.Parameter) {
		toSerialize["parameter"] = o.Parameter
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

func (o *CausingError) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"errorCode",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCausingError := _CausingError{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCausingError)

	if err != nil {
		return err
	}

	*o = CausingError(varCausingError)

	return err
}

type NullableCausingError struct {
	value *CausingError
	isSet bool
}

func (v NullableCausingError) Get() *CausingError {
	return v.value
}

func (v *NullableCausingError) Set(val *CausingError) {
	v.value = val
	v.isSet = true
}

func (v NullableCausingError) IsSet() bool {
	return v.isSet
}

func (v *NullableCausingError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCausingError(val *CausingError) *NullableCausingError {
	return &NullableCausingError{value: val, isSet: true}
}

func (v NullableCausingError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCausingError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


