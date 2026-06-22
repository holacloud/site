
# Managing Projects and Services

Projects are the top-level organizational unit in HolaCloud. They group together service instances and routing settings. This guide covers what is currently documented for the Console.

## Creating and Switching Between Projects

To create a new project, open the project selector in the sidebar and click **New Project**. Provide a name and optional description. Once created, the Console switches to the new project automatically.

The project creation API is currently not covered by the Console documentation.

To switch projects, click the current project name in the sidebar and select a different one from the dropdown.

```bash
# List all accessible projects
curl -X GET "https://console.hola.cloud/fakeapi/projectsapi/v0/projects"
```

```json
[
  { "id": "project-00000000-0000-0000-0000-000000000001", "name": "Hello" }
]
```

## Managing Service Instances per Project

Each project contains its own instances of HolaCloud services. Navigate to the service section (e.g., Databases, Functions, Files, Config, Logs, Queues, Scheduler) to manage instances within the current project.

For example, to list InceptionDB databases in the current project:

```bash
curl -X GET "https://console.hola.cloud/fakeapi/inceptionapi/v1/databases"
```

```json
[
  { "id": "00000000-0000-0000-0000-000000000001", "name": "db1", "owners_length": 1 }
]
```

The same pattern applies to Lambda functions, Files buckets, Config stores, InstantLogs streams, Queues, and Scheduler jobs.

## Project API Coverage

Console project administration APIs are currently not covered by these docs. API key management is implemented by the serviceprojects API under `/v0/apikeys`.

## Import/Export Configuration as JSON

Import/export endpoints are currently not covered by the Console documentation.

## Summary

Projects in HolaCloud provide isolated environments for your applications. Through the Console you can switch projects and inspect the service data exposed by the configured APIs.
