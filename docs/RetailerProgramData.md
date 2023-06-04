# RetailerProgramData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | Pointer to [**[]RetailerProgramDataAddress**](RetailerProgramDataAddress.md) | Ship From - Certain retailers want to display the ship from address as their return facility, not Shipbob&#39;s warehouse address        ///  Mark For Address - Final destination address | [optional] 
**CustomerTicketNumber** | Pointer to **string** | Customer Ticket Number | [optional] 
**DeliveryDate** | Pointer to **NullableTime** | Expected delivery date | [optional] 
**Department** | Pointer to **string** | Identifies a merchant&#39;s store department | [optional] 
**MarkForStore** | Pointer to **string** | Store Number | [optional] 
**PurchaseOrderNumber** | **string** | First initial documentation sent from buyer to seller with item(s) and quantities. | 
**RetailerProgramType** | **string** | Identifies retailer-merchant combination | 

## Methods

### NewRetailerProgramData

`func NewRetailerProgramData(purchaseOrderNumber string, retailerProgramType string, ) *RetailerProgramData`

NewRetailerProgramData instantiates a new RetailerProgramData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetailerProgramDataWithDefaults

`func NewRetailerProgramDataWithDefaults() *RetailerProgramData`

NewRetailerProgramDataWithDefaults instantiates a new RetailerProgramData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *RetailerProgramData) GetAddresses() []RetailerProgramDataAddress`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *RetailerProgramData) GetAddressesOk() (*[]RetailerProgramDataAddress, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *RetailerProgramData) SetAddresses(v []RetailerProgramDataAddress)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *RetailerProgramData) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetCustomerTicketNumber

`func (o *RetailerProgramData) GetCustomerTicketNumber() string`

GetCustomerTicketNumber returns the CustomerTicketNumber field if non-nil, zero value otherwise.

### GetCustomerTicketNumberOk

`func (o *RetailerProgramData) GetCustomerTicketNumberOk() (*string, bool)`

GetCustomerTicketNumberOk returns a tuple with the CustomerTicketNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerTicketNumber

`func (o *RetailerProgramData) SetCustomerTicketNumber(v string)`

SetCustomerTicketNumber sets CustomerTicketNumber field to given value.

### HasCustomerTicketNumber

`func (o *RetailerProgramData) HasCustomerTicketNumber() bool`

HasCustomerTicketNumber returns a boolean if a field has been set.

### GetDeliveryDate

`func (o *RetailerProgramData) GetDeliveryDate() time.Time`

GetDeliveryDate returns the DeliveryDate field if non-nil, zero value otherwise.

### GetDeliveryDateOk

`func (o *RetailerProgramData) GetDeliveryDateOk() (*time.Time, bool)`

GetDeliveryDateOk returns a tuple with the DeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDate

`func (o *RetailerProgramData) SetDeliveryDate(v time.Time)`

SetDeliveryDate sets DeliveryDate field to given value.

### HasDeliveryDate

`func (o *RetailerProgramData) HasDeliveryDate() bool`

HasDeliveryDate returns a boolean if a field has been set.

### SetDeliveryDateNil

`func (o *RetailerProgramData) SetDeliveryDateNil(b bool)`

 SetDeliveryDateNil sets the value for DeliveryDate to be an explicit nil

### UnsetDeliveryDate
`func (o *RetailerProgramData) UnsetDeliveryDate()`

UnsetDeliveryDate ensures that no value is present for DeliveryDate, not even an explicit nil
### GetDepartment

`func (o *RetailerProgramData) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *RetailerProgramData) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *RetailerProgramData) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *RetailerProgramData) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetMarkForStore

`func (o *RetailerProgramData) GetMarkForStore() string`

GetMarkForStore returns the MarkForStore field if non-nil, zero value otherwise.

### GetMarkForStoreOk

`func (o *RetailerProgramData) GetMarkForStoreOk() (*string, bool)`

GetMarkForStoreOk returns a tuple with the MarkForStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkForStore

`func (o *RetailerProgramData) SetMarkForStore(v string)`

SetMarkForStore sets MarkForStore field to given value.

### HasMarkForStore

`func (o *RetailerProgramData) HasMarkForStore() bool`

HasMarkForStore returns a boolean if a field has been set.

### GetPurchaseOrderNumber

`func (o *RetailerProgramData) GetPurchaseOrderNumber() string`

GetPurchaseOrderNumber returns the PurchaseOrderNumber field if non-nil, zero value otherwise.

### GetPurchaseOrderNumberOk

`func (o *RetailerProgramData) GetPurchaseOrderNumberOk() (*string, bool)`

GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderNumber

`func (o *RetailerProgramData) SetPurchaseOrderNumber(v string)`

SetPurchaseOrderNumber sets PurchaseOrderNumber field to given value.


### GetRetailerProgramType

`func (o *RetailerProgramData) GetRetailerProgramType() string`

GetRetailerProgramType returns the RetailerProgramType field if non-nil, zero value otherwise.

### GetRetailerProgramTypeOk

`func (o *RetailerProgramData) GetRetailerProgramTypeOk() (*string, bool)`

GetRetailerProgramTypeOk returns a tuple with the RetailerProgramType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetailerProgramType

`func (o *RetailerProgramData) SetRetailerProgramType(v string)`

SetRetailerProgramType sets RetailerProgramType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


