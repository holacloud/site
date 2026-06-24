# 服务元数据

这些端点暴露服务元数据和诊断信息。

## 路由树

```http
GET /doc HTTP/1.1
Host: api.hola.cloud
```

返回服务注册路由的文本树。

## API 根路径

```http
GET /v1 HTTP/1.1
Host: api.hola.cloud
```

返回 v1 欢迎响应。

## Release

```http
GET /release HTTP/1.1
Host: api.hola.cloud
```

返回已部署服务的 release 元数据。

## OpenAPI

```http
GET /openapi.json HTTP/1.1
Host: api.hola.cloud
```

返回 API 提供的 OpenAPI 文档。

## Debug 数据库

```http
GET /v1/debug/dbs HTTP/1.1
Host: api.hola.cloud
```

返回 debug 数据库列表。该端点用于诊断，不适合应用流量。
