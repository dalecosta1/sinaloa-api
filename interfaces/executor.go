// interfaces/executor.go
package interfaces

import "encoding/json"

type PythonExecutor interface {
    ExecutePyFn(rawJSON json.RawMessage, pyFile string, pyFunction string) (interface{}, error)
}
