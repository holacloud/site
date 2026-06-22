# 运行

Run 是 HolaCloud 的容器执行服务。它通过兼容 Docker Registry v2 的接口管理 Docker 容器，让您能够在隔离环境中部署和运行应用程序。

## 功能特性

### 兼容 Docker Registry v2
Run 是完全兼容的 Docker Registry v2。使用标准 Docker 命令推送和拉取镜像——无需自定义工具。

```bash
docker login run.hola.cloud
docker build -t run.hola.cloud/my-project/my-app:latest .
docker push run.hola.cloud/my-project/my-app:latest
```

### 控制台管理
通过 HolaCloud 控制台管理容器、镜像和部署。启动、停止和重启容器，查看日志，监控资源使用情况——全部通过 Web 界面完成。

### 环境变量
使用环境变量配置容器。在部署时通过控制台或 API 设置环境变量，实现灵活的配置驱动部署。

```bash
curl -X POST "https://api.hola.cloud/v1/run/deploy" \
  -H "Authorization: Bearer your-token" \
  -d '{
    "image": "run.hola.cloud/my-project/my-app:latest",
    "env": {
      "DATABASE_URL": "postgres://...",
      "LOG_LEVEL": "debug"
    }
  }'
```

### 持久卷
为有状态工作负载的容器附加持久卷。卷由 HolaCloud 的分布式存储支持，可在容器重启后保留数据。

## 使用场景

### 完整应用部署
将 Web 应用程序、API 和微服务部署为 Docker 容器。Run 负责网络、健康检查和自动重启。

### 开发环境
为每个分支或开发人员创建隔离的开发环境。使用 Run 创建与生产环境匹配的临时容器。

### CI/CD 流水线
将 Run 集成到您的 CI/CD 工作流中。在构建过程中将镜像推送到注册表，然后自动部署到预发布或生产环境。

## 快速入门

1. **认证**到注册表：
   ```bash
   docker login run.hola.cloud
   ```

2. **构建并推送**镜像：
   ```bash
   docker build -t run.hola.cloud/my-project/my-app:latest .
   docker push run.hola.cloud/my-project/my-app:latest
   ```

3. **部署**：通过控制台或 API 部署：
   ```bash
   curl -X POST "https://api.hola.cloud/v1/run/deploy" \
     -H "Authorization: Bearer your-token" \
     -d '{"image": "run.hola.cloud/my-project/my-app:latest"}'
   ```

您的容器将启动，并在分配的端点上可用。
