package models

import (
    "encoding/json"
    "log"
    "time"
)

// ProductPricePoint represents a ProductPricePoint struct.
type ProductPricePoint struct {
    Id                      *int                `json:"id,omitempty"`
    // The product price point name
    Name                    *string             `json:"name,omitempty"`
    // The product price point API handle
    Handle                  *string             `json:"handle,omitempty"`
    // The product price point price, in integer cents
    PriceInCents            *int64              `json:"price_in_cents,omitempty"`
    // The numerical interval. i.e. an interval of ‘30’ coupled with an interval_unit of day would mean this product price point would renew every 30 days
    Interval                *int                `json:"interval,omitempty"`
    // A string representing the interval unit for this product price point, either month or day
    IntervalUnit            *IntervalUnit       `json:"interval_unit,omitempty"`
    // The product price point trial price, in integer cents
    TrialPriceInCents       *int64              `json:"trial_price_in_cents,omitempty"`
    // The numerical trial interval. i.e. an interval of ‘30’ coupled with a trial_interval_unit of day would mean this product price point trial would last 30 days
    TrialInterval           *int                `json:"trial_interval,omitempty"`
    // A string representing the trial interval unit for this product price point, either month or day
    TrialIntervalUnit       *IntervalUnit       `json:"trial_interval_unit,omitempty"`
    TrialType               *string             `json:"trial_type,omitempty"`
    // reserved for future use
    IntroductoryOffer       *bool               `json:"introductory_offer,omitempty"`
    // The product price point initial charge, in integer cents
    InitialChargeInCents    *int64              `json:"initial_charge_in_cents,omitempty"`
    InitialChargeAfterTrial *bool               `json:"initial_charge_after_trial,omitempty"`
    // The numerical expiration interval. i.e. an expiration_interval of ‘30’ coupled with an expiration_interval_unit of day would mean this product price point would expire after 30 days
    ExpirationInterval      *int                `json:"expiration_interval,omitempty"`
    // A string representing the expiration interval unit for this product price point, either month or day
    ExpirationIntervalUnit  *IntervalUnit       `json:"expiration_interval_unit,omitempty"`
    // The product id this price point belongs to
    ProductId               *int                `json:"product_id,omitempty"`
    // Timestamp indicating when this price point was archived
    ArchivedAt              Optional[time.Time] `json:"archived_at"`
    // Timestamp indicating when this price point was created
    CreatedAt               *time.Time          `json:"created_at,omitempty"`
    // Timestamp indicating when this price point was last updated
    UpdatedAt               *time.Time          `json:"updated_at,omitempty"`
    // Whether or not to use the site's exchange rate or define your own pricing when your site has multiple currencies defined.
    UseSiteExchangeRate     *bool               `json:"use_site_exchange_rate,omitempty"`
    // The type of price point
    Type                    *PricePointType     `json:"type,omitempty"`
    // Whether or not the price point includes tax
    TaxIncluded             *bool               `json:"tax_included,omitempty"`
    // The subscription id this price point belongs to
    SubscriptionId          Optional[int]       `json:"subscription_id"`
    // An array of currency pricing data is available when multiple currencies are defined for the site. It varies based on the use_site_exchange_rate setting for the price point. This parameter is present only in the response of read endpoints, after including the appropriate query parameter.
    CurrencyPrices          []CurrencyPrice     `json:"currency_prices,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ProductPricePoint.
// It customizes the JSON marshaling process for ProductPricePoint objects.
func (p *ProductPricePoint) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the ProductPricePoint object to a map representation for JSON marshaling.
func (p *ProductPricePoint) toMap() map[string]any {
    structMap := make(map[string]any)
    if p.Id != nil {
        structMap["id"] = p.Id
    }
    if p.Name != nil {
        structMap["name"] = p.Name
    }
    if p.Handle != nil {
        structMap["handle"] = p.Handle
    }
    if p.PriceInCents != nil {
        structMap["price_in_cents"] = p.PriceInCents
    }
    if p.Interval != nil {
        structMap["interval"] = p.Interval
    }
    if p.IntervalUnit != nil {
        structMap["interval_unit"] = p.IntervalUnit
    }
    if p.TrialPriceInCents != nil {
        structMap["trial_price_in_cents"] = p.TrialPriceInCents
    }
    if p.TrialInterval != nil {
        structMap["trial_interval"] = p.TrialInterval
    }
    if p.TrialIntervalUnit != nil {
        structMap["trial_interval_unit"] = p.TrialIntervalUnit
    }
    if p.TrialType != nil {
        structMap["trial_type"] = p.TrialType
    }
    if p.IntroductoryOffer != nil {
        structMap["introductory_offer"] = p.IntroductoryOffer
    }
    if p.InitialChargeInCents != nil {
        structMap["initial_charge_in_cents"] = p.InitialChargeInCents
    }
    if p.InitialChargeAfterTrial != nil {
        structMap["initial_charge_after_trial"] = p.InitialChargeAfterTrial
    }
    if p.ExpirationInterval != nil {
        structMap["expiration_interval"] = p.ExpirationInterval
    }
    if p.ExpirationIntervalUnit != nil {
        structMap["expiration_interval_unit"] = p.ExpirationIntervalUnit
    }
    if p.ProductId != nil {
        structMap["product_id"] = p.ProductId
    }
    if p.ArchivedAt.IsValueSet() {
        var ArchivedAtVal *string = nil
        if p.ArchivedAt.Value() != nil {
            val := p.ArchivedAt.Value().Format(time.RFC3339)
            ArchivedAtVal = &val
        }
        structMap["archived_at"] = ArchivedAtVal
    }
    if p.CreatedAt != nil {
        structMap["created_at"] = p.CreatedAt.Format(time.RFC3339)
    }
    if p.UpdatedAt != nil {
        structMap["updated_at"] = p.UpdatedAt.Format(time.RFC3339)
    }
    if p.UseSiteExchangeRate != nil {
        structMap["use_site_exchange_rate"] = p.UseSiteExchangeRate
    }
    if p.Type != nil {
        structMap["type"] = p.Type
    }
    if p.TaxIncluded != nil {
        structMap["tax_included"] = p.TaxIncluded
    }
    if p.SubscriptionId.IsValueSet() {
        structMap["subscription_id"] = p.SubscriptionId.Value()
    }
    if p.CurrencyPrices != nil {
        structMap["currency_prices"] = p.CurrencyPrices
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ProductPricePoint.
// It customizes the JSON unmarshaling process for ProductPricePoint objects.
func (p *ProductPricePoint) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id                      *int             `json:"id,omitempty"`
        Name                    *string          `json:"name,omitempty"`
        Handle                  *string          `json:"handle,omitempty"`
        PriceInCents            *int64           `json:"price_in_cents,omitempty"`
        Interval                *int             `json:"interval,omitempty"`
        IntervalUnit            *IntervalUnit    `json:"interval_unit,omitempty"`
        TrialPriceInCents       *int64           `json:"trial_price_in_cents,omitempty"`
        TrialInterval           *int             `json:"trial_interval,omitempty"`
        TrialIntervalUnit       *IntervalUnit    `json:"trial_interval_unit,omitempty"`
        TrialType               *string          `json:"trial_type,omitempty"`
        IntroductoryOffer       *bool            `json:"introductory_offer,omitempty"`
        InitialChargeInCents    *int64           `json:"initial_charge_in_cents,omitempty"`
        InitialChargeAfterTrial *bool            `json:"initial_charge_after_trial,omitempty"`
        ExpirationInterval      *int             `json:"expiration_interval,omitempty"`
        ExpirationIntervalUnit  *IntervalUnit    `json:"expiration_interval_unit,omitempty"`
        ProductId               *int             `json:"product_id,omitempty"`
        ArchivedAt              Optional[string] `json:"archived_at"`
        CreatedAt               *string          `json:"created_at,omitempty"`
        UpdatedAt               *string          `json:"updated_at,omitempty"`
        UseSiteExchangeRate     *bool            `json:"use_site_exchange_rate,omitempty"`
        Type                    *PricePointType  `json:"type,omitempty"`
        TaxIncluded             *bool            `json:"tax_included,omitempty"`
        SubscriptionId          Optional[int]    `json:"subscription_id"`
        CurrencyPrices          []CurrencyPrice  `json:"currency_prices,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    p.Id = temp.Id
    p.Name = temp.Name
    p.Handle = temp.Handle
    p.PriceInCents = temp.PriceInCents
    p.Interval = temp.Interval
    p.IntervalUnit = temp.IntervalUnit
    p.TrialPriceInCents = temp.TrialPriceInCents
    p.TrialInterval = temp.TrialInterval
    p.TrialIntervalUnit = temp.TrialIntervalUnit
    p.TrialType = temp.TrialType
    p.IntroductoryOffer = temp.IntroductoryOffer
    p.InitialChargeInCents = temp.InitialChargeInCents
    p.InitialChargeAfterTrial = temp.InitialChargeAfterTrial
    p.ExpirationInterval = temp.ExpirationInterval
    p.ExpirationIntervalUnit = temp.ExpirationIntervalUnit
    p.ProductId = temp.ProductId
    p.ArchivedAt.ShouldSetValue(temp.ArchivedAt.IsValueSet())
    if temp.ArchivedAt.Value() != nil {
        ArchivedAtVal, err := time.Parse(time.RFC3339, (*temp.ArchivedAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse archived_at as % s format.", time.RFC3339)
        }
        p.ArchivedAt.SetValue(&ArchivedAtVal)
    }
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        p.CreatedAt = &CreatedAtVal
    }
    if temp.UpdatedAt != nil {
        UpdatedAtVal, err := time.Parse(time.RFC3339, *temp.UpdatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
        }
        p.UpdatedAt = &UpdatedAtVal
    }
    p.UseSiteExchangeRate = temp.UseSiteExchangeRate
    p.Type = temp.Type
    p.TaxIncluded = temp.TaxIncluded
    p.SubscriptionId = temp.SubscriptionId
    p.CurrencyPrices = temp.CurrencyPrices
    return nil
}
