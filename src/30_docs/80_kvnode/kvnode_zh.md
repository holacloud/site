# Kvnode

KVNode 是一个复制式分布式键值存储系统，专为高可用性和低延迟访问而设计。它是 HolaCloud 生态系统的一部分，提供简单的 HTTP API 用于存储和检索 JSON 值。

## 主要特性

### 基于 WAL 的持久化
每次写入在应用到内存存储之前都会先记录到预写日志（WAL）中。这确保了持久性，并实现了节点间的无缝复制。

### 可插拔的 WAL 后端
KVNode 支持多种 WAL 后端——内存、Kafka、PostgreSQL、Redis 和 MongoDB——您可以选择最适合您基础设施的持久化层。

### 多语言 SDK
官方 SDK 支持 Go、Python、Java、JavaScript、Kotlin、PHP 和 Node.js，让 KVNode 能够轻松集成到任何技术栈中。

### 强一致性
根据 WAL 后端的不同，写入以同步或异步方式复制，单键操作具有线性化语义。

## 使用场景

- **配置存储**：以键值对的形式存储应用程序配置，并在节点间自动复制。
- **功能开关**：集中管理功能开关，实时传播变更。
- **服务发现**：以低延迟读取的方式注册和发现服务端点。
- **会话存储**：存储用户会话，支持快速查找和内置复制以实现容错。

## API 概述

KVNode 在 `https://api.hola.cloud` 提供 RESTful API。所有修改数据的端点都需要通过 API 密钥和密钥或 Glue 网关认证头进行内部认证。

| 方法 | 路径 | 描述 | 认证 |
|--------|------|------|------|
| GET | /healthz | 健康检查 | 公开 |
| GET | /readyz | 就绪检查（检查父节点连接） | 公开 |
| GET | /v1/status | 节点状态（集合、复制、运行时间） | 内部 |
| GET | /v1/metrics | 节点指标（写入、读取） | 内部 |
| GET | /v1/collections | 列出集合 | 内部 |
| POST | /v1/collections | 创建集合 | 内部 |
| DELETE | /v1/collections/{col} | 删除集合 | 内部 |
| GET | /v1/collections/{col}/keys | 列出键（限制、前缀） | 内部 |
| GET | /v1/collections/{col}/keys/* | 获取键值 | 内部 |
| POST | /v1/collections/{col}/keys/* | 设置键值 | 内部 |
| DELETE | /v1/collections/{col}/keys/* | 删除键 | 内部 |
| POST | /v1/replicate | 复制流（NDJSON） | 内部 |
