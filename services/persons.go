package services

import (
	"friends-api/errors"
	"friends-api/models"
	"log"
)

// GetPerson function
func GetPerson(pq *models.PersonsQuery) (*models.Persons, *errors.RequestError) {
	var result models.Persons
	session, err := getSession()
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		return &result, errors.ErrDatabaseConnection
	}

	err = session.
		DB("friends").
		C("people").
		Find(pq.Find).
		One(&result)

	if err != nil {
		log.Printf("ERROR: %v\n", err)
		log.Printf("Trying to find person: %v\n", pq.Find)
		return &result, errors.ErrDatabaseQuery
	}

	return &result, nil
}

// GetPersons function
func GetPersons(pq *models.PersonsQuery) (*[]models.Persons, *errors.RequestError) {
	var results []models.Persons
	session, err := getSession()
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		return &results, errors.ErrDatabaseConnection
	}

	err = session.
		DB("friends").
		C("people").
		Find(pq.Find).
		Select(pq.Include).
		Sort(pq.OrderBy...).
		Limit(pq.Limit).
		All(&results)

	if err != nil {
		log.Printf("ERROR: %v\n", err)
		return &results, errors.ErrDatabaseQuery
	}

	return &results, nil
}
