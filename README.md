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
go test config_test.go
```

# Configuration

```
config/
└── test1
    ├── bot_agent.txt   - patterns to match bot user agents
    ├── bot_url.txt     - list of urls where to forward requests matched by (bot_agent.txt)
    └── default_url.txt - list of urls where to forward not matched requests

```

#
