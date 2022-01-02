package models

type Detail struct {
	DetailID int    `json:"detail_id" gorm:"not null;primary_key;AUTO_INCREMENT"`
	UserID   string `json:"user_id"   gorm:"not null;type:varchar(36);unique"`

	Asset      string `json:"asset"      gorm:"not null;type:text"    form:"asset"      binding:"required"`
	RealEstate int    `json:"realestate" gorm:"not null;type:tinyint" form:"realestate" binding:"required"`

	FamilyCertificate  string `json:"family_certificate"  gorm:"type:text" form:"family_certificate"`
	SalaryCertificate  string `json:"salary_certificate"  gorm:"type:text" form:"salary_certificate"`
	SchoolCertificate  string `json:"school_certificate"  gorm:"type:text" form:"school_certificate"`
	CompanyCertificate string `json:"company_certificate" gorm:"type:text" form:"company_certificate"`
	VehicleCertificate string `json:"vehicle_ceritiface"  gorm:"type:text" form:"vehicle_certificate"`

	PartnerAgeMax      int    `json:"partner_age_max"     gorm:"type:tinyint"  form:"partner_age_max"`
	PartnerAgeMin      int    `json:"partner_age_min"     gorm:"type:tinyint"  form:"partner_age_min"`
	PartnerHeightMax   int    `json:"partner_height_max"  gorm:"type:smallint" form:"partner_height_max"`
	PartnerHeightMin   int    `json:"partner_height_min"  gorm:"type:smallint" form:"partner_height_min"`
	PartnerScool       int    `json:"partner_school"      gorm:"type:tinyint"  form:"partner_school"`
	PartnerJob         int    `json:"partner_job"         gorm:"type:tinyint"  form:"partner_job"`
	PartnerSalary      int    `json:"partner_salary"      gorm:"type:tinyint"  form:"partner_salary"`
	PartnerAsset       int    `json:"partner_asset"       gorm:"type:text"     form:"partner_asset"`
	PartnerSmoke       int    `json:"partner_smoke"       gorm:"type:tinyint"  form:"partner_smoke"`
	PartnerDrink       int    `json:"partner_drink"       gorm:"type:tinyint"  form:"partner_drink"`
	PartnerMarriage    int    `json:"partner_marriage"    gorm:"type:tinyint"  form:"partner_marriage"`
	PartnerBodies      string `json:"partner_bodies"      gorm:"type:text"     form:"partner_bodies"`
	PartnerDescription string `json:"partner_description" gorm:"type:text"     form:"partner_description"`
}
