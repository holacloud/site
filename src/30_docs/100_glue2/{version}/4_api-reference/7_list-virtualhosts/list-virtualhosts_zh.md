# GET /v0/virtualhosts

返回当前虚拟主机路由表。

## 描述

此端点列出所有配置的虚拟主机及其对应的后端服务目标。HolaCloud 控制台使用它来显示路由配置。

## 身份验证

无。此端点为公开端点。

## 请求

无需请求体。

## 示例

```bash
curl -X GET "https://api.hola.cloud/v0/virtualhosts"
```

## 响应

```json
{
  "virtualhosts": [
    {
      "host": "my-project.hola.cloud",
      "target": "inceptiondb",
      "port": 8080,
      "tls": true
    },
    {
      "host": "api.my-project.hola.cloud",
      "target": "lambda",
      "port": 8081,
      "tls": true
    }
  ]
}
```

## 错误码

| 状态码 | 描述 |
|--------|------|
| 200 | 成功返回路由表 |
