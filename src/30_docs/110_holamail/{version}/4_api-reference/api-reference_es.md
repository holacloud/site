# Referencia de la API de Holamail

Holamail es un servicio de correo electrónico basado en **SMTP** — no expone una API HTTP. Los clientes se conectan usando el protocolo SMTP estándar en el puerto 2525.

## Referencia SMTP

Para una guía completa de comandos SMTP, ejemplos de conexión y mejores prácticas, consulte la [Referencia del Protocolo SMTP](./smtp-reference).

## Detalles de Conexión

| Configuración | Valor |
|---------------|-------|
| Host | `smtp.hola.cloud` |
| Puerto | `2525` |
| Protocolo | SMTP (sin TLS requerido) |
| Autenticación | Ninguna (relé abierto para proyectos autenticados) |

## Endpoints

| Método | Ruta | Descripción |
|--------|------|-------------|
| SMTP | `2525` | Conexión SMTP para enviar correo electrónico |
