# Getting Started

This guide walks you through retrieving and updating your configuration using the Config API. Before you begin, you will need an API key with the appropriate permissions.

## Step 1: Retrieve Your Configuration

Use the `GET /v1/config` endpoint to fetch your current configuration:

```bash
curl "https://api.hola.cloud/v1/config" \
  -H "Api-Key: your-api-key" \
  -H "Api-Secret: your-api-secret"
```

Expected response:

```json
{
  "project": "my-app",
  "environment": "production",
  "services": {
    "database": {
      "host": "db.example.com",
      "port": 5432
    },
    "cache": {
      "host": "redis.example.com",
      "port": 6379
    }
  },
  "features": {
    "new-checkout": false
  }
}
```

## Step 2: Update a Specific Value

Use `PATCH /v1/config` with a partial JSON payload to update only the fields you need. For example, to enable the new checkout feature:

```bash
curl -X PATCH "https://api.hola.cloud/v1/config" \
  -H "Api-Key: your-api-key" \
  -H "Api-Secret: your-api-secret" \
  -H "Content-Type: application/json" \
  -d '{
    "features": {
      "new-checkout": true
    }
  }'
```

Expected response:

```json
{
  "project": "my-app",
  "environment": "production",
  "services": {
    "database": {
      "host": "db.example.com",
      "port": 5432
    },
    "cache": {
      "host": "redis.example.com",
      "port": 6379
    }
  },
  "features": {
    "new-checkout": true
  }
}
```

## Step 3: Verify the Update

Call `GET /v1/config` again to confirm your changes were applied:

```bash
curl "https://api.hola.cloud/v1/config" \
  -H "Api-Key: your-api-key" \
  -H "Api-Secret: your-api-secret"
```

The response should now show `"new-checkout": true`.

## Summary

In just a few steps you retrieved a configuration, applied a partial update, and verified the change. The Config API makes it easy to manage runtime settings without redeploying your application.
