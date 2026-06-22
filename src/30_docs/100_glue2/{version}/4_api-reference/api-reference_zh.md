# Glue2 API 参考

Glue2 是 HolaCloud 的中央 API 网关。大多数端点是公开的，但 `/v0/secret` 下的端点需要 glueauth.Require 身份验证。

## 端点

| 方法 | 路径 | 描述 |
|--------|------|-------------|
| `GET` | `/version` | 返回网关版本 |
| `GET` | `/v0/virtualhosts` | 列出路由表 |
| `GET` | `/v0/stats` | 获取流量统计 |
| `GET` | `/v0/status` | 检查后端服务健康状态 |
| `GET` | `/v0/secret` | 受保护的测试端点 |
| `GET` | `/openapi.json` | OpenAPI 规范 |
