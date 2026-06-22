# Sending Email

Holamail is a lightweight SMTP listener. It accepts the basic SMTP commands `HELO`/`EHLO`, `MAIL FROM`, `RCPT TO`, `DATA`, and `QUIT`, then logs the message. It does not deliver email to external recipients.

## Connection Details

| Parameter | Public test host | Local instance |
|-----------|------------------|----------------|
| Host | `smtp.testmail.hola.cloud` | `localhost` |
| Port | `25` | `2525` |
| Encryption | None | None |
| Authentication | None | None |

Holamail does not support STARTTLS or SMTP AUTH.

## SMTP Protocol Flow

The raw SMTP conversation follows this sequence:

```text
HELO client.example.com
MAIL FROM:<noreply@example.com>
RCPT TO:<user@example.com>
DATA
Subject: Hello from Holamail

This message will be logged by Holamail.
.
QUIT
```

## Examples

### Using curl

```bash
curl --url smtp://smtp.testmail.hola.cloud:25 \
     --mail-from noreply@example.com \
     --mail-rcpt user@example.com \
     --upload-file email.txt
```

Where `email.txt` contains:

```text
Subject: Hello from Holamail

This message will be logged by Holamail.
```

### Plain TCP Session

```bash
telnet localhost 2525
```

Then enter:

```text
HELO client.example.com
MAIL FROM:<noreply@example.com>
RCPT TO:<user@example.com>
DATA
Subject: Local Holamail test

Hello from a local Holamail listener.
.
QUIT
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
        "This message will be logged by Holamail.\r\n")

    err := smtp.SendMail("smtp.testmail.hola.cloud:25", nil,
        "noreply@example.com", to, msg)
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
msg.set_content("This message will be logged by Holamail.")
msg["Subject"] = "Hello from Holamail"
msg["From"] = "noreply@example.com"
msg["To"] = "user@example.com"

with smtplib.SMTP("smtp.testmail.hola.cloud", 25) as server:
    server.send_message(msg)
```

## Expected Behavior

Holamail accepts syntactically valid SMTP input and logs the message. It is useful for confirming that an application can speak SMTP without sending real email.

Unsupported features include external delivery, HTTP APIs, STARTTLS, AUTH, rate limiting, templates, analytics, and tracking.
