package models

import (
    "encoding/json"
)

// CreateAllocation represents a CreateAllocation struct.
type CreateAllocation struct {
    // The allocated quantity to which to set the line-items allocated quantity. By default, this is an integer. If decimal allocations are enabled for the component, it will be a decimal number. For On/Off components, use 1for on and 0 for off.
    Quantity                 float64               `json:"quantity"`
    // (required for the multiple allocations endpoint) The id associated with the component for which the allocation is being made
    ComponentId              *int                  `json:"component_id,omitempty"`
    // A memo to record along with the allocation
    Memo                     *string               `json:"memo,omitempty"`
    // The scheme used if the proration is a downgrade. Defaults to the site setting if one is not provided.
    ProrationDowngradeScheme *string               `json:"proration_downgrade_scheme,omitempty"` // Deprecated
    // The scheme used if the proration is an upgrade. Defaults to the site setting if one is not provided.
    ProrationUpgradeScheme   *string               `json:"proration_upgrade_scheme,omitempty"`   // Deprecated
    // If the change in cost is an upgrade, this determines if the charge should accrue to the next renewal or if capture should be attempted immediately. Defaults to the site setting if one is not provided.
    AccrueCharge             *bool                 `json:"accrue_charge,omitempty"`
    // The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
    // Available values: `full`, `prorated`, `none`.
    DowngradeCredit          Optional[CreditType]  `json:"downgrade_credit"`
    // The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
    // Available values: `full`, `prorated`, `none`.
    UpgradeCharge            Optional[CreditType]  `json:"upgrade_charge"`
    // If set to true, if the immediate component payment fails, initiate dunning for the subscription. 
    // Otherwise, leave the charges on the subscription to pay for at renewal. Defaults to false.
    InitiateDunning          *bool                 `json:"initiate_dunning,omitempty"`
    // Price point that the allocation should be charged at. Accepts either the price point's id (integer) or handle (string). When not specified, the default price point will be used.
    PricePointId             Optional[interface{}] `json:"price_point_id"`
    // This attribute is particularly useful when you need to align billing events for different components on distinct schedules within a subscription. Please note this only works for site with Multifrequency enabled
    BillingSchedule          *BillingSchedule      `json:"billing_schedule,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateAllocation.
// It customizes the JSON marshaling process for CreateAllocation objects.
func (c *CreateAllocation) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateAllocation object to a map representation for JSON marshaling.
func (c *CreateAllocation) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["quantity"] = c.Quantity
    if c.ComponentId != nil {
        structMap["component_id"] = c.ComponentId
    }
    if c.Memo != nil {
        structMap["memo"] = c.Memo
    }
    if c.ProrationDowngradeScheme != nil {
        structMap["proration_downgrade_scheme"] = c.ProrationDowngradeScheme
    }
    if c.ProrationUpgradeScheme != nil {
        structMap["proration_upgrade_scheme"] = c.ProrationUpgradeScheme
    }
    if c.AccrueCharge != nil {
        structMap["accrue_charge"] = c.AccrueCharge
    }
    if c.DowngradeCredit.IsValueSet() {
        structMap["downgrade_credit"] = c.DowngradeCredit.Value()
    }
    if c.UpgradeCharge.IsValueSet() {
        structMap["upgrade_charge"] = c.UpgradeCharge.Value()
    }
    if c.InitiateDunning != nil {
        structMap["initiate_dunning"] = c.InitiateDunning
    }
    if c.PricePointId.IsValueSet() {
        structMap["price_point_id"] = c.PricePointId.Value()
    }
    if c.BillingSchedule != nil {
        structMap["billing_schedule"] = c.BillingSchedule.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateAllocation.
// It customizes the JSON unmarshaling process for CreateAllocation objects.
func (c *CreateAllocation) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Quantity                 float64               `json:"quantity"`
        ComponentId              *int                  `json:"component_id,omitempty"`
        Memo                     *string               `json:"memo,omitempty"`
        ProrationDowngradeScheme *string               `json:"proration_downgrade_scheme,omitempty"`
        ProrationUpgradeScheme   *string               `json:"proration_upgrade_scheme,omitempty"`
        AccrueCharge             *bool                 `json:"accrue_charge,omitempty"`
        DowngradeCredit          Optional[CreditType]  `json:"downgrade_credit"`
        UpgradeCharge            Optional[CreditType]  `json:"upgrade_charge"`
        InitiateDunning          *bool                 `json:"initiate_dunning,omitempty"`
        PricePointId             Optional[interface{}] `json:"price_point_id"`
        BillingSchedule          *BillingSchedule      `json:"billing_schedule,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.Quantity = temp.Quantity
    c.ComponentId = temp.ComponentId
    c.Memo = temp.Memo
    c.ProrationDowngradeScheme = temp.ProrationDowngradeScheme
    c.ProrationUpgradeScheme = temp.ProrationUpgradeScheme
    c.AccrueCharge = temp.AccrueCharge
    c.DowngradeCredit = temp.DowngradeCredit
    c.UpgradeCharge = temp.UpgradeCharge
    c.InitiateDunning = temp.InitiateDunning
    c.PricePointId = temp.PricePointId
    c.BillingSchedule = temp.BillingSchedule
    return nil
}
