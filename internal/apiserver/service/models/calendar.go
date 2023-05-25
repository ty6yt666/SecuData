package models

type Calendar struct {
	Date     string   `json:"date" gorm:"column:date"`
	Trading  []string `json:"trading" gorm:"column:trading"`
	Clearing []string `json:"clearing" gorm:"column:clearing"`
}

type CalendarList struct {
	Items []*Calendar `json:"items"`
}

func (c *Calendar) TableName() string {
	return "calendar"
}
