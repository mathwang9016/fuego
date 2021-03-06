package fuego

// PanicNoSuchElement signifies that the iterator does not have
// the requested element.
const PanicNoSuchElement = "No such element"

// A Maybe is a maybe monad.
type Maybe struct {
	value   Entry
	isEmpty bool
}

// MaybeNone is a Maybe that does not have a value.
func MaybeNone() Maybe {
	return Maybe{
		value:   nil,
		isEmpty: true,
	}
}

// MaybeSome creates a new Maybe with the given value.
// Note: MaybeOf(nil) == None() whereas MaybeSome(nil) == MaybeSome(nil).
func MaybeSome(i Entry) Maybe {
	return Maybe{
		value:   i,
		isEmpty: false,
	}
}

// MaybeOf creates a new Maybe with the given value.
// If the value is nil then return None otherwise Some(value).
// Note: MaybeOf(nil) == None() whereas MaybeSome(nil) == MaybeSome(nil).
func MaybeOf(i Entry) Maybe {
	if i == nil || i.Equal(nil) {
		return MaybeNone()
	}
	return MaybeSome(i)
}

// IsEmpty returns true when this Maybe does not have
// a value.
func (m Maybe) IsEmpty() bool { return m.isEmpty }

// Filter    \
// Map        > TODO: can we use Stream for those?
// FlatMap   /

// Get the value of this Maybe or panic if none exists.
func (m Maybe) Get() Entry {
	if m.IsEmpty() {
		panic(PanicNoSuchElement) // TODO: return MaybeNone??
	}
	return m.value
}

// GetOrElse gets the value of this Maybe or the given Entry if none exists.
func (m Maybe) GetOrElse(e Entry) Entry {
	if m.IsEmpty() {
		return e
	}
	return m.value
}

// OrElse returns this Maybe or the given Maybe if this Maybe is empty.
func (m Maybe) OrElse(other Maybe) Maybe {
	if m.IsEmpty() {
		return other
	}
	return m
}
