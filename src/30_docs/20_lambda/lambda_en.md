<h1>Lambda</h1>

HolaCloud Lambda is a serverless compute service that lets you run code in response to events without provisioning or managing servers. Upload your functions in JavaScript, Go, or Python, and HolaCloud handles everything from auto-scaling to high availability.

## Key Features

### Multi-Language Support
Write functions in the language of your choice. HolaCloud Lambda supports JavaScript (Node.js), Go, and Python runtimes out of the box, with a consistent invocation interface across all languages.

### Auto-Scaling
Lambda functions scale automatically based on incoming request volume. Whether you have one request per day or thousands per second, HolaCloud adjusts capacity instantly — you only pay for what you use.

### Webhook Support
Every Lambda function is automatically reachable via a public URL. Use it as a webhook endpoint for external services like Stripe, GitHub, or Slack, or integrate it with your own applications.

### Pay-Per-Use Billing
No idle costs. You are billed only for the compute time your code consumes during execution, rounded up to the nearest millisecond.

## Use Cases

- **APIs**: Build lightweight REST or GraphQL backends without managing infrastructure.
- **Webhooks**: Receive and process events from third-party services with a dedicated public endpoint.
- **Background Jobs**: Offload tasks like email sending, image processing, or report generation.
- **Data Processing**: Transform, filter, or enrich data streams in real time.

## How It Fits in the HolaCloud Ecosystem

Lambda works alongside other HolaCloud services to power full-stack serverless applications:

- **InceptionDB** — store and query persistent data from your functions.
- **HolaAuth** — authenticate users and protect endpoints.
- **Mux Router** — map custom domains to specific Lambda functions and owners.

All services share a common API gateway, authentication model, and CLI tooling, making it straightforward to compose them into a unified application.
