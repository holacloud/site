# 流式插入

`insertStream` 和 `insertFullduplex` 是用于流式 JSON 输入的实验性插入端点。

## 认证

需要 `Api-Key` 和 `Api-Secret`，或者在数据库 owner 被允许时使用 Glue owner token。

## Insert Stream

`insertStream` 会接管 HTTP 连接，并使用 chunked 输出返回 `202 Accepted`。

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/events:insertStream HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/jsonl

{"id":"evt-1","type":"signup"}
{"id":"evt-2","type":"login"}
```

## Insert Full Duplex

`insertFullduplex` 启用 HTTP full duplex，并在读取请求体时把每个已插入文档编码为 JSON。

```http
POST /v1/databases/a1b2c3d4-e5f6-7890-abcd-ef1234567890/collections/events:insertFullduplex HTTP/1.1
Host: api.hola.cloud
Api-Key: 1abbe476-6ad6-4b97-9cca-6deb6ab2901d
Api-Secret: 4bda6d52-762b-4e5d-bed7-85614c13b8bf
Content-Type: application/jsonl

{"id":"evt-1","type":"signup"}
{"id":"evt-2","type":"login"}
```

两个端点的响应体都是包含已插入文档的 JSON Lines；`insertStream` 也可能返回逐文档错误字符串。
