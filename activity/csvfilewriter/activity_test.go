package csvfilewriter

import (
	"testing"

	"encoding/json"
	"fmt"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("Error during execution: %v", r)
		}
	}()

	act := &Activity{}
	tc := test.NewActivityContext(act.Metadata())

	tc.SetInput("filename", "/tmp/fileOut.parquet")
	//tc.SetInput("jsonstring", "[{\"CELL_ID\":\"1\", \"CNTR_CODE\":\"1\", \"LAU_CODE\":\"1\", \"LAU_NAME\":\"1\", \"SUB_LAU_CODE\":\"1\", \"SUB_LAU_NAME\":\"1\"}  ]")

	rec := `
	{
		"CELL_ID":"22",
		"CNTR_CODE":"11",
		"LAU_CODE":"13",
		"LAU_NAME":"14",
		"SUB_LAU_CODE":"15",
		"SUB_LAU_NAME":"16"
	}
`

	tc.SetInput("jsonstring", rec)

	//	tc.SetInput("jsonschema", "{       \"Tag\":\"name=parquet-go-root repetitiontype=OPTIONAL\",       \"Fields\":[		    {\"Tag\":\"name=CELL_ID, type=BYTE_ARRAY, repetitiontype=OPTIONAL\"},		    {\"Tag\":\"name=CNTR_CODE, type=BYTE_ARRAY, repetitiontype=OPTIONAL\"},		    {\"Tag\":\"name=LAU_CODE, type=BYTE_ARRAY, repetitiontype=OPTIONAL\"},		    {\"Tag\":\"name=LAU_NAME, type=BYTE_ARRAY, repetitiontype=OPTIONAL\"},		    {\"Tag\":\"name=SUB_LAU_CODE, type=BYTE_ARRAY, repetitiontype=OPTIONAL\"},	    {\"Tag\":\"name=SUB_LAU_NAME, type=BYTE_ARRAY, repetitiontype=OPTIONAL\"},{\"Tag\":\"name=TS_MOD, type=INT64\"}]}")

	schema := `
	{       "Tag":"name=parquet-go-root repetitiontype=REQUIRED",       
	"Fields":[		    {"Tag":"name=CELL_ID, type=BYTE_ARRAY, repetitiontype=OPTIONAL"},		    {"Tag":"name=CNTR_CODE, type=BYTE_ARRAY, repetitiontype=OPTIONAL"},		    {"Tag":"name=LAU_CODE, type=BYTE_ARRAY, repetitiontype=OPTIONAL"},		    {"Tag":"name=LAU_NAME, type=BYTE_ARRAY, repetitiontype=OPTIONAL"},		    {"Tag":"name=SUB_LAU_CODE, type=BYTE_ARRAY, repetitiontype=OPTIONAL"},	    {"Tag":"name=SUB_LAU_NAME, type=BYTE_ARRAY, repetitiontype=OPTIONAL"},
	{"Tag":"name=TS_MOD, type=INT64 , repetitiontype=OPTIONAL"}]}`

	tc.SetInput("jsonschema", schema)

	done, err := act.Eval(tc)
	assert.True(t, done)
	assert.Nil(t, err)

	//check result attr
	aOutput := &Output{}
	err = tc.GetOutputObject(aOutput)
	assert.Nil(t, err)

	b, _ := json.Marshal(aOutput.Result)
	fmt.Println(string(b))
}
