# Enviando Correos

Holamail es un servidor SMTP ligero para enviar correos electrónicos transaccionales. Implementa el protocolo SMTP estándar (HELO/EHLO, MAIL FROM, RCPT TO, DATA, QUIT) y escucha en el puerto TCP 2525. No requiere TLS.

## Detalles de Conexión

| Parámetro | Valor |
|-----------|-------|
| Host | `smtp.hola.cloud` |
| Puerto | `2525` |
| Cifrado | Ninguno (SMTP simple) |

## Flujo del Protocolo SMTP

La conversación SMTP cruda sigue esta secuencia:

```
HELO cliente.ejemplo.com
MAIL FROM:<noreply@holacloud.app>
RCPT TO:<usuario@ejemplo.com>
DATA
Subject: Hola desde Holamail

Este es el cuerpo del correo.
.
QUIT
```

## Ejemplos

### Usando curl

curl soporta SMTP de forma nativa:

```bash
curl --mail-from noreply@holacloud.app \
     --mail-rcpt usuario@ejemplo.com \
     --upload-file correo.txt \
     smtp://smtp.hola.cloud:2525
```

Donde `correo.txt` contiene:

```
Subject: Hola desde Holamail

Este es el cuerpo del correo.
```

### Usando openssl s_client

Habla SMTP manualmente sobre una conexión TCP simple:

```bash
exec 3<>/dev/tcp/smtp.hola.cloud/2525
echo "HELO cliente.ejemplo.com" >&3
echo "MAIL FROM:<noreply@holacloud.app>" >&3
echo "RCPT TO:<usuario@ejemplo.com>" >&3
echo "DATA" >&3
echo "Subject: Hola" >&3
echo "" >&3
echo "¡Hola desde Holamail!" >&3
echo "." >&3
echo "QUIT" >&3
cat <&3
```

### Usando net/smtp de Go

```go
package main

import (
    "log"
    "net/smtp"
)

func main() {
    para := []string{"usuario@ejemplo.com"}
    msg := []byte("To: usuario@ejemplo.com\r\n" +
        "Subject: Hola desde Holamail\r\n" +
        "\r\n" +
        "Este es el cuerpo del correo.\r\n")

    err := smtp.SendMail("smtp.hola.cloud:2525", nil,
        "noreply@holacloud.app", para, msg)
    if err != nil {
        log.Fatal(err)
    }
}
```

### Usando smtplib de Python

```python
import smtplib
from email.message import EmailMessage

msg = EmailMessage()
msg.set_content("Este es el cuerpo del correo.")
msg["Subject"] = "Hola desde Holamail"
msg["From"] = "noreply@holacloud.app"
msg["To"] = "usuario@ejemplo.com"

with smtplib.SMTP("smtp.hola.cloud", 2525) as servidor:
    servidor.send_message(msg)
```

## Casos de Uso

- **Restablecimiento de contraseñas** — entrega enlaces seguros de restablecimiento bajo demanda.
- **Notificaciones** — envía alertas, recordatorios y actualizaciones de estado.
- **Alertas del sistema** — activa correos desde sistemas de monitoreo (Nagios, Prometheus, Grafana).
- **Integración con Lambda** — invoca Holamail desde una función Lambda para enviar correos programáticamente:

```python
import smtplib
from email.message import EmailMessage

def handler(event, context):
    msg = EmailMessage()
    msg.set_content(event.get("body", ""))
    msg["Subject"] = event.get("subject", "Notificación")
    msg["From"] = "noreply@holacloud.app"
    msg["To"] = event["to"]

    with smtplib.SMTP("smtp.hola.cloud", 2525) as servidor:
        servidor.send_message(msg)

    return {"statusCode": 200, "body": "Correo enviado"}
```

## Manejo de Errores

Holamail devuelve códigos de estado SMTP estándar:

| Código | Significado |
|--------|-------------|
| 250 | Solicitud completada |
| 550 | Buzón no disponible (destinatario inválido) |
| 551 | Destinatario rechazado |
| 552 | Almacenamiento excedido |
| 553 | Dirección inválida |
| 554 | Transacción fallida |

Cuando una dirección de destinatario no es válida, el servidor responde con `550`. Maneja los errores de conexión (timeout, rechazo) en el código de tu aplicación — generalmente indican problemas de red o que el servicio está temporalmente no disponible.
