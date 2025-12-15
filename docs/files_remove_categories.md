# POST /api/v1/files/:id/categories/remove

Method: POST

URL: /api/v1/files/:id/categories/remove

Auth: Bearer (required)

Path params:

- `id` (file id)

Request JSON:

```json
{ "category_ids": [1,2] }
```

Response (200):

```json
{ "message": "categories removed successfully" }
```
