## Fintual CLI project

This app is a convinient wrap around the fintual API:
- [Fintual API](https://fintual.cl/api-docs/index.html)

### Development useful resources

- [Log rocket request tutorial](https://blog.logrocket.com/making-http-requests-in-go/)
- [Tutorial for ini files](https://www.kelche.co/blog/go/ini/)


### Project layout

fintual-cli/
├── cmd/
│   ├─- user/
│   │   ├─- get.go
│   │   ├─- set.go
│   │   └── user.go
│   ├── banks.go
│   └── root.go
├── internals/
│   ├─- models/
│   │   ├── bank_model.go
│   │   └── config.go
│   ├─- repositories/
│   │   ├─- config_repository.go
│   │   └── fintual_repository.go
├── tests/
├── go.mod
├── go.sum
└── main.go