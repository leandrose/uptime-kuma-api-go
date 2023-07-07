# UptimeKuma API in GO

This project is a REST API implementation for UptimeKuma, a monitoring and uptime
service. It is built using the Go programming language (Golang).

# Installation

- Clone the repository: git clone https://github.com/your/uptime-kuma-api-go.git
- Navigate to the project directory: cd uptime-kuma-api-go
- Install the dependencies: go install
- Run: go build main serve

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
| Get All Status Pages | GET    | ❌        |
| Add Status Page      | POST   | ❌        |
| Get Status Page      | GET    | ❌        |
| Save Status Page     | POST   | ❌        |
| Delete Status Page   | DELETE | ❌        |
| Post Incident        | POST   | ❌        |
| Unpin Incident       | DELETE | ❌        |

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