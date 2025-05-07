package repositories

import (
	"godp.abdanhafidz.com/models"
)

func BulkCreateProvince(provinces []models.RegionProvince) Repository[[]models.RegionProvince, []models.RegionProvince] {
	repo := Construct[[]models.RegionProvince, []models.RegionProvince](
		provinces,
	)

	Create(repo)
	return *repo
}

func BulkCreateCity(cities []models.RegionCity) Repository[[]models.RegionCity, []models.RegionCity] {
	repo := Construct[[]models.RegionCity, []models.RegionCity](
		cities,
	)

	Create(repo)
	return *repo
}

func GetListProvinces() Repository[models.RegionProvince, []models.RegionProvince] {
	repo := Construct[models.RegionProvince, []models.RegionProvince](
		models.RegionProvince{},
	)

	repo.Transactions(
		Find[models.RegionProvince, []models.RegionProvince],
	)
	return *repo
}

func GetListCitiesByProvinceId(provinceId uint) Repository[models.RegionCity, []models.RegionCity] {
	repo := Construct[models.RegionCity, []models.RegionCity](
		models.RegionCity{
			ProvinceId: provinceId,
		},
	)
	repo.Transactions(
		WhereGivenConstructor[models.RegionCity, []models.RegionCity],
		Find[models.RegionCity, []models.RegionCity],
	)
	return *repo
}
