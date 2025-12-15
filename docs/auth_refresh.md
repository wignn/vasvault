# POST /api/v1/refresh

Method: POST

URL: /api/v1/refresh

Auth: No (uses refresh token)

Request JSON:

```json
{ "refresh_token": "<refresh>" }
```

Response (200):

```json
{ "access_token": "<jwt>", "refresh_token": "<refresh>" }
```
