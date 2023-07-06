# UptimeKuma API in GO

This project is a REST API implementation for UptimeKuma, a monitoring and uptime
service. It is built using the Go programming language (Golang).

# Installation

- Clone the repository: git clone https://github.com/your/uptime-kuma-api-go.git
- Navigate to the project directory: cd uptime-kuma-api-go
- Install the dependencies: go install
- Run: go build main serve

# Support

| Resource                   | Method | Support? |
|----------------------------|--------|----------|
| - Ping Average             |        | [X]      |
| Get Ping Average           | GET    | [X]      |
| Get Ping Average           | GET    | [X]      |
| - Uptime                   |        | [X]      |
| Get Uptime                 | GET    | [X]      |
| Get Uptimes                | GET    | [X]      |
| - Heartbeat List           |        | [-]      |
| Heartbeat                  | GET    | [ ]      |
| Heartbeats                 | GET    | [ ]      |
| - Important Heartbeat List |        | [-]      |
| Important Heartbeat        | GET    | [ ]      |
| Important Heartbeats       | GET    | [ ]      |
| - Monitor List             |        | [-]      |
| Get All Monitors           | GET    | [X]      |
| Create Monitor             | POST   | [X]      |
| Get Monitor                | GET    | [X]      |
| Delete Monitor             | DELETE | [X]      |
| Update Monitor             | PATCH  | [X]      |
| Pause Monitor              | PATCH  | [ ]      |
| Resume Monitor             | PATCH  | [ ]      |
| Monitor Beats              | GET    | [ ]      |
| Add Tag Monitor            | POST   | [ ]      |
| Remove Tag Monitor         | DELETE | [ ]      |
| - API Key List             |        | [-]      |
| API Key                    | GET    | [ ]      |
| - Docker Host List         |        | [-]      |
| Docker Host                | GET    | [ ]      |
| - Info                     |        | [-]      |
| Get Info                   | GET    | [X]      |
| - Maintenance List         |        | [-]      |
| Get Maintenances           | GET    | [ ]      |
| Create Maintenance         | POST   | [ ]      |
| Get Maintenance            | GET    | [ ]      |
| Delete Maintenance         | DELETE | [ ]      |
| Update Maintenance         | PATCH  | [ ]      |
| Pause Maintenance          | POST   | [ ]      |
| Resume Maintenance         | POST   | [ ]      |
| Add Monitor Maintenance    | GET    | [ ]      |
| Add Monitor Maintenance    | POST   | [ ]      |
| - Notification List        |        | [-]      |
| - Proxy List               |        | [-]      |
| - Status Page List         |        | [-]      |
| Get All Status Pages       | GET    | [ ]      |
| Add Status Page            | POST   | [ ]      |
| Get Status Page            | GET    | [ ]      |
| Save Status Page           | POST   | [ ]      |
| Delete Status Page         | DELETE | [ ]      |
| Post Incident              | POST   | [ ]      |
| Unpin Incident             | DELETE | [ ]      |
| - Cert Info                |        | [-]      |
| Get Cert Info              | GET    | [ ]      |
| - Tags                     |        | [X]      |
| Get Tags                   | GET    | [X]      |
| Add Tag                    | POST   | [X]      |
| Get Tag                    | GET    | [X]      |
| Delete Tag                 | DELETE | [X]      |

# Contributing

Contributions are welcome! If you have any ideas, suggestions, or bug reports, 
please open an issue or submit a pull request.

# License

This project is licensed under the MIT License.

# References

- UptimeKuma: https://uptimekuma.com
- Go Documentation: https://golang.org/doc/