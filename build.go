package astreams

import (
	"time"
)

// ArticleBuilder
type ArticleBuilder struct {
	article Article
}

func NewArticle() *ArticleBuilder {
	return &ArticleBuilder{
		article: Article{Type: "Article"},
	}
}

func (ab *ArticleBuilder) ID(id string) *ArticleBuilder {
	ab.article.ID = id
	return ab
}

func (ab *ArticleBuilder) Name(name string) *ArticleBuilder {
	ab.article.Name = name
	return ab
}

func (ab *ArticleBuilder) Content(content string) *ArticleBuilder {
	ab.article.Content = content
	return ab
}

func (ab *ArticleBuilder) Published(t *time.Time) *ArticleBuilder {
	ab.article.Published = t
	return ab
}

func (ab *ArticleBuilder) Build() *Article {
	return &ab.article
}

// DocumentBuilder
type DocumentBuilder struct {
	document Document
}

func NewDocument() *DocumentBuilder {
	return &DocumentBuilder{
		document: Document{Type: "Document"},
	}
}

func (db *DocumentBuilder) ID(id string) *DocumentBuilder {
	db.document.ID = id
	return db
}

func (db *DocumentBuilder) Name(name string) *DocumentBuilder {
	db.document.Name = name
	return db
}

func (db *DocumentBuilder) Content(content string) *DocumentBuilder {
	db.document.Content = content
	return db
}

func (db *DocumentBuilder) Published(t *time.Time) *DocumentBuilder {
	db.document.Published = t
	return db
}

func (db *DocumentBuilder) Build() *Document {
	return &db.document
}

// AudioBuilder
type AudioBuilder struct {
	audio Audio
}

func NewAudio() *AudioBuilder {
	return &AudioBuilder{
		audio: Audio{Type: "Audio"},
	}
}

func (ab *AudioBuilder) ID(id string) *AudioBuilder {
	ab.audio.ID = id
	return ab
}

func (ab *AudioBuilder) Name(name string) *AudioBuilder {
	ab.audio.Name = name
	return ab
}

func (ab *AudioBuilder) Content(content string) *AudioBuilder {
	ab.audio.Content = content
	return ab
}

func (ab *AudioBuilder) Published(t *time.Time) *AudioBuilder {
	ab.audio.Published = t
	return ab
}

func (ab *AudioBuilder) Build() *Audio {
	return &ab.audio
}

// VideoBuilder
type VideoBuilder struct {
	video Video
}

func NewVideo() *VideoBuilder {
	return &VideoBuilder{
		video: Video{Type: "Video"},
	}
}

func (vb *VideoBuilder) ID(id string) *VideoBuilder {
	vb.video.ID = id
	return vb
}

func (vb *VideoBuilder) Name(name string) *VideoBuilder {
	vb.video.Name = name
	return vb
}

func (vb *VideoBuilder) Content(content string) *VideoBuilder {
	vb.video.Content = content
	return vb
}

func (vb *VideoBuilder) Published(t *time.Time) *VideoBuilder {
	vb.video.Published = t
	return vb
}

func (vb *VideoBuilder) Build() *Video {
	return &vb.video
}

// ImageBuilder
type ImageBuilder struct {
	image Image
}

func NewImage() *ImageBuilder {
	return &ImageBuilder{
		image: Image{Document: Document{Type: "Image"}},
	}
}

func (ib *ImageBuilder) ID(id string) *ImageBuilder {
	ib.image.ID = id
	return ib
}

func (ib *ImageBuilder) Name(name string) *ImageBuilder {
	ib.image.Name = name
	return ib
}

func (ib *ImageBuilder) Published(t *time.Time) *ImageBuilder {
	ib.image.Published = t
	return ib
}

func (ib *ImageBuilder) Width(width int) *ImageBuilder {
	ib.image.Width = width
	return ib
}

func (ib *ImageBuilder) Height(height int) *ImageBuilder {
	ib.image.Height = height
	return ib
}

func (ib *ImageBuilder) Build() *Image {
	return &ib.image
}

// EventBuilder
type EventBuilder struct {
	event Event
}

func NewEvent() *EventBuilder {
	return &EventBuilder{
		event: Event{Type: "Event"},
	}
}

func (eb *EventBuilder) ID(id string) *EventBuilder {
	eb.event.ID = id
	return eb
}

func (eb *EventBuilder) Name(name string) *EventBuilder {
	eb.event.Name = name
	return eb
}

func (eb *EventBuilder) Content(content string) *EventBuilder {
	eb.event.Content = content
	return eb
}

func (eb *EventBuilder) Published(t *time.Time) *EventBuilder {
	eb.event.Published = t
	return eb
}

func (eb *EventBuilder) StartTime(t *time.Time) *EventBuilder {
	eb.event.StartTime = t
	return eb
}

func (eb *EventBuilder) EndTime(t *time.Time) *EventBuilder {
	eb.event.EndTime = t
	return eb
}

func (eb *EventBuilder) Build() *Event {
	return &eb.event
}

// NoteBuilder
type NoteBuilder struct {
	note Note
}

func NewNote() *NoteBuilder {
	return &NoteBuilder{
		note: Note{Type: "Note"},
	}
}

func (nb *NoteBuilder) ID(id string) *NoteBuilder {
	nb.note.ID = id
	return nb
}

func (nb *NoteBuilder) Content(content string) *NoteBuilder {
	nb.note.Content = content
	return nb
}

func (nb *NoteBuilder) Published(t *time.Time) *NoteBuilder {
	nb.note.Published = t
	return nb
}

func (nb *NoteBuilder) To(urls ...string) *NoteBuilder {
	if len(urls) > 0 {
		targets := make(ObjectOrLink, len(urls))
		for i, url := range urls {
			targets[i] = Link{Type: "Link", Href: url}
		}
		nb.note.To = &ObjectOrLinkOrString{Target: targets}
	}
	return nb
}

func (nb *NoteBuilder) Build() *Note {
	return &nb.note
}

// PageBuilder
type PageBuilder struct {
	page Page
}

func NewPage() *PageBuilder {
	return &PageBuilder{
		page: Page{Type: "Page"},
	}
}

func (pb *PageBuilder) ID(id string) *PageBuilder {
	pb.page.ID = id
	return pb
}

func (pb *PageBuilder) Name(name string) *PageBuilder {
	pb.page.Name = name
	return pb
}

func (pb *PageBuilder) Content(content string) *PageBuilder {
	pb.page.Content = content
	return pb
}

func (pb *PageBuilder) Published(t *time.Time) *PageBuilder {
	pb.page.Published = t
	return pb
}

func (pb *PageBuilder) Build() *Page {
	return &pb.page
}

// CollectionBuilder
type CollectionBuilder struct {
	collection Collection
}

func NewCollection() *CollectionBuilder {
	return &CollectionBuilder{
		collection: Collection{Object: Object{Type: "Collection"}},
	}
}

func (cb *CollectionBuilder) ID(id string) *CollectionBuilder {
	cb.collection.ID = id
	return cb
}

func (cb *CollectionBuilder) TotalItems(totalItems int) *CollectionBuilder {
	cb.collection.TotalItems = totalItems
	return cb
}

func (cb *CollectionBuilder) First(url string) *CollectionBuilder {
	cb.collection.First = &StringWithCollectionPage{URL: url}
	return cb
}

func (cb *CollectionBuilder) Last(url string) *CollectionBuilder {
	cb.collection.Last = &StringWithCollectionPage{URL: url}
	return cb
}

func (cb *CollectionBuilder) Build() *Collection {
	return &cb.collection
}

// CollectionPageBuilder
type CollectionPageBuilder struct {
	collectionPage CollectionPage
}

func NewCollectionPage() *CollectionPageBuilder {
	return &CollectionPageBuilder{
		collectionPage: CollectionPage{Collection: Collection{Object: Object{Type: "CollectionPage"}}},
	}
}

func (cpb *CollectionPageBuilder) ID(id string) *CollectionPageBuilder {
	cpb.collectionPage.ID = id
	return cpb
}

func (cpb *CollectionPageBuilder) PartOf(url string) *CollectionPageBuilder {
	cpb.collectionPage.PartOf = &ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return cpb
}

func (cpb *CollectionPageBuilder) Next(url string) *CollectionPageBuilder {
	cpb.collectionPage.Next = &ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return cpb
}

func (cpb *CollectionPageBuilder) Prev(url string) *CollectionPageBuilder {
	cpb.collectionPage.Prev = &ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return cpb
}

func (cpb *CollectionPageBuilder) Items(urls ...string) *CollectionPageBuilder {
	if len(urls) > 0 {
		targets := make(ObjectOrLink, len(urls))
		for i, url := range urls {
			targets[i] = Link{Type: "Link", Href: url}
		}
		cpb.collectionPage.Items = &ObjectOrLinkOrString{Target: targets}
	}
	return cpb
}

func (cpb *CollectionPageBuilder) Build() *CollectionPage {
	return &cpb.collectionPage
}

// OrderedCollectionBuilder
type OrderedCollectionBuilder struct {
	orderedCollection OrderedCollection
}

func NewOrderedCollection() *OrderedCollectionBuilder {
	return &OrderedCollectionBuilder{
		orderedCollection: OrderedCollection{Object: Object{Type: "OrderedCollection"}},
	}
}

func (ocb *OrderedCollectionBuilder) ID(id string) *OrderedCollectionBuilder {
	ocb.orderedCollection.ID = id
	return ocb
}

func (ocb *OrderedCollectionBuilder) TotalItems(totalItems int) *OrderedCollectionBuilder {
	ocb.orderedCollection.TotalItems = totalItems
	return ocb
}

func (ocb *OrderedCollectionBuilder) First(url string) *OrderedCollectionBuilder {
	ocb.orderedCollection.First = &StringWithOrderedCollectionPage{URL: url}
	return ocb
}

func (ocb *OrderedCollectionBuilder) Last(url string) *OrderedCollectionBuilder {
	ocb.orderedCollection.Last = &StringWithOrderedCollectionPage{URL: url}
	return ocb
}

func (ocb *OrderedCollectionBuilder) Build() *OrderedCollection {
	return &ocb.orderedCollection
}

// OrderedCollectionPageBuilder
type OrderedCollectionPageBuilder struct {
	orderedCollectionPage OrderedCollectionPage
}

func NewOrderedCollectionPage() *OrderedCollectionPageBuilder {
	return &OrderedCollectionPageBuilder{
		orderedCollectionPage: OrderedCollectionPage{OrderedCollection: OrderedCollection{Object: Object{Type: "OrderedCollectionPage"}}},
	}
}

func (ocpb *OrderedCollectionPageBuilder) ID(id string) *OrderedCollectionPageBuilder {
	ocpb.orderedCollectionPage.ID = id
	return ocpb
}

func (ocpb *OrderedCollectionPageBuilder) PartOf(url string) *OrderedCollectionPageBuilder {
	ocpb.orderedCollectionPage.PartOf = &ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return ocpb
}

func (ocpb *OrderedCollectionPageBuilder) Next(url string) *OrderedCollectionPageBuilder {
	ocpb.orderedCollectionPage.Next = &ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return ocpb
}

func (ocpb *OrderedCollectionPageBuilder) Prev(url string) *OrderedCollectionPageBuilder {
	ocpb.orderedCollectionPage.Prev = &ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return ocpb
}

func (ocpb *OrderedCollectionPageBuilder) OrderedItems(urls ...string) *OrderedCollectionPageBuilder {
	if len(urls) > 0 {
		targets := make(ObjectOrLink, len(urls))
		for i, url := range urls {
			targets[i] = Link{Type: "Link", Href: url}
		}
		ocpb.orderedCollectionPage.OrderedItems = &ObjectOrLinkOrString{Target: targets}
	}
	return ocpb
}

func (ocpb *OrderedCollectionPageBuilder) StartIndex(index uint) *OrderedCollectionPageBuilder {
	ocpb.orderedCollectionPage.StartIndex = index
	return ocpb
}

func (ocpb *OrderedCollectionPageBuilder) Build() *OrderedCollectionPage {
	return &ocpb.orderedCollectionPage
}

// LocationBuilder
type LocationBuilder struct {
	location Location
}

func NewLocation() *LocationBuilder {
	return &LocationBuilder{
		location: Location{Object: Object{Type: "Location"}},
	}
}

func (lb *LocationBuilder) ID(id string) *LocationBuilder {
	lb.location.ID = id
	return lb
}

func (lb *LocationBuilder) Name(name string) *LocationBuilder {
	lb.location.Name = name
	return lb
}

func (lb *LocationBuilder) Latitude(lat float32) *LocationBuilder {
	lb.location.Latitude = lat
	return lb
}

func (lb *LocationBuilder) Longitude(lon float32) *LocationBuilder {
	lb.location.Longitude = lon
	return lb
}

func (lb *LocationBuilder) Build() *Location {
	return &lb.location
}

// ProfileBuilder
type ProfileBuilder struct {
	profile Profile
}

func NewProfile() *ProfileBuilder {
	return &ProfileBuilder{
		profile: Profile{Object: Object{Type: "Profile"}},
	}
}

func (pb *ProfileBuilder) ID(id string) *ProfileBuilder {
	pb.profile.ID = id
	return pb
}

func (pb *ProfileBuilder) Name(name string) *ProfileBuilder {
	pb.profile.Name = name
	return pb
}

func (pb *ProfileBuilder) Describes(url string) *ProfileBuilder {
	pb.profile.Describes = ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return pb
}

func (pb *ProfileBuilder) Build() *Profile {
	return &pb.profile
}

// TombstoneBuilder
type TombstoneBuilder struct {
	tombstone Tombstone
}

func NewTombstone() *TombstoneBuilder {
	return &TombstoneBuilder{
		tombstone: Tombstone{Object: Object{Type: "Tombstone"}},
	}
}

func (tb *TombstoneBuilder) ID(id string) *TombstoneBuilder {
	tb.tombstone.ID = id
	return tb
}

func (tb *TombstoneBuilder) Deleted(t *time.Time) *TombstoneBuilder {
	tb.tombstone.Deleted = t
	return tb
}

func (tb *TombstoneBuilder) FormerType(formerType string) *TombstoneBuilder {
	tb.tombstone.FormerType = formerType
	return tb
}

func (tb *TombstoneBuilder) Build() *Tombstone {
	return &tb.tombstone
}

// RelationshipBuilder
type RelationshipBuilder struct {
	relationship Relationship
}

func NewRelationship() *RelationshipBuilder {
	return &RelationshipBuilder{
		relationship: Relationship{Object: Object{Type: "Relationship"}},
	}
}

func (rb *RelationshipBuilder) ID(id string) *RelationshipBuilder {
	rb.relationship.ID = id
	return rb
}

func (rb *RelationshipBuilder) Relationship(rel string) *RelationshipBuilder {
	rb.relationship.Relationship = rel
	return rb
}

func (rb *RelationshipBuilder) Subject(url string) *RelationshipBuilder {
	rb.relationship.Subject = &ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return rb
}

func (rb *RelationshipBuilder) Object(url string) *RelationshipBuilder {
	rb.relationship.RelationshipObject = &ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return rb
}

func (rb *RelationshipBuilder) Build() *Relationship {
	return &rb.relationship
}

// QuestionBuilder
type QuestionBuilder struct {
	question Question
}

func NewQuestion() *QuestionBuilder {
	return &QuestionBuilder{
		question: Question{IntransitiveActivity: IntransitiveActivity{Object: Object{Type: "Question"}}},
	}
}

func (qb *QuestionBuilder) ID(id string) *QuestionBuilder {
	qb.question.ID = id
	return qb
}

func (qb *QuestionBuilder) Content(content string) *QuestionBuilder {
	qb.question.Content = content
	return qb
}

func (qb *QuestionBuilder) OneOf(urls ...string) *QuestionBuilder {
	if len(urls) > 0 {
		targets := make(ObjectOrLink, len(urls))
		for i, url := range urls {
			targets[i] = Link{Type: "Link", Href: url}
		}
		qb.question.OneOf = &ObjectOrLinkOrString{Target: targets}
	}
	return qb
}

func (qb *QuestionBuilder) AnyOf(urls ...string) *QuestionBuilder {
	if len(urls) > 0 {
		targets := make(ObjectOrLink, len(urls))
		for i, url := range urls {
			targets[i] = Link{Type: "Link", Href: url}
		}
		qb.question.AnyOf = &ObjectOrLinkOrString{Target: targets}
	}
	return qb
}

func (qb *QuestionBuilder) Closed(t *time.Time) *QuestionBuilder {
	qb.question.Closed = t
	return qb
}

func (qb *QuestionBuilder) Build() *Question {
	return &qb.question
}

// ApplicationBuilder
type ApplicationBuilder struct {
	application Application
}

func NewApplication() *ApplicationBuilder {
	return &ApplicationBuilder{
		application: Application{Object: Object{Type: "Application"}},
	}
}

func (ab *ApplicationBuilder) ID(id string) *ApplicationBuilder {
	ab.application.ID = id
	return ab
}

func (ab *ApplicationBuilder) Name(name string) *ApplicationBuilder {
	ab.application.Name = name
	return ab
}

func (ab *ApplicationBuilder) Username(username string) *ApplicationBuilder {
	ab.application.PreferredUsername = username
	return ab
}

func (ab *ApplicationBuilder) Inbox(url string) *ApplicationBuilder {
	ab.application.Inbox = &StringWithOrderedCollection{URL: url}
	return ab
}

func (ab *ApplicationBuilder) Outbox(url string) *ApplicationBuilder {
	ab.application.Outbox = &StringWithOrderedCollection{URL: url}
	return ab
}

func (ab *ApplicationBuilder) Build() *Application {
	return &ab.application
}

// GroupBuilder
type GroupBuilder struct {
	group Group
}

func NewGroup() *GroupBuilder {
	return &GroupBuilder{
		group: Group{Object: Object{Type: "Group"}},
	}
}

func (gb *GroupBuilder) ID(id string) *GroupBuilder {
	gb.group.ID = id
	return gb
}

func (gb *GroupBuilder) Name(name string) *GroupBuilder {
	gb.group.Name = name
	return gb
}

func (gb *GroupBuilder) Username(username string) *GroupBuilder {
	gb.group.PreferredUsername = username
	return gb
}

func (gb *GroupBuilder) Inbox(url string) *GroupBuilder {
	gb.group.Inbox = &StringWithOrderedCollection{URL: url}
	return gb
}

func (gb *GroupBuilder) Outbox(url string) *GroupBuilder {
	gb.group.Outbox = &StringWithOrderedCollection{URL: url}
	return gb
}

func (gb *GroupBuilder) Build() *Group {
	return &gb.group
}

// OrganizationBuilder
type OrganizationBuilder struct {
	organization Organization
}

func NewOrganization() *OrganizationBuilder {
	return &OrganizationBuilder{
		organization: Organization{Object: Object{Type: "Organization"}},
	}
}

func (ob *OrganizationBuilder) ID(id string) *OrganizationBuilder {
	ob.organization.ID = id
	return ob
}

func (ob *OrganizationBuilder) Name(name string) *OrganizationBuilder {
	ob.organization.Name = name
	return ob
}

func (ob *OrganizationBuilder) Username(username string) *OrganizationBuilder {
	ob.organization.PreferredUsername = username
	return ob
}

func (ob *OrganizationBuilder) Inbox(url string) *OrganizationBuilder {
	ob.organization.Inbox = &StringWithOrderedCollection{URL: url}
	return ob
}

func (ob *OrganizationBuilder) Outbox(url string) *OrganizationBuilder {
	ob.organization.Outbox = &StringWithOrderedCollection{URL: url}
	return ob
}

func (ob *OrganizationBuilder) Build() *Organization {
	return &ob.organization
}

// PersonBuilder
type PersonBuilder struct {
	person Person
}

func NewPerson() *PersonBuilder {
	return &PersonBuilder{
		person: Person{Object: Object{Type: "Person"}},
	}
}

func (pb *PersonBuilder) ID(id string) *PersonBuilder {
	pb.person.ID = id
	return pb
}

func (pb *PersonBuilder) Name(name string) *PersonBuilder {
	pb.person.Name = name
	return pb
}

func (pb *PersonBuilder) Username(username string) *PersonBuilder {
	pb.person.PreferredUsername = username
	return pb
}

func (pb *PersonBuilder) Summary(summary string) *PersonBuilder {
	pb.person.Summary = summary
	return pb
}

func (pb *PersonBuilder) Published(t *time.Time) *PersonBuilder {
	pb.person.Published = t
	return pb
}

func (pb *PersonBuilder) Inbox(url string) *PersonBuilder {
	pb.person.Inbox = &StringWithOrderedCollection{URL: url}
	return pb
}

func (pb *PersonBuilder) Outbox(url string) *PersonBuilder {
	pb.person.Outbox = &StringWithOrderedCollection{URL: url}
	return pb
}

func (pb *PersonBuilder) Followers(url string) *PersonBuilder {
	pb.person.Followers = &StringWithOrderedCollection{URL: url}
	return pb
}

func (pb *PersonBuilder) Following(url string) *PersonBuilder {
	pb.person.Following = &StringWithOrderedCollection{URL: url}
	return pb
}

func (pb *PersonBuilder) Build() *Person {
	return &pb.person
}

// ServiceBuilder
type ServiceBuilder struct {
	service Service
}

func NewService() *ServiceBuilder {
	return &ServiceBuilder{
		service: Service{Object: Object{Type: "Service"}},
	}
}

func (sb *ServiceBuilder) ID(id string) *ServiceBuilder {
	sb.service.ID = id
	return sb
}

func (sb *ServiceBuilder) Name(name string) *ServiceBuilder {
	sb.service.Name = name
	return sb
}

func (sb *ServiceBuilder) Username(username string) *ServiceBuilder {
	sb.service.PreferredUsername = username
	return sb
}

func (sb *ServiceBuilder) Inbox(url string) *ServiceBuilder {
	sb.service.Inbox = &StringWithOrderedCollection{URL: url}
	return sb
}

func (sb *ServiceBuilder) Outbox(url string) *ServiceBuilder {
	sb.service.Outbox = &StringWithOrderedCollection{URL: url}
	return sb
}

func (sb *ServiceBuilder) Build() *Service {
	return &sb.service
}

// AcceptBuilder
type AcceptBuilder struct {
	accept Accept
}

func NewAccept() *AcceptBuilder {
	return &AcceptBuilder{
		accept: Accept{Object: Object{Type: "Accept"}},
	}
}

func (ab *AcceptBuilder) ID(id string) *AcceptBuilder {
	ab.accept.ID = id
	return ab
}

func (ab *AcceptBuilder) Actor(url string) *AcceptBuilder {
	ab.accept.Actor = &ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return ab
}

func (ab *AcceptBuilder) Object(obj interface{}) *AcceptBuilder {
	switch v := obj.(type) {
	case *Note:
		ab.accept.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	case *Person:
		ab.accept.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	default:
		ab.accept.ActivityObject = &ObjectOrLinkOrString{
			URL: []string{obj.(string)},
		}
	}
	return ab
}

func (ab *AcceptBuilder) Published(t *time.Time) *AcceptBuilder {
	ab.accept.Published = t
	return ab
}

func (ab *AcceptBuilder) Build() *Accept {
	return &ab.accept
}

// AddBuilder
type AddBuilder struct {
	add Add
}

func NewAdd() *AddBuilder {
	return &AddBuilder{
		add: Add{Object: Object{Type: "Add"}},
	}
}

func (ab *AddBuilder) ID(id string) *AddBuilder {
	ab.add.ID = id
	return ab
}

func (ab *AddBuilder) Actor(url string) *AddBuilder {
	ab.add.Actor = &ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return ab
}

func (ab *AddBuilder) Object(obj interface{}) *AddBuilder {
	switch v := obj.(type) {
	case *Note:
		ab.add.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	case *Person:
		ab.add.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	default:
		ab.add.ActivityObject = &ObjectOrLinkOrString{
			URL: []string{obj.(string)},
		}
	}
	return ab
}

func (ab *AddBuilder) Published(t *time.Time) *AddBuilder {
	ab.add.Published = t
	return ab
}

func (ab *AddBuilder) Build() *Add {
	return &ab.add
}

// AnnounceBuilder
type AnnounceBuilder struct {
	announce Announce
}

func NewAnnounce() *AnnounceBuilder {
	return &AnnounceBuilder{
		announce: Announce{Object: Object{Type: "Announce"}},
	}
}

func (ab *AnnounceBuilder) ID(id string) *AnnounceBuilder {
	ab.announce.ID = id
	return ab
}

func (ab *AnnounceBuilder) Actor(url string) *AnnounceBuilder {
	ab.announce.Actor = &ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return ab
}

func (ab *AnnounceBuilder) Object(obj interface{}) *AnnounceBuilder {
	switch v := obj.(type) {
	case *Note:
		ab.announce.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	case *Person:
		ab.announce.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	default:
		ab.announce.ActivityObject = &ObjectOrLinkOrString{
			URL: []string{obj.(string)},
		}
	}
	return ab
}

func (ab *AnnounceBuilder) Published(t *time.Time) *AnnounceBuilder {
	ab.announce.Published = t
	return ab
}

func (ab *AnnounceBuilder) Build() *Announce {
	return &ab.announce
}

// ArriveBuilder
type ArriveBuilder struct {
	arrive Arrive
}

func NewArrive() *ArriveBuilder {
	return &ArriveBuilder{
		arrive: Arrive{Object: Object{Type: "Arrive"}},
	}
}

func (ab *ArriveBuilder) ID(id string) *ArriveBuilder {
	ab.arrive.ID = id
	return ab
}

func (ab *ArriveBuilder) Actor(url string) *ArriveBuilder {
	ab.arrive.Actor = &ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return ab
}

func (ab *ArriveBuilder) Published(t *time.Time) *ArriveBuilder {
	ab.arrive.Published = t
	return ab
}

func (ab *ArriveBuilder) Build() *Arrive {
	return &ab.arrive
}

// BlockBuilder
type BlockBuilder struct {
	block Block
}

func NewBlock() *BlockBuilder {
	return &BlockBuilder{
		block: Block{Object: Object{Type: "Block"}},
	}
}

func (bb *BlockBuilder) ID(id string) *BlockBuilder {
	bb.block.ID = id
	return bb
}

func (bb *BlockBuilder) Actor(url string) *BlockBuilder {
	bb.block.Actor = &ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return bb
}

func (bb *BlockBuilder) Object(obj interface{}) *BlockBuilder {
	switch v := obj.(type) {
	case *Note:
		bb.block.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	case *Person:
		bb.block.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	default:
		bb.block.ActivityObject = &ObjectOrLinkOrString{
			URL: []string{obj.(string)},
		}
	}
	return bb
}

func (bb *BlockBuilder) Published(t *time.Time) *BlockBuilder {
	bb.block.Published = t
	return bb
}

func (bb *BlockBuilder) Build() *Block {
	return &bb.block
}

// CreateBuilder
type CreateBuilder struct {
	create Create
}

func NewCreate() *CreateBuilder {
	return &CreateBuilder{
		create: Create{Object: Object{Type: "Create"}},
	}
}

func (cb *CreateBuilder) ID(id string) *CreateBuilder {
	cb.create.ID = id
	return cb
}

func (cb *CreateBuilder) Actor(url string) *CreateBuilder {
	cb.create.Actor = &ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return cb
}

func (cb *CreateBuilder) Object(obj interface{}) *CreateBuilder {
	switch v := obj.(type) {
	case *Note:
		cb.create.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	case *Person:
		cb.create.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	default:
		cb.create.ActivityObject = &ObjectOrLinkOrString{
			URL: []string{obj.(string)},
		}
	}
	return cb
}

func (cb *CreateBuilder) Published(t *time.Time) *CreateBuilder {
	cb.create.Published = t
	return cb
}

func (cb *CreateBuilder) Build() *Create {
	return &cb.create
}

// DeleteBuilder
type DeleteBuilder struct {
	delete Delete
}

func NewDelete() *DeleteBuilder {
	return &DeleteBuilder{
		delete: Delete{Object: Object{Type: "Delete"}},
	}
}

func (db *DeleteBuilder) ID(id string) *DeleteBuilder {
	db.delete.ID = id
	return db
}

func (db *DeleteBuilder) Actor(url string) *DeleteBuilder {
	db.delete.Actor = &ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return db
}

func (db *DeleteBuilder) Object(obj interface{}) *DeleteBuilder {
	switch v := obj.(type) {
	case *Note:
		db.delete.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	case *Person:
		db.delete.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	default:
		db.delete.ActivityObject = &ObjectOrLinkOrString{
			URL: []string{obj.(string)},
		}
	}
	return db
}

func (db *DeleteBuilder) Published(t *time.Time) *DeleteBuilder {
	db.delete.Published = t
	return db
}

func (db *DeleteBuilder) Build() *Delete {
	return &db.delete
}

// DislikeBuilder
type DislikeBuilder struct {
	dislike Dislike
}

func NewDislike() *DislikeBuilder {
	return &DislikeBuilder{
		dislike: Dislike{Object: Object{Type: "Dislike"}},
	}
}

func (db *DislikeBuilder) ID(id string) *DislikeBuilder {
	db.dislike.ID = id
	return db
}

func (db *DislikeBuilder) Actor(url string) *DislikeBuilder {
	db.dislike.Actor = &ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return db
}

func (db *DislikeBuilder) Object(obj interface{}) *DislikeBuilder {
	switch v := obj.(type) {
	case *Note:
		db.dislike.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	case *Person:
		db.dislike.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	default:
		db.dislike.ActivityObject = &ObjectOrLinkOrString{
			URL: []string{obj.(string)},
		}
	}
	return db
}

func (db *DislikeBuilder) Published(t *time.Time) *DislikeBuilder {
	db.dislike.Published = t
	return db
}

func (db *DislikeBuilder) Build() *Dislike {
	return &db.dislike
}

// FollowBuilder
type FollowBuilder struct {
	follow Follow
}

func NewFollow() *FollowBuilder {
	return &FollowBuilder{
		follow: Follow{Object: Object{Type: "Follow"}},
	}
}

func (fb *FollowBuilder) ID(id string) *FollowBuilder {
	fb.follow.ID = id
	return fb
}

func (fb *FollowBuilder) Actor(url string) *FollowBuilder {
	fb.follow.Actor = &ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return fb
}

func (fb *FollowBuilder) Object(url string) *FollowBuilder {
	fb.follow.ActivityObject = &ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return fb
}

func (fb *FollowBuilder) Published(t *time.Time) *FollowBuilder {
	fb.follow.Published = t
	return fb
}

func (fb *FollowBuilder) Build() *Follow {
	return &fb.follow
}

// FlagBuilder
type FlagBuilder struct {
	flag Flag
}

func NewFlag() *FlagBuilder {
	return &FlagBuilder{
		flag: Flag{Object: Object{Type: "Flag"}},
	}
}

func (fb *FlagBuilder) ID(id string) *FlagBuilder {
	fb.flag.ID = id
	return fb
}

func (fb *FlagBuilder) Actor(url string) *FlagBuilder {
	fb.flag.Actor = &ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return fb
}

func (fb *FlagBuilder) Object(obj interface{}) *FlagBuilder {
	switch v := obj.(type) {
	case *Note:
		fb.flag.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	case *Person:
		fb.flag.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	default:
		fb.flag.ActivityObject = &ObjectOrLinkOrString{
			URL: []string{obj.(string)},
		}
	}
	return fb
}

func (fb *FlagBuilder) Published(t *time.Time) *FlagBuilder {
	fb.flag.Published = t
	return fb
}

func (fb *FlagBuilder) Build() *Flag {
	return &fb.flag
}

// IgnoreBuilder
type IgnoreBuilder struct {
	ignore Ignore
}

func NewIgnore() *IgnoreBuilder {
	return &IgnoreBuilder{
		ignore: Ignore{Object: Object{Type: "Ignore"}},
	}
}

func (ib *IgnoreBuilder) ID(id string) *IgnoreBuilder {
	ib.ignore.ID = id
	return ib
}

func (ib *IgnoreBuilder) Actor(url string) *IgnoreBuilder {
	ib.ignore.Actor = &ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return ib
}

func (ib *IgnoreBuilder) Object(obj interface{}) *IgnoreBuilder {
	switch v := obj.(type) {
	case *Note:
		ib.ignore.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	case *Person:
		ib.ignore.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	default:
		ib.ignore.ActivityObject = &ObjectOrLinkOrString{
			URL: []string{obj.(string)},
		}
	}
	return ib
}

func (ib *IgnoreBuilder) Published(t *time.Time) *IgnoreBuilder {
	ib.ignore.Published = t
	return ib
}

func (ib *IgnoreBuilder) Build() *Ignore {
	return &ib.ignore
}

// InviteBuilder
type InviteBuilder struct {
	invite Invite
}

func NewInvite() *InviteBuilder {
	return &InviteBuilder{
		invite: Invite{Object: Object{Type: "Invite"}},
	}
}

func (ib *InviteBuilder) ID(id string) *InviteBuilder {
	ib.invite.ID = id
	return ib
}

func (ib *InviteBuilder) Actor(url string) *InviteBuilder {
	ib.invite.Actor = &ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return ib
}

func (ib *InviteBuilder) Object(obj interface{}) *InviteBuilder {
	switch v := obj.(type) {
	case *Note:
		ib.invite.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	case *Person:
		ib.invite.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	default:
		ib.invite.ActivityObject = &ObjectOrLinkOrString{
			URL: []string{obj.(string)},
		}
	}
	return ib
}

func (ib *InviteBuilder) Published(t *time.Time) *InviteBuilder {
	ib.invite.Published = t
	return ib
}

func (ib *InviteBuilder) Build() *Invite {
	return &ib.invite
}

// JoinBuilder
type JoinBuilder struct {
	join Join
}

func NewJoin() *JoinBuilder {
	return &JoinBuilder{
		join: Join{Object: Object{Type: "Join"}},
	}
}

func (jb *JoinBuilder) ID(id string) *JoinBuilder {
	jb.join.ID = id
	return jb
}

func (jb *JoinBuilder) Actor(url string) *JoinBuilder {
	jb.join.Actor = &ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return jb
}

func (jb *JoinBuilder) Object(obj interface{}) *JoinBuilder {
	switch v := obj.(type) {
	case *Note:
		jb.join.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	case *Person:
		jb.join.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	default:
		jb.join.ActivityObject = &ObjectOrLinkOrString{
			URL: []string{obj.(string)},
		}
	}
	return jb
}

func (jb *JoinBuilder) Published(t *time.Time) *JoinBuilder {
	jb.join.Published = t
	return jb
}

func (jb *JoinBuilder) Build() *Join {
	return &jb.join
}

// LeaveBuilder
type LeaveBuilder struct {
	leave Leave
}

func NewLeave() *LeaveBuilder {
	return &LeaveBuilder{
		leave: Leave{Object: Object{Type: "Leave"}},
	}
}

func (lb *LeaveBuilder) ID(id string) *LeaveBuilder {
	lb.leave.ID = id
	return lb
}

func (lb *LeaveBuilder) Actor(url string) *LeaveBuilder {
	lb.leave.Actor = &ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return lb
}

func (lb *LeaveBuilder) Object(obj interface{}) *LeaveBuilder {
	switch v := obj.(type) {
	case *Note:
		lb.leave.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	case *Person:
		lb.leave.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	default:
		lb.leave.ActivityObject = &ObjectOrLinkOrString{
			URL: []string{obj.(string)},
		}
	}
	return lb
}

func (lb *LeaveBuilder) Published(t *time.Time) *LeaveBuilder {
	lb.leave.Published = t
	return lb
}

func (lb *LeaveBuilder) Build() *Leave {
	return &lb.leave
}

// LikeBuilder
type LikeBuilder struct {
	like Like
}

func NewLike() *LikeBuilder {
	return &LikeBuilder{
		like: Like{Object: Object{Type: "Like"}},
	}
}

func (lb *LikeBuilder) ID(id string) *LikeBuilder {
	lb.like.ID = id
	return lb
}

func (lb *LikeBuilder) Actor(url string) *LikeBuilder {
	lb.like.Actor = &ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return lb
}

func (lb *LikeBuilder) Object(obj interface{}) *LikeBuilder {
	switch v := obj.(type) {
	case *Note:
		lb.like.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	case *Person:
		lb.like.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	default:
		lb.like.ActivityObject = &ObjectOrLinkOrString{
			URL: []string{obj.(string)},
		}
	}
	return lb
}

func (lb *LikeBuilder) Published(t *time.Time) *LikeBuilder {
	lb.like.Published = t
	return lb
}

func (lb *LikeBuilder) Build() *Like {
	return &lb.like
}

// ListenBuilder
type ListenBuilder struct {
	listen Listen
}

func NewListen() *ListenBuilder {
	return &ListenBuilder{
		listen: Listen{Object: Object{Type: "Listen"}},
	}
}

func (lb *ListenBuilder) ID(id string) *ListenBuilder {
	lb.listen.ID = id
	return lb
}

func (lb *ListenBuilder) Actor(url string) *ListenBuilder {
	lb.listen.Actor = &ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return lb
}

func (lb *ListenBuilder) Object(obj interface{}) *ListenBuilder {
	switch v := obj.(type) {
	case *Note:
		lb.listen.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	case *Person:
		lb.listen.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	default:
		lb.listen.ActivityObject = &ObjectOrLinkOrString{
			URL: []string{obj.(string)},
		}
	}
	return lb
}

func (lb *ListenBuilder) Published(t *time.Time) *ListenBuilder {
	lb.listen.Published = t
	return lb
}

func (lb *ListenBuilder) Build() *Listen {
	return &lb.listen
}

// MoveBuilder
type MoveBuilder struct {
	move Move
}

func NewMove() *MoveBuilder {
	return &MoveBuilder{
		move: Move{Object: Object{Type: "Move"}},
	}
}

func (mb *MoveBuilder) ID(id string) *MoveBuilder {
	mb.move.ID = id
	return mb
}

func (mb *MoveBuilder) Actor(url string) *MoveBuilder {
	mb.move.Actor = &ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return mb
}

func (mb *MoveBuilder) Object(obj interface{}) *MoveBuilder {
	switch v := obj.(type) {
	case *Note:
		mb.move.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	case *Person:
		mb.move.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	default:
		mb.move.ActivityObject = &ObjectOrLinkOrString{
			URL: []string{obj.(string)},
		}
	}
	return mb
}

func (mb *MoveBuilder) Published(t *time.Time) *MoveBuilder {
	mb.move.Published = t
	return mb
}

func (mb *MoveBuilder) Build() *Move {
	return &mb.move
}

// OfferBuilder
type OfferBuilder struct {
	offer Offer
}

func NewOffer() *OfferBuilder {
	return &OfferBuilder{
		offer: Offer{Object: Object{Type: "Offer"}},
	}
}

func (ob *OfferBuilder) ID(id string) *OfferBuilder {
	ob.offer.ID = id
	return ob
}

func (ob *OfferBuilder) Actor(url string) *OfferBuilder {
	ob.offer.Actor = &ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return ob
}

func (ob *OfferBuilder) Object(obj interface{}) *OfferBuilder {
	switch v := obj.(type) {
	case *Note:
		ob.offer.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	case *Person:
		ob.offer.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	default:
		ob.offer.ActivityObject = &ObjectOrLinkOrString{
			URL: []string{obj.(string)},
		}
	}
	return ob
}

func (ob *OfferBuilder) Published(t *time.Time) *OfferBuilder {
	ob.offer.Published = t
	return ob
}

func (ob *OfferBuilder) Build() *Offer {
	return &ob.offer
}

// ReadBuilder
type ReadBuilder struct {
	read Read
}

func NewRead() *ReadBuilder {
	return &ReadBuilder{
		read: Read{Object: Object{Type: "Read"}},
	}
}

func (rb *ReadBuilder) ID(id string) *ReadBuilder {
	rb.read.ID = id
	return rb
}

func (rb *ReadBuilder) Actor(url string) *ReadBuilder {
	rb.read.Actor = &ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return rb
}

func (rb *ReadBuilder) Object(obj interface{}) *ReadBuilder {
	switch v := obj.(type) {
	case *Note:
		rb.read.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	case *Person:
		rb.read.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	default:
		rb.read.ActivityObject = &ObjectOrLinkOrString{
			URL: []string{obj.(string)},
		}
	}
	return rb
}

func (rb *ReadBuilder) Published(t *time.Time) *ReadBuilder {
	rb.read.Published = t
	return rb
}

func (rb *ReadBuilder) Build() *Read {
	return &rb.read
}

// RejectBuilder
type RejectBuilder struct {
	reject Reject
}

func NewReject() *RejectBuilder {
	return &RejectBuilder{
		reject: Reject{Object: Object{Type: "Reject"}},
	}
}

func (rb *RejectBuilder) ID(id string) *RejectBuilder {
	rb.reject.ID = id
	return rb
}

func (rb *RejectBuilder) Actor(url string) *RejectBuilder {
	rb.reject.Actor = &ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return rb
}

func (rb *RejectBuilder) Object(obj interface{}) *RejectBuilder {
	switch v := obj.(type) {
	case *Note:
		rb.reject.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	case *Person:
		rb.reject.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	default:
		rb.reject.ActivityObject = &ObjectOrLinkOrString{
			URL: []string{obj.(string)},
		}
	}
	return rb
}

func (rb *RejectBuilder) Published(t *time.Time) *RejectBuilder {
	rb.reject.Published = t
	return rb
}

func (rb *RejectBuilder) Build() *Reject {
	return &rb.reject
}

// RemoveBuilder
type RemoveBuilder struct {
	remove Remove
}

func NewRemove() *RemoveBuilder {
	return &RemoveBuilder{
		remove: Remove{Object: Object{Type: "Remove"}},
	}
}

func (rb *RemoveBuilder) ID(id string) *RemoveBuilder {
	rb.remove.ID = id
	return rb
}

func (rb *RemoveBuilder) Actor(url string) *RemoveBuilder {
	rb.remove.Actor = &ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return rb
}

func (rb *RemoveBuilder) Object(obj interface{}) *RemoveBuilder {
	switch v := obj.(type) {
	case *Note:
		rb.remove.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	case *Person:
		rb.remove.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	default:
		rb.remove.ActivityObject = &ObjectOrLinkOrString{
			URL: []string{obj.(string)},
		}
	}
	return rb
}

func (rb *RemoveBuilder) Published(t *time.Time) *RemoveBuilder {
	rb.remove.Published = t
	return rb
}

func (rb *RemoveBuilder) Build() *Remove {
	return &rb.remove
}

// TentativeAcceptBuilder
type TentativeAcceptBuilder struct {
	tentativeAccept TentativeAccept
}

func NewTentativeAccept() *TentativeAcceptBuilder {
	return &TentativeAcceptBuilder{
		tentativeAccept: TentativeAccept{Object: Object{Type: "TentativeAccept"}},
	}
}

func (tab *TentativeAcceptBuilder) ID(id string) *TentativeAcceptBuilder {
	tab.tentativeAccept.ID = id
	return tab
}

func (tab *TentativeAcceptBuilder) Actor(url string) *TentativeAcceptBuilder {
	tab.tentativeAccept.Actor = &ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return tab
}

func (tab *TentativeAcceptBuilder) Object(obj interface{}) *TentativeAcceptBuilder {
	switch v := obj.(type) {
	case *Note:
		tab.tentativeAccept.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	case *Person:
		tab.tentativeAccept.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	default:
		tab.tentativeAccept.ActivityObject = &ObjectOrLinkOrString{
			URL: []string{obj.(string)},
		}
	}
	return tab
}

func (tab *TentativeAcceptBuilder) Published(t *time.Time) *TentativeAcceptBuilder {
	tab.tentativeAccept.Published = t
	return tab
}

func (tab *TentativeAcceptBuilder) Build() *TentativeAccept {
	return &tab.tentativeAccept
}

// TentativeRejectBuilder
type TentativeRejectBuilder struct {
	tentativeReject TentativeReject
}

func NewTentativeReject() *TentativeRejectBuilder {
	return &TentativeRejectBuilder{
		tentativeReject: TentativeReject{Object: Object{Type: "TentativeReject"}},
	}
}

func (trb *TentativeRejectBuilder) ID(id string) *TentativeRejectBuilder {
	trb.tentativeReject.ID = id
	return trb
}

func (trb *TentativeRejectBuilder) Actor(url string) *TentativeRejectBuilder {
	trb.tentativeReject.Actor = &ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return trb
}

func (trb *TentativeRejectBuilder) Object(obj interface{}) *TentativeRejectBuilder {
	switch v := obj.(type) {
	case *Note:
		trb.tentativeReject.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	case *Person:
		trb.tentativeReject.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	default:
		trb.tentativeReject.ActivityObject = &ObjectOrLinkOrString{
			URL: []string{obj.(string)},
		}
	}
	return trb
}

func (trb *TentativeRejectBuilder) Published(t *time.Time) *TentativeRejectBuilder {
	trb.tentativeReject.Published = t
	return trb
}

func (trb *TentativeRejectBuilder) Build() *TentativeReject {
	return &trb.tentativeReject
}

// TravelBuilder
type TravelBuilder struct {
	travel Travel
}

func NewTravel() *TravelBuilder {
	return &TravelBuilder{
		travel: Travel{Object: Object{Type: "Travel"}},
	}
}

func (tb *TravelBuilder) ID(id string) *TravelBuilder {
	tb.travel.ID = id
	return tb
}

func (tb *TravelBuilder) Actor(url string) *TravelBuilder {
	tb.travel.Actor = &ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return tb
}

func (tb *TravelBuilder) Published(t *time.Time) *TravelBuilder {
	tb.travel.Published = t
	return tb
}

func (tb *TravelBuilder) Build() *Travel {
	return &tb.travel
}

// UndoBuilder
type UndoBuilder struct {
	undo Undo
}

func NewUndo() *UndoBuilder {
	return &UndoBuilder{
		undo: Undo{Object: Object{Type: "Undo"}},
	}
}

func (ub *UndoBuilder) ID(id string) *UndoBuilder {
	ub.undo.ID = id
	return ub
}

func (ub *UndoBuilder) Actor(url string) *UndoBuilder {
	ub.undo.Actor = &ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return ub
}

func (ub *UndoBuilder) Object(obj interface{}) *UndoBuilder {
	switch v := obj.(type) {
	case *Note:
		ub.undo.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	case *Person:
		ub.undo.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	default:
		ub.undo.ActivityObject = &ObjectOrLinkOrString{
			URL: []string{obj.(string)},
		}
	}
	return ub
}

func (ub *UndoBuilder) Published(t *time.Time) *UndoBuilder {
	ub.undo.Published = t
	return ub
}

func (ub *UndoBuilder) Build() *Undo {
	return &ub.undo
}

// UpdateBuilder
type UpdateBuilder struct {
	update Update
}

func NewUpdate() *UpdateBuilder {
	return &UpdateBuilder{
		update: Update{Object: Object{Type: "Update"}},
	}
}

func (ub *UpdateBuilder) ID(id string) *UpdateBuilder {
	ub.update.ID = id
	return ub
}

func (ub *UpdateBuilder) Actor(url string) *UpdateBuilder {
	ub.update.Actor = &ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return ub
}

func (ub *UpdateBuilder) Object(obj interface{}) *UpdateBuilder {
	switch v := obj.(type) {
	case *Note:
		ub.update.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	case *Person:
		ub.update.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	default:
		ub.update.ActivityObject = &ObjectOrLinkOrString{
			URL: []string{obj.(string)},
		}
	}
	return ub
}

func (ub *UpdateBuilder) Published(t *time.Time) *UpdateBuilder {
	ub.update.Published = t
	return ub
}

func (ub *UpdateBuilder) Build() *Update {
	return &ub.update
}

// ViewBuilder
type ViewBuilder struct {
	view View
}

func NewView() *ViewBuilder {
	return &ViewBuilder{
		view: View{Object: Object{Type: "View"}},
	}
}

func (vb *ViewBuilder) ID(id string) *ViewBuilder {
	vb.view.ID = id
	return vb
}

func (vb *ViewBuilder) Actor(url string) *ViewBuilder {
	vb.view.Actor = &ObjectOrLinkOrString{
		Target: ObjectOrLink{
			Link{Type: "Link", Href: url},
		},
	}
	return vb
}

func (vb *ViewBuilder) Object(obj interface{}) *ViewBuilder {
	switch v := obj.(type) {
	case *Note:
		vb.view.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	case *Person:
		vb.view.ActivityObject = &ObjectOrLinkOrString{
			Target: ObjectOrLink{*v},
		}
	default:
		vb.view.ActivityObject = &ObjectOrLinkOrString{
			URL: []string{obj.(string)},
		}
	}
	return vb
}

func (vb *ViewBuilder) Published(t *time.Time) *ViewBuilder {
	vb.view.Published = t
	return vb
}

func (vb *ViewBuilder) Build() *View {
	return &vb.view
}
