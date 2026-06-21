<h1>管理Lambda函数</h1>

创建 Lambda 函数后，您可以通过 API 更新代码、配置环境变量、查看状态或删除函数。

## 函数结构

每个 Lambda 函数必须导出一个默认处理器，该处理器接收请求对象并返回响应。不同运行时的处理器签名略有不同：

### JavaScript (Node.js)

```javascript
export default (req) => {
  // req.body      - 解析后的请求体
  // req.headers   - 请求头对象
  // req.query     - 解析后的查询参数
  // req.method    - HTTP 方法

  return {
    body: "响应内容",
    status_code: 200,
    headers: { "X-Custom": "value" }
  };
};
```

### Go

```go
package main

type Request struct {
    Body    interface{} `json:"body"`
    Headers map[string]string `json:"headers"`
    Method  string `json:"method"`
}

type Response struct {
    Body    interface{} `json:"body"`
    StatusCode int `json:"status_code"`
}

func Handler(req Request) (Response, error) {
    return Response{
        Body:       "来自Go的问候",
        StatusCode: 200,
    }, nil
}
```

### Python

```python
def handler(req):
    # req["body"]      - 请求体
    # req["headers"]   - 请求头字典
    # req["method"]    - HTTP 方法
    return {
        "body": "来自Python的问候",
        "status_code": 200
    }
```

## 更新函数代码

使用 `PATCH /api/v0/lambdas/{id}` 修改已有函数。可以更新 `name`、`runtime`、`code` 和 `environment` 字段。

```bash
curl -X PATCH "https://api.hola.cloud/api/v0/lambdas/您的函数ID" \
  -H "Api-Key: 您的_API_KEY" \
  -H "Api-Secret: 您的_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "code": "export default (req) => { return { body: \"已更新!\" } }",
    "environment": {
      "DB_URL": "https://db.example.com",
      "LOG_LEVEL": "debug"
    }
  }'
```

## 设置环境变量

环境变量通过创建或更新函数时的 `environment` 字段以键值对对象的形式传入。它们在运行时注入到函数进程中，可通过各语言的标准机制访问（Node.js 使用 `process.env`，Go 使用 `os.Getenv`，Python 使用 `os.environ`）。

```bash
curl -X PATCH "https://api.hola.cloud/api/v0/lambdas/您的函数ID" \
  -H "Api-Key: 您的_API_KEY" \
  -H "Api-Secret: 您的_API_SECRET" \
  -H "Content-Type: application/json" \
  -d '{
    "environment": {
      "MY_VAR": "my-value"
    }
  }'
```

## 查看函数详情与状态

获取单个函数的完整信息：

```bash
curl "https://api.hola.cloud/api/v0/lambdas/您的函数ID" \
  -H "Api-Key: 您的_API_KEY" \
  -H "Api-Secret: 您的_API_SECRET"
```

预期响应：

```json
{
  "id": "f3b2c1a0-1234-5678-9abc-def012345678",
  "name": "hello-world",
  "runtime": "js",
  "status": "active",
  "environment": {
    "LOG_LEVEL": "debug"
  },
  "created_at": "2025-06-21T12:00:00Z",
  "updated_at": "2025-06-21T14:30:00Z"
}
```

`status` 字段可能为 `active`、`inactive` 或 `error`。

## 列出所有函数

```bash
curl "https://api.hola.cloud/api/v0/lambdas" \
  -H "Api-Key: 您的_API_KEY" \
  -H "Api-Secret: 您的_API_SECRET"
```

## 删除函数

永久删除函数及其所有关联数据：

```bash
curl -X DELETE "https://api.hola.cloud/api/v0/lambdas/您的函数ID" \
  -H "Api-Key: 您的_API_KEY" \
  -H "Api-Secret: 您的_API_SECRET"
```

删除成功返回 `200 OK`，无响应体。
