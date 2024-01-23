package multimutex

// MultiMutex is a mutex that allows for multiple concurrent locks, up to a specified maximum.
type MultiMutex struct {
	max int           // The maximum allowed concurrency. Set to -1 to disable MultiMutex.
	ch  chan struct{} // The channel used for managing concurrency.
}

// NewMultiMutex creates a new MultiMutex with the specified maximum concurrency.
// If maxConcurrent is set to -1, MultiMutex is disabled.
func NewMultiMutex(maxConcurrent int) *MultiMutex {
	mu := &MultiMutex{
		max: maxConcurrent,
		ch:  nil,
	}

	if maxConcurrent != -1 {
		mu.ch = make(chan struct{}, maxConcurrent)
	}

	return mu
}

// Lock acquires a lock. If MultiMutex is disabled, it does nothing.
func (m *MultiMutex) Lock() {
	if m.max == -1 {
		return
	}
	m.ch <- struct{}{}
}

// Unlock releases a lock. If MultiMutex is disabled, it does nothing.
func (m *MultiMutex) Unlock() {
	if m.max == -1 {
		return
	}
	<-m.ch
}
