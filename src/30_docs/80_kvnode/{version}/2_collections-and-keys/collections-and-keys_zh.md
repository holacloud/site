# 集合与键

KVNode 将数据组织到**集合**中，每个集合包含**键值对**。本文档详细介绍集合管理和键操作。

## 集合管理

### 创建集合

```bash
curl -X POST "https://api.hola.cloud/v1/collections" \
  -H "apikey: 你的API密钥" \
  -H "secret: 你的API密钥secret" \
  -d '{"name": "my-collection"}'
```

响应：

```json
{"ok":true,"collection":"my-collection"}
```

集合会立即创建。如果集合已存在，则该操作是幂等的。

### 列出集合

```bash
curl "https://api.hola.cloud/v1/collections" \
  -H "apikey: 你的API密钥" \
  -H "secret: 你的API密钥secret"
```

响应：

```json
{"collections":["my-collection","config","sessions"]}
```

### 删除集合

```bash
curl -X DELETE "https://api.hola.cloud/v1/collections/my-collection" \
  -H "apikey: 你的API密钥" \
  -H "secret: 你的API密钥secret"
```

响应：

```json
{"ok":true,"collection":"my-collection"}
```

删除集合将移除其中存储的所有键。此操作不可撤销。

## 键操作

KVNode 的键是字符串，可以包含任何 UTF-8 字符。值必须是有效的 JSON。

### 设置键

```bash
curl -X POST "https://api.hola.cloud/v1/collections/my-collection/keys/settings:theme" \
  -H "apikey: 你的API密钥" \
  -H "secret: 你的API密钥secret" \
  -d '{"value": {"mode": "dark", "fontSize": 14}}'
```

响应包含序列号和版本号：

```json
{"ok":true,"seq":3,"version":1}
```

每次成功写入都会返回单调递增的 `seq`（全局序列号）和 `version`（每个键的版本号），可用于乐观并发控制。

### 获取键

```bash
curl "https://api.hola.cloud/v1/collections/my-collection/keys/settings:theme" \
  -H "apikey: 你的API密钥" \
  -H "secret: 你的API密钥secret"
```

响应：

```json
{"key":"settings:theme","value":{"mode":"dark","fontSize":14},"version":1,"updatedAt":"2025-01-01T00:00:00Z"}
```

尝试获取不存在的键将返回 404 错误。

### 删除键

```bash
curl -X DELETE "https://api.hola.cloud/v1/collections/my-collection/keys/settings:theme" \
  -H "apikey: 你的API密钥" \
  -H "secret: 你的API密钥secret"
```

响应：

```json
{"ok":true,"seq":4,"version":2}
```

### 按前缀和限制列出键

```bash
curl "https://api.hola.cloud/v1/collections/my-collection/keys?prefix=settings:&limit=20" \
  -H "apikey: 你的API密钥" \
  -H "secret: 你的API密钥secret"
```

`prefix` 和 `limit` 都是可选的查询参数：
- `prefix` — 过滤以指定字符串开头的键
- `limit` — 限制返回的记录数量（默认：无限制）

响应：

```json
[
  {"key":"settings:theme","value":{"mode":"dark","fontSize":14},"version":1,"updatedAt":"2025-01-01T00:00:00Z"},
  {"key":"settings:locale","value":"en-US","version":1,"updatedAt":"2025-01-01T00:00:01Z"}
]
```

## 键命名约定

虽然 KVNode 接受任何 UTF-8 字符串作为键，但我们建议遵循以下约定：

- **使用冒号命名空间**：`app:module:key`（例如 `config:database:host`）
- **避免特殊字符**，以免与 URL 编码冲突（`/`、`?`、`#`）
- **保持键名可读**——有意义的名称便于调试
- **保持一致性**——选择命名模式并在所有集合中统一使用

## 值类型

所有值必须是有效的 JSON。KVNode 支持：

- **对象**：`{"user": "alice", "role": "admin"}`
- **数组**：`["item1", "item2"]`
- **字符串**：`"hello world"`
- **数字**：`42` 或 `3.14`
- **布尔值**：`true` 或 `false`
- **空值**：`null`

最大值大小由 WAL 后端配置决定。
