package astreams

import "reflect"

// ActivityStreamer can be either a (sub)type of an 'Object' or 'Link'
type ActivityStreamer interface {
	IsObject() bool
	IsLink() bool
	GetObject() *Object
	GetLink() *Link
}

// ConcreteType returns both, the type name obtained using reflection,
// as well as the Type property of the Object/Link.
// The object's own Type property is going to be more specific, so use that where useful.
func ConcreteType(t ActivityStreamer) (reflectType, astreamsType string) {
	if t.IsLink() {
		return reflect.TypeOf(t).Name(), t.GetLink().Type
	}
	return reflect.TypeOf(t).Name(), t.GetObject().Type
}

// Implements 'ActivityStreamer' interface for 'Object'
func (o Object) IsObject() bool {
	return true
}

func (o Object) IsLink() bool {
	return false
}

func (o Object) GetObject() *Object {
	return &o
}

func (o Object) GetLink() *Link {
	return nil
}

// Implements 'ActivityStreamer' interface for 'Link'
func (l Link) IsObject() bool {
	return false
}

func (l Link) IsLink() bool {
	return true
}

func (l Link) GetObject() *Object {
	return nil
}

func (l Link) GetLink() *Link {
	return &l
}
