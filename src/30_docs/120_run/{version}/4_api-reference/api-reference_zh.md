# Run API 参考

Run 是 HolaCloud 的容器执行服务。它提供用于容器管理的 REST API 和用于镜像存储的 Docker Registry v2 兼容接口。

## 端点

| 方法 | 路径 | 描述 |
|--------|------|-------------|
| `GET` | `/version` | 返回 Run 服务版本 |
| `GET` | `/api/console` | 控制台数据 — 仓库、镜像、状态 |
| `POST` | `/api/console/start` | 启动容器 |
| `POST` | `/api/console/stop` | 停止容器 |
| `POST` | `/api/console/rollback` | 回滚到之前的镜像 |
| `PUT` | `/api/console/env` | 保存环境变量 |
| `PUT` | `/api/console/volumes` | 保存卷配置 |
| `GET` | `/v2/` | Docker Registry API 版本检查 |
| `HEAD` | `/v2/*` | Docker Registry blob/清单检查 |
| `POST` | `/v2/*/blobs/uploads/` | Docker Registry blob 上传 |
| `PATCH` | `/v2/*/blobs/uploads/:uuid` | Docker Registry blob 上传分块 |
| `PUT` | `/v2/*/blobs/uploads/:uuid` | Docker Registry blob 上传完成 |
| `PUT` | `/v2/*/manifests/:tag` | Docker Registry 清单推送 |

## 身份验证

控制台 API 端点（`/api/console/*`）为公开端点。Docker Registry 端点需要使用 DockerAuth 机制进行基本身份验证。
