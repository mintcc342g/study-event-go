// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"study-event-go/ent/charmmodel"
	"study-event-go/types"
	"time"

	"entgo.io/ent/dialect/sql"
)

// CharmModel is the model entity for the CharmModel schema.
type CharmModel struct {
	config `json:"-"`
	// ID of the ent.
	ID types.CharmModelID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// CreatorID holds the value of the "creator_id" field.
	CreatorID types.CharmCreatorID `json:"creator_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Type holds the value of the "type" field.
	Type types.CharmModelType `json:"type,omitempty"`
	// Generation holds the value of the "generation" field.
	Generation types.CharmModelGeneration `json:"generation,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CharmModel) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case charmmodel.FieldID, charmmodel.FieldCreatorID, charmmodel.FieldType, charmmodel.FieldGeneration:
			values[i] = new(sql.NullInt64)
		case charmmodel.FieldName:
			values[i] = new(sql.NullString)
		case charmmodel.FieldCreatedAt, charmmodel.FieldUpdatedAt, charmmodel.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CharmModel", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CharmModel fields.
func (cm *CharmModel) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case charmmodel.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cm.ID = types.CharmModelID(value.Int64)
		case charmmodel.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				cm.CreatedAt = value.Time
			}
		case charmmodel.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				cm.UpdatedAt = value.Time
			}
		case charmmodel.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				cm.DeletedAt = new(time.Time)
				*cm.DeletedAt = value.Time
			}
		case charmmodel.FieldCreatorID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field creator_id", values[i])
			} else if value.Valid {
				cm.CreatorID = types.CharmCreatorID(value.Int64)
			}
		case charmmodel.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				cm.Name = value.String
			}
		case charmmodel.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				cm.Type = types.CharmModelType(value.Int64)
			}
		case charmmodel.FieldGeneration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field generation", values[i])
			} else if value.Valid {
				cm.Generation = types.CharmModelGeneration(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this CharmModel.
// Note that you need to call CharmModel.Unwrap() before calling this method if this CharmModel
// was returned from a transaction, and the transaction was committed or rolled back.
func (cm *CharmModel) Update() *CharmModelUpdateOne {
	return (&CharmModelClient{config: cm.config}).UpdateOne(cm)
}

// Unwrap unwraps the CharmModel entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cm *CharmModel) Unwrap() *CharmModel {
	tx, ok := cm.config.driver.(*txDriver)
	if !ok {
		panic("ent: CharmModel is not a transactional entity")
	}
	cm.config.driver = tx.drv
	return cm
}

// String implements the fmt.Stringer.
func (cm *CharmModel) String() string {
	var builder strings.Builder
	builder.WriteString("CharmModel(")
	builder.WriteString(fmt.Sprintf("id=%v", cm.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(cm.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(cm.UpdatedAt.Format(time.ANSIC))
	if v := cm.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", creator_id=")
	builder.WriteString(fmt.Sprintf("%v", cm.CreatorID))
	builder.WriteString(", name=")
	builder.WriteString(cm.Name)
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", cm.Type))
	builder.WriteString(", generation=")
	builder.WriteString(fmt.Sprintf("%v", cm.Generation))
	builder.WriteByte(')')
	return builder.String()
}

// CharmModels is a parsable slice of CharmModel.
type CharmModels []*CharmModel

func (cm CharmModels) config(cfg config) {
	for _i := range cm {
		cm[_i].config = cfg
	}
}
