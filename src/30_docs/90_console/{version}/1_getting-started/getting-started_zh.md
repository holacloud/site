
# 控制台快速入门

控制台是所有 HolaCloud 服务的统一控制平面。它提供了一个集中的 Web 界面，用于管理数据库、函数、文件、配置、日志、队列等——全部通过一个仪表板完成。

![控制台仪表板](img/console.png)

## 登录

访问 [https://console.hola.cloud](https://console.hola.cloud) 并使用您的 HolaCloud 账户凭据登录。如果您使用的是演示或开发模式，模拟 API 端点将处于活动状态，让您无需真实后端即可探索所有功能。

```bash
# 示例：通过模拟 API 检查登录会话
curl -X GET "https://console.hola.cloud/api/v1/session" \
  -H "Authorization: Bearer <您的令牌>"
```

预期响应：

```json
{
  "user": "demo@holacloud.com",
  "project": "default",
  "role": "admin",
  "authenticated": true
}
```

## 导航仪表板

登录后，主仪表板将显示：

- **服务状态** — 每个 HolaCloud 服务的健康指示器（InceptionDB、Lambda、Files、Config、InstantLogs、Queues、Scheduler）。
- **最近活动** — 项目中执行的操作的时间顺序信息流。
- **关键指标** — 使用摘要，如请求计数、存储大小和函数调用次数。

使用左侧边栏在各个部分之间切换：仪表板、数据库、函数、文件、配置、日志、队列、调度程序和项目设置。

## 查看服务状态和最近活动

仪表板上的每张服务卡片都显示其当前状态（运行正常、降级或离线）。点击卡片可进入该服务的详细视图。最近活动信息流可按服务类型和时间范围进行筛选。

```bash
# 通过模拟 API 获取项目活动
curl -X GET "https://console.hola.cloud/api/v1/projects/default/activity" \
  -H "Authorization: Bearer <您的令牌>"
```

```json
{
  "activities": [
    { "type": "database.create", "service": "inceptiondb", "timestamp": "2026-06-20T10:30:00Z" },
    { "type": "function.invoke",  "service": "lambda",      "timestamp": "2026-06-20T10:28:00Z" },
    { "type": "file.upload",      "service": "files",       "timestamp": "2026-06-20T10:25:00Z" }
  ]
}
```

## 切换项目

如果您属于多个项目，请使用侧边栏顶部的项目选择器。整个界面将切换到所选项目的上下文，包括服务、账单和团队成员。

```bash
# 列出可用项目
curl -X GET "https://console.hola.cloud/api/v1/projects" \
  -H "Authorization: Bearer <您的令牌>"
```

```json
{
  "projects": [
    { "id": "default", "name": "Default Project", "role": "admin" },
    { "id": "staging", "name": "Staging Environment", "role": "developer" }
  ]
}
```

## 访问文档和支持

控制台包含一个内置帮助菜单（顶部栏中的问号图标），提供以下链接：

- 本文档网站
- API 参考
- 社区论坛
- 支持工单系统
- 服务状态页面

## 总结

控制台是所有 HolaCloud 服务的统一控制平面。无论您是在生产环境中运行，还是在演示模式下探索，控制台都能让您从一个地方全面了解和掌控您的云资源。
