# 摄入帧事件

`POST /v1/ingest/events` 接受 logframe 编码事件。每个 frame 的 metadata 必须包含 `project_id`。

```json
{ "project_id": "project-1" }
```

响应：

```json
{
  "events": 2,
  "bytes": 128
}
```
