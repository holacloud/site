# 配置

Config 是 Hola.Cloud 生态系统中的小型 JSON 配置 API。它存储配置条目，并允许已认证用户读取或更新运行时配置。

## 关键优势

### 基于条目的配置
v0 API 将配置存储为包含 `id` 和 `entries` 的 Thing。v1 用户 API 将当前用户配置返回为 `entries` 映射。

### 基于 JSON
所有配置都存储为 JSON 文档，易于读写，并能与任何编程语言或工具链集成。模式灵活，您可以存储任何有效的 JSON 结构。

### 简单更新
更新会替换或修补 API 接受的条目数据。

### 易于集成
Config 可与 Hola.Cloud 生态系统中的其他服务（如 InceptionDB、Lambda 和 InstantLogs）无缝集成。在服务之间集中配置可减少重复工作，并保持基础设施的一致性。

### 安全访问
该服务提供两个 API 接口。v0 API（`/v0/configs`）管理配置 Thing。v1 用户 API（`/v1/config`）需要 `X-Glue-Authentication`，用于读取或更新已认证用户的运行时配置。

## API 概览

| 方法 | 路径 | 描述 | 认证 |
|--------|------|-------------|------|
| GET | `/v0/configs` | 列出所有配置（管理） | 公开 |
| POST | `/v0/configs` | 创建新配置（管理） | 公开 |
| GET | `/v0/configs/{id}` | 根据 ID 获取配置（管理） | 公开 |
| DELETE | `/v0/configs/{id}` | 删除配置（管理） | 公开 |
| PATCH | `/v0/configs/{id}` | 部分更新配置（管理） | 公开 |
| GET | `/v1/config` | 获取当前用户配置 `entries` 映射 | Glue 认证 |
| PATCH | `/v1/config` | 更新当前用户配置 `entries` 映射 | Glue 认证 |

基础 URL：`https://api.hola.cloud`

## 最佳使用场景

### 运行时配置
在启动时或运行期间从 `/v1/config` 获取已认证用户的配置。

### 功能开关
将功能开关存储为配置项，并在运行时更新而无需修改代码。
