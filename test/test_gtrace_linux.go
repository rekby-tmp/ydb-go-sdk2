// Code generated by gtrace. DO NOT EDIT.

// +build linux

package test

// Compose returns a new ConditionalBuildTrace which has functional fields composed
// both from t and x.
func (t ConditionalBuildTrace) Compose(x ConditionalBuildTrace) (ret ConditionalBuildTrace) {
	switch {
	case t.OnSomething == nil:
		ret.OnSomething = x.OnSomething
	case x.OnSomething == nil:
		ret.OnSomething = t.OnSomething
	default:
		h1 := t.OnSomething
		h2 := x.OnSomething
		ret.OnSomething = func() {
			h1()
			h2()
		}
	}
	return ret
}
func (t ConditionalBuildTrace) onSomething() {
	fn := t.OnSomething
	if fn == nil {
		return
	}
	fn()
}
