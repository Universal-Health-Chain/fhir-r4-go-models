package models

import "time"

type BiologicallyDerivedProduct struct {
	ResourceType    DomainResource                         `bson:"resourceType" json:"resourceType"`
	Identifier      []Identifier                           `bson:"identifier" json:"identifier"`
	ProductCategory string                                 `bson:"productCategory" json:"productCategory"`
	ProductCode     CodeableConcept                        `bson:"productCode" json:"productCode"`
	Status          string                                 `bson:"status" json:"status"`
	Request         Reference                              `bson:"request" json:"request"`
	Quantity        int                                    `bson:"quantity" json:"quantity"`
	Parent          []Reference                            `bson:"parent" json:"parent"`
	Collection      CollectionBiologicallyDerivedProduct   `bson:"collection" json:"collection"`
	Processing      []ProcessingBiologicallyDerivedProduct `bson:"processing" json:"processing"`
	Manipulation    ManipulationBiologicallyDerivedProduct `bson:"manipulation" json:"manipulation"`
	Storage         []StorageBiologicallyDerivedProduct    `bson:"storage" json:"storage"`
}

type CollectionBiologicallyDerivedProduct struct {
	BackboneElement
	Collector         Reference `bson:"collector" json:"collector"`
	Source            Reference `bson:"procedure" json:"procedure"`
	CollectedDateTime time.Time `bson:"collectedDateTime" json:"collectedDateTime"`
	CollectedPeriod   Period    `bson:"collectedPeriod" json:"collectedPeriod"`
}

type ProcessingBiologicallyDerivedProduct struct {
	BackboneElement
	Description  string          `bson:"description" json:"description"`
	Procedure    CodeableConcept `bson:"procedure" json:"procedure"`
	Additive     Reference       `bson:"additive" json:"additive"`
	TimeDateTime time.Time       `bson:"timeDateTime" json:"timeDateTime"`
	TimePeriod   Period          `bson:"timePeriod" json:"timePeriod"`
}

type ManipulationBiologicallyDerivedProduct struct {
	BackboneElement
	Description  string    `bson:"description" json:"description"`
	TimeDateTime time.Time `bson:"timeDateTime" json:"timeDateTime"`
	TimePeriod   Period    `bson:"timePeriod" json:"timePeriod"`
}

type StorageBiologicallyDerivedProduct struct {
	BackboneElement
	Description string `bson:"description" json:"description"`
	Temperature string `bson:"temperature" json:"temperature"`
	Scale       string `bson:"scale" json:"scale"`
	Duration    string `bson:"duration" json:"duration"`
}
