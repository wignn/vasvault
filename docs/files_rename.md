
# PUT /api/v1/files/:id

Method: PUT (application/json)

URL: /api/v1/files/:id

Auth: Bearer (required)

Body (application/json):

- `new_name` (string, required) — desired new filename. If the provided name has no extension, the original file extension is preserved and appended automatically.

Example request:

```json
{ "new_name": "report-2025" }
```

Behavior:

- Only the file owner may rename the file. The server verifies ownership and returns an error if the authenticated user does not own the file.
- The server renames the file on disk (moves it under the uploads base path) and updates the file metadata in the database (`filename`, `filepath`).
- If a file already exists with the target name in the uploads folder, the request fails with a `400` error to avoid overwriting.
- The server preserves the original file extension if the `new_name` does not include an extension.

Successful response (200):

```json
{
  "id": 15,
  "user_id": 11,
  "workspace_id": 3,
  "file_name": "report-2025.pdf",
  "file_path": "./uploads/report-2025.pdf",
  "mime_type": "application/pdf",
  "size": 12345,
  "categories": [],
  "created_at": "2025-12-17T12:34:56Z"
}
```

Errors:

- `400 Bad Request` — invalid `file id`, invalid JSON, target filename already exists, or unauthorized (file not owned by user).
- `404 Not Found` — file not found.
- `500 Internal Server Error` — filesystem or database errors while renaming.

Notes for integrators:

- Callers should provide only the desired name (without path). The server stores files inside the configured uploads directory.
- If you need to preserve the original filename extension, you may omit the extension in `new_name` — the server will append the original extension automatically.

See implementation: [internal/handlers/file_handler.go](internal/handlers/file_handler.go#L1) and [internal/services/file_service.go](internal/services/file_service.go#L1).
