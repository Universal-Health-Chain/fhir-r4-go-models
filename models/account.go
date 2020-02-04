package models

type Account struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier                `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status         string                      `bson:"status,omitempty" json:"status,omitempty"`
	Type           *CodeableConcept            `bson:"type,omitempty" json:"type,omitempty"`
	Name           string                      `bson:"name,omitempty" json:"name,omitempty"`
	Subject        *Reference                  `bson:"subject,omitempty" json:"subject,omitempty"`
	Period         *Period                     `bson:"period,omitempty" json:"period,omitempty"`
	Active         *Period                     `bson:"active,omitempty" json:"active,omitempty"`
	Balance        *Quantity                   `bson:"balance,omitempty" json:"balance,omitempty"`
	Coverage       []AccountCoverageComponent  `bson:"coverage,omitempty" json:"coverage,omitempty"`
	Owner          *Reference                  `bson:"owner,omitempty" json:"owner,omitempty"`
	Description    string                      `bson:"description,omitempty" json:"description,omitempty"`
	Guarantor      []AccountGuarantorComponent `bson:"guarantor,omitempty" json:"guarantor,omitempty"`
}

type AccountCoverageComponent struct {
	BackboneElement `bson:",inline"`
	Coverage        *Reference `bson:"coverage,omitempty" json:"coverage,omitempty"`
	Priority        *uint32    `bson:"priority,omitempty" json:"priority,omitempty"`
}

type AccountGuarantorComponent struct {
	BackboneElement `bson:",inline"`
	Party           *Reference `bson:"party,omitempty" json:"party,omitempty"`
	OnHold          *bool      `bson:"onHold,omitempty" json:"onHold,omitempty"`
	Period          *Period    `bson:"period,omitempty" json:"period,omitempty"`
}
