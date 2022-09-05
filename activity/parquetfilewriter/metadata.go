package parquetfilewriter

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
}

type Input struct {
	ParquetFile string `md:"filename,required"`
	JSONString  string `md:"jsonstring,required"`
	JSONSchema  string `md:"jsonschema"` //,required
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"filename":   r.ParquetFile,
		"jsonstring": r.JSONString,
		"jsonschema": r.JSONSchema,
	}
}
func (r *Input) FromMap(values map[string]interface{}) error {

	var err error
	r.ParquetFile, err = coerce.ToString(values["filename"])
	if err != nil {
		return err
	}
	r.JSONString, err = coerce.ToString(values["jsonstring"])
	if err != nil {
		return err
	}

	r.JSONSchema, err = coerce.ToString(values["jsonschema"])
	if err != nil {
		return err
	}
	return nil
}

type Output struct {
	Result string `md:"result"`
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"result": o.Result,
	}
}
func (o *Output) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["result"])
	o.Result = strVal
	return nil
}
