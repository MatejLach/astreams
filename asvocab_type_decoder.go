package astreams

import (
	"bytes"
	"encoding/json"
	"errors"
)

// Decode the supplied JSON into the appropriate type based on its "type" property
func decodeASType(data []byte, jsonType string) (Targeter, error) {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	switch jsonType {
	case "Activity":
		activity := Activity{}
		if err := decoder.Decode(&activity); err != nil {
			return nil, err
		}
		return Targeter(activity), nil
	case "Application":
		application := Application{}
		if err := decoder.Decode(&application); err != nil {
			return nil, err
		}
		return Targeter(application), nil
	case "Article":
		article := Article{}
		if err := decoder.Decode(&article); err != nil {
			return nil, err
		}
		return Targeter(article), nil
	case "Audio":
		audio := Audio{}
		if err := decoder.Decode(&audio); err != nil {
			return nil, err
		}
		return Targeter(audio), nil
	case "Collection":
		collection := Collection{}
		if err := decoder.Decode(&collection); err != nil {
			return nil, err
		}
		return Targeter(collection), nil
	case "CollectionPage":
		colPage := CollectionPage{}
		if err := decoder.Decode(&colPage); err != nil {
			return nil, err
		}
		return Targeter(colPage), nil
	case "Relationship":
		relationship := Relationship{}
		if err := decoder.Decode(&relationship); err != nil {
			return nil, err
		}
		return Targeter(relationship), nil
	case "Document":
		document := Document{}
		if err := decoder.Decode(&document); err != nil {
			return nil, err
		}
		return Targeter(document), nil
	case "Event":
		event := Event{}
		if err := decoder.Decode(&event); err != nil {
			return nil, err
		}
		return Targeter(event), nil
	case "Group":
		group := Group{}
		if err := decoder.Decode(&group); err != nil {
			return nil, err
		}
		return Targeter(group), nil
	case "Image":
		image := Image{}
		if err := decoder.Decode(&image); err != nil {
			return nil, err
		}
		return Targeter(image), nil
	case "IntransitiveActivity":
		intransitiveActivity := IntransitiveActivity{}
		if err := decoder.Decode(&intransitiveActivity); err != nil {
			return nil, err
		}
		return Targeter(intransitiveActivity), nil
	case "Note":
		note := Note{}
		if err := decoder.Decode(&note); err != nil {
			return nil, err
		}
		return Targeter(note), nil
	case "Object":
		object := Object{}
		if err := decoder.Decode(&object); err != nil {
			return nil, err
		}
		return Targeter(object), nil
	case "OrderedCollection":
		ordCollection := OrderedCollection{}
		if err := decoder.Decode(&ordCollection); err != nil {
			return nil, err
		}
		return Targeter(ordCollection), nil
	case "OrderedCollectionPage":
		ordColPage := OrderedCollectionPage{}
		if err := decoder.Decode(&ordColPage); err != nil {
			return nil, err
		}
		return Targeter(ordColPage), nil
	case "Organization":
		organization := Organization{}
		if err := decoder.Decode(&organization); err != nil {
			return nil, err
		}
		return Targeter(organization), nil
	case "Page":
		page := Page{}
		if err := decoder.Decode(&page); err != nil {
			return nil, err
		}
		return Targeter(page), nil
	case "Person":
		person := Person{}
		if err := decoder.Decode(&person); err != nil {
			return nil, err
		}
		return Targeter(person), nil
	case "Place":
		place := Place{}
		if err := decoder.Decode(&place); err != nil {
			return nil, err
		}
		return Targeter(place), nil
	case "Profile":
		profile := Profile{}
		if err := decoder.Decode(&profile); err != nil {
			return nil, err
		}
		return Targeter(profile), nil
	case "Question":
		question := Question{}
		if err := decoder.Decode(&question); err != nil {
			return nil, err
		}
		return Targeter(question), nil
	case "Service":
		service := Service{}
		if err := decoder.Decode(&service); err != nil {
			return nil, err
		}
		return Targeter(service), nil
	case "Tombstone":
		tombStone := Tombstone{}
		if err := decoder.Decode(&tombStone); err != nil {
			return nil, err
		}
		return Targeter(tombStone), nil
	case "Video":
		video := Video{}
		if err := decoder.Decode(&video); err != nil {
			return nil, err
		}
		return Targeter(video), nil
	case "Accept":
		accept := Accept{}
		if err := decoder.Decode(&accept); err != nil {
			return nil, err
		}
		return Targeter(accept), nil
	case "Add":
		addActivity := Add{}
		if err := decoder.Decode(&addActivity); err != nil {
			return nil, err
		}
		return Targeter(addActivity), nil
	case "Announce":
		announce := Announce{}
		if err := decoder.Decode(&announce); err != nil {
			return nil, err
		}
		return Targeter(announce), nil
	case "Arrive":
		arrive := Arrive{}
		if err := decoder.Decode(&arrive); err != nil {
			return nil, err
		}
		return Targeter(arrive), nil
	case "Block":
		block := Block{}
		if err := decoder.Decode(&block); err != nil {
			return nil, err
		}
		return Targeter(block), nil
	case "Create":
		create := Create{}
		if err := decoder.Decode(&create); err != nil {
			return nil, err
		}
		return Targeter(create), nil
	case "Delete":
		deleteA := Delete{}
		if err := decoder.Decode(&deleteA); err != nil {
			return nil, err
		}
		return Targeter(deleteA), nil
	case "Follow":
		follow := Follow{}
		if err := decoder.Decode(&follow); err != nil {
			return nil, err
		}
		return Targeter(follow), nil
	case "Flag":
		flag := Flag{}
		if err := decoder.Decode(&flag); err != nil {
			return nil, err
		}
		return Targeter(flag), nil
	case "Ignore":
		ignore := Ignore{}
		if err := decoder.Decode(&ignore); err != nil {
			return nil, err
		}
		return Targeter(ignore), nil
	case "Invite":
		invite := Invite{}
		if err := decoder.Decode(&invite); err != nil {
			return nil, err
		}
		return Targeter(invite), nil
	case "Join":
		join := Join{}
		if err := decoder.Decode(&join); err != nil {
			return nil, err
		}
		return Targeter(join), nil
	case "Leave":
		leave := Leave{}
		if err := decoder.Decode(&leave); err != nil {
			return nil, err
		}
		return Targeter(leave), nil
	case "Like":
		like := Like{}
		if err := decoder.Decode(&like); err != nil {
			return nil, err
		}
		return Targeter(like), nil
	case "Listen":
		listen := Listen{}
		if err := decoder.Decode(&listen); err != nil {
			return nil, err
		}
		return Targeter(listen), nil
	case "Move":
		move := Move{}
		if err := decoder.Decode(&move); err != nil {
			return nil, err
		}
		return Targeter(move), nil
	case "Offer":
		offer := Offer{}
		if err := decoder.Decode(&offer); err != nil {
			return nil, err
		}
		return Targeter(offer), nil
	case "Read":
		read := Read{}
		if err := decoder.Decode(&read); err != nil {
			return nil, err
		}
		return Targeter(read), nil
	case "Reject":
		reject := Reject{}
		if err := decoder.Decode(&reject); err != nil {
			return nil, err
		}
		return Targeter(reject), nil
	case "Remove":
		remove := Remove{}
		if err := decoder.Decode(&remove); err != nil {
			return nil, err
		}
		return Targeter(remove), nil
	case "TentativeAccept":
		tentativeAccept := TentativeAccept{}
		if err := decoder.Decode(&tentativeAccept); err != nil {
			return nil, err
		}
		return Targeter(tentativeAccept), nil
	case "TentativeReject":
		tentativeReject := TentativeReject{}
		if err := decoder.Decode(&tentativeReject); err != nil {
			return nil, err
		}
		return Targeter(tentativeReject), nil
	case "Travel":
		travel := Travel{}
		if err := decoder.Decode(&travel); err != nil {
			return nil, err
		}
		return Targeter(travel), nil
	case "Undo":
		undo := Undo{}
		if err := decoder.Decode(&undo); err != nil {
			return nil, err
		}
		return Targeter(undo), nil
	case "Update":
		update := Update{}
		if err := decoder.Decode(&update); err != nil {
			return nil, err
		}
		return Targeter(update), nil
	case "View":
		view := View{}
		if err := decoder.Decode(&view); err != nil {
			return nil, err
		}
		return Targeter(view), nil
	case "Link":
		link := Link{}
		if err := decoder.Decode(&link); err != nil {
			return nil, err
		}
		return Targeter(link), nil
	case "Mention":
		mention := Mention{}
		if err := decoder.Decode(&mention); err != nil {
			return nil, err
		}
		return Targeter(mention), nil
	case "Hashtag":
		hashtag := Hashtag{}
		if err := decoder.Decode(&hashtag); err != nil {
			return nil, err
		}
		return Targeter(hashtag), nil
	case "PropertyValue":
		propertyValue := PropertyValue{}
		if err := decoder.Decode(&propertyValue); err != nil {
			return nil, err
		}
		return Targeter(propertyValue), nil
	}
	return nil, errors.New("unsupported ActivityStreams type, or invalid AS document")
}
