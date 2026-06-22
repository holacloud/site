# 列出Schedulers

返回所有调度器的列表。

## 身份验证

此端点为公开，无需身份验证。

## 查询参数

| 参数 | 类型 | 描述 |
|------|------|------|
| search | string | 按显示名称过滤调度器（部分匹配） |
| q | string | 通用搜索查询 |

## 请求示例

```bash
curl "https://api.hola.cloud/schedulers?search=my"
```

## 响应示例

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "schedulers": [
    {
      "id": "sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890",
      "display_name": "my-scheduler",
      "ready": true,
      "scheduled": 5,
      "inflight": 0,
      "created_at": "2025-06-20T10:00:00Z",
      "updated_at": "2025-06-21T08:30:00Z"
    }
  ],
  "total": 1
}
```

## 错误代码

| 状态 | 代码 | 描述 |
|------|------|------|
| 500 | internal_error | 服务器内部错误 |
