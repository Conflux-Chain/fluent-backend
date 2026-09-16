# Fluent Backend Service - Agent Guide

This file is the shared project guidance for AI coding agents. Keep project-wide rules here and put detailed business behavior in the relevant document under `docs/`.

## Project Overview

Fluent Backend is a Go REST service for browser-extension wallets on Conflux eSpace. Its main concerns are EIP-7702 account upgrades, EIP-4337 paymaster sponsorship, Gas Tank settlement, and ERC20-based gas payment.

- Language: Go 1.23
- HTTP framework: Gin
- Persistence: GORM
- CLI and configuration: Cobra and Viper-based utilities
- Chain integration: go-ethereum and web3go
- API documentation: swaggo/OpenAPI

Treat transaction validation, signing, sponsorship, nonce handling, and private-key use as security-sensitive behavior.

## Architecture Rules

Follow the module boundaries and dependency direction documented in [`docs/architecture.md`](docs/architecture.md). Keep HTTP concerns in `api`, business decisions in `service`, persistence in `store`, and background processing in `worker`. Do not move business validation into controllers merely to shorten a service implementation.

`cmd/start.go` is the application composition root. Preserve configuration-driven feature dependencies when changing initialization or route registration. A nil optional service means that its routes and workers must remain disabled.

## Development Workflow

Use the narrowest relevant validation first, then run the repository-wide checks before completing a substantial change.

```bash
# Format changed Go files
gofmt -w <changed-go-files>

# Test a changed package first
go test ./service

# Repository checks used by CI
go test ./...
go build ./...
```

Use the equivalent package path for focused tests. Add or update tests when changing validation, encoding, persistence, configuration behavior, or transaction flow. Do not require network access in unit tests unless the test is explicitly an integration test.

After changing Swagger annotations, regenerate the OpenAPI output with:

```bash
swag init --parseDependency
```

## Generated Files

Do not manually edit generated output:

- `docs/docs.go`
- `docs/swagger.json`
- `docs/swagger.yaml`
- Go contract bindings in `contract/` that begin with `Code generated - DO NOT EDIT.`

Change the source annotation or ABI and regenerate instead. The repository does not currently document a canonical contract-binding generation command; determine and document the expected tool version and command before regenerating bindings. Handwritten helpers such as `contract/userop_extension.go` are valid editing surfaces.

## Change Rules

- Follow existing package boundaries, error handling, naming, and test patterns.
- Prefer small, focused changes over unrelated refactoring.
- Preserve public APIs unless the task explicitly requires a breaking change.
- Update `.env.example` and configuration documentation when configuration fields or enablement rules change.
- Update feature documentation when business flow, accepted risk, API behavior, or operational assumptions change.
- Keep comments focused on non-obvious constraints and reasons, not line-by-line narration.
- Do not introduce a new abstraction unless it removes meaningful duplication or enforces an established boundary.

## Security Constraints

- Never commit or print private keys, credentials, production RPC URLs, signed production transactions, or other secrets.
- Never weaken authorization, signature, chain ID, nonce, delegation, whitelist, balance, gas-cost, calldata, or simulation checks without explicit justification and focused tests.
- Preserve serialized sponsor transaction submission where it protects nonce allocation.
- Treat rate limiting and blacklisting as risk controls, not complete security boundaries.
- Preserve the binding between signed data and its chain, EntryPoint, paymaster, sender, validity window, and complete payload where applicable.
- Use test-only keys and synthetic transaction data in tests and examples.
- Escalate changes to signing or sponsorship rules for human security review.

## Definition of Done

A change is complete when all applicable items are satisfied:

- Changed Go files are formatted.
- Focused tests for the affected package pass.
- `go test ./...` and `go build ./...` pass, or any environment blocker is reported.
- New or changed behavior has appropriate regression coverage.
- Generated files are regenerated from their source rather than hand-edited.
- API, configuration, feature, and operational documentation are updated when their contracts change.
- No secrets or production-sensitive data are introduced.

## Documentation Index

- Project overview and setup: [`README.md`](README.md)
- System architecture and module boundaries: [`docs/architecture.md`](docs/architecture.md)
- Gas Tank behavior: [`docs/features/gas-tank.md`](docs/features/gas-tank.md)
- Token Pay behavior and accepted risks: [`docs/features/token-pay.md`](docs/features/token-pay.md)
- Verifying Paymaster behavior: [`docs/features/verifying-paymaster.md`](docs/features/verifying-paymaster.md)
- API schema: [`docs/swagger.yaml`](docs/swagger.yaml)

When documentation and code disagree, verify behavior with tests and the controlling implementation, then update the stale documentation in the same change.
