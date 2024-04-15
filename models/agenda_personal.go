package models

type Contacto struct {
	Id                int
	Nombre            string
	NumeroDocumento   string
	Direccion         string
	Activo            bool
	FechaCreacion     string
	FechaModificacion string
	AgendaId          *Agenda
}
