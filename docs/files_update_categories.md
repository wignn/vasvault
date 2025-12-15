# PUT /api/v1/files/:id/categories

Method: PUT

URL: /api/v1/files/:id/categories

Auth: Bearer (required)

Path params:

- `id` (file id)

Request JSON:

```json
{ "category_ids": [3,4] }
```

Response (200):

```json
{ "message": "categories updated successfully" }
```
