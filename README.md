# astreams

[![License: AGPL v3](https://img.shields.io/badge/License-AGPL%20v3-blue.svg)](LICENSE)

> A *hand-crafted* implementation of the [Activity Streams 2.0](https://www.w3.org/TR/activitystreams-core) specification in **Go**, especially suitable for projects implementing [ActivityPub](https://activitypub.rocks).

## Usage

[![Documentation](http://img.shields.io/badge/godoc-reference-5272B4.svg)](https://pkg.go.dev/github.com/MatejLach/astreams?tab=doc)

Example Activity Streams 2.0 object:

```go
object := `
{
  "@context": "https://www.w3.org/ns/activitystreams",
  "summary": "Martin's recent activities",
  "type": "Collection",
  "totalItems": 1,
  "items" : [
    {
      "type": "Add",
      "published": "2011-02-10T15:04:55Z",
      "generator": "http://example.org/activities-app",
      "nameMap": {
        "en": "Martin added a new image to his album.",
        "ga": "Martin phost le fisean nua a albam."
      },
      "actor": {
        "type": "Person",
        "id": "http://www.test.example/martin",
        "name": "Martin Smith",
        "url": "http://example.org/martin",
        "image": {
          "type": "Link",
          "href": "http://example.org/martin/image",
          "mediaType": "image/jpeg",
          "width": 250,
          "height": 250
        }
      },
      "object" : {
        "name": "My fluffy cat",
        "type": "Image",
        "id": "http://example.org/album/máiréad.jpg",
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
`
```

> Unmarshaling Activity Streams 2.0 JSON objects into Go objects:

```go
// Unmarshal from AS 2.0 JSON
var exobj astreams.Collection
err := json.Unmarshal([]byte(object), &exobj)
if err != nil {
	log.Println(err)
}
```
↑ An example of unmarshaling into an [astreams.Collection](https://godoc.org/github.com/MatejLach/astreams#Collection) object.

> Marshaling Go objects into Activity Streams 2.0 JSON objects:

```go
// Marshal to AS 2.0 JSON
jsB, err := json.Marshal(&exobj) // note the use of a pointer receiver here!
if err != nil {
	log.Println(err)
}
```
↑ Marshaling the same [astreams.Collection](https://godoc.org/github.com/MatejLach/astreams#Collection) back into the original ActivityStreams 2.0 JSON is equally as straightforward.

## Contributing

Contributions are welcome. If you've found a bug, or have a feature request, don't hesitate to [open an issue](https://github.com/MatejLach/astreams/issues/new).
