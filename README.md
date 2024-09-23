# Golang Template

This is a sample repository to create a new golang project with a simple api server.

## Features

- [x] Config
- [x] Logger
- [ ] HTTP Server
- [ ] Database
- [ ] Kafka


## Development

### Configuration

Configuration is done using the `config.yaml` file using the [viper](https://github.com/spf13/viper) library.

Environment variable can be set before running the server with prefix `APP_` to override the yaml file values.

Validation of config is done by verifying valus with the help of struct tags. 

### Running the server

#### Local Development

```bash
make local-run
```

### Logger

Logging is done using the [zerolog](https://github.com/rs/zerolog) library.


## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
