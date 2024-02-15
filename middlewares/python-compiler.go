package middlewares

import (
    "encoding/json"
    "fmt"
    "unsafe"
    "bytes"
    "os/exec"
    "github.com/dalecosta1/sinaloa-api/interfaces" // Assuming interfaces package exists and is used elsewhere appropriately
)

/*
#cgo pkg-config: python3
#include <Python.h>
*/
import "C"

type PythonMiddleware struct{}

func NewPythonMiddleware() *PythonMiddleware {
    return &PythonMiddleware{}
}

// Ensure PythonMiddleware implements interfaces.PythonExecutor.
var _ interfaces.PythonExecutor = &PythonMiddleware{}

func (p *PythonMiddleware) ExecutePyFn(
    rawJSON json.RawMessage,
    pyFile string,
    pyFunction string,
) (interface{}, error) {
    C.Py_Initialize()

    jsonStr := C.CString(string(rawJSON))
    defer C.free(unsafe.Pointer(jsonStr))

    moduleName := C.CString(pyFile)
    funcName := C.CString(pyFunction)
    defer func() {
        C.free(unsafe.Pointer(moduleName))
        C.free(unsafe.Pointer(funcName))
    }()

    pModule := C.PyImport_ImportModule(moduleName)
    if pModule == nil {
        C.PyErr_Print()
        C.Py_Finalize()
        return nil, fmt.Errorf("failed to load Python module %s", pyFile)
    }

    pFunc := C.PyObject_GetAttrString(pModule, funcName)
    if pFunc == nil || C.PyCallable_Check(pFunc) == 0 {
        if pFunc != nil {
            C.Py_DecRef(pFunc)
        }
        C.Py_DecRef(pModule)
        C.Py_Finalize()
        return nil, fmt.Errorf("failed to load function %s from module %s", pyFunction, pyFile)
    }

    pArgs := C.PyTuple_New(1)
    pValue := C.PyUnicode_FromString(jsonStr)
    C.PyTuple_SetItem(pArgs, 0, pValue) // Tuple_SetItem steals a reference

    pResult := C.PyObject_CallObject(pFunc, pArgs)
    if pResult == nil {
        C.PyErr_Print()
        C.Py_DecRef(pArgs)
        C.Py_DecRef(pFunc)
        C.Py_DecRef(pModule)
        C.Py_Finalize()
        return nil, fmt.Errorf("Python function call failed (%s from module %s)", pyFunction, pyFile)
    }

    var goResult interface{}
    if C.PyUnicode_Check(pResult) > 0 {
        pyStr := C.PyUnicode_AsUTF8String(pResult)
        if pyStr == nil {
            C.Py_DecRef(pResult)
            C.Py_Finalize()
            return nil, fmt.Errorf("failed to convert Python return value to string")
        }
        cStr := C.PyBytes_AsString(pyStr)
        resultStr := C.GoString(cStr)
        C.Py_DecRef(pyStr)

        if err := json.Unmarshal([]byte(resultStr), &goResult); err != nil {
            C.Py_DecRef(pResult)
            C.Py_Finalize()
            return nil, fmt.Errorf("failed to unmarshal Python return value: %v", err)
        }
    } else {
        // Handle other return types or set an error
    }

    C.Py_DecRef(pResult)
    C.Py_DecRef(pArgs)
    C.Py_DecRef(pFunc)
    C.Py_DecRef(pModule)
    C.Py_Finalize()

    return goResult, nil
}

func ExecuteCmd(command string, args ...string) (string, error) {
    cmd := exec.Command(command, args...)
    var out bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = &out
    err := cmd.Run()
    if err != nil {
        return "", err
    }
    return out.String(), nil
}
