## Drone

- [Drone](#drone)
  - [一、部署](#一部署)

### 一、部署

```yaml
version: '3'
services:
  drone-server:
    image: drone/drone:1.0.0-rc.5
    ports:
      - "8000:80"
    volumes:
      - ./data/drone:/var/lib/drone/
      - /var/run/docker.sock:/var/run/docker.sock
    environment:
      - DRONE_OPEN=true
      - DRONE_SERVER_PROTO=http
      - DRONE_SERVER_HOST=47.98.48.32:8000
      - DRONE_DEBUG=true
      - DRONE_GIT_ALWAYS_AUTH=false
      - DRONE_GITHUB_SERVER=https://github.com
      - DRONE_GITHUB_CLIENT_ID=d0a6948ae6aa19101178
      - DRONE_GITHUB_CLIENT_SECRET=eb606a31a23af0c820bf1c0644044eba90d01fd8
      - DRONE_GITHUB=true
      - DRONE_PROVIDER=github
      - DRONE_DATABASE_DATASOURCE=/var/lib/drone/drone.sqlite
      - DRONE_DATABASE_DRIVER=sqlite3
      - DRONE_RPC_SECRET=ALQU2M0KdptXUdTPKcEw
   
  drone-agent:
    image: drone/agent:1.0.0-rc.5
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    depends_on:
      - drone-server
    environment:
      - DRONE_RPC_SERVER=http://47.98.48.32:8000
      - DRONE_RPC_SECRET=ALQU2M0KdptXUdTPKcEw
      - DRONE_DEBUG=true
```

