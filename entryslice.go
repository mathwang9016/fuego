package fuego

// EntrySlice is an Entry for '[]Entry'.
type EntrySlice []Entry

// TODO: implement Stream() (see EntryMap)

// Hash returns a hash for this Entry.
func (es EntrySlice) Hash() uint32 {
	if es == nil || len(es) == 0 {
		return 0
	}

	result := uint32(1)

	for _, element := range es {
		var h uint32
		if element != nil {
			h = element.Hash()
		}
		result = 31*result + h
	}

	return result
}

// Equal returns true if this type is equal to 'e'.
func (es EntrySlice) Equal(e Entry) bool {
	if _, ok := e.(EntrySlice); !ok {
		return false
	}

	return es.Hash() == e.Hash()
}
