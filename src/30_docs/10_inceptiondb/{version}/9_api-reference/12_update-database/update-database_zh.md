# 重命名数据库

`PATCH /v1/databases/{databaseId}` · 登录会话

这些账户端点需要已登录 hola.cloud 用户的会话，而不是 API key。[控制台](https://inceptiondb.hola.cloud/)会替你完成，位置：**API keys & settings → General**。

| 请求体字段 | 类型 | 含义 |
|---|---|---|
| `name` | string | 新名称。id 不会改变。 |

请求体：

```json
{"name": "my-app-prod"}
```

响应 `200 OK`：

```json
{"id":"df355e63-4eb4-45cd-bcef-790755609208","name":"my-app-prod","creation_date":"2026-10-28T05:27:03.539038505Z","owners":["owner-1"],"api_keys":[]}
```
