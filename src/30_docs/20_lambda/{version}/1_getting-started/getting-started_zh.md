<h1>Lambda快速入门</h1>

本指南将引导您创建第一个 Lambda 函数、调用它并检查执行结果。

## 前提条件

- 拥有 HolaCloud 账户及 API 凭证（Api-Key 和 Api-Secret）。
- 已安装 [curl](https://curl.se/)。

## 第一步：获取 API 凭证

登录 HolaCloud 控制台，进入 API 密钥管理页面。生成一个新的密钥对，您将获得一个 **Api-Key** 和一个 **Api-Secret**。请妥善保管这些值，它们用于验证所有管理请求。

## 第二步：创建 Lambda 函数

创建一个用 JavaScript 编写的 "hello-world" 函数：

```bash
curl -X POST "https://api.hola.cloud/api/v0/lambdas" \
  -H "Api-Key: 您的_API_KEY" \
  -H "Api-Secret: 您的_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "hello-world",
    "runtime": "js",
    "code": "export default (req) => { return { body: \"Hello, World!\" } }"
  }'
```

预期响应：

```json
{
  "id": "f3b2c1a0-1234-5678-9abc-def012345678",
  "name": "hello-world",
  "runtime": "js",
  "status": "active",
  "created_at": "2025-06-21T12:00:00Z"
}
```

保存返回的 `id`——调用和管理函数时需要使用。

## 第三步：列出 Lambda 函数

通过列出账户下所有 Lambda 函数来确认创建成功：

```bash
curl "https://api.hola.cloud/api/v0/lambdas" \
  -H "Api-Key: 您的_API_KEY" \
  -H "Api-Secret: 您的_API_SECRET"
```

## 第四步：运行函数

使用管理端点同步调用函数：

```bash
curl -X POST "https://api.hola.cloud/api/v0/run/您的函数ID" \
  -H "Api-Key: 您的_API_KEY" \
  -H "Api-Secret: 您的_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{}'
```

预期响应：

```json
{
  "body": "Hello, World!",
  "status_code": 200
}
```

您也可以使用公共端点（无需认证）调用函数：

```bash
curl -X POST "https://api.hola.cloud/run/您的函数ID" \
  -H "Content-Type: application/json" \
  -d '{}'
```

## 第五步：查看活跃调用

使用公共端点列出当前正在运行的调用：

```bash
curl "https://api.hola.cloud/ongoing"
```

## 下一步

- 了解如何更新函数代码和设置环境变量，请参阅[管理Lambda函数](../2_managing-functions/managing-functions_zh.md)。
- 探索包括 webhook 和 mux 路由器在内的调用模式，请参阅[调用Lambda函数](../3_invoking-functions/invoking-functions_zh.md)。
