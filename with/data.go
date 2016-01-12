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
	return fields.ToFields()
}

func (f *Data) With(key string, value interface{}) Data {
	n := f.Copy()
	n[key] = value
	return n
}

func (f *Data) WithAll(data Data) Data {
	n := f.Copy()
	for k, v := range data {
		n[k] = v
	}
	return n
}

func (f *Data) Copy() Data {
	data := Data{}
	for k, v := range (*f) {
		data[k] = v
	}
	return data
}
