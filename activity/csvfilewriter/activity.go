package csvfilewriter

import (
	"bufio"
	"compress/gzip"
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/metadata"
)

// Activity is a CSV writer
type Activity struct {
	settings *Settings
}

func init() {
	_ = activity.Register(&Activity{}, New)
}

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// New create a new  activity
func New(ctx activity.InitContext) (activity.Activity, error) {
	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Debugf("Settings separator: %s", s.Separator)

	act := &Activity{}
	act.settings = s
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

	Filename := in.Filename
	FileColumns := in.FileColumns
	FileContent := in.FileContent

	ctx.Logger().Infof("Processing file: %s", Filename)

	if FileColumns != nil {

		//Read mapped values
		ctx.Logger().Debugf("FileColumns: %s", FileColumns)
		ctx.Logger().Debugf("FileContent: %s", FileContent) //è un array di mappe

		// create file
		f, err := os.Create(Filename)
		if err != nil {
			log.Fatal(err)
		}
		// remember to close the file
		defer f.Close()

		// create new buffer
		var writer io.Writer

		if a.settings.Compress {
			awriter := gzip.NewWriter(f)
			defer awriter.Close()
			writer = awriter

		} else {
			awriter := bufio.NewWriter(f)
			defer awriter.Flush()
			writer = awriter
		}

		w := csv.NewWriter(writer)
		defer w.Flush()

		switch a.settings.Separator {
		case "PIPE":
			w.Comma = '|'
		case ",":
			w.Comma = ','
		case ";":
			w.Comma = ';'
		case "TAB":
			w.Comma = '\t'

		}

		//--
		if FileContent != nil {
			for i := 0; i < len(FileContent); i++ {
				row := FileContent[i].(map[string]interface{})
				record := []string{}

				for j := 0; j < len(FileColumns); j++ {
					col := FileColumns[j].(map[string]interface{})
					name := col["Name"].(string)

					var val string
					var err error

					val, err = coerce.ToString(row[name])
					if err != nil {
						return false, err
					}

					record = append(record, val)

				}

				if err := w.Write(record); err != nil {
					ctx.Logger().Errorf("error writing record to csv: %s", err)
					return false, err

				}
			}
		}

		//--

		if err := w.Error(); err != nil {
			ctx.Logger().Errorf("Write file Error %s", err)
			return false, err
		}

		//--

	}
	ctx.Logger().Infof("File completed %s", Filename)

	return true, nil
}
