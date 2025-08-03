# astreams

[![License: AGPL v3](https://img.shields.io/badge/License-AGPL%20v3-blue.svg)](LICENSE)

> A *hand-crafted* implementation of the [Activity Streams 2.0](https://www.w3.org/TR/activitystreams-core) specification in **Go**, especially suitable for projects implementing [ActivityPub](https://activitypub.rocks).

## 🚀 Quick Start

### Basic Usage

```go
package main

import (
    "encoding/json"
    "fmt"
    "strings"

    "your-module/astreams"
)

func main() {
    // Decode a Note object
    noteJSON := `{
        "type": "Note",
        "content": "Hello, world!",
        "published": "2023-01-01T00:00:00Z"
    }`

    var note astreams.Note
    err := json.Unmarshal([]byte(noteJSON), &note)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Note content: %s\n", note.Content)

    // Encode back to JSON
    encoded, err := json.Marshal(&note)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(encoded))
}
```

## 📚 Core Types

### Base Types

- **`Object`** - The base ActivityStreams object type
- **`Link`** - The base Link type
- **`Actor`** - Actor types (Person, Organization, etc.)
- **`Activity`** - Activity types (Create, Like, Follow, etc.)
- **`IntransitiveActivity`** - Intransitive activities (Arrive, Travel, etc.)

### Collections

- **`Collection`** - Standard collection
- **`CollectionPage`** - Paginated collection
- **`OrderedCollection`** - Ordered collection
- **`OrderedCollectionPage`** - Paginated ordered collection

### Special Types

- **`ObjectOrLink`** - Slice that can contain any ActivityStreams object or link
- **`ObjectOrLinkOrString`** - Can hold objects/links or string URLs
- **`Icons`** - Icon representation supporting both URL and Object forms
- **`Location`** - Geographic location data

## 🔧 Working with Objects

### Decoding

```go
// Decode any ActivityStreams object using the generic decoder
func decodeObject(jsonData []byte) (astreams.ObjectLinker, error) {
    var obj astreams.ObjectOrLink
    err := json.Unmarshal(jsonData, &obj)
    if err != nil {
        return nil, err
    }

    if len(obj) > 0 {
        return obj[0], nil
    }
    return nil, errors.New("empty object")
}

// Decode specific types
func decodeNote(jsonData []byte) (*astreams.Note, error) {
    var note astreams.Note
    err := json.Unmarshal(jsonData, &note)
    return &note, err
}
```

### Encoding

```go
// Encode any ActivityStreams object
func encodeObject(obj astreams.ObjectLinker) ([]byte, error) {
    return json.Marshal(obj)
}

// Using the generic encoder
func encodeGeneric[T astreams.ActivityStreamer](toEncode T) ([]byte, error) {
    return astreams.EncodeJSON(toEncode)
}
```

## 📦 Working with Collections

## 📤 Heterogeneous slices

Slices can contain any mix of `Object` or `Link` (sub)types:

```go
// ObjectOrLink handles heterogeneous arrays
var items astreams.ObjectOrLink
items = append(items, astreams.Note{Content: "Hello"})
items = append(items, astreams.Person{Name: "Alice"})

jsonBytes, err := json.Marshal(items)
// Results in JSON array with mixed types
```

### ObjectOrLink Slices

```go
// Create and populate a heterogeneous collection
var items astreams.ObjectOrLink

// Add different types of objects
note := astreams.Note{Content: "Hello"}
person := astreams.Person{Name: "Alice"}

items = append(items, note)
items = append(items, person)

// Encode the collection
jsonBytes, err := json.Marshal(items)
```

### Collections with Pagination

```go
// Create a Collection
collection := astreams.Collection{
    TotalItems: 100,
    Items: &astreams.ObjectOrLinkOrString{
        Target: astreams.ObjectOrLink{
            astreams.Note{Content: "First note"},
            astreams.Note{Content: "Second note"},
        },
    },
}
```

## 🔄 Generic Helper Functions

### DecodeJSON

```go
// Decode any valid ActivityStreams object
note, err := astreams.DecodeJSON[astreams.Note](strings.NewReader(noteJSON))
if err != nil {
    // handle error
}

// Works with any AS type
person, err := astreams.DecodeJSON[astreams.Person](strings.NewReader(personJSON))
activity, err := astreams.DecodeJSON[astreams.Create](strings.NewReader(activityJSON))
```

### EncodeJSON

```go
// Generic encoder for any valid ActivityStreams object
note := astreams.Note{Content: "Hello"}
bytes, err := astreams.EncodeJSON(note)
if err != nil {
    // handle error
}
```

## 📡 Working with Actors

### Person Example

```go
person := astreams.Person{
    ID:                "https://example.com/users/alice",
    Type:              "Person",
    Name:              "Alice Smith",
    PreferredUsername: "alice",
    Summary:           "A person on the web",
    Inbox: &astreams.StringWithOrderedCollection{
        URL: "https://example.com/users/alice/inbox",
    },
    Outbox: &astreams.StringWithOrderedCollection{
        URL: "https://example.com/users/alice/outbox",
    },
}
```

### Actor Collections

```go
// Create a followers collection
followers := astreams.Collection{
    TotalItems: 42,
    First: &astreams.StringWithCollectionPage{
        CollectionPage: astreams.CollectionPage{
            // pagination details
        },
        URL: "https://example.com/users/alice/followers?page=1",
    },
}
```

## 📝 Working with Activities

### Create Activity Example

```go
createActivity := astreams.Create{
    Type: "Create",
    Actor: &astreams.ObjectOrLinkOrString{
        Target: astreams.ObjectOrLink{
            astreams.Person{
                ID:   "https://example.com/users/alice",
                Name: "Alice Smith",
            },
        },
    },
    ActivityObject: &astreams.ObjectOrLinkOrString{
        Target: astreams.ObjectOrLink{
            astreams.Note{
                Content: "Hello, world!",
                Published: &time.Now(),
            },
        },
    },
}
```

## 🔍 Type Checking and Conversion

### Using ConcreteType

```go
func processObject(obj astreams.ObjectLinker) {
    reflectType, asType := astreams.ConcreteType(obj)
    fmt.Printf("Reflection type: %s, AS type: %s\n", reflectType, asType)

    switch asType {
    case "Note":
        if note, ok := obj.(astreams.Note); ok {
            // Handle Note
        }
    case "Person":
        if person, ok := obj.(astreams.Person); ok {
            // Handle Person
        }
    }
}
```

### Type Assertions

```go
// Safe type assertions
func asNote(obj astreams.ObjectLinker) (*astreams.Note, bool) {
    if n, ok := obj.(astreams.Note); ok {
        return &n, true
    }
    return nil, false
}

func asPerson(obj astreams.ObjectLinker) (*astreams.Person, bool) {
    if p, ok := obj.(astreams.Person); ok {
        return &p, true
    }
    return nil, false
}
```

### Type Analysis

```go
func decodePayload(jsonData []byte) (astreams.JsonPayload, error) {
    reader := bytes.NewReader(jsonData)
    payloadType, err := astreams.DecodePayloadObjectType(reader)
    if err != nil {
        return payloadType, err
    }

    switch payloadType.Type {
    case "Note":
        // Handle Note
    case "Person":
        // Handle Person
    default:
        // Handle generic object or unknown type
    }

    return payloadType, nil
}
```

## 📋 Supported Types

### Object Types
- `Object`, `Link`
- `Article`, `Document`, `Audio`, `Video`, `Image`
- `Event`, `Note`, `Page`
- `Collection`, `CollectionPage`
- `OrderedCollection`, `OrderedCollectionPage`
- `Location`, `Profile`, `Tombstone`
- `Relationship`, `Question`

### Actor Types
- `Application`, `Group`, `Organization`, `Person`, `Service`

### Activity Types
- `Accept`, `Add`, `Announce`, `Arrive`, `Block`, `Create`, `Delete`
- `Dislike`, `Follow`, `Flag`, `Ignore`, `Invite`, `Join`, `Leave`
- `Like`, `Listen`, `Move`, `Offer`, `Read`, `Reject`, `Remove`
- `TentativeAccept`, `TentativeReject`, `Travel`, `Undo`, `Update`, `View`

## Contributing

Contributions are welcome. If you've found a bug, or have a feature request, don't hesitate to [open an issue](https://github.com/MatejLach/astreams/issues/new).
