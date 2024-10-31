package services

import "trocup-message/repository"

func DeleteMessage(id string) error {
	return repository.DeleteMessage(id)
}
