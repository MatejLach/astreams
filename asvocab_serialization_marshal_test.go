package astreams

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

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
        "type": "Person",
        "id": "http://www.test.example/martin",
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
        "type": "Image",
        "id": "http://example.org/album/máiréad.jpg",
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
        "type": "Collection",
        "id": "http://example.org/album/",
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
		got, err := json.MarshalIndent(collection, "", "  ")
		if err != nil {
			t.Fatal(err)
		}
		if string(got) != strings.TrimSpace(want[idx]) {
			fmt.Println(string(got))
			t.Fatal("marshaled Collection JSON not as expected")
		}
	}
}
