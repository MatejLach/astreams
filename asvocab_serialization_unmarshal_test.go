package astreams

import (
	"encoding/json"
	"testing"
)

func TestUnmarshalJSON_ObjectOrLink(t *testing.T) {
	var obj ObjectOrLink
	var err error
	testCases := []string{
		`{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Object",
  "id": "http://www.test.example/object/1",
  "name": "A Simple, non-specific object"
}`,

		`{"@context": "https://www.w3.org/ns/activitystreams",
 "type": "Person",
 "id": "https://social.example/alyssa/",
 "name": "Alyssa P. Hacker",
 "preferredUsername": "alyssa",
 "summary": "Lisp enthusiast hailing from MIT",
 "inbox": "https://social.example/alyssa/inbox/",
 "outbox": "https://social.example/alyssa/outbox/",
 "followers": "https://social.example/alyssa/followers/",
 "following": "https://social.example/alyssa/following/",
 "liked": "https://social.example/alyssa/liked/"}`,

		`{"@context": "https://www.w3.org/ns/activitystreams",
 "type": "Note",
 "to": ["https://chatty.example/ben/"],
 "attributedTo": "https://social.example/alyssa/",
 "content": "Say, did you finish reading that book I lent you?"}`,

		`{
  "@context": ["https://www.w3.org/ns/activitystreams",
               {"@language": "en"}],
  "type": "Note",
  "id": "http://postparty.example/p/2415",
  "content": "<p>I <em>really</em> like strawberries!</p>",
  "source": {
    "content": "I *really* like strawberries!",
    "mediaType": "text/markdown"}
}`,
	}

	for _, tc := range testCases {
		err = json.Unmarshal([]byte(tc), &obj)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestUnmarshalJSON_ObjectOrLinkOrString(t *testing.T) {
	var obj ObjectOrLinkOrString
	var err error
	testCases := []string{
		`{
  "@context": [
    "https://www.w3.org/ns/activitystreams",
    {
      "@language": "en"
    }],
  "type": "Object",
  "name": "This is the title"
}`,

		`{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "A note",
  "type": "Note",
  "content": "My dog has fleas."
}`,

		`{
  "@context": {
     "@vocab": "https://www.w3.org/ns/activitystreams",
     "ext": "https://canine-extension.example/terms/",
     "@language": "en"
  },
  "summary": "A note",
  "type": "Note",
  "content": "My dog has fleas."
}`,

		`{
  "@context": [
     "https://www.w3.org/ns/activitystreams",
     {
      "css": "http://www.w3.org/ns/oa#styledBy"
     }
  ],
  "summary": "A note",
  "type": "Note",
  "content": "My dog has fleas."
}`,

		`{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "A simple note",
  "type": "Note",
  "content": "A simple note",
  "icon": [
    {
      "type": "Image",
      "summary": "Note (16x16)",
      "url": "http://example.org/note1.png",
      "width": 16,
      "height": 16
    },
    {
      "type": "Image",
      "summary": "Note (32x32)",
      "url": "http://example.org/note2.png",
      "width": 32,
      "height": 32
    }
  ]
}`,

		`{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "A simple note",
  "type": "Note",
  "content": "This is all there is.",
  "inReplyTo": {
    "summary": "Previous note",
    "type": "Note",
    "content": "What else is there?"
  }
}`,

		`{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Person",
  "name": "Sally",
  "location": {
    "name": "Over the Arabian Sea, east of Socotra Island Nature Sanctuary",
    "type": "Place",
    "longitude": 12.34,
    "latitude": 56.78,
    "altitude": 90,
    "units": "m"
  }
}`,

		`{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Image",
  "summary": "Picture of Sally",
  "url": "http://example.org/sally.jpg",
  "tag": [
    {
      "type": "Person",
      "id": "http://sally.example.org",
      "name": "Sally"
    }
  ]
}`,

		`{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Sally offered the post to John",
  "type": "Offer",
  "actor": "http://sally.example.org",
  "object": "http://example.org/posts/1",
  "target": {
    "type": "Person",
    "name": "John"
  }
}`,
	}

	for _, tc := range testCases {
		err = json.Unmarshal([]byte(tc), &obj)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestUnmarshalJSON_Link(t *testing.T) {
	var lnk Link
	var err error
	testCases := []string{
		`{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Link",
  "href": "http://example.org/abc",
  "hreflang": "en",
  "mediaType": "text/html",
  "name": "An example link"
}`,
	}

	for _, tc := range testCases {
		err = json.Unmarshal([]byte(tc), &lnk)
		if err != nil {
			t.Fatal(err)
		}
	}
}
