# Secreto

Endpoint protegido de prueba que valida la autenticación.

## Descripción

Este endpoint requiere un token de autenticación válido. Se utiliza para verificar que la capa de autenticación funciona correctamente.

## Autenticación

Requiere autenticación `glueauth.Require` mediante una clave API o sesión válida.

## Solicitud

No se requiere cuerpo en la solicitud.

## Ejemplo

```bash
curl -X GET "https://api.hola.cloud/v0/secret" \
  -H "Authorization: Bearer <api-key>"
```

## Respuesta

```text
OK
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 200 | Autenticación exitosa |
| 401 | Autenticación faltante o inválida |
