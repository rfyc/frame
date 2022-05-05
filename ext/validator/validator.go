package validator

import "fmt"

type IRule interface {
	Validate() (bool, error)
}

type IRules []IRule

var rules = map[string]IRule{}

func Register(name string, rule IRule) {
	rules[name] = rule
}

type iAction interface {
	Rules() IRules
}

func Rules(action interface{}) IRules {
	if act, ok := action.(iAction); ok {
		return act.Rules()
	}
	return IRules{}
}

func Validate(rules IRules) (bool, error) {
	for _, rule := range rules {
		if ok, err := rule.Validate(); false == ok {
			return ok, err
		}
	}
	return true, nil
}

func errorf(rule_error error, errmsg string, args ...interface{}) error {
	if rule_error == nil {
		return fmt.Errorf(errmsg, args...)
	} else {
		return fmt.Errorf("%w: "+errmsg, append([]interface{}{rule_error}, args...))
	}
}
