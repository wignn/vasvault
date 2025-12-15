# POST /api/v1/workspaces

Method: POST

URL: /api/v1/workspaces

Auth: Bearer (required)

Request JSON:

```json
{ "name": "Team Vault", "description": "Shared storage" }
```

Response (201):

```json
{ "id":1, "name":"Team Vault", "description":"Shared storage", "owner_id":1 }
```
