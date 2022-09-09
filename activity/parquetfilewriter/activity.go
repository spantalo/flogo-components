package parquetfilewriter

import (
	"encoding/json"
	"fmt"

	//"github.com/project-flogo/core/activity"
	//"github.com/project-flogo/core/data"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/logger"

	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/writer"
)

var log = logger.GetLogger("flogo-parquet-filewriter")

// PutActivity is a stub for your Activity implementation
type ParquetFileWriterActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	//TODO REMOVE
	//log.SetLogLevel(logger.DebugLevel)
	//
	return &ParquetFileWriterActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *ParquetFileWriterActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *ParquetFileWriterActivity) Eval(ctx activity.Context) (done bool, err error) {

	// Get the runtime values
	jsonSchema := ctx.GetInput("jsonschema").(string)
	jsonString := "-" //ctx.GetInput("jsonstring").(string)
	parquetFile := ctx.GetInput("filename").(string)
	log.Infof("### Processing file: %s", parquetFile)

	if ctx.GetInput("ContentJson") != nil {
		messageJSONObj := ctx.GetInput("ContentJson").(*data.ComplexObject)
		log.Infof("### messageJSONObj Val: %s", messageJSONObj.Value)
		log.Infof("### messageJSONObj Schema: %s", messageJSONObj.Metadata)

		buffer, err := json.Marshal(messageJSONObj.Value)
		if err != nil {
			log.Errorf("Failed to decode map input for reason [%s]", err)
			return false, fmt.Errorf("failed to decode map input for reason [%s]", err)
		}
		log.Infof("### Buffer: %s", buffer)
		jsonString = string(buffer)

	}

	//log.Debugf("jsonSchema %s", jsonSchema)
	log.Infof("### jsonString: %s", jsonString)

	//--

	//write

	fw, err := local.NewLocalFileWriter(parquetFile)
	if err != nil {
		log.Errorf("Can't create file %s", err)
		return false, err
	}

	pw, err := writer.NewJSONWriter(jsonSchema, fw, 4)
	if err != nil {
		log.Errorf("Can't create JSON writer %s", err)
		return false, err
	}

	if err = pw.Write(jsonString); err != nil {
		log.Errorf("Write file Error %s", err)
	}

	if err = pw.WriteStop(); err != nil {
		log.Errorf("Close file Error %s", err)
	}

	fw.Close()

	log.Infof("File completed %s", parquetFile)

	ctx.SetOutput("result", "OK")

	if err != nil {
		return false, err
	}
	return true, nil
}
