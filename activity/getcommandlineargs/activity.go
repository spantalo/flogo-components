package getcommandlineargs

import (
	"os"

	"github.com/project-flogo/core/activity"
)

// Activity is a Parquet parser
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

	output := &Output{}
	output.Args = os.Args

	err = ctx.SetOutputObject(output)
	if err != nil {
		return false, err
	}

	return true, nil
}
