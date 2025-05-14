package logging

type Category string
type SubCategory string

var subCategoryToCategory = map[SubCategory]Category{
	Startup:              General,
	ExternalService:      General,
	Config:               General,
	Migration:            DataBase,
	Select:               DataBase,
	Rollback:             DataBase,
	Update:               DataBase,
	Delete:               DataBase,
	Insert:               DataBase,
	Api:                  Internal,
	HashPassword:         Internal,
	DefaultValueNotFound: Internal,
	FailedToCreate:       Internal,
	MobileValidation:     Validation,
	PasswordValidation:   Validation,
	RemoveFile:           IO,
}

const (
	General    Category = "General"
	DataBase   Category = "DataBase"
	Internal   Category = "Internal"
	Validation Category = "Validation"
	IO         Category = "IO"
)

const (
	Startup         SubCategory = "Startup"
	ExternalService SubCategory = "ExternalService"
	Config          SubCategory = "Config"

	Migration SubCategory = "Migration"
	Select    SubCategory = "Select"
	Rollback  SubCategory = "Rollback"
	Update    SubCategory = "Update"
	Delete    SubCategory = "Delete"
	Insert    SubCategory = "Insert"

	Api                  SubCategory = "Api"
	HashPassword         SubCategory = "HashPassword"
	DefaultValueNotFound SubCategory = "DefaultValueNotFound"
	FailedToCreate       SubCategory = "FailedToCreate"

	MobileValidation   SubCategory = "MobileValidation"
	PasswordValidation SubCategory = "PasswordValidation"

	RemoveFile SubCategory = "RemoveFile"
)
