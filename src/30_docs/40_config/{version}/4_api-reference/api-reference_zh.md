# Config API 参考

基础 URL：`https://api.hola.cloud`

## 认证

`/v1/config` 需要 `X-Glue-Authentication`。`/v0/configs` 直接管理配置 Thing。

## 端点

| 方法 | 路径 | 描述 | 认证 |
|--------|------|-------------|------|
| GET | `/v0/configs` | 列出配置 Thing | 公开 |
| POST | `/v0/configs` | 创建包含 `entries` 的配置 Thing | 公开 |
| GET | `/v0/configs/{id}` | 按 ID 获取配置 Thing | 公开 |
| DELETE | `/v0/configs/{id}` | 删除配置 Thing | 公开 |
| PATCH | `/v0/configs/{id}` | 修补配置 Thing 的 `entries` | 公开 |
| GET | `/v1/config` | 获取已认证用户配置的 `entries` 映射 | Glue 认证 |
| PATCH | `/v1/config` | 更新已认证用户配置的 `entries` 映射 | Glue 认证 |

## 结构

v0 配置 Thing：

```json
{
  "id": "cfg_abc123",
  "entries": {
    "feature.newCheckout": true
  }
}
```

v1 用户配置：

```json
{
  "entries": {
    "feature.newCheckout": true
  }
}
```
