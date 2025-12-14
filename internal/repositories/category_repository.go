package repositories

import (
	"vasvault/internal/models"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Create(category *models.Category) error {
	return r.db.Create(category).Error
}

func (r *CategoryRepository) GetByName(userID uint, name string) (*models.Category, error) {
	var category models.Category
	err := r.db.Where("user_id = ? AND name = ?", userID, name).First(&category).Error
	return &category, err
}

func (r *CategoryRepository) GetByID(id uint) (*models.Category, error) {
	var category models.Category
	err := r.db.First(&category, id).Error
	return &category, err
}

func (r *CategoryRepository) GetByIDAndUser(userID uint, id uint) (*models.Category, error) {
	var category models.Category
	err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// list, filter n sort
func (r *CategoryRepository) List(userID uint, search string) ([]models.Category, error) {
	var categories []models.Category

	query := r.db.Where("user_id = ?", userID)

	// Optional search
	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	// Default sort: latest
	err := query.Order("created_at DESC").
		Select("id, name, color, created_at, updated_at").
		Find(&categories).Error

	return categories, err
}


//func findbyid
func (r *CategoryRepository) FindByID(userID, categoryID uint) (*models.Category, error) {
	var category models.Category
	err := r.db.
		Preload("Files"). // 🔥 INI PRELOAD
		Where("id = ? AND user_id = ?", categoryID, userID).
		First(&category).Error
	return &category, err
}

// func (r *CategoryRepository) FindByID(userID, categoryID uint) (*models.Category, error) {
// 	var category models.Category
// 	err := r.db.
// 		Where("id = ? AND user_id = ?", categoryID, userID).
// 		First(&category).Error
// 	return &category, err
// }

func (r *CategoryRepository) ExistsByName(
	userID uint,
	name string,
	excludeID uint,
) (bool, error) {
	var count int64
	err := r.db.Model(&models.Category{}).
		Where("user_id = ? AND name = ? AND id <> ?", userID, name, excludeID).
		Count(&count).Error
	return count > 0, err
}


func (r *CategoryRepository) Update(category *models.Category) error {
	return r.db.Save(category).Error
}

//delete category
func (r *CategoryRepository) Delete(category *models.Category) error {
	return r.db.Delete(category).Error
}

