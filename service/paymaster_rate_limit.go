package service

import (
	"fmt"
	"time"

	"github.com/Conflux-Chain/fluent-backend/contract"
	"github.com/Conflux-Chain/fluent-backend/store"
)

type LimitConfig struct {
	// MaxPending is the maximum number of pending user operations allowed per sender
	// and per IP address. If set to a non-positive value, this limit is disabled.
	MaxPending int64 `default:"100"`

	// ValidUntils defines the limits for user operations based on their signature expiration time.
	// The key is a descriptive name (e.g., "daily", "monthly"), and the value specifies the duration
	// and the maximum number of user operations allowed within that duration. For example, use below
	// configuration for both daily and monthly limits:
	//
	// "daily": { 24h, 10}
	// "monthly": { 720h, 100 }
	ValidUntils map[string]struct {
		Duration time.Duration
		Max      int64
	}
}

// Limiter is an interface for limiting the signed user operations.
type Limiter interface {
	Limit(userOp *contract.PackedUserOperation, ip string) error
}

func NewLimiter(config LimitConfig, store *store.Store) Limiter {
	var limiters []Limiter

	if config.MaxPending > 0 {
		limiters = append(limiters, NewPendingCountLimiter(config.MaxPending, store))
	}

	for _, v := range config.ValidUntils {
		limiters = append(limiters, NewValidUntilLimiter(v.Duration, v.Max, store))
	}

	return CompositeLimiter(limiters)
}

// CompositeLimiter is a composite of multiple Limiter instances.
type CompositeLimiter []Limiter

func (limiter CompositeLimiter) Limit(userOp *contract.PackedUserOperation, ip string) error {
	for _, v := range limiter {
		if err := v.Limit(userOp, ip); err != nil {
			return err
		}
	}

	return nil
}

// PendingCountLimiter limits the number of pending user operations per sender and/or per IP address.
type PendingCountLimiter struct {
	maxPendings int64
	store       *store.Store
}

func NewPendingCountLimiter(maxPendings int64, store *store.Store) *PendingCountLimiter {
	return &PendingCountLimiter{
		maxPendings: maxPendings,
		store:       store,
	}
}

func (limiter *PendingCountLimiter) Limit(userOp *contract.PackedUserOperation, ip string) error {
	pendingsBySender, err := limiter.store.UserOp.GetPendingCount(userOp.Sender)
	if err != nil {
		return err
	}

	if pendingsBySender >= limiter.maxPendings {
		return ErrVerifyingPaymasterTooManyOps.WithData(fmt.Sprintf(
			"Exceeds the maximum %v pending user operations by sender %v", limiter.maxPendings, userOp.Sender,
		))
	}

	pendingsByIP, err := limiter.store.UserOp.GetPendingCountByIP(ip)
	if err != nil {
		return err
	}

	if pendingsByIP >= limiter.maxPendings {
		return ErrVerifyingPaymasterTooManyOps.WithData(fmt.Sprintf(
			"Exceeds the maximum %v pending user operations by IP %v", limiter.maxPendings, ip,
		))
	}

	return nil
}

// ValidUntilLimiter limits the number of user operations based on signature expiration time.
type ValidUntilLimiter struct {
	duration time.Duration
	max      int64
	store    *store.Store
}

func NewValidUntilLimiter(duration time.Duration, max int64, store *store.Store) *ValidUntilLimiter {
	return &ValidUntilLimiter{
		duration: duration,
		max:      max,
		store:    store,
	}
}

func (limiter *ValidUntilLimiter) Limit(userOp *contract.PackedUserOperation, ip string) error {
	since := time.Now().Add(-limiter.duration)

	count, err := limiter.store.UserOp.GetCountByValidUntil(userOp.Sender, since)
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
