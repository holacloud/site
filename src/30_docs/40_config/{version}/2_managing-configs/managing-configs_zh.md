# 管理 Config

v0 API 直接管理配置 Thing。配置 Thing 包含 `id` 和 `entries` 对象。

## 列出

```bash
curl "https://api.hola.cloud/v0/configs"
```

```json
[
  {
    "id": "cfg_abc123",
    "entries": {
      "feature.newCheckout": true
    }
  }
]
```

## 创建

```bash
curl -X POST "https://api.hola.cloud/v0/configs" \
  -H "Content-Type: application/json" \
  -d '{"entries":{"feature.newCheckout":true}}'
```

## 获取

```bash
curl "https://api.hola.cloud/v0/configs/cfg_abc123"
```

## 修补

```bash
curl -X PATCH "https://api.hola.cloud/v0/configs/cfg_abc123" \
  -H "Content-Type: application/json" \
  -d '{"entries":{"feature.newCheckout":false}}'
```

## 删除

```bash
curl -X DELETE "https://api.hola.cloud/v0/configs/cfg_abc123"
```
