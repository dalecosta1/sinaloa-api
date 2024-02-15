// service/argocd_manager_service.go
package service

import (
    "encoding/json"
    "github.com/dalecosta1/sinaloa-api/interfaces"
)

type ArgocdManagerService interface {
    MultiActions(rawJSON json.RawMessage) (interface{}, error)
}

type argocdManagerService struct {
    pyExecutor interfaces.PythonExecutor
}

func NewArgocdManagerService(executor interfaces.PythonExecutor) ArgocdManagerService {
    return &argocdManagerService{
        pyExecutor: executor,
    }
}

func (service *argocdManagerService) MultiActions(rawJSON json.RawMessage) (interface{}, error) {
    return service.pyExecutor.ExecutePyFn(rawJSON, "api_module", "multi_actions")
}