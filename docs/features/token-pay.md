# Token Pay

Token Pay allows an EOA to submit a business transaction while paying the native-token gas fee with a supported ERC20 token. The backend sponsors the native gas and receives the ERC20 payment from the user.

## Transaction Flow

The client submits two user-signed transactions:

1. **Transfer-token transaction**: transfers the ERC20 gas payment to the configured backend recipient.
2. **Business transaction**: performs the user's intended operation.

The backend processes them in this order:

1. Validates both signed raw transactions.
2. Sends a funding transaction from the sponsor account.
3. Waits for the funding transaction to confirm.
4. Broadcasts the transfer-token transaction.
5. Broadcasts the business transaction immediately after the transfer-token transaction is accepted for submission.
6. Waits for the transaction receipts. The backend requires the transfer-token transaction to execute successfully, but currently does not act on the business transaction's execution result.

The transfer-token and business transactions are signed by the user before submission. The backend does not have the user's private key.

## API Contract

- `GET /api/tokenpay/config` returns supported tokens, the payment recipient, and gas and price policy parameters.
- `GET /api/tokenpay/price?token=0x...` returns the number of token smallest units equivalent to one native token.
- `POST /api/tokenpay/submit` accepts `rawTransferTokenTx` and `rawBusinessTx` as signed raw transactions.

The first token returned by the configuration endpoint is the default token used by the built-in test client for quoting and payment.

## Known Risks

The implementation accepts a bounded sponsor-loss risk during the interval between backend funding and confirmed transaction execution.

### Risk 1: Token Transfer Before Execution

After the backend funds the user's account but before the transfer-token transaction executes, the user can transfer the ERC20 tokens away. The subsequent transfer may fail, while the sponsor has already paid native gas.

### Risk 2: User Pre-emption

The user may submit a competing transaction that changes nonce ordering or otherwise prevents the backend's transfer-token transaction from executing as intended. This can also leave the sponsor with an unrecovered gas cost.

## Current Mitigations

- Validate both signed transactions before funding.
- Apply per-IP rate limiting to the sponsor endpoint.
- Blacklist the sender address and client IP when transfer-token submission fails with selected transaction-pool errors, or when the transfer-token transaction executes unsuccessfully before timeout.
- Keep the sponsor account funded with a controlled balance and monitor it continuously.

These controls reduce exposure but do not make the flow atomic and do not eliminate the underlying race conditions.

## Operational Monitoring

Monitor at least the following signals:

- sponsor account native-token balance;
- failed transfer-token transactions by address and IP;
- funding transactions that do not lead to successful token transfer;
- unusual increases in rejected or blacklisted requests; and
- aggregate sponsor losses.

Review blacklist decisions for false positives and adjust rate limits or the feature's enabled state when observed losses exceed the accepted operating boundary.

## Related Code

- Token Pay service: [`service/token_pay.go`](../../service/token_pay.go)
- Token Pay configuration: [`service/token_pay_config.go`](../../service/token_pay_config.go)
- Token Pay API controller: [`api/controller_token_pay.go`](../../api/controller_token_pay.go)
- Token Pay API routes: [`api/route.go`](../../api/route.go)
