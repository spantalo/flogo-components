package filemanager

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	Action string `md:"action,required"`
}

type Input struct {
	File string `md:"file,required,allowed(COPY, MOVE, LIST, DELETE, MKDIR)"`
	To   string `md:"to"`
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"file": r.File,
		"to":   r.To,
	}
}
func (r *Input) FromMap(values map[string]interface{}) error {

	var err error
	r.File, err = coerce.ToString(values["file"])
	if err != nil {
		return err
	}
	r.To, err = coerce.ToString(values["to"])
	if err != nil {
		return err
	}
	return nil
}

type Output struct {
	Result string        `md:"result"`
	Error  string        `md:"error"`
	Files  []interface{} `md:"files"`
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"result": o.Result,
		"error":  o.Error,
		"files":  o.Files,
	}
}
func (o *Output) FromMap(values map[string]interface{}) error {

	var err error
	o.Result, err = coerce.ToString(values["result"])
	if err != nil {
		return err
	}
	o.Error, err = coerce.ToString(values["error"])
	if err != nil {
		return err
	}
	o.Files, err = coerce.ToArray(values["files"])
	if err != nil {
		return err
	}

	return nil
}
