# PUT /api/v1/categories/:id

Method: PUT

URL: /api/v1/categories/:id

Auth: Bearer (required)

Path params:

- `id` (category id)

Request JSON (any of fields):

```json
{ "name": "Receipts", "color": "#00FF00" }
```

Response (200):

```json
{ "id":1, "name":"Receipts", "color":"#00FF00" }
```
