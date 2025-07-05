# Coroot Webhook Proxy

This app receives webhooks from Coroot.
It receives JSON data. It checks the status and message. Then it sends message to VK Teams.

## How it works

1. Coroot sends a POST-response.
2. The app reads the JSON.
3. It creates a text message.
4. It sends the message to VK Teams.

## How to run

1. Build the Docker image:
```bash
docker build -t coroot-webhook-proxy .
```
2. Run in Kubernetes. Use Deployment, Service, Ingress.
3. Set environment variables:
`VK_URL` - API VK Teams URL
`VK_CHAT_ID` - Chat ID
`VK_TOKEN` - Access token

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
Deployed: city-gateway-01:Deployment:city-gateway
- 💔 Latency: 93.37% of requests faster 500ms (objective: 99%)
- 🎉 CPU usage: -7% compared to the previous deployment
- 💔 Memory usage: +6% compared to the previous deployment
- 💔 Logs: the number of errors in the logs has increased by 336%

http://coroot.apps.citydrive-main.dev.citydrive.tech/p/elf2lyow/app/city-gateway-01:Deployment:city-gateway/Deployments#c855dcb5b:1744807395
```

## Health
The app responds on patch `/health`
Kubernetes use this endpoint patch to check if the app is working.

## Schema

coroot -> coroot-webhook-proxy -> VK Teams

## Coroot documentation

Link - https://docs.coroot.com/alerting/webhook/
