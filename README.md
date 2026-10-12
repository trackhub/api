# Regenerating gqlgen
```
go tool gqlgen
```

# Regenerating gorm models
```
go tool gorm gen -i gorm/model/ -o gorm/generated
```

# Docker debug

example vscode configuration
```
    "configurations": [
        {
            "name": "Connect to a docker server",
            "type": "go",
            "request": "attach",
            "mode": "remote",
            "remotePath": "/app",
            "port": 4000,
            "host": "172.19.0.4"
        },
    ]
```
