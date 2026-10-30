# 删除数据库

`DELETE /v1/databases/{databaseId}` · 仅限控制台

仅限控制台：此端点需要已登录控制台的用户会话。`Api-Key` 和 `Api-Secret` 对它无效，也没有受支持的方式从你自己的代码调用它。这里记录它是为了说明控制台做了什么。[控制台](https://inceptiondb.hola.cloud/)会替你完成，位置：**API keys & settings → Delete database**。

删除数据库及其所有集合、文档和 API key，无法撤销。响应 `204 No Content`。
