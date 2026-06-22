# SDK 参考

KVNode 在 `kvnode/sdk/python/configservice.py` 中包含本地 Python configservice SDK。它使用同一个 `/v1/collections...` REST API 和复制流。

## 包

| 语言 | 包 | 说明 |
|------|----|------|
| Python | `ConfigServiceClient` from `kvnode/sdk/python/configservice.py` | 仓库内本地包，不是已发布的外部包 |

## 连接

```python
from configservice import ConfigServiceClient

client = ConfigServiceClient(
    base_url="https://api.hola.cloud",
    mode="http",
    node="sdk-python",
)
```

## 操作

```python
out = client.set("my-collection", "user:alice", {"role": "admin"})
# {"ok": true, "seq": 1, "version": 1}

entry = client.get("my-collection", "user:alice")
keys = client.list_keys("my-collection", prefix="user:", limit=100)
# {"items": [...], "next": "..."}

out = client.delete("my-collection", "user:alice")
# {"ok": true, "seq": 2, "version": 2}
```

## 同步

`sync` 读取 `POST /v1/replicate` NDJSON，并应用大写命令，例如 `SET`、`DELETE`、`BASELINE_BEGIN` 和 `BASELINE_END`。

## 认证

服务只检查是否存在 `X-Glue-Authentication`，或同时存在 `apikey` 和 `secret`。缺少认证 header 时返回 `403 forbidden`。
