
# Update Segment Request

## Structure

`UpdateSegmentRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Segment` | [`models.UpdateSegment`](../../doc/models/update-segment.md) | Required | - |

## Example (as JSON)

```json
{
  "segment": {
    "pricing_scheme": "stairstep",
    "prices": [
      {
        "starting_quantity": 64,
        "ending_quantity": 38,
        "unit_price": {
          "key1": "val1",
          "key2": "val2"
        }
      },
      {
        "starting_quantity": 64,
        "ending_quantity": 38,
        "unit_price": {
          "key1": "val1",
          "key2": "val2"
        }
      },
      {
        "starting_quantity": 64,
        "ending_quantity": 38,
        "unit_price": {
          "key1": "val1",
          "key2": "val2"
        }
      }
    ]
  }
}
```

