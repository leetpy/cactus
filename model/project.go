package model

type ProjectModel struct {
	BaseModel
	Name        string `json:"project_name" gorm:"column:name;not null"`
	Description string `gorm:"column:description;null"`
}

func (p *ProjectModel) TableName() string {
	return "project"
}

func (p *ProjectModel) Create() error {
	return nil
}
