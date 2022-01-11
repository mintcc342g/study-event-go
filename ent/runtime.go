// Code generated by entc, DO NOT EDIT.

package ent

import (
	"study-event-go/ent/garden"
	"study-event-go/ent/lily"
	"study-event-go/ent/lilyskill"
	"study-event-go/ent/mentorship"
	"study-event-go/ent/schema"
	"study-event-go/ent/skill"
	"study-event-go/types"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	gardenFields := schema.Garden{}.Fields()
	_ = gardenFields
	// gardenDescCreatedAt is the schema descriptor for created_at field.
	gardenDescCreatedAt := gardenFields[1].Descriptor()
	// garden.DefaultCreatedAt holds the default value on creation for the created_at field.
	garden.DefaultCreatedAt = gardenDescCreatedAt.Default.(func() time.Time)
	// gardenDescUpdatedAt is the schema descriptor for updated_at field.
	gardenDescUpdatedAt := gardenFields[2].Descriptor()
	// garden.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	garden.DefaultUpdatedAt = gardenDescUpdatedAt.Default.(func() time.Time)
	// garden.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	garden.UpdateDefaultUpdatedAt = gardenDescUpdatedAt.UpdateDefault.(func() time.Time)
	// gardenDescName is the schema descriptor for name field.
	gardenDescName := gardenFields[4].Descriptor()
	// garden.NameValidator is a validator for the "name" field. It is called by the builders before save.
	garden.NameValidator = gardenDescName.Validators[0].(func(string) error)
	lilyFields := schema.Lily{}.Fields()
	_ = lilyFields
	// lilyDescCreatedAt is the schema descriptor for created_at field.
	lilyDescCreatedAt := lilyFields[1].Descriptor()
	// lily.DefaultCreatedAt holds the default value on creation for the created_at field.
	lily.DefaultCreatedAt = lilyDescCreatedAt.Default.(func() time.Time)
	// lilyDescUpdatedAt is the schema descriptor for updated_at field.
	lilyDescUpdatedAt := lilyFields[2].Descriptor()
	// lily.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	lily.DefaultUpdatedAt = lilyDescUpdatedAt.Default.(func() time.Time)
	// lily.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	lily.UpdateDefaultUpdatedAt = lilyDescUpdatedAt.UpdateDefault.(func() time.Time)
	// lilyDescCauseOfDeletion is the schema descriptor for cause_of_deletion field.
	lilyDescCauseOfDeletion := lilyFields[4].Descriptor()
	// lily.DefaultCauseOfDeletion holds the default value on creation for the cause_of_deletion field.
	lily.DefaultCauseOfDeletion = types.CauseOfDeletion(lilyDescCauseOfDeletion.Default.(uint32))
	// lilyDescFirstName is the schema descriptor for first_name field.
	lilyDescFirstName := lilyFields[6].Descriptor()
	// lily.FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	lily.FirstNameValidator = lilyDescFirstName.Validators[0].(func(string) error)
	// lilyDescLastName is the schema descriptor for last_name field.
	lilyDescLastName := lilyFields[8].Descriptor()
	// lily.LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	lily.LastNameValidator = lilyDescLastName.Validators[0].(func(string) error)
	// lilyDescRank is the schema descriptor for rank field.
	lilyDescRank := lilyFields[9].Descriptor()
	// lily.DefaultRank holds the default value on creation for the rank field.
	lily.DefaultRank = lilyDescRank.Default.(uint32)
	// lilyDescGardenID is the schema descriptor for garden_id field.
	lilyDescGardenID := lilyFields[10].Descriptor()
	// lily.DefaultGardenID holds the default value on creation for the garden_id field.
	lily.DefaultGardenID = types.GardenID(lilyDescGardenID.Default.(uint64))
	// lilyDescLegionID is the schema descriptor for legion_id field.
	lilyDescLegionID := lilyFields[11].Descriptor()
	// lily.DefaultLegionID holds the default value on creation for the legion_id field.
	lily.DefaultLegionID = types.LegionID(lilyDescLegionID.Default.(uint64))
	lilyskillFields := schema.LilySkill{}.Fields()
	_ = lilyskillFields
	// lilyskillDescCreatedAt is the schema descriptor for created_at field.
	lilyskillDescCreatedAt := lilyskillFields[2].Descriptor()
	// lilyskill.DefaultCreatedAt holds the default value on creation for the created_at field.
	lilyskill.DefaultCreatedAt = lilyskillDescCreatedAt.Default.(func() time.Time)
	// lilyskillDescUpdatedAt is the schema descriptor for updated_at field.
	lilyskillDescUpdatedAt := lilyskillFields[3].Descriptor()
	// lilyskill.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	lilyskill.DefaultUpdatedAt = lilyskillDescUpdatedAt.Default.(func() time.Time)
	// lilyskill.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	lilyskill.UpdateDefaultUpdatedAt = lilyskillDescUpdatedAt.UpdateDefault.(func() time.Time)
	mentorshipFields := schema.Mentorship{}.Fields()
	_ = mentorshipFields
	// mentorshipDescCreatedAt is the schema descriptor for created_at field.
	mentorshipDescCreatedAt := mentorshipFields[1].Descriptor()
	// mentorship.DefaultCreatedAt holds the default value on creation for the created_at field.
	mentorship.DefaultCreatedAt = mentorshipDescCreatedAt.Default.(func() time.Time)
	// mentorshipDescUpdatedAt is the schema descriptor for updated_at field.
	mentorshipDescUpdatedAt := mentorshipFields[2].Descriptor()
	// mentorship.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	mentorship.DefaultUpdatedAt = mentorshipDescUpdatedAt.Default.(func() time.Time)
	// mentorship.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	mentorship.UpdateDefaultUpdatedAt = mentorshipDescUpdatedAt.UpdateDefault.(func() time.Time)
	// mentorshipDescName is the schema descriptor for name field.
	mentorshipDescName := mentorshipFields[4].Descriptor()
	// mentorship.NameValidator is a validator for the "name" field. It is called by the builders before save.
	mentorship.NameValidator = mentorshipDescName.Validators[0].(func(string) error)
	skillFields := schema.Skill{}.Fields()
	_ = skillFields
	// skillDescCreatedAt is the schema descriptor for created_at field.
	skillDescCreatedAt := skillFields[1].Descriptor()
	// skill.DefaultCreatedAt holds the default value on creation for the created_at field.
	skill.DefaultCreatedAt = skillDescCreatedAt.Default.(func() time.Time)
	// skillDescUpdatedAt is the schema descriptor for updated_at field.
	skillDescUpdatedAt := skillFields[2].Descriptor()
	// skill.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	skill.DefaultUpdatedAt = skillDescUpdatedAt.Default.(func() time.Time)
	// skill.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	skill.UpdateDefaultUpdatedAt = skillDescUpdatedAt.UpdateDefault.(func() time.Time)
	// skillDescName is the schema descriptor for name field.
	skillDescName := skillFields[4].Descriptor()
	// skill.NameValidator is a validator for the "name" field. It is called by the builders before save.
	skill.NameValidator = skillDescName.Validators[0].(func(string) error)
	// skillDescType is the schema descriptor for type field.
	skillDescType := skillFields[5].Descriptor()
	// skill.DefaultType holds the default value on creation for the type field.
	skill.DefaultType = types.SkillType(skillDescType.Default.(uint32))
}
