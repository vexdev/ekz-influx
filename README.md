# EKZ-Influx

Simple InfluxDB exporter for EKZ smart meter. Available as a Docker container.

Environment variables used:

- `EKZ_USERNAME`: Username to connect to the my EKZ portal
- `EKZ_PASSWORD`: Password to connect to the my EKZ portal
- `EKZ_INSTALLATION_ID`: Installation ID of the smart meter, this is also called "Anlage" in the my EKZ portal
- `INFLUX_URL`: URL of the InfluxDB server
- `INFLUX_TOKEN`: Token to write to the InfluxDB server
- `INFLUX_ORG`: Organization to write to in the InfluxDB server
- `INFLUX_BUCKET`: Bucket to write to in the InfluxDB server

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
      - INFLUX_URL=http://influxdb:8086
      - INFLUX_TOKEN=your_token
      - INFLUX_ORG=your_org
      - INFLUX_BUCKET=your_bucket
    restart: unless-stopped
```