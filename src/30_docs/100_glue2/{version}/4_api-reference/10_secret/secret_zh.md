# 密钥

用于验证认证的受保护测试端点。

## 描述

此端点需要有效的认证令牌，用于验证认证层是否正常工作。

## 认证

需要通过有效的 API 密钥或会话进行 `glueauth.Require` 认证。

## 请求

无需请求体。

## 示例

```bash
curl -X GET "https://api.hola.cloud/v0/secret" \
  -H "Authorization: Bearer <api-key>"
```

## 响应

```text
OK
```

## 错误代码

| 代码 | 描述 |
|------|------|
| 200 | 认证成功 |
| 401 | 缺少或无效的认证 |
