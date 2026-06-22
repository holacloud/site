# Run

Run 是 HolaCloud 的容器镜像和 Console 控制服务。它提供面向 push 的 Docker Registry v2 子集，以及用于查看仓库和修改运行配置的小型 Console API。

## 功能

### 面向 push 的 Registry v2 子集

Run 支持在 `/v2` 下上传 blob 和 manifest 所需的 endpoints。它不是完整 Docker Registry 实现，也不应被描述为通用 pull registry。

```bash
docker login run.hola.cloud
docker build -t run.hola.cloud/my-project/my-app:latest .
docker push run.hola.cloud/my-project/my-app:latest
```

### Console API

Console API 使用 repository 以及 image reference 或 digest：

- `GET /version`
- `GET /api/console?repository=`
- `POST /api/console/start`
- `POST /api/console/stop`
- `POST /api/console/rollback`
- `PUT /api/console/env`
- `PUT /api/console/volumes`

不存在 `/v1/run/deploy`、`/api/console/run`、push/exec API，也没有 `container_id` 工作流。

## 快速开始

```bash
docker login run.hola.cloud
docker build -t run.hola.cloud/my-project/my-app:latest .
docker push run.hola.cloud/my-project/my-app:latest
curl -X POST "https://api.hola.cloud/api/console/start" \
  -H "Content-Type: application/json" \
  -d '{"repository": "my-project/my-app", "reference": "latest"}'
```
