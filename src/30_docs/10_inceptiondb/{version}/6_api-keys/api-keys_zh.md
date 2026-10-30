# API key

你的应用通过 **API key** 访问数据库：一个密钥和一个 secret，通过两个请求头发送。数据库、API key 和所有者都在 [InceptionDB 控制台](https://inceptiondb.hola.cloud/) 中管理，使用你的 hola.cloud 账户登录。

## 创建数据库

在控制台中打开数据库选择器，选择 **New database**。大约一秒即可就绪。它的 ID 显示在概览页面上：这就是 `INCEPTIONDB_DATABASE` 的值。

## 创建 API key

1. 选择数据库并打开 **API keys & settings**。
2. 点击 **New API key**，以将使用它的应用命名，例如 `backend-production`。
3. 复制 secret。它**只显示一次**：InceptionDB 只保存其哈希值。

控制台会给出可以直接使用的凭据，例如作为环境变量：

```bash
export INCEPTIONDB_DATABASE="your-database-id"
export INCEPTIONDB_API_KEY="your-api-key"
export INCEPTIONDB_API_SECRET="your-api-secret"
```

一个 API key 可以访问其数据库中的所有集合，且只能访问这个数据库。

## 使用

在每个请求的 `Api-Key` 和 `Api-Secret` 请求头中发送密钥和 secret：

```bash
curl "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

```json
[{"defaults":{"id":"uuid()"},"indexes":0,"name":"customers","total":2}]
```

密钥或 secret 错误时返回 `403 Forbidden`。

## 保证安全

- InceptionDB 是服务器到服务器的 API。不要在浏览器或移动应用中运行的代码里调用它：任何人都能在那里读到 secret。如果网页需要数据，让它调用你自己的后端，由后端保存 secret 并调用 InceptionDB。
- 把 secret 保存在服务器上，放在环境变量或密钥管理器中。不要放在代码、代码仓库、网页或移动应用中：任何看到它的人都能读取和修改你的数据。
- 每个应用或环境使用一个 API key。这样撤销其中一个不会影响其他的。
- 如果 secret 泄露，创建一个新密钥并部署，然后在 **API keys & settings** 中**撤销**旧密钥。被撤销的密钥立即失效。

## 所有者

所有者是通过控制台管理数据库的 hola.cloud 用户：管理其集合、API key 和其他所有者。在 **API keys & settings** 中添加或移除。数据库始终至少保留一个所有者，且你不能移除自己。

## 账户操作：仅限控制台

创建、重命名和删除数据库，以及管理 API key 和所有者，都在控制台中完成。[API 参考](/zh/docs/inceptiondb/v2/api-reference/) 列出了这些端点，用于说明控制台做了什么，但它们不用于自动化：

- `Api-Key` 和 `Api-Secret` 对它们无效。API key 只能用于其数据库中的集合。
- 它们需要已登录控制台的用户会话，在控制台之外没有受支持的获取方式。

## 查看控制台操作的代码

控制台中的每个操作都是一个 HTTP 请求，控制台会显示如何在你的代码中发出它。点击侧边菜单底部带有 `{ }` 图标的 **API calls**。它会列出最近的请求；选择其中一个，即可看到 curl、Go、Python 等语言的版本，可直接复制。每个新请求还会在控制台右边缘的窄栏中显示为一个彩色圆点；点击圆点会打开同样的视图。
