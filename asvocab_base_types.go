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
	Target ObjectOrLink
	URL    []string
}

// StringWithOrderedCollectionPage can store a string URL pointing to an OrderedCollectionPage, which can itself be stored in the struct
type StringWithOrderedCollectionPage struct {
	OrdCollectionPage OrderedCollectionPage
	URL               string
}

// StringWithOrderedCollection can store a string URL pointing to an OrderedCollection, which can itself be stored in the struct
type StringWithOrderedCollection struct {
	OrdCollection OrderedCollection
	URL           string
}

// StringWithCollectionPage can store a string URL pointing to a CollectionPage, which can itself be stored in the struct
type StringWithCollectionPage struct {
	CollectionPage CollectionPage
	URL            string
}

// StringWithCollection can store a string URL pointing to a Collection, which can itself be stored in the struct
type StringWithCollection struct {
	Collection Collection
	URL        string
}

// Icons is a type representing the possible ActivityPub icon values
type Icons struct {
	Icons []Icon
	URL   string
}

type EndpointsOrString struct {
	Endpoints Endpoint
	URL       string
}

// https://www.w3.org/TR/activitystreams-vocabulary

// Object represents the base ActivityStreams Object and all of its properties
// Most of the other types extend Object
type Object struct {
	AlsoKnownAs  *ObjectOrLinkOrString `json:"alsoKnownAs,omitempty"`
	ASContext    any                   `json:"@context,omitempty"`
	ASLanguage   string                `json:"@language,omitempty"`
	AtID         string                `json:"@id,omitempty"`
	AtType       string                `json:"@type,omitempty"`
	Attachment   *ObjectOrLinkOrString `json:"attachment,omitempty"`
	AttributedTo *ObjectOrLinkOrString `json:"attributedTo,omitempty"`
	Audience     *ObjectOrLinkOrString `json:"audience,omitempty"`
	Bcc          *ObjectOrLinkOrString `json:"bcc,omitempty"`
	Bto          *ObjectOrLinkOrString `json:"bto,omitempty"`
	Cc           *ObjectOrLinkOrString `json:"cc,omitempty"`
	Content      string                `json:"content,omitempty"` // needs to be parsed safely ie by https://golang.org/pkg/html/template
	ContentMap   map[string]string     `json:"contentMap,omitempty"`
	Context      *ObjectOrLinkOrString `json:"context,omitempty"`
	Conversation string                `json:"conversation,omitempty"`
	Duration     string                `json:"duration,omitempty"` // xsd:duration
	EndTime      *time.Time            `json:"endTime,omitempty"`
	Generator    *ObjectOrLinkOrString `json:"generator,omitempty"`
	Icon         *ObjectOrLinkOrString `json:"icon,omitempty"`
	ID           string                `json:"id,omitempty"`
	Image        *ObjectOrLinkOrString `json:"image,omitempty"`
	InReplyTo    *ObjectOrLinkOrString `json:"inReplyTo,omitempty"`
	Likes        *ObjectOrLinkOrString `json:"likes,omitempty"`
	Location     *ObjectOrLinkOrString `json:"location,omitempty"`
	MediaType    string                `json:"mediaType,omitempty"`
	MovedTo      *ObjectOrLinkOrString `json:"movedTo,omitempty"`
	Name         string                `json:"name,omitempty"`
	NameMap      map[string]string     `json:"nameMap,omitempty"`
	Preview      *ObjectOrLinkOrString `json:"preview,omitempty"`
	Published    *time.Time            `json:"published,omitempty"`
	Replies      *OrderedCollection    `json:"replies,omitempty"`
	Schema       string                `json:"schema,omitempty"`
	Sensitive    bool                  `json:"sensitive,omitempty"`
	Shares       *ObjectOrLinkOrString `json:"shares,omitempty"`
	Source       *ObjectOrLinkOrString `json:"source,omitempty"`
	StartTime    *time.Time            `json:"startTime,omitempty"`
	Summary      string                `json:"summary,omitempty"`
	SummaryMap   map[string]string     `json:"summaryMap,omitempty"`
	Tag          *ObjectOrLinkOrString `json:"tag,omitempty"`
	To           *ObjectOrLinkOrString `json:"to,omitempty"`
	Type         string                `json:"type"`
	Updated      *time.Time            `json:"updated,omitempty"`
	URL          *ObjectOrLinkOrString `json:"url,omitempty"`
}

// Link represents the base ActivityStreams Link and all of its properties
// Other Link types extend it
type Link struct {
	ASContext  any                   `json:"@context,omitempty"`
	ASLanguage string                `json:"@language,omitempty"`
	Height     int                   `json:"height,omitempty"`
	Href       string                `json:"href,omitempty"`
	Hreflang   string                `json:"hreflang,omitempty"`
	MediaType  string                `json:"mediaType,omitempty"`
	Name       string                `json:"name,omitempty"`
	NameMap    map[string]string     `json:"nameMap,omitempty"`
	Preview    *ObjectOrLinkOrString `json:"preview,omitempty"`
	Published  *time.Time            `json:"published,omitempty"`
	Rel        []string              `json:"rel,omitempty"`
	Type       string                `json:"type"`
	Value      string                `json:"value,omitempty"`
	Width      int                   `json:"width,omitempty"`
}

type Location struct {
	Object

	Altitude  int     `json:"altitude,omitempty"`
	Latitude  float32 `json:"latitude,omitempty"`
	Longitude float32 `json:"longitude,omitempty"`
	Units     string  `json:"units,omitempty"`
}

type Endpoint struct {
	OauthAuthorizationEndpoint string                       `json:"oauthAuthorizationEndpoint,omitempty"`
	OauthTokenEndpoint         string                       `json:"oauthTokenEndpoint,omitempty"`
	ProvideClientKey           string                       `json:"provideClientKey,omitempty"`
	ProxyUrl                   string                       `json:"proxyUrl,omitempty"`
	SharedInbox                *StringWithOrderedCollection `json:"sharedInbox,omitempty"`
	SignClientKey              string                       `json:"signClientKey,omitempty"`
}

type Icon struct {
	Object

	Height int `json:"height"`
	Width  int `json:"width"`
}

type Actor struct {
	Object

	Devices                   *StringWithCollection        `json:"devices,omitempty"`
	Discoverable              *bool                        `json:"discoverable,omitempty"`
	EndpointsOrURI            *EndpointsOrString           `json:"endpoints,omitempty"`
	Featured                  *StringWithOrderedCollection `json:"featured,omitempty"`
	FeaturedTags              *StringWithCollection        `json:"featuredTags,omitempty"`
	Followers                 *StringWithOrderedCollection `json:"followers,omitempty"`
	Following                 *StringWithOrderedCollection `json:"following,omitempty"`
	Inbox                     *StringWithOrderedCollection `json:"inbox,omitempty"`
	Indexable                 *bool                        `json:"indexable,omitempty"`
	Liked                     *StringWithOrderedCollection `json:"liked,omitempty"`
	ManuallyApprovesFollowers *bool                        `json:"manuallyApprovesFollowers,omitempty"`
	Memorial                  *bool                        `json:"memorial,omitempty"`
	Outbox                    *StringWithOrderedCollection `json:"outbox,omitempty"`
	PreferredUsername         string                       `json:"preferredUsername,omitempty"`
	PublicKey                 *PublicKey                   `json:"publicKey,omitempty"`
	Streams                   *ObjectOrLinkOrString        `json:"streams,omitempty"`
}

type PublicKey struct {
	ID           string `json:"id"`
	Owner        string `json:"owner"`
	PublicKeyPem string `json:"publicKeyPem"`
}

type Activity struct {
	Object

	Actor          *ObjectOrLinkOrString `json:"actor,omitempty"`
	ActivityObject *ObjectOrLinkOrString `json:"object,omitempty"` // the 'object' property of Activity
	Instrument     *ObjectOrLinkOrString `json:"instrument,omitempty"`
	Origin         *ObjectOrLinkOrString `json:"origin,omitempty"`
	Result         *ObjectOrLinkOrString `json:"result,omitempty"`
	Target         *ObjectOrLinkOrString `json:"target,omitempty"`
}

type IntransitiveActivity struct {
	Object

	Actor      *ObjectOrLinkOrString `json:"actor,omitempty"`
	Instrument *ObjectOrLinkOrString `json:"instrument,omitempty"`
	Origin     *ObjectOrLinkOrString `json:"origin,omitempty"`
	Result     *ObjectOrLinkOrString `json:"result,omitempty"`
	Target     *ObjectOrLinkOrString `json:"target,omitempty"`
}

type CollectionPage struct {
	Collection

	Next   *ObjectOrLinkOrString `json:"next,omitempty"`
	PartOf *ObjectOrLinkOrString `json:"partOf,omitempty"`
	Prev   *ObjectOrLinkOrString `json:"prev,omitempty"`
}

type Collection struct {
	Object

	Current    *ObjectOrLinkOrString `json:"current,omitempty"`
	First      *ObjectOrLinkOrString `json:"first,omitempty"`
	Items      *ObjectOrLinkOrString `json:"items,omitempty"`
	Last       *ObjectOrLinkOrString `json:"last,omitempty"`
	TotalItems int                   `json:"totalItems,omitempty"`
}

// OrderedCollectionPage implements https://golang.org/pkg/sort/#Interface
type OrderedCollectionPage struct {
	OrderedCollection

	Next       *ObjectOrLinkOrString `json:"next,omitempty"`
	PartOf     *ObjectOrLinkOrString `json:"partOf,omitempty"`
	Prev       *ObjectOrLinkOrString `json:"prev,omitempty"`
	StartIndex uint                  `json:"startIndex,omitempty"`
}

// OrderedCollection implements https://golang.org/pkg/sort/#Interface
type OrderedCollection struct {
	Object

	Current      *StringWithOrderedCollectionPage `json:"current,omitempty"`
	First        *StringWithOrderedCollectionPage `json:"first,omitempty"`
	Last         *StringWithOrderedCollectionPage `json:"last,omitempty"`
	OrderedItems *ObjectOrLinkOrString            `json:"orderedItems"`
	TotalItems   int                              `json:"totalItems"`
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
