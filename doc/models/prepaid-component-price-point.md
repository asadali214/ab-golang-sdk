
# Prepaid Component Price Point

## Structure

`PrepaidComponentPricePoint`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `*string` | Optional | - |
| `Handle` | `*string` | Optional | - |
| `PricingScheme` | [`*models.PricingScheme`](../../doc/models/pricing-scheme.md) | Optional | The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes. |
| `Prices` | [`[]models.Price`](../../doc/models/price.md) | Optional | - |
| `OveragePricing` | [`*models.OveragePricing`](../../doc/models/overage-pricing.md) | Optional | - |

## Example (as JSON)

```json
{
  "name": "name8",
  "handle": "handle4",
  "pricing_scheme": "stairstep",
  "prices": [
    {
      "starting_quantity": {
        "key1": "val1",
        "key2": "val2"
      },
      "ending_quantity": {
        "key1": "val1",
        "key2": "val2"
      },
      "unit_price": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "starting_quantity": {
        "key1": "val1",
        "key2": "val2"
      },
      "ending_quantity": {
        "key1": "val1",
        "key2": "val2"
      },
      "unit_price": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "overage_pricing": {
    "pricing_scheme": "stairstep",
    "prices": [
      {
        "starting_quantity": {
          "key1": "val1",
          "key2": "val2"
        },
        "ending_quantity": {
          "key1": "val1",
          "key2": "val2"
        },
        "unit_price": {
          "key1": "val1",
          "key2": "val2"
        }
      }
    ]
  }
}
```

