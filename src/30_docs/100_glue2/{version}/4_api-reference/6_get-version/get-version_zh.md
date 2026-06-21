# GET /version

返回 Glue2 网关的当前版本。

## 描述

版本端点提供有关正在运行的网关实例的构建和发布信息。

## 身份验证

无。此端点为公开端点。

## 请求

无需请求体。

## 示例

```bash
curl -X GET "https://api.hola.cloud/version"
```

## 响应

```json
{
  "service": "glue2",
  "version": "2.3.1",
  "commit": "a1b2c3d4",
  "build_time": "2026-06-20T12:00:00Z"
}
```

## 错误码

| 状态码 | 描述 |
|--------|------|
| 200 | 成功返回版本信息 |
