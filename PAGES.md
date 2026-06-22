# Site Pages Review

Scope: source content under `src/*`. Review English pages first. Spanish and Chinese pages are synchronized only after the English source page is corrected.

Status columns:
- `Audited`: checked against the service implementation or site generator.
- `Fixed EN`: English source corrected when the audit found wrong or invented content.
- `Synced ES/ZH`: Spanish and Chinese sources updated from the corrected English page.

## General Pages

| Page | Source | Audited | Fixed EN | Synced ES/ZH | Notes |
|---|---|---:|---:|---:|---|
| Home | `src/home_en.html` | yes | yes | partial | Removed unsupported SLA/support/regions/compliance/AI claims in EN/ES and false metrics in ZH. |
| Pricing | `src/20_pricing/pricing_en.html` | yes | yes | yes | EN/ES now state pricing is not published yet; no ZH source exists. |
| Docs index | `src/30_docs/docs_en.md` | yes | yes | yes | Service descriptions aligned with audited behavior in EN/ES/ZH. |

## Service Landings

| Service | Source | Audited | Fixed EN | Synced ES/ZH | Notes |
|---|---|---:|---:|---:|---|
| InceptionDB | `src/50_inceptiondb/inceptiondb_en.html` | yes | yes | yes | Corrected to JSON documents, collections, action endpoints, and real auth split. |
| InceptionDB Introduction | `src/50_inceptiondb/10_introduction/introduction_en.html` | yes | yes | n/a | Added basic implementation-backed introduction. |
| Lambda | `src/51_lambda/lambda_en.html` | yes | yes | yes | Corrected to JavaScript/static assets and real `/api/v0/lambdas` model. |
| Files | `src/52_files/files_en.html` | yes | yes | yes | Corrected to buckets/files REST model and real auth/header shape. |
| Config | `src/53_config/config_en.html` | yes | yes | yes | Corrected to v1 user config and v0 entries model. |
| Logs | `src/54_logs/logs_en.html` | yes | yes | yes | Corrected landing auth and ingest example; feature copy still needs deeper docs alignment. |
| Queues | `src/55_queues/queues_en.html` | yes | yes | yes | Corrected to simple Tailon queue endpoints. |
| Scheduler | `src/56_scheduler/scheduler_en.html` | yes | yes | yes | Corrected to delayed tasks, leases, snapshots, and real endpoints. |
| KVNode | `src/57_kvnode/kvnode_en.html` | yes | yes | yes | Corrected to collections/keys API; removed Raft/TTL claims. |
| Console | `src/58_console/console_en.html` | yes | yes | yes | Corrected landing to current console UI/fake API/version endpoint. |
| Glue2 | `src/59_glue2/glue2_en.html` | yes | yes | yes | Corrected quick start to real gateway diagnostics; feature copy still needs deeper docs alignment. |
| HolaMail | `src/60_holamail/holamail_en.html` | yes | yes | yes | Corrected to basic SMTP test server and removed HTTP send API. |
| Run | `src/61_run/run_en.html` | yes | yes | yes | Corrected to repository console API and push-oriented registry subset. |

## Documentation Groups

| Service | Source Group | Audited | Fixed EN | Synced ES/ZH | Notes |
|---|---|---:|---:|---:|---|
| InceptionDB | `src/30_docs/10_inceptiondb/**/*_en.md` | yes | yes | yes | Corrected auth split, action endpoints, DB/collection response shapes, and removed invented capabilities. |
| Lambda | `src/30_docs/20_lambda/**/*_en.md` | yes | yes | yes | Corrected auth, language/model fields, run/mux methods, and removed unsupported runtime/env/status claims. |
| Files | `src/30_docs/30_files/**/*_en.md` | yes | yes | yes | Corrected REST model, auth, bucket/file schemas, and removed S3/SLA/versioning/range claims. |
| Config | `src/30_docs/40_config/**/*_en.md` | yes | yes | yes | Corrected v0 Thing and v1 user config models; removed hierarchy/versioning/rollback claims. |
| Logs | `src/30_docs/50_logs/**/*_en.md` | yes | yes | yes | Corrected auth modes, raw/framed ingest, regex/follow, stats, and no retention claims. |
| Queues | `src/30_docs/60_queues/**/*_en.md` | yes | yes | yes | Corrected Tailon in-memory queue endpoints, NDJSON reads, and unsupported delete/list-clients notes. |
| Scheduler | `src/30_docs/70_scheduler/**/*_en.md` | yes | yes | yes | Corrected delayed task queue contract, leases, labels, responses, SSE snapshots, and errors. |
| KVNode | `src/30_docs/80_kvnode/**/*_en.md` | yes | yes | yes | Corrected collection/key API, auth behavior, replication commands, responses, and SDK names. |
| Console | `src/30_docs/90_console/**/*_en.md` | yes | yes | yes | Corrected fake API routes and version endpoint; removed invented `/api/v1` project/session docs. |
| Glue2 | `src/30_docs/100_glue2/**/*_en.md` | yes | yes | yes | Corrected gateway diagnostics, auth headers, stats/status shapes, and API key scope docs. |
| HolaMail | `src/30_docs/110_holamail/**/*_en.md` | yes | yes | yes | Corrected to basic SMTP listener; removed HTTP API/STARTTLS/AUTH/templates/tracking claims. |
| Run | `src/30_docs/120_run/**/*_en.md` | yes | yes | yes | Corrected console API payloads and push-oriented registry subset; removed invented run/container endpoints. |

## Audit Sources

| Service | Implementation Source |
|---|---|
| InceptionDB | `../inceptiondb/`, `../inceptiondb-saas/` |
| Lambda | `../lambdas/` |
| Files | `../files/` |
| Config | `../config/` |
| Logs | `../instantlogs-saas/`, `../instantlogs/` |
| Queues | `../tailon/` |
| Scheduler | `../scheduler/` |
| KVNode | `../kvnode/` |
| Console | `../console/`, `../run/` |
| Glue2 | `../glue2/` |
| HolaMail | `../holamail/` |
| Run | `../run/` |
