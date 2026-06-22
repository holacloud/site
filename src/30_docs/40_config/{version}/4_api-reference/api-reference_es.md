# Referencia API de Config

URL base: `https://api.hola.cloud`

## Autenticación

`/v1/config` requiere `X-Glue-Authentication`. `/v0/configs` gestiona Things de configuración directamente.

## Endpoints

| Método | Ruta | Descripción | Auth |
|--------|------|-------------|------|
| GET | `/v0/configs` | Listar Things de configuración | Pública |
| POST | `/v0/configs` | Crear un Thing de configuración con `entries` | Pública |
| GET | `/v0/configs/{id}` | Obtener un Thing de configuración por ID | Pública |
| DELETE | `/v0/configs/{id}` | Eliminar un Thing de configuración | Pública |
| PATCH | `/v0/configs/{id}` | Parchear los `entries` de un Thing de configuración | Pública |
| GET | `/v1/config` | Obtener el mapa `entries` de configuración del usuario autenticado | Auth Glue |
| PATCH | `/v1/config` | Actualizar el mapa `entries` de configuración del usuario autenticado | Auth Glue |

## Formas

Thing de configuración v0:

```json
{
  "id": "cfg_abc123",
  "entries": {
    "feature.newCheckout": true
  }
}
```

Configuración de usuario v1:

```json
{
  "entries": {
    "feature.newCheckout": true
  }
}
```
