package models

type AllergyIntolerance struct {
	DomainResource     `bson:",inline"`
	Identifier         []Identifier                          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	ClinicalStatus     string                                `bson:"clinicalStatus,omitempty" json:"clinicalStatus,omitempty"`
	VerificationStatus string                                `bson:"verificationStatus,omitempty" json:"verificationStatus,omitempty"`
	Type               string                                `bson:"type,omitempty" json:"type,omitempty"`
	Category           []string                              `bson:"category,omitempty" json:"category,omitempty"`
	Criticality        string                                `bson:"criticality,omitempty" json:"criticality,omitempty"`
	Code               *CodeableConcept                      `bson:"code,omitempty" json:"code,omitempty"`
	Patient            *Reference                            `bson:"patient,omitempty" json:"patient,omitempty"`
	OnsetDateTime      *FHIRDateTime                         `bson:"onsetDateTime,omitempty" json:"onsetDateTime,omitempty"`
	OnsetAge           *Quantity                             `bson:"onsetAge,omitempty" json:"onsetAge,omitempty"`
	OnsetPeriod        *Period                               `bson:"onsetPeriod,omitempty" json:"onsetPeriod,omitempty"`
	OnsetRange         *Range                                `bson:"onsetRange,omitempty" json:"onsetRange,omitempty"`
	OnsetString        string                                `bson:"onsetString,omitempty" json:"onsetString,omitempty"`
	AssertedDate       *FHIRDateTime                         `bson:"assertedDate,omitempty" json:"assertedDate,omitempty"`
	Recorder           *Reference                            `bson:"recorder,omitempty" json:"recorder,omitempty"`
	Asserter           *Reference                            `bson:"asserter,omitempty" json:"asserter,omitempty"`
	LastOccurrence     *FHIRDateTime                         `bson:"lastOccurrence,omitempty" json:"lastOccurrence,omitempty"`
	Note               []Annotation                          `bson:"note,omitempty" json:"note,omitempty"`
	Reaction           []AllergyIntoleranceReactionComponent `bson:"reaction,omitempty" json:"reaction,omitempty"`
}

type AllergyIntoleranceReactionComponent struct {
	BackboneElement `bson:",inline"`
	Substance       *CodeableConcept  `bson:"substance,omitempty" json:"substance,omitempty"`
	Manifestation   []CodeableConcept `bson:"manifestation,omitempty" json:"manifestation,omitempty"`
	Description     string            `bson:"description,omitempty" json:"description,omitempty"`
	Onset           *FHIRDateTime     `bson:"onset,omitempty" json:"onset,omitempty"`
	Severity        string            `bson:"severity,omitempty" json:"severity,omitempty"`
	ExposureRoute   *CodeableConcept  `bson:"exposureRoute,omitempty" json:"exposureRoute,omitempty"`
	Note            []Annotation      `bson:"note,omitempty" json:"note,omitempty"`
}
