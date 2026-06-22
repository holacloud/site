# 列出Buckets

列出经过身份验证的账户的所有存储桶。

## 身份验证

需要 `Api-Key` 和 `Api-Secret` 标头。

## 查询参数

| 参数 | 类型 | 默认值 | 描述 |
|-----------|------|---------|-------------|
| `limit` | integer | 100 | 返回的最大存储桶数量 |
| `offset` | integer | 0 | 跳过的存储桶数量 |

## 请求

```bash
curl "https://api.hola.cloud/v1/buckets?limit=10&offset=0" \
  -H "Api-Key: 您的API密钥" \
  -H "Api-Secret: 您的API密钥密码"
```

## 响应

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "buckets": [
    {
      "id": "bkt_abc123",
      "name": "my-first-bucket",
      "createdAt": "2026-06-21T10:00:00Z",
      "size": 1048576,
      "fileCount": 5,
      "public": false
    },
    {
      "id": "bkt_def456",
      "name": "assets",
      "createdAt": "2026-06-20T08:30:00Z",
      "size": 52428800,
      "fileCount": 42,
      "public": true
    }
  ],
  "total": 2
}
```

## 错误代码

| 状态 | 代码 | 描述 |
|--------|------|-------------|
| 401 | Unauthorized | 缺少或无效的 API 凭证 |
| 403 | Forbidden | 权限不足 |
| 500 | Internal Server Error | 发生意外错误 |
