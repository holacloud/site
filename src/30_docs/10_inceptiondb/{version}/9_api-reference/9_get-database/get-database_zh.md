# 获取数据库

`GET /v1/databases/{databaseId}` · 登录会话

这些账户端点需要已登录 hola.cloud 用户的会话，而不是 API key。[控制台](https://inceptiondb.hola.cloud/)会替你完成，位置：**API keys & settings**。

返回数据库及其所有者和 API key（从不包含密钥的 secret）。

```json
{"id":"df355e63-4eb4-45cd-bcef-790755609208","name":"my-app","creation_date":"2026-10-28T05:27:03.539038505Z","owners":["owner-1"],"api_keys":[{"name":"backend","key":"4fdd0267-fded-4d51-b1a2-e7674e02126e","creation_date":"2026-10-28T05:30:12Z"}]}
```
