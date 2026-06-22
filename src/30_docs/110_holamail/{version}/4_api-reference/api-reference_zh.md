# Holamail API 参考

Holamail 不暴露 HTTP API。它是一个明文 SMTP listener，接受基础 SMTP 会话并记录消息。

## SMTP 参考

命令示例请参见 [SMTP 协议参考](./smtp-reference)。

## 连接信息

| 设置 | 公共测试主机 | 本地实例 |
|------|--------------|----------|
| Host | `smtp.testmail.hola.cloud` | `localhost` |
| Port | `25` | `2525` |
| Protocol | Plain SMTP | Plain SMTP |
| Auth | None | None |

## 支持的命令

| 命令 | 描述 |
|------|------|
| `HELO` / `EHLO` | 开始 SMTP 会话 |
| `MAIL FROM` | 设置 envelope 发件人 |
| `RCPT TO` | 设置 envelope 收件人 |
| `DATA` | 发送消息头和正文，以 `.` 结束 |
| `QUIT` | 关闭会话 |

## 不支持

Holamail 不提供外部投递、HTTP API、STARTTLS、AUTH、限速、模板、分析或追踪。
