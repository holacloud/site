# Referencia SMTP

Holamail implementa un listener SMTP básico sin cifrado. Acepta mensajes SMTP y los registra; no entrega correo externamente.

## Conexión

Host público de pruebas:

```bash
telnet smtp.testmail.hola.cloud 25
```

Instancia local:

```bash
telnet localhost 2525
```

STARTTLS y SMTP AUTH no están soportados.

## Comandos SMTP

### HELO / EHLO

Inicia la sesión SMTP.

```text
EHLO client.example.com
```

### MAIL FROM

Especifica la dirección del remitente.

```text
MAIL FROM:<noreply@example.com>
```

### RCPT TO

Especifica la dirección del destinatario.

```text
RCPT TO:<user@example.com>
```

### DATA

Inicia el contenido del mensaje. Termina con una línea que contenga solo un punto (`.`).

```text
DATA
From: noreply@example.com
To: user@example.com
Subject: Prueba de Holamail

Este mensaje será registrado por Holamail.
.
```

### QUIT

Cierra la sesión SMTP.

```text
QUIT
```

## Ejemplo de sesión completa

```bash
printf 'EHLO client.example.com\r\nMAIL FROM:<noreply@example.com>\r\nRCPT TO:<user@example.com>\r\nDATA\r\nFrom: noreply@example.com\r\nTo: user@example.com\r\nSubject: Test\r\n\r\nHola!\r\n.\r\nQUIT\r\n' | nc localhost 2525
```

## Códigos habituales

| Código | Descripción |
|--------|-------------|
| `220` | Servicio listo |
| `221` | Cierre de canal |
| `250` | Acción completada |
| `354` | Inicio de entrada de datos |
| `500` | Error de sintaxis o comando no reconocido |
| `502` | Comando no implementado |
| `503` | Secuencia de comandos incorrecta |

## Alcance

Holamail solo registra mensajes aceptados. No incluye entrega externa, APIs HTTP, STARTTLS, AUTH, rate limiting, plantillas, analíticas ni tracking.
