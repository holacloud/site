# Replication

KVNode 通过**主从复制**模型实现高可用性。写入主（领导者）节点的更改通过基于 NDJSON 的复制协议实时传输到副本（跟随者）节点。

## 复制工作原理

1. 客户端向集群中的任意节点写入键值对。
2. 接收节点将写入持久化到其 WAL 中，并应用到内存存储。
3. 如果节点配置了**父节点**，则通过 `/v1/replicate` 端点将写入向上游转发。
4. 如果节点有**子节点**，则通过复制流将写入广播到所有已连接的副本。
5. 每个副本将更改应用到自己的存储中，保持最终一致的数据副本。

## 复制协议

复制端点（`POST /v1/replicate`）使用 **NDJSON**（换行符分隔的 JSON）作为传输格式。该协议支持以下命令类型：

| 命令 | 描述 |
|---------|------|
| `set` | 键被创建或更新 |
| `delete` | 键被删除 |
| `baseline_begin` | 完整集合快照开始 |
| `baseline_end` | 完整集合快照结束 |
| `ping` | 心跳保持连接 |

### 示例：复制流

```
{"type":"baseline_begin","collection":"my-collection","seq":0}
{"type":"set","collection":"my-collection","key":"user:alice","value":{"role":"admin"},"version":1,"timestamp":1700000000}
{"type":"baseline_end","collection":"my-collection","seq":1}
{"type":"set","collection":"my-collection","key":"user:bob","value":{"role":"user"},"version":1,"timestamp":1700000001}
{"type":"ping","timestamp":1700000010}
```

## 节点状态

`/v1/status` 端点暴露复制拓扑信息：

```bash
curl "https://api.hola.cloud/v1/status" \
  -H "apikey: 你的API密钥" \
  -H "secret: 你的API密钥secret"
```

响应：

```json
{
  "node": "node-1",
  "role": "primary",
  "parent": "",
  "collections": {
    "my-collection": {
      "keys": 42,
      "lastSeq": 128,
      "walSizeBytes": 65536
    }
  },
  "children": 2,
  "uptimeSeconds": 86400,
  "replication": {
    "enabled": true,
    "parent_connected": false
  }
}
```

## 节点指标

通过 `/v1/metrics` 端点监控性能：

```bash
curl "https://api.hola.cloud/v1/metrics" \
  -H "apikey: 你的API密钥" \
  -H "secret: 你的API密钥secret"
```

响应：

```json
{
  "writes_total": 15000,
  "reads_total": 82000,
  "replication_commands_sent_total": 15000,
  "replication_commands_received_total": 0,
  "children_connected": 2,
  "parent_connected": false
}
```

## 就绪检查

`/readyz` 端点验证节点是否已准备好处理流量。对于副本，这意味着父节点连接必须已建立且 WAL 已完全重放。对于主节点，启动后立即可用。

```bash
curl "https://api.hola.cloud/readyz"
```

未就绪时的响应：

```json
{"ok":false,"node":"node-2","role":"replica","ready":false,"checks":{"wal_replayed":true,"parent_connected":false}}
```

## 可插拔的 WAL 后端

KVNode 支持多种 WAL 后端，可在启动时选择：

| 后端 | 描述 | 使用场景 |
|---------|-------------|----------|
| **内存** | 进程内缓冲区（非持久化） | 开发、测试、临时缓存 |
| **Kafka** | 通过 Apache Kafka 的持久化日志 | 高吞吐量生产部署 |
| **PostgreSQL** | WAL 存储在 PostgreSQL 表中 | 已在用 Postgres 的环境 |
| **Redis** | WAL 存储在 Redis 流中 | 低延迟、以 Redis 为中心的架构 |
| **MongoDB** | WAL 存储在 MongoDB 集合中 | 以 MongoDB 为中心的部署 |

每个后端实现相同的 WAL 接口，因此在它们之间切换无需修改应用程序代码。
