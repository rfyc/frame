package validator

type Method struct {
	Func func() error
}

func (this *Method) Validate() (bool, error) {

	if this != nil {
		if err := this.Func(); err != nil {
			return false, err
		}

	}
	return true, nil
}
