package filemanager

import (
	"io"
	"os"
	"time"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
)

type Activity struct {
	settings *Settings
	logger   log.Logger
}

/*
func init() {
	_ = activity.Register(&Activity{}, New)
}*/

type ListFile struct {
	Name    string
	Size    int64
	ModTime time.Time
	IsDir   bool
}

var activityMd = activity.ToMetadata(&Input{}, &Output{})

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

	act := &Activity{logger: ctx.Logger(), settings: s}
	return act, nil
}

// Eval implements activity.Activity.Eval
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return false, err
	}

	output := &Output{}

	file := input.File
	to := input.To
	action := a.settings.Action

	switch action {
	case "COPY":

		// Open original file
		originalFile, err := os.Open(file)
		if err != nil {
			a.logger.Errorf("Cannot %s file %s to %s error:", action, file, to, err.Error())
			output.Result = "KO"
			output.Error = err.Error()
			break
		}
		defer originalFile.Close()

		// Create new file
		newFile, err := os.Create(to)
		if err != nil {
			a.logger.Errorf("Cannot %s file %s to %s error:", action, file, to, err.Error())
			output.Result = "KO"
			output.Error = err.Error()
			break
		}
		defer newFile.Close()

		// Copy the bytes to destination from source
		bytesWritten, err := io.Copy(newFile, originalFile)
		if err != nil {
			a.logger.Errorf("Cannot %s file %s to %s error:", action, file, to, err.Error())
			output.Result = "KO"
			output.Error = err.Error()
			break
		}
		a.logger.Debugf("Copied %s file size: %d", file, bytesWritten)

		// Commit the file contents
		// Flushes memory to disk
		err = newFile.Sync()
		if err != nil {
			a.logger.Errorf("Cannot %s file %s to %s error:", action, file, to, err.Error())
			output.Result = "KO"
			output.Error = err.Error()
			break
		}
		output.Result = "OK"

	case "MOVE":
		err := os.Rename(file, to)
		if err != nil {
			a.logger.Errorf("Cannot %s file %s to %s error:", action, file, to, err.Error())
			output.Result = "KO"
			output.Error = err.Error()
			break
		}
		output.Result = "OK"

	case "LIST":

		// Open original file
		originalFile, err := os.Open(file)
		if err != nil {
			a.logger.Errorf("Cannot %s file %s error:", action, file, err.Error())
			output.Result = "KO"
			output.Error = err.Error()
			break
		}
		defer originalFile.Close()

		files, err := originalFile.Readdir(0)
		if err != nil {
			a.logger.Errorf("Cannot %s file %s error:", action, file, err.Error())
			output.Result = "KO"
			output.Error = err.Error()
			break
		}

		var lista []interface{}

		for _, info := range files {
			lista = append(lista, ListFile{
				info.Name(),
				info.Size(),
				info.ModTime(),
				info.IsDir(),
			})
		}

		output.Result = "OK"
		output.Files = lista

	case "DELETE":
		err := os.Remove(file)
		if err != nil {
			a.logger.Errorf("Cannot %s file %s error:", action, file, err.Error())
			output.Result = "KO"
			output.Error = err.Error()
			break
		}
		output.Result = "OK"
	}

	err = ctx.SetOutputObject(output)
	if err != nil {
		return false, err
	}

	return true, err
}
