package modelos

type Mujer struct {
	Hombre
}

func (m *Mujer) Respirar() bool { return true }
func (m *Mujer) Comer() bool    { return false }
func (m *Mujer) Pensar() bool   { return true }
func (m *Mujer) Sexo() string   { return "Mujer" }
