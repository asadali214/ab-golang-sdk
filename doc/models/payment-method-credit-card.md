
# Payment Method Credit Card

## Structure

`PaymentMethodCreditCard`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CardBrand` | `string` | Required | - |
| `CardExpiration` | `*string` | Optional | - |
| `LastFour` | `Optional[string]` | Optional | - |
| `MaskedCardNumber` | `string` | Required | - |
| `Type` | [`models.InvoiceEventPaymentMethod`](../../doc/models/invoice-event-payment-method.md) | Required | - |

## Example (as JSON)

```json
{
  "card_brand": "card_brand4",
  "card_expiration": "card_expiration2",
  "last_four": "last_four4",
  "masked_card_number": "masked_card_number0",
  "type": "external"
}
```

