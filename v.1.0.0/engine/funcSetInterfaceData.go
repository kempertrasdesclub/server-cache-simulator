package engine

import (
	"github.com/kempertrasdesclub/server-cache-simulator/v.1.0.0/interfaces"
)

// SetInterfaceData (Português): Recebe o ponteiro do objeto responsável pela criação dos dados
func (e *Engine) SetInterfaceData(data interfaces.Data) {
	e.data = data
}
