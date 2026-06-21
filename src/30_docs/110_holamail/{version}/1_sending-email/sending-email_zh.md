# 使用Holamail发送邮件

Holamail 是一个轻量级的 SMTP 服务器，用于发送事务性电子邮件。它实现了标准的 SMTP 协议（HELO/EHLO、MAIL FROM、RCPT TO、DATA、QUIT），监听 TCP 端口 2525，无需 TLS。

## 连接信息

| 参数 | 值 |
|------|-----|
| 主机 | `smtp.hola.cloud` |
| 端口 | `2525` |
| 加密 | 无（纯 SMTP） |

## SMTP 协议流程

原始 SMTP 对话遵循以下顺序：

```
HELO client.example.com
MAIL FROM:<noreply@holacloud.app>
RCPT TO:<user@example.com>
DATA
Subject: 来自Holamail的问候

这是邮件正文。
.
QUIT
```

## 示例

### 使用 curl

curl 原生支持 SMTP：

```bash
curl --mail-from noreply@holacloud.app \
     --mail-rcpt user@example.com \
     --upload-file email.txt \
     smtp://smtp.hola.cloud:2525
```

其中 `email.txt` 包含：

```
Subject: 来自Holamail的问候

这是邮件正文。
```

### 使用 openssl s_client

通过纯 TCP 连接手动进行 SMTP 通信：

```bash
exec 3<>/dev/tcp/smtp.hola.cloud/2525
echo "HELO client.example.com" >&3
echo "MAIL FROM:<noreply@holacloud.app>" >&3
echo "RCPT TO:<user@example.com>" >&3
echo "DATA" >&3
echo "Subject: 你好" >&3
echo "" >&3
echo "来自Holamail的问候！" >&3
echo "." >&3
echo "QUIT" >&3
cat <&3
```

### 使用 Go 的 net/smtp

```go
package main

import (
    "log"
    "net/smtp"
)

func main() {
    to := []string{"user@example.com"}
    msg := []byte("To: user@example.com\r\n" +
        "Subject: 来自Holamail的问候\r\n" +
        "\r\n" +
        "这是邮件正文。\r\n")

    err := smtp.SendMail("smtp.hola.cloud:2525", nil,
        "noreply@holacloud.app", to, msg)
    if err != nil {
        log.Fatal(err)
    }
}
```

### 使用 Python 的 smtplib

```python
import smtplib
from email.message import EmailMessage

msg = EmailMessage()
msg.set_content("这是邮件正文。")
msg["Subject"] = "来自Holamail的问候"
msg["From"] = "noreply@holacloud.app"
msg["To"] = "user@example.com"

with smtplib.SMTP("smtp.hola.cloud", 2525) as server:
    server.send_message(msg)
```

## 使用场景

- **密码重置** — 按需发送安全的重置链接。
- **通知** — 发送警报、提醒和状态更新。
- **系统告警** — 从监控系统（Nagios、Prometheus、Grafana）触发邮件。
- **Lambda 集成** — 从 Lambda 函数调用 Holamail 以编程方式发送邮件：

```python
import smtplib
from email.message import EmailMessage

def handler(event, context):
    msg = EmailMessage()
    msg.set_content(event.get("body", ""))
    msg["Subject"] = event.get("subject", "通知")
    msg["From"] = "noreply@holacloud.app"
    msg["To"] = event["to"]

    with smtplib.SMTP("smtp.hola.cloud", 2525) as server:
        server.send_message(msg)

    return {"statusCode": 200, "body": "邮件已发送"}
```

## 错误处理

Holamail 返回标准 SMTP 状态码：

| 状态码 | 含义 |
|--------|------|
| 250 | 请求完成 |
| 550 | 邮箱不可用（收件人无效） |
| 551 | 收件人被拒绝 |
| 552 | 超出存储配额 |
| 553 | 地址无效 |
| 554 | 事务失败 |

当收件人地址无效时，服务器返回 `550`。请在你的应用程序代码中处理连接错误（超时、拒绝连接）——这些通常表示网络问题或服务暂时不可用。
