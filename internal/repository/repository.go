package repository

type EntityRepository interface {
	findOneById(id int)
	findAll()
}
