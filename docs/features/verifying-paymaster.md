# Verifying Paymaster

The Verifying Paymaster sponsors EIP-4337 UserOperations after the backend validates the operation and signs the paymaster payload. The on-chain contract verifies the signer, validity period, and sender delegation before allowing the EntryPoint to use the paymaster.

## Enablement

The service is enabled only when the paymaster address, smart-account whitelist, and contract whitelist are all configured. Initialization must also succeed with a chain-authorized signer.

## Responsibilities

The design keeps changeable sponsorship rules in the backend and keeps cryptographic and execution-critical checks in the contract:

- The backend validates the UserOperation and signs the paymaster hash.
- The contract verifies the authorized signer, validity period, and that the delegation in paymaster data matches the sender's actual delegation.
- The worker indexes finalized sponsorship events and stores UserOperation records used by rate limiting and operational reporting.

The signed hash includes the complete paymaster data, including the sender delegation, and is bound to the chain, EntryPoint, and paymaster domain.

## Backend Validation

Before signing, the backend checks:

- UserOperation structure, paymaster address, and EIP-7702 init-code rules;
- maximum gas cost and configured finalized-UserOperation limits;
- `execute` or `executeBatch` calldata and every execution against the configured execution policies;
- the delegation in paymaster data against the paymaster's smart-account whitelist;
- paymaster pause state; and
- paymaster deposit balance.

### Execution Policies

Execution policies are evaluated as alternatives. An execution is eligible for sponsorship when at least one policy allows it:

- The target-contract policy allows calls whose target is in `ContractWhitelist`.
- Optional DeFi policies allow product-specific calls after validating their calldata and token rules.

DeFi policies are configured under the `DeFi` section. The currently supported products are Uniswap V2 and Uniswap V3. When a router is not configured, no policy-specific validation or sponsorship rule is enabled.

### Uniswap V2 Policy

Set `DeFi.Uniswap.V2.Router` to enable the Uniswap V2 policy. The router is authorized by this product-specific configuration and must not also appear in `ContractWhitelist`; initialization fails if the address is duplicated. At startup, the backend reads and caches the router's WETH address. Initialization fails if that call fails.

The policy supports these router methods:

- `swapExactTokensForTokens`;
- `swapTokensForExactTokens`;
- `swapExactTokensForETH`;
- `swapTokensForExactETH`;
- `swapExactETHForTokens`; and
- `swapETHForExactTokens`.

Other router methods, malformed calldata, and paths containing fewer than two tokens are not eligible for sponsorship. Only the first and last tokens in the path are checked; intermediate tokens are not checked.

The input and output rules are:

- Token-to-token swaps require both path endpoints in `ContractWhitelist`. WETH is treated like any other ERC20 token in these methods and must be explicitly whitelisted when used as an endpoint.
- Token-to-ETH swaps require the input token in `ContractWhitelist` and the final path token to equal the router's WETH address.
- ETH-to-token swaps require the first path token to equal the router's WETH address and the output token in `ContractWhitelist`.

Token-input methods require a zero `Execution.Value`. ETH-input methods require a positive `Execution.Value`. Any policy evaluation error is treated as a rejected execution. `executeBatch` applies the same policy evaluation independently to every execution in the batch.

### Uniswap V3 Policy

Set `DeFi.Uniswap.V3.Router` to enable the Uniswap V3 policy. The router must not also appear in `ContractWhitelist`. The backend reads and caches WETH9 from the configured router at startup.

The policy supports `exactInputSingle`, `exactInput`, `exactOutputSingle`, and `exactOutput`. Other router methods, including `multicall`, are rejected. For the single-token methods, the backend reads `TokenIn` and `TokenOut` from the decoded parameters. For the path methods, it requires Uniswap V3 packed path encoding with at least 43 bytes and a layout of a 20-byte token followed by one or more 3-byte fees and 20-byte tokens. For `exactInput`, the first and final 20-byte tokens in the encoded path are treated as the input and output endpoints. For `exactOutput`, the path is interpreted in reverse: the final token is the input endpoint and the first token is the output endpoint.

Each endpoint must either be present in `ContractWhitelist` or equal the router's cached WETH9 address. The policy only validates the router method, calldata decoding, path shape, and token endpoints. It does not validate `Execution.Value`, so native-value requirements are not enforced by this V3 policy. `executeBatch` applies the same policy evaluation independently to every execution in the batch.

## Paymaster Data

Paymaster data uses the same encoding in the backend and contract. It contains:

- paymaster address;
- paymaster verification gas limit;
- paymaster post-operation gas limit;
- sender delegation;
- validity period; and
- backend signature.

The backend's stub endpoint returns the paymaster address and the data needed for gas estimation. The sign endpoint returns the final signed paymaster data.

## EIP-7702 Delegation

The paymaster data contains the sender's delegation address. The backend encodes this address using the same layout as the contract. When the client supplies the zero address as delegation, the backend reads the current delegation from chain state.

The signing service verifies that the delegation is included in the paymaster's whitelist. The request must include a `delegation` field; use the zero address when the backend should read the current delegation from chain state.

When a bundle carries a new authorization, the client supplies the intended delegation address and the UserOperation uses the EIP-7702 init-code marker. This address is a signing-time input, not proof that the authorization will be included. The bundler must include a matching authorization.

The deployed contract is expected to check during validation that the delegation in paymaster data matches the sender's actual delegation. The whitelist check remains an off-chain signing policy. Because the signed hash covers the complete paymaster data, the delegation cannot be changed after signing.

## Finalized Soft Limits

The backend does not store signed UserOperations because a client may never submit them. The worker stores finalized sponsorship events, and rate limits query those records.

These are intentionally soft limits: multiple signatures may be issued before any UserOperation finalizes, so the configured count can temporarily be exceeded. The current design accepts this because signing QPS is low and exposure is reduced by:

- per-IP limiting;
- a short `SignatureTimeout`;
- per-operation `MaxGasCost` in the chain's native smallest unit; and
- the paymaster deposit balance.

Per-IP limiting is not a strict security boundary. The business must accept the worst-case cost of all valid signatures issued during one validity window:

```text
temporary exposure ~= valid signatures issued during the validity window * MaxGasCost
```

Monitor signing volume, worker lag, finalized sponsorships, and deposit depletion. Stricter off-chain controls can be added later without moving sponsorship policy into the contract.

## Related API

- `GET /api/aa/paymaster/config`
- `POST /api/aa/paymaster/stub`
- `POST /api/aa/paymaster/sign`

## Related Code

- Verifying Paymaster service: [`service/paymaster_verifying.go`](../../service/paymaster_verifying.go)
- Verifying Paymaster API controller: [`api/controller_verifying_paymaster.go`](../../api/controller_verifying_paymaster.go)
- Verifying Paymaster API routes: [`api/route.go`](../../api/route.go)
- Verifying Paymaster contract ABI: [`contract/VerifyingPaymaster.abi.json`](../../contract/VerifyingPaymaster.abi.json)
