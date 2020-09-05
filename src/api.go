package mapatlapi

import "github.com/abrie/mapatlapi/internal/geocoder"
import "github.com/abrie/mapatlapi/internal/record"

func SearchByAddress(address string) (*geocoder.Result, error) {
	request := geocoder.Request{Address: address}
	submitter := geocoder.ProductionSubmitter{}
	return submitter.Submit(&request)
}

func FetchRecord(refId int) (*record.Result, error) {
	request := record.Request{Ref_ID: refId}
	submitter := record.ProductionSubmitter{}
	return submitter.Submit(&request)
}
