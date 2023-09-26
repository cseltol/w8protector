package validation

import validation "github.com/go-ozzo/ozzo-validation"

func RequiredIf(cond bool) validation.RuleFunc {
	return func(val interface{}) error {
		if cond {
			return validation.Validate(val, validation.Required)
		}
		return nil
	}
}
