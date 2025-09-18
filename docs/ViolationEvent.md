# ViolationEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ViolationType**](ViolationType.md) |  | 
**ScheduleViolationTypes** | Pointer to [**[]ScheduleViolationType**](ScheduleViolationType.md) | Contains the list of violated schedule restrictions at this event. Only present if **type&#x3D;SCHEDULE**.  | [optional] 
**ViolatedVehicleProperties** | Pointer to [**[]ViolatedVehicleProperty**](ViolatedVehicleProperty.md) | Contains the list of violated vehicle properties at this event. Only present if **type&#x3D;VEHICLE_PROPERTY**.  | [optional] 
**Temporary** | Pointer to **bool** | States whether the cause of the violation is only temporary, e.g. a road construction site. In such cases  see the corresponding traffic event for more information.  | [optional] 
**TimeDomain** | Pointer to **string** | Returns the time validity of the violated restriction in the GDF time domain format if the restriction is only active at certain times. See [here](./concepts/date-and-time#gdftimedomains) for more information. | [optional] 
**AccessType** | Pointer to [**AccessType**](AccessType.md) |  | [optional] 
**RelatedEventIndex** | Pointer to **int32** | For **accessType** _ENTER_ this index points to the corresponding event with **accessType** _EXIT_ and vice-versa. Not present otherwise. | [optional] 
**Polyline** | Pointer to **string** | The polyline of the violation in the format specified by **options[polylineFormat]**. Only present for **accessType** _ENTER_. Requires _VIOLATION_EVENTS_POLYLINE_ to be requested. | [optional] 

## Methods

### NewViolationEvent

`func NewViolationEvent(type_ ViolationType, ) *ViolationEvent`

NewViolationEvent instantiates a new ViolationEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViolationEventWithDefaults

`func NewViolationEventWithDefaults() *ViolationEvent`

NewViolationEventWithDefaults instantiates a new ViolationEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ViolationEvent) GetType() ViolationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ViolationEvent) GetTypeOk() (*ViolationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ViolationEvent) SetType(v ViolationType)`

SetType sets Type field to given value.


### GetScheduleViolationTypes

`func (o *ViolationEvent) GetScheduleViolationTypes() []ScheduleViolationType`

GetScheduleViolationTypes returns the ScheduleViolationTypes field if non-nil, zero value otherwise.

### GetScheduleViolationTypesOk

`func (o *ViolationEvent) GetScheduleViolationTypesOk() (*[]ScheduleViolationType, bool)`

GetScheduleViolationTypesOk returns a tuple with the ScheduleViolationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleViolationTypes

`func (o *ViolationEvent) SetScheduleViolationTypes(v []ScheduleViolationType)`

SetScheduleViolationTypes sets ScheduleViolationTypes field to given value.

### HasScheduleViolationTypes

`func (o *ViolationEvent) HasScheduleViolationTypes() bool`

HasScheduleViolationTypes returns a boolean if a field has been set.

### GetViolatedVehicleProperties

`func (o *ViolationEvent) GetViolatedVehicleProperties() []ViolatedVehicleProperty`

GetViolatedVehicleProperties returns the ViolatedVehicleProperties field if non-nil, zero value otherwise.

### GetViolatedVehiclePropertiesOk

`func (o *ViolationEvent) GetViolatedVehiclePropertiesOk() (*[]ViolatedVehicleProperty, bool)`

GetViolatedVehiclePropertiesOk returns a tuple with the ViolatedVehicleProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolatedVehicleProperties

`func (o *ViolationEvent) SetViolatedVehicleProperties(v []ViolatedVehicleProperty)`

SetViolatedVehicleProperties sets ViolatedVehicleProperties field to given value.

### HasViolatedVehicleProperties

`func (o *ViolationEvent) HasViolatedVehicleProperties() bool`

HasViolatedVehicleProperties returns a boolean if a field has been set.

### GetTemporary

`func (o *ViolationEvent) GetTemporary() bool`

GetTemporary returns the Temporary field if non-nil, zero value otherwise.

### GetTemporaryOk

`func (o *ViolationEvent) GetTemporaryOk() (*bool, bool)`

GetTemporaryOk returns a tuple with the Temporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporary

`func (o *ViolationEvent) SetTemporary(v bool)`

SetTemporary sets Temporary field to given value.

### HasTemporary

`func (o *ViolationEvent) HasTemporary() bool`

HasTemporary returns a boolean if a field has been set.

### GetTimeDomain

`func (o *ViolationEvent) GetTimeDomain() string`

GetTimeDomain returns the TimeDomain field if non-nil, zero value otherwise.

### GetTimeDomainOk

`func (o *ViolationEvent) GetTimeDomainOk() (*string, bool)`

GetTimeDomainOk returns a tuple with the TimeDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeDomain

`func (o *ViolationEvent) SetTimeDomain(v string)`

SetTimeDomain sets TimeDomain field to given value.

### HasTimeDomain

`func (o *ViolationEvent) HasTimeDomain() bool`

HasTimeDomain returns a boolean if a field has been set.

### GetAccessType

`func (o *ViolationEvent) GetAccessType() AccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *ViolationEvent) GetAccessTypeOk() (*AccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *ViolationEvent) SetAccessType(v AccessType)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *ViolationEvent) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetRelatedEventIndex

`func (o *ViolationEvent) GetRelatedEventIndex() int32`

GetRelatedEventIndex returns the RelatedEventIndex field if non-nil, zero value otherwise.

### GetRelatedEventIndexOk

`func (o *ViolationEvent) GetRelatedEventIndexOk() (*int32, bool)`

GetRelatedEventIndexOk returns a tuple with the RelatedEventIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedEventIndex

`func (o *ViolationEvent) SetRelatedEventIndex(v int32)`

SetRelatedEventIndex sets RelatedEventIndex field to given value.

### HasRelatedEventIndex

`func (o *ViolationEvent) HasRelatedEventIndex() bool`

HasRelatedEventIndex returns a boolean if a field has been set.

### GetPolyline

`func (o *ViolationEvent) GetPolyline() string`

GetPolyline returns the Polyline field if non-nil, zero value otherwise.

### GetPolylineOk

`func (o *ViolationEvent) GetPolylineOk() (*string, bool)`

GetPolylineOk returns a tuple with the Polyline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolyline

`func (o *ViolationEvent) SetPolyline(v string)`

SetPolyline sets Polyline field to given value.

### HasPolyline

`func (o *ViolationEvent) HasPolyline() bool`

HasPolyline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


