
# Lambda API 参考

基础 URL：`https://api.hola.cloud`

## 身份验证

管理端点需要 `Api-Key` 和 `Api-Secret` 请求头。公共调用端点（`/run/{id}` 和 `/mux/{owner}/*`）无需身份验证。

所有请求必须通过 HTTPS 发送。

## 端点

| 方法 | 路径 | 描述 |
|--------|------|------|
| GET | `/api/v0/lambdas` | 列出所有 Lambda 函数 |
| POST | `/api/v0/lambdas` | 创建新的 Lambda 函数 |
| GET | `/api/v0/lambdas/{id}` | 根据 ID 获取 Lambda 函数 |
| PATCH | `/api/v0/lambdas/{id}` | 修改 Lambda 函数 |
| DELETE | `/api/v0/lambdas/{id}` | 删除 Lambda 函数 |
| POST | `/api/v0/run/{id}` | 运行 Lambda 函数（管理员） |
| POST | `/run/{id}` | 运行 Lambda 函数（公共） |
| GET | `/mux/{owner}/*` | 路由到所有者的 Lambda 函数 |

## 常见错误码

| 状态码 | 描述 |
|--------|------|
| 400 | 错误请求 — 语法错误或参数无效 |
| 401 | 未授权 — 缺少或无效的 API 凭证 |
| 403 | 禁止访问 — 凭证无权限访问该资源 |
| 404 | 未找到 — 请求的资源不存在 |
| 409 | 冲突 — 资源已存在 |
| 500 | 服务器内部错误 |
