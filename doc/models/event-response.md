
# Event Response

## Structure

`EventResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Event` | [`models.Event`](../../doc/models/event.md) | Required | - |

## Example (as JSON)

```json
{
  "event": {
    "id": 242,
    "key": "key0",
    "message": "message0",
    "subscription_id": 96,
    "customer_id": 24,
    "created_at": "2016-03-13T12:52:32.123Z",
    "event_specific_data": {
      "key1": "val1",
      "key2": "val2"
    }
  }
}
```

