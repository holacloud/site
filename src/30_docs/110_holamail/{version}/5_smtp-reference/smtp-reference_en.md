# SMTP Protocol Reference

Holamail implements the standard SMTP protocol for sending transactional emails. Clients connect on port 2525 and issue SMTP commands.

## Connection

Connect to `smtp.hola.cloud` on port `2525`:

```bash
openssl s_client -connect smtp.hola.cloud:2525 -starttls smtp
```

Or without TLS:

```bash
telnet smtp.hola.cloud 2525
```

## SMTP Commands

### HELO / EHLO

Initiate the SMTP session. `EHLO` enables extended SMTP features.

```http
EHLO client.example.com
```

Response:
```
250-smtp.hola.cloud Hello client.example.com
250-SIZE 35882577
250 PIPELINING
```

### MAIL FROM

Specify the sender address.

```http
MAIL FROM:<noreply@holacloud.app>
```

Response:
```
250 OK
```

### RCPT TO

Specify a recipient. Repeat for multiple recipients.

```http
RCPT TO:<user@example.com>
```

Response:
```
250 OK
```

### DATA

Begin the message content. End with a line containing only a period (`.`).

```http
DATA
```

Response:
```
354 Start mail input
```

Then send headers and body:

```http
From: noreply@holacloud.app
To: user@example.com
Subject: Welcome to HolaCloud

Hello! Thanks for signing up.
.
```

Response:
```
250 OK: queued as abc123
```

### QUIT

Terminate the SMTP session.

```http
QUIT
```

Response:
```
221 Bye
```

## Full SMTP Session Example

```bash
echo -e "EHLO client.example.com\r\nMAIL FROM:<noreply@holacloud.app>\r\nRCPT TO:<user@example.com>\r\nDATA\r\nFrom: noreply@holacloud.app\r\nTo: user@example.com\r\nSubject: Test\r\n\r\nHello!\r\n.\r\nQUIT\r\n" | openssl s_client -connect smtp.hola.cloud:2525 -starttls smtp
```

## Example Go Code

```go
package main

import (
	"log"
	"net/smtp"
)

func main() {
	from := "noreply@holacloud.app"
	to := []string{"user@example.com"}
	msg := []byte("From: noreply@holacloud.app\r\n" +
		"To: user@example.com\r\n" +
		"Subject: Welcome to HolaCloud\r\n\r\n" +
		"Hello! Thanks for signing up.\r\n")

	addr := "smtp.hola.cloud:2525"
	if err := smtp.SendMail(addr, nil, from, to, msg); err != nil {
		log.Fatal(err)
	}
}
```

## Error Codes

| Code | Description |
|------|-------------|
| 211 | System status help reply |
| 214 | Help message |
| 220 | Service ready |
| 221 | Service closing transmission channel |
| 235 | Authentication successful |
| 250 | Requested action completed |
| 251 | User not local; will forward |
| 252 | Cannot verify user; will attempt delivery |
| 354 | Start mail input |
| 421 | Service not available |
| 450 | Mailbox unavailable |
| 451 | Local error processing request |
| 452 | Insufficient storage |
| 500 | Syntax error, command unrecognized |
| 501 | Syntax error in parameters |
| 502 | Command not implemented |
| 503 | Bad sequence of commands |
| 504 | Command parameter not implemented |
| 550 | Mailbox not found |
| 551 | User not local |
| 552 | Exceeded storage allocation |
| 553 | Mailbox name not allowed |
| 554 | Transaction failed |

## Best Practices

- **Use meaningful subjects** to help recipients identify your email and avoid spam filters.
- **Set a valid From address** that matches your domain to improve deliverability.
- **Handle errors gracefully** — check SMTP response codes and retry on transient failures (4xx).
- **Use EHLO** instead of HELO for extended SMTP features.
- **Pipelining** — send multiple commands without waiting for each response to reduce latency.
- **Rate limiting** — avoid sending bulk email too quickly; Holamail may throttle excessive connections.
