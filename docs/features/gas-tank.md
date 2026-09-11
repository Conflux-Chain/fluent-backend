# Gas Tank Paymaster

Gas Tank allows smart accounts to pay EIP-4337 gas fees with ERC20 tokens. The paymaster pays the native-token gas first and settles the ERC20 charge after the UserOperation completes.

## Payment Modes

- `REFUND` (`0`): The user has already deposited ERC20 tokens into the paymaster. Validation reserves the user's balance. After execution, `postOp` settles the actual gas cost and returns the unused amount.
- `CREDIT` (`1`): The user has not deposited ERC20 tokens yet. The paymaster sponsors the UserOperation first. The UserOperation then executes `approve` and `depositToken`, allowing `postOp` to settle the charge from the newly deposited balance.

## Backend Flow

1. The client calls the Gas Tank stub endpoint for `CREDIT` or `REFUND` mode.
2. The backend validates the token and, for `REFUND`, checks that the sender has a non-zero Gas Tank balance.
3. The client places the returned paymaster data into the UserOperation and estimates gas.
4. The client submits the estimated UserOperation to the Gas Tank signing endpoint.
5. The backend calculates the maximum token cost, validates the UserOperation, and returns signed paymaster data.
6. The EntryPoint executes the UserOperation and the paymaster settles the ERC20 charge in `postOp`.

## CREDIT Mode Risk

The backend balance check does not lock the user's ERC20 tokens. After receiving a valid signature and before the UserOperation is included, the user can transfer the tokens away.

In that case, the paymaster may already have paid for validation, account execution, and `postOp`, while the `depositToken` call fails. The paymaster cannot collect the ERC20 charge and incurs bad debt.

The following measures do not eliminate this race condition:

- shortening the signature validity period;
- checking the token balance again before signing;
- binding more calldata into the signature; or
- blacklisting the user after a later failure.

## Production Policy

**Do not enable `CREDIT` mode in production.**

Use `REFUND` mode in production so that the paymaster controls the ERC20 settlement balance before sponsoring gas. `CREDIT` mode is limited to testing or controlled environments where sponsored-gas losses are explicitly accepted.

## Related Code

- Gas Tank service: [`service/paymaster_gas_tank.go`](../../service/paymaster_gas_tank.go)
- Gas Tank API routes: [`api/route.go`](../../api/route.go)
- Gas Tank request models: [`api/models.go`](../../api/models.go)
- Gas Tank contract ABI: [`contract/GasTankPaymaster.abi.json`](../../contract/GasTankPaymaster.abi.json)
