package semaphore

import (
	"context"
	"math"

	"golang.org/x/sync/semaphore"
)

// Semaphore represents a custom semaphore with optional capacity.
type Semaphore struct {
	cap int64           // Maximum capacity of the semaphore, -1 indicates no limit.
	ctx context.Context // Context for the semaphore.
	sem *semaphore.Weighted
}

// Acquire acquires n units from the semaphore if the capacity is not unlimited.
func (s *Semaphore) Acquire(n int64) error {
	if s.cap != -1 {
		return s.sem.Acquire(s.ctx, n)
	}
	return nil
}

// Release releases n units back to the semaphore if the capacity is not unlimited.
func (s *Semaphore) Release(n int64) {
	if s.cap != -1 {
		s.sem.Release(n)
	}
}

// TryAcquire attempts to acquire n units from the semaphore without blocking if the capacity is not unlimited.
func (s *Semaphore) TryAcquire(n int64) bool {
	if s.cap != -1 {
		return s.sem.TryAcquire(n)
	}
	return true
}

// Resize adjusts the maximum capacity of the semaphore. Acquires or releases permits accordingly.
func (s *Semaphore) Resize(newCap int64) error {
	var err error
	if newCap != -1 {
		if newCap > s.cap {
			err = s.Acquire(newCap - s.cap)
			if err != nil {
				return err
			}
		} else if newCap < s.cap {
			s.Release(s.cap - newCap)
		}
	}
	s.cap = newCap
	return nil
}

// NewSemaphore creates and initializes a new Semaphore with the specified initial capacity.
func NewSemaphore(initialCap int64) *Semaphore {
	sem := Semaphore{
		ctx: context.Background(),
		sem: semaphore.NewWeighted(math.MaxInt64),
	}
	_ = sem.Resize(initialCap)
	return &sem
}
