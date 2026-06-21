# KVNode快速入门

本指南将带您了解 KVNode 的基本操作：健康检查、集合管理以及键值操作。

## 健康检查

验证节点是否正常运行：

```bash
curl "https://api.hola.cloud/healthz"
```

预期响应：

```json
{"ok":true,"node":"node-1","role":"primary"}
```

## 就绪检查

检查节点是否已准备好处理流量（检查副本的父节点连接）：

```bash
curl "https://api.hola.cloud/readyz"
```

就绪时的预期响应：

```json
{"ok":true,"node":"node-1","role":"primary","ready":true,"checks":{"wal_replayed":true,"parent_connected":true}}
```

## 创建集合

集合是键值对的容器。通过 POST 请求创建一个集合：

```bash
curl -X POST "https://api.hola.cloud/v1/collections" \
  -H "apikey: 你的API密钥" \
  -H "secret: 你的API密钥secret" \
  -d '{"name": "my-collection"}'
```

预期响应：

```json
{"ok":true,"collection":"my-collection"}
```

## 设置键值对

在集合中的某个键下存储 JSON 值：

```bash
curl -X POST "https://api.hola.cloud/v1/collections/my-collection/keys/user:alice" \
  -H "apikey: 你的API密钥" \
  -H "secret: 你的API密钥secret" \
  -d '{"value": {"user": "alice", "role": "admin"}}'
```

预期响应：

```json
{"ok":true,"seq":1,"version":1}
```

## 获取键

检索特定键的值：

```bash
curl "https://api.hola.cloud/v1/collections/my-collection/keys/user:alice" \
  -H "apikey: 你的API密钥" \
  -H "secret: 你的API密钥secret"
```

预期响应：

```json
{"key":"user:alice","value":{"user":"alice","role":"admin"},"version":1,"updatedAt":"2025-01-01T00:00:00Z"}
```

## 按前缀列出键

列出集合中的所有键，可按前缀过滤并限制数量：

```bash
curl "https://api.hola.cloud/v1/collections/my-collection/keys?prefix=user:&limit=10" \
  -H "apikey: 你的API密钥" \
  -H "secret: 你的API密钥secret"
```

预期响应：

```json
[{"key":"user:alice","value":{"user":"alice","role":"admin"},"version":1}]
```

## 删除键

从集合中删除一个键：

```bash
curl -X DELETE "https://api.hola.cloud/v1/collections/my-collection/keys/user:alice" \
  -H "apikey: 你的API密钥" \
  -H "secret: 你的API密钥secret"
```

预期响应：

```json
{"ok":true,"seq":2,"version":2}
```

## 总结

您已经完成了 KVNode 的基本操作：健康检查、集合创建以及键值对的完整 CRUD。接下来，可以探索复制配置和多语言 SDK 集成。
