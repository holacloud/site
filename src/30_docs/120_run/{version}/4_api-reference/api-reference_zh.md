# Run API 参考

Run 提供一个小型 Console API，以及 `/v2` 下的面向 push 的 Docker Registry v2 子集。

## Console API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/version` | 以纯文本返回服务版本 |
| `GET` | `/api/console?repository=` | 返回某个仓库的 Console 数据 |
| `POST` | `/api/console/start` | 使用 reference 或 digest 启动仓库 |
| `POST` | `/api/console/stop` | 停止仓库 |
| `POST` | `/api/console/rollback` | 回滚仓库到 reference 或 digest |
| `PUT` | `/api/console/env` | 保存仓库环境变量 |
| `PUT` | `/api/console/volumes` | 保存仓库卷配置 |

Console API 不使用 `container_id`，也不暴露 `/v1/run/deploy`、`/v1/run/push`、`/v1/run/exec`、`/api/console/run` 或类似 run/exec endpoints。

## Registry Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/v2/` | Registry API 检查 |
| `HEAD` | `/v2/*` | push 过程中检查 blob 或 manifest |
| `POST` | `/v2/*/blobs/uploads/` | 开始 blob 上传 |
| `PATCH` | `/v2/*/blobs/uploads/:uuid` | 上传 blob 数据 |
| `PUT` | `/v2/*/blobs/uploads/:uuid` | 使用 digest 完成 blob 上传 |
| `PUT` | `/v2/*/manifests/:reference` | 上传 manifest |

这是面向 push 的子集，不是完整 pull-capable registry API。
