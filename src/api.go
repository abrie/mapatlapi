package mapatlapi

import "context"
import "github.com/abrie/mapatlapi/internal/geocoder"
import "github.com/abrie/mapatlapi/internal/record"

func SearchByAddress(address string) (*geocoder.Result, error) {
	request := geocoder.Request{Address: address}
	submitter := geocoder.ProductionSubmitter{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return submitter.SubmitWithContext(ctx, &request)
}

func FetchRecord(refId int) (*record.Result, error) {
	request := record.Request{Ref_ID: refId}
	submitter := record.ProductionSubmitter{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return submitter.SubmitWithContext(ctx, &request)
}
