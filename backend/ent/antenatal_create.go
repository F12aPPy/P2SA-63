// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/F12aPPy/app/ent/antenatal"
	"github.com/F12aPPy/app/ent/babystatus"
	"github.com/F12aPPy/app/ent/patient"
	"github.com/F12aPPy/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// AntenatalCreate is the builder for creating a Antenatal entity.
type AntenatalCreate struct {
	config
	mutation *AntenatalMutation
	hooks    []Hook
}

// SetAddedTime sets the added_time field.
func (ac *AntenatalCreate) SetAddedTime(t time.Time) *AntenatalCreate {
	ac.mutation.SetAddedTime(t)
	return ac
}

// SetUserID sets the user edge to User by id.
func (ac *AntenatalCreate) SetUserID(id int) *AntenatalCreate {
	ac.mutation.SetUserID(id)
	return ac
}

// SetNillableUserID sets the user edge to User by id if the given value is not nil.
func (ac *AntenatalCreate) SetNillableUserID(id *int) *AntenatalCreate {
	if id != nil {
		ac = ac.SetUserID(*id)
	}
	return ac
}

// SetUser sets the user edge to User.
func (ac *AntenatalCreate) SetUser(u *User) *AntenatalCreate {
	return ac.SetUserID(u.ID)
}

// SetPatientID sets the patient edge to Patient by id.
func (ac *AntenatalCreate) SetPatientID(id int) *AntenatalCreate {
	ac.mutation.SetPatientID(id)
	return ac
}

// SetNillablePatientID sets the patient edge to Patient by id if the given value is not nil.
func (ac *AntenatalCreate) SetNillablePatientID(id *int) *AntenatalCreate {
	if id != nil {
		ac = ac.SetPatientID(*id)
	}
	return ac
}

// SetPatient sets the patient edge to Patient.
func (ac *AntenatalCreate) SetPatient(p *Patient) *AntenatalCreate {
	return ac.SetPatientID(p.ID)
}

// SetBabystatusID sets the babystatus edge to Babystatus by id.
func (ac *AntenatalCreate) SetBabystatusID(id int) *AntenatalCreate {
	ac.mutation.SetBabystatusID(id)
	return ac
}

// SetNillableBabystatusID sets the babystatus edge to Babystatus by id if the given value is not nil.
func (ac *AntenatalCreate) SetNillableBabystatusID(id *int) *AntenatalCreate {
	if id != nil {
		ac = ac.SetBabystatusID(*id)
	}
	return ac
}

// SetBabystatus sets the babystatus edge to Babystatus.
func (ac *AntenatalCreate) SetBabystatus(b *Babystatus) *AntenatalCreate {
	return ac.SetBabystatusID(b.ID)
}

// Mutation returns the AntenatalMutation object of the builder.
func (ac *AntenatalCreate) Mutation() *AntenatalMutation {
	return ac.mutation
}

// Save creates the Antenatal in the database.
func (ac *AntenatalCreate) Save(ctx context.Context) (*Antenatal, error) {
	if _, ok := ac.mutation.AddedTime(); !ok {
		return nil, &ValidationError{Name: "added_time", err: errors.New("ent: missing required field \"added_time\"")}
	}
	var (
		err  error
		node *Antenatal
	)
	if len(ac.hooks) == 0 {
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AntenatalMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ac.mutation = mutation
			node, err = ac.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AntenatalCreate) SaveX(ctx context.Context) *Antenatal {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ac *AntenatalCreate) sqlSave(ctx context.Context) (*Antenatal, error) {
	a, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	a.ID = int(id)
	return a, nil
}

func (ac *AntenatalCreate) createSpec() (*Antenatal, *sqlgraph.CreateSpec) {
	var (
		a     = &Antenatal{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: antenatal.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: antenatal.FieldID,
			},
		}
	)
	if value, ok := ac.mutation.AddedTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: antenatal.FieldAddedTime,
		})
		a.AddedTime = value
	}
	if nodes := ac.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   antenatal.UserTable,
			Columns: []string{antenatal.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.PatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   antenatal.PatientTable,
			Columns: []string{antenatal.PatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.BabystatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   antenatal.BabystatusTable,
			Columns: []string{antenatal.BabystatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: babystatus.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return a, _spec
}