
# InceptionDB API 参考

基础 URL：`https://api.hola.cloud`

## 身份验证

InceptionDB 管理端点需要通过两个请求头进行身份验证：

- `Api-Key` — 您的 API 密钥（UUID）
- `Api-Secret` — 您的 API 密钥密码（UUID）
- `X-Project` — 将请求限定到特定项目

所有请求必须通过 HTTPS 发送。

## 端点

| 方法 | 路径 | 描述 |
|--------|------|------|
| GET | `/v1/databases` | 列出所有数据库 |
| POST | `/v1/databases` | 创建新数据库 |
| GET | `/v1/databases/{id}` | 根据 ID 获取数据库 |
| DELETE | `/v1/databases/{id}` | 删除数据库 |
| PATCH | `/v1/databases/{id}` | 更新数据库 |
| GET | `/v1/databases/{id}/collections` | 列出数据库中的集合 |
| POST | `/v1/databases/{id}/collections` | 创建集合 |
| POST | `/v1/databases/{id}/collections/{col}` | 文档操作（插入、查找、删除、更新） |

## 常见错误码

| 状态码 | 描述 |
|--------|------|
| 400 | 错误请求 — 语法错误或参数无效 |
| 401 | 未授权 — 缺少或无效的 API 凭证 |
| 403 | 禁止访问 — 凭证无权限访问该资源 |
| 404 | 未找到 — 请求的资源不存在 |
| 409 | 冲突 — 资源已存在 |
| 500 | 服务器内部错误 |
