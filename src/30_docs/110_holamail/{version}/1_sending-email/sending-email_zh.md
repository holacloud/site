# 发送邮件

Holamail 是一个轻量级 SMTP listener。它接受基础命令 `HELO`/`EHLO`、`MAIL FROM`、`RCPT TO`、`DATA` 和 `QUIT`，然后记录消息。它不会向外部收件人投递邮件。

## 连接信息

| 参数 | 公共测试主机 | 本地实例 |
|------|--------------|----------|
| Host | `smtp.testmail.hola.cloud` | `localhost` |
| Port | `25` | `2525` |
| Encryption | None | None |
| Authentication | None | None |

Holamail 不支持 STARTTLS 或 SMTP AUTH。

## SMTP 流程

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

## 示例

### 使用 curl

```bash
curl --url smtp://smtp.testmail.hola.cloud:25 \
     --mail-from noreply@example.com \
     --mail-rcpt user@example.com \
     --upload-file email.txt
```

`email.txt` 内容：

```text
Subject: Hello from Holamail

This message will be logged by Holamail.
```

### 明文 TCP 会话

```bash
telnet localhost 2525
```

然后输入：

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

## 预期行为

Holamail 接受语法有效的 SMTP 输入并记录消息。它适合验证应用能否使用 SMTP，而不会发送真实邮件。

不支持外部投递、HTTP API、STARTTLS、AUTH、限速、模板、分析或追踪。
