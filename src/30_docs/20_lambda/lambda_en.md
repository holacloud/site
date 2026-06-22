<h1>Lambda</h1>

HolaCloud Lambda lets you publish small HTTP handlers that run inside HolaCloud. A lambda stores its source code, route metadata, owner, and project association, then responds through the Lambda API or public run routes.

## Key Features

### Supported Languages
Lambda accepts these `language` values: `javascript`, `static-html`, `static-css`, and `static-js`.

### HTTP Routing
Each lambda can store a `method` and `path`. Use these fields to describe how the lambda should be reached through the mux router or your application routing layer.

### Public and Admin Invocation
Use authenticated admin endpoints to create, inspect, update, delete, and run lambdas. Use public run and mux endpoints when a lambda must be callable by webhooks, browsers, or other external clients.

## Use Cases

- **HTTP APIs**: Build small request handlers backed by HolaCloud services.
- **Webhooks**: Receive third-party events through a public `/run/{lambda_id}` URL.
- **Static Responses**: Serve HTML, CSS, or client-side JavaScript snippets with the static language modes.
- **Mux Routing**: Route owner-scoped paths through `/mux/{owner_id}/*`.

## Lambda Fields

Lambda records use these fields: `id`, `created_timestamp`, `owner`, `project_id`, `name`, `language`, `code`, `method`, and `path`.
