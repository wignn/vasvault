package services

import (
	"errors"
	"vasvault/internal/models"
	"vasvault/internal/repositories"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) List(userID uint, search string) ([]models.Category, error) {
	return s.repo.List(userID, search)
}

func (s *CategoryService) GetByID(id uint) (*models.Category, error) {
    category, err := s.repo.GetByID(id)
    if err != nil {
        return nil, errors.New("category not found")
    }
    return category, nil
}

func (s *CategoryService) Detail(userID uint, id uint) (*models.Category, error) {
    category, err := s.repo.GetByIDAndUser(userID, id)
    if err != nil {
        return nil, errors.New("category not found")
    }
    return category, nil
}




func (s *CategoryService) Create(name string, color string, userID uint) (*models.Category, error) {
	if name == "" {
		return nil, errors.New("category name is required")
	}


	existing, err := s.repo.GetByName(userID, name)
	if err == nil && existing.ID != 0 {
		return nil, errors.New("category already exists")
	}

	// Default color
	if color == "" {
		color = "#3B82F6"
	}

	category := &models.Category{
		Name:   name,
		Color:  color,
		UserID: userID,
	}
	
	err = s.repo.Create(category)
	if err != nil {
		return nil, err
	}

	return category, nil
}


//update func
func (s *CategoryService) Update(
	userID uint,
	categoryID uint,
	name string,
	color string,
) (*models.Category, error) {

	category, err := s.repo.FindByID(userID, categoryID)
	if err != nil {
		return nil, errors.New("category not found")
	}

	exists, err := s.repo.ExistsByName(userID, name, categoryID)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("category name already exists")
	}

	category.Name = name
	if color != "" {
		category.Color = color
	}

	if err := s.repo.Update(category); err != nil {
		return nil, err
	}

	return category, nil
}

// delete
func (s *CategoryService) Delete(userID, categoryID uint) error {
	// 1. pastikan category milik user
	category, err := s.repo.FindByID(userID, categoryID)
	if err != nil {
		return errors.New("category not found")
	}

	// 2. cek apakah category masih dipakai file
	if len(category.Files) > 0 {
		return errors.New("category is still used by files")
	}

	// 3. delete
	return s.repo.Delete(category)
}
