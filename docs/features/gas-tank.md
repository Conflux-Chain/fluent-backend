# Gas Tank Paymaster

Gas Tank allows smart accounts to pay EIP-4337 gas fees with ERC20 tokens. The user deposits tokens into the Gas Tank paymaster in advance. The paymaster sponsors the native-token gas first, then settles the ERC20 charge from the user's deposited balance after the UserOperation completes.

## Enablement

Gas Tank routes are registered when `GasTank.Address` is set. The configured Gas Tank contract must have the requested token enabled, and the token must also have a price configured in the price oracle so the backend can calculate the token cost.

## Payment Mechanism

Gas Tank uses the `REFUND` mechanism only. The user must deposit ERC20 tokens into the paymaster before requesting sponsorship. The backend checks the user's available balance, excluding any amount already reserved for withdrawal, before signing. The deployed paymaster contract settles the actual gas cost from that balance and returns the unused amount during `postOp`.

## Backend Flow

1. The user deposits ERC20 tokens into the Gas Tank paymaster and leaves enough available balance to cover the UserOperation.
2. The client calls `POST /aa/gastank/stub` with the smart account sender and ERC20 token address.
3. The backend validates that the token is allowed and that the sender has a non-zero available Gas Tank balance.
4. The client places the returned paymaster data into the UserOperation and estimates gas.
5. The client submits the estimated UserOperation to `POST /aa/gastank/sign`.
6. The backend validates the UserOperation and paymaster data, calculates the maximum token cost from the price oracle, and verifies that the available balance covers that cost.
7. The backend updates `maxTokenCost` and returns signed paymaster data. The client replaces the UserOperation's paymaster data with the signed value before submission.
8. The EntryPoint executes the UserOperation and the deployed paymaster contract settles the ERC20 charge in `postOp`, returning any unused amount according to the contract's settlement logic.

## Balance and Settlement

The backend does not transfer or lock tokens during the stub or signing requests. The available balance is the account's Gas Tank balance minus its pending withdrawal amount. The balance can therefore change between signing and UserOperation inclusion; clients and operators should account for this when choosing how much balance to keep available.

The signed paymaster data contains the token, the calculated maximum token cost, the validity window, and the paymaster signature. The token cost is calculated from the UserOperation's maximum gas cost and the configured token price.

## Related Code

- Gas Tank service: [`service/paymaster_gas_tank.go`](../../service/paymaster_gas_tank.go)
- Gas Tank API routes: [`api/route.go`](../../api/route.go)
- Gas Tank request models: [`api/models.go`](../../api/models.go)
- Gas Tank contract ABI: [`contract/GasTankPaymaster.abi.json`](../../contract/GasTankPaymaster.abi.json)
