package astreams

import "time"

// Extended 'Object' types
type Article = Object

type Document = Object

type Audio = Document

type Video = Document

type Image struct {
	Document

	Height int `json:"height,omitempty"`
	Width  int `json:"width,omitempty"`
}

type Event = Object

type Note = Object

type Page = Document

// Extended 'Link' types
type Mention = Link

type Hashtag = Link

type PropertyValue = Link

type Place struct {
	Location

	Accuracy float32 `json:"accuracy,omitempty"`
	Radius   float32 `json:"radius,omitempty"`
}

type Profile struct {
	Object

	Describes ObjectOrLinkOrString `json:"describes,omitempty"`
}

type Tombstone struct {
	Object

	Deleted    *time.Time `json:"deleted,omitempty"`
	FormerType string     `json:"formerType,omitempty"`
}

type Relationship struct {
	Object

	Relationship       string                `json:"relationship,omitempty"`
	RelationshipObject *ObjectOrLinkOrString `json:"object,omitempty"`
	Subject            *ObjectOrLinkOrString `json:"subject,omitempty"`
}

// Extended 'Actor' types
type Application = Actor

type Group = Actor

type Organization = Actor

type Person = Actor

type Service = Actor

// Extended 'Activity' types
type Accept = Activity

type Add = Activity

type Announce = Activity

type Arrive = IntransitiveActivity

type Ignore = Activity

type Block = Ignore

type Create = Activity

type Delete = Activity

type Dislike = Activity

type Flag = Activity

type Follow = Activity

type Offer = Activity

type Invite = Offer

type Join = Activity

type Leave = Activity

type Like = Activity

type Listen = Activity

type Move = Activity

type Question struct {
	IntransitiveActivity

	// only OneOf or AnyOf can be set, not both
	AnyOf  *ObjectOrLinkOrString `json:"anyOf,omitempty"`
	Closed *time.Time            `json:"closed,omitempty"`
	OneOf  *ObjectOrLinkOrString `json:"oneOf,omitempty"`
}

type Reject = Activity

type Read = Activity

type Remove = Activity

type TentativeReject = Reject

type TentativeAccept = Accept

type Travel = IntransitiveActivity

type Undo = Activity

type Update = Activity

type View = Activity
