package services

import (
	"github.com/gosimple/slug"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/repositories"
)

type OptionService struct {
	Service[[]models.OptionsRequest, models.OptionsResponse]
}

type OptionValueService struct {
	Service[models.OptionCategory, models.Options]
}

func (s *OptionService) Create() {
	optionsResult := []models.Options{}
	for _, option := range s.Constructor {
		OptionCategoryRepo := repositories.CreateOptionCategory(
			models.OptionCategory{
				OptionName: option.OptionName,
				OptionSlug: slug.Make(option.OptionName),
			},
		)
		if OptionCategoryRepo.RowsError != nil {
			OptionCategoryRepo.Transaction.Rollback()
			s.Error = OptionCategoryRepo.RowsError
			return
		}
		optionsValueResult := []models.OptionValues{}
		for _, value := range option.OptionValue {
			OptionValuesRepo := repositories.CreateOptionValues(
				models.OptionValues{
					OptionCategoryId: OptionCategoryRepo.Result.Id,
					OptionValue:      value,
				},
			)

			if OptionValuesRepo.RowsError != nil {
				OptionValuesRepo.Transaction.Rollback()
				s.Error = OptionValuesRepo.RowsError
				return
			}
			optionsValueResult = append(optionsValueResult, OptionValuesRepo.Result)
		}
		optionsResult = append(optionsResult, models.Options{OptionCategory: OptionCategoryRepo.Result, OptionValues: optionsValueResult})
	}
	s.Result = models.OptionsResponse{
		Options: optionsResult,
	}
}

func (s *OptionValueService) Retrieve() {
	optionCategory := repositories.GetOptionCategoryBySlug(s.Constructor.OptionSlug)
	if optionCategory.RowsError != nil {
		s.Error = optionCategory.RowsError
		return
	}
	optionValues := repositories.GetOptionValuesByCategoryId(optionCategory.Result.Id)
	if optionValues.RowsError != nil {
		s.Error = optionValues.RowsError
		return
	}
	s.Result = models.Options{
		OptionCategory: optionCategory.Result,
		OptionValues:   optionValues.Result,
	}

}
