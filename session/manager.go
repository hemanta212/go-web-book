package session

import (
	"fmt"
	"sync"
)

type Manager struct {
	cookieName  string     // private cookiename
	lock        sync.Mutex // protects session
	provider    Provider
	maxlifetime int64
}

func NewManager(provideName, cookieName string, maxlifetime int64) (*Manager, error) {
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("Session : unknown provide %q (forgotten import?)", provideName)
	}
	return &Manager{cookieName: cookieName, provider: provider, maxlifetime: maxlifetime}, nil
}

type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
}

type Session interface {
	Set(key, value interface{}) error       // set session value
	Get(key, value interface{}) interface{} // get session value
	Delete(key interface{}) error           // delete session value
	SessionId() string                      // back current sessionID
}

var provides = make(map[string]Provider)

// Register make a session provider available by the provided name.
// if
