# 列出Files

列出存储桶中的文件。通配符（`*`）可以替换为可选的前缀路径以过滤结果。

## 身份验证

需要 `Api-Key` 和 `Api-Secret` 标头。

## 路径参数

| 参数 | 类型 | 描述 |
|-----------|------|-------------|
| `id` | string | 存储桶 ID（例如 `bkt_abc123`） |

## 查询参数

| 参数 | 类型 | 默认值 | 描述 |
|-----------|------|---------|-------------|
| `prefix` | string | "" | 筛选以此前缀开头的文件 |
| `delimiter` | string | "" | 使用此分隔符分组结果（例如 `/`） |
| `maxKeys` | integer | 1000 | 返回的最大文件数量 |

## 请求

```bash
curl "https://api.hola.cloud/v1/buckets/bkt_abc123/list/*?prefix=images/&maxKeys=50" \
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
  "files": [
    {
      "path": "images/logo.png",
      "size": 24576,
      "contentType": "image/png",
      "modifiedAt": "2026-06-21T11:00:00Z",
      "etag": "\"abc123def456\""
    },
    {
      "path": "images/banner.jpg",
      "size": 102400,
      "contentType": "image/jpeg",
      "modifiedAt": "2026-06-21T10:30:00Z",
      "etag": "\"789012ghi345\""
    }
  ],
  "prefix": "images/",
  "total": 2
}
```

## 错误代码

| 状态 | 代码 | 描述 |
|--------|------|-------------|
| 401 | Unauthorized | 缺少或无效的 API 凭证 |
| 404 | Not Found | 指定的存储桶不存在 |
| 500 | Internal Server Error | 发生意外错误 |
