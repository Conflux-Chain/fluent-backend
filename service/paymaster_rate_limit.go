package service

import (
	"fmt"
	"time"

	"github.com/Conflux-Chain/fluent-backend/contract"
	"github.com/Conflux-Chain/fluent-backend/store"
)

type LimitConfig struct {
	// Finalized defines the limits for finalized user operations on chain.
	// The key is a descriptive name (e.g., "daily", "monthly"), and the value specifies the duration
	// and the maximum number of user operations allowed within that duration. For example, use below
	// configuration for both daily and monthly limits:
	//
	// "daily": { 24h, 10 } means a maximum of 10 finalized user operations are allowed within a 24-hour period.
	// "monthly": { 720h, 100 } means a maximum of 100 finalized user operations are allowed within a 720-hour period.
	Finalized map[string]struct {
		Duration time.Duration
		Max      int64
	}
}

// Limiter is an interface for limiting the signed user operations.
type Limiter interface {
	Limit(userOp *contract.PackedUserOperation) error
}

func NewLimiter(config LimitConfig, store *store.Store) Limiter {
	var limiters []Limiter

	for _, v := range config.Finalized {
		limiters = append(limiters, NewFinalizedLimiter(v.Duration, v.Max, store))
	}

	return CompositeLimiter(limiters)
}

// CompositeLimiter is a composite of multiple Limiter instances.
type CompositeLimiter []Limiter

func (limiter CompositeLimiter) Limit(userOp *contract.PackedUserOperation) error {
	for _, v := range limiter {
		if err := v.Limit(userOp); err != nil {
			return err
		}
	}

	return nil
}

// FinalizedLimiter limits the number of finalized user operations on chain.
type FinalizedLimiter struct {
	duration time.Duration
	max      int64
	store    *store.Store
}

func NewFinalizedLimiter(duration time.Duration, max int64, store *store.Store) *FinalizedLimiter {
	return &FinalizedLimiter{
		duration: duration,
		max:      max,
		store:    store,
	}
}

func (limiter *FinalizedLimiter) Limit(userOp *contract.PackedUserOperation) error {
	since := time.Now().Add(-limiter.duration)

	count, err := limiter.store.UserOp.GetCountByBlockTimestamp(userOp.Sender, since)
	if err != nil {
		return err
	}

	if count >= limiter.max {
		return ErrVerifyingPaymasterTooManyOps.WithData(fmt.Sprintf(
			"Exceeds the maximum %v user operations in %v", limiter.max, limiter.duration,
		))
	}

	return nil
}
