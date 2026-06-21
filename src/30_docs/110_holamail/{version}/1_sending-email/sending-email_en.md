# Sending Email with Holamail

Holamail is a lightweight SMTP server for sending transactional emails. It implements raw SMTP protocol (HELO/EHLO, MAIL FROM, RCPT TO, DATA, QUIT) and listens on TCP port 2525. No TLS is required.

## Connection Details

| Parameter | Value |
|-----------|-------|
| Host | `smtp.hola.cloud` |
| Port | `2525` |
| Encryption | None (plain SMTP) |

## SMTP Protocol Flow

The raw SMTP conversation follows this sequence:

```
HELO client.example.com
MAIL FROM:<noreply@holacloud.app>
RCPT TO:<user@example.com>
DATA
Subject: Hello from Holamail

This is the email body.
.
QUIT
```

## Examples

### Using curl

curl supports SMTP natively:

```bash
curl --mail-from noreply@holacloud.app \
     --mail-rcpt user@example.com \
     --upload-file email.txt \
     smtp://smtp.hola.cloud:2525
```

Where `email.txt` contains:

```
Subject: Hello from Holamail

This is the email body.
```

### Using openssl s_client

Manually speak SMTP over a plain TCP connection:

```bash
openssl s_client -connect smtp.hola.cloud:2525 -starttls smtp
```

Or for plain SMTP (no STARTTLS):

```bash
exec 3<>/dev/tcp/smtp.hola.cloud/2525
echo "HELO client.example.com" >&3
echo "MAIL FROM:<noreply@holacloud.app>" >&3
echo "RCPT TO:<user@example.com>" >&3
echo "DATA" >&3
echo "Subject: Hello" >&3
echo "" >&3
echo "Hello from Holamail!" >&3
echo "." >&3
echo "QUIT" >&3
cat <&3
```

### Using Go's net/smtp

```go
package main

import (
    "log"
    "net/smtp"
)

func main() {
    to := []string{"user@example.com"}
    msg := []byte("To: user@example.com\r\n" +
        "Subject: Hello from Holamail\r\n" +
        "\r\n" +
        "This is the email body.\r\n")

    err := smtp.SendMail("smtp.hola.cloud:2525", nil,
        "noreply@holacloud.app", to, msg)
    if err != nil {
        log.Fatal(err)
    }
}
```

### Using Python's smtplib

```python
import smtplib
from email.message import EmailMessage

msg = EmailMessage()
msg.set_content("This is the email body.")
msg["Subject"] = "Hello from Holamail"
msg["From"] = "noreply@holacloud.app"
msg["To"] = "user@example.com"

with smtplib.SMTP("smtp.hola.cloud", 2525) as server:
    server.send_message(msg)
```

## Use Cases

- **Password Resets** — deliver secure reset links on demand.
- **Notifications** — send alerts, reminders, and status updates.
- **System Alerts** — trigger emails from monitoring systems (Nagios, Prometheus, Grafana).
- **Lambda Integration** — invoke Holamail from a Lambda function to send emails programmatically:

```python
import smtplib
from email.message import EmailMessage

def handler(event, context):
    msg = EmailMessage()
    msg.set_content(event.get("body", ""))
    msg["Subject"] = event.get("subject", "Notification")
    msg["From"] = "noreply@holacloud.app"
    msg["To"] = event["to"]

    with smtplib.SMTP("smtp.hola.cloud", 2525) as server:
        server.send_message(msg)

    return {"statusCode": 200, "body": "Email sent"}
```

## Error Handling

Holamail returns standard SMTP status codes:

| Code | Meaning |
|------|---------|
| 250 | Request completed |
| 550 | Mailbox unavailable (invalid recipient) |
| 551 | Recipient rejected |
| 552 | Exceeded storage allocation |
| 553 | Address invalid |
| 554 | Transaction failed |

When a recipient address is invalid, the server returns `550`. Handle connection errors (timeout, refused) in your application code — they typically indicate network issues or the service being temporarily unavailable.
