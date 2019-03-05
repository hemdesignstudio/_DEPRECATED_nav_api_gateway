package salesinvoice

func removeField(name string, fields interface{}) interface{} {
	_fields := fields.([]string)
	var _index int

	for index, value := range _fields {
		if name == value {
			_index = index
		}
	}

	// Deleting element without caring about order
	//If the type of the element is a pointer or a struct with pointer fields,
	// which need to be garbage collected, the implementations of deleting an element
	// have a potential memory leak problem:
	// some elements with values are still referenced by slice a and
	// thus can not be collected.
	// although we are dealing with a []string,
	// It is still better to be oin the safe side and implement it this way
	copy(_fields[_index:], _fields[_index+1:])
	_fields[len(_fields)-1] = ""
	_fields = _fields[:len(_fields)-1]

	return _fields
}
