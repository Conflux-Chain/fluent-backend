# Fluent Backend

REST backend service for browser-extension wallets on Conflux eSpace. It provides EIP-7702 account upgrades, EIP-4337 paymaster services, and ERC20-sponsored native-gas payment.

## Features

- **EIP-7702 Account Abstraction**: accepts a signed authorization, submits the type-4 set-code transaction with the service account as fee payer, and provides a status endpoint.
- **Verifying Paymaster**: validates and signs UserOperation paymaster data subject to delegation, contract, account, gas-cost, and deposit policies. See the [Verifying Paymaster documentation](docs/features/verifying-paymaster.md).
- **Gas Tank**: prepares and signs ERC20 paymaster data for `REFUND` and `CREDIT` modes. `CREDIT` mode is not suitable for production; see the [Gas Tank documentation](docs/features/gas-tank.md).
- **Token Pay**: sponsors native gas for a pair of user-signed transactions, one ERC20 payment and one business transaction. See the [Token Pay documentation](docs/features/token-pay.md).

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

Run `swag init` after changing API annotations. The generated files should not be edited manually.

When `SwaggerEnabled` is enabled in the API configuration, the interactive Swagger UI is available at `/swagger/index.html`.

## Business Errors

Business errors are defined in [service/errors.go](service/errors.go).

## CLI Commands

The binary includes auxiliary commands for testing and development. Run `./fluent-backend --help` or `./fluent-backend <command> --help` for command usage and options.

Commands that sign transactions require a private key. Use test keys only and avoid passing production keys on the command line.
