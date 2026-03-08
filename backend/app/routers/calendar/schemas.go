package calendar

type createEvent struct {
	HiveID     uint   `json:"hive_id" validate:"required,gt=0"`
	EventTitle string `json:"event_title" validate:"required,min=1,max=255"`
	// Expected from frontend: 2026-02-15T18:30:00+05:30
	EventTimestamp string `json:"event_timestamp" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	// Expected from frontend: Asia/Kolkata
	EventOriginTimezone string `json:"event_origin_timezone" validate:"required,timezone"`
}
