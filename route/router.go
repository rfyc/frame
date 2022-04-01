package route

import (
	"strings"
)

type IRoute interface {
	Parse(path string) (IController, IAction)
	Default(IController)
	RegisterController(name string, exec IController)
	RegisterAction(controller, action string, exec IAction)
}

type router struct {
	actions     map[string]IAction
	controllers map[string]IController
}

func (this *router) Parse(path string) (execController IController, execAction IAction) {

	var (
		action     = ""
		controller = ""
		paths      = strings.Split(strings.Trim(path, "/"), "/")
		path_len   = len(paths)
	)

	//******** name ********//
	switch {
	case path_len == 1:
		controller = "/" + paths[0]
		action = "index"
	case path_len > 1:
		controller = "/" + strings.Join(paths[0:path_len-1], "/")
		action = paths[path_len-1]
	}
	//******** found controller ********//
	execController = this.controllers[strings.ToLower(controller)]
	//******** found action ********//
	execAction = this.actions[strings.ToLower(action)]

	return
}

func (this *router) RegisterController(name string, controller IController) {
	if len(this.controllers) == 0 {
		this.controllers = map[string]IController{}
	}
	this.controllers[name] = controller
}

func (this *router) RegisterAction(controller_name, action_name string, action IAction) {
	if len(this.actions) == 0 {
		this.actions = map[string]IAction{}
	}
	this.actions[controller_name+"::"+action_name] = action
}
