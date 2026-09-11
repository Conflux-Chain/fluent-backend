# Verifying Paymaster Design

## Responsibilities

The design keeps the contract simple and places changeable sponsorship rules in the backend:

- The backend validates the UserOperation and signs the paymaster hash.
- The contract verifies the signer, validity period, and that the delegation in paymaster data matches the sender's actual delegation.
- The worker indexes finalized sponsorships for accounting and rate limiting.

The paymaster hash includes the complete paymaster data, including the sender delegation, and binds the chain, EntryPoint, and paymaster domain.

Paymaster data uses the same encoding in the backend and contract: paymaster address, gas limits, delegation, validity period, and signature.

## Backend Checks

Before signing, the backend checks:

- UserOperation structure, paymaster address, and EIP-7702 init-code rules.
- Maximum gas cost and finalized-UserOperation limits.
- `execute` or `executeBatch` calldata and every target against the contract whitelist.
- The delegation in paymaster data against the paymaster's smart-account whitelist.
- Paymaster pause state and deposit balance.

## EIP-7702

The paymaster data contains the sender's delegation address. The backend `Stub` interface encodes this address using the same layout as the contract; when it is not supplied, the backend reads it from chain state. The signing service verifies that the address is in the paymaster's whitelist.

When the bundle carries a new authorization, the client supplies the intended delegation address and the UserOperation uses the EIP-7702 init-code marker. This address is only a signing-time value, not proof of authorization. The bundler must include a matching authorization. During validation, the contract only checks that the delegation in paymaster data matches the sender's actual delegation; whitelist validation is performed by the backend signing service.

The paymaster hash covers the complete paymaster data, so the delegation cannot be changed after signing. The execution-time contract check also protects against omitted or changed delegations.

## Finalized Soft Limits

Signed UserOperations are not stored because users may never submit them. The worker instead stores finalized sponsorship events, and rate limits query those records.

These are intentionally soft limits: multiple signatures may be issued before any operation finalizes, so the configured count can be temporarily exceeded. This is accepted at the current stage because signing QPS is low, while exposure is reduced by per-IP limiting, a short `SignatureTimeout`, per-operation `MaxGasCost`, and the paymaster deposit balance.

Per-IP limiting is not a strict security boundary. The business must accept the worst-case cost of all valid signatures issued during one validity window:

```text
temporary exposure ~= valid signatures issued during the validity window * MaxGasCost
```

Signing volume, worker lag, finalized sponsorships, and deposit depletion should be monitored. Stricter off-chain controls can be added later without moving sponsorship policy into the contract.
