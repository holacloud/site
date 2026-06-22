# SMTP 参考

Holamail 实现基础的明文 SMTP listener。它接受 SMTP 消息并记录日志，不会外部投递邮件。

## 连接

公共测试主机：

```bash
telnet smtp.testmail.hola.cloud 25
```

本地实例：

```bash
telnet localhost 2525
```

不支持 STARTTLS 和 SMTP AUTH。

## SMTP 命令

### HELO / EHLO

开始 SMTP 会话。

```text
EHLO client.example.com
```

### MAIL FROM

指定发件人地址。

```text
MAIL FROM:<noreply@example.com>
```

### RCPT TO

指定收件人地址。

```text
RCPT TO:<user@example.com>
```

### DATA

开始消息内容。用只包含点号 (`.`) 的一行结束。

```text
DATA
From: noreply@example.com
To: user@example.com
Subject: Holamail test

This message will be logged by Holamail.
.
```

### QUIT

结束 SMTP 会话。

```text
QUIT
```

## 完整会话示例

```bash
printf 'EHLO client.example.com\r\nMAIL FROM:<noreply@example.com>\r\nRCPT TO:<user@example.com>\r\nDATA\r\nFrom: noreply@example.com\r\nTo: user@example.com\r\nSubject: Test\r\n\r\nHello!\r\n.\r\nQUIT\r\n' | nc localhost 2525
```

## 常见状态码

| Code | Description |
|------|-------------|
| `220` | Service ready |
| `221` | Service closing transmission channel |
| `250` | Requested action completed |
| `354` | Start mail input |
| `500` | Syntax error or unrecognized command |
| `502` | Command not implemented |
| `503` | Bad command sequence |

## 范围

Holamail 只记录已接受的消息。它不包含外部投递、HTTP API、STARTTLS、AUTH、限速、模板、分析或追踪。
