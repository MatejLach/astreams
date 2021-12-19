package astreams

import (
	"bytes"
	"encoding/json"
	"errors"
)

// Decode the supplied JSON into the appropriate type based on its "type" property
func decodeASType(data []byte, jsonType string) (ObjectLinker, error) {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	switch jsonType {
	case "Activity":
		activity := Activity{}
		if err := decoder.Decode(&activity); err != nil {
			return nil, err
		}
		return ObjectLinker(activity), nil
	case "Application":
		application := Application{}
		if err := decoder.Decode(&application); err != nil {
			return nil, err
		}
		return ObjectLinker(application), nil
	case "Article":
		article := Article{}
		if err := decoder.Decode(&article); err != nil {
			return nil, err
		}
		return ObjectLinker(article), nil
	case "Audio":
		audio := Audio{}
		if err := decoder.Decode(&audio); err != nil {
			return nil, err
		}
		return ObjectLinker(audio), nil
	case "Collection":
		collection := Collection{}
		if err := decoder.Decode(&collection); err != nil {
			return nil, err
		}
		return ObjectLinker(collection), nil
	case "CollectionPage":
		colPage := CollectionPage{}
		if err := decoder.Decode(&colPage); err != nil {
			return nil, err
		}
		return ObjectLinker(colPage), nil
	case "Relationship":
		relationship := Relationship{}
		if err := decoder.Decode(&relationship); err != nil {
			return nil, err
		}
		return ObjectLinker(relationship), nil
	case "Document":
		document := Document{}
		if err := decoder.Decode(&document); err != nil {
			return nil, err
		}
		return ObjectLinker(document), nil
	case "Event":
		event := Event{}
		if err := decoder.Decode(&event); err != nil {
			return nil, err
		}
		return ObjectLinker(event), nil
	case "Group":
		group := Group{}
		if err := decoder.Decode(&group); err != nil {
			return nil, err
		}
		return ObjectLinker(group), nil
	case "Image":
		image := Image{}
		if err := decoder.Decode(&image); err != nil {
			return nil, err
		}
		return ObjectLinker(image), nil
	case "IntransitiveActivity":
		intransitiveActivity := IntransitiveActivity{}
		if err := decoder.Decode(&intransitiveActivity); err != nil {
			return nil, err
		}
		return ObjectLinker(intransitiveActivity), nil
	case "Note":
		note := Note{}
		if err := decoder.Decode(&note); err != nil {
			return nil, err
		}
		return ObjectLinker(note), nil
	case "Object":
		object := Object{}
		if err := decoder.Decode(&object); err != nil {
			return nil, err
		}
		return ObjectLinker(object), nil
	case "OrderedCollection":
		ordCollection := OrderedCollection{}
		if err := decoder.Decode(&ordCollection); err != nil {
			return nil, err
		}
		return ObjectLinker(ordCollection), nil
	case "OrderedCollectionPage":
		ordColPage := OrderedCollectionPage{}
		if err := decoder.Decode(&ordColPage); err != nil {
			return nil, err
		}
		return ObjectLinker(ordColPage), nil
	case "Organization":
		organization := Organization{}
		if err := decoder.Decode(&organization); err != nil {
			return nil, err
		}
		return ObjectLinker(organization), nil
	case "Page":
		page := Page{}
		if err := decoder.Decode(&page); err != nil {
			return nil, err
		}
		return ObjectLinker(page), nil
	case "Person":
		person := Person{}
		if err := decoder.Decode(&person); err != nil {
			return nil, err
		}
		return ObjectLinker(person), nil
	case "Place":
		place := Place{}
		if err := decoder.Decode(&place); err != nil {
			return nil, err
		}
		return ObjectLinker(place), nil
	case "Profile":
		profile := Profile{}
		if err := decoder.Decode(&profile); err != nil {
			return nil, err
		}
		return ObjectLinker(profile), nil
	case "Question":
		question := Question{}
		if err := decoder.Decode(&question); err != nil {
			return nil, err
		}
		return ObjectLinker(question), nil
	case "Service":
		service := Service{}
		if err := decoder.Decode(&service); err != nil {
			return nil, err
		}
		return ObjectLinker(service), nil
	case "Tombstone":
		tombStone := Tombstone{}
		if err := decoder.Decode(&tombStone); err != nil {
			return nil, err
		}
		return ObjectLinker(tombStone), nil
	case "Video":
		video := Video{}
		if err := decoder.Decode(&video); err != nil {
			return nil, err
		}
		return ObjectLinker(video), nil
	case "Accept":
		accept := Accept{}
		if err := decoder.Decode(&accept); err != nil {
			return nil, err
		}
		return ObjectLinker(accept), nil
	case "Add":
		addActivity := Add{}
		if err := decoder.Decode(&addActivity); err != nil {
			return nil, err
		}
		return ObjectLinker(addActivity), nil
	case "Announce":
		announce := Announce{}
		if err := decoder.Decode(&announce); err != nil {
			return nil, err
		}
		return ObjectLinker(announce), nil
	case "Arrive":
		arrive := Arrive{}
		if err := decoder.Decode(&arrive); err != nil {
			return nil, err
		}
		return ObjectLinker(arrive), nil
	case "Block":
		block := Block{}
		if err := decoder.Decode(&block); err != nil {
			return nil, err
		}
		return ObjectLinker(block), nil
	case "Create":
		create := Create{}
		if err := decoder.Decode(&create); err != nil {
			return nil, err
		}
		return ObjectLinker(create), nil
	case "Delete":
		deleteA := Delete{}
		if err := decoder.Decode(&deleteA); err != nil {
			return nil, err
		}
		return ObjectLinker(deleteA), nil
	case "Follow":
		follow := Follow{}
		if err := decoder.Decode(&follow); err != nil {
			return nil, err
		}
		return ObjectLinker(follow), nil
	case "Flag":
		flag := Flag{}
		if err := decoder.Decode(&flag); err != nil {
			return nil, err
		}
		return ObjectLinker(flag), nil
	case "Ignore":
		ignore := Ignore{}
		if err := decoder.Decode(&ignore); err != nil {
			return nil, err
		}
		return ObjectLinker(ignore), nil
	case "Invite":
		invite := Invite{}
		if err := decoder.Decode(&invite); err != nil {
			return nil, err
		}
		return ObjectLinker(invite), nil
	case "Join":
		join := Join{}
		if err := decoder.Decode(&join); err != nil {
			return nil, err
		}
		return ObjectLinker(join), nil
	case "Leave":
		leave := Leave{}
		if err := decoder.Decode(&leave); err != nil {
			return nil, err
		}
		return ObjectLinker(leave), nil
	case "Like":
		like := Like{}
		if err := decoder.Decode(&like); err != nil {
			return nil, err
		}
		return ObjectLinker(like), nil
	case "Listen":
		listen := Listen{}
		if err := decoder.Decode(&listen); err != nil {
			return nil, err
		}
		return ObjectLinker(listen), nil
	case "Move":
		move := Move{}
		if err := decoder.Decode(&move); err != nil {
			return nil, err
		}
		return ObjectLinker(move), nil
	case "Offer":
		offer := Offer{}
		if err := decoder.Decode(&offer); err != nil {
			return nil, err
		}
		return ObjectLinker(offer), nil
	case "Read":
		read := Read{}
		if err := decoder.Decode(&read); err != nil {
			return nil, err
		}
		return ObjectLinker(read), nil
	case "Reject":
		reject := Reject{}
		if err := decoder.Decode(&reject); err != nil {
			return nil, err
		}
		return ObjectLinker(reject), nil
	case "Remove":
		remove := Remove{}
		if err := decoder.Decode(&remove); err != nil {
			return nil, err
		}
		return ObjectLinker(remove), nil
	case "TentativeAccept":
		tentativeAccept := TentativeAccept{}
		if err := decoder.Decode(&tentativeAccept); err != nil {
			return nil, err
		}
		return ObjectLinker(tentativeAccept), nil
	case "TentativeReject":
		tentativeReject := TentativeReject{}
		if err := decoder.Decode(&tentativeReject); err != nil {
			return nil, err
		}
		return ObjectLinker(tentativeReject), nil
	case "Travel":
		travel := Travel{}
		if err := decoder.Decode(&travel); err != nil {
			return nil, err
		}
		return ObjectLinker(travel), nil
	case "Undo":
		undo := Undo{}
		if err := decoder.Decode(&undo); err != nil {
			return nil, err
		}
		return ObjectLinker(undo), nil
	case "Update":
		update := Update{}
		if err := decoder.Decode(&update); err != nil {
			return nil, err
		}
		return ObjectLinker(update), nil
	case "View":
		view := View{}
		if err := decoder.Decode(&view); err != nil {
			return nil, err
		}
		return ObjectLinker(view), nil
	case "Link":
		link := Link{}
		if err := decoder.Decode(&link); err != nil {
			return nil, err
		}
		return ObjectLinker(link), nil
	case "Mention":
		mention := Mention{}
		if err := decoder.Decode(&mention); err != nil {
			return nil, err
		}
		return ObjectLinker(mention), nil
	case "Hashtag":
		hashtag := Hashtag{}
		if err := decoder.Decode(&hashtag); err != nil {
			return nil, err
		}
		return ObjectLinker(hashtag), nil
	case "PropertyValue":
		propertyValue := PropertyValue{}
		if err := decoder.Decode(&propertyValue); err != nil {
			return nil, err
		}
		return ObjectLinker(propertyValue), nil
	}
	return nil, errors.New("unsupported ActivityStreams type, or invalid AS document")
}
