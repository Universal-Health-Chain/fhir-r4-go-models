// Copyright (c) 2011-2017, HL7, Inc & The MITRE Corporation
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification,
// are permitted provided that the following conditions are met:
//
//     * Redistributions of source code must retain the above copyright notice, this
//       list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above copyright notice,
//       this list of conditions and the following disclaimer in the documentation
//       and/or other materials provided with the distribution.
//     * Neither the name of HL7 nor the names of its contributors may be used to
//       endorse or promote products derived from this software without specific
//       prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
// IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT,
// INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
// NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
// PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
// WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Observation struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn        []Reference  `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	// added partOf in FHIR R4
	PartOf   *Reference        `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status   string            `bson:"status,omitempty" json:"status,omitempty"`
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	Code     *CodeableConcept  `bson:"code,omitempty" json:"code,omitempty"`
	Subject  *Reference        `bson:"subject,omitempty" json:"subject,omitempty"`
	// added focus in FHIR r4
	Focus *Reference `bson:"focus,omitempty" json:"focus,omitempty"`
	// context is renamed to encounter in FHIR R4
	// Context			*Reference			`bson:"context,omitempty" json:"context,omitempty"`
	Encounter         *Reference    `bson:"encounter,omitempty" json:"encounter,omitempty"`
	EffectiveDateTime *FHIRDateTime `bson:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty"`
	EffectivePeriod   *Period       `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	// added effectiveTiming and effectiveInstant
	EffectiveTiming      *Timing          `bson:"effectiveTiming,omitempty" json:"effectiveTiming,omitempty"`
	EffectiveInstant     *FHIRDateTime    `bson:"effectiveInstant,omitempty" json:"effectiveInstant,omitempty"`
	Issued               *FHIRDateTime    `bson:"issued,omitempty" json:"issued,omitempty"`
	Performer            []Reference      `bson:"performer,omitempty" json:"performer,omitempty"`
	ValueQuantity        *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueCodeableConcept *CodeableConcept `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueString          string           `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueBoolean         *bool            `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	// valueInteger added in FHIR R4
	ValueInteger     *int         `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueRange       *Range       `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
	ValueRatio       *Ratio       `bson:"valueRatio,omitempty" json:"valueRatio,omitempty"`
	ValueSampledData *SampledData `bson:"valueSampledData,omitempty" json:"valueSampledData,omitempty"`
	// ValueAttachment  *Attachment      `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueTime        *FHIRDateTime      `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueDateTime    *FHIRDateTime      `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValuePeriod      *Period            `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty"`
	DataAbsentReason *CodeableConcept   `bson:"dataAbsentReason,omitempty" json:"dataAbsentReason,omitempty"`
	Interpretation   []*CodeableConcept `bson:"interpretation,omitempty" json:"interpretation,omitempty"`
	// comment is renamed to note in FHIR R4
	// Comment              string           `bson:"comment,omitempty" json:"comment,omitempty"`
	Note           string                               `bson:"note,omitempty" json:"note,omitempty"`
	BodySite       *CodeableConcept                     `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Method         *CodeableConcept                     `bson:"method,omitempty" json:"method,omitempty"`
	Specimen       *Reference                           `bson:"specimen,omitempty" json:"specimen,omitempty"`
	Device         *Reference                           `bson:"device,omitempty" json:"device,omitempty"`
	ReferenceRange []ObservationReferenceRangeComponent `bson:"referenceRange,omitempty" json:"referenceRange,omitempty"`
	// Related              []ObservationRelatedComponent        `bson:"related,omitempty" json:"related,omitempty"`
	// added hasMember in FHIR R4
	HasMemeber *Reference `bson:"hasMember,omitempty" json:"hasMember,omitempty"`
	// added derivedFrom in FHIR R4
	DerivedFrom *Reference                      `bson:"derivedFrom,omitempty" json:"derivedFrom,omitempty"`
	Component   []ObservationComponentComponent `bson:"component,omitempty" json:"component,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Observation) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Observation"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Observation), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Observation) GetBSON() (interface{}, error) {
	x.ResourceType = "Observation"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "observation" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type observation Observation

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Observation) UnmarshalJSON(data []byte) (err error) {
	x2 := observation{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i], err = MapToResource(x2.Contained[i], true)
				if err != nil {
					return err
				}
			}
		}
		*x = Observation(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Observation) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Observation"
	} else if x.ResourceType != "Observation" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Observation, instead received %s", x.ResourceType))
	}
	return nil
}

type ObservationReferenceRangeComponent struct {
	BackboneElement `bson:",inline"`
	Low             *Quantity         `bson:"low,omitempty" json:"low,omitempty"`
	High            *Quantity         `bson:"high,omitempty" json:"high,omitempty"`
	Type            *CodeableConcept  `bson:"type,omitempty" json:"type,omitempty"`
	AppliesTo       []CodeableConcept `bson:"appliesTo,omitempty" json:"appliesTo,omitempty"`
	Age             *Range            `bson:"age,omitempty" json:"age,omitempty"`
	Text            string            `bson:"text,omitempty" json:"text,omitempty"`
}

type ObservationRelatedComponent struct {
	BackboneElement `bson:",inline"`
	Type            string     `bson:"type,omitempty" json:"type,omitempty"`
	Target          *Reference `bson:"target,omitempty" json:"target,omitempty"`
}

type ObservationComponentComponent struct {
	BackboneElement      `bson:",inline"`
	Code                 *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	ValueQuantity        *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueCodeableConcept *CodeableConcept `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueString          string           `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueBoolean         *bool            `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	// valueInteger added in FHIR R4
	ValueInteger     *int         `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueRange       *Range       `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
	ValueRatio       *Ratio       `bson:"valueRatio,omitempty" json:"valueRatio,omitempty"`
	ValueSampledData *SampledData `bson:"valueSampledData,omitempty" json:"valueSampledData,omitempty"`
	// ValueAttachment      *Attachment                          `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueTime        *FHIRDateTime                        `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueDateTime    *FHIRDateTime                        `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValuePeriod      *Period                              `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty"`
	DataAbsentReason *CodeableConcept                     `bson:"dataAbsentReason,omitempty" json:"dataAbsentReason,omitempty"`
	Interpretation   *CodeableConcept                     `bson:"interpretation,omitempty" json:"interpretation,omitempty"`
	ReferenceRange   []ObservationReferenceRangeComponent `bson:"referenceRange,omitempty" json:"referenceRange,omitempty"`
}
