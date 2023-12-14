package repository

import (
	"context"
	"errors"
	"wanderer/features/locations"
	"wanderer/helpers/filters"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"gorm.io/gorm"
)

func NewLocationRepository(mysqlDB *gorm.DB, cld *cloudinary.Cloudinary) locations.Repository {
	return &locationRepository{
		mysqlDB: mysqlDB,
		cld:     cld,
	}
}

type locationRepository struct {
	mysqlDB *gorm.DB
	cld     *cloudinary.Cloudinary
}

func (repo *locationRepository) GetAll(ctx context.Context, flt filters.Filter) ([]locations.Location, error) {
	var data []Location

	qry := repo.mysqlDB

	if flt.Search.Keyword != "" {
		qry = qry.Where("name like ?", "%"+flt.Search.Keyword+"%")
	}

	if flt.Pagination.Limit != 0 {
		qry = qry.Limit(flt.Pagination.Limit)
	}

	qry = qry.Find(&data)
	if qry.Error != nil {
		return nil, qry.Error
	}

	var result []locations.Location
	for _, location := range data {
		result = append(result, *location.ToEntity())
	}

	return result, nil
}

func (repo *locationRepository) Create(ctx context.Context, data locations.Location) error {
	var mod = new(Location)
	mod.FromEntity(data)

	if mod.ImageRaw != nil {
		UniqueFilename := true
		res, err := repo.cld.Upload.Upload(ctx, mod.ImageRaw, uploader.UploadParams{
			UniqueFilename: &UniqueFilename,
			Folder:         "locations",
		})

		if err != nil {
			return err
		}

		mod.ImageUrl = res.URL
	}

	qry := repo.mysqlDB.Create(mod)
	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected == 0 {
		return errors.New("failed to create location")
	}

	return nil
}

func (repo *locationRepository) Update(ctx context.Context, id uint, data locations.Location) error {
	var mod = new(Location)
	mod.FromEntity(data)

	if mod.ImageRaw != nil {
		UniqueFilename := true
		res, err := repo.cld.Upload.Upload(ctx, mod.ImageRaw, uploader.UploadParams{
			UniqueFilename: &UniqueFilename,
			Folder:         "locations",
		})

		if err != nil {
			return err
		}

		mod.ImageUrl = res.URL
	}

	qry := repo.mysqlDB.Where(&Location{Id: id}).Updates(mod)
	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected == 0 {
		return errors.New("not found")
	}

	return nil
}

func (repo *locationRepository) Delete(ctx context.Context, id uint) error {
	qry := repo.mysqlDB.Where(&Location{Id: id}).Delete(&Location{})
	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected == 0 {
		return errors.New("not found")
	}

	return nil
}
