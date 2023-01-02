package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("name"),
		field.Bool("isActivated").Default(true),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		//ent에서는 To를 정의하는 스키마 즉 여기선 User가
		//참조 관계의 주인이라고 정의함
		//일반적인 JPA의 방식과는 반대
		edge.To("products", TourProduct.Type),
	}
}
