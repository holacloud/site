# 删除数据库

`DELETE /v1/databases/{databaseId}` · 登录会话

这些账户端点需要已登录 hola.cloud 用户的会话，而不是 API key。[控制台](https://inceptiondb.hola.cloud/)会替你完成，位置：**API keys & settings → Delete database**。

删除数据库及其所有集合、文档和 API key，无法撤销。响应 `204 No Content`。
