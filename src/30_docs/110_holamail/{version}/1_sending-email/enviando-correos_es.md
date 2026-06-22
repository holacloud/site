# Enviando correos

Holamail es un listener SMTP ligero. Acepta los comandos básicos `HELO`/`EHLO`, `MAIL FROM`, `RCPT TO`, `DATA` y `QUIT`, y luego registra el mensaje. No entrega correo a destinatarios externos.

## Detalles de conexión

| Parámetro | Host público de pruebas | Instancia local |
|-----------|--------------------------|-----------------|
| Host | `smtp.testmail.hola.cloud` | `localhost` |
| Puerto | `25` | `2525` |
| Cifrado | Ninguno | Ninguno |
| Autenticación | Ninguna | Ninguna |

Holamail no soporta STARTTLS ni SMTP AUTH.

## Flujo SMTP

```text
HELO client.example.com
MAIL FROM:<noreply@example.com>
RCPT TO:<user@example.com>
DATA
Subject: Hola desde Holamail

Este mensaje será registrado por Holamail.
.
QUIT
```

## Ejemplos

### Con curl

```bash
curl --url smtp://smtp.testmail.hola.cloud:25 \
     --mail-from noreply@example.com \
     --mail-rcpt user@example.com \
     --upload-file email.txt
```

Donde `email.txt` contiene:

```text
Subject: Hola desde Holamail

Este mensaje será registrado por Holamail.
```

### Sesión TCP simple

```bash
telnet localhost 2525
```

Luego ingresa:

```text
HELO client.example.com
MAIL FROM:<noreply@example.com>
RCPT TO:<user@example.com>
DATA
Subject: Prueba local de Holamail

Hola desde un listener local de Holamail.
.
QUIT
```

## Comportamiento esperado

Holamail acepta entrada SMTP sintácticamente válida y registra el mensaje. Es útil para confirmar que una aplicación puede hablar SMTP sin enviar correo real.

No soporta entrega externa, APIs HTTP, STARTTLS, AUTH, rate limiting, plantillas, analíticas ni tracking.
