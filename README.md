# EKZ-Influx

Simple InfluxDB exporter for EKZ smart meter. Available as a Docker container.

Environment variables used:

- `EKZ_USERNAME`: Username to connect to the my EKZ portal
- `EKZ_PASSWORD`: Password to connect to the my EKZ portal
- `EKZ_INSTALLATION_ID`: Installation ID of the smart meter, this is also called "Anlage" in the my EKZ portal
- `INFLUXDB_URL`: URL of the InfluxDB server
- `INFLUXDB_TOKEN`: Token to write to the InfluxDB server
- `INFLUXDB_ORG`: Organization to write to in the InfluxDB server
- `INFLUXDB_BUCKET`: Bucket to write to in the InfluxDB server

## Docker Compose

```yaml
version: '3.7'

services:
  ekz-influx:
    image: vexdev/ekz-influx:latest
    container_name: ekz-influx
    environment:
      - EKZ_USERNAME=your_username
      - EKZ_PASSWORD=your_password
      - EKZ_INSTALLATION_ID=your_installation_id
      - INFLUXDB_URL=http://influxdb:8086
      - INFLUXDB_TOKEN=your_token
      - INFLUXDB_ORG=your_org
      - INFLUXDB_BUCKET=your_bucket
    restart: unless-stopped
```