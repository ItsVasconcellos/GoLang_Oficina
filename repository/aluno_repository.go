package repository

import (
	"github.com/ItsVasconcellos/oficina/database"
	"github.com/ItsVasconcellos/oficina/model"
)

func InsertUsuario(usuario model.Usuario) error {
	db := database.GrabDB()

	statement := db.MustBegin()

	statement.MustExec("INSERT INTO usuarios VALUES($1,$2)", usuario.Nome, usuario.RA)
	if err := statement.Commit(); err != nil {
		return err
	}

	return nil
}
