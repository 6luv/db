// Code generated by ent, DO NOT EDIT.

package ent

import (
	"db/ent/tourproduct"
	"db/ent/user"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// TourProduct is the model entity for the TourProduct schema.
type TourProduct struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Price holds the value of the "price" field.
	Price int `json:"price,omitempty"`
	// ForSale holds the value of the "forSale" field.
	ForSale bool `json:"forSale,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TourProductQuery when eager-loading is set.
	Edges         TourProductEdges `json:"edges"`
	user_products *string
}

// TourProductEdges holds the relations/edges for other nodes in the graph.
type TourProductEdges struct {
	// Manager holds the value of the manager edge.
	Manager *User `json:"manager,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ManagerOrErr returns the Manager value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TourProductEdges) ManagerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Manager == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Manager, nil
	}
	return nil, &NotLoadedError{edge: "manager"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TourProduct) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tourproduct.FieldForSale:
			values[i] = new(sql.NullBool)
		case tourproduct.FieldID, tourproduct.FieldPrice:
			values[i] = new(sql.NullInt64)
		case tourproduct.FieldName:
			values[i] = new(sql.NullString)
		case tourproduct.ForeignKeys[0]: // user_products
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TourProduct", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TourProduct fields.
func (tp *TourProduct) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tourproduct.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tp.ID = int(value.Int64)
		case tourproduct.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				tp.Name = value.String
			}
		case tourproduct.FieldPrice:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				tp.Price = int(value.Int64)
			}
		case tourproduct.FieldForSale:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field forSale", values[i])
			} else if value.Valid {
				tp.ForSale = value.Bool
			}
		case tourproduct.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_products", values[i])
			} else if value.Valid {
				tp.user_products = new(string)
				*tp.user_products = value.String
			}
		}
	}
	return nil
}

// QueryManager queries the "manager" edge of the TourProduct entity.
func (tp *TourProduct) QueryManager() *UserQuery {
	return (&TourProductClient{config: tp.config}).QueryManager(tp)
}

// Update returns a builder for updating this TourProduct.
// Note that you need to call TourProduct.Unwrap() before calling this method if this TourProduct
// was returned from a transaction, and the transaction was committed or rolled back.
func (tp *TourProduct) Update() *TourProductUpdateOne {
	return (&TourProductClient{config: tp.config}).UpdateOne(tp)
}

// Unwrap unwraps the TourProduct entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tp *TourProduct) Unwrap() *TourProduct {
	_tx, ok := tp.config.driver.(*txDriver)
	if !ok {
		panic("ent: TourProduct is not a transactional entity")
	}
	tp.config.driver = _tx.drv
	return tp
}

// String implements the fmt.Stringer.
func (tp *TourProduct) String() string {
	var builder strings.Builder
	builder.WriteString("TourProduct(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tp.ID))
	builder.WriteString("name=")
	builder.WriteString(tp.Name)
	builder.WriteString(", ")
	builder.WriteString("price=")
	builder.WriteString(fmt.Sprintf("%v", tp.Price))
	builder.WriteString(", ")
	builder.WriteString("forSale=")
	builder.WriteString(fmt.Sprintf("%v", tp.ForSale))
	builder.WriteByte(')')
	return builder.String()
}

// TourProducts is a parsable slice of TourProduct.
type TourProducts []*TourProduct

func (tp TourProducts) config(cfg config) {
	for _i := range tp {
		tp[_i].config = cfg
	}
}
