package csvfilewriter

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	Separator string `md:"Separator,required"`
	Compress  bool   `md:"Compress,required"`
}

type Input struct {
	Filename    string        `md:"Filename,required"`
	FileColumns []interface{} `md:"FileColumns,required"`
	FileContent []interface{} `md:"FileContent,required"`

	//FileContent object like: {map[string]interface {}{"CELL_ID":"64A2D87"}, map[string]interface {}{"CELL_ID":"BBB2-75A5"},...

}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Filename":    r.Filename,
		"FileColumns": r.FileColumns,
		"FileContent": r.FileContent,
	}
}

func (r *Input) FromMap(values map[string]interface{}) error {

	var err error
	r.Filename, err = coerce.ToString(values["Filename"])
	if err != nil {
		return err
	}

	r.FileColumns, err = coerce.ToArray(values["FileColumns"])
	if err != nil {
		return err
	}

	r.FileContent, err = coerce.ToArray(values["FileContent"])
	if err != nil {
		return err
	}

	return nil
}

type Output struct {
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{}
}
func (o *Output) FromMap(values map[string]interface{}) error {
	return nil
}
