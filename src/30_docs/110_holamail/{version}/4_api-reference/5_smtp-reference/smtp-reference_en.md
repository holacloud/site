# SMTP Reference

Holamail implements a basic plain-SMTP listener. It accepts SMTP messages and logs them; it does not deliver mail externally.

## Connection

Public test host:

```bash
telnet smtp.testmail.hola.cloud 25
```

Local instance:

```bash
telnet localhost 2525
```

STARTTLS and SMTP AUTH are not supported.

## SMTP Commands

### HELO / EHLO

Initiate the SMTP session.

```text
EHLO client.example.com
```

Typical response:

```text
250 Hello client.example.com
```

### MAIL FROM

Specify the sender address.

```text
MAIL FROM:<noreply@example.com>
```

Typical response:

```text
250 OK
```

### RCPT TO

Specify a recipient address.

```text
RCPT TO:<user@example.com>
```

Typical response:

```text
250 OK
```

### DATA

Begin the message content. End with a line containing only a period (`.`).

```text
DATA
```

Typical response:

```text
354 End data with <CR><LF>.<CR><LF>
```

Then send headers and body:

```text
From: noreply@example.com
To: user@example.com
Subject: Holamail test

This message will be logged by Holamail.
.
```

Typical response:

```text
250 OK
```

### QUIT

Terminate the SMTP session.

```text
QUIT
```

Typical response:

```text
221 Bye
```

## Full SMTP Session Example

```bash
printf 'EHLO client.example.com\r\nMAIL FROM:<noreply@example.com>\r\nRCPT TO:<user@example.com>\r\nDATA\r\nFrom: noreply@example.com\r\nTo: user@example.com\r\nSubject: Test\r\n\r\nHello!\r\n.\r\nQUIT\r\n' | nc localhost 2525
```

## Example Go Code

```go
package main

import (
    "log"
    "net/smtp"
)

func main() {
    from := "noreply@example.com"
    to := []string{"user@example.com"}
    msg := []byte("From: noreply@example.com\r\n" +
        "To: user@example.com\r\n" +
        "Subject: Holamail test\r\n\r\n" +
        "This message will be logged by Holamail.\r\n")

    if err := smtp.SendMail("smtp.testmail.hola.cloud:25", nil, from, to, msg); err != nil {
        log.Fatal(err)
    }
}
```

## Error Codes

| Code | Description |
|------|-------------|
| `220` | Service ready |
| `221` | Service closing transmission channel |
| `250` | Requested action completed |
| `354` | Start mail input |
| `500` | Syntax error or unrecognized command |
| `502` | Command not implemented |
| `503` | Bad command sequence |

## Scope

Holamail logs accepted messages only. It does not include external delivery, HTTP APIs, STARTTLS, AUTH, rate limiting, templates, analytics, or tracking.
