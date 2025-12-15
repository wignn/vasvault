# PUT /api/v1/workspaces/:id

Method: PUT

URL: /api/v1/workspaces/:id

Auth: Bearer (required)

Request JSON:

```json
{ "name": "New Name", "description": "Updated" }
```

Response (200):

```json
{ "id":1, "name":"New Name", "description":"Updated" }
```
