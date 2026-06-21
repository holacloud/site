# GET /fakeapi/*

通配符端点，在开发和演示模式下为所有控制台后端服务提供模拟数据。

## 描述

`/fakeapi/*` 路由匹配 `/fakeapi/` 下的任何路径，并返回模拟响应。这使得控制台 UI 无需真实后端即可运行。路径结构镜像了真实 API 的布局，因此 UI 代码在开发和生产环境之间保持一致。

## 子路由

| 子路径 | 描述 |
|--------|------|
| `/fakeapi/auth/me` | 返回当前模拟用户会话 |
| `/fakeapi/inceptionapi` | 模拟 InceptionDB API 响应 |
| `/fakeapi/lambdasapi` | 模拟 Lambda 函数管理 |
| `/fakeapi/projectsapi` | 模拟项目设置和列表 |
| `/fakeapi/filesapi` | 模拟文件存储操作 |

## 身份验证

无。这些模拟端点无需身份验证。

## 示例

```bash
curl -X GET "https://api.hola.cloud/fakeapi/auth/me"
```

```json
{
  "user": "demo@holacloud.com",
  "name": "演示用户",
  "project": "default",
  "role": "admin",
  "authenticated": true
}
```

```bash
curl -X GET "https://api.hola.cloud/fakeapi/projectsapi"
```

```json
{
  "projects": [
    {
      "id": "default",
      "name": "默认项目",
      "plan": "free",
      "region": "us-east"
    }
  ]
}
```

```bash
curl -X GET "https://api.hola.cloud/fakeapi/lambdasapi"
```

```json
{
  "functions": [
    {
      "id": "fn-001",
      "name": "hello-world",
      "runtime": "nodejs18",
      "status": "active"
    }
  ]
}
```

## 错误码

| 状态码 | 描述 |
|--------|------|
| 200 | 成功返回模拟数据 |
| 404 | `/fakeapi/` 下的路径无法识别 |
