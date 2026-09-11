# Verifying Paymaster

The Verifying Paymaster sponsors EIP-4337 UserOperations after the backend validates the operation and signs the paymaster payload. The on-chain contract verifies the signer, validity period, and sender delegation before allowing the EntryPoint to use the paymaster.

## Responsibilities

The design keeps changeable sponsorship rules in the backend and keeps cryptographic and execution-critical checks in the contract:

- The backend validates the UserOperation and signs the paymaster hash.
- The contract verifies the authorized signer, validity period, and that the delegation in paymaster data matches the sender's actual delegation.
- The worker indexes finalized sponsorship events for accounting and rate limiting.

The signed hash includes the complete paymaster data, including the sender delegation, and is bound to the chain, EntryPoint, and paymaster domain.

## Backend Validation

Before signing, the backend checks:

- UserOperation structure, paymaster address, and EIP-7702 init-code rules;
- maximum gas cost and finalized-UserOperation limits;
- `execute` or `executeBatch` calldata and every target against the contract whitelist;
- the delegation in paymaster data against the paymaster's smart-account whitelist;
- paymaster pause state; and
- paymaster deposit balance.

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

The paymaster data contains the sender's delegation address. The backend encodes this address using the same layout as the contract. When the client does not supply a delegation, the backend reads it from chain state.

The signing service verifies that the delegation is included in the paymaster's whitelist.

When a bundle carries a new authorization, the client supplies the intended delegation address and the UserOperation uses the EIP-7702 init-code marker. This address is a signing-time input, not proof that the authorization will be included. The bundler must include a matching authorization.

During validation, the contract checks that the delegation in paymaster data matches the sender's actual delegation. The whitelist check remains an off-chain signing policy. Because the signed hash covers the complete paymaster data, the delegation cannot be changed after signing.

## Finalized Soft Limits

The backend does not store signed UserOperations because a client may never submit them. The worker stores finalized sponsorship events, and rate limits query those records.

These are intentionally soft limits: multiple signatures may be issued before any UserOperation finalizes, so the configured count can temporarily be exceeded. The current design accepts this because signing QPS is low and exposure is reduced by:

- per-IP limiting;
- a short `SignatureTimeout`;
- per-operation `MaxGasCost`; and
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
