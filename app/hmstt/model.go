package hmstt

type hmsttState struct {
	Key   string `gorm:"primaryKey"`
	Value string
}

func (hmsttState) TableName() string {
	return "hmstt_states"
}
