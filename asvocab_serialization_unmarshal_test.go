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

		`
{
  "@context": [
    "https://www.w3.org/ns/activitystreams",
    "https://w3id.org/security/v1",
    {
      "manuallyApprovesFollowers": "as:manuallyApprovesFollowers",
      "toot": "http://joinmastodon.org/ns#",
      "featured": {
        "@id": "toot:featured",
        "@type": "@id"
      },
      "alsoKnownAs": {
        "@id": "as:alsoKnownAs",
        "@type": "@id"
      },
      "movedTo": {
        "@id": "as:movedTo",
        "@type": "@id"
      },
      "schema": "http://schema.org#",
      "PropertyValue": "schema:PropertyValue",
      "value": "schema:value",
      "Hashtag": "as:Hashtag",
      "Emoji": "toot:Emoji",
      "IdentityProof": "toot:IdentityProof",
      "focalPoint": {
        "@container": "@list",
        "@id": "toot:focalPoint"
      }
    }
  ],
  "id": "https://social.matej-lach.me/users/MatejLach",
  "type": "Person",
  "following": "https://social.matej-lach.me/users/MatejLach/following",
  "followers": "https://social.matej-lach.me/users/MatejLach/followers",
  "inbox": "https://social.matej-lach.me/users/MatejLach/inbox",
  "outbox": "https://social.matej-lach.me/users/MatejLach/outbox",
  "featured": "https://social.matej-lach.me/users/MatejLach/collections/featured",
  "preferredUsername": "MatejLach",
  "name": "Matej Ľach  ✅",
  "summary": "<p>Free software enthusiast, <a href=\"https://social.matej-lach.me/tags/golang\" class=\"mention hashtag\" rel=\"tag\">#<span>golang</span></a>, <a href=\"https://social.matej-lach.me/tags/rustlang\" class=\"mention hashtag\" rel=\"tag\">#<span>rustlang</span></a>, <a href=\"https://social.matej-lach.me/tags/jvm\" class=\"mention hashtag\" rel=\"tag\">#<span>jvm</span></a> &amp; <a href=\"https://social.matej-lach.me/tags/swiftlang\" class=\"mention hashtag\" rel=\"tag\">#<span>swiftlang</span></a>  . Working on a question/answer <a href=\"https://social.matej-lach.me/tags/activitypub\" class=\"mention hashtag\" rel=\"tag\">#<span>ActivityPub</span></a> server. <a href=\"https://social.matej-lach.me/tags/systemd\" class=\"mention hashtag\" rel=\"tag\">#<span>systemd</span></a> aficionado :-)</p>",
  "url": "https://social.matej-lach.me/@MatejLach",
  "manuallyApprovesFollowers": false,
  "publicKey": {
    "id": "https://social.matej-lach.me/users/MatejLach#main-key",
    "owner": "https://social.matej-lach.me/users/MatejLach",
    "publicKeyPem": "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAm9ir9e6zclOuWFN+mrtV\nqjz+qNYdgYYA7TGLWf4heNjdQRP0wOTjSY5mXST95aKY8h30LSyPAI01/GLrPsme\nm+uSgr59VX3dyvDHYiOSSuMl8lkFMGthxWlKXcUb7vFcJnHRr4q5TUZ+J4wjEksw\nWmqK5me2Lnt+wVQnWXplfnknVJaZvCPEfRWVVu53lgSfTkF+rO4Bl6osw2TrIj3T\n8MoOpnGKXSTGuL86cAQAkxbJcqkFeM/ksojVBqVpGn+xdQuOf62j6mFzZl4B9wfo\nKah+O7zbgvJEHhSmvSZlo8b9YJre0YbAJBCcQnAyj3m2oSEwIKz20jjTapsiAJFS\nTwIDAQAB\n-----END PUBLIC KEY-----\n"
  },
  "tag": [
    {
      "type": "Hashtag",
      "href": "https://social.matej-lach.me/explore/golang",
      "name": "#golang"
    },
    {
      "type": "Hashtag",
      "href": "https://social.matej-lach.me/explore/activitypub",
      "name": "#activitypub"
    },
    {
      "type": "Hashtag",
      "href": "https://social.matej-lach.me/explore/rustlang",
      "name": "#rustlang"
    },
    {
      "type": "Hashtag",
      "href": "https://social.matej-lach.me/explore/swiftlang",
      "name": "#swiftlang"
    },
    {
      "type": "Hashtag",
      "href": "https://social.matej-lach.me/explore/jvm",
      "name": "#jvm"
    },
    {
      "type": "Hashtag",
      "href": "https://social.matej-lach.me/explore/systemd",
      "name": "#systemd"
    }
  ],
  "attachment": [
    {
      "type": "PropertyValue",
      "name": "Blog",
      "value": "<a href=\"https://blog.matej-lach.me\" rel=\"me nofollow noopener\" target=\"_blank\"><span class=\"invisible\">https://</span><span class=\"\">blog.matej-lach.me</span><span class=\"invisible\"></span></a>"
    }
  ],
  "endpoints": {
    "sharedInbox": "https://social.matej-lach.me/inbox"
  },
  "icon": {
    "type": "Image",
    "mediaType": "image/png",
    "url": "https://social.matej-lach.me/system/accounts/avatars/000/000/001/original/6e9242b03795bf80.png?1509060490"
  },
  "image": {
    "type": "Image",
    "mediaType": "image/png",
    "url": "https://social.matej-lach.me/system/accounts/headers/000/000/001/original/0865cab093a8367b.png?1536786967"
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
