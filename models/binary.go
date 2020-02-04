package models

type Binary struct {
	Resource        `bson:",inline"`
	ContentType     string     `bson:"contentType,omitempty" json:"contentType,omitempty"`
	SecurityContext *Reference `bson:"securityContext,omitempty" json:"securityContext,omitempty"`
	Content         string     `bson:"content,omitempty" json:"content,omitempty"`
}
