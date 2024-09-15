package astreams

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
		require.NoError(t, err)
		got, err := json.MarshalIndent(&orderedCollectionPage, "", "  ")
		require.NoError(t, err)

		assert.JSONEq(t, want[idx], string(got))
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
		require.NoError(t, err)
		got, err := json.MarshalIndent(&collection, "", "  ")
		require.NoError(t, err)

		assert.JSONEq(t, want[idx], string(got))
	}
}

func TestMarshalJSON_Actor(t *testing.T) {
	want := []string{
		`{
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
      "featuredTags": {
        "@id": "toot:featuredTags",
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
      "discoverable": "toot:discoverable",
      "Device": "toot:Device",
      "Ed25519Signature": "toot:Ed25519Signature",
      "Ed25519Key": "toot:Ed25519Key",
      "Curve25519Key": "toot:Curve25519Key",
      "EncryptedMessage": "toot:EncryptedMessage",
      "publicKeyBase64": "toot:publicKeyBase64",
      "deviceId": "toot:deviceId",
      "claim": {
        "@type": "@id",
        "@id": "toot:claim"
      },
      "fingerprintKey": {
        "@type": "@id",
        "@id": "toot:fingerprintKey"
      },
      "identityKey": {
        "@type": "@id",
        "@id": "toot:identityKey"
      },
      "devices": {
        "@type": "@id",
        "@id": "toot:devices"
      },
      "messageFranking": "toot:messageFranking",
      "messageType": "toot:messageType",
      "cipherText": "toot:cipherText",
      "suspended": "toot:suspended",
      "memorial": "toot:memorial",
      "indexable": "toot:indexable",
      "Hashtag": "as:Hashtag",
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
  "featuredTags": "https://social.matej-lach.me/users/MatejLach/collections/tags",
  "preferredUsername": "MatejLach",
  "name": "Matej Ľach  ✅",
  "summary": "<p>Free software enthusiast, <a href=\"https://social.matej-lach.me/tags/golang\" class=\"mention hashtag\" rel=\"tag\">#<span>golang</span></a>, <a href=\"https://social.matej-lach.me/tags/rustlang\" class=\"mention hashtag\" rel=\"tag\">#<span>rustlang</span></a>, <a href=\"https://social.matej-lach.me/tags/swiftlang\" class=\"mention hashtag\" rel=\"tag\">#<span>swiftlang</span></a>  . Working on a question/answer <a href=\"https://social.matej-lach.me/tags/ActivityPub\" class=\"mention hashtag\" rel=\"tag\">#<span>ActivityPub</span></a> server. <a href=\"https://social.matej-lach.me/tags/systemd\" class=\"mention hashtag\" rel=\"tag\">#<span>systemd</span></a> aficionado :-)</p>",
  "url": "https://social.matej-lach.me/@MatejLach",
  "manuallyApprovesFollowers": false,
  "discoverable": true,
  "indexable": false,
  "published": "2017-10-26T00:00:00Z",
  "memorial": false,
  "devices": "https://social.matej-lach.me/users/MatejLach/collections/devices",
  "publicKey": {
    "id": "https://social.matej-lach.me/users/MatejLach#main-key",
    "owner": "https://social.matej-lach.me/users/MatejLach",
    "publicKeyPem": "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAm9ir9e6zclOuWFN+mrtV\nqjz+qNYdgYYA7TGLWf4heNjdQRP0wOTjSY5mXST95aKY8h30LSyPAI01/GLrPsme\nm+uSgr59VX3dyvDHYiOSSuMl8lkFMGthxWlKXcUb7vFcJnHRr4q5TUZ+J4wjEksw\nWmqK5me2Lnt+wVQnWXplfnknVJaZvCPEfRWVVu53lgSfTkF+rO4Bl6osw2TrIj3T\n8MoOpnGKXSTGuL86cAQAkxbJcqkFeM/ksojVBqVpGn+xdQuOf62j6mFzZl4B9wfo\nKah+O7zbgvJEHhSmvSZlo8b9YJre0YbAJBCcQnAyj3m2oSEwIKz20jjTapsiAJFS\nTwIDAQAB\n-----END PUBLIC KEY-----\n"
  },
  "tag": [
    {
      "type": "Hashtag",
      "href": "https://social.matej-lach.me/tags/golang",
      "name": "#golang"
    },
    {
      "type": "Hashtag",
      "href": "https://social.matej-lach.me/tags/activitypub",
      "name": "#activitypub"
    },
    {
      "type": "Hashtag",
      "href": "https://social.matej-lach.me/tags/rustlang",
      "name": "#rustlang"
    },
    {
      "type": "Hashtag",
      "href": "https://social.matej-lach.me/tags/swiftlang",
      "name": "#swiftlang"
    },
    {
      "type": "Hashtag",
      "href": "https://social.matej-lach.me/tags/systemd",
      "name": "#systemd"
    }
  ],
  "attachment": [],
  "endpoints": {
    "sharedInbox": "https://social.matej-lach.me/inbox"
  },
  "icon": {
    "type": "Image",
    "mediaType": "image/png",
    "url": "https://social.matej-lach.me/system/accounts/avatars/000/000/001/original/6e9242b03795bf80.png"
  },
  "image": {
    "type": "Image",
    "mediaType": "image/png",
    "url": "https://social.matej-lach.me/system/accounts/headers/000/000/001/original/f18240c45b0ac254.png"
  }
}`,
	}

	for idx := range want {
		// Prepare the object to marshal first
		person := Person{}
		err := json.Unmarshal([]byte(want[idx]), &person)
		require.NoError(t, err)
		got, err := json.MarshalIndent(&person, "", "  ")
		require.NoError(t, err)

		assert.JSONEq(t, want[idx], string(got))
	}
}
