# 获取版本

返回 Run 服务的当前版本。

## 描述

提供有关正在运行的 Run 服务实例的版本和构建信息。

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
  "service": "run",
  "version": "1.5.2",
  "commit": "e5f6g7h8",
  "build_time": "2026-06-20T12:00:00Z"
}
```

## 错误码

| 状态码 | 描述 |
|--------|------|
| 200 | 成功返回版本信息 |
