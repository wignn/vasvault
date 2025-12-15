# GET /api/v1/files/:id

Method: GET

URL: /api/v1/files/:id

Auth: Bearer (required)

Path params:

- `id` (file id)

Response (200):

```json
{ "id":1, "file_name":"abc.pdf", "size":12345, "mime_type":"application/pdf" }
```
