package parquetfilewriter

import (
	"encoding/json"
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

	ctx.Logger().Infof("Processing file: %s", parquetFile)

	// Get the runtime values
	in := &Input{}
	err = ctx.GetInputObject(in)
	if err != nil {
		return false, err
	}

	jsonSchema := in.JSONSchema
	jsonString := in.JSONString
	parquetFile := in.ParquetFile

	if ctx.GetInput("ContentJson") != nil {
		messageJSONObj := ctx.GetInput("ContentJson").(*data.ComplexObject)
		buffer, err := json.Marshal(messageJSONObj.Value)
		if err != nil {
			ctx.Logger().Errorf("Failed to decode map input for reason [%s]", err)
			return false, fmt.Errorf("Failed to decode map input for reason [%s]", err)
		}
		jsonString = string(buffer)
	}

	//ctx.Logger().Debugf("jsonSchema %s", jsonSchema)
	ctx.Logger().Infof("jsonString: %s", jsonString)

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

	if err = pw.WriteStop(); err != nil {
		ctx.Logger().Errorf("Close file Error %s", err)
	}

	fw.Close()

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
