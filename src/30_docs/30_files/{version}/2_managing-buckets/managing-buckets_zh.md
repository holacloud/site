# 管理存储桶

存储桶是 HolaCloud Files 中的基本容器。本指南涵盖所有存储桶操作：创建、列出、查看、修改和删除。

## 存储桶命名规则

创建存储桶时，名称必须遵循以下规则：

- 长度必须在 3 到 63 个字符之间
- 只能包含小写字母、数字和连字符 (-)
- 必须以字母或数字开头和结尾
- 不能格式化为 IP 地址
- 必须在所有 HolaCloud 客户中全局唯一

## 创建存储桶

使用 POST 请求创建新存储桶：

```bash
curl -X POST "https://api.hola.cloud/v1/buckets" \
  -H "Api-Key: 您的API密钥" \
  -H "Api-Secret: 您的API密钥密码" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "my-assets-bucket",
    "metadata": {
      "environment": "production",
      "project": "my-app"
    }
  }'
```

等效 HTTP 请求：

```http
POST /v1/buckets HTTP/1.1
Host: api.hola.cloud
Api-Key: 您的API密钥
Api-Secret: 您的API密钥密码
Content-Type: application/json

{
  "name": "my-assets-bucket",
  "metadata": {
    "environment": "production",
    "project": "my-app"
  }
}
```

预期响应 (HTTP 201)：

```json
{
  "id": "bkt_xyz789",
  "name": "my-assets-bucket",
  "metadata": {
    "environment": "production",
    "project": "my-app"
  },
  "createdAt": "2026-06-21T10:00:00Z",
  "size": 0,
  "fileCount": 0
}
```

## 列出存储桶

列出您账户中的所有存储桶：

```bash
curl "https://api.hola.cloud/v1/buckets" \
  -H "Api-Key: 您的API密钥" \
  -H "Api-Secret: 您的API密钥密码"
```

预期响应：

```json
[
  {
    "id": "bkt_abc123",
    "name": "my-first-bucket",
    "createdAt": "2026-06-21T09:00:00Z",
    "size": 1024,
    "fileCount": 3
  },
  {
    "id": "bkt_xyz789",
    "name": "my-assets-bucket",
    "createdAt": "2026-06-21T10:00:00Z",
    "size": 0,
    "fileCount": 0
  }
]
```

## 获取存储桶详情

通过 ID 检索特定存储桶的详细信息：

```bash
curl "https://api.hola.cloud/v1/buckets/bkt_xyz789" \
  -H "Api-Key: 您的API密钥" \
  -H "Api-Secret: 您的API密钥密码"
```

预期响应：

```json
{
  "id": "bkt_xyz789",
  "name": "my-assets-bucket",
  "metadata": {
    "environment": "production",
    "project": "my-app"
  },
  "createdAt": "2026-06-21T10:00:00Z",
  "modifiedAt": "2026-06-21T10:00:00Z",
  "size": 0,
  "fileCount": 0
}
```

## 修改存储桶

使用 PATCH 更新存储桶元数据：

```bash
curl -X PATCH "https://api.hola.cloud/v1/buckets/bkt_xyz789" \
  -H "Api-Key: 您的API密钥" \
  -H "Api-Secret: 您的API密钥密码" \
  -H "Content-Type: application/json" \
  -d '{
    "metadata": {
      "environment": "staging"
    }
  }'
```

预期响应：

```json
{
  "id": "bkt_xyz789",
  "name": "my-assets-bucket",
  "metadata": {
    "environment": "staging",
    "project": "my-app"
  },
  "createdAt": "2026-06-21T10:00:00Z",
  "modifiedAt": "2026-06-21T11:00:00Z",
  "size": 0,
  "fileCount": 0
}
```

## 删除存储桶

存储桶必须**为空**（不包含文件）才能删除。

```bash
curl -X DELETE "https://api.hola.cloud/v1/buckets/bkt_xyz789" \
  -H "Api-Key: 您的API密钥" \
  -H "Api-Secret: 您的API密钥密码"
```

预期响应：HTTP 204 无内容。

### 错误：存储桶不为空

如果尝试删除仍包含文件的存储桶，您将收到：

```json
{
  "error": {
    "code": "BUCKET_NOT_EMPTY",
    "message": "无法删除存储桶 'my-assets-bucket'：该存储桶包含 3 个文件。请先删除所有文件再删除存储桶。"
  }
}
```

## 错误处理

### 存储桶名称重复

```json
{
  "error": {
    "code": "BUCKET_ALREADY_EXISTS",
    "message": "名称为 'my-assets-bucket' 的存储桶已存在。"
  }
}
```

### 无效的存储桶名称

```json
{
  "error": {
    "code": "INVALID_BUCKET_NAME",
    "message": "存储桶名称长度必须在 3 到 63 个字符之间，只能包含小写字母、数字和连字符，并且以字母或数字开头和结尾。"
  }
}
```

### 存储桶未找到

```json
{
  "error": {
    "code": "BUCKET_NOT_FOUND",
    "message": "未找到存储桶 'bkt_nonexistent'。"
  }
}
```

## HTTP 状态码汇总

| 状态码 | 描述 |
|--------|------|
| 200 | 成功 (GET, PATCH) |
| 201 | 已创建 (POST) |
| 204 | 无内容 (DELETE) |
| 400 | 错误请求 -- 输入无效 |
| 404 | 未找到 -- 存储桶不存在 |
| 409 | 冲突 -- 存储桶已存在或不为空 |
| 500 | 服务器内部错误 |
