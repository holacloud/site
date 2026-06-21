# Referencia del Protocolo SMTP

Holamail implementa el protocolo SMTP estándar para el envío de correos electrónicos transaccionales. Los clientes se conectan al puerto 2525 y emiten comandos SMTP.

## Conexión

Conéctese a `smtp.hola.cloud` en el puerto `2525`:

```bash
openssl s_client -connect smtp.hola.cloud:2525 -starttls smtp
```

O sin TLS:

```bash
telnet smtp.hola.cloud 2525
```

## Comandos SMTP

### HELO / EHLO

Inicia la sesión SMTP. `EHLO` habilita características SMTP extendidas.

```http
EHLO client.example.com
```

Respuesta:
```
250-smtp.hola.cloud Hello client.example.com
250-SIZE 35882577
250 PIPELINING
```

### MAIL FROM

Especifica la dirección del remitente.

```http
MAIL FROM:<noreply@holacloud.app>
```

Respuesta:
```
250 OK
```

### RCPT TO

Especifica un destinatario. Repita para múltiples destinatarios.

```http
RCPT TO:<user@example.com>
```

Respuesta:
```
250 OK
```

### DATA

Inicia el contenido del mensaje. Termine con una línea que contenga solo un punto (`.`).

```http
DATA
```

Respuesta:
```
354 Start mail input
```

Luego envíe los encabezados y el cuerpo:

```http
From: noreply@holacloud.app
To: user@example.com
Subject: Welcome to HolaCloud

Hello! Thanks for signing up.
.
```

Respuesta:
```
250 OK: queued as abc123
```

### QUIT

Termina la sesión SMTP.

```http
QUIT
```

Respuesta:
```
221 Bye
```

## Ejemplo Completo de Sesión SMTP

```bash
echo -e "EHLO client.example.com\r\nMAIL FROM:<noreply@holacloud.app>\r\nRCPT TO:<user@example.com>\r\nDATA\r\nFrom: noreply@holacloud.app\r\nTo: user@example.com\r\nSubject: Test\r\n\r\nHello!\r\n.\r\nQUIT\r\n" | openssl s_client -connect smtp.hola.cloud:2525 -starttls smtp
```

## Ejemplo en Go

```go
package main

import (
	"log"
	"net/smtp"
)

func main() {
	from := "noreply@holacloud.app"
	to := []string{"user@example.com"}
	msg := []byte("From: noreply@holacloud.app\r\n" +
		"To: user@example.com\r\n" +
		"Subject: Welcome to HolaCloud\r\n\r\n" +
		"Hello! Thanks for signing up.\r\n")

	addr := "smtp.hola.cloud:2525"
	if err := smtp.SendMail(addr, nil, from, to, msg); err != nil {
		log.Fatal(err)
	}
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 211 | Respuesta de ayuda de estado del sistema |
| 214 | Mensaje de ayuda |
| 220 | Servicio listo |
| 221 | Servicio cerrando canal de transmisión |
| 235 | Autenticación exitosa |
| 250 | Acción solicitada completada |
| 251 | Usuario no local; se reenviará |
| 252 | No se puede verificar usuario; se intentará entrega |
| 354 | Iniciar entrada de correo |
| 421 | Servicio no disponible |
| 450 | Buzón no disponible |
| 451 | Error local procesando solicitud |
| 452 | Almacenamiento insuficiente |
| 500 | Error de sintaxis, comando no reconocido |
| 501 | Error de sintaxis en parámetros |
| 502 | Comando no implementado |
| 503 | Secuencia incorrecta de comandos |
| 504 | Parámetro de comando no implementado |
| 550 | Buzón no encontrado |
| 551 | Usuario no local |
| 552 | Excedida la asignación de almacenamiento |
| 553 | Nombre de buzón no permitido |
| 554 | Transacción fallida |

## Mejores Prácticas

- **Use asuntos significativos** para ayudar a los destinatarios a identificar su correo y evitar filtros de spam.
- **Establezca una dirección From válida** que coincida con su dominio para mejorar la capacidad de entrega.
- **Maneje los errores correctamente** — verifique los códigos de respuesta SMTP y reintente en fallos transitorios (4xx).
- **Use EHLO** en lugar de HELO para características SMTP extendidas.
- **Pipelining** — envíe múltiples comandos sin esperar cada respuesta para reducir latencia.
- **Límite de tasa** — evite enviar correos masivos demasiado rápido; Holamail puede limitar conexiones excesivas.
