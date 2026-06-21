
# POST /v1/ingest/events

使用 logframe 协议接收帧事件。此端点接受带有元数据和消息长度前缀的二进制协议。

## 认证

需要数据凭据：

- `X-Instantlogs-Event-Secret: <secret>` 或 `Authorization: Bearer <secret>`

## 请求体

请求体使用 logframe 协议格式：

```
len_metadata:len_message:metadata_json:message\n
```

每个帧包含：
- `len_metadata` — 元数据 JSON 的长度（ASCII 数字）
- `len_message` — 消息的长度（ASCII 数字）
- `metadata_json` — 包含事件元数据的 JSON 对象
- `message` — 日志消息内容
- `\n` — 换行符分隔符

## 请求

```bash
# 帧：metadata={"service":"web"} + message="用户已登录"
# len_metadata=18, len_message=15
# 格式：18:15:{"service":"web"}用户已登录\n
printf '18:15:{"service":"web"}用户已登录\n' | \
curl -X POST "https://api.hola.cloud/v1/ingest/events" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456" \
  -H "Content-Type: application/octet-stream" \
  --data-binary @-
```

```http
POST /v1/ingest/events HTTP/1.1
Host: api.hola.cloud
X-Instantlogs-Event-Secret: lgs_abc123def456
Content-Type: application/octet-stream

18:15:{"service":"web"}用户已登录\n
```

## 响应

```json
{
  "ingested": 1
}
```

## 错误代码

| 代码 | 描述 |
|------|------|
| 400 | 帧格式错误 — 长度前缀无效或缺少分隔符 |
| 401 | 缺少或无效的事件密钥 |
| 413 | 负载过大 |
