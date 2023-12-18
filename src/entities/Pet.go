package entities

type Pet struct{
	ID 				string 		`json:"id" gorm:"primaryKey"`
	Nickname 		string 		`json:"nickname"`
	BreedType		string		`json:"breedType"`
	Size			string		`json:"size"`
	Gender			string		`json:"gender"`
	ThumbnailImage 	string		`json:"thumbnailImage" gorm:"type:text"`
	Images			[]string	`json:"images" gorm:"type:text[]"`
	Description 	string		`json:"description" gorm:"type:text"`
	Author			User		`json:"author`
	CreatedAt   	time.Time 	`json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   	time.Time 	`json:"updatedAt" gorm:"autoUpdateTime"`
}

type PetList []*Pet

func (p *Pet) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New().String()
	return
}