# Gas Tank Paymaster

Gas Tank allows smart accounts to pay EIP-4337 gas fees with ERC20 tokens. The paymaster pays the native-token gas first and settles the ERC20 charge after the UserOperation completes.

## Enablement

Gas Tank routes are registered only when at least one price-oracle USDT token is configured and `GasTank.Address` is set. Both `CREDIT` and `REFUND` routes are exposed when the service is enabled; the backend does not enforce a mode-level production switch.

## Payment Modes

- `REFUND` (`0`): The user has already deposited ERC20 tokens into the paymaster. The backend verifies the sender's balance before signing. The deployed paymaster contract settles the actual gas cost and returns the unused amount during `postOp`.
- `CREDIT` (`1`): The user has not deposited ERC20 tokens yet. The paymaster sponsors the UserOperation first. The UserOperation then executes `approve` and `depositToken`, allowing the deployed paymaster contract to settle the charge from the newly deposited balance during `postOp`.

## Backend Flow

1. The client calls the Gas Tank stub endpoint for `CREDIT` or `REFUND` mode.
2. The backend validates the token and, for `REFUND`, the stub checks for a non-zero Gas Tank balance. The signing endpoint verifies that the balance covers the calculated maximum token cost.
3. The client places the returned paymaster data into the UserOperation and estimates gas.
4. The client submits the estimated UserOperation to the Gas Tank signing endpoint.
5. The backend calculates the maximum token cost, validates the UserOperation, and returns signed paymaster data.
6. The EntryPoint executes the UserOperation and the deployed paymaster contract settles the ERC20 charge in `postOp`.

## CREDIT Mode Risk

The backend balance check does not lock the user's ERC20 tokens. After receiving a valid signature and before the UserOperation is included, the user can transfer the tokens away.

In that case, the paymaster may already have paid for validation, account execution, and `postOp`, while the `depositToken` call fails. The paymaster cannot collect the ERC20 charge and incurs bad debt.

The following measures do not eliminate this race condition:

- shortening the signature validity period;
- checking the token balance again before signing;
- binding more calldata into the signature; or
- blacklisting the user after a later failure.

## Production Policy

**Do not expose `CREDIT` mode in production.**

Use `REFUND` mode in production so that the paymaster controls the ERC20 settlement balance before sponsoring gas. Because the current backend does not enforce this policy, production deployments must restrict access to the `CREDIT` endpoint operationally. `CREDIT` mode is limited to testing or controlled environments where sponsored-gas losses are explicitly accepted.

## Related Code

- Gas Tank service: [`service/paymaster_gas_tank.go`](../../service/paymaster_gas_tank.go)
- Gas Tank API routes: [`api/route.go`](../../api/route.go)
- Gas Tank request models: [`api/models.go`](../../api/models.go)
- Gas Tank contract ABI: [`contract/GasTankPaymaster.abi.json`](../../contract/GasTankPaymaster.abi.json)
