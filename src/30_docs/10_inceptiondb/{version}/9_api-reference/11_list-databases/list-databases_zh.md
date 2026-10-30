# 列出数据库

`GET /v1/databases` · 仅限控制台

仅限控制台：此端点需要已登录控制台的用户会话。`Api-Key` 和 `Api-Secret` 对它无效，也没有受支持的方式从你自己的代码调用它。这里记录它是为了说明控制台做了什么。[控制台](https://inceptiondb.hola.cloud/)会替你完成，位置：数据库选择器。

返回你拥有的数据库。

```json
[{"id":"df355e63-4eb4-45cd-bcef-790755609208","name":"my-app","owners_length":1}]
```
