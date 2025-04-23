package repositories

import (
	"godp.abdanhafidz.com/models"
)

func CreateOptionCategory(categories models.OptionCategory) Repository[models.OptionCategory, models.OptionCategory] {
	repo := Construct[models.OptionCategory, models.OptionCategory](
		categories,
	)
	Create(repo)
	return *repo
}

func CreateOptionValues(values models.OptionValues) Repository[models.OptionValues, models.OptionValues] {
	repo := Construct[models.OptionValues, models.OptionValues](
		values,
	)
	Create(repo)
	return *repo
}

func GetOptionCategoryBySlug(slug string) Repository[models.OptionCategory, models.OptionCategory] {
	repo := Construct[models.OptionCategory, models.OptionCategory](
		models.OptionCategory{OptionSlug: slug},
	)
	repo.Transactions(
		WhereGivenConstructor[models.OptionCategory, models.OptionCategory],
		Find[models.OptionCategory, models.OptionCategory],
	)
	return *repo
}

func GetOptionValuesByCategoryId(categoryId uint) Repository[models.OptionValues, []models.OptionValues] {
	repo := Construct[models.OptionValues, []models.OptionValues](
		models.OptionValues{OptionCategoryId: categoryId},
	)
	repo.Transactions(
		WhereGivenConstructor[models.OptionValues, []models.OptionValues],
		Find[models.OptionValues, []models.OptionValues],
	)
	return *repo
}
