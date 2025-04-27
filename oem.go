package rfms

// Known rFMS OEMs.
const (
	Scania      = "SCANIA"
	VolvoTrucks = "VOLVO TRUCKS"
	Daimler     = "DAIMLER"
	Iveco       = "IVECO"
	Daf         = "DAF"
	Man         = "MAN"
	Renault     = "RENAULT TRUCKS"
	VDL         = "VDL"
	VolvoBuses  = "VOLVO BUSES"
	IvecoBuses  = "IVECO BUS"
	Irisbus     = "IRISBUS"
)

// Known rFMS base URLs.
const (
	ScaniaBaseURL      = "https://dataaccess.scania.com/rfms4"
	ScaniaAuthBaseURL  = "https://dataaccess.scania.com/auth"
	VolvoTrucksBaseURL = "https://api.volvotrucks.com/rfms"
)
