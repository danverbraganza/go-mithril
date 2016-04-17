package moria

import "github.com/gopherjs/gopherjs/js"

type Prop struct {
	*js.Object
}

func StringProp(target *string) (getterSetter *js.Object) {
	return js.MakeFunc(func(this *js.Object, args []*js.Object) interface{} {
		if len(args) != 0 && args[0] != nil {
			*target = args[0].String()
		}
		return *target
	})
}

func BoolProp(target *bool) (getterSetter *js.Object) {
	return js.MakeFunc(func(this *js.Object, args []*js.Object) interface{} {
		if len(args) != 0 && args[0] != nil {
			*target = args[0].Bool()
		}
		return *target
	})
}

// TODO(danver): Fill in all the convenience methods here for defining props.
