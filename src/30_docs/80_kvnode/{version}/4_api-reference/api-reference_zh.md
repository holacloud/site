# KVNode API 参考

基础 URL：`https://api.hola.cloud`

| 方法 | 路径 | 描述 | 认证 |
|------|------|------|------|
| GET | /healthz | 健康检查 | 公开 |
| GET | /readyz | 就绪检查 | 公开 |
| GET | /v1/status | 节点状态 | 内部 |
| GET | /v1/metrics | 节点指标 | 内部 |
| GET | /v1/collections | 列出集合 | 内部 |
| POST | /v1/collections | 创建集合 | 内部 |
| DELETE | /v1/collections/{col} | 删除集合 | 内部 |
| GET | /v1/collections/{col}/keys | 列出键 | 内部 |
| POST | /v1/collections/{col}/keys/* | 设置键 | 内部 |
| GET | /v1/collections/{col}/keys/* | 获取键 | 内部 |
| DELETE | /v1/collections/{col}/keys/* | 删除键 | 内部 |

## 身份验证

健康检查端点为公开。所有其他端点需要通过以下方式之一进行身份验证：

- `X-Glue-Authentication` 头部
- `apikey` 和 `secret` 头部

## 错误代码

| 状态 | 代码 | 描述 |
|------|------|------|
| 400 | invalid_request | 请求体或参数无效 |
| 401 | unauthorized | 缺少或无效的身份验证 |
| 404 | not_found | 未找到集合或键 |
| 409 | conflict | 集合已存在 |
| 500 | internal_error | 服务器内部错误 |
