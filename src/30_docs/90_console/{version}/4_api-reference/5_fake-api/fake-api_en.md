# Fake API

Fixed mock endpoints that serve Console development and demo data.

## Description

The Console service exposes only the fake API routes listed below. These endpoints allow the UI to run without the real backend services. Paths not listed here are not part of the Console fake API.

## Sub-routes

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/fakeapi/authapi/auth/me` | Returns the current mock user |
| `GET` | `/fakeapi/inceptionapi/v1/databases` | Lists mock databases |
| `GET` | `/fakeapi/inceptionapi/v1/databases/{databaseid}/collections` | Lists mock collections for a database |
| `GET` | `/fakeapi/lambdasapi/api/v0/lambdas` | Lists mock lambdas |
| `GET` | `/fakeapi/projectsapi/v0/projects` | Lists mock projects |
| `GET` | `/fakeapi/projectsapi/v0/projects/{projectid}` | Returns one mock project |
| `GET` | `/fakeapi/filesapi/v1/buckets` | Lists mock buckets |

## Authentication

None. These mock endpoints are unauthenticated.

## Example

```bash
curl -X GET "https://api.hola.cloud/fakeapi/authapi/auth/me"
```

```json
{
  "email": "fulanez@gmail.com",
  "id": "user-1234",
  "nick": "fulanez",
  "picture": "https://lh3.googleusercontent.com/a-/ACNPEu8h8e8wfcL1Hsk4JbFsROFXIBNIllOFLLPnWEDWV2s=s96-c"
}
```

```bash
curl -X GET "https://api.hola.cloud/fakeapi/projectsapi/v0/projects"
```

```json
[
  {
    "id": "project-00000000-0000-0000-0000-000000000001",
    "name": "Hello",
    "host": "red-pepper.holacloud.app",
    "owners": ["user1"]
  }
]
```

```bash
curl -X GET "https://api.hola.cloud/fakeapi/lambdasapi/api/v0/lambdas"
```

```json
[
  {
    "id": "a663e1f1-e5cb-4fc2-b846-53f2cc7574c9",
    "method": "GET",
    "name": "Index",
    "path": "/"
  }
]
```

## Error Codes

| Code | Description |
|------|-------------|
| 200 | Mock data returned successfully |
| 404 | Path under `/fakeapi/` not recognized |
