# VehicleParametersAtWaypoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageFuelConsumption** | Pointer to **float64** | The average fuel consumption of the vehicle. Depending on the **fuelType** [l/100km] for liquid fuel types or [kg/100km] for gaseous fuel types.  Supported for **engineType** _COMBUSTION_  or _HYBRID_. Relevant for &#x60;emissions&#x60;.  | [optional] 
**AverageElectricityConsumption** | Pointer to **float64** | The average electricity consumption of the vehicle if the **engineType** is _ELECTRIC_ or _HYBRID_ [kWh/100km].  This parameter is not used, if a model of an electric vehicle is used. In this case the average electricity  consumption from the model based consumption calculation is used for the emissions calculation.  Supported for **engineType** _ELECTRIC_ or _HYBRID_. Relevant for &#x60;emissions&#x60;.  | [optional] 
**LoadWeight** | Pointer to **int32** | The weight of the vehicle&#39;s load [kg].  Relevant for &#x60;routing&#x60;, &#x60;emissions&#x60;, &#x60;range calculation&#x60;.  | [optional] 
**AxleWeight** | Pointer to **int32** | The maximum distributed weight that may be supported by an axle of the vehicle [kg].  Relevant for &#x60;routing&#x60;, &#x60;toll&#x60;.  | [optional] 
**Height** | Pointer to **int32** | The height of the vehicle [cm].  Relevant for &#x60;routing&#x60;, &#x60;toll&#x60;.  | [optional] 
**Length** | Pointer to **int32** | The length of the vehicle [cm].  Relevant for &#x60;routing&#x60;, &#x60;toll&#x60;.  | [optional] 
**Width** | Pointer to **int32** | The width of the vehicle [cm].  Relevant for &#x60;routing&#x60;, &#x60;toll&#x60;.  | [optional] 
**HazardousMaterials** | Pointer to [**[]HazardousMaterials**](HazardousMaterials.md) | The list of hazardous materials the vehicle has loaded. Use a list with only the _NONE_ value to specify that no hazardous material is loaded from that waypoint on. An empty list means that the hazardous materials don&#39;t change at the waypoint. If _NONE_ is specified along with other hazardous materials it is ignored.  Relevant for &#x60;routing&#x60;.  | [optional] 
**TunnelRestrictionCode** | Pointer to [**TunnelRestrictionCode**](TunnelRestrictionCode.md) |  | [optional] 
**TruckRoutes** | Pointer to [**[]TruckRoutes**](TruckRoutes.md) | The list of truck routes the vehicle has to follow.  * &#x60;DE_LKWUEBERLSTVAUSNV&#x60;  Preferred routes for long trucks in Germany, also known as Lang-LKW.  * &#x60;NL_LZV&#x60;  Preferred routes for long trucks in the Netherlands, also known as LZV (Langere en Zwaardere Vrachtautocombinatie).  * &#x60;NZ_HPMV&#x60;  The network for High Productivity Motor Vehicles (HPMV) carrying the maximum loads available under a permit (New Zealand Transport Agency).  * &#x60;SE_BK_1&#x60;  Public roads and bridges that support up to 64 t total permitted weight (Swedish Transport Administration).  * &#x60;SE_BK_2&#x60;  Public roads and bridges that support up to 51.4 t total permitted weight.  Actual limit depends on wheelbase and axle weight (Swedish Transport Administration).  * &#x60;SE_BK_3&#x60;  Public roads and bridges that support up to 37.5 t total permitted weight.  Actual limit depends on wheelbase and axle weight (Swedish Transport Administration).  * &#x60;SE_BK_4&#x60;  Public roads and bridges that support up to 74 t total permitted weight (draft summer 2018, Swedish Transport Administration).  * &#x60;US_STAA&#x60;  Routes that belong to the highway network as defined by the Surface Transportation Assistance Act in the US.  * &#x60;US_TD&#x60;  Part of a state-designated highway network for trucks in the US.  * &#x60;AU_B_DOUBLE&#x60;  B-Double routes as defined in Australia.  * &#x60;AU_B_DOUBLE_HML&#x60;  Routes for B-Double vehicle combinations operating at Higher Mass Limits (HML) (Australian Transport Administration).  * &#x60;AU_B_TRIPLE&#x60;  B-Triple routes as defined in Australia.  * &#x60;AU_B_TRIPLE_HML&#x60;  Routes for B-Triple vehicle combinations operating at Higher Mass Limits (HML) (Australian Transport Administration).  * &#x60;AU_AB_TRIPLE&#x60;  Routes for AB-Triple vehicle combinations operating (Australian Transport Administration).  * &#x60;AU_AB_TRIPLE_HML&#x60;  Routes for AB-Triple vehicle combinations operating at Higher Mass Limits (HML) (Australian Transport Administration).  * &#x60;GENERAL_TRUCK_ROUTES&#x60;  General routes designated for trucks, for example to prevent trucks routing through city centres when they are on transit.  * &#x60;NONE&#x60;  Use a list with only the _NONE_ value for the truck routes  at a waypoint to specify that no truck route must be used from that waypoint on. An empty list of truck routes  at a waypoint means that the truck routes don&#39;t change at the waypoint. If _NONE_ is specified along with other  truck routes it is ignored.  This parameter will be ignored for non-truck profiles, e.g. _EUR_CAR_, _EUR_VAN_, _USA_1_PICKUP_ or _AUS_LCV_LIGHT_COMMERCIAL_ and for routing modes other than _FAST_. Relevant for &#x60;routing&#x60;.  | [optional] 

## Methods

### NewVehicleParametersAtWaypoint

`func NewVehicleParametersAtWaypoint() *VehicleParametersAtWaypoint`

NewVehicleParametersAtWaypoint instantiates a new VehicleParametersAtWaypoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleParametersAtWaypointWithDefaults

`func NewVehicleParametersAtWaypointWithDefaults() *VehicleParametersAtWaypoint`

NewVehicleParametersAtWaypointWithDefaults instantiates a new VehicleParametersAtWaypoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageFuelConsumption

`func (o *VehicleParametersAtWaypoint) GetAverageFuelConsumption() float64`

GetAverageFuelConsumption returns the AverageFuelConsumption field if non-nil, zero value otherwise.

### GetAverageFuelConsumptionOk

`func (o *VehicleParametersAtWaypoint) GetAverageFuelConsumptionOk() (*float64, bool)`

GetAverageFuelConsumptionOk returns a tuple with the AverageFuelConsumption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageFuelConsumption

`func (o *VehicleParametersAtWaypoint) SetAverageFuelConsumption(v float64)`

SetAverageFuelConsumption sets AverageFuelConsumption field to given value.

### HasAverageFuelConsumption

`func (o *VehicleParametersAtWaypoint) HasAverageFuelConsumption() bool`

HasAverageFuelConsumption returns a boolean if a field has been set.

### GetAverageElectricityConsumption

`func (o *VehicleParametersAtWaypoint) GetAverageElectricityConsumption() float64`

GetAverageElectricityConsumption returns the AverageElectricityConsumption field if non-nil, zero value otherwise.

### GetAverageElectricityConsumptionOk

`func (o *VehicleParametersAtWaypoint) GetAverageElectricityConsumptionOk() (*float64, bool)`

GetAverageElectricityConsumptionOk returns a tuple with the AverageElectricityConsumption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageElectricityConsumption

`func (o *VehicleParametersAtWaypoint) SetAverageElectricityConsumption(v float64)`

SetAverageElectricityConsumption sets AverageElectricityConsumption field to given value.

### HasAverageElectricityConsumption

`func (o *VehicleParametersAtWaypoint) HasAverageElectricityConsumption() bool`

HasAverageElectricityConsumption returns a boolean if a field has been set.

### GetLoadWeight

`func (o *VehicleParametersAtWaypoint) GetLoadWeight() int32`

GetLoadWeight returns the LoadWeight field if non-nil, zero value otherwise.

### GetLoadWeightOk

`func (o *VehicleParametersAtWaypoint) GetLoadWeightOk() (*int32, bool)`

GetLoadWeightOk returns a tuple with the LoadWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadWeight

`func (o *VehicleParametersAtWaypoint) SetLoadWeight(v int32)`

SetLoadWeight sets LoadWeight field to given value.

### HasLoadWeight

`func (o *VehicleParametersAtWaypoint) HasLoadWeight() bool`

HasLoadWeight returns a boolean if a field has been set.

### GetAxleWeight

`func (o *VehicleParametersAtWaypoint) GetAxleWeight() int32`

GetAxleWeight returns the AxleWeight field if non-nil, zero value otherwise.

### GetAxleWeightOk

`func (o *VehicleParametersAtWaypoint) GetAxleWeightOk() (*int32, bool)`

GetAxleWeightOk returns a tuple with the AxleWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAxleWeight

`func (o *VehicleParametersAtWaypoint) SetAxleWeight(v int32)`

SetAxleWeight sets AxleWeight field to given value.

### HasAxleWeight

`func (o *VehicleParametersAtWaypoint) HasAxleWeight() bool`

HasAxleWeight returns a boolean if a field has been set.

### GetHeight

`func (o *VehicleParametersAtWaypoint) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *VehicleParametersAtWaypoint) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *VehicleParametersAtWaypoint) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *VehicleParametersAtWaypoint) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetLength

`func (o *VehicleParametersAtWaypoint) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *VehicleParametersAtWaypoint) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *VehicleParametersAtWaypoint) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *VehicleParametersAtWaypoint) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetWidth

`func (o *VehicleParametersAtWaypoint) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *VehicleParametersAtWaypoint) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *VehicleParametersAtWaypoint) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *VehicleParametersAtWaypoint) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHazardousMaterials

`func (o *VehicleParametersAtWaypoint) GetHazardousMaterials() []HazardousMaterials`

GetHazardousMaterials returns the HazardousMaterials field if non-nil, zero value otherwise.

### GetHazardousMaterialsOk

`func (o *VehicleParametersAtWaypoint) GetHazardousMaterialsOk() (*[]HazardousMaterials, bool)`

GetHazardousMaterialsOk returns a tuple with the HazardousMaterials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHazardousMaterials

`func (o *VehicleParametersAtWaypoint) SetHazardousMaterials(v []HazardousMaterials)`

SetHazardousMaterials sets HazardousMaterials field to given value.

### HasHazardousMaterials

`func (o *VehicleParametersAtWaypoint) HasHazardousMaterials() bool`

HasHazardousMaterials returns a boolean if a field has been set.

### GetTunnelRestrictionCode

`func (o *VehicleParametersAtWaypoint) GetTunnelRestrictionCode() TunnelRestrictionCode`

GetTunnelRestrictionCode returns the TunnelRestrictionCode field if non-nil, zero value otherwise.

### GetTunnelRestrictionCodeOk

`func (o *VehicleParametersAtWaypoint) GetTunnelRestrictionCodeOk() (*TunnelRestrictionCode, bool)`

GetTunnelRestrictionCodeOk returns a tuple with the TunnelRestrictionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelRestrictionCode

`func (o *VehicleParametersAtWaypoint) SetTunnelRestrictionCode(v TunnelRestrictionCode)`

SetTunnelRestrictionCode sets TunnelRestrictionCode field to given value.

### HasTunnelRestrictionCode

`func (o *VehicleParametersAtWaypoint) HasTunnelRestrictionCode() bool`

HasTunnelRestrictionCode returns a boolean if a field has been set.

### GetTruckRoutes

`func (o *VehicleParametersAtWaypoint) GetTruckRoutes() []TruckRoutes`

GetTruckRoutes returns the TruckRoutes field if non-nil, zero value otherwise.

### GetTruckRoutesOk

`func (o *VehicleParametersAtWaypoint) GetTruckRoutesOk() (*[]TruckRoutes, bool)`

GetTruckRoutesOk returns a tuple with the TruckRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruckRoutes

`func (o *VehicleParametersAtWaypoint) SetTruckRoutes(v []TruckRoutes)`

SetTruckRoutes sets TruckRoutes field to given value.

### HasTruckRoutes

`func (o *VehicleParametersAtWaypoint) HasTruckRoutes() bool`

HasTruckRoutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


