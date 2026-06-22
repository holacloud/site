# 模拟Api

固定端点，在开发和演示模式下为控制台提供模拟数据。

## 描述

控制台只公开下面列出的 fake API 路由。未列出的路径不属于当前控制台 fake API。

## 子路由

| 方法 | 路径 | 描述 |
|------|------|------|
| `GET` | `/fakeapi/authapi/auth/me` | 返回当前模拟用户 |
| `GET` | `/fakeapi/inceptionapi/v1/databases` | 列出模拟数据库 |
| `GET` | `/fakeapi/inceptionapi/v1/databases/{databaseid}/collections` | 列出模拟集合 |
| `GET` | `/fakeapi/lambdasapi/api/v0/lambdas` | 列出模拟 lambdas |
| `GET` | `/fakeapi/projectsapi/v0/projects` | 列出模拟项目 |
| `GET` | `/fakeapi/projectsapi/v0/projects/{projectid}` | 返回一个模拟项目 |
| `GET` | `/fakeapi/filesapi/v1/buckets` | 列出模拟 buckets |

## 身份验证

无。这些模拟端点无需身份验证。

## 示例

```bash
curl -X GET "https://api.hola.cloud/fakeapi/authapi/auth/me"
```

```json
{
  "email": "fulanez@gmail.com",
  "id": "user-1234",
  "nick": "fulanez"
}
```

```bash
curl -X GET "https://api.hola.cloud/fakeapi/projectsapi/v0/projects"
```

```json
[
  { "id": "project-00000000-0000-0000-0000-000000000001", "name": "Hello" }
]
```

```bash
curl -X GET "https://api.hola.cloud/fakeapi/lambdasapi/api/v0/lambdas"
```

```json
[
  { "id": "a663e1f1-e5cb-4fc2-b846-53f2cc7574c9", "method": "GET", "name": "Index", "path": "/" }
]
```

## 错误码

| 状态码 | 描述 |
|--------|------|
| 200 | 成功返回模拟数据 |
| 404 | `/fakeapi/` 下的路径无法识别 |
