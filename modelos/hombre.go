package modelos

type Hombre struct {
	Edad       int
	Altura     float32
	Peso       float32
	Respirando bool
	Pensando   bool
	Comiendo   bool
	Vivo       bool
}

func (h *Hombre) Respirar() bool { return false }
func (h *Hombre) Comer() bool    { return true }
func (h *Hombre) Pensar() bool   { return false }
func (h *Hombre) Sexo() string   { return "Hombre" }
