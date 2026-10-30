# 获取数据库

`GET /v1/databases/{databaseId}` · 仅限控制台

仅限控制台：此端点需要已登录控制台的用户会话。`Api-Key` 和 `Api-Secret` 对它无效，也没有受支持的方式从你自己的代码调用它。这里记录它是为了说明控制台做了什么。[控制台](https://inceptiondb.hola.cloud/)会替你完成，位置：**API keys & settings**。

返回数据库及其所有者和 API key（从不包含密钥的 secret）。

```json
{"id":"df355e63-4eb4-45cd-bcef-790755609208","name":"my-app","creation_date":"2026-10-28T05:27:03.539038505Z","owners":["owner-1"],"api_keys":[{"name":"backend","key":"4fdd0267-fded-4d51-b1a2-e7674e02126e","creation_date":"2026-10-28T05:30:12Z"}]}
```
