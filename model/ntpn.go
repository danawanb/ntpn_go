package model

type NTPN struct {
	KodeBilling string `json:"kode_billing"`
	Ntpn        string `json:"ntpn"`
	Nilai       string `json:"nilai"`
	Akun        string `json:"akun"`
	Ket         string `json:"ket"`
}

type MPNCookie struct {
	Name  string
	Value string
}
