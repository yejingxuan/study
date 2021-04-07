#   CICD

## 一、drone部署

### 1、docker部署
```yaml
version: '2'

services:
  drone-server:
    image: drone/drone:1
    ports:
      - 8000:80
    volumes:
      - ./:/data
    restart: always
    environment:
      - DRONE_SERVER_HOST=47.98.48.32:8000
      - DRONE_SERVER_PROTO=http
      - DRONE_RPC_SECRET=123456

      # GitHub Config
      - DRONE_GITHUB_SERVER=https://github.com
      - DRONE_GITHUB_CLIENT_ID=d0a6948ae6aa19101178
      - DRONE_GITHUB_CLIENT_SECRET=eb606a31a23af0c820bf1c0644044eba90d01fd8

      - DRONE_LOGS_PRETTY=true
      - DRONE_LOGS_COLOR=true

  # runner for docker version
  drone-runner:
    image: drone/drone-runner-docker:1
    restart: always
    depends_on:
      - drone-server
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    environment:
      - DRONE_RPC_HOST=drone-server
      - DRONE_RPC_PROTO=http
      - DRONE_RPC_SECRET=123456
      - DRONE_RUNNER_CAPACITY=3
```