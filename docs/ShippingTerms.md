# ShippingTerms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CarrierType** | Pointer to **NullableString** | Identifies whether to ship parcel or freight.              Parcel: Smaller, light weight boxes.              Freight: Larger boxes, usually transported by truckload. | [optional] 
**PaymentTerm** | Pointer to **NullableString** | Identifies the party responsible for shipping charges.              Collect: The person/entity receiving the product pays the shipping charges.              ThirdParty: Another party pays for the shipping charges (not Shipbob) [parcel only].              Prepaid: The shipper pays the shipping charges (Shipbob or merchant).              MerchantResponsible: The merchant will be responsible for uploading shipping labels or booking freight transportation. | [optional] 

## Methods

### NewShippingTerms

`func NewShippingTerms() *ShippingTerms`

NewShippingTerms instantiates a new ShippingTerms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingTermsWithDefaults

`func NewShippingTermsWithDefaults() *ShippingTerms`

NewShippingTermsWithDefaults instantiates a new ShippingTerms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrierType

`func (o *ShippingTerms) GetCarrierType() string`

GetCarrierType returns the CarrierType field if non-nil, zero value otherwise.

### GetCarrierTypeOk

`func (o *ShippingTerms) GetCarrierTypeOk() (*string, bool)`

GetCarrierTypeOk returns a tuple with the CarrierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierType

`func (o *ShippingTerms) SetCarrierType(v string)`

SetCarrierType sets CarrierType field to given value.

### HasCarrierType

`func (o *ShippingTerms) HasCarrierType() bool`

HasCarrierType returns a boolean if a field has been set.

### SetCarrierTypeNil

`func (o *ShippingTerms) SetCarrierTypeNil(b bool)`

 SetCarrierTypeNil sets the value for CarrierType to be an explicit nil

### UnsetCarrierType
`func (o *ShippingTerms) UnsetCarrierType()`

UnsetCarrierType ensures that no value is present for CarrierType, not even an explicit nil
### GetPaymentTerm

`func (o *ShippingTerms) GetPaymentTerm() string`

GetPaymentTerm returns the PaymentTerm field if non-nil, zero value otherwise.

### GetPaymentTermOk

`func (o *ShippingTerms) GetPaymentTermOk() (*string, bool)`

GetPaymentTermOk returns a tuple with the PaymentTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTerm

`func (o *ShippingTerms) SetPaymentTerm(v string)`

SetPaymentTerm sets PaymentTerm field to given value.

### HasPaymentTerm

`func (o *ShippingTerms) HasPaymentTerm() bool`

HasPaymentTerm returns a boolean if a field has been set.

### SetPaymentTermNil

`func (o *ShippingTerms) SetPaymentTermNil(b bool)`

 SetPaymentTermNil sets the value for PaymentTerm to be an explicit nil

### UnsetPaymentTerm
`func (o *ShippingTerms) UnsetPaymentTerm()`

UnsetPaymentTerm ensures that no value is present for PaymentTerm, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


