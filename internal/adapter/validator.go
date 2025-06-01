package adapter

func WithValidator(v Validator) Option {
	return func(a *Adapter) {
		a.Validator = v
	}
}
