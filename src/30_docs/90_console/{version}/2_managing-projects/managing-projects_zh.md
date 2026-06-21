
# 管理项目和服务

项目是 HolaCloud 中的顶级组织单元。它们将服务实例、团队成员、API 密钥和账单集中在一起。本指南介绍如何通过控制台创建和管理项目。

## 创建和切换项目

要创建新项目，请打开侧边栏中的项目选择器，点击**新建项目**。提供名称和可选描述。创建后，控制台会自动切换到新项目。

```bash
# 通过控制台 API 创建项目
curl -X POST "https://console.hola.cloud/api/v1/projects" \
  -H "Authorization: Bearer <您的令牌>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "我的生产应用",
    "description": "我的应用程序的生产环境"
  }'
```

```json
{
  "id": "prod-app-123",
  "name": "我的生产应用",
  "description": "我的应用程序的生产环境",
  "createdAt": "2026-06-20T12:00:00Z"
}
```

要切换项目，请点击侧边栏中的当前项目名称，然后从下拉菜单中选择其他项目。

```bash
# 列出所有可访问的项目
curl -X GET "https://console.hola.cloud/api/v1/projects" \
  -H "Authorization: Bearer <您的令牌>"
```

```json
{
  "projects": [
    { "id": "default", "name": "Default", "role": "admin" },
    { "id": "prod-app-123", "name": "我的生产应用", "role": "admin" }
  ]
}
```

## 管理每个项目的服务实例

每个项目都包含自己的 HolaCloud 服务实例。导航到相应的服务部分（例如数据库、函数、文件、配置、日志、队列、调度程序）来管理当前项目中的实例。

例如，列出当前项目中的 InceptionDB 数据库：

```bash
curl -X GET "https://console.hola.cloud/api/v1/projects/prod-app-123/databases" \
  -H "Authorization: Bearer <您的令牌>"
```

```json
{
  "databases": [
    { "id": "db-001", "name": "users-db", "engine": "inceptiondb", "status": "active" },
    { "id": "db-002", "name": "analytics-db", "engine": "inceptiondb", "status": "active" }
  ]
}
```

同样的模式适用于 Lambda 函数、Files 存储桶、Config 存储、InstantLogs 流、队列和调度程序作业。

## 每个项目的 API 密钥管理

API 密钥用于验证对 HolaCloud 服务的编程访问。在**项目设置 > API 密钥**中管理它们。

```bash
# 创建 API 密钥
curl -X POST "https://console.hola.cloud/api/v1/projects/prod-app-123/api-keys" \
  -H "Authorization: Bearer <您的令牌>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "CI/CD 密钥",
    "scopes": ["lambda:invoke", "files:read", "config:read"]
  }'
```

```json
{
  "id": "key-abc123",
  "name": "CI/CD 密钥",
  "key": "hc_api_abc123...",
  "scopes": ["lambda:invoke", "files:read", "config:read"],
  "createdAt": "2026-06-20T14:00:00Z"
}
```

您可以随时从同一设置页面撤销密钥。

## 查看项目账单和使用情况

**账单**部分显示当前使用指标、发票历史记录以及按服务划分的成本明细。

```bash
curl -X GET "https://console.hola.cloud/api/v1/projects/prod-app-123/billing" \
  -H "Authorization: Bearer <您的令牌>"
```

```json
{
  "currentMonth": {
    "period": "2026-06",
    "total": 125.40,
    "services": {
      "inceptiondb": 45.00,
      "lambda": 60.00,
      "files": 15.40,
      "config": 5.00
    }
  },
  "invoices": [
    { "id": "inv-001", "period": "2026-05", "total": 110.20, "status": "paid" }
  ]
}
```

## 团队成员管理

从**项目设置 > 团队**邀请团队成员。分配角色：管理员、开发者或查看者。

```bash
curl -X POST "https://console.hola.cloud/api/v1/projects/prod-app-123/team" \
  -H "Authorization: Bearer <您的令牌>" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "同事@公司.com",
    "role": "developer"
  }'
```

```json
{
  "id": "team-456",
  "email": "同事@公司.com",
  "role": "developer",
  "status": "pending"
}
```

## 以 JSON 格式导入/导出配置

将您的完整项目配置（服务设置、API 密钥元数据、团队成员）导出为单个 JSON 文件。导航到**项目设置 > 导入/导出**。

```bash
# 导出配置
curl -X GET "https://console.hola.cloud/api/v1/projects/prod-app-123/export" \
  -H "Authorization: Bearer <您的令牌>"
```

```json
{
  "project": { "name": "我的生产应用" },
  "services": {
    "inceptiondb": { "databases": [...] },
    "lambda": { "functions": [...] },
    "files": { "buckets": [...] },
    "config": { "entries": [...] }
  },
  "team": [...]
}
```

要导入，请通过控制台 UI 或 API 上传相同格式的 JSON 文件：

```bash
curl -X POST "https://console.hola.cloud/api/v1/projects/prod-app-123/import" \
  -H "Authorization: Bearer <您的令牌>" \
  -H "Content-Type: application/json" \
  -d @project-config.json
```

```json
{
  "status": "imported",
  "servicesUpdated": ["inceptiondb", "lambda", "files", "config"]
}
```

## 总结

HolaCloud 中的项目为您的应用程序提供了隔离的环境。通过控制台，您可以创建和切换项目、管理服务实例、API 密钥、账单、团队成员，以及导入/导出配置——全部通过一个界面完成。
