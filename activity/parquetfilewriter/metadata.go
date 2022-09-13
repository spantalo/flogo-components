package parquetfilewriter

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
}

type Input struct {
	ParquetFile     string        `md:"filename,required"`
	CompressionType string        `md:"CompressionType,required"`
	FileColumns     []interface{} `md:"FileColumns,required"`
	FileContent     []interface{} `md:"FileContent,required"`

	//FileContent object like: {map[string]interface {}{"CELL_ID":"64A2D87"}, map[string]interface {}{"CELL_ID":"BBB2-75A5"},...

}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"filename":        r.ParquetFile,
		"CompressionType": r.CompressionType,
		"FileColumns":     r.FileColumns,
		"FileContent":     r.FileContent,
	}
}
func (r *Input) FromMap(values map[string]interface{}) error {

	var err error
	r.ParquetFile, err = coerce.ToString(values["filename"])
	if err != nil {
		return err
	}

	r.CompressionType, err = coerce.ToString(values["CompressionType"])
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
