# Holamail API Reference

Holamail does not expose an HTTP API. It is a plain SMTP listener that accepts a basic SMTP conversation and logs the message.

## SMTP Reference

For command examples, see the [SMTP Protocol Reference](./smtp-reference).

## Connection Details

| Setting | Public test host | Local instance |
|---------|------------------|----------------|
| Host | `smtp.testmail.hola.cloud` | `localhost` |
| Port | `25` | `2525` |
| Protocol | Plain SMTP | Plain SMTP |
| Auth | None | None |

## Supported Commands

| Command | Description |
|---------|-------------|
| `HELO` / `EHLO` | Start the SMTP session |
| `MAIL FROM` | Set the sender envelope address |
| `RCPT TO` | Set the recipient envelope address |
| `DATA` | Send message headers and body, ending with `.` |
| `QUIT` | Close the session |

## Not Supported

Holamail does not provide external delivery, an HTTP API, STARTTLS, AUTH, rate limiting, templates, analytics, or tracking.
