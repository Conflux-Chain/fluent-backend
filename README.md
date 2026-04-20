# Fluent Backend

REST backend service that enables browser-extension wallets to upgrade an **EOA (Externally Owned Account)** to an **EIP-7702 Account Abstraction smart-contract wallet** on Conflux eSpace.

The service accepts a pre-signed EIP-7702 authorization message from the wallet, broadcasts the corresponding type-4 set-code transaction using its own **fee-payer account**, and lets the wallet poll for the on-chain result. So, the user's EOA never needs to hold native tokens for gas.

## Configuration

Please refer to `.env.example` file to prepare a `.env` file for service configurations.

## Running the Service

```bash
# Build
go build

# Run (auto load .env file)
./fluent-backend
```

## Business errors

There are some business errors for wallet to handle, see all defined errors [here](./service/errors.go).

## CLI Helper: Generate Auth Message

A built-in `auth` sub-command generates and prints a signed EIP-7702 authorization message, useful for testing:

```bash
./fluent-backend auth \
  --url      https://evmtestnet.confluxrpc.com \
  --key      0x<YOUR_PRIVATE_KEY> \
  --contract 0x<DELEGATED_CONTRACT_ADDRESS>
```

The output is a JSON object ready to POST to `/api/aa/auth`.

> **Warning:** Pass your private key via an environment variable or a secrets file in production to avoid shell history exposure.

---

## Swagger UI

When `api.swaggerEnabled: true`, the interactive Swagger UI is available at `swagger/index.html`.
