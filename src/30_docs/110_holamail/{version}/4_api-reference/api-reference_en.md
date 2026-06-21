# Holamail API Reference

Holamail is an **SMTP-based email service** — it does not expose an HTTP API. Clients connect using the standard SMTP protocol on port 2525.

## SMTP Reference

For a complete guide to SMTP commands, connection examples, and best practices, see the [SMTP Protocol Reference](./smtp-reference).

## Connection Details

| Setting | Value |
|---------|-------|
| Host | `smtp.hola.cloud` |
| Port | `2525` |
| Protocol | SMTP (no TLS required) |
| Auth | None (open relay for authenticated projects) |

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| SMTP | `2525` | Raw SMTP connection for sending email |
