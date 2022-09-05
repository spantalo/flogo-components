package parseparquet

import (
	"encoding/json"
	"fmt"

	"github.com/project-flogo/core/activity"
	//"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"

	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/reader"
)

// ActivityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-flogo-parseparquet")

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

	activityLog.Debugf("Processing file: %s, [%s-%s] ", parquetFile, initRow, maxRows)

	//--
	/* 	var file, error = os.OpenFile(parquetFile, os.O_RDONLY, 0)
	   	if error != nil {

	   		return false, fmt.Errorf("error opening the specified file: %v", error)
	   	} */

	// - ultima versione

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
		return false, fmt.Errorf("Cannot read parquet file: %v", error)
	}

	jsonBs, error := json.Marshal(res)
	if error != nil {
		return false, fmt.Errorf("Cannot create JSON output: %v", error)
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

	// -------

	//--
	/* test con statistiche
	stt, errStats := file.Stat()
	if errStats != nil {
		return false, fmt.Errorf("error opening the specified file: %v", errStats)
	}

	size := stt.Size()

	f, errParquetFile := parquet.OpenFile(file, size)
	if errParquetFile != nil {
		return false, fmt.Errorf("error opening the specified file: %v", errParquetFile)
	}
	*/
	//	fmt.Println(f.Schema().Columns())
	//--

	//	reader := parquet.NewReader(file)

	// OK VECCHIA

	/* 	reader := parquet.NewGenericReader[any](file)

	   	if initRow > 0 {
	   		reader.SeekToRow(int64(initRow))
	   	}

	   	rows := []any{}
	   	for {

	   		if len(rows) >= maxRows {
	   			break
	   		}

	   		//row := []any{}
	   		row := make([]any, 1)

	   		_, err := reader.Read(row)

	   		fmt.Println("ROW:")
	   		fmt.Println(row)

	   		if err != nil {
	   			if err == io.EOF {
	   				break
	   			}
	   		}

	   		//rows = append(rows, row)
	   	}

	   	//-
	   	activityLog.Debugf("Parsed Object from parquetFile: %s", rows)

	   	fmt.Println("ROWS")
	   	fmt.Println(rows)

	   	ctx.SetOutput(ovOutput, rows)

	   	return true, nil */

	// fine ok vecchia

	//--

	//var reader io.Reader

	//	if txt, ok := ctx.GetInput(ivCSV).(string); ok && len(txt) > 0 {
	//		reader = strings.NewReader(txt)
	//	} else

	/*
		if file, ok := ctx.GetInput(ivparquetFile).(string); ok {
			osFile, err := os.Open(file)
			if err != nil {
				return false, fmt.Errorf("error opening the specified file: %v", err)
			}
			reader = bufio.NewReader(osFile)
		} else {
			return false, fmt.Errorf("A parquetFilename must be supplied")
		}


		r := csv.NewReader(reader)
	*/

	//r.LazyQuotes = true
	//r.Comma = comma
	//r.Comment = '#'

	/*
		 	for {
				record, err := r.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					activityLog.Errorf("Failed to read csv string: %s", err)
					return false, err
				}

				if len(record) != len(fieldNames) {
					activityLog.Error("Mismatch between number of fields and field names specified")
					return false, fmt.Errorf("Fields supplied do not match total fields in csv. Expected %d but got %d", len(fieldNames), len(record))
				}

				field := make(map[string]interface{})

				for i := 0; i < len(record); i++ {
					field[fieldNames[i].(string)] = record[i]
				}

				obj = append(obj, field)
			}

			activityLog.Debugf("Parsed Object from parquetFile: %s", obj)
			ctx.SetOutput(ovOutput, obj)

			return true, nil
	*/
}

//----------- ORI

/*func init() {
	_ = activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}
*/

// New optional factory method, should be used if one activity instance per configuration is desired
/*func New(ctx activity.InitContext) (activity.Activity, error) {
}
*/

// Eval implements api.Activity.Eval - Logs the Message
/* func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}

	ctx.Logger().Debugf("Input: %s", input.AnInput)

	output := &Output{AnOutput: input.AnInput}
	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}
*/
