package config

// Keep datacenters state
type DCStates struct {
	MCITehran  int
	MCIMashhad int
	MCITabriz  int
	MCIEsfehan int
	MCIShiraz  int
	Afranet    int
	Irancell   int
	Mobinnet   int
	Shatel     int
}

func New() DCStates {
	return DCStates{
		MCITehran:  0,
		MCIMashhad: 0,
		MCITabriz:  0,
		MCIEsfehan: 0,
		MCIShiraz:  0,
		Afranet:    0,
		Irancell:   0,
		Mobinnet:   0,
		Shatel:     0,
	}
}
