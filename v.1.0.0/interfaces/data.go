package interfaces

import (
	"github.com/kempertrasdesclub/server-cache-simulator/v.1.0.0/data"
)

// Data (Português): Interface usada para popular a cache
type Data interface {
	// NewData (Português): Monta o conteúdo de uma chave a ser populada na cache.
	//   key:     chave da cache
	//   dataKey: conteúdo da chave
	//   err:     indicador de erro
	NewData() (key string, dataKey data.Cache, err error)
}
