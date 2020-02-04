package models

type AdverseEvent struct {
	DomainResource        `bson:",inline"`
	Identifier            *Identifier                          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Category              string                               `bson:"category,omitempty" json:"category,omitempty"`
	Type                  *CodeableConcept                     `bson:"type,omitempty" json:"type,omitempty"`
	Subject               *Reference                           `bson:"subject,omitempty" json:"subject,omitempty"`
	Date                  *FHIRDateTime                        `bson:"date,omitempty" json:"date,omitempty"`
	Reaction              []Reference                          `bson:"reaction,omitempty" json:"reaction,omitempty"`
	Location              *Reference                           `bson:"location,omitempty" json:"location,omitempty"`
	Seriousness           *CodeableConcept                     `bson:"seriousness,omitempty" json:"seriousness,omitempty"`
	Outcome               *CodeableConcept                     `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Recorder              *Reference                           `bson:"recorder,omitempty" json:"recorder,omitempty"`
	EventParticipant      *Reference                           `bson:"eventParticipant,omitempty" json:"eventParticipant,omitempty"`
	Description           string                               `bson:"description,omitempty" json:"description,omitempty"`
	SuspectEntity         []AdverseEventSuspectEntityComponent `bson:"suspectEntity,omitempty" json:"suspectEntity,omitempty"`
	SubjectMedicalHistory []Reference                          `bson:"subjectMedicalHistory,omitempty" json:"subjectMedicalHistory,omitempty"`
	ReferenceDocument     []Reference                          `bson:"referenceDocument,omitempty" json:"referenceDocument,omitempty"`
	Study                 []Reference                          `bson:"study,omitempty" json:"study,omitempty"`
}

type AdverseEventSuspectEntityComponent struct {
	BackboneElement             `bson:",inline"`
	Instance                    *Reference       `bson:"instance,omitempty" json:"instance,omitempty"`
	Causality                   string           `bson:"causality,omitempty" json:"causality,omitempty"`
	CausalityAssessment         *CodeableConcept `bson:"causalityAssessment,omitempty" json:"causalityAssessment,omitempty"`
	CausalityProductRelatedness string           `bson:"causalityProductRelatedness,omitempty" json:"causalityProductRelatedness,omitempty"`
	CausalityMethod             *CodeableConcept `bson:"causalityMethod,omitempty" json:"causalityMethod,omitempty"`
	CausalityAuthor             *Reference       `bson:"causalityAuthor,omitempty" json:"causalityAuthor,omitempty"`
	CausalityResult             *CodeableConcept `bson:"causalityResult,omitempty" json:"causalityResult,omitempty"`
}
