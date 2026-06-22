# Lambda API 参考

基础 URL：`https://api.hola.cloud`

## 认证

管理端点需要 `X-Glue-Authentication` header。公开调用端点（`/run/{lambda_id}` 和 `/mux/{owner_id}/*`）不需要认证。

所有请求都必须通过 HTTPS 发送。

## Lambda 对象

| 字段 | 类型 | 描述 |
|------|------|------|
| `id` | string | Lambda 标识符 |
| `created_timestamp` | number | Unix timestamp 格式的创建时间 |
| `owner` | string | 所有者标识符 |
| `project_id` | string | 项目标识符 |
| `name` | string | Lambda 名称 |
| `language` | string | `javascript`、`static-html`、`static-css` 或 `static-js` |
| `code` | string | 源码或静态内容 |
| `method` | string | 与 lambda 关联的 HTTP 方法 |
| `path` | string | 与 lambda 关联的 HTTP path |

## 端点

| 方法 | Path | 描述 |
|------|------|------|
| POST | `/api/v0/lambdas` | 创建 lambda |
| GET | `/api/v0/lambdas` | 列出 lambdas |
| GET | `/api/v0/lambdas/{lambda_id}` | 按 ID 获取 lambda |
| PATCH | `/api/v0/lambdas/{lambda_id}` | 修改 lambda |
| DELETE | `/api/v0/lambdas/{lambda_id}` | 删除 lambda |
| ANY | `/api/v0/run/{lambda_id}` | 使用认证运行 lambda |
| ANY | `/run/{lambda_id}` | 公开运行 lambda |
| ANY | `/mux/{owner_id}/*` | 通过所有者范围路由请求 |
| GET | `/ongoing` | 列出当前运行中的调用 |
| GET | `/me` | 返回已认证用户 |
| GET | `/openapi` | 返回 OpenAPI 文档 |

## 常见错误码

| 代码 | 描述 |
|------|------|
| 400 | 请求无效 |
| 401 | 认证缺失或无效 |
| 403 | 禁止访问 |
| 404 | 资源未找到 |
| 500 | 内部服务器错误 |
