# Coroot Webhook Proxy
This app receives webhooks from [Coroot](https://github.com/coroot/coroot).
It receives JSON data. It checks the status and message. Then it sends message to [VK Teams](https://teams.vk.com/botapi/).

## How it works
1. Coroot sends a POST-response to Coroot Webhook Proxy.
2. The app reads the JSON.
3. It creates a text message.
4. It sends the message to VK Teams.

## How to run
1. Clone this repo.
2. Build the Docker image:
```bash
docker build -t coroot-webhook-proxy:v0.1.0 .
```
or pull image from [github](https://github.com/Gakhramanzode/coroot-webhook-proxy/pkgs/container/coroot-webhook-proxy).

3. Run in Docker:
```bash
docker run --name coroot-webhook-proxy -p 8080:8080 \
-e VK_URL=vk \
-e VK_CHAT_ID=chat \
-e VK_TOKEN=token \
coroot-webhook-proxy:v0.1.0
```
or run in Kubernetes. Use Deployment, Service, Ingress, Secrets.

4. Set environment variables:
- `VK_URL` - API VK Teams URL
- `VK_CHAT_ID` - Chat ID
- `VK_TOKEN` - Access token

and some envs for Coroot OpenTelemetry Integration (more info about Coroot [tracing](https://docs.coroot.com/tracing/overview)):
- `OTEL_SERVICE_NAME`
- `OTEL_EXPORTER_OTLP_TRACES_ENDPOINT`
- `OTEL_EXPORTER_OTLP_LOGS_ENDPOINT`
- `OTEL_EXPORTER_OTLP_PROTOCOL`
- `OTEL_METRICS_EXPORTER`
- `OTEL_EXPORTER_OTLP_HEADERS`

## Example
Example JSON from Coroot:
```json
{
  "status": "Deployed",
  "application": "default:Deployment:app1",
  "version": "123ab456: app:v1.8.2",
  "summary": [
    "💔 Availability: 87% (objective: 99%)",
    "💔 CPU usage: +21% (+$37/mo) compared to the previous deployment",
    "🎉 Memory: looks like the memory leak has been fixed"
  ],
  "url": "http://127.0.0.1:8080/p/x0xwl4jz/app/default:Deployment:app1/Deployments#123ab456:123"
}
```
Example message for VK Teams:
```
Deployed: gateway:Deployment:gateway
- 💔 Latency: 93.37% of requests faster 500ms (objective: 99%)
- 🎉 CPU usage: -7% compared to the previous deployment
- 💔 Memory usage: +6% compared to the previous deployment
- 💔 Logs: the number of errors in the logs has increased by 336%

http://coroot.apps.tech/p/elf2lyow/app/gateway:Deployment:gateway/Deployments#c855dcb5b:1744807395
```

## Health
The app responds on patch `/health`
Kubernetes use this endpoint patch to check if the app is working.

## Schema
coroot -> coroot-webhook-proxy -> VK Teams

## Coroot documentation
Link - https://docs.coroot.com/alerting/webhook/
