package parquetfilewriter

import (
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"
	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/parquet"
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

	parquetFile := in.ParquetFile
	FileColumns := in.FileColumns
	FileContent := in.FileContent
	CompressionType := in.CompressionType

	ctx.Logger().Infof("Processing file: %s", parquetFile)

	if FileColumns != nil {

		//Read mapped values
		ctx.Logger().Debugf("FileColumns: %s", FileColumns)
		ctx.Logger().Debugf("FileContent: %s", FileContent) //è un array di mappe

		md := []string{}
		for j := 0; j < len(FileColumns); j++ {
			col := FileColumns[j].(map[string]interface{})
			md = append(md, "name="+col["Name"].(string)+", type="+col["Type"].(string))
		}

		ctx.Logger().Debugf("Model: %s", md)

		//creazione del file
		fw, err := local.NewLocalFileWriter(parquetFile)
		if err != nil {
			ctx.Logger().Errorf("Can't create file %s", err)
			return false, err
		}

		pw, err := writer.NewCSVWriter(md, fw, 4)
		if err != nil {
			ctx.Logger().Errorf("Can't create file writer %s", err)
			return false, err
		}

		switch CompressionType {
		case "UNCOMPRESSED":
			pw.CompressionType = parquet.CompressionCodec_UNCOMPRESSED
		case "SNAPPY":
			pw.CompressionType = parquet.CompressionCodec_SNAPPY
		case "GZIP":
			pw.CompressionType = parquet.CompressionCodec_GZIP
		case "LZ4":
			pw.CompressionType = parquet.CompressionCodec_LZ4
		case "ZSTD":
			pw.CompressionType = parquet.CompressionCodec_ZSTD

		}

		if FileContent != nil {
			for i := 0; i < len(FileContent); i++ {
				row := FileContent[i].(map[string]interface{})

				el := []any{}
				for j := 0; j < len(FileColumns); j++ {
					col := FileColumns[j].(map[string]interface{})
					name := col["Name"].(string)

					var val any
					var err error

					switch col["Type"].(string) {
					case "BYTE_ARRAY":
						val, err = coerce.ToString(row[name])
					case "INT32":
						val, err = coerce.ToInt32(row[name])
					case "INT64":
						val, err = coerce.ToInt64(row[name])
					case "FLOAT":
						val, err = coerce.ToFloat32(row[name])
					case "BOOLEAN":
						val, err = coerce.ToBool(row[name])
					case "DOUBLE":
						val, err = coerce.ToFloat64(row[name])
					}

					if err != nil {
						return false, err
					}
					el = append(el, val)
				}
				//ctx.Logger().Warnf("Writing: %s", el)
				if err = pw.Write(el); err != nil {
					ctx.Logger().Warnf("Row: %s", row)
					ctx.Logger().Warnf("FileColumns: %s", FileColumns)
					ctx.Logger().Errorf("Write file Error %s", err)
					return false, err
				}
			}
		}
		if err = pw.WriteStop(); err != nil {
			ctx.Logger().Errorf("Close file Error %s", err)
		}
		fw.Close()
	}

	ctx.Logger().Infof("File completed %s", parquetFile)

	output := &Output{}
	output.Result = "OK"

	err = ctx.SetOutputObject(output)
	if err != nil {
		return false, err
	}
	return true, nil
}
