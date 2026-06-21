# Holamail

Holamail 是一个简单的 SMTP 邮件服务，用于发送事务性电子邮件。它在 2525 端口上监听原始 SMTP 连接，并将消息投递给收件人。

## 工作原理

Holamail 通过 **SMTP 协议**在 2525 端口上接受连接。客户端使用标准 SMTP 进行连接——任何具有 SMTP 客户端库的语言或框架都可以通过 Holamail 发送电子邮件，无需额外的 SDK。

```
┌──────────────┐   SMTP :2525    ┌───────────┐
│  应用程序    │───────────────▶│  Holamail  │────▶  收件人
└──────────────┘                 └───────────┘
```

## 使用场景

- **事务性邮件**：订单确认、账户验证、密码重置。
- **通知**：由应用程序事件触发的警报、提醒、状态更新。
- **密码重置**：安全地向用户投递密码重置链接。
- **系统警报**：从监控或 CI/CD 流水线发送自动警报。

## 与其他服务的集成

Holamail 与其他 HolaCloud 服务协同工作：

- **Lambda** — 从无服务器函数触发邮件发送。
- **控制台** — 管理邮件模板和查看投递日志。
- **InceptionDB** — 存储邮件模板和投递历史。

## 快速入门

使用任何 SMTP 客户端连接到 Holamail 的 2525 端口：

```bash
# 使用 swaks（SMTP 瑞士军刀）
swaks --to user@example.com \
      --from noreply@holacloud.app \
      --server smtp.hola.cloud \
      --port 2525 \
      --header "Subject: 欢迎使用 HolaCloud" \
      --body "您好！感谢您的注册。"
```

或使用 Python 的 `smtplib`：

```python
import smtplib
from email.message import EmailMessage

msg = EmailMessage()
msg.set_content("您好！感谢您的注册。")
msg["Subject"] = "欢迎使用 HolaCloud"
msg["From"] = "noreply@holacloud.app"
msg["To"] = "user@example.com"

with smtplib.SMTP("smtp.hola.cloud", 2525) as s:
    s.send_message(msg)
```
