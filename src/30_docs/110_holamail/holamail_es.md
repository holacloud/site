# Holamail

Holamail es un listener SMTP pequeño para flujos de desarrollo y prueba. Acepta una conversación SMTP básica, registra el mensaje recibido en logs y no entrega correo a destinatarios externos.

## Qué implementa

Holamail acepta conexiones SMTP sin cifrado y soporta estos comandos:

- `HELO` / `EHLO`
- `MAIL FROM`
- `RCPT TO`
- `DATA`
- `QUIT`

No ofrece API HTTP, entrega externa, STARTTLS, AUTH, rate limiting, plantillas, analíticas ni tracking.

## Cómo funciona

Las aplicaciones se conectan con un cliente SMTP estándar. Holamail registra el sobre y el contenido del mensaje para inspección.

```
┌────────────┐    SMTP    ┌───────────┐
│ Aplicación │───────────▶│ Holamail  │──▶ registra mensaje
└────────────┘            └───────────┘
```

## Primeros pasos

Usa el host público de pruebas o una instancia local:

```bash
swaks --to user@example.com \
      --from noreply@example.com \
      --server smtp.testmail.hola.cloud \
      --port 25 \
      --header "Subject: Prueba de Holamail" \
      --body "Hola desde Holamail."
```

Para un listener local:

```bash
swaks --to user@example.com \
      --from noreply@example.com \
      --server localhost \
      --port 2525 \
      --header "Subject: Prueba local de Holamail" \
      --body "Hola desde un listener local de Holamail."
```
