package shopping

type addItem struct {
	HiveID uint   `json:"hive_id" validate:"required,gt=0"`
	Item   string `json:"item" validate:"required,min=1"`
}
