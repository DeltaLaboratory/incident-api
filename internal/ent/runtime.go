// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/DeltaLaboratory/incident-api/internal/ent/help"
	"github.com/DeltaLaboratory/incident-api/internal/ent/incident"
	"github.com/DeltaLaboratory/incident-api/internal/ent/schema"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	helpFields := schema.Help{}.Fields()
	_ = helpFields
	// helpDescCreatedAt is the schema descriptor for created_at field.
	helpDescCreatedAt := helpFields[8].Descriptor()
	// help.DefaultCreatedAt holds the default value on creation for the created_at field.
	help.DefaultCreatedAt = helpDescCreatedAt.Default.(func() time.Time)
	// helpDescID is the schema descriptor for id field.
	helpDescID := helpFields[0].Descriptor()
	// help.DefaultID holds the default value on creation for the id field.
	help.DefaultID = helpDescID.Default.(func() uuid.UUID)
	incidentFields := schema.Incident{}.Fields()
	_ = incidentFields
	// incidentDescImage is the schema descriptor for image field.
	incidentDescImage := incidentFields[6].Descriptor()
	// incident.DefaultImage holds the default value on creation for the image field.
	incident.DefaultImage = incidentDescImage.Default.([]byte)
	// incidentDescVote is the schema descriptor for vote field.
	incidentDescVote := incidentFields[7].Descriptor()
	// incident.DefaultVote holds the default value on creation for the vote field.
	incident.DefaultVote = incidentDescVote.Default.(int)
	// incidentDescVoteFilter is the schema descriptor for vote_filter field.
	incidentDescVoteFilter := incidentFields[8].Descriptor()
	// incident.DefaultVoteFilter holds the default value on creation for the vote_filter field.
	incident.DefaultVoteFilter = incidentDescVoteFilter.Default.(func() []byte)
	// incidentDescCreatedAt is the schema descriptor for created_at field.
	incidentDescCreatedAt := incidentFields[9].Descriptor()
	// incident.DefaultCreatedAt holds the default value on creation for the created_at field.
	incident.DefaultCreatedAt = incidentDescCreatedAt.Default.(func() time.Time)
	// incidentDescID is the schema descriptor for id field.
	incidentDescID := incidentFields[0].Descriptor()
	// incident.DefaultID holds the default value on creation for the id field.
	incident.DefaultID = incidentDescID.Default.(func() uuid.UUID)
}
