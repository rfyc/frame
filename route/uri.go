package route

import (
	"strings"
)

type RegistHandler interface {
	Parse(path string) (IController, IAction)
	Controller(exec IController, name ...string)
	Action(exec IAction, controller, action string)
}

type DefaultRegister struct {
	actions           map[string]IAction
	controllers       map[string]IController
	defaultController IController
}

func (this *DefaultRegister) Parse(path string) (execController IController, execAction IAction) {

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

func (this *DefaultRegister) Controller(controller IController, names ...string) {
	if len(this.controllers) == 0 {
		this.controllers = map[string]IController{}
	}
	if len(names) > 0 {
		for _, name := range names {
			this.controllers[name] = controller
		}
	} else {
		this.defaultController = controller
	}
}

func (this *DefaultRegister) Action(action IAction, controller_name, action_name string) {
	if len(this.actions) == 0 {
		this.actions = map[string]IAction{}
	}
	this.actions[controller_name+"::"+action_name] = action
}
