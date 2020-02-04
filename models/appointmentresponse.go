package models

type AppointmentResponse struct {
	DomainResource    `bson:",inline"`
	Identifier        []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Appointment       *Reference        `bson:"appointment,omitempty" json:"appointment,omitempty"`
	Start             *FHIRDateTime     `bson:"start,omitempty" json:"start,omitempty"`
	End               *FHIRDateTime     `bson:"end,omitempty" json:"end,omitempty"`
	ParticipantType   []CodeableConcept `bson:"participantType,omitempty" json:"participantType,omitempty"`
	Actor             *Reference        `bson:"actor,omitempty" json:"actor,omitempty"`
	ParticipantStatus string            `bson:"participantStatus,omitempty" json:"participantStatus,omitempty"`
	Comment           string            `bson:"comment,omitempty" json:"comment,omitempty"`
}
