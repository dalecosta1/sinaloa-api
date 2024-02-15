package interfaces

type Executor interface {
    ExecuteCmd(jsonStr string, pyFile string, args string) (interface{}, error)
}
