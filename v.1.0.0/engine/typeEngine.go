package engine

import (
	"github.com/kempertrasdesclub/server-cache-simulator/v.1.0.0/data"
	"github.com/kempertrasdesclub/server-cache-simulator/v.1.0.0/interfaces"
	"sync"
)

// Engine (Português): Objeto principal do framework engine
type Engine struct {
	sizeOfData   int
	sizeOfEvents int

	data         interfaces.Data
	interactions []interfaces.Interactions

	doNotRepeatKey map[string]bool

	eventList []Event
	cache     map[string]data.Cache

	totalSetAllCache   int
	totalSetOne        int
	totalInvalidateKey int
	totalInvalidateAll int
	totalGetAll        int
	totalGetKey        int

	setAllCache   float64
	setOne        float64
	invalidateKey float64
	invalidateAll float64
	getAll        float64
	getKey        float64

	mutex sync.RWMutex
}
