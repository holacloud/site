# Holamail

Holamail es un servicio SMTP sencillo para enviar correos electrónicos transaccionales. Escucha conexiones SMTP en el puerto 2525 y entrega los mensajes a sus destinatarios.

## Cómo Funciona

Holamail acepta conexiones mediante el **protocolo SMTP** en el puerto 2525. Los clientes se conectan usando SMTP estándar — cualquier lenguaje o framework con una librería cliente SMTP puede enviar correos a través de Holamail sin SDKs adicionales.

```
┌──────────────┐   SMTP :2525    ┌───────────┐
│  Aplicación  │───────────────▶│  Holamail  │────▶  Destinatario
└──────────────┘                 └───────────┘
```

## Casos de Uso

- **Correos Transaccionales**: Confirmaciones de pedido, verificación de cuenta, restablecimiento de contraseña.
- **Notificaciones**: Alertas, recordatorios, actualizaciones de estado generadas por eventos de la aplicación.
- **Restablecimiento de Contraseña**: Entrega segura de enlaces de restablecimiento a los usuarios.
- **Alertas del Sistema**: Envío automatizado de alertas desde monitoreo o pipelines CI/CD.

## Integración con Otros Servicios

Holamail funciona junto a otros servicios de HolaCloud:

- **Lambda** — envía correos desde funciones serverless.
- **Consola** — administra plantillas de correo y visualiza registros de entrega.
- **InceptionDB** — almacena plantillas de correo e historial de envíos.

## Primeros Pasos

Conéctate a Holamail en el puerto 2525 usando cualquier cliente SMTP:

```bash
# Usando swaks (Swiss Army Knife para SMTP)
swaks --to usuario@ejemplo.com \
      --from noreply@holacloud.app \
      --server smtp.hola.cloud \
      --port 2525 \
      --header "Subject: Bienvenido a HolaCloud" \
      --body "¡Hola! Gracias por registrarte."
```

O usando `smtplib` de Python:

```python
import smtplib
from email.message import EmailMessage

msg = EmailMessage()
msg.set_content("¡Hola! Gracias por registrarte.")
msg["Subject"] = "Bienvenido a HolaCloud"
msg["From"] = "noreply@holacloud.app"
msg["To"] = "usuario@ejemplo.com"

with smtplib.SMTP("smtp.hola.cloud", 2525) as s:
    s.send_message(msg)
```
