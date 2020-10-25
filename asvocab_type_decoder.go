package astreams

import (
	"bytes"
	"encoding/json"
	"errors"
)

// Decode the supplied JSON into the appropriate type based on its "type" property
func decodeASType(data []byte, jsonType string) (ActivityStreamer, error) {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	switch jsonType {
	case "Activity":
		activity := Activity{}
		if err := decoder.Decode(&activity); err != nil {
			return nil, err
		}
		return ActivityStreamer(activity), nil
	case "Application":
		application := Application{}
		if err := decoder.Decode(&application); err != nil {
			return nil, err
		}
		return ActivityStreamer(application), nil
	case "Article":
		article := Article{}
		if err := decoder.Decode(&article); err != nil {
			return nil, err
		}
		return ActivityStreamer(article), nil
	case "Audio":
		audio := Audio{}
		if err := decoder.Decode(&audio); err != nil {
			return nil, err
		}
		return ActivityStreamer(audio), nil
	case "Collection":
		collection := Collection{}
		if err := decoder.Decode(&collection); err != nil {
			return nil, err
		}
		return ActivityStreamer(collection), nil
	case "CollectionPage":
		colPage := CollectionPage{}
		if err := decoder.Decode(&colPage); err != nil {
			return nil, err
		}
		return ActivityStreamer(colPage), nil
	case "Relationship":
		relationship := Relationship{}
		if err := decoder.Decode(&relationship); err != nil {
			return nil, err
		}
		return ActivityStreamer(relationship), nil
	case "Document":
		document := Document{}
		if err := decoder.Decode(&document); err != nil {
			return nil, err
		}
		return ActivityStreamer(document), nil
	case "Event":
		event := Event{}
		if err := decoder.Decode(&event); err != nil {
			return nil, err
		}
		return ActivityStreamer(event), nil
	case "Group":
		group := Group{}
		if err := decoder.Decode(&group); err != nil {
			return nil, err
		}
		return ActivityStreamer(group), nil
	case "Image":
		image := Image{}
		if err := decoder.Decode(&image); err != nil {
			return nil, err
		}
		return ActivityStreamer(image), nil
	case "IntransitiveActivity":
		intransitiveActivity := IntransitiveActivity{}
		if err := decoder.Decode(&intransitiveActivity); err != nil {
			return nil, err
		}
		return ActivityStreamer(intransitiveActivity), nil
	case "Note":
		note := Note{}
		if err := decoder.Decode(&note); err != nil {
			return nil, err
		}
		return ActivityStreamer(note), nil
	case "Object":
		object := Object{}
		if err := decoder.Decode(&object); err != nil {
			return nil, err
		}
		return ActivityStreamer(object), nil
	case "OrderedCollection":
		ordCollection := OrderedCollection{}
		if err := decoder.Decode(&ordCollection); err != nil {
			return nil, err
		}
		return ActivityStreamer(ordCollection), nil
	case "OrderedCollectionPage":
		ordColPage := OrderedCollectionPage{}
		if err := decoder.Decode(&ordColPage); err != nil {
			return nil, err
		}
		return ActivityStreamer(ordColPage), nil
	case "Organization":
		organization := Organization{}
		if err := decoder.Decode(&organization); err != nil {
			return nil, err
		}
		return ActivityStreamer(organization), nil
	case "Page":
		page := Page{}
		if err := decoder.Decode(&page); err != nil {
			return nil, err
		}
		return ActivityStreamer(page), nil
	case "Person":
		person := Person{}
		if err := decoder.Decode(&person); err != nil {
			return nil, err
		}
		return ActivityStreamer(person), nil
	case "Place":
		place := Place{}
		if err := decoder.Decode(&place); err != nil {
			return nil, err
		}
		return ActivityStreamer(place), nil
	case "Profile":
		profile := Profile{}
		if err := decoder.Decode(&profile); err != nil {
			return nil, err
		}
		return ActivityStreamer(profile), nil
	case "Question":
		question := Question{}
		if err := decoder.Decode(&question); err != nil {
			return nil, err
		}
		return ActivityStreamer(question), nil
	case "Service":
		service := Service{}
		if err := decoder.Decode(&service); err != nil {
			return nil, err
		}
		return ActivityStreamer(service), nil
	case "Tombstone":
		tombStone := Tombstone{}
		if err := decoder.Decode(&tombStone); err != nil {
			return nil, err
		}
		return ActivityStreamer(tombStone), nil
	case "Video":
		video := Video{}
		if err := decoder.Decode(&video); err != nil {
			return nil, err
		}
		return ActivityStreamer(video), nil
	case "Accept":
		accept := Accept{}
		if err := decoder.Decode(&accept); err != nil {
			return nil, err
		}
		return ActivityStreamer(accept), nil
	case "Add":
		addActivity := Add{}
		if err := decoder.Decode(&addActivity); err != nil {
			return nil, err
		}
		return ActivityStreamer(addActivity), nil
	case "Announce":
		announce := Announce{}
		if err := decoder.Decode(&announce); err != nil {
			return nil, err
		}
		return ActivityStreamer(announce), nil
	case "Arrive":
		arrive := Arrive{}
		if err := decoder.Decode(&arrive); err != nil {
			return nil, err
		}
		return ActivityStreamer(arrive), nil
	case "Block":
		block := Block{}
		if err := decoder.Decode(&block); err != nil {
			return nil, err
		}
		return ActivityStreamer(block), nil
	case "Create":
		create := Create{}
		if err := decoder.Decode(&create); err != nil {
			return nil, err
		}
		return ActivityStreamer(create), nil
	case "Delete":
		deleteA := Delete{}
		if err := decoder.Decode(&deleteA); err != nil {
			return nil, err
		}
		return ActivityStreamer(deleteA), nil
	case "Follow":
		follow := Follow{}
		if err := decoder.Decode(&follow); err != nil {
			return nil, err
		}
		return ActivityStreamer(follow), nil
	case "Flag":
		flag := Flag{}
		if err := decoder.Decode(&flag); err != nil {
			return nil, err
		}
		return ActivityStreamer(flag), nil
	case "Ignore":
		ignore := Ignore{}
		if err := decoder.Decode(&ignore); err != nil {
			return nil, err
		}
		return ActivityStreamer(ignore), nil
	case "Invite":
		invite := Invite{}
		if err := decoder.Decode(&invite); err != nil {
			return nil, err
		}
		return ActivityStreamer(invite), nil
	case "Join":
		join := Join{}
		if err := decoder.Decode(&join); err != nil {
			return nil, err
		}
		return ActivityStreamer(join), nil
	case "Leave":
		leave := Leave{}
		if err := decoder.Decode(&leave); err != nil {
			return nil, err
		}
		return ActivityStreamer(leave), nil
	case "Like":
		like := Like{}
		if err := decoder.Decode(&like); err != nil {
			return nil, err
		}
		return ActivityStreamer(like), nil
	case "Listen":
		listen := Listen{}
		if err := decoder.Decode(&listen); err != nil {
			return nil, err
		}
		return ActivityStreamer(listen), nil
	case "Move":
		move := Move{}
		if err := decoder.Decode(&move); err != nil {
			return nil, err
		}
		return ActivityStreamer(move), nil
	case "Offer":
		offer := Offer{}
		if err := decoder.Decode(&offer); err != nil {
			return nil, err
		}
		return ActivityStreamer(offer), nil
	case "Read":
		read := Read{}
		if err := decoder.Decode(&read); err != nil {
			return nil, err
		}
		return ActivityStreamer(read), nil
	case "Reject":
		reject := Reject{}
		if err := decoder.Decode(&reject); err != nil {
			return nil, err
		}
		return ActivityStreamer(reject), nil
	case "Remove":
		remove := Remove{}
		if err := decoder.Decode(&remove); err != nil {
			return nil, err
		}
		return ActivityStreamer(remove), nil
	case "TentativeAccept":
		tentativeAccept := TentativeAccept{}
		if err := decoder.Decode(&tentativeAccept); err != nil {
			return nil, err
		}
		return ActivityStreamer(tentativeAccept), nil
	case "TentativeReject":
		tentativeReject := TentativeReject{}
		if err := decoder.Decode(&tentativeReject); err != nil {
			return nil, err
		}
		return ActivityStreamer(tentativeReject), nil
	case "Travel":
		travel := Travel{}
		if err := decoder.Decode(&travel); err != nil {
			return nil, err
		}
		return ActivityStreamer(travel), nil
	case "Undo":
		undo := Undo{}
		if err := decoder.Decode(&undo); err != nil {
			return nil, err
		}
		return ActivityStreamer(undo), nil
	case "Update":
		update := Update{}
		if err := decoder.Decode(&update); err != nil {
			return nil, err
		}
		return ActivityStreamer(update), nil
	case "View":
		view := View{}
		if err := decoder.Decode(&view); err != nil {
			return nil, err
		}
		return ActivityStreamer(view), nil
	case "Link":
		link := Link{}
		if err := decoder.Decode(&link); err != nil {
			return nil, err
		}
		return ActivityStreamer(link), nil
	case "Mention":
		mention := Mention{}
		if err := decoder.Decode(&mention); err != nil {
			return nil, err
		}
		return ActivityStreamer(mention), nil
	case "Hashtag":
		hashtag := Hashtag{}
		if err := decoder.Decode(&hashtag); err != nil {
			return nil, err
		}
		return ActivityStreamer(hashtag), nil
	case "PropertyValue":
		propertyValue := PropertyValue{}
		if err := decoder.Decode(&propertyValue); err != nil {
			return nil, err
		}
		return ActivityStreamer(propertyValue), nil
	}
	return nil, errors.New("unsupported ActivityStreams type, or invalid AS document")
}
