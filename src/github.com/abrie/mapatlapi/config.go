package mapatlapi

type Config struct {
	GeocoderEndpoint string
	LocationEndpoint string
	PlacesEndpoint   string
}

func New() Config {
	return Config{
		GeocoderEndpoint: "https://egis.atlantaga.gov/arc/rest/services/WebLocators/TrAddrPointS/GeocodeServer/findAddressCandidates",
		LocationEndpoint: "http://egis.atlantaga.gov/app/home/php/egisws.php",
		PlacesEndpoint:   "http://egis.atlantaga.gov/app/home/php/egispoi.php"}
}
