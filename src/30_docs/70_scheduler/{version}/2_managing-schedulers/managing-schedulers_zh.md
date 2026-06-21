# 管理调度器

本文档涵盖调度器的生命周期：创建（含元数据）、列表（含搜索和过滤）、更新配置、运行状况监控和删除。

## 创建调度器

使用显示名称创建新的调度器：

```bash
curl -X POST "https://api.hola.cloud/schedulers" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "display_name": "email-scheduler"
  }'
```

```http
POST /schedulers HTTP/1.1
Host: api.hola.cloud
X-API-Key: YOUR_API_KEY
Content-Type: application/json

{
  "display_name": "email-scheduler"
}
```

预期响应：

```json
{
  "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "display_name": "email-scheduler",
  "task_count": 0,
  "status": "active"
}
```

## 列出调度器

列出您账户中的所有调度器（公共）：

```bash
curl "https://api.hola.cloud/schedulers"
```

```http
GET /schedulers HTTP/1.1
Host: api.hola.cloud
```

预期响应：

```json
[
  {
    "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "display_name": "email-scheduler",
    "task_count": 5,
    "status": "active"
  },
  {
    "id": "sched-b2c3d4e5-f6a7-8901-bcde-f12345678901",
    "display_name": "report-generator",
    "task_count": 0,
    "status": "active"
  }
]
```

## 获取调度器详情

通过ID检索单个调度器：

```bash
curl "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890"
```

```http
GET /schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
```

预期响应：

```json
{
  "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "display_name": "email-scheduler",
  "task_count": 5,
  "status": "active"
}
```

## 更新调度器

更新显示名称或其他元数据：

```bash
curl -X PATCH "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "X-API-Key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "display_name": "email-scheduler-v2"
  }'
```

```http
PATCH /schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
X-API-Key: YOUR_API_KEY
Content-Type: application/json

{
  "display_name": "email-scheduler-v2"
}
```

预期响应：

```json
{
  "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "display_name": "email-scheduler-v2",
  "task_count": 5,
  "status": "active"
}
```

PUT可以与PATCH互换使用。

## 运行状况监控

随时检查调度器的运行状态：

```bash
curl "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/health"
```

```http
GET /schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/health HTTP/1.1
Host: api.hola.cloud
```

预期响应：

```json
{
  "status": "ok",
  "uptime_seconds": 123456
}
```

非`ok`状态表示调度器不可用或处于降级状态。

## 删除调度器

永久删除调度器及其所有任务：

```bash
curl -X DELETE "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "X-API-Key: YOUR_API_KEY"
```

```http
DELETE /schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890 HTTP/1.1
Host: api.hola.cloud
X-API-Key: YOUR_API_KEY
```

预期响应：HTTP `204 No Content`。

删除后，调度器ID将不再有效，所有待处理任务将被丢弃。
