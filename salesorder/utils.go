package salesorder

func removeField(name string, fields interface{}) interface{} {
	_fields := fields.([]string)
	var _index int

	for index, value := range _fields {
		if name == value {
			_index = index
		}
	}

	copy(_fields[_index:], _fields[_index+1:])
	_fields[len(_fields)-1] = "" // or the zero value of T
	_fields = _fields[:len(_fields)-1]

	return _fields
}
