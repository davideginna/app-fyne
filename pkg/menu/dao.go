package menu

import (
	"time"
)

type (
	Menu struct {
		Referente  utente
		DataOrdine time.Time
		Ordine     ordine
	}

	utente struct {
		Nome     string
		Cognome  string
		Telefono string
	}

	ordine struct {
		Antipasto []portata
		Primo     []portata
		Secondo   []portata
		Dolce     []portata
		Bevanda   []bevanda
	}

	portata struct {
		Nome   string
		Prezzo float32
		Qt     int
		Note   string
	}

	bevanda struct {
		Nome   string
		Prezzo float32
		Qt     int
		Note   string
		Ml     int
	}
)
