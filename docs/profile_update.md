# PUT /api/v1/profile

Method: PUT

URL: /api/v1/profile

Auth: Bearer (required)

Request JSON (all fields optional):

```json
{ "username": "newname", "email": "new@example.com", "password": "newpass" }
```

Response (200):

```json
{ "id":1, "username":"newname", "email":"new@example.com" }
```
