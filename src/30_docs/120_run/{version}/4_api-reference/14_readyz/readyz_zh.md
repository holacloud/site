# 就绪检查

返回服务的就绪状态。

## 请求

```http
GET /readyz
```

## 示例

```bash
curl "https://api.hola.cloud/readyz"
```

## 响应

服务就绪可接收流量时返回 `ok` 及 200 状态码。

```text
ok
```
