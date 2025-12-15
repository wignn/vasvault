# POST /api/v1/login

Method: POST

URL: /api/v1/login

Auth: No

Request JSON:

```json
{
  "email": "user@example.com",
  "password": "secret"
}
```

Response (200):

```json
{
  "user": {"id":1,"username":"alice","email":"user@example.com"},
  "token": {"access_token":"<jwt>","refresh_token":"<refresh>"}
}
```
