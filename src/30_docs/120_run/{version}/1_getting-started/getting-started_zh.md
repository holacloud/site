# GettingStarted

Run 是一个容器执行服务，内置 Docker Registry v2 兼容 API。您可以使用它部署、运行和管理带有环境变量和卷的容器镜像。

## 身份认证

Run 使用 **DockerAuth** 配合基本认证（Basic Auth）进行注册表身份验证。使用您的 HolaCloud API 密钥和 API 密钥 Secret 作为用户名和密码。

### 登录到注册表

```bash
docker login https://registry.hola.cloud
```

系统将提示您输入 HolaCloud API 密钥（用户名）和 API 密钥 Secret（密码）。

## 构建和推送镜像

创建一个简单的应用程序和 `Dockerfile`：

```dockerfile
FROM alpine:latest
COPY app.sh /app.sh
RUN chmod +x /app.sh
CMD ["/app.sh"]
```

构建并标记镜像：

```bash
docker build -t registry.hola.cloud/my-project/my-app:v1 .
```

推送到注册表：

```bash
docker push registry.hola.cloud/my-project/my-app:v1
```

## 启动容器

通过 Console API 从已推送的镜像启动容器：

```bash
curl -X POST "https://api.hola.cloud/api/console/run" \
  -H "Api-Key: 你的API密钥" \
  -H "Api-Secret: 你的API密钥Secret" \
  -H "Content-Type: application/json" \
  -d '{
    "image": "registry.hola.cloud/my-project/my-app:v1",
    "env": {
      "LOG_LEVEL": "info",
      "DATABASE_URL": "postgres://..."
    },
    "volumes": [
      {
        "container_path": "/data",
        "volume_name": "my-data"
      }
    ]
  }'
```

预期响应：

```json
{
  "id": "c8a9f3b2-1234-5678-9abc-def012345678",
  "image": "registry.hola.cloud/my-project/my-app:v1",
  "status": "running",
  "created_at": "2025-06-21T12:00:00Z"
}
```

## 停止正在运行的容器

通过容器 ID 停止容器：

```bash
curl -X DELETE "https://api.hola.cloud/api/console/run/c8a9f3b2-1234-5678-9abc-def012345678" \
  -H "Api-Key: 你的API密钥" \
  -H "Api-Secret: 你的API密钥Secret"
```

## 回滚到之前的镜像版本

通过指定不同的标签来部署之前的镜像版本：

```bash
curl -X PUT "https://api.hola.cloud/api/console/run/c8a9f3b2-1234-5678-9abc-def012345678" \
  -H "Api-Key: 你的API密钥" \
  -H "Api-Secret: 你的API密钥Secret" \
  -H "Content-Type: application/json" \
  -d '{
    "image": "registry.hola.cloud/my-project/my-app:v0"
  }'
```

## 管理环境变量

更新运行中容器的环境变量：

```bash
curl -X PUT "https://api.hola.cloud/api/console/env" \
  -H "Api-Key: 你的API密钥" \
  -H "Api-Secret: 你的API密钥Secret" \
  -H "Content-Type: application/json" \
  -d '{
    "env": {
      "LOG_LEVEL": "debug",
      "DATABASE_URL": "postgres://..."
    }
  }'
```

## 管理卷

将卷挂载到容器：

```bash
curl -X PUT "https://api.hola.cloud/api/console/volumes" \
  -H "Api-Key: 你的API密钥" \
  -H "Api-Secret: 你的API密钥Secret" \
  -H "Content-Type: application/json" \
  -d '{
    "volumes": [
      {
        "container_path": "/data",
        "volume_name": "my-data",
        "read_only": false
      }
    ]
  }'
```

## 检查运行状态

查询所有容器的当前状态：

```bash
curl "https://api.hola.cloud/api/console/run" \
  -H "Api-Key: 你的API密钥" \
  -H "Api-Secret: 你的API密钥Secret"
```

预期响应：

```json
[
  {
    "id": "c8a9f3b2-1234-5678-9abc-def012345678",
    "image": "registry.hola.cloud/my-project/my-app:v1",
    "status": "running",
    "env": { "LOG_LEVEL": "info" },
    "volumes": [{ "container_path": "/data", "volume_name": "my-data" }],
    "created_at": "2025-06-21T12:00:00Z"
  }
]
```

## 完整工作流程示例

完整的端到端工作流程：

1. **登录**到注册表：
   ```bash
   docker login https://registry.hola.cloud
   ```

2. **构建**你的镜像：
   ```bash
   docker build -t registry.hola.cloud/my-project/my-app:v1 .
   ```

3. **推送**镜像：
   ```bash
   docker push registry.hola.cloud/my-project/my-app:v1
   ```

4. **启动**容器：
   ```bash
   curl -X POST "https://api.hola.cloud/api/console/run" \
     -H "Api-Key: 你的API密钥" \
     -H "Api-Secret: 你的API密钥Secret" \
     -H "Content-Type: application/json" \
     -d '{"image": "registry.hola.cloud/my-project/my-app:v1"}'
   ```

5. **验证**它正在运行：
   ```bash
   curl "https://api.hola.cloud/api/console/run" \
     -H "Api-Key: 你的API密钥" \
     -H "Api-Secret: 你的API密钥Secret"
   ```

6. **停止**容器：
   ```bash
   curl -X DELETE "https://api.hola.cloud/api/console/run/CONTAINER_ID" \
     -H "Api-Key: 你的API密钥" \
     -H "Api-Secret: 你的API密钥Secret"
   ```
