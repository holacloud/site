# GettingStarted

本指南将引导您完成 HolaCloud Files 的基本操作：创建存储桶、上传文件、列出内容、下载和删除。

## 前提条件

在开始之前，您需要：

- 一个拥有 API 密钥和 API 密钥密码的 HolaCloud 账户
- 在您的机器上安装 `curl`

您的 API 凭证作为标头在每个请求中传递：

```
Api-Key: your-api-key
Api-Secret: your-api-secret
```

所有请求使用基础 URL `https://api.hola.cloud`。

## 第一步：创建存储桶

存储桶是文件的容器。每个存储桶必须具有全局唯一的名称。

```bash
curl -X POST "https://api.hola.cloud/v1/buckets" \
  -H "Api-Key: 您的API密钥" \
  -H "Api-Secret: 您的API密钥密码" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "我的第一个存储桶"
  }'
```

预期响应：

```json
{
  "id": "bkt_abc123",
  "name": "我的第一个存储桶",
  "createdAt": "2026-06-21T10:00:00Z",
  "size": 0,
  "fileCount": 0
}
```

## 第二步：上传文件

使用 PUT 将文本文件上传到存储桶。文件路径在 URL 中 `/files/` 之后指定。

```bash
echo "你好，HolaCloud Files！" > hello.txt

curl -X PUT "https://api.hola.cloud/v1/buckets/bkt_abc123/files/hello.txt" \
  -H "Api-Key: 您的API密钥" \
  -H "Api-Secret: 您的API密钥密码" \
  -H "Content-Type: text/plain" \
  --data-binary @hello.txt
```

预期响应：

```json
{
  "path": "hello.txt",
  "size": 24,
  "contentType": "text/plain",
  "uploadedAt": "2026-06-21T10:01:00Z",
  "etag": "\"d41d8cd98f00b204e9800998ecf8427e\""
}
```

## 第三步：列出存储桶内容

列出存储桶中的所有文件，可使用可选的前缀过滤器：

```bash
curl "https://api.hola.cloud/v1/buckets/bkt_abc123/list/*" \
  -H "Api-Key: 您的API密钥" \
  -H "Api-Secret: 您的API密钥密码"
```

预期响应：

```json
{
  "files": [
    {
      "path": "hello.txt",
      "size": 24,
      "contentType": "text/plain",
      "modifiedAt": "2026-06-21T10:01:00Z"
    }
  ],
  "prefix": "",
  "total": 1
}
```

## 第四步：下载文件

使用 GET 下载文件：

```bash
curl "https://api.hola.cloud/v1/buckets/bkt_abc123/files/hello.txt" \
  -H "Api-Key: 您的API密钥" \
  -H "Api-Secret: 您的API密钥密码" \
  -o downloaded_hello.txt

cat downloaded_hello.txt
```

输出：

```
你好，HolaCloud Files！
```

## 第五步：删除文件

使用 DELETE 删除文件：

```bash
curl -X DELETE "https://api.hola.cloud/v1/buckets/bkt_abc123/files/hello.txt" \
  -H "Api-Key: 您的API密钥" \
  -H "Api-Secret: 您的API密钥密码"
```

预期响应：HTTP 204 无内容。

## 第六步：删除存储桶

存储桶必须为空才能删除。由于我们已删除文件，此操作将成功：

```bash
curl -X DELETE "https://api.hola.cloud/v1/buckets/bkt_abc123" \
  -H "Api-Key: 您的API密钥" \
  -H "Api-Secret: 您的API密钥密码"
```

预期响应：HTTP 204 无内容。

## 总结

您已成功创建存储桶、上传文件、列出存储桶内容、下载文件并清理资源。现在您可以开始将 HolaCloud Files 集成到您的应用程序中。
