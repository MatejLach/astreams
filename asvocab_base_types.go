package astreams

import (
	"sort"
	"time"
)

// ObjectOrLink is a wrapper type that represents any valid ActivityStreams 2.0 object or link
// Single objects are still presented as slices, but with a single element
type ObjectOrLink []ObjectLinker

// ObjectOrLinkOrString is a type that can either represent simple string URL(s) or an Object/Link slice
type ObjectOrLinkOrString struct {
	URL    []string
	Target ObjectOrLink
}

// StringWithOrderedCollectionPage can store a string URL pointing to an OrderedCollectionPage, which can itself be stored in the struct
type StringWithOrderedCollectionPage struct {
	URL               string
	OrdCollectionPage OrderedCollectionPage
}

// StringWithOrderedCollection can store a string URL pointing to an OrderedCollection, which can itself be stored in the struct
type StringWithOrderedCollection struct {
	URL           string
	OrdCollection OrderedCollection
}

// StringWithCollectionPage can store a string URL pointing to a CollectionPage, which can itself be stored in the struct
type StringWithCollectionPage struct {
	URL            string
	CollectionPage CollectionPage
}

// StringWithCollection can store a string URL pointing to a Collection, which can itself be stored in the struct
type StringWithCollection struct {
	URL        string
	Collection Collection
}

// Icons is a type representing the possible ActivityPub icon values
type Icons struct {
	URL   string
	Icons []Icon
}

type EndpointsOrString struct {
	URL       string
	Endpoints Endpoint
}

// https://www.w3.org/TR/activitystreams-vocabulary

// Object represents the base ActivityStreams Object and all of its properties
// Most of the other types extend Object
type Object struct {
	ASContext    *ObjectOrLinkOrString `json:"@context,omitempty"`
	ASLanguage   string                `json:"@language,omitempty"`
	Schema       string                `json:"schema,omitempty"`
	ID           string                `json:"id,omitempty"`
	Type         string                `json:"type,omitempty"`
	AtType       string                `json:"@type,omitempty"`
	Summary      string                `json:"summary,omitempty"`
	SummaryMap   map[string]string     `json:"summaryMap,omitempty"`
	AtID         string                `json:"@id,omitempty"`
	Name         string                `json:"name,omitempty"`
	NameMap      map[string]string     `json:"nameMap,omitempty"`
	AlsoKnownAs  *ObjectOrLinkOrString `json:"alsoKnownAs,omitempty"`
	MovedTo      *ObjectOrLinkOrString `json:"movedTo,omitempty"`
	Attachment   *ObjectOrLinkOrString `json:"attachment,omitempty"`
	AttributedTo *ObjectOrLinkOrString `json:"attributedTo,omitempty"`
	Audience     *ObjectOrLinkOrString `json:"audience,omitempty"`
	Content      string                `json:"content,omitempty"` // needs to be parsed safely ie by https://golang.org/pkg/html/template
	ContentMap   map[string]string     `json:"contentMap,omitempty"`
	Source       *ObjectOrLinkOrString `json:"source,omitempty"`
	Context      *ObjectOrLinkOrString `json:"context,omitempty"`
	StartTime    *time.Time            `json:"startTime,omitempty"`
	EndTime      *time.Time            `json:"endTime,omitempty"`
	Generator    *ObjectOrLinkOrString `json:"generator,omitempty"`
	Featured     *ObjectOrLinkOrString `json:"featured,omitempty"`
	Likes        *ObjectOrLinkOrString `json:"likes,omitempty"`
	Shares       *ObjectOrLinkOrString `json:"shares,omitempty"`
	Icon         *ObjectOrLinkOrString `json:"icon,omitempty"`
	Image        *ObjectOrLinkOrString `json:"image,omitempty"`
	InReplyTo    *ObjectOrLinkOrString `json:"inReplyTo,omitempty"`
	Location     *ObjectOrLinkOrString `json:"location,omitempty"`
	Preview      *ObjectOrLinkOrString `json:"preview,omitempty"`
	Published    *time.Time            `json:"published,omitempty"`
	Replies      *OrderedCollection    `json:"replies,omitempty"`
	Tag          *ObjectOrLinkOrString `json:"tag,omitempty"`
	Updated      *time.Time            `json:"updated,omitempty"`
	URL          *ObjectOrLinkOrString `json:"url,omitempty"`
	To           *ObjectOrLinkOrString `json:"to,omitempty"`
	Bto          *ObjectOrLinkOrString `json:"bto,omitempty"`
	Cc           *ObjectOrLinkOrString `json:"cc,omitempty"`
	Bcc          *ObjectOrLinkOrString `json:"bcc,omitempty"`
	MediaType    string                `json:"mediaType,omitempty"`
	Duration     string                `json:"duration,omitempty"` // xsd:duration
}

// Link represents the base ActivityStreams Link and all of its properties
// Other Link types extend it
type Link struct {
	ASContext  *ObjectOrLinkOrString `json:"@context,omitempty"`
	ASLanguage string                `json:"@language,omitempty"`
	Type       string                `json:"type"`
	Href       string                `json:"href,omitempty"`
	Hreflang   string                `json:"hreflang,omitempty"`
	Rel        []string              `json:"rel,omitempty"`
	MediaType  string                `json:"mediaType,omitempty"`
	Height     int                   `json:"height,omitempty"`
	Width      int                   `json:"width,omitempty"`
	Name       string                `json:"name,omitempty"`
	NameMap    map[string]string     `json:"nameMap,omitempty"`
	Value      string                `json:"value,omitempty"`
	Preview    *ObjectOrLinkOrString `json:"preview,omitempty"`
	Published  *time.Time            `json:"published,omitempty"`
}

type Location struct {
	Object
	Longitude float32 `json:"longitude,omitempty"`
	Latitude  float32 `json:"latitude,omitempty"`
	Altitude  int     `json:"altitude,omitempty"`
	Units     string  `json:"units,omitempty"`
}

type Endpoint struct {
	ProxyUrl                   string                       `json:"proxyUrl,omitempty"`
	OauthAuthorizationEndpoint string                       `json:"oauthAuthorizationEndpoint,omitempty"`
	OauthTokenEndpoint         string                       `json:"oauthTokenEndpoint,omitempty"`
	ProvideClientKey           string                       `json:"provideClientKey,omitempty"`
	SignClientKey              string                       `json:"signClientKey,omitempty"`
	SharedInbox                *StringWithOrderedCollection `json:"sharedInbox,omitempty"`
}

type Icon struct {
	Object
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Actor struct {
	Object
	PublicKey                 *PublicKey                       `json:"publicKey,omitempty"`
	Inbox                     *StringWithOrderedCollectionPage `json:"inbox,omitempty"`
	Outbox                    *StringWithOrderedCollectionPage `json:"outbox,omitempty"`
	Followers                 *StringWithOrderedCollectionPage `json:"followers,omitempty"`
	Following                 *StringWithOrderedCollectionPage `json:"following,omitempty"`
	Liked                     *StringWithOrderedCollectionPage `json:"liked,omitempty"`
	Streams                   *ObjectOrLinkOrString            `json:"streams,omitempty"`
	PreferredUsername         string                           `json:"preferredUsername,omitempty"`
	ManuallyApprovesFollowers bool                             `json:"manuallyApprovesFollowers,omitempty"`
	EndpointsOrURI            *EndpointsOrString               `json:"endpoints,omitempty"`
}

type Activity struct {
	Object
	Actor          *ObjectOrLinkOrString `json:"actor,omitempty"`
	ActivityObject *ObjectOrLinkOrString `json:"object,omitempty"` // the 'object' property of Activity
	Target         *ObjectOrLinkOrString `json:"target,omitempty"`
	Result         *ObjectOrLinkOrString `json:"result,omitempty"`
	Origin         *ObjectOrLinkOrString `json:"origin,omitempty"`
	Instrument     *ObjectOrLinkOrString `json:"instrument,omitempty"`
}

type IntransitiveActivity struct {
	Object
	Actor      *ObjectOrLinkOrString `json:"actor,omitempty"`
	Target     *ObjectOrLinkOrString `json:"target,omitempty"`
	Result     *ObjectOrLinkOrString `json:"result,omitempty"`
	Origin     *ObjectOrLinkOrString `json:"origin,omitempty"`
	Instrument *ObjectOrLinkOrString `json:"instrument,omitempty"`
}

// CollectionPage is provided for spec compliance, prefer OrderedCollectionPage
type CollectionPage struct {
	Collection
	PartOf *ObjectOrLinkOrString `json:"partOf,omitempty"`
	Next   *ObjectOrLinkOrString `json:"next,omitempty"`
	Prev   *ObjectOrLinkOrString `json:"prev,omitempty"`
}

// Collection is provided for spec compliance, prefer OrderedCollection
type Collection struct {
	Object
	TotalItems int                   `json:"totalItems,omitempty"`
	Current    *ObjectOrLinkOrString `json:"current,omitempty"`
	First      *ObjectOrLinkOrString `json:"first,omitempty"`
	Last       *ObjectOrLinkOrString `json:"last,omitempty"`
	Items      *ObjectOrLinkOrString `json:"items,omitempty"`
}

// OrderedCollectionPage implements https://golang.org/pkg/sort/#Interface
type OrderedCollectionPage struct {
	OrderedCollection
	PartOf     *ObjectOrLinkOrString `json:"partOf,omitempty"`
	Next       *ObjectOrLinkOrString `json:"next,omitempty"`
	Prev       *ObjectOrLinkOrString `json:"prev,omitempty"`
	StartIndex uint                  `json:"startIndex,omitempty"`
}

// OrderedCollection implements https://golang.org/pkg/sort/#Interface
type OrderedCollection struct {
	Object
	TotalItems   int                   `json:"totalItems"`
	Current      *ObjectOrLinkOrString `json:"current,omitempty"`
	First        *ObjectOrLinkOrString `json:"first,omitempty"`
	Last         *ObjectOrLinkOrString `json:"last,omitempty"`
	OrderedItems *ObjectOrLinkOrString `json:"orderedItems"`
}

func (oc OrderedCollection) Len() int {
	if len(oc.OrderedItems.URL) > 0 {
		return len(oc.OrderedItems.URL)
	}

	return len(oc.OrderedItems.Target)
}

func (oc OrderedCollection) Less(i, j int) bool {
	if len(oc.OrderedItems.Target) > 0 {
		if oc.OrderedItems.Target[i].IsObject() && oc.OrderedItems.Target[j].IsObject() {
			return oc.OrderedItems.Target[j].GetObject().Published.Before(*oc.OrderedItems.Target[i].GetObject().Published)
		} else if oc.OrderedItems.Target[i].IsObject() && oc.OrderedItems.Target[j].IsLink() {
			return oc.OrderedItems.Target[j].GetObject().Published.Before(*oc.OrderedItems.Target[i].GetLink().Published)
		} else if oc.OrderedItems.Target[i].IsLink() && oc.OrderedItems.Target[j].IsLink() {
			return oc.OrderedItems.Target[j].GetLink().Published.Before(*oc.OrderedItems.Target[i].GetLink().Published)
		} else if oc.OrderedItems.Target[i].IsLink() && oc.OrderedItems.Target[j].IsObject() {
			return oc.OrderedItems.Target[j].GetLink().Published.Before(*oc.OrderedItems.Target[i].GetObject().Published)
		}
	}

	return false
}

func (oc OrderedCollection) Swap(i, j int) {
	if len(oc.OrderedItems.URL) > 0 {
		oc.OrderedItems.URL[i], oc.OrderedItems.URL[j] = oc.OrderedItems.URL[j], oc.OrderedItems.URL[i]
	} else {
		oc.OrderedItems.Target[i], oc.OrderedItems.Target[j] = oc.OrderedItems.Target[j], oc.OrderedItems.Target[i]
	}
}

// SortByUpdated sorts OrderedCollection objects by Updated rather than Published date
func (oc OrderedCollection) SortByUpdated() {
	sort.Slice(oc.OrderedItems.Target, func(i, j int) bool {
		if len(oc.OrderedItems.Target) > 0 {
			if oc.OrderedItems.Target[i].IsObject() && oc.OrderedItems.Target[j].IsObject() {
				return oc.OrderedItems.Target[j].GetObject().Updated.Before(*oc.OrderedItems.Target[i].GetObject().Updated)
			} else if oc.OrderedItems.Target[i].IsObject() && oc.OrderedItems.Target[j].IsLink() {
				return oc.OrderedItems.Target[j].GetObject().Updated.Before(*oc.OrderedItems.Target[i].GetLink().Published)
			} else if oc.OrderedItems.Target[i].IsLink() && oc.OrderedItems.Target[j].IsLink() {
				return oc.OrderedItems.Target[j].GetLink().Published.Before(*oc.OrderedItems.Target[i].GetLink().Published)
			} else if oc.OrderedItems.Target[i].IsLink() && oc.OrderedItems.Target[j].IsObject() {
				return oc.OrderedItems.Target[j].GetLink().Published.Before(*oc.OrderedItems.Target[i].GetObject().Updated)
			}
		}

		return false

	})
}
