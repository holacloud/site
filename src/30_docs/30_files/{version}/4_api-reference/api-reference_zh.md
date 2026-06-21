# Files API 参考

基础 URL：`https://api.hola.cloud`

所有端点都需要通过 `Api-Key` 和 `Api-Secret` 标头进行身份验证。

## 端点

| 方法 | 路径 | 描述 |
|--------|------|-------------|
| GET | `/v1/buckets` | 列出所有存储桶 |
| POST | `/v1/buckets` | 创建新的存储桶 |
| GET | `/v1/buckets/{id}` | 获取存储桶详情 |
| PATCH | `/v1/buckets/{id}` | 修改存储桶元数据 |
| DELETE | `/v1/buckets/{id}` | 删除存储桶 |
| GET | `/v1/buckets/{id}/list/*` | 列出存储桶中的文件 |
| PUT | `/v1/buckets/{id}/files/*` | 上传文件 |
| GET | `/v1/buckets/{id}/files/*` | 下载文件 |
| DELETE | `/v1/buckets/{id}/files/*` | 删除文件 |
| HEAD | `/v1/buckets/{id}/files/*` | 获取文件元数据 |

## 常见错误

| 状态 | 代码 | 描述 |
|--------|------|-------------|
| 400 | Bad Request | 请求体或参数无效 |
| 401 | Unauthorized | 缺少或无效的 API 凭证 |
| 403 | Forbidden | 权限不足 |
| 404 | Not Found | 指定的资源不存在 |
| 409 | Conflict | 资源已存在或操作冲突 |
| 413 | Payload Too Large | 文件超过最大大小限制（5 GB） |
| 429 | Too Many Requests | 超出速率限制 |
| 500 | Internal Server Error | 发生意外错误 |
