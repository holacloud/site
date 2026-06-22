
# Getting Started with the Console

The Console is the single control plane for all HolaCloud services. It provides a unified web interface to manage databases, functions, files, configuration, logs, queues, and more — all from a central dashboard.

![Console Dashboard](img/console.png)

## Logging In

Navigate to [https://console.hola.cloud](https://console.hola.cloud) and sign in with your HolaCloud account credentials. If you are using demo or development mode, mock API endpoints are active, allowing you to explore all features without a real backend.

```bash
# Example: check login session via the mock API
curl -X GET "https://console.hola.cloud/fakeapi/authapi/auth/me"
```

Expected response:

```json
{
  "email": "fulanez@gmail.com",
  "id": "user-1234",
  "nick": "fulanez"
}
```

## Navigating the Dashboard

After logging in, the main dashboard displays:

- **Service Status** — health indicators for each active HolaCloud service (InceptionDB, Lambda, Files, Config, InstantLogs, Queues, Scheduler).
- **Recent Activity** — a chronological feed of actions performed in your project.
- **Key Metrics** — usage summaries such as request counts, storage size, and function invocations.

Use the left sidebar to jump between sections: Dashboard, Databases, Functions, Files, Config, Logs, Queues, Scheduler, and Project Settings.

## Viewing Service Status

Each service card on the dashboard shows the data returned by the configured service APIs. Project activity APIs are not exposed by the current Console fake API.

## Switching Between Projects

If you belong to multiple projects, use the project selector at the top of the sidebar. The entire UI switches context to the selected project and its services.

```bash
# List available projects
curl -X GET "https://console.hola.cloud/fakeapi/projectsapi/v0/projects"
```

```json
[
  { "id": "project-00000000-0000-0000-0000-000000000001", "name": "Hello" }
]
```

## Accessing Documentation and Support

The Console includes a built-in help menu (question-mark icon in the top bar) with links to:

- This documentation site
- API reference
- Community forum
- Support ticket system
- Service status page

## Summary

The Console acts as the single control plane for all HolaCloud services. Whether you are running in production or exploring in demo mode, the Console gives you full visibility and control over your cloud resources from one place.
