package portfolio

type portfolio struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" binding:required`
}
