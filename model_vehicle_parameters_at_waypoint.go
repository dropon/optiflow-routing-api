/*
Routing

With the Routing service you can calculate routes from A to B taking into account vehicle-specific restrictions, traffic situations, toll, emissions, drivers' working hours, service times and opening intervals.

API version: 1.33
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package optiflow_routing

import (
	"encoding/json"
)

// checks if the VehicleParametersAtWaypoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VehicleParametersAtWaypoint{}

// VehicleParametersAtWaypoint The vehicle parameters that change at a waypoint. A vehicle parameter specified at a waypoint overrides the setting from the **profile** and **vehicle**. It is valid until the end of the route unless it is changed again at a subsequent waypoint. 
type VehicleParametersAtWaypoint struct {
	// The average fuel consumption of the vehicle. Depending on the **fuelType** [l/100km] for liquid fuel types or [kg/100km] for gaseous fuel types.  Supported for **engineType** _COMBUSTION_  or _HYBRID_. Relevant for `emissions`. 
	AverageFuelConsumption *float64 `json:"averageFuelConsumption,omitempty"`
	// The average electricity consumption of the vehicle if the **engineType** is _ELECTRIC_ or _HYBRID_ [kWh/100km].  This field is not used, if a model of an electric vehicle is used. In this case the average electricity  consumption from the model based consumption calculation is used for the emissions calculation.  Supported for **engineType** _ELECTRIC_ or _HYBRID_. Relevant for `emissions`. 
	AverageElectricityConsumption *float64 `json:"averageElectricityConsumption,omitempty"`
	// The weight of the vehicle's load [kg].  Relevant for `routing`, `emissions`, `range calculation`. 
	LoadWeight *int32 `json:"loadWeight,omitempty"`
	// The maximum distributed weight that may be supported by an axle of the vehicle [kg].  Relevant for `routing`, `toll`. 
	AxleWeight *int32 `json:"axleWeight,omitempty"`
	// The height of the vehicle [cm].  Relevant for `routing`, `toll`. 
	Height *int32 `json:"height,omitempty"`
	// The length of the vehicle [cm].  Relevant for `routing`, `toll`. 
	Length *int32 `json:"length,omitempty"`
	// The width of the vehicle [cm].  Relevant for `routing`, `toll`. 
	Width *int32 `json:"width,omitempty"`
	// The list of hazardous materials the vehicle has loaded. Use a list with only the _NONE_ value to specify that no hazardous material is loaded from that waypoint on. An empty list means that the hazardous materials don't change at the waypoint. If _NONE_ is specified along with other hazardous materials it is ignored.  Relevant for `routing`. 
	HazardousMaterials []HazardousMaterials `json:"hazardousMaterials,omitempty"`
	TunnelRestrictionCode *TunnelRestrictionCode `json:"tunnelRestrictionCode,omitempty"`
	// The list of truck routes the vehicle has to follow. This parameter will be ignored for non-truck profiles such as EUR_CAR, EUR_VAN, USA_1_PICKUP or AUS_LCV_LIGHT_COMMERCIAL.  * `DE_LKWUEBERLSTVAUSNV`  Preferred routes for long trucks in Germany, also known as Lang-LKW.  * `NL_LZV`  Preferred routes for long trucks in the Netherlands, also known as LZV (Langere en Zwaardere Vrachtautocombinatie).  * `NZ_HPMV`  The network for High Productivity Motor Vehicles (HPMV) carrying the maximum loads available under a permit (New Zealand Transport Agency).  * `SE_BK_1`  Public roads and bridges that support up to 64 t total permitted weight (Swedish Transport Administration).  * `SE_BK_2`  Public roads and bridges that support up to 51.4 t total permitted weight.  Actual limit depends on wheelbase and axle weight (Swedish Transport Administration).  * `SE_BK_3`  Public roads and bridges that support up to 37.5 t total permitted weight.  Actual limit depends on wheelbase and axle weight (Swedish Transport Administration).  * `SE_BK_4`  Public roads and bridges that support up to 74 t total permitted weight (draft summer 2018, Swedish Transport Administration).  * `US_STAA`  Routes that belong to the highway network as defined by the Surface Transportation Assistance Act in the US.  * `US_TD`  Part of a state-designated highway network for trucks in the US.  * `AU_B_DOUBLE`  B-Double routes as defined in Australia.  * `AU_B_DOUBLE_HML`  Routes for B-Double vehicle combinations operating at Higher Mass Limits (HML) (Australian Transport Administration).  * `AU_B_TRIPLE`  B-Triple routes as defined in Australia.  * `AU_B_TRIPLE_HML`  Routes for B-Triple vehicle combinations operating at Higher Mass Limits (HML) (Australian Transport Administration).  * `AU_AB_TRIPLE`  Routes for AB-Triple vehicle combinations operating (Australian Transport Administration).  * `AU_AB_TRIPLE_HML`  Routes for AB-Triple vehicle combinations operating at Higher Mass Limits (HML) (Australian Transport Administration).  * `NONE`  Use a list with only the _NONE_ value for the truck routes  at a waypoint to specify that no truck route must be used from that waypoint on. An empty list of truck routes  at a waypoint means that the truck routes don't change at the waypoint. If _NONE_ is specified along with other  truck routes it is ignored.  Relevant for `routing`. Cannot be used with **options[routingMode]=MONETARY**. 
	TruckRoutes []TruckRoutes `json:"truckRoutes,omitempty"`
}

// NewVehicleParametersAtWaypoint instantiates a new VehicleParametersAtWaypoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVehicleParametersAtWaypoint() *VehicleParametersAtWaypoint {
	this := VehicleParametersAtWaypoint{}
	return &this
}

// NewVehicleParametersAtWaypointWithDefaults instantiates a new VehicleParametersAtWaypoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVehicleParametersAtWaypointWithDefaults() *VehicleParametersAtWaypoint {
	this := VehicleParametersAtWaypoint{}
	return &this
}

// GetAverageFuelConsumption returns the AverageFuelConsumption field value if set, zero value otherwise.
func (o *VehicleParametersAtWaypoint) GetAverageFuelConsumption() float64 {
	if o == nil || IsNil(o.AverageFuelConsumption) {
		var ret float64
		return ret
	}
	return *o.AverageFuelConsumption
}

// GetAverageFuelConsumptionOk returns a tuple with the AverageFuelConsumption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleParametersAtWaypoint) GetAverageFuelConsumptionOk() (*float64, bool) {
	if o == nil || IsNil(o.AverageFuelConsumption) {
		return nil, false
	}
	return o.AverageFuelConsumption, true
}

// HasAverageFuelConsumption returns a boolean if a field has been set.
func (o *VehicleParametersAtWaypoint) HasAverageFuelConsumption() bool {
	if o != nil && !IsNil(o.AverageFuelConsumption) {
		return true
	}

	return false
}

// SetAverageFuelConsumption gets a reference to the given float64 and assigns it to the AverageFuelConsumption field.
func (o *VehicleParametersAtWaypoint) SetAverageFuelConsumption(v float64) {
	o.AverageFuelConsumption = &v
}

// GetAverageElectricityConsumption returns the AverageElectricityConsumption field value if set, zero value otherwise.
func (o *VehicleParametersAtWaypoint) GetAverageElectricityConsumption() float64 {
	if o == nil || IsNil(o.AverageElectricityConsumption) {
		var ret float64
		return ret
	}
	return *o.AverageElectricityConsumption
}

// GetAverageElectricityConsumptionOk returns a tuple with the AverageElectricityConsumption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleParametersAtWaypoint) GetAverageElectricityConsumptionOk() (*float64, bool) {
	if o == nil || IsNil(o.AverageElectricityConsumption) {
		return nil, false
	}
	return o.AverageElectricityConsumption, true
}

// HasAverageElectricityConsumption returns a boolean if a field has been set.
func (o *VehicleParametersAtWaypoint) HasAverageElectricityConsumption() bool {
	if o != nil && !IsNil(o.AverageElectricityConsumption) {
		return true
	}

	return false
}

// SetAverageElectricityConsumption gets a reference to the given float64 and assigns it to the AverageElectricityConsumption field.
func (o *VehicleParametersAtWaypoint) SetAverageElectricityConsumption(v float64) {
	o.AverageElectricityConsumption = &v
}

// GetLoadWeight returns the LoadWeight field value if set, zero value otherwise.
func (o *VehicleParametersAtWaypoint) GetLoadWeight() int32 {
	if o == nil || IsNil(o.LoadWeight) {
		var ret int32
		return ret
	}
	return *o.LoadWeight
}

// GetLoadWeightOk returns a tuple with the LoadWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleParametersAtWaypoint) GetLoadWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.LoadWeight) {
		return nil, false
	}
	return o.LoadWeight, true
}

// HasLoadWeight returns a boolean if a field has been set.
func (o *VehicleParametersAtWaypoint) HasLoadWeight() bool {
	if o != nil && !IsNil(o.LoadWeight) {
		return true
	}

	return false
}

// SetLoadWeight gets a reference to the given int32 and assigns it to the LoadWeight field.
func (o *VehicleParametersAtWaypoint) SetLoadWeight(v int32) {
	o.LoadWeight = &v
}

// GetAxleWeight returns the AxleWeight field value if set, zero value otherwise.
func (o *VehicleParametersAtWaypoint) GetAxleWeight() int32 {
	if o == nil || IsNil(o.AxleWeight) {
		var ret int32
		return ret
	}
	return *o.AxleWeight
}

// GetAxleWeightOk returns a tuple with the AxleWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleParametersAtWaypoint) GetAxleWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.AxleWeight) {
		return nil, false
	}
	return o.AxleWeight, true
}

// HasAxleWeight returns a boolean if a field has been set.
func (o *VehicleParametersAtWaypoint) HasAxleWeight() bool {
	if o != nil && !IsNil(o.AxleWeight) {
		return true
	}

	return false
}

// SetAxleWeight gets a reference to the given int32 and assigns it to the AxleWeight field.
func (o *VehicleParametersAtWaypoint) SetAxleWeight(v int32) {
	o.AxleWeight = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *VehicleParametersAtWaypoint) GetHeight() int32 {
	if o == nil || IsNil(o.Height) {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleParametersAtWaypoint) GetHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *VehicleParametersAtWaypoint) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *VehicleParametersAtWaypoint) SetHeight(v int32) {
	o.Height = &v
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *VehicleParametersAtWaypoint) GetLength() int32 {
	if o == nil || IsNil(o.Length) {
		var ret int32
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleParametersAtWaypoint) GetLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.Length) {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *VehicleParametersAtWaypoint) HasLength() bool {
	if o != nil && !IsNil(o.Length) {
		return true
	}

	return false
}

// SetLength gets a reference to the given int32 and assigns it to the Length field.
func (o *VehicleParametersAtWaypoint) SetLength(v int32) {
	o.Length = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *VehicleParametersAtWaypoint) GetWidth() int32 {
	if o == nil || IsNil(o.Width) {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleParametersAtWaypoint) GetWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *VehicleParametersAtWaypoint) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *VehicleParametersAtWaypoint) SetWidth(v int32) {
	o.Width = &v
}

// GetHazardousMaterials returns the HazardousMaterials field value if set, zero value otherwise.
func (o *VehicleParametersAtWaypoint) GetHazardousMaterials() []HazardousMaterials {
	if o == nil || IsNil(o.HazardousMaterials) {
		var ret []HazardousMaterials
		return ret
	}
	return o.HazardousMaterials
}

// GetHazardousMaterialsOk returns a tuple with the HazardousMaterials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleParametersAtWaypoint) GetHazardousMaterialsOk() ([]HazardousMaterials, bool) {
	if o == nil || IsNil(o.HazardousMaterials) {
		return nil, false
	}
	return o.HazardousMaterials, true
}

// HasHazardousMaterials returns a boolean if a field has been set.
func (o *VehicleParametersAtWaypoint) HasHazardousMaterials() bool {
	if o != nil && !IsNil(o.HazardousMaterials) {
		return true
	}

	return false
}

// SetHazardousMaterials gets a reference to the given []HazardousMaterials and assigns it to the HazardousMaterials field.
func (o *VehicleParametersAtWaypoint) SetHazardousMaterials(v []HazardousMaterials) {
	o.HazardousMaterials = v
}

// GetTunnelRestrictionCode returns the TunnelRestrictionCode field value if set, zero value otherwise.
func (o *VehicleParametersAtWaypoint) GetTunnelRestrictionCode() TunnelRestrictionCode {
	if o == nil || IsNil(o.TunnelRestrictionCode) {
		var ret TunnelRestrictionCode
		return ret
	}
	return *o.TunnelRestrictionCode
}

// GetTunnelRestrictionCodeOk returns a tuple with the TunnelRestrictionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleParametersAtWaypoint) GetTunnelRestrictionCodeOk() (*TunnelRestrictionCode, bool) {
	if o == nil || IsNil(o.TunnelRestrictionCode) {
		return nil, false
	}
	return o.TunnelRestrictionCode, true
}

// HasTunnelRestrictionCode returns a boolean if a field has been set.
func (o *VehicleParametersAtWaypoint) HasTunnelRestrictionCode() bool {
	if o != nil && !IsNil(o.TunnelRestrictionCode) {
		return true
	}

	return false
}

// SetTunnelRestrictionCode gets a reference to the given TunnelRestrictionCode and assigns it to the TunnelRestrictionCode field.
func (o *VehicleParametersAtWaypoint) SetTunnelRestrictionCode(v TunnelRestrictionCode) {
	o.TunnelRestrictionCode = &v
}

// GetTruckRoutes returns the TruckRoutes field value if set, zero value otherwise.
func (o *VehicleParametersAtWaypoint) GetTruckRoutes() []TruckRoutes {
	if o == nil || IsNil(o.TruckRoutes) {
		var ret []TruckRoutes
		return ret
	}
	return o.TruckRoutes
}

// GetTruckRoutesOk returns a tuple with the TruckRoutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleParametersAtWaypoint) GetTruckRoutesOk() ([]TruckRoutes, bool) {
	if o == nil || IsNil(o.TruckRoutes) {
		return nil, false
	}
	return o.TruckRoutes, true
}

// HasTruckRoutes returns a boolean if a field has been set.
func (o *VehicleParametersAtWaypoint) HasTruckRoutes() bool {
	if o != nil && !IsNil(o.TruckRoutes) {
		return true
	}

	return false
}

// SetTruckRoutes gets a reference to the given []TruckRoutes and assigns it to the TruckRoutes field.
func (o *VehicleParametersAtWaypoint) SetTruckRoutes(v []TruckRoutes) {
	o.TruckRoutes = v
}

func (o VehicleParametersAtWaypoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VehicleParametersAtWaypoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AverageFuelConsumption) {
		toSerialize["averageFuelConsumption"] = o.AverageFuelConsumption
	}
	if !IsNil(o.AverageElectricityConsumption) {
		toSerialize["averageElectricityConsumption"] = o.AverageElectricityConsumption
	}
	if !IsNil(o.LoadWeight) {
		toSerialize["loadWeight"] = o.LoadWeight
	}
	if !IsNil(o.AxleWeight) {
		toSerialize["axleWeight"] = o.AxleWeight
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.Length) {
		toSerialize["length"] = o.Length
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.HazardousMaterials) {
		toSerialize["hazardousMaterials"] = o.HazardousMaterials
	}
	if !IsNil(o.TunnelRestrictionCode) {
		toSerialize["tunnelRestrictionCode"] = o.TunnelRestrictionCode
	}
	if !IsNil(o.TruckRoutes) {
		toSerialize["truckRoutes"] = o.TruckRoutes
	}
	return toSerialize, nil
}

type NullableVehicleParametersAtWaypoint struct {
	value *VehicleParametersAtWaypoint
	isSet bool
}

func (v NullableVehicleParametersAtWaypoint) Get() *VehicleParametersAtWaypoint {
	return v.value
}

func (v *NullableVehicleParametersAtWaypoint) Set(val *VehicleParametersAtWaypoint) {
	v.value = val
	v.isSet = true
}

func (v NullableVehicleParametersAtWaypoint) IsSet() bool {
	return v.isSet
}

func (v *NullableVehicleParametersAtWaypoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVehicleParametersAtWaypoint(val *VehicleParametersAtWaypoint) *NullableVehicleParametersAtWaypoint {
	return &NullableVehicleParametersAtWaypoint{value: val, isSet: true}
}

func (v NullableVehicleParametersAtWaypoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVehicleParametersAtWaypoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


