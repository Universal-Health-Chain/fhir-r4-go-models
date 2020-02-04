package models

type Appointment struct {
	DomainResource        `bson:",inline"`
	Identifier            []Identifier                      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status                string                            `bson:"status,omitempty" json:"status,omitempty"`
	ServiceCategory       *CodeableConcept                  `bson:"serviceCategory,omitempty" json:"serviceCategory,omitempty"`
	ServiceType           []CodeableConcept                 `bson:"serviceType,omitempty" json:"serviceType,omitempty"`
	Specialty             []CodeableConcept                 `bson:"specialty,omitempty" json:"specialty,omitempty"`
	AppointmentType       *CodeableConcept                  `bson:"appointmentType,omitempty" json:"appointmentType,omitempty"`
	Reason                []CodeableConcept                 `bson:"reason,omitempty" json:"reason,omitempty"`
	Indication            []Reference                       `bson:"indication,omitempty" json:"indication,omitempty"`
	Priority              *uint32                           `bson:"priority,omitempty" json:"priority,omitempty"`
	Description           string                            `bson:"description,omitempty" json:"description,omitempty"`
	SupportingInformation []Reference                       `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
	Start                 *FHIRDateTime                     `bson:"start,omitempty" json:"start,omitempty"`
	End                   *FHIRDateTime                     `bson:"end,omitempty" json:"end,omitempty"`
	MinutesDuration       *uint32                           `bson:"minutesDuration,omitempty" json:"minutesDuration,omitempty"`
	Slot                  []Reference                       `bson:"slot,omitempty" json:"slot,omitempty"`
	Created               *FHIRDateTime                     `bson:"created,omitempty" json:"created,omitempty"`
	Comment               string                            `bson:"comment,omitempty" json:"comment,omitempty"`
	IncomingReferral      []Reference                       `bson:"incomingReferral,omitempty" json:"incomingReferral,omitempty"`
	Participant           []AppointmentParticipantComponent `bson:"participant,omitempty" json:"participant,omitempty"`
	RequestedPeriod       []Period                          `bson:"requestedPeriod,omitempty" json:"requestedPeriod,omitempty"`
}

type AppointmentParticipantComponent struct {
	BackboneElement `bson:",inline"`
	Type            []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Actor           *Reference        `bson:"actor,omitempty" json:"actor,omitempty"`
	Required        string            `bson:"required,omitempty" json:"required,omitempty"`
	Status          string            `bson:"status,omitempty" json:"status,omitempty"`
}
