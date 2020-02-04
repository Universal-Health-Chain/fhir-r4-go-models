package models

type AuditEvent struct {
	DomainResource `bson:",inline"`
	Type           *Coding                     `bson:"type,omitempty" json:"type,omitempty"`
	Subtype        []Coding                    `bson:"subtype,omitempty" json:"subtype,omitempty"`
	Action         string                      `bson:"action,omitempty" json:"action,omitempty"`
	Recorded       *FHIRDateTime               `bson:"recorded,omitempty" json:"recorded,omitempty"`
	Outcome        string                      `bson:"outcome,omitempty" json:"outcome,omitempty"`
	OutcomeDesc    string                      `bson:"outcomeDesc,omitempty" json:"outcomeDesc,omitempty"`
	PurposeOfEvent []CodeableConcept           `bson:"purposeOfEvent,omitempty" json:"purposeOfEvent,omitempty"`
	Agent          []AuditEventAgentComponent  `bson:"agent,omitempty" json:"agent,omitempty"`
	Source         *AuditEventSourceComponent  `bson:"source,omitempty" json:"source,omitempty"`
	Entity         []AuditEventEntityComponent `bson:"entity,omitempty" json:"entity,omitempty"`
}

type AuditEventAgentComponent struct {
	BackboneElement `bson:",inline"`
	Role            []CodeableConcept                `bson:"role,omitempty" json:"role,omitempty"`
	Reference       *Reference                       `bson:"reference,omitempty" json:"reference,omitempty"`
	UserId          *Identifier                      `bson:"userId,omitempty" json:"userId,omitempty"`
	AltId           string                           `bson:"altId,omitempty" json:"altId,omitempty"`
	Name            string                           `bson:"name,omitempty" json:"name,omitempty"`
	Requestor       *bool                            `bson:"requestor,omitempty" json:"requestor,omitempty"`
	Location        *Reference                       `bson:"location,omitempty" json:"location,omitempty"`
	Policy          []string                         `bson:"policy,omitempty" json:"policy,omitempty"`
	Media           *Coding                          `bson:"media,omitempty" json:"media,omitempty"`
	Network         *AuditEventAgentNetworkComponent `bson:"network,omitempty" json:"network,omitempty"`
	PurposeOfUse    []CodeableConcept                `bson:"purposeOfUse,omitempty" json:"purposeOfUse,omitempty"`
}

type AuditEventAgentNetworkComponent struct {
	BackboneElement `bson:",inline"`
	Address         string `bson:"address,omitempty" json:"address,omitempty"`
	Type            string `bson:"type,omitempty" json:"type,omitempty"`
}

type AuditEventSourceComponent struct {
	BackboneElement `bson:",inline"`
	Site            string      `bson:"site,omitempty" json:"site,omitempty"`
	Identifier      *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type            []Coding    `bson:"type,omitempty" json:"type,omitempty"`
}

type AuditEventEntityComponent struct {
	BackboneElement `bson:",inline"`
	Identifier      *Identifier                       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Reference       *Reference                        `bson:"reference,omitempty" json:"reference,omitempty"`
	Type            *Coding                           `bson:"type,omitempty" json:"type,omitempty"`
	Role            *Coding                           `bson:"role,omitempty" json:"role,omitempty"`
	Lifecycle       *Coding                           `bson:"lifecycle,omitempty" json:"lifecycle,omitempty"`
	SecurityLabel   []Coding                          `bson:"securityLabel,omitempty" json:"securityLabel,omitempty"`
	Name            string                            `bson:"name,omitempty" json:"name,omitempty"`
	Description     string                            `bson:"description,omitempty" json:"description,omitempty"`
	Query           string                            `bson:"query,omitempty" json:"query,omitempty"`
	Detail          []AuditEventEntityDetailComponent `bson:"detail,omitempty" json:"detail,omitempty"`
}

type AuditEventEntityDetailComponent struct {
	BackboneElement `bson:",inline"`
	Type            string `bson:"type,omitempty" json:"type,omitempty"`
	Value           string `bson:"value,omitempty" json:"value,omitempty"`
}
