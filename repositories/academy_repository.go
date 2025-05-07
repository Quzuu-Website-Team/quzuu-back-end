package repositories

import "godp.abdanhafidz.com/models"

func GetAllAcademy() Repository[models.Academy, []models.Academy] {
	repo := Construct[models.Academy, []models.Academy](
		models.Academy{},
	)
	repo.Transactions(
		WhereGivenConstructor[models.Academy, []models.Academy],
		Find[models.Academy, []models.Academy],
	)
	return *repo
}

func GetAcademyDataBySlug(slug string) Repository[models.Academy, models.Academy] {
	repo := Construct[models.Academy, models.Academy](
		models.Academy{Slug: slug},
	)
	repo.Transactions(
		WhereGivenConstructor[models.Academy, models.Academy],
		Find[models.Academy, models.Academy],
	)
	return *repo
}

func GetAllAcademyMaterialsByAcademyId(acaddemyId uint) Repository[models.AcademyMaterial, []models.AcademyMaterial] {
	repo := Construct[models.AcademyMaterial, []models.AcademyMaterial](
		models.AcademyMaterial{AcademyId: acaddemyId},
	)
	repo.Transactions(
		WhereGivenConstructor[models.AcademyMaterial, []models.AcademyMaterial],
		Find[models.AcademyMaterial, []models.AcademyMaterial],
	)
	return *repo
}

func GetAllAcademyContentsByMaterialID(materialId uint) Repository[models.AcademyContent, []models.AcademyContent] {
	repo := Construct[models.AcademyContent, []models.AcademyContent](
		models.AcademyContent{AcademyMaterialId: materialId},
	)
	repo.Transactions(
		WhereGivenConstructor[models.AcademyContent, []models.AcademyContent],
		Find[models.AcademyContent, []models.AcademyContent],
	)
	return *repo
}

func CreateAcademy(academies models.Academy) Repository[models.Academy, models.Academy] {
	repo := Construct[models.Academy, models.Academy](
		academies,
	)

	Create(repo)
	return *repo
}

func CreateAcademyMaterial(academyMaterial models.AcademyMaterial) Repository[models.AcademyMaterial, models.AcademyMaterial] {
	repo := Construct[models.AcademyMaterial, models.AcademyMaterial](
		academyMaterial,
	)

	Create(repo)
	return *repo
}

func CreateAcademyContent(academyContent models.AcademyContent) Repository[models.AcademyContent, models.AcademyContent] {
	repo := Construct[models.AcademyContent, models.AcademyContent](
		academyContent,
	)
	Create(repo)
	return *repo
}
