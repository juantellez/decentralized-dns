package dnsclient

var reservedTLDs = map[string]bool{
	"com": true,
	"net": true,
	"org": true,
	"edu": true,
	"gov": true,
	"mil": true,
	"mx":  true,
	"us":  true,
	"uk":  true,
	"io":  true,
}
