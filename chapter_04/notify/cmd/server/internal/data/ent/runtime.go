// Code generated by entc, DO NOT EDIT.

package ent

import (
	"notify-server/internal/data/ent/schema"
	"notify-server/internal/data/ent/tag"
	"notify-server/internal/data/ent/template"
	"notify-server/internal/data/ent/user"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	tagFields := schema.Tag{}.Fields()
	_ = tagFields
	// tagDescCreatedAt is the schema descriptor for created_at field.
	tagDescCreatedAt := tagFields[0].Descriptor()
	// tag.DefaultCreatedAt holds the default value on creation for the created_at field.
	tag.DefaultCreatedAt = tagDescCreatedAt.Default.(func() time.Time)
	// tagDescUpdatedAt is the schema descriptor for updated_at field.
	tagDescUpdatedAt := tagFields[1].Descriptor()
	// tag.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	tag.DefaultUpdatedAt = tagDescUpdatedAt.Default.(func() time.Time)
	// tag.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	tag.UpdateDefaultUpdatedAt = tagDescUpdatedAt.UpdateDefault.(func() time.Time)
	// tagDescUUID is the schema descriptor for uuid field.
	tagDescUUID := tagFields[2].Descriptor()
	// tag.DefaultUUID holds the default value on creation for the uuid field.
	tag.DefaultUUID = tagDescUUID.Default.(func() uuid.UUID)
	// tagDescDesc is the schema descriptor for desc field.
	tagDescDesc := tagFields[3].Descriptor()
	// tag.DescValidator is a validator for the "desc" field. It is called by the builders before save.
	tag.DescValidator = func() func(string) error {
		validators := tagDescDesc.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(desc string) error {
			for _, fn := range fns {
				if err := fn(desc); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// tagDescName is the schema descriptor for name field.
	tagDescName := tagFields[4].Descriptor()
	// tag.NameValidator is a validator for the "name" field. It is called by the builders before save.
	tag.NameValidator = func() func(string) error {
		validators := tagDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// tagDescStatus is the schema descriptor for status field.
	tagDescStatus := tagFields[5].Descriptor()
	// tag.DefaultStatus holds the default value on creation for the status field.
	tag.DefaultStatus = tagDescStatus.Default.(int)
	// tag.StatusValidator is a validator for the "status" field. It is called by the builders before save.
	tag.StatusValidator = tagDescStatus.Validators[0].(func(int) error)
	templateFields := schema.Template{}.Fields()
	_ = templateFields
	// templateDescCreatedAt is the schema descriptor for created_at field.
	templateDescCreatedAt := templateFields[0].Descriptor()
	// template.DefaultCreatedAt holds the default value on creation for the created_at field.
	template.DefaultCreatedAt = templateDescCreatedAt.Default.(func() time.Time)
	// templateDescUpdatedAt is the schema descriptor for updated_at field.
	templateDescUpdatedAt := templateFields[1].Descriptor()
	// template.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	template.DefaultUpdatedAt = templateDescUpdatedAt.Default.(func() time.Time)
	// template.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	template.UpdateDefaultUpdatedAt = templateDescUpdatedAt.UpdateDefault.(func() time.Time)
	// templateDescUUID is the schema descriptor for uuid field.
	templateDescUUID := templateFields[2].Descriptor()
	// template.DefaultUUID holds the default value on creation for the uuid field.
	template.DefaultUUID = templateDescUUID.Default.(func() uuid.UUID)
	// templateDescDesc is the schema descriptor for desc field.
	templateDescDesc := templateFields[3].Descriptor()
	// template.DescValidator is a validator for the "desc" field. It is called by the builders before save.
	template.DescValidator = func() func(string) error {
		validators := templateDescDesc.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(desc string) error {
			for _, fn := range fns {
				if err := fn(desc); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// templateDescName is the schema descriptor for name field.
	templateDescName := templateFields[4].Descriptor()
	// template.NameValidator is a validator for the "name" field. It is called by the builders before save.
	template.NameValidator = func() func(string) error {
		validators := templateDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// templateDescStatus is the schema descriptor for status field.
	templateDescStatus := templateFields[6].Descriptor()
	// template.DefaultStatus holds the default value on creation for the status field.
	template.DefaultStatus = templateDescStatus.Default.(int)
	// template.StatusValidator is a validator for the "status" field. It is called by the builders before save.
	template.StatusValidator = templateDescStatus.Validators[0].(func(int) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescUUID is the schema descriptor for uuid field.
	userDescUUID := userFields[2].Descriptor()
	// user.DefaultUUID holds the default value on creation for the uuid field.
	user.DefaultUUID = userDescUUID.Default.(func() uuid.UUID)
	// userDescMobile is the schema descriptor for mobile field.
	userDescMobile := userFields[3].Descriptor()
	// user.MobileValidator is a validator for the "mobile" field. It is called by the builders before save.
	user.MobileValidator = func() func(string) error {
		validators := userDescMobile.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(mobile string) error {
			for _, fn := range fns {
				if err := fn(mobile); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[4].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = func() func(string) error {
		validators := userDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[5].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = func() func(string) error {
		validators := userDescEmail.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(email string) error {
			for _, fn := range fns {
				if err := fn(email); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescStatus is the schema descriptor for status field.
	userDescStatus := userFields[6].Descriptor()
	// user.DefaultStatus holds the default value on creation for the status field.
	user.DefaultStatus = userDescStatus.Default.(int)
	// user.StatusValidator is a validator for the "status" field. It is called by the builders before save.
	user.StatusValidator = userDescStatus.Validators[0].(func(int) error)
}
