# Holamail

Holamail is a simple SMTP email service for sending transactional emails. It listens for raw SMTP connections on port 2525 and delivers messages to their recipients.

## How It Works

Holamail accepts connections via the **SMTP protocol** on port 2525. Clients connect using standard SMTP — any language or framework with an SMTP client library can send email through Holamail without additional SDKs.

```
┌──────────────┐    SMTP :2525    ┌───────────┐
│  Application │────────────────▶│  Holamail  │────▶  Recipient
└──────────────┘                  └───────────┘
```

## Use Cases

- **Transactional Emails**: Order confirmations, account verification, password resets.
- **Notifications**: Alerts, reminders, status updates triggered by application events.
- **Password Resets**: Securely deliver password reset links to users.
- **System Alerts**: Send automated alerts from monitoring or CI/CD pipelines.

## Integration with Other Services

Holamail works alongside other HolaCloud services:

- **Lambda** — trigger email sends from serverless functions.
- **Console** — manage email templates and view delivery logs.
- **InceptionDB** — store email templates and delivery history.

## Getting Started

Connect to Holamail on port 2525 using any SMTP client:

```bash
# Using swaks (Swiss Army Knife for SMTP)
swaks --to user@example.com \
      --from noreply@holacloud.app \
      --server smtp.hola.cloud \
      --port 2525 \
      --header "Subject: Welcome to HolaCloud" \
      --body "Hello! Thanks for signing up."
```

Or using Python's `smtplib`:

```python
import smtplib
from email.message import EmailMessage

msg = EmailMessage()
msg.set_content("Hello! Thanks for signing up.")
msg["Subject"] = "Welcome to HolaCloud"
msg["From"] = "noreply@holacloud.app"
msg["To"] = "user@example.com"

with smtplib.SMTP("smtp.hola.cloud", 2525) as s:
    s.send_message(msg)
```
