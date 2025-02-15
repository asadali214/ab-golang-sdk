
# Allocation Preview

## Structure

`AllocationPreview`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `StartDate` | `*time.Time` | Optional | - |
| `EndDate` | `*time.Time` | Optional | - |
| `SubtotalInCents` | `*int64` | Optional | - |
| `TotalTaxInCents` | `*int64` | Optional | - |
| `TotalDiscountInCents` | `*int64` | Optional | - |
| `TotalInCents` | `*int64` | Optional | - |
| `Direction` | [`*models.AllocationPreviewDirection`](../../doc/models/allocation-preview-direction.md) | Optional | - |
| `ProrationScheme` | `*string` | Optional | - |
| `LineItems` | [`[]models.AllocationPreviewLineItem`](../../doc/models/allocation-preview-line-item.md) | Optional | - |
| `AccrueCharge` | `*bool` | Optional | - |
| `Allocations` | [`[]models.AllocationPreviewItem`](../../doc/models/allocation-preview-item.md) | Optional | - |
| `PeriodType` | `*string` | Optional | - |
| `ExistingBalanceInCents` | `*int64` | Optional | An integer representing the amount of the subscription's current balance |

## Example (as JSON)

```json
{
  "start_date": "2016-03-13T12:52:32.123Z",
  "end_date": "2016-03-13T12:52:32.123Z",
  "subtotal_in_cents": 4,
  "total_tax_in_cents": 128,
  "total_discount_in_cents": 122
}
```

