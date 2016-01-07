package with

type Data map[string]interface{}

type FieldsConverter interface {
	ToFields() Data
}

func Field(key string, value interface{}) Data {
	i := make(Data, 3)
	return i.With(key, value)
}

func Fields(fields FieldsConverter) Data {
	toFields := fields.ToFields()
	i := make(Data, len(toFields) + 1)
	return i.WithAll(toFields)
}

func (f Data) With(key string, value interface{}) Data {
	f[key] = value
	return f
}

func (f Data) WithAll(data Data) Data {
	for k, v := range data {
		f[k] = v
	}
	return f
}
