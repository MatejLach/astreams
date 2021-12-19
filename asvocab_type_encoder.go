package astreams

import (
	"encoding/json"
	"errors"
	"fmt"
)

func encodeASType(t ObjectLinker) ([]byte, error) {
	datatType, _ := ConcreteType(t)
	switch datatType {
	case "Activity":
		if activity, ok := t.(Activity); ok {
			marshalB, err := json.Marshal(activity)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Application":
		if application, ok := t.(Application); ok {
			marshalB, err := json.Marshal(application)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Article":
		if article, ok := t.(Article); ok {
			marshalB, err := json.Marshal(article)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Audio":
		if audio, ok := t.(Audio); ok {
			marshalB, err := json.Marshal(audio)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Collection":
		if collection, ok := t.(Collection); ok {
			marshalB, err := json.Marshal(collection)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "CollectionPage":
		if colPage, ok := t.(CollectionPage); ok {
			marshalB, err := json.Marshal(colPage)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Relationship":
		if relationship, ok := t.(Relationship); ok {
			marshalB, err := json.Marshal(relationship)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Document":
		if document, ok := t.(Document); ok {
			marshalB, err := json.Marshal(document)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Event":
		if event, ok := t.(Event); ok {
			marshalB, err := json.Marshal(event)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Group":
		if group, ok := t.(Group); ok {
			marshalB, err := json.Marshal(group)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Image":
		if image, ok := t.(Image); ok {
			marshalB, err := json.Marshal(image)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "IntransitiveActivity":
		if intransitiveActivity, ok := t.(IntransitiveActivity); ok {
			marshalB, err := json.Marshal(intransitiveActivity)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Note":
		if note, ok := t.(Note); ok {
			marshalB, err := json.Marshal(note)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Object":
		if object, ok := t.(Object); ok {
			marshalB, err := json.Marshal(object)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "OrderedCollection":
		if ordCollection, ok := t.(OrderedCollection); ok {
			marshalB, err := json.Marshal(ordCollection)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "OrderedCollectionPage":
		if ordColPage, ok := t.(OrderedCollectionPage); ok {
			marshalB, err := json.Marshal(ordColPage)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Organization":
		if organization, ok := t.(Organization); ok {
			marshalB, err := json.Marshal(organization)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Page":
		if page, ok := t.(Page); ok {
			marshalB, err := json.Marshal(page)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Person":
		if person, ok := t.(Person); ok {
			marshalB, err := json.Marshal(person)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Place":
		if place, ok := t.(Place); ok {
			marshalB, err := json.Marshal(place)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Profile":
		if profile, ok := t.(Profile); ok {
			marshalB, err := json.Marshal(profile)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Question":
		if question, ok := t.(Question); ok {
			marshalB, err := json.Marshal(question)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Service":
		if service, ok := t.(Service); ok {
			marshalB, err := json.Marshal(service)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Tombstone":
		if tombstone, ok := t.(Tombstone); ok {
			marshalB, err := json.Marshal(tombstone)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Video":
		if video, ok := t.(Video); ok {
			marshalB, err := json.Marshal(video)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Accept":
		if accept, ok := t.(Accept); ok {
			marshalB, err := json.Marshal(accept)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Add":
		if add, ok := t.(Add); ok {
			marshalB, err := json.Marshal(add)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Announce":
		if announce, ok := t.(Announce); ok {
			marshalB, err := json.Marshal(announce)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Arrive":
		if arrive, ok := t.(Arrive); ok {
			marshalB, err := json.Marshal(arrive)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Block":
		if block, ok := t.(Block); ok {
			marshalB, err := json.Marshal(block)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Create":
		if create, ok := t.(Create); ok {
			marshalB, err := json.Marshal(create)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Delete":
		if del, ok := t.(Delete); ok {
			marshalB, err := json.Marshal(del)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Dislike":
		if dislike, ok := t.(Dislike); ok {
			marshalB, err := json.Marshal(dislike)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Follow":
		if follow, ok := t.(Follow); ok {
			marshalB, err := json.Marshal(follow)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Flag":
		if flag, ok := t.(Flag); ok {
			marshalB, err := json.Marshal(flag)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Ignore":
		if ignore, ok := t.(Ignore); ok {
			marshalB, err := json.Marshal(ignore)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Invite":
		if invite, ok := t.(Invite); ok {
			marshalB, err := json.Marshal(invite)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Join":
		if join, ok := t.(Join); ok {
			marshalB, err := json.Marshal(join)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Leave":
		if leave, ok := t.(Leave); ok {
			marshalB, err := json.Marshal(leave)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Like":
		if like, ok := t.(Like); ok {
			marshalB, err := json.Marshal(like)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Listen":
		if listen, ok := t.(Listen); ok {
			marshalB, err := json.Marshal(listen)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Move":
		if move, ok := t.(Move); ok {
			marshalB, err := json.Marshal(move)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Offer":
		if offer, ok := t.(Offer); ok {
			marshalB, err := json.Marshal(offer)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Read":
		if read, ok := t.(Read); ok {
			marshalB, err := json.Marshal(read)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Reject":
		if reject, ok := t.(Reject); ok {
			marshalB, err := json.Marshal(reject)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Remove":
		if remove, ok := t.(Remove); ok {
			marshalB, err := json.Marshal(remove)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "TentativeAccept":
		if tentativeAccept, ok := t.(TentativeAccept); ok {
			marshalB, err := json.Marshal(tentativeAccept)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "TentativeReject":
		if tentativeReject, ok := t.(TentativeReject); ok {
			marshalB, err := json.Marshal(tentativeReject)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Travel":
		if travel, ok := t.(Travel); ok {
			marshalB, err := json.Marshal(travel)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Undo":
		if undo, ok := t.(Undo); ok {
			marshalB, err := json.Marshal(undo)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Update":
		if update, ok := t.(Update); ok {
			marshalB, err := json.Marshal(update)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "View":
		if view, ok := t.(View); ok {
			marshalB, err := json.Marshal(view)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Link":
		if link, ok := t.(Link); ok {
			marshalB, err := json.Marshal(link)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Mention":
		if mention, ok := t.(Mention); ok {
			marshalB, err := json.Marshal(mention)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "Hashtag":
		if hashtag, ok := t.(Hashtag); ok {
			marshalB, err := json.Marshal(hashtag)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	case "PropertyValue":
		if propertyValue, ok := t.(Mention); ok {
			marshalB, err := json.Marshal(propertyValue)
			if err != nil {
				return []byte{}, err
			}
			return marshalB, nil
		}
		return []byte{}, fmt.Errorf("failed to Marshal %s", datatType)
	}
	return []byte{}, errors.New("unsupported ActivityStreams type, or invalid AS document")
}
