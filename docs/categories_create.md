# POST /api/v1/categories

Method: POST

URL: /api/v1/categories

Auth: Bearer (required)

Request JSON:

```json
{ "name": "Invoices", "color": "#FF0000" }
```

Response (201):

```json
{ "id":1, "name":"Invoices", "color":"#FF0000", "user_id":1 }
```
