# Files

HolaCloud Files 是一项兼容 S3 的对象存储服务，专为随时存储和检索任意数量的数据而设计。它提供了简单的 RESTful API 用于管理存储桶和文件，内置可扩展性、持久性和安全性。

## 主要特性

### 兼容 S3 的 API
HolaCloud Files 提供熟悉的兼容 S3 的 API，使迁移现有应用程序或与使用 S3 协议的工具集成变得容易。

### 可扩展性
Files 自动扩展以处理从几 GB 到 PB 级的任意数据量，无需手动配置或容量规划。

### 持久性和可用性
数据在多个可用区中冗余存储，确保 99.999999999% 的持久性和 99.99% 的可用性。

### 安全性
所有数据在静态时使用 AES-256 加密，在传输中使用 TLS 1.3 加密。通过具有细粒度权限的 API 密钥控制访问。

### 版本控制支持
该服务设计支持版本控制，允许您保留、检索和恢复存储桶中每个文件的每个版本。

## 使用场景

### 备份和灾难恢复
Files 为数据库备份、系统快照和灾难恢复归档提供可靠的目标。其持久性保证确保您的数据能够承受硬件故障和区域中断。

### Web 应用程序的静态资源
以低延迟和高吞吐量托管图像、CSS、JavaScript 和其他静态资源。Files 可与 HolaCloud CDN 无缝集成，实现全球内容交付。

### 用户生成的内容
通过可扩展且安全的存储层处理来自用户的文件上传 -- 个人资料图片、文档、视频等。

### 数据湖和分析
以 PB 级规模存储结构化和非结构化数据，用于分析工作负载。Files 与 HolaCloud Compute 和 HolaCloud Analytics 集成，用于数据处理流水线。

### 媒体存储和流媒体
为媒体应用程序存储音频、视频和图像文件。Files 支持范围请求，实现高效的部分内容交付。

## 与 HolaCloud 集成

HolaCloud Files 可与其它 HolaCloud 服务无缝协作：

- **HolaCloud Compute** -- 将 Files 卷挂载到您的计算实例以实现持久存储。
- **HolaCloud CDN** -- 通过边缘缓存全球交付文件内容。
- **HolaCloud Logs** -- 将访问日志直接归档到 Files 存储桶。
- **HolaCloud InceptionDB** -- 将数据库备份快照存储在 Files 中。

## API 参考

所有端点使用基础 URL `https://api.hola.cloud`，并需要通过 `Api-Key` 和 `Api-Secret` 标头进行身份验证。

| 方法 | 路径 | 描述 |
|--------|------|-------------|
| GET | /v1/buckets | 列出所有存储桶 |
| POST | /v1/buckets | 创建新存储桶 |
| GET | /v1/buckets/{id} | 获取存储桶详情 |
| PATCH | /v1/buckets/{id} | 修改存储桶元数据 |
| DELETE | /v1/buckets/{id} | 删除存储桶 |
| GET | /v1/buckets/{id}/list/* | 列出存储桶中的文件 |
| PUT | /v1/buckets/{id}/files/* | 上传文件 |
| GET | /v1/buckets/{id}/files/* | 下载文件 |
| DELETE | /v1/buckets/{id}/files/* | 删除文件 |
| HEAD | /v1/buckets/{id}/files/* | 获取文件元数据 |
