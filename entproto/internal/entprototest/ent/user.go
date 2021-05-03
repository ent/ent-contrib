// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/contrib/entproto/internal/entprototest/ent/image"
	"entgo.io/contrib/entproto/internal/entprototest/ent/user"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserName holds the value of the "user_name" field.
	UserName string `json:"user_name,omitempty"`
	// Status holds the value of the "status" field.
	Status user.Status `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges            UserEdges `json:"edges"`
	user_profile_pic *uuid.UUID
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// BlogPosts holds the value of the blog_posts edge.
	BlogPosts []*BlogPost `json:"blog_posts,omitempty"`
	// ProfilePic holds the value of the profile_pic edge.
	ProfilePic *Image `json:"profile_pic,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// BlogPostsOrErr returns the BlogPosts value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) BlogPostsOrErr() ([]*BlogPost, error) {
	if e.loadedTypes[0] {
		return e.BlogPosts, nil
	}
	return nil, &NotLoadedError{edge: "blog_posts"}
}

// ProfilePicOrErr returns the ProfilePic value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) ProfilePicOrErr() (*Image, error) {
	if e.loadedTypes[1] {
		if e.ProfilePic == nil {
			// The edge profile_pic was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: image.Label}
		}
		return e.ProfilePic, nil
	}
	return nil, &NotLoadedError{edge: "profile_pic"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldUserName, user.FieldStatus:
			values[i] = new(sql.NullString)
		case user.ForeignKeys[0]: // user_profile_pic
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldUserName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_name", values[i])
			} else if value.Valid {
				u.UserName = value.String
			}
		case user.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				u.Status = user.Status(value.String)
			}
		case user.ForeignKeys[0]:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_profile_pic", values[i])
			} else if value != nil {
				u.user_profile_pic = value
			}
		}
	}
	return nil
}

// QueryBlogPosts queries the "blog_posts" edge of the User entity.
func (u *User) QueryBlogPosts() *BlogPostQuery {
	return (&UserClient{config: u.config}).QueryBlogPosts(u)
}

// QueryProfilePic queries the "profile_pic" edge of the User entity.
func (u *User) QueryProfilePic() *ImageQuery {
	return (&UserClient{config: u.config}).QueryProfilePic(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", user_name=")
	builder.WriteString(u.UserName)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", u.Status))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
