# simple-tds

Service for routing HTTP requests based on different patterns


# Development

Building:
```
go build
```

Run:
```
./simple-tds
```

Access:
```
curl -v http://localhost:8000/
curl -v http://localhost:8000/home
curl -v http://localhost:8000/go/test1
curl  -A "Potato" -v http://localhost:8000/go/test1
```

Redirect logic test:
```
go test
```

# Configuration

Variables:
```
HTTP_PORT - listening HTTP port (default: 8000)
CONFIG_RELOAD_INTERVAL - configuration reload interval (default: 60)
CONFIG_DIR  - location of configuration directories
```

Configuration directory:
```
config/
└── test1
    ├── bot_agent.txt   - patterns to match bot user agents
    ├── bot_url.txt     - list of urls where to forward requests matched by (bot_agent.txt)
    └── default_url.txt - list of urls where to forward not matched requests

```

# Deployment

Build/re-use docker images

https://hub.docker.com/r/denis256/simple-tds
