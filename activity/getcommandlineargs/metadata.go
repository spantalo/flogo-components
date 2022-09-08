package getcommandlineargs

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
}

type Input struct {
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{}
}
func (r *Input) FromMap(values map[string]interface{}) error {
	return nil
}

type Output struct {
	Args interface{} `md:"args"`
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"args": o.Args,
	}
}
func (o *Output) FromMap(values map[string]interface{}) error {
	args, err := coerce.ToArray(values["args"])
	o.Args = args
	return err
}
