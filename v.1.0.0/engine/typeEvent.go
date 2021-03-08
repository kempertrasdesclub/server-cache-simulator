package engine

import (
	"github.com/kempertrasdesclub/server-cache-simulator/v.1.0.0/data"
	"github.com/kempertrasdesclub/server-cache-simulator/v.1.0.0/statistics"
)

// Event (Português): Arquiva os eventos a serem testados
type Event struct {
	DataCache data.Cache
	Event     statistics.CacheEvent
	Key       string
}
