package rfms

// Known rFMS OEMs.
const (
	DAF         = "DAF"
	Daimler     = "DAIMLER"
	Irisbus     = "IRISBUS"
	Iveco       = "IVECO"
	IvecoBuses  = "IVECO BUS"
	MAN         = "MAN"
	Renault     = "RENAULT TRUCKS"
	Scania      = "SCANIA"
	VDL         = "VDL"
	VolvoBuses  = "VOLVO BUSES"
	VolvoTrucks = "VOLVO TRUCKS"
)

// Known rFMS base URLs.
const (
	ScaniaBaseURL      = "https://dataaccess.scania.com/rfms4"
	ScaniaAuthBaseURL  = "https://dataaccess.scania.com/auth"
	VolvoTrucksBaseURL = "https://api.volvotrucks.com/rfms"
)
