# 健康检查

返回服务的健康状态。

## 请求

```http
GET /healthz
```

## 示例

```bash
curl "https://api.hola.cloud/healthz"
```

## 响应

服务健康时返回 `ok` 及 200 状态码。

```text
ok
```
