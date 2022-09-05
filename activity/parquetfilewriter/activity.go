package parquetfilewriter

import (
	"fmt"

	"github.com/project-flogo/core/activity"
	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/writer"
)

// Activity is a Parquet writer
type Activity struct {
}

func init() {
	_ = activity.Register(&Activity{}, New)
}

var activityMd = activity.ToMetadata(&Input{}, &Output{})

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// New create a new  activity
func New(ctx activity.InitContext) (activity.Activity, error) {
	act := &Activity{}
	return act, nil
}

// Eval implements activity.Activity.Eval
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	// Get the runtime values
	in := &Input{}
	err = ctx.GetInputObject(in)
	if err != nil {
		return false, err
	}

	jsonSchema := in.JSONSchema
	jsonString := in.JSONString
	parquetFile := in.ParquetFile

	ctx.Logger().Infof("Processing file: %s", parquetFile)

	//--

	//write

	fw, err := local.NewLocalFileWriter(parquetFile)
	if err != nil {
		ctx.Logger().Errorf("Can't create file %s", err)
		return false, err
	}

	pw, err := writer.NewJSONWriter(jsonSchema, fw, 4)
	if err != nil {
		ctx.Logger().Errorf("Can't create JSON writer %s", err)
		return false, err
	}

	if err = pw.Write(jsonString); err != nil {
		ctx.Logger().Errorf("Write file Error %s", err)
	}

	fmt.Println(">>>>1")
	fmt.Println(jsonSchema)
	fmt.Println(jsonString)

	if err = pw.WriteStop(); err != nil {
		ctx.Logger().Errorf("Close file Error %s", err)
	}

	fmt.Println(">>>>2")

	fw.Close()

	fmt.Println(">>>>3")

	ctx.Logger().Infof("File completed %s", parquetFile)

	//--
	output := &Output{}
	output.Result = "OK"

	err = ctx.SetOutputObject(output)
	if err != nil {
		return false, err
	}
	return true, nil
}
