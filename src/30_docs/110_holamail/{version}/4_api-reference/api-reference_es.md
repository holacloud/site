# Referencia de API de Holamail

Holamail no expone una API HTTP. Es un listener SMTP simple que acepta una conversación SMTP básica y registra el mensaje.

## Referencia SMTP

Para ejemplos de comandos, consulta la [Referencia del protocolo SMTP](./smtp-reference).

## Detalles de conexión

| Opción | Host público de pruebas | Instancia local |
|--------|--------------------------|-----------------|
| Host | `smtp.testmail.hola.cloud` | `localhost` |
| Puerto | `25` | `2525` |
| Protocolo | SMTP sin cifrado | SMTP sin cifrado |
| Auth | Ninguna | Ninguna |

## Comandos soportados

| Comando | Descripción |
|---------|-------------|
| `HELO` / `EHLO` | Inicia la sesión SMTP |
| `MAIL FROM` | Define el remitente del sobre |
| `RCPT TO` | Define el destinatario del sobre |
| `DATA` | Envía cabeceras y cuerpo, terminando con `.` |
| `QUIT` | Cierra la sesión |

## No soportado

Holamail no ofrece entrega externa, API HTTP, STARTTLS, AUTH, rate limiting, plantillas, analíticas ni tracking.
