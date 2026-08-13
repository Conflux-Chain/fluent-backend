# Gas Tank Paymaster Design

Gas Tank allows smart accounts to pay EIP-4337 gas fees with ERC20 tokens. The paymaster pays the native-token gas first, then settles the ERC20 charge.

## Modes

- `REFUND` (`0`): The user has already deposited ERC20 into the paymaster. Validation reserves the balance, and `postOp` settles the actual gas cost and returns the unused amount.
- `CREDIT` (`1`): The user has not deposited ERC20. The paymaster sponsors gas first; the UserOperation then executes `approve + depositToken`, and `postOp` settles from the deposited balance.

## CREDIT Mode Security Risk

The backend checks the user's ERC20 balance before issuing a paymaster signature, but this check does not lock the tokens. After obtaining the signature and before the UserOperation is included, the user can transfer the ERC20 away.

The paymaster has already sponsored validation, account execution, and `postOp` gas. However, `depositToken` then fails, so the paymaster cannot collect ERC20 and incurs bad debt. Shortening signature validity, checking the balance again, binding calldata, or blacklisting later cannot eliminate this risk.

## Production Requirement

**Do not enable `CREDIT` mode in production.**

`CREDIT` mode may only be used in testing or controlled scenarios where sponsored-gas losses are acceptable. Production must use `REFUND` mode so that the paymaster controls the ERC20 settlement balance before sponsoring gas.
