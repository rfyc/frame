package validator

type Method func() error

func (this *Method) Validate() (bool, error) {

	if this != nil {
		if err := (*this)(); err != nil {
			return false, err
		}

	}
	return true, nil
}
