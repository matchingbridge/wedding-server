package models

type User struct {
	UserID string `json:"user_id" gorm:"not null;type:varchar(36);primary_key" form:"user_id" binding:"required,max=36"`
	Email  string `json:"email"   gorm:"not null;type:text"                    form:"email"   binding:"required"`
	Name   string `json:"name"    gorm:"not null;type:varchar(12)"             form:"name"    binding:"required,max=12"`
	Phone  string `json:"phone"   gorm:"not null;type:varchar(12);unique"      form:"phone"   binding:"required,max=12"`

	Face1 string `json:"face1"   gorm:"not null;type:text" form:"face1"  binding:"required"`
	Face2 string `json:"face2"   gorm:"not null;type:text" form:"face2"  binding:"required"`
	Body1 string `json:"body1"   gorm:"not null;type:text" form:"body1"  binding:"required"`
	Body2 string `json:"body2"   gorm:"not null;type:text" form:"body2"  binding:"required"`
	Video string `json:"video"    gorm:"not null;type:text" form:"video"   binding:"required"`

	Introduction string `json:"introduction" gorm:"not null;type:text" form:"introduction"`

	Area    int    `json:"area"    gorm:"not null;type:tinyint" form:"area"`
	Address string `json:"address" gorm:"not null;type:text"    form:"address" binding:"required"`
	Detail  string `json:"detail"  gorm:"type:text"             form:"detail"`

	Gender    int    `json:"gender"    gorm:"not null;type:tinyint"    form:"gender"    binding:"required"`
	Birthday  string `json:"birthday"  gorm:"not null;type:varchar(6)" form:"birthday"  binding:"required"`
	Marriage  int    `json:"marriage"  gorm:"not null;type:tinyint"    form:"marriage"  binding:"required"`
	Height    int    `json:"height"    gorm:"not null;type:smallint"   form:"height"    binding:"required"`
	Weight    int    `json:"weight"    gorm:"not null;type:smallint"   form:"weight"    binding:"required"`
	Smoke     bool   `json:"smoke"     gorm:"not null;type:boolean"    form:"smoke"     binding:"required"`
	Drink     bool   `json:"drink"     gorm:"not null;type:boolean"    form:"drink"     binding:"required"`
	Religion  int    `json:"religion"  gorm:"not null;type:tinyint"    form:"religion"  binding:"required"`
	BodyType  string `json:"bodytype"  gorm:"not null;type:text"       form:"bodytype"  binding:"required"`
	Character string `json:"character" gorm:"not null;type:text"       form:"character" binding:"required"`

	Scholar int    `json:"scholar" gorm:"not null;type:tinyint" form:"scholar" binding:"required"`
	School  string `json:"school"  gorm:"not null;type:text"    form:"school"  binding:"required"`
	Major   string `json:"major"   gorm:"not null;type:text"    form:"major"   binding:"required"`

	Job        int    `json:"job"        gorm:"not null;type:tinyint" form:"job"         binding:"required"`
	Company    string `json:"company"    gorm:"not null;type:text"    form:"company"     binding:"required"`
	Workplace  string `json:"workplace"  gorm:"not null;type:text"    form:"workplace"   binding:"required"`
	Employment int    `json:"employment" gorm:"not null;type:tinyint" form:"employtment" binding:"required"`
	Salary     int    `json:"salary"     gorm:"not null;type:tinyint" form:"salary"      binding:"required"`

	Distance bool   `json:"distance" gorm:"not null;type:boolean"    form:"distance" binding:"required"`
	Period   string `json:"period"   gorm:"not null;type:varchar(6)" form:"period"   binding:"required"`

	PartnerCharacters string `json:"partner_characters" gorm:"type:text" form:"partner_characters"`
	Vehicle           string `json:"vehicle"  gorm:"not null;type:text" form:"vehicle" binding:"required"`
}
