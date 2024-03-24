package astreams

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestMarshalJSON_OrderedCollectionPage(t *testing.T) {
	want := []string{
		`{
  "@context": "https://www.w3.org/ns/activitystreams",
  "id": "https://social.matej-lach.me/users/MatejLach/following?page=1",
  "type": "OrderedCollectionPage",
  "totalItems": 329,
  "orderedItems": [
    "https://suma-ev.social/users/christian",
    "https://mastodon.social/users/Daojoan",
    "https://23.social/users/pantierra",
    "https://famichiki.jp/users/tsturm",
    "https://bikeshed.vibber.net/users/brooke",
    "https://idlethumbs.social/users/ja2ke",
    "https://oisaur.com/users/renchap",
    "https://social.alternativebit.fr/users/picnoir",
    "https://swiss.social/users/lx",
    "https://social.lol/users/whakkee",
    "https://oldbytes.space/users/aperezdc",
    "https://tech.lgbt/users/jaycie"
  ],
  "next": "https://social.matej-lach.me/users/MatejLach/following?page=2",
  "partOf": "https://social.matej-lach.me/users/MatejLach/following"
}`,
	}

	for idx := range want {
		// Prepare the object to marshal first
		orderedCollectionPage := OrderedCollectionPage{}
		err := json.Unmarshal([]byte(want[idx]), &orderedCollectionPage)
		if err != nil {
			t.Fatal(err)
		}
		got, err := json.MarshalIndent(&orderedCollectionPage, "", "  ")
		if err != nil {
			t.Fatal(err)
		}
		if string(got) != strings.TrimSpace(want[idx]) {
			t.Fatalf("marshaled OrderedCollectionPage JSON not as expected, wanted '%s', got '%s'", strings.TrimSpace(want[idx]), string(got))
		}
	}
}

func TestMarshalJSON_Collection(t *testing.T) {
	want := []string{
		`
{
  "@context": "https://www.w3.org/ns/activitystreams",
  "type": "Collection",
  "summary": "Martin's recent activities",
  "totalItems": 1,
  "items": [
    {
      "type": "Add",
      "nameMap": {
        "en": "Martin added a new image to his album.",
        "ga": "Martin phost le fisean nua a albam."
      },
      "generator": "http://example.org/activities-app",
      "published": "2011-02-10T15:04:55Z",
      "actor": {
        "id": "http://www.test.example/martin",
        "type": "Person",
        "name": "Martin Smith",
        "image": {
          "type": "Link",
          "href": "http://example.org/martin/image",
          "mediaType": "image/jpeg",
          "height": 250,
          "width": 250
        },
        "url": "http://example.org/martin"
      },
      "object": {
        "id": "http://example.org/album/máiréad.jpg",
        "type": "Image",
        "name": "My fluffy cat",
        "preview": {
          "type": "Link",
          "href": "http://example.org/album/máiréad.jpg",
          "mediaType": "image/jpeg"
        },
        "url": [
          {
            "type": "Link",
            "href": "http://example.org/album/máiréad.jpg",
            "mediaType": "image/jpeg"
          },
          {
            "type": "Link",
            "href": "http://example.org/album/máiréad.png",
            "mediaType": "image/png"
          }
        ]
      },
      "target": {
        "id": "http://example.org/album/",
        "type": "Collection",
        "nameMap": {
          "en": "Martin's Photo Album",
          "ga": "Grianghraif Mairtin"
        },
        "image": {
          "type": "Link",
          "href": "http://example.org/album/thumbnail.jpg",
          "mediaType": "image/jpeg"
        }
      }
    }
  ]
}
`,
	}

	for idx := range want {
		// Prepare the object to marshal first
		collection := Collection{}
		err := json.Unmarshal([]byte(want[idx]), &collection)
		if err != nil {
			t.Fatal(err)
		}
		got, err := json.MarshalIndent(&collection, "", "  ")
		if err != nil {
			t.Fatal(err)
		}
		if string(got) != strings.TrimSpace(want[idx]) {
			t.Fatalf("marshaled Collection JSON not as expected, wanted '%s', got '%s'", strings.TrimSpace(want[idx]), string(got))
		}
	}
}

func TestMarshalJSON_Actor(t *testing.T) {
	want := []string{
		`
{
  "@context": [
    "https://www.w3.org/ns/activitystreams",
    "https://w3id.org/security/v1",
    {
      "schema": "http://schema.org#",
      "alsoKnownAs": {
        "@type": "@id",
        "@id": "as:alsoKnownAs"
      },
      "movedTo": {
        "@type": "@id",
        "@id": "as:movedTo"
      }
    }
  ],
  "id": "https://social.matej-lach.me/users/MatejLach",
  "type": "Person",
  "summary": "\u003cp\u003eFree software enthusiast, \u003ca href=\"https://social.matej-lach.me/tags/golang\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003egolang\u003c/span\u003e\u003c/a\u003e, \u003ca href=\"https://social.matej-lach.me/tags/rustlang\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003erustlang\u003c/span\u003e\u003c/a\u003e, \u003ca href=\"https://social.matej-lach.me/tags/jvm\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003ejvm\u003c/span\u003e\u003c/a\u003e \u0026amp; \u003ca href=\"https://social.matej-lach.me/tags/swiftlang\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003eswiftlang\u003c/span\u003e\u003c/a\u003e  . Working on a question/answer \u003ca href=\"https://social.matej-lach.me/tags/ActivityPub\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003eActivityPub\u003c/span\u003e\u003c/a\u003e server. \u003ca href=\"https://social.matej-lach.me/tags/systemd\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003esystemd\u003c/span\u003e\u003c/a\u003e aficionado :-)\u003c/p\u003e",
  "name": "Matej Ľach  ✅",
  "attachment": {
    "type": "PropertyValue",
    "name": "Blog",
    "value": "\u003ca href=\"https://blog.matej-lach.me\" rel=\"me nofollow noopener\" target=\"_blank\"\u003e\u003cspan class=\"invisible\"\u003ehttps://\u003c/span\u003e\u003cspan class=\"\"\u003eblog.matej-lach.me\u003c/span\u003e\u003cspan class=\"invisible\"\u003e\u003c/span\u003e\u003c/a\u003e"
  },
  "featured": "https://social.matej-lach.me/users/MatejLach/collections/featured",
  "icon": {
    "type": "Image",
    "url": "https://social.matej-lach.me/system/accounts/avatars/000/000/001/original/6e9242b03795bf80.png?1509060490",
    "mediaType": "image/png"
  },
  "image": {
    "type": "Image",
    "url": "https://social.matej-lach.me/system/accounts/headers/000/000/001/original/f18240c45b0ac254.png?1576335545",
    "mediaType": "image/png"
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
  "url": "https://social.matej-lach.me/@MatejLach",
  "publicKey": {
    "id": "https://social.matej-lach.me/users/MatejLach#main-key",
    "owner": "https://social.matej-lach.me/users/MatejLach",
    "publicKeyPem": "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAm9ir9e6zclOuWFN+mrtV\nqjz+qNYdgYYA7TGLWf4heNjdQRP0wOTjSY5mXST95aKY8h30LSyPAI01/GLrPsme\nm+uSgr59VX3dyvDHYiOSSuMl8lkFMGthxWlKXcUb7vFcJnHRr4q5TUZ+J4wjEksw\nWmqK5me2Lnt+wVQnWXplfnknVJaZvCPEfRWVVu53lgSfTkF+rO4Bl6osw2TrIj3T\n8MoOpnGKXSTGuL86cAQAkxbJcqkFeM/ksojVBqVpGn+xdQuOf62j6mFzZl4B9wfo\nKah+O7zbgvJEHhSmvSZlo8b9YJre0YbAJBCcQnAyj3m2oSEwIKz20jjTapsiAJFS\nTwIDAQAB\n-----END PUBLIC KEY-----\n"
  },
  "inbox": "https://social.matej-lach.me/users/MatejLach/inbox",
  "outbox": "https://social.matej-lach.me/users/MatejLach/outbox",
  "followers": "https://social.matej-lach.me/users/MatejLach/followers",
  "following": "https://social.matej-lach.me/users/MatejLach/following",
  "preferredUsername": "MatejLach",
  "endpoints": {
    "sharedInbox": "https://social.matej-lach.me/inbox"
  }
}
`,
	}

	for idx := range want {
		// Prepare the object to marshal first
		person := Person{}
		err := json.Unmarshal([]byte(want[idx]), &person)
		if err != nil {
			t.Fatal(err)
		}
		got, err := json.MarshalIndent(&person, "", "  ")
		if err != nil {
			t.Fatal(err)
		}
		if string(got) != strings.TrimSpace(want[idx]) {
			t.Fatalf("marshaled Actor JSON not as expected, wanted '%s', got '%s'", strings.TrimSpace(want[idx]), string(got))
		}
	}
}
