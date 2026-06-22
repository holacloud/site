# Holamail

Holamail is a small SMTP listener for development and test workflows. It accepts a basic SMTP conversation, logs the received message, and does not deliver mail to external recipients.

## What It Implements

Holamail accepts raw SMTP connections and supports these commands:

- `HELO` / `EHLO`
- `MAIL FROM`
- `RCPT TO`
- `DATA`
- `QUIT`

It does not provide an HTTP API, external delivery, STARTTLS, AUTH, rate limiting, templates, analytics, or tracking.

## How It Works

Applications connect with a plain SMTP client. Holamail records the envelope and message data in its logs for inspection.

```
┌──────────────┐    SMTP    ┌───────────┐
│  Application │───────────▶│ Holamail  │──▶ logs message
└──────────────┘            └───────────┘
```

## Getting Started

Use the public test host or a local instance:

```bash
swaks --to user@example.com \
      --from noreply@example.com \
      --server smtp.testmail.hola.cloud \
      --port 25 \
      --header "Subject: Holamail test" \
      --body "Hello from Holamail."
```

For a local listener:

```bash
swaks --to user@example.com \
      --from noreply@example.com \
      --server localhost \
      --port 2525 \
      --header "Subject: Local Holamail test" \
      --body "Hello from a local Holamail listener."
```
