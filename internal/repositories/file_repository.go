package repositories

import (
	"vasvault/internal/models"

	"gorm.io/gorm"
)

type FileRepositoryInterface interface {
	Create(file *models.File) error
	FindByID(id uint) (*models.File, error)
	FindByIDWithCategories(id uint) (*models.File, error)
	ListUserFiles(userID uint) ([]models.File, error)
	ListUserFilesWithCategories(userID uint) ([]models.File, error)
	Delete(fileID uint) error
	AssignCategories(fileID uint, categoryIDs []uint) error
	RemoveCategories(fileID uint, categoryIDs []uint) error
	ClearAllCategories(fileID uint) error
}

type FileRepository struct {
	db *gorm.DB
}

// Create implements FileRepositoryInterface.
func (r *FileRepository) Create(file *models.File) error {
	return r.db.Create(file).Error
}

func NewFileRepository(db *gorm.DB) *FileRepository {
	return &FileRepository{db: db}
}

func (r *FileRepository) Upload(file *models.File) error {
	return r.db.Create(file).Error
}

func (r *FileRepository) FindByID(id uint) (*models.File, error) {
	var file models.File
	if err := r.db.Where("id = ?", id).First(&file).Error; err != nil {
		return nil, err
	}
	return &file, nil
}

func (r *FileRepository) FindByIDWithCategories(id uint) (*models.File, error) {
	var file models.File
	if err := r.db.Preload("Categories").Where("id = ?", id).First(&file).Error; err != nil {
		return nil, err
	}
	return &file, nil
}

func (r *FileRepository) ListUserFiles(userID uint) ([]models.File, error) {
	var files []models.File
	if err := r.db.Where("user_id = ?", userID).Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}

func (r *FileRepository) ListUserFilesWithCategories(userID uint) ([]models.File, error) {
	var files []models.File
	if err := r.db.Preload("Categories").Where("user_id = ?", userID).Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}

func (r *FileRepository) Delete(fileID uint) error {
	return r.db.Delete(&models.File{}, fileID).Error
}

// AssignCategories menambahkan kategori ke file
func (r *FileRepository) AssignCategories(fileID uint, categoryIDs []uint) error {
	var file models.File
	if err := r.db.First(&file, fileID).Error; err != nil {
		return err
	}

	var categories []models.Category
	if err := r.db.Find(&categories, categoryIDs).Error; err != nil {
		return err
	}

	return r.db.Model(&file).Association("Categories").Append(&categories)
}

// RemoveCategories menghapus kategori dari file
func (r *FileRepository) RemoveCategories(fileID uint, categoryIDs []uint) error {
	var file models.File
	if err := r.db.First(&file, fileID).Error; err != nil {
		return err
	}

	var categories []models.Category
	if err := r.db.Find(&categories, categoryIDs).Error; err != nil {
		return err
	}

	return r.db.Model(&file).Association("Categories").Delete(&categories)
}

// ClearAllCategories menghapus semua kategori dari file
func (r *FileRepository) ClearAllCategories(fileID uint) error {
	var file models.File
	if err := r.db.First(&file, fileID).Error; err != nil {
		return err
	}

	return r.db.Model(&file).Association("Categories").Clear()
}
