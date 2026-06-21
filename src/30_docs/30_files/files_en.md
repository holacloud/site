# Files

HolaCloud Files is an S3-compatible object storage service designed for storing and retrieving any amount of data at any time. It provides a simple RESTful API for managing buckets and files, with built-in scalability, durability, and security.

## Key Features

### S3-Compatible API
HolaCloud Files exposes a familiar S3-compatible API, making it easy to migrate existing applications or integrate with tools that speak the S3 protocol.

### Scalability
Files scales automatically to handle any amount of data, from a few gigabytes to petabytes, without any manual provisioning or capacity planning.

### Durability and Availability
Data is stored redundantly across multiple availability zones, ensuring 99.999999999% durability and 99.99% availability.

### Security
All data is encrypted at rest using AES-256 and in transit using TLS 1.3. Access is controlled via API keys with fine-grained permissions.

### Versioning-Ready
The service is designed with versioning support, allowing you to preserve, retrieve, and restore every version of every file in your buckets.

## Use Cases

### Backups and Disaster Recovery
Files provides a reliable destination for database backups, system snapshots, and disaster recovery archives. Its durability guarantees ensure your data survives hardware failures and regional outages.

### Static Assets for Web Applications
Host images, CSS, JavaScript, and other static assets with low latency and high throughput. Files integrates seamlessly with HolaCloud CDN for global content delivery.

### User-Generated Content
Handle file uploads from your users -- profile pictures, documents, videos, and more -- with a scalable, secure storage layer.

### Data Lakes and Analytics
Store structured and unstructured data at petabyte scale for analytics workloads. Files integrates with HolaCloud Compute and HolaCloud Analytics for data processing pipelines.

### Media Storage and Streaming
Store audio, video, and image files for media applications. Files supports range requests for efficient partial content delivery.

## Integration with HolaCloud

HolaCloud Files works seamlessly with other HolaCloud services:

- **HolaCloud Compute** -- Attach Files volumes to your compute instances for persistent storage.
- **HolaCloud CDN** -- Deliver file content globally through edge caching.
- **HolaCloud Logs** -- Archive access logs directly to Files buckets.
- **HolaCloud InceptionDB** -- Store database backup snapshots in Files.

## API Reference

All endpoints are under the base URL `https://api.hola.cloud` and require authentication via `Api-Key` and `Api-Secret` headers.

| Method | Path | Description |
|--------|------|-------------|
| GET | /v1/buckets | List all buckets |
| POST | /v1/buckets | Create a new bucket |
| GET | /v1/buckets/{id} | Get bucket details |
| PATCH | /v1/buckets/{id} | Modify bucket metadata |
| DELETE | /v1/buckets/{id} | Delete a bucket |
| GET | /v1/buckets/{id}/list/* | List files in a bucket |
| PUT | /v1/buckets/{id}/files/* | Upload a file |
| GET | /v1/buckets/{id}/files/* | Download a file |
| DELETE | /v1/buckets/{id}/files/* | Delete a file |
| HEAD | /v1/buckets/{id}/files/* | Get file metadata |
