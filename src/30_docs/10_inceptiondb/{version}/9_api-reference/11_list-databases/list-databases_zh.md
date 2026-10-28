# 列出数据库

`GET /v1/databases` · 登录会话

这些账户端点需要已登录 hola.cloud 用户的会话，而不是 API key。[控制台](https://inceptiondb.hola.cloud/)会替你完成，位置：数据库选择器。

返回你拥有的数据库。

```json
[{"id":"df355e63-4eb4-45cd-bcef-790755609208","name":"my-app","owners_length":1}]
```
