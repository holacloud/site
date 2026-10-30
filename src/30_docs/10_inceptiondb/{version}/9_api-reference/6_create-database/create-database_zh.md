# 创建数据库

`POST /v1/databases` · 仅限控制台

仅限控制台：此端点需要已登录控制台的用户会话。`Api-Key` 和 `Api-Secret` 对它无效，也没有受支持的方式从你自己的代码调用它。这里记录它是为了说明控制台做了什么。[控制台](https://inceptiondb.hola.cloud/)会替你完成，位置：**数据库选择器 → New database**。

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
