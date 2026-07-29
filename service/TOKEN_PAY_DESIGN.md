# Token Pay Security Design

## Overview

The token pay feature allows users to pay gas fees using ERC20 tokens instead of native currency (CFX), enabling a frictionless experience for new users who have 0 CFX but hold supported ERC20 tokens.

## Known Security Risks (By Design)

The current token pay implementation has two potential attack vectors that could lead to backend gas fee losses:

### Risk 1: Token Transfer Before Transaction Execution

After the backend transfers tokens to the user's address to cover gas fees, but before the transfer transaction is executed on-chain, a malicious user could transfer the ERC20 tokens to another address.

### Risk 2: User Pre-emption

After the backend initiates a token transfer, the user could construct their own transaction to:

- Replace the transfer transaction (e.g., via transaction mempool manipulation)
- Transfer the ERC20 tokens away before the backend's transfer is confirmed

## Why These Risks Are Accepted

The design decision to accept these security risks is based on the following factors:

1. **Low Per-Incident Cost**: Each individual gas fee loss is minimal. The financial impact of a single successful attack is bounded and acceptable.

2. **User Blacklisting**: When a transfer fails due to balance or nonce-related issues, the backend automatically blacklists the user's address and client IP. This prevents repeated attacks from the same source and increases the cost of attacking.

3. **Service Fee Buffer**: The backend is configured to collect a small service fee (by charging slightly more ERC20 tokens than required). This fee margin can cover losses from occasional fraudulent attempts.

4. **Sponsor Account Monitoring**: The backend sponsor account that pays gas fees is maintained with only a minimal CFX balance. The account balance is actively monitored with alerts set up for low balances to trigger timely refills. This limits the total exposure even in case of large-scale coordinated attacks from multiple IP addresses.

## Risk-Benefit Analysis

The token pay feature provides significant value to legitimate users:

- Users can access and use the dApp without owning any native currency
- Dramatically improves user onboarding experience
- Reduces friction for token holders

The risk of gas fee losses is negligible compared to these benefits, especially given the mitigation strategies in place.

## Recommendations for Monitoring

- Monitor sponsor account balance continuously
- Track failed token transfer attempts by address and IP
- Alert on unusual patterns of failures
- Periodically review blacklist for false positives
- Adjust service fee percentage if loss patterns emerge
