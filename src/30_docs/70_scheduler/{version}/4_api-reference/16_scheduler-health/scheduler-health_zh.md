# 调度器健康检查

返回调度器的健康状态。

## 身份验证

此端点为公开，无需身份验证。

## 路径参数

| 参数 | 类型 | 描述 |
|------|------|------|
| id | string | 调度器的唯一标识符 |

## 请求示例

```bash
curl "https://api.hola.cloud/schedulers/sched-a1b2c3d4-e5f6-7890-abcd-ef1234567890/health"
```

## 响应示例

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

```json
{
  "status": "ok",
  "uptime_seconds": 12345
}
```

## 错误代码

| 状态 | 代码 | 描述 |
|------|------|------|
| 404 | not_found | 未找到调度器 |
| 500 | internal_error | 服务器内部错误 |
