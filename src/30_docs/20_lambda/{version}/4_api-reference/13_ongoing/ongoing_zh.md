# 列出正在运行的调用

返回当前已认证账户正在运行的 lambda 调用列表。

## 认证

需要 `X-Glue-Authentication`。

## HTTP 请求

```http
GET /api/v0/ongoing HTTP/1.1
Host: api.hola.cloud
X-Glue-Authentication: YOUR_TOKEN
```

## 示例

```bash
curl -X GET "https://api.hola.cloud/api/v0/ongoing" \
  -H "X-Glue-Authentication: YOUR_TOKEN"
```

## 响应

```json
[
  {
    "id": "inv_abc123",
    "lambda_id": "f1a2b3c4-d5e6-7890-abcd-ef0123456789",
    "lambda_name": "hello-world",
    "started_at": "2025-07-01T12:00:00Z",
    "method": "POST",
    "path": "/hello-world"
  }
]
```

## 错误码

| 代码 | 描述 |
|------|------|
| 401 | 认证缺失或无效 |
