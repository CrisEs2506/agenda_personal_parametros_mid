package models

type Agenda struct {
	Id                int
	Nombre            string
	Descripcion       string
	NumeroDocumento   string
	Activo            bool
	FechaCreacion     string
	FechaModificacion string
}
