# SMTP 协议参考

Holamail 实现了用于发送事务性电子邮件的标准 SMTP 协议。客户端通过端口 2525 连接并发出 SMTP 命令。

## 连接

连接到 `smtp.hola.cloud` 的 `2525` 端口：

```bash
openssl s_client -connect smtp.hola.cloud:2525 -starttls smtp
```

或无需 TLS：

```bash
telnet smtp.hola.cloud 2525
```

## SMTP 命令

### HELO / EHLO

启动 SMTP 会话。`EHLO` 启用扩展 SMTP 功能。

```http
EHLO client.example.com
```

响应：
```
250-smtp.hola.cloud Hello client.example.com
250-SIZE 35882577
250 PIPELINING
```

### MAIL FROM

指定发件人地址。

```http
MAIL FROM:<noreply@holacloud.app>
```

响应：
```
250 OK
```

### RCPT TO

指定收件人。多个收件人可重复此命令。

```http
RCPT TO:<user@example.com>
```

响应：
```
250 OK
```

### DATA

开始消息内容。以仅包含一个点号（`.`）的行结束。

```http
DATA
```

响应：
```
354 Start mail input
```

然后发送标头和正文：

```http
From: noreply@holacloud.app
To: user@example.com
Subject: Welcome to HolaCloud

Hello! Thanks for signing up.
.
```

响应：
```
250 OK: queued as abc123
```

### QUIT

终止 SMTP 会话。

```http
QUIT
```

响应：
```
221 Bye
```

## 完整 SMTP 会话示例

```bash
echo -e "EHLO client.example.com\r\nMAIL FROM:<noreply@holacloud.app>\r\nRCPT TO:<user@example.com>\r\nDATA\r\nFrom: noreply@holacloud.app\r\nTo: user@example.com\r\nSubject: Test\r\n\r\nHello!\r\n.\r\nQUIT\r\n" | openssl s_client -connect smtp.hola.cloud:2525 -starttls smtp
```

## Go 代码示例

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

## 错误码

| 代码 | 描述 |
|------|-------------|
| 211 | 系统状态帮助回复 |
| 214 | 帮助消息 |
| 220 | 服务就绪 |
| 221 | 服务关闭传输通道 |
| 235 | 认证成功 |
| 250 | 请求的操作已完成 |
| 251 | 用户非本地；将转发 |
| 252 | 无法验证用户；将尝试投递 |
| 354 | 开始邮件输入 |
| 421 | 服务不可用 |
| 450 | 邮箱不可用 |
| 451 | 处理请求时本地错误 |
| 452 | 存储空间不足 |
| 500 | 语法错误，命令无法识别 |
| 501 | 参数语法错误 |
| 502 | 命令未实现 |
| 503 | 命令顺序错误 |
| 504 | 命令参数未实现 |
| 550 | 邮箱未找到 |
| 551 | 用户非本地 |
| 552 | 超出存储分配 |
| 553 | 邮箱名称不允许 |
| 554 | 事务失败 |

## 最佳实践

- **使用有意义的主题** 帮助收件人识别您的邮件并避免垃圾邮件过滤器。
- **设置有效的发件人地址** 与您的域名匹配以提高送达率。
- **优雅地处理错误** — 检查 SMTP 响应码并在临时故障（4xx）时重试。
- **使用 EHLO** 替代 HELO 以获取扩展 SMTP 功能。
- **流水线处理** — 无需等待每个响应即可发送多个命令以减少延迟。
- **速率限制** — 避免过快发送批量邮件；Holamail 可能会限制过多连接。
