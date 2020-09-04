package mapatlapi

import "github.com/abrie/mapatlapi/internal/geocoder"

func SearchByAddress(address string) (*geocoder.Result, error) {
	request := geocoder.Request{Address: address}
	submitter := geocoder.ProductionSubmitter{}
	return submitter.Submit(&request)
}
