package chore

type createChore struct {
	HiveID uint   `json:"hive_id" validate:"required,gt=0"`
	Chore  string `json:"chore" validate:"required,min=1"`
}
