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
    pyExecutor interfaces.Executor
}

func NewArgocdManagerService(executor interfaces.Executor) ArgocdManagerService {
    return &argocdManagerService{
        pyExecutor: executor,
    }
}

func (service *argocdManagerService) MultiActions(rawJSON json.RawMessage) (interface{}, error) {
    return service.pyExecutor.ExecuteCmd(string(rawJSON), "multi_actions.py", "python3")
}