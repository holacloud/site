# 快速开始

Run 提供面向 push 的 Docker Registry v2 子集，以及用于仓库状态、start、stop、rollback、环境变量和卷的 Console API。

## 认证

Registry endpoints 使用通过 `docker login` 配置的 Basic Auth。

```bash
docker login run.hola.cloud
```

## 构建并推送镜像

```bash
docker build -t run.hola.cloud/my-project/my-app:v1 .
docker push run.hola.cloud/my-project/my-app:v1
```

## 读取 Console 数据

```bash
curl "https://api.hola.cloud/api/console?repository=my-project/my-app"
```

## 启动仓库

```bash
curl -X POST "https://api.hola.cloud/api/console/start" \
  -H "Content-Type: application/json" \
  -d '{
    "repository": "my-project/my-app",
    "reference": "v1"
  }'
```

也可以使用 digest：

```json
{
  "repository": "my-project/my-app",
  "digest": "sha256:a1b2c3d4..."
}
```

## 停止仓库

```bash
curl -X POST "https://api.hola.cloud/api/console/stop" \
  -H "Content-Type: application/json" \
  -d '{"repository": "my-project/my-app"}'
```

## 回滚

```bash
curl -X POST "https://api.hola.cloud/api/console/rollback" \
  -H "Content-Type: application/json" \
  -d '{"repository": "my-project/my-app", "reference": "v0"}'
```

## 环境变量

```bash
curl -X PUT "https://api.hola.cloud/api/console/env" \
  -H "Content-Type: application/json" \
  -d '{
    "repository": "my-project/my-app",
    "env": [
      {"key": "LOG_LEVEL", "desired_value": "debug"}
    ]
  }'
```

## 卷

```bash
curl -X PUT "https://api.hola.cloud/api/console/volumes" \
  -H "Content-Type: application/json" \
  -d '{
    "repository": "my-project/my-app",
    "volumes": [
      {"name": "my-data", "target": "/data", "mode": "rw"}
    ]
  }'
```
