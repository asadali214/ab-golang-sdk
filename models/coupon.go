package models

import (
    "encoding/json"
    "log"
    "time"
)

// Coupon represents a Coupon struct.
type Coupon struct {
    Id                            *int                `json:"id,omitempty"`
    Name                          *string             `json:"name,omitempty"`
    Code                          *string             `json:"code,omitempty"`
    Description                   *string             `json:"description,omitempty"`
    Amount                        Optional[float64]   `json:"amount"`
    AmountInCents                 Optional[int]       `json:"amount_in_cents"`
    ProductFamilyId               *int                `json:"product_family_id,omitempty"`
    ProductFamilyName             Optional[string]    `json:"product_family_name"`
    StartDate                     *time.Time          `json:"start_date,omitempty"`
    EndDate                       Optional[time.Time] `json:"end_date"`
    Percentage                    Optional[string]    `json:"percentage"`
    Recurring                     *bool               `json:"recurring,omitempty"`
    RecurringScheme               *RecurringScheme    `json:"recurring_scheme,omitempty"`
    DurationPeriodCount           Optional[int]       `json:"duration_period_count"`
    DurationInterval              Optional[int]       `json:"duration_interval"`
    DurationIntervalUnit          Optional[string]    `json:"duration_interval_unit"`
    DurationIntervalSpan          Optional[string]    `json:"duration_interval_span"`
    AllowNegativeBalance          *bool               `json:"allow_negative_balance,omitempty"`
    ArchivedAt                    Optional[time.Time] `json:"archived_at"`
    ConversionLimit               Optional[string]    `json:"conversion_limit"`
    Stackable                     *bool               `json:"stackable,omitempty"`
    CompoundingStrategy           *interface{}        `json:"compounding_strategy,omitempty"`
    UseSiteExchangeRate           *bool               `json:"use_site_exchange_rate,omitempty"`
    CreatedAt                     *time.Time          `json:"created_at,omitempty"`
    UpdatedAt                     *time.Time          `json:"updated_at,omitempty"`
    DiscountType                  *DiscountType       `json:"discount_type,omitempty"`
    ExcludeMidPeriodAllocations   *bool               `json:"exclude_mid_period_allocations,omitempty"`
    ApplyOnCancelAtEndOfPeriod    *bool               `json:"apply_on_cancel_at_end_of_period,omitempty"`
    ApplyOnSubscriptionExpiration *bool               `json:"apply_on_subscription_expiration,omitempty"`
    CouponRestrictions            []CouponRestriction `json:"coupon_restrictions,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for Coupon.
// It customizes the JSON marshaling process for Coupon objects.
func (c *Coupon) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the Coupon object to a map representation for JSON marshaling.
func (c *Coupon) toMap() map[string]any {
    structMap := make(map[string]any)
    if c.Id != nil {
        structMap["id"] = c.Id
    }
    if c.Name != nil {
        structMap["name"] = c.Name
    }
    if c.Code != nil {
        structMap["code"] = c.Code
    }
    if c.Description != nil {
        structMap["description"] = c.Description
    }
    if c.Amount.IsValueSet() {
        structMap["amount"] = c.Amount.Value()
    }
    if c.AmountInCents.IsValueSet() {
        structMap["amount_in_cents"] = c.AmountInCents.Value()
    }
    if c.ProductFamilyId != nil {
        structMap["product_family_id"] = c.ProductFamilyId
    }
    if c.ProductFamilyName.IsValueSet() {
        structMap["product_family_name"] = c.ProductFamilyName.Value()
    }
    if c.StartDate != nil {
        structMap["start_date"] = c.StartDate.Format(time.RFC3339)
    }
    if c.EndDate.IsValueSet() {
        var EndDateVal *string = nil
        if c.EndDate.Value() != nil {
            val := c.EndDate.Value().Format(time.RFC3339)
            EndDateVal = &val
        }
        structMap["end_date"] = EndDateVal
    }
    if c.Percentage.IsValueSet() {
        structMap["percentage"] = c.Percentage.Value()
    }
    if c.Recurring != nil {
        structMap["recurring"] = c.Recurring
    }
    if c.RecurringScheme != nil {
        structMap["recurring_scheme"] = c.RecurringScheme
    }
    if c.DurationPeriodCount.IsValueSet() {
        structMap["duration_period_count"] = c.DurationPeriodCount.Value()
    }
    if c.DurationInterval.IsValueSet() {
        structMap["duration_interval"] = c.DurationInterval.Value()
    }
    if c.DurationIntervalUnit.IsValueSet() {
        structMap["duration_interval_unit"] = c.DurationIntervalUnit.Value()
    }
    if c.DurationIntervalSpan.IsValueSet() {
        structMap["duration_interval_span"] = c.DurationIntervalSpan.Value()
    }
    if c.AllowNegativeBalance != nil {
        structMap["allow_negative_balance"] = c.AllowNegativeBalance
    }
    if c.ArchivedAt.IsValueSet() {
        var ArchivedAtVal *string = nil
        if c.ArchivedAt.Value() != nil {
            val := c.ArchivedAt.Value().Format(time.RFC3339)
            ArchivedAtVal = &val
        }
        structMap["archived_at"] = ArchivedAtVal
    }
    if c.ConversionLimit.IsValueSet() {
        structMap["conversion_limit"] = c.ConversionLimit.Value()
    }
    if c.Stackable != nil {
        structMap["stackable"] = c.Stackable
    }
    if c.CompoundingStrategy != nil {
        structMap["compounding_strategy"] = c.CompoundingStrategy
    }
    if c.UseSiteExchangeRate != nil {
        structMap["use_site_exchange_rate"] = c.UseSiteExchangeRate
    }
    if c.CreatedAt != nil {
        structMap["created_at"] = c.CreatedAt.Format(time.RFC3339)
    }
    if c.UpdatedAt != nil {
        structMap["updated_at"] = c.UpdatedAt.Format(time.RFC3339)
    }
    if c.DiscountType != nil {
        structMap["discount_type"] = c.DiscountType
    }
    if c.ExcludeMidPeriodAllocations != nil {
        structMap["exclude_mid_period_allocations"] = c.ExcludeMidPeriodAllocations
    }
    if c.ApplyOnCancelAtEndOfPeriod != nil {
        structMap["apply_on_cancel_at_end_of_period"] = c.ApplyOnCancelAtEndOfPeriod
    }
    if c.ApplyOnSubscriptionExpiration != nil {
        structMap["apply_on_subscription_expiration"] = c.ApplyOnSubscriptionExpiration
    }
    if c.CouponRestrictions != nil {
        structMap["coupon_restrictions"] = c.CouponRestrictions
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Coupon.
// It customizes the JSON unmarshaling process for Coupon objects.
func (c *Coupon) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id                            *int                `json:"id,omitempty"`
        Name                          *string             `json:"name,omitempty"`
        Code                          *string             `json:"code,omitempty"`
        Description                   *string             `json:"description,omitempty"`
        Amount                        Optional[float64]   `json:"amount"`
        AmountInCents                 Optional[int]       `json:"amount_in_cents"`
        ProductFamilyId               *int                `json:"product_family_id,omitempty"`
        ProductFamilyName             Optional[string]    `json:"product_family_name"`
        StartDate                     *string             `json:"start_date,omitempty"`
        EndDate                       Optional[string]    `json:"end_date"`
        Percentage                    Optional[string]    `json:"percentage"`
        Recurring                     *bool               `json:"recurring,omitempty"`
        RecurringScheme               *RecurringScheme    `json:"recurring_scheme,omitempty"`
        DurationPeriodCount           Optional[int]       `json:"duration_period_count"`
        DurationInterval              Optional[int]       `json:"duration_interval"`
        DurationIntervalUnit          Optional[string]    `json:"duration_interval_unit"`
        DurationIntervalSpan          Optional[string]    `json:"duration_interval_span"`
        AllowNegativeBalance          *bool               `json:"allow_negative_balance,omitempty"`
        ArchivedAt                    Optional[string]    `json:"archived_at"`
        ConversionLimit               Optional[string]    `json:"conversion_limit"`
        Stackable                     *bool               `json:"stackable,omitempty"`
        CompoundingStrategy           *interface{}        `json:"compounding_strategy,omitempty"`
        UseSiteExchangeRate           *bool               `json:"use_site_exchange_rate,omitempty"`
        CreatedAt                     *string             `json:"created_at,omitempty"`
        UpdatedAt                     *string             `json:"updated_at,omitempty"`
        DiscountType                  *DiscountType       `json:"discount_type,omitempty"`
        ExcludeMidPeriodAllocations   *bool               `json:"exclude_mid_period_allocations,omitempty"`
        ApplyOnCancelAtEndOfPeriod    *bool               `json:"apply_on_cancel_at_end_of_period,omitempty"`
        ApplyOnSubscriptionExpiration *bool               `json:"apply_on_subscription_expiration,omitempty"`
        CouponRestrictions            []CouponRestriction `json:"coupon_restrictions,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.Id = temp.Id
    c.Name = temp.Name
    c.Code = temp.Code
    c.Description = temp.Description
    c.Amount = temp.Amount
    c.AmountInCents = temp.AmountInCents
    c.ProductFamilyId = temp.ProductFamilyId
    c.ProductFamilyName = temp.ProductFamilyName
    if temp.StartDate != nil {
        StartDateVal, err := time.Parse(time.RFC3339, *temp.StartDate)
        if err != nil {
            log.Fatalf("Cannot Parse start_date as % s format.", time.RFC3339)
        }
        c.StartDate = &StartDateVal
    }
    c.EndDate.ShouldSetValue(temp.EndDate.IsValueSet())
    if temp.EndDate.Value() != nil {
        EndDateVal, err := time.Parse(time.RFC3339, (*temp.EndDate.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse end_date as % s format.", time.RFC3339)
        }
        c.EndDate.SetValue(&EndDateVal)
    }
    c.Percentage = temp.Percentage
    c.Recurring = temp.Recurring
    c.RecurringScheme = temp.RecurringScheme
    c.DurationPeriodCount = temp.DurationPeriodCount
    c.DurationInterval = temp.DurationInterval
    c.DurationIntervalUnit = temp.DurationIntervalUnit
    c.DurationIntervalSpan = temp.DurationIntervalSpan
    c.AllowNegativeBalance = temp.AllowNegativeBalance
    c.ArchivedAt.ShouldSetValue(temp.ArchivedAt.IsValueSet())
    if temp.ArchivedAt.Value() != nil {
        ArchivedAtVal, err := time.Parse(time.RFC3339, (*temp.ArchivedAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse archived_at as % s format.", time.RFC3339)
        }
        c.ArchivedAt.SetValue(&ArchivedAtVal)
    }
    c.ConversionLimit = temp.ConversionLimit
    c.Stackable = temp.Stackable
    c.CompoundingStrategy = temp.CompoundingStrategy
    c.UseSiteExchangeRate = temp.UseSiteExchangeRate
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        c.CreatedAt = &CreatedAtVal
    }
    if temp.UpdatedAt != nil {
        UpdatedAtVal, err := time.Parse(time.RFC3339, *temp.UpdatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
        }
        c.UpdatedAt = &UpdatedAtVal
    }
    c.DiscountType = temp.DiscountType
    c.ExcludeMidPeriodAllocations = temp.ExcludeMidPeriodAllocations
    c.ApplyOnCancelAtEndOfPeriod = temp.ApplyOnCancelAtEndOfPeriod
    c.ApplyOnSubscriptionExpiration = temp.ApplyOnSubscriptionExpiration
    c.CouponRestrictions = temp.CouponRestrictions
    return nil
}
