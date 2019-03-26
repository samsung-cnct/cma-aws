package cluster

import "sync"

// ErrorValue contains the error and command output
type ErrorValue struct {
	CmdError  error
	CmdOutput string
}

// ErrorMap is used to temporarily store errors from eksctl commands
type ErrorMap struct {
	sync.RWMutex
	// key is string of: cluster-name + datacenter-region + aws-access-key-id
	// value contains error output
	internal map[string]ErrorValue
}

func NewErrorMap() *ErrorMap {
	return &ErrorMap{
		internal: make(map[string]ErrorValue),
	}
}

func (em *ErrorMap) Load(key string) (value ErrorValue, found bool) {
	em.RLock()
	result, ok := em.internal[key]
	em.RUnlock()
	return result, ok
}

func (em *ErrorMap) Store(key string, value ErrorValue) {
	em.Lock()
	em.internal[key] = value
	em.Unlock()
}

func (em *ErrorMap) Delete(key string) {
	em.Lock()
	delete(em.internal, key)
	em.Unlock()
}
