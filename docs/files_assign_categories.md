# POST /api/v1/files/:id/categories/assign

Method: POST

URL: /api/v1/files/:id/categories/assign

Auth: Bearer (required)

Path params:

- `id` (file id)

Request JSON:

```json
{ "category_ids": [1,2] }
```

Response (200):

```json
{ "message": "categories assigned successfully" }
```
