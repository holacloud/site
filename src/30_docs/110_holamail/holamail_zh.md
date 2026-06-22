# Holamail

Holamail 是一个用于开发和测试流程的小型 SMTP listener。它接受基础 SMTP 会话，记录收到的消息，但不会向外部收件人投递邮件。

## 实现范围

Holamail 接受明文 SMTP 连接，并支持以下命令：

- `HELO` / `EHLO`
- `MAIL FROM`
- `RCPT TO`
- `DATA`
- `QUIT`

它不提供 HTTP API、外部投递、STARTTLS、AUTH、限速、模板、分析或追踪功能。

## 工作方式

应用使用 SMTP 客户端连接。Holamail 会把 envelope 和消息内容写入日志，便于检查。

```
┌─────────────┐    SMTP    ┌───────────┐
│ Application │───────────▶│ Holamail  │──▶ log message
└─────────────┘            └───────────┘
```

## 快速开始

使用公共测试主机：

```bash
swaks --to user@example.com \
      --from noreply@example.com \
      --server smtp.testmail.hola.cloud \
      --port 25 \
      --header "Subject: Holamail test" \
      --body "Hello from Holamail."
```

或使用本地 listener：

```bash
swaks --to user@example.com \
      --from noreply@example.com \
      --server localhost \
      --port 2525 \
      --header "Subject: Local Holamail test" \
      --body "Hello from a local Holamail listener."
```
