# Verifying Paymaster Design

## Responsibilities

The design keeps the contract simple and places changeable sponsorship rules in the backend:

- The backend validates the UserOperation and signs the paymaster hash.
- The contract verifies the signer, validity period, and sender's delegated implementation at execution time.
- The worker indexes finalized sponsorships for accounting and rate limiting.

The paymaster hash must bind all sponsorship-relevant fields and the chain, EntryPoint, and paymaster domain.

## Backend Checks

Before signing, the backend checks:

- UserOperation structure, paymaster address, and EIP-7702 init-code rules.
- Maximum gas cost and finalized-UserOperation limits.
- `execute` or `executeBatch` calldata and every target against the contract whitelist.
- The delegated implementation against the paymaster's smart-account whitelist.
- Paymaster pause state and deposit balance.

## EIP-7702

For an existing delegation, `delegatedContract` is zero and `initCode` must be empty. The backend reads the sender's delegation from chain state and validates its implementation.

When the bundle carries a new authorization, the client supplies the intended `delegatedContract` and the UserOperation uses the EIP-7702 init-code marker. This address is only a signing-time hint, not proof of authorization. The bundler must include a matching authorization, and the contract must reject execution unless the sender code points to an approved implementation.

The execution-time contract check protects against omitted, replaced, or changed delegations after signing.

## Finalized Soft Limits

Signed UserOperations are not stored because users may never submit them. The worker instead stores finalized sponsorship events, and rate limits query those records.

These are intentionally soft limits: multiple signatures may be issued before any operation finalizes, so the configured count can be temporarily exceeded. This is accepted at the current stage because signing QPS is low, while exposure is reduced by per-IP limiting, a short `SignatureTimeout`, per-operation `MaxGasCost`, and the paymaster deposit balance.

Per-IP limiting is not a strict security boundary. The business must accept the worst-case cost of all valid signatures issued during one validity window:

```text
temporary exposure ~= valid signatures issued during the validity window * MaxGasCost
```

Signing volume, worker lag, finalized sponsorships, and deposit depletion should be monitored. Stricter off-chain controls can be added later without moving sponsorship policy into the contract.
