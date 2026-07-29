# Fluent Backend

REST backend service that enables browser-extension wallets to upgrade an **EOA (Externally Owned Account)** to an **EIP-7702 Account Abstraction smart-contract wallet** on Conflux eSpace.

The service includes three major capability groups:

1. **EIP-7702 Account Abstraction upgrade**
2. **Gas Tank paymaster signing flow**
3. **Token Pay sponsored transaction flow**

For account-abstraction upgrade, the service accepts a pre-signed EIP-7702 authorization message from the wallet, broadcasts the corresponding type-4 set-code transaction using its own **fee-payer account**, and lets the wallet poll for the on-chain result. So, the user's EOA never needs to hold native tokens for gas.

For Gas Tank and Token Pay, the service helps clients estimate/sign/pay gas in ERC20-denominated flows.

## Configuration

Please refer to `.env.example` file to prepare a `.env` file for service configurations.

At minimum, make sure you configure:

- RPC endpoint and sponsor private key (service signer)
- EIP-7702 delegated contract (optional restriction)
- Gas Tank paymaster address
- Token Pay recipient and supported token list

## Running the Service

```bash
# Build
go build

# Run (auto load .env file)
./fluent-backend
```

## Feature Overview

### 1) Account Abstraction (EIP-7702)

- `POST /api/aa/auth`: submit a signed EIP-7702 authorization
- `GET /api/aa/auth/:txHash`: query execution status/result

### 2) Gas Tank

Gas Tank APIs help wallet clients prepare and sign `paymasterData` for UserOperation.

- `POST /api/aa/gastank/prepare/credit`
  - Prepare `paymasterData` for CREDIT mode (deposit flow)
- `POST /api/aa/gastank/prepare/refund`
  - Prepare `paymasterData` for REFUND mode (spend existing Gas Tank balance)
- `POST /api/aa/gastank/sign`
  - Calculate `maxTokenCost`, sign paymaster payload, and return final `paymasterData`

Typical wallet flow:

1. Call `prepare/credit` or `prepare/refund`.
2. Put returned `paymasterData` into UserOperation and estimate gas.
3. Call `sign` with estimated UserOperation.
4. Replace UserOperation paymaster data with returned signed payload and submit.

### 3) Token Pay

Token Pay lets users submit two signed txs while backend sponsors native gas:

- tx A: transfer ERC20 token to backend recipient (gas payment)
- tx B: user business transaction

APIs:

- `GET /api/tokenpay/config`
  - Return supported tokens, recipient, and gas/price policy params
- `GET /api/tokenpay/price?token=0x...`
  - Return token units equivalent to 1 ETH (integer in token smallest unit)
- `POST /api/tokenpay/submit`
  - Submit both signed raw txs: `rawTransferTokenTx` and `rawBusinessTx`

Submit behavior:

1. Backend validates both signed raw txs.
2. Backend sends funding ETH tx from sponsor account.
3. After funding confirms, backend broadcasts transfer-token tx then business tx.
4. Completion is considered successful only when both transactions execute successfully.

## Business errors

There are some business errors for wallet to handle, see all defined errors on [github](./service/errors.go).

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

## CLI Helper: Test Token Pay End-to-End

A built-in `test-token-pay` sub-command performs an end-to-end token-pay test flow:

```bash
./fluent-backend test-token-pay \
  --backend https://api-testnet.fluentwallet.com/api \
  --rpc     https://evmtestnet.confluxrpc.com \
  --key     0x<FAUCET_PRIVATE_KEY>
```

It will:

1. Create a random test account
2. Faucet supported token to that account
3. Build transfer-token/business tx pair
4. Simulate both txs
5. Submit to `/api/tokenpay/submit`
6. Wait for business tx receipt

---

## Swagger UI

When `api.swaggerEnabled: true`, the interactive Swagger UI is available at `swagger/index.html`.
