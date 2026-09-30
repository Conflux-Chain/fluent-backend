# Fluent Backend

REST backend service for browser-extension wallets on Conflux eSpace. It provides EIP-4337 paymaster services and ERC20-based gas payment through Gas Tank. Free EOA-to-smart-account upgrades use the EIP-4337 flow with EIP-7702 delegation.

## Features

- **Verifying Paymaster**: validates and signs UserOperation paymaster data subject to delegation, contract, account, gas-cost, and deposit policies. See the [Verifying Paymaster documentation](docs/features/verifying-paymaster.md).
- **Gas Tank**: prepares and signs ERC20 paymaster data using the `REFUND` mechanism. Users must deposit ERC20 tokens into the Gas Tank paymaster in advance, and the backend validates the available balance before signing. See the [Gas Tank documentation](docs/features/gas-tank.md).

## Legacy Features

**Token Pay** was an early interim solution and is no longer used in production. Its code, API, and CLI helpers remain in the repository for legacy reference and non-production use. Leave `APP_SERVICE_TOKENPAY_RECIPIENT` unset or zero in production to keep its service and routes disabled. See the [legacy Token Pay documentation](docs/features/token-pay.md).

## Configuration

Configure the required environment variables in a `.env` file before deployment. The supported configuration items and their comments are documented in the [`.env.example`](.env.example) file.

Keep sponsor private keys outside source control and do not expose them in shell history.

## Running the Service

Before building and running the service, make sure that:

- Go `1.23.0` or a compatible Go 1.23 toolchain is installed.
- The required variables in [`.env.example`](.env.example) are configured in `.env`, including a reachable Conflux eSpace RPC endpoint.
- The database settings are configured for the store used by the service.
- A sponsor private key and sufficient on-chain funds are available when the enabled features need to submit transactions or sponsor gas.

```bash
# Build
go build

# Run
./fluent-backend
```

Run the test suite with:

```bash
go test ./...
```

## API Documentation

API documentation is generated from the controller code annotations. The generated OpenAPI files are available in the [`docs`](docs) directory:

- [OpenAPI YAML](docs/swagger.yaml)
- [OpenAPI JSON](docs/swagger.json)

Run `swag init --parseDependency` after changing API annotations. The generated files should not be edited manually.

When `SwaggerEnabled` is enabled in the API configuration, the interactive Swagger UI is available at `/swagger/index.html`.

## Project Documentation

- [Architecture](docs/architecture.md)
- [Gas Tank](docs/features/gas-tank.md)
- [Legacy Token Pay](docs/features/token-pay.md)
- [Verifying Paymaster](docs/features/verifying-paymaster.md)

## Business Errors

Business errors are defined in [service/errors.go](service/errors.go).

## CLI Commands

The binary includes auxiliary commands for testing and development. Run `./fluent-backend --help` or `./fluent-backend <command> --help` for command usage and options.

Commands that sign transactions require a private key. Use test keys only and avoid passing production keys on the command line.
