# UptimeKuma API in GO

This project is a REST API implementation for UptimeKuma, a monitoring and uptime
service. It is built using the Go programming language (Golang).

# Installation

- Clone the repository: git clone https://github.com/your/uptime-kuma-api-go.git
- Navigate to the project directory: cd uptime-kuma-api-go
- Install the dependencies: go install
- Run: go build main serve

# How to use this image

## Start a ```uptime-kuma-api-go``` server instance

Starting a Uptime Kuma Api GO instance is simple:

```$ docker run --name uptimekuma-api-go leandrose/uptime-kuma-api-go:latest```

## via docker-compose

Example docker-compose.yaml:

```
version: '3'

services:
    uptimekumaapi:
        image: leandrose/uptime-kuma-api-go:latest
        environment:
            UPTIMEKUMA_URI: wss://DOMAIN/socket.io/?EIO=4&transport=websocket
            UPTIMEKUMA_USERNAME: user
            UPTIMEKUMA_PASSWORD: pass
        ports:
            - 3000:3000
        restart: always

```

# Environment Variables

### UPTIMEKUMA_URI

The UPTIMEKUMA_URI environment variable represents the URI of the WebSocket connection
for UptimeKuma. It defines the endpoint where the client application will establish a
connection to the UptimeKuma server.

Example usage:

```
UPTIMEKUMA_URI=wss://DOMAIN/socket.io/?EIO=4&transport=websocket
```

### UPTIMEKUMA_USERNAME

The UPTIMEKUMA_USERNAME environment variable is used to specify the username for
accessing the UptimeKuma WebSocket. It is the credential used by the client application
to authenticate and interact with the UptimeKuma server.

### UPTIMEKUMA_PASSWORD

The UPTIMEKUMA_PASSWORD environment variable is used to set the password for accessing
the UptimeKuma WebSocket. It works in conjunction with the UPTIMEKUMA_USERNAME to
authenticate the client application with the UptimeKuma server.

# Support

## API Key List

| Resource | Method | Support? |
|----------|--------|----------|
| API Key  | GET    | ❌        |

## Cert Info

| Resource      | Method | Support? |
|---------------|--------|----------|
| Get Cert Info | GET    | ❌        |

## Docker Host List

| Resource    | Method | Support? |
|-------------|--------|----------|
| Docker Host | GET    | ❌        |

## Heartbeat List

| Resource   | Method | Support? |
|------------|--------|----------|
| Heartbeats | GET    | ✅️       |

## Important Heartbeat List

| Resource             | Method | Support? |
|----------------------|--------|----------|
| Important Heartbeats | GET    | ❌        |

## Info

| Resource | Method | Support? |
|----------|--------|----------|
| Get Info | GET    | ✅️       |

## Maintenance List

| Resource                | Method | Support? |
|-------------------------|--------|----------|
| Get Maintenances        | GET    | ❌        |
| Create Maintenance      | POST   | ❌        |
| Get Maintenance         | GET    | ❌        |
| Delete Maintenance      | DELETE | ❌        |
| Update Maintenance      | PATCH  | ❌        |
| Pause Maintenance       | POST   | ❌        |
| Resume Maintenance      | POST   | ❌        |
| Add Monitor Maintenance | GET    | ❌        |
| Add Monitor Maintenance | POST   | ❌        |

## Monitor List

| Resource           | Method | Support? |
|--------------------|--------|----------|
| Get All Monitors   | GET    | ✅️       |
| Create Monitor     | POST   | ✅️       |
| Get Monitor        | GET    | ✅️       |
| Delete Monitor     | DELETE | ✅️       |
| Update Monitor     | PATCH  | ✅️       |
| Pause Monitor      | PATCH  | ✅️       |
| Resume Monitor     | PATCH  | ✅️       |
| Monitor Beats      | GET    | ❌        |
| Add Tag Monitor    | POST   | ✅️       |
| Remove Tag Monitor | DELETE | ✅️       |

## Notification List

| Resource            | Method | Support? |
|---------------------|--------|----------|
| Get Notifications   | GET    | ✅️       |
| Create Notification | GET    | ✅️       |
| Delete Notification | GET    | ✅️       |

## Ping Average

| Resource             | Method | Support? |
|----------------------|--------|----------|
| Get All Ping Average | GET    | ✅️       |
| Get Ping Average     | GET    | ✅️       |

## Proxy List

| Resource | Method | Support? |
|----------|--------|----------|

## Status Page List

| Resource             | Method | Support? |
|----------------------|--------|----------|
| Get All Status Pages | GET    | ✅️       |
| Add Status Page      | POST   | ✅️       |
| Get Status Page      | GET    | ✅️       |
| Save Status Page     | POST   | ✅️       |
| Delete Status Page   | DELETE | ✅️       |
| Post Incident        | POST   | ✅️       |
| Unpin Incident       | DELETE | ✅️       |

## Uptime

| Resource    | Method | Support? |
|-------------|--------|----------|
| Get Uptime  | GET    | ✅️       |
| Get Uptimes | GET    | ✅️       |

## Tags

| Resource   | Method | Support? |
|------------|--------|----------|
| Get Tags   | GET    | ✅️       |
| Add Tag    | POST   | ✅️       |
| Get Tag    | GET    | ✅️       |
| Delete Tag | DELETE | ✅️       |

# Contributing

Contributions are welcome! If you have any ideas, suggestions, or bug reports,
please open an issue or submit a pull request.

# License

This project is licensed under the MIT License.

# References

- UptimeKuma: https://uptimekuma.com
- Go Documentation: https://golang.org/doc/