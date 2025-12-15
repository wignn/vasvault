# POST /api/v1/register

Method: POST

URL: /api/v1/register

Auth: No

Request JSON:

```json
{
  "username": "alice",
  "email": "user@example.com",
  "password": "secret"
}
```

Response (201):

```json
{ "id": 1, "username": "alice", "email": "user@example.com" }
```
