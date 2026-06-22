<h1>Lambda</h1>

HolaCloud Lambda 用于发布在 HolaCloud 内运行的小型 HTTP 处理器。一个 lambda 会保存源码、路由元数据、所有者和项目关联，并通过 Lambda API 或公开运行路由返回响应。

## 主要功能

### 支持的语言
Lambda 接受这些 `language` 值：`javascript`、`static-html`、`static-css` 和 `static-js`。

### HTTP 路由
每个 lambda 可以保存 `method` 和 `path`。这些字段用于描述 lambda 如何通过 mux router 或应用路由层访问。

### 公开和管理调用
使用经过认证的管理端点创建、查看、更新、删除和运行 lambdas。当 lambda 需要被 webhook、浏览器或其他外部客户端调用时，使用公开运行路由和 mux 路由。

## 使用场景

- **HTTP API**：构建由 HolaCloud 服务支持的小型请求处理器。
- **Webhooks**：通过公开的 `/run/{lambda_id}` URL 接收第三方事件。
- **静态响应**：使用静态语言模式提供 HTML、CSS 或客户端 JavaScript 片段。
- **Mux 路由**：通过 `/mux/{owner_id}/*` 按所有者路由 path。

## Lambda 字段

Lambda 记录使用这些字段：`id`、`created_timestamp`、`owner`、`project_id`、`name`、`language`、`code`、`method` 和 `path`。
