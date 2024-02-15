package middlewares

import (
    "bytes"
    "encoding/json"
    "os/exec"
    "fmt"
)

// PythonMiddleware represents a middleware for executing Python scripts.
type PythonMiddleware struct{}

// NewPythonMiddleware creates a new instance of PythonMiddleware.
func NewPythonMiddleware() *PythonMiddleware {
    return &PythonMiddleware{}
}

// Executes a Python script and returns its output.
func (p *PythonMiddleware) ExecuteCmd(jsonStr string, pyFile string, cmd string) (interface{}, error) {
    // Build the command to execute the Python script.
    cmdArgs := append([]string{pyFile}, jsonStr)
    c := exec.Command(cmd, cmdArgs...)

    // Capture the output of the script.
    var out bytes.Buffer
    c.Stdout = &out

    // Run the command.
    err := c.Run()
    if err != nil {
        return nil, fmt.Errorf("failed to execute python script: %v", err)
    }

    // Here, we assume the script returns JSON that we want to unmarshal into a Go data structure.
    // Adjust the unmarshalling according to your needs.
    var result interface{}
    if err := json.Unmarshal(out.Bytes(), &result); err != nil {
        return nil, fmt.Errorf("failed to unmarshal script output: %v", err)
    }

    return result, nil
}
