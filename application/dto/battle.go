package dto

import "study_event_go/types"

// Alarm ...
type Alarm struct {
	GardenID     uint64
	CaveLocation string
	Huges        []*Huge
	TotalCount   uint32
	AlertLevel   types.AlertLevel
}

// Huge ...
type Huge struct {
	Class types.HugeClass
	Type  types.HugeType
}