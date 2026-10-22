package main

type Runner interface {
	Run() string
}

type ValueRunner struct{}

func (r ValueRunner) Run() string {
	return "ValueRunner"
}

var _ Runner = ValueRunner{}

type PointerRunner struct{}

func (r *PointerRunner) Run() string {
	return "PointerRunner"
}

var _ Runner = (*PointerRunner)(nil)

// Execute вызывает обычное поведение полиморфно через интерфейс без type switch
func Execute(r Runner) string {
	return r.Run()
}

func Describe(v any) string {
	switch v.(type) {
	case *PointerRunner:
		return "PointerRunner"
	case ValueRunner:
		return "ValueRunner"
	default:
		return "unknown"
	}
}
