package models

import (
	"time"
)

type BiologicallyDerivedProduct struct {
	ResourceType    DomainResource                         `bson:"resourceType,omitempty" json:"resourceType,omitempty"`
	Identifier      []Identifier                           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	ProductCategory string                                 `bson:"productCategory,omitempty" json:"productCategory,omitempty"`
	ProductCode     CodeableConcept                        `bson:"productCode,omitempty" json:"productCode,omitempty"`
	Status          string                                 `bson:"status,omitempty" json:"status,omitempty"`
	Request         Reference                              `bson:"request,omitempty" json:"request,omitempty"`
	Quantity        int                                    `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Parent          []Reference                            `bson:"parent,omitempty" json:"parent,omitempty"`
	Collection      CollectionBiologicallyDerivedProduct   `bson:"collection,omitempty" json:"collection,omitempty"`
	Processing      []ProcessingBiologicallyDerivedProduct `bson:"processing,omitempty" json:"processing,omitempty"`
	Manipulation    ManipulationBiologicallyDerivedProduct `bson:"manipulation,omitempty" json:"manipulation,omitempty"`
	Storage         []StorageBiologicallyDerivedProduct    `bson:"storage,omitempty" json:"storage,omitempty"`
}

type CollectionBiologicallyDerivedProduct struct {
	BackboneElement
	Collector         Reference `bson:"collector,omitempty" json:"collector,omitempty"`
	Source            Reference `bson:"procedure,omitempty" json:"procedure,omitempty"`
	CollectedDateTime time.Time `bson:"collectedDateTime,omitempty" json:"collectedDateTime,omitempty"`
	CollectedPeriod   Period    `bson:"collectedPeriod,omitempty" json:"collectedPeriod,omitempty"`
}

type ProcessingBiologicallyDerivedProduct struct {
	BackboneElement
	Description  string          `bson:"description,omitempty" json:"description,omitempty"`
	Procedure    CodeableConcept `bson:"procedure,omitempty" json:"procedure,omitempty"`
	Additive     Reference       `bson:"additive,omitempty" json:"additive,omitempty"`
	TimeDateTime time.Time       `bson:"timeDateTime,omitempty" json:"timeDateTime,omitempty"`
	TimePeriod   Period          `bson:"timePeriod,omitempty" json:"timePeriod,omitempty"`
}

type ManipulationBiologicallyDerivedProduct struct {
	BackboneElement
	Description  string    `bson:"description,omitempty" json:"description,omitempty"`
	TimeDateTime time.Time `bson:"timeDateTime,omitempty" json:"timeDateTime,omitempty"`
	TimePeriod   Period    `bson:"timePeriod,omitempty" json:"timePeriod,omitempty"`
}

type StorageBiologicallyDerivedProduct struct {
	BackboneElement
	Description string `bson:"description,omitempty" json:"description,omitempty"`
	Temperature string `bson:"temperature,omitempty" json:"temperature,omitempty"`
	Scale       string `bson:"scale,omitempty" json:"scale,omitempty"`
	Duration    string `bson:"duration,omitempty" json:"duration,omitempty"`
}
