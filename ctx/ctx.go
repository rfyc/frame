package ctx

import "context"

type Ctx struct {
	Action     string
	Controller string
	Context    context.Context
	Input      *Input
	Output     *Output
	Server     *Server
}
