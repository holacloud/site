# 创建数据库

`POST /v1/databases` · 登录会话

这些账户端点需要已登录 hola.cloud 用户的会话，而不是 API key。[控制台](https://inceptiondb.hola.cloud/)会替你完成，位置：**数据库选择器 → New database**。

| 请求体字段 | 类型 | 含义 |
|---|---|---|
| `name` | string | 数据库名称。 |

请求体：

```json
{"name": "my-app"}
```

响应 `201 Created`。`id` 就是 `INCEPTIONDB_DATABASE` 的值：

```json
{"id":"df355e63-4eb4-45cd-bcef-790755609208","name":"my-app","creation_date":"2026-10-28T05:27:03.539038505Z"}
```
