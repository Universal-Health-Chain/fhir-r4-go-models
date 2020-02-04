package models

type ActivityDefinition struct {
	DomainResource         `bson:",inline"`
	Url                    string                                    `bson:"url,omitempty" json:"url,omitempty"`
	Identifier             []Identifier                              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version                string                                    `bson:"version,omitempty" json:"version,omitempty"`
	Name                   string                                    `bson:"name,omitempty" json:"name,omitempty"`
	Title                  string                                    `bson:"title,omitempty" json:"title,omitempty"`
	Status                 string                                    `bson:"status,omitempty" json:"status,omitempty"`
	Experimental           *bool                                     `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date                   *FHIRDateTime                             `bson:"date,omitempty" json:"date,omitempty"`
	Publisher              string                                    `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Description            string                                    `bson:"description,omitempty" json:"description,omitempty"`
	Purpose                string                                    `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Usage                  string                                    `bson:"usage,omitempty" json:"usage,omitempty"`
	ApprovalDate           *FHIRDateTime                             `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	LastReviewDate         *FHIRDateTime                             `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	EffectivePeriod        *Period                                   `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	UseContext             []UsageContext                            `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept                         `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Topic                  []CodeableConcept                         `bson:"topic,omitempty" json:"topic,omitempty"`
	Contributor            []Contributor                             `bson:"contributor,omitempty" json:"contributor,omitempty"`
	Contact                []ContactDetail                           `bson:"contact,omitempty" json:"contact,omitempty"`
	Copyright              string                                    `bson:"copyright,omitempty" json:"copyright,omitempty"`
	RelatedArtifact        []RelatedArtifact                         `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	Library                []Reference                               `bson:"library,omitempty" json:"library,omitempty"`
	Kind                   string                                    `bson:"kind,omitempty" json:"kind,omitempty"`
	Code                   *CodeableConcept                          `bson:"code,omitempty" json:"code,omitempty"`
	TimingTiming           *Timing                                   `bson:"timingTiming,omitempty" json:"timingTiming,omitempty"`
	TimingDateTime         *FHIRDateTime                             `bson:"timingDateTime,omitempty" json:"timingDateTime,omitempty"`
	TimingPeriod           *Period                                   `bson:"timingPeriod,omitempty" json:"timingPeriod,omitempty"`
	TimingRange            *Range                                    `bson:"timingRange,omitempty" json:"timingRange,omitempty"`
	Location               *Reference                                `bson:"location,omitempty" json:"location,omitempty"`
	Participant            []ActivityDefinitionParticipantComponent  `bson:"participant,omitempty" json:"participant,omitempty"`
	ProductReference       *Reference                                `bson:"productReference,omitempty" json:"productReference,omitempty"`
	ProductCodeableConcept *CodeableConcept                          `bson:"productCodeableConcept,omitempty" json:"productCodeableConcept,omitempty"`
	Quantity               *Quantity                                 `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Dosage                 []Dosage                                  `bson:"dosage,omitempty" json:"dosage,omitempty"`
	BodySite               []CodeableConcept                         `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Transform              *Reference                                `bson:"transform,omitempty" json:"transform,omitempty"`
	DynamicValue           []ActivityDefinitionDynamicValueComponent `bson:"dynamicValue,omitempty" json:"dynamicValue,omitempty"`
}

type ActivityDefinitionParticipantComponent struct {
	BackboneElement `bson:",inline"`
	Type            string           `bson:"type,omitempty" json:"type,omitempty"`
	Role            *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
}

type ActivityDefinitionDynamicValueComponent struct {
	BackboneElement `bson:",inline"`
	Description     string `bson:"description,omitempty" json:"description,omitempty"`
	Path            string `bson:"path,omitempty" json:"path,omitempty"`
	Language        string `bson:"language,omitempty" json:"language,omitempty"`
	Expression      string `bson:"expression,omitempty" json:"expression,omitempty"`
}
