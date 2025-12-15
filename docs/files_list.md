# GET /api/v1/files

Method: GET

URL: /api/v1/files

Auth: Bearer (required)

Query params:

- `categoryId` (optional)

Response (200):

```json
[
  { "id":1, "file_name":"abc.pdf", "size":12345, "created_at":"2025-12-01T12:00:00Z" }
]
```
