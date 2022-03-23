package validator

type IRule interface {
	Valiedate() bool
}

type IRules []IRule

var rules = map[string]IRule{}

func Register(name string, rule IRule) {
	rules[name] = rule
}

func Rules(obj interface{}) IRules {

	return nil
}

func Validate(rules IRules) (bool, error) {

	return true, nil
}
