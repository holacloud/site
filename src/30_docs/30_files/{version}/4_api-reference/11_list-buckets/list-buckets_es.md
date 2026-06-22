# Listar Buckets

Listar todos los buckets de la cuenta autenticada.

## Autenticación

Requiere los encabezados `Api-Key` y `Api-Secret`.

## Parámetros de Consulta

| Parámetro | Tipo | Valor por Defecto | Descripción |
|-----------|------|-------------------|-------------|
| `limit` | integer | 100 | Número máximo de buckets a devolver |
| `offset` | integer | 0 | Número de buckets a omitir |

## Solicitud

```bash
curl "https://api.hola.cloud/v1/buckets?limit=10&offset=0" \
  -H "Api-Key: SU_API_KEY" \
  -H "Api-Secret: SU_API_SECRET"
```

## Respuesta

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "buckets": [
    {
      "id": "bkt_abc123",
      "name": "mi-primer-bucket",
      "createdAt": "2026-06-21T10:00:00Z",
      "size": 1048576,
      "fileCount": 5,
      "public": false
    },
    {
      "id": "bkt_def456",
      "name": "assets",
      "createdAt": "2026-06-20T08:30:00Z",
      "size": 52428800,
      "fileCount": 42,
      "public": true
    }
  ],
  "total": 2
}
```

## Códigos de Error

| Estado | Código | Descripción |
|--------|--------|-------------|
| 401 | Unauthorized | Credenciales API faltantes o inválidas |
| 403 | Forbidden | Permisos insuficientes |
| 500 | Internal Server Error | Ocurrió un error inesperado |
