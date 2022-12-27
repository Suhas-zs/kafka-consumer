package transformers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/google/uuid"
	"github.int.mcafee.com/mcafee/bulk-processing-framework/pkg/models"
)

const (
	TEN_LAKH = 1000000
)

type transformer struct{}

func New() transformer {
	return transformer{}
}

func (t transformer) CustomTransformer(ctx *gofr.Context, record models.Record) (*models.Record, error) {
	ctx.Logger.Debugf("This is the record before transformation: %v", record)

	conversionTime(&record)
	record.Value["correlation_id"] = uuid.New()

	ctx.Logger.Debugf("This is the record after transformation: %v", record)
	return &record, nil
}

func conversionTime(record *models.Record) {
	if record.Value["creation_time"] != nil {
		val, ok := record.Value["creation_time"].(string)
		if ok {
			timer, _ := strconv.Atoi(val)
			record.Value["crdt"] = parseTime(time.UnixMilli(int64(timer)).UTC())
		}
	}

	if record.Value["changed_time"] != nil {
		val, ok := record.Value["changed_time"].(string)
		if ok {
			timer, _ := strconv.Atoi(val)
			record.Value["chdt"] = parseTime(time.UnixMilli(int64(timer)).UTC())
		}
	}

	record.Value["hit_time"] = parseTime(time.Now().UTC())
}

func parseTime(t time.Time) string {
	newString := fmt.Sprintf("%d.%dZ", t.Second(), t.Nanosecond()/TEN_LAKH)
	respTime := strings.Replace(t.Format(time.RFC3339), strconv.Itoa(t.Second())+"Z", newString, 1)

	return respTime
}
