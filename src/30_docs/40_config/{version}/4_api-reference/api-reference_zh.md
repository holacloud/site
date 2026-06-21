# Config API 参考

基础 URL：`https://api.hola.cloud`

Config 服务提供两个 API 接口：

- **管理 API**（`/v0/configs`）— 公开访问，用于管理配置
- **用户 API**（`/v1/config`）— 需要 `Api-Key` 和 `Api-Secret` 标头，用于运行时配置访问

## 端点

| 方法 | 路径 | 描述 | 认证 |
|--------|------|-------------|------|
| GET | `/v0/configs` | 列出所有配置（管理） | 公开 |
| POST | `/v0/configs` | 创建新配置（管理） | 公开 |
| GET | `/v0/configs/{id}` | 按 ID 获取配置（管理） | 公开 |
| DELETE | `/v0/configs/{id}` | 删除配置（管理） | 公开 |
| PATCH | `/v0/configs/{id}` | 部分更新配置（管理） | 公开 |
| GET | `/v1/config` | 获取当前用户配置 | API 密钥 |
| PATCH | `/v1/config` | 更新当前用户配置 | API 密钥 |

## 常见错误

| 状态 | 代码 | 描述 |
|--------|------|-------------|
| 400 | Bad Request | 请求体或参数无效 |
| 401 | Unauthorized | 缺少或无效的 API 凭证 |
| 404 | Not Found | 指定的资源不存在 |
| 409 | Conflict | 资源已存在或操作冲突 |
| 500 | Internal Server Error | 发生意外错误 |
