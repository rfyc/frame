package route

import (
	"net/http"
)

type Handler struct {
	actions     map[string]IAction
	controllers map[string]IController
}

func (this *Handler) RegisterController(name string, controller IController) {
	if len(this.controllers) == 0 {
		this.controllers = map[string]IController{}
	}
	this.controllers[name] = controller
}

func (this *Handler) RegisterAction(controller_name, action_name string, action IAction) {
	if len(this.actions) == 0 {
		this.actions = map[string]IAction{}
	}
	this.actions[controller_name+"::"+action_name] = action
}

func (this *Handler) ServeHTTP(response http.ResponseWriter, request *http.Request) {

}

func (this *Handler) ServeTCP() {

}
