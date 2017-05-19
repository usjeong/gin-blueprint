# gin-blueprint
golang web api blueprint based gin-gonic, viper, sqlx, glide

[![Go Report Card](https://goreportcard.com/badge/github.com/usjeong/gin-blueprint)](https://goreportcard.com/report/github.com/usjeong/gin-blueprint)

## Structure
```bash
.
├── api 
│   ├── api.go // app main
│   ├── context.go // gin handler
│   └── model
│       └── model.go // database
├── conf
│   └── config.go // global configurations, injectable dependency objects
├── glide.yaml // package manage configurations
├── main.go 
└── tests
    └── api_test.go
```
