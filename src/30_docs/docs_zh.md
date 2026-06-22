# 文档

欢迎使用 HolaCloud 文档。在这里，你可以找到平台服务的指南和 API 参考。

## 快速入门

如果你刚接触 HolaCloud，建议从[控制台](/zh/docs/console/)和各服务指南开始。部分服务仍在演进，因此 API 参考应结合当前实现说明一起阅读。

每项服务都包含入门指南、概念文档和 API 参考，涵盖 endpoint 详情、请求示例和错误码。

## 服务

### [InceptionDB](/zh/docs/inceptiondb/)
NoSQL 文档数据库。使用动态结构、集合和索引存储并查询 JSON 文档。

### [Lambda](/zh/docs/lambda/)
函数执行平台。将 JavaScript 函数和静态资源部署到 HTTP endpoint 后面。

### [Files](/zh/docs/files/)
文件 API。按 bucket 组织文件，并支持上传、下载、列表、查看和删除。

### [Config](/zh/docs/config/)
配置 API。存储和更新用户级 JSON 配置条目。

### [InstantLogs](/zh/docs/logs/)
实时日志采集与流式传输。接收 raw 或 framed logs，使用正则过滤并流式读取事件。

### [Tailon](/zh/docs/queues/)
消息队列 API。创建队列，写入 JSON 消息，并从命名队列读取换行分隔的 JSON。

### [Scheduler](/zh/docs/scheduler/)
延迟任务队列。入队一次性延迟任务，使用 lease 预留任务，并流式读取 scheduler snapshots。

### [KVNode](/zh/docs/kvnode/)
键值 API，包含 collections、keys、节点状态、指标和 root/replica 同步。

### [控制台](/zh/docs/console/)
HolaCloud 服务的 Web console surface，包括当前 UI 使用的 fake API endpoints。

### [Glue2](/zh/docs/glue2/)
基于 host 的 gateway 和服务 glue 层，提供 virtual hosts、认证上下文注入、状态和流量统计。

### [Holamail](/zh/docs/holamail/)
基础 SMTP 测试服务器。接收并记录 SMTP 消息，用于测试和开发。

### [Run](/zh/docs/run/)
容器 runtime 控制，包含面向 push 的 Docker Registry 子集。启动、停止和回滚 repository runtimes。

## API 参考

每项服务都包含 API 参考章节：

- 端点描述与 HTTP 方法
- 认证要求
- JSON 格式的请求和响应示例
- 错误码表
- 每个端点的 curl 示例

浏览侧边栏以访问任意服务，或使用顶部导航栏中的产品链接。
