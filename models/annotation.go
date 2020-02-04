package models

type Annotation struct {
	AuthorReference *Reference    `bson:"authorReference,omitempty" json:"authorReference,omitempty"`
	AuthorString    string        `bson:"authorString,omitempty" json:"authorString,omitempty"`
	Time            *FHIRDateTime `bson:"time,omitempty" json:"time,omitempty"`
	// text changed from string to markdown in FHIR R4
	Text string `bson:"text,omitempty" json:"text,omitempty"`
}
