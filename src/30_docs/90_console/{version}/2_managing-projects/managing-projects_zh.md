# 管理项目和服务

项目用于组合服务实例和路由配置。控制台的项目创建、计费、成员、导入/导出和 API 密钥端点当前不在本文档范围内。

## 切换项目

请使用控制台中的项目选择器。在演示/开发模式中，项目列表来自真实的 fake API：

```bash
curl -X GET "https://console.hola.cloud/fakeapi/projectsapi/v0/projects"
```

## 可用模拟服务

```bash
curl -X GET "https://console.hola.cloud/fakeapi/inceptionapi/v1/databases"
curl -X GET "https://console.hola.cloud/fakeapi/lambdasapi/api/v0/lambdas"
curl -X GET "https://console.hola.cloud/fakeapi/filesapi/v1/buckets"
```

## API 范围

生产项目 API 当前不在这些文档范围内。真实 API 密钥管理属于 serviceprojects 的 `/v0/apikeys`。
