package parseparquet

import (
	"encoding/json"
	"fmt"

	"github.com/project-flogo/core/activity"
	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/reader"
)

/*
const (
	ivinitRow     = "initrow"
	ivmaxRows     = "maxrows"
	ivparquetFile = "parquetfile"
	ovOutput      = "output"
)*/

// ParseCSVActivity is a stub for your Activity implementation
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

	initRow := ctx.GetInput("initrow").(int)
	maxRows := ctx.GetInput("maxrows").(int)
	parquetFile := ctx.GetInput("filename").(string)

	ctx.Logger().Info("Processing file:" + parquetFile + " - " + fmt.Sprint(initRow) + " - " + fmt.Sprint(maxRows))

	fr, error := local.NewLocalFileReader(parquetFile)
	if error != nil {
		return false, fmt.Errorf("error opening the specified file: %v", error)
	}

	pr, error := reader.NewParquetReader(fr, nil, 4)
	if error != nil {
		return false, fmt.Errorf("error reading the specified file: %v", error)
	}

	num := int(pr.GetNumRows())

	res, error := pr.ReadByNumber(num)
	if error != nil {
		return false, fmt.Errorf("cannot read parquet file: %v", error)
	}

	jsonBs, error := json.Marshal(res)
	if error != nil {
		return false, fmt.Errorf("cannot create JSON output: %v", error)
	}

	fmt.Println("JSON:")
	fmt.Println(string(jsonBs))

	pr.ReadStop()
	fr.Close()

	output := &Output{}
	output.Output = string(jsonBs)

	err = ctx.SetOutputObject(output)
	if err != nil {
		return false, err
	}

	return true, nil

}
