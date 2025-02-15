package models

import (
    "encoding/json"
)

// CustomerAttributes represents a CustomerAttributes struct.
type CustomerAttributes struct {
    // The first name of the customer. Required when creating a customer via attributes.
    FirstName    *string           `json:"first_name,omitempty"`
    // The last name of the customer. Required when creating a customer via attributes.
    LastName     *string           `json:"last_name,omitempty"`
    // The email address of the customer. Required when creating a customer via attributes.
    Email        *string           `json:"email,omitempty"`
    // A list of emails that should be cc’d on all customer communications. Optional.
    CcEmails     *string           `json:"cc_emails,omitempty"`
    // The organization/company of the customer. Optional.
    Organization *string           `json:"organization,omitempty"`
    // A customer “reference”, or unique identifier from your app, stored in Chargify. Can be used so that you may reference your customer’s within Chargify using the same unique value you use in your application. Optional.
    Reference    *string           `json:"reference,omitempty"`
    // (Optional) The customer’s shipping street address (i.e. “123 Main St.”).
    Address      *string           `json:"address,omitempty"`
    // (Optional) Second line of the customer’s shipping address i.e. “Apt. 100”
    Address2     Optional[string]  `json:"address_2"`
    // (Optional) The customer’s shipping address city (i.e. “Boston”).
    City         *string           `json:"city,omitempty"`
    // (Optional) The customer’s shipping address state (i.e. “MA”). This must conform to the [ISO_3166-1](https://en.wikipedia.org/wiki/ISO_3166-1#Current_codes) in order to be valid for tax locale purposes.
    State        *string           `json:"state,omitempty"`
    // (Optional) The customer’s shipping address zip code (i.e. “12345”).
    Zip          *string           `json:"zip,omitempty"`
    // (Optional) The customer shipping address country, required in [ISO_3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) format (i.e. “US”).
    Country      *string           `json:"country,omitempty"`
    // (Optional) The phone number of the customer.
    Phone        *string           `json:"phone,omitempty"`
    Verified     *bool             `json:"verified,omitempty"`
    // (Optional) The tax_exempt status of the customer. Acceptable values are true or 1 for true and false or 0 for false.
    TaxExempt    *bool             `json:"tax_exempt,omitempty"`
    // (Optional) Supplying the VAT number allows EU customer’s to opt-out of the Value Added Tax assuming the merchant address and customer billing address are not within the same EU country. It’s important to omit the country code from the VAT number upon entry. Otherwise, taxes will be assessed upon the purchase.
    VatNumber    *string           `json:"vat_number,omitempty"`
    // (Optional) A set of key/value pairs representing custom fields and their values. Metafields will be created “on-the-fly” in your site for a given key, if they have not been created yet.
    Metafields   map[string]string `json:"metafields,omitempty"`
    // The parent ID in Chargify if applicable. Parent is another Customer object.
    ParentId     Optional[int]     `json:"parent_id"`
}

// MarshalJSON implements the json.Marshaler interface for CustomerAttributes.
// It customizes the JSON marshaling process for CustomerAttributes objects.
func (c *CustomerAttributes) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CustomerAttributes object to a map representation for JSON marshaling.
func (c *CustomerAttributes) toMap() map[string]any {
    structMap := make(map[string]any)
    if c.FirstName != nil {
        structMap["first_name"] = c.FirstName
    }
    if c.LastName != nil {
        structMap["last_name"] = c.LastName
    }
    if c.Email != nil {
        structMap["email"] = c.Email
    }
    if c.CcEmails != nil {
        structMap["cc_emails"] = c.CcEmails
    }
    if c.Organization != nil {
        structMap["organization"] = c.Organization
    }
    if c.Reference != nil {
        structMap["reference"] = c.Reference
    }
    if c.Address != nil {
        structMap["address"] = c.Address
    }
    if c.Address2.IsValueSet() {
        structMap["address_2"] = c.Address2.Value()
    }
    if c.City != nil {
        structMap["city"] = c.City
    }
    if c.State != nil {
        structMap["state"] = c.State
    }
    if c.Zip != nil {
        structMap["zip"] = c.Zip
    }
    if c.Country != nil {
        structMap["country"] = c.Country
    }
    if c.Phone != nil {
        structMap["phone"] = c.Phone
    }
    if c.Verified != nil {
        structMap["verified"] = c.Verified
    }
    if c.TaxExempt != nil {
        structMap["tax_exempt"] = c.TaxExempt
    }
    if c.VatNumber != nil {
        structMap["vat_number"] = c.VatNumber
    }
    if c.Metafields != nil {
        structMap["metafields"] = c.Metafields
    }
    if c.ParentId.IsValueSet() {
        structMap["parent_id"] = c.ParentId.Value()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CustomerAttributes.
// It customizes the JSON unmarshaling process for CustomerAttributes objects.
func (c *CustomerAttributes) UnmarshalJSON(input []byte) error {
    temp := &struct {
        FirstName    *string           `json:"first_name,omitempty"`
        LastName     *string           `json:"last_name,omitempty"`
        Email        *string           `json:"email,omitempty"`
        CcEmails     *string           `json:"cc_emails,omitempty"`
        Organization *string           `json:"organization,omitempty"`
        Reference    *string           `json:"reference,omitempty"`
        Address      *string           `json:"address,omitempty"`
        Address2     Optional[string]  `json:"address_2"`
        City         *string           `json:"city,omitempty"`
        State        *string           `json:"state,omitempty"`
        Zip          *string           `json:"zip,omitempty"`
        Country      *string           `json:"country,omitempty"`
        Phone        *string           `json:"phone,omitempty"`
        Verified     *bool             `json:"verified,omitempty"`
        TaxExempt    *bool             `json:"tax_exempt,omitempty"`
        VatNumber    *string           `json:"vat_number,omitempty"`
        Metafields   map[string]string `json:"metafields,omitempty"`
        ParentId     Optional[int]     `json:"parent_id"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.FirstName = temp.FirstName
    c.LastName = temp.LastName
    c.Email = temp.Email
    c.CcEmails = temp.CcEmails
    c.Organization = temp.Organization
    c.Reference = temp.Reference
    c.Address = temp.Address
    c.Address2 = temp.Address2
    c.City = temp.City
    c.State = temp.State
    c.Zip = temp.Zip
    c.Country = temp.Country
    c.Phone = temp.Phone
    c.Verified = temp.Verified
    c.TaxExempt = temp.TaxExempt
    c.VatNumber = temp.VatNumber
    c.Metafields = temp.Metafields
    c.ParentId = temp.ParentId
    return nil
}
