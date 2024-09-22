package astreams

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
)

// Implements https://golang.org/pkg/encoding/json/#Unmarshaler
func (ol *ObjectOrLink) UnmarshalJSON(data []byte) error {
	if bytes.HasPrefix(bytes.TrimSpace(data), []byte{'['}) {
		var rawJSONSlice []json.RawMessage
		if err := json.Unmarshal(data, &rawJSONSlice); err != nil {
			return err
		}
		for _, jsonObj := range rawJSONSlice {
			rawJsonMap := make(map[string]interface{})
			if err := json.Unmarshal(jsonObj, &rawJsonMap); err != nil {
				return err
			}
			if key, ok := rawJsonMap["type"]; ok {
				if objectType, ok := key.(string); ok {
					astype, err := decodeASType(jsonObj, objectType)
					if err != nil {
						return err
					}
					*ol = append(*ol, astype)
				}
			} else {
				// assuming a base Object if there's not a more specific type
				// a Link without an explicit type is considered invalid
				asObject := Object{}
				if err := json.Unmarshal(data, &asObject); err != nil {
					return err
				}
				*ol = append(*ol, asObject)
			}
		}
	} else if bytes.HasPrefix(bytes.TrimSpace(data), []byte{'{'}) {
		rawJsonMap := make(map[string]interface{})
		if err := json.Unmarshal(data, &rawJsonMap); err != nil {
			return err
		}
		if key, ok := rawJsonMap["type"]; ok {
			if objectType, ok := key.(string); ok {
				astype, err := decodeASType(data, objectType)
				if err != nil {
					return err
				}
				*ol = append(*ol, astype)
			}
		} else {
			// assuming a base Object if there's not a more specific type
			// a Link without an explicit type is considered invalid
			asObject := Object{}
			if err := json.Unmarshal(data, &asObject); err != nil {
				return err
			}
			*ol = append(*ol, asObject)
		}
	}
	return nil
}

// Implements https://golang.org/pkg/encoding/json/#Unmarshaler
func (ics *Icons) UnmarshalJSON(data []byte) error {
	if bytes.HasPrefix(bytes.TrimSpace(data), []byte{'['}) {
		var rawJSONSlice []json.RawMessage
		if err := json.Unmarshal(data, &rawJSONSlice); err != nil {
			return err
		}
		for _, bts := range rawJSONSlice {
			if bytes.HasPrefix(bytes.TrimSpace(bts), []byte{'"'}) {
				if err := json.Unmarshal(bts, &ics.URL); err != nil {
					return err
				}
				if _, err := url.ParseRequestURI(ics.URL); err != nil {
					return err
				}
			}
			if bytes.HasPrefix(bytes.TrimSpace(bts), []byte{'{'}) {
				decoder := json.NewDecoder(bytes.NewReader(bts))
				icon := Icon{}
				if err := decoder.Decode(&icon); err != nil {
					return err
				}
				ics.Icons = append(ics.Icons, icon)
			}
		}
	} else if bytes.HasPrefix(bytes.TrimSpace(data), []byte{'{'}) {
		decoder := json.NewDecoder(bytes.NewReader(data))
		icon := Icon{}
		if err := decoder.Decode(&icon); err != nil {
			return err
		}
		ics.Icons = append(ics.Icons, icon)
	} else if bytes.HasPrefix(bytes.TrimSpace(data), []byte{'"'}) {
		if err := json.Unmarshal(data, &ics.URL); err != nil {
			return err
		}
		if _, err := url.ParseRequestURI(ics.URL); err != nil {
			return err
		}
	}
	return nil
}

// Implements https://golang.org/pkg/encoding/json/#Unmarshaler
func (ols *ObjectOrLinkOrString) UnmarshalJSON(data []byte) error {
	if bytes.HasPrefix(bytes.TrimSpace(data), []byte{'['}) {
		var rawJSONSlice []json.RawMessage
		if err := json.Unmarshal(data, &rawJSONSlice); err != nil {
			return err
		}
		for _, bts := range rawJSONSlice {
			if bytes.HasPrefix(bytes.TrimSpace(bts), []byte{'"'}) {
				var strlink string
				if err := json.Unmarshal(bts, &strlink); err != nil {
					return err
				}
				if _, err := url.ParseRequestURI(strlink); err != nil {
					return err
				}
				ols.URL = append(ols.URL, strlink)
			}
			if bytes.HasPrefix(bytes.TrimSpace(bts), []byte{'{'}) {
				if err := json.Unmarshal(bts, &ols.Target); err != nil {
					return err
				}
			}
		}
	} else if bytes.HasPrefix(bytes.TrimSpace(data), []byte{'{'}) {
		if err := json.Unmarshal(data, &ols.Target); err != nil {
			return err
		}
	} else if bytes.HasPrefix(bytes.TrimSpace(data), []byte{'"'}) {
		var strlink string
		if err := json.Unmarshal(data, &strlink); err != nil {
			return err
		}
		if _, err := url.ParseRequestURI(strlink); err != nil {
			return err
		}
		ols.URL = append(ols.URL, strlink)
	}
	return nil
}

// Implements https://golang.org/pkg/encoding/json/#Unmarshaler
func (socp *StringWithOrderedCollectionPage) UnmarshalJSON(data []byte) error {
	if bytes.HasPrefix(bytes.TrimSpace(data), []byte{'"'}) {
		if err := json.Unmarshal(data, &socp.URL); err != nil {
			return err
		}
		if _, err := url.ParseRequestURI(socp.URL); err != nil {
			return err
		}
	} else if bytes.HasPrefix(bytes.TrimSpace(data), []byte{'{'}) {
		if err := json.Unmarshal(data, &socp.OrdCollectionPage); err != nil {
			return err
		}
	}
	return nil
}

// Implements https://golang.org/pkg/encoding/json/#Unmarshaler
func (soc *StringWithOrderedCollection) UnmarshalJSON(data []byte) error {
	if bytes.HasPrefix(bytes.TrimSpace(data), []byte{'"'}) {
		if err := json.Unmarshal(data, &soc.URL); err != nil {
			return err
		}
		if _, err := url.ParseRequestURI(soc.URL); err != nil {
			return err
		}
	} else if bytes.HasPrefix(bytes.TrimSpace(data), []byte{'{'}) {
		if err := json.Unmarshal(data, &soc.OrdCollection); err != nil {
			return err
		}
	}
	return nil
}

// Implements https://golang.org/pkg/encoding/json/#Unmarshaler
func (scp *StringWithCollectionPage) UnmarshalJSON(data []byte) error {
	if bytes.HasPrefix(bytes.TrimSpace(data), []byte{'"'}) {
		if err := json.Unmarshal(data, &scp.URL); err != nil {
			return err
		}
		if _, err := url.ParseRequestURI(scp.URL); err != nil {
			return err
		}
	} else if bytes.HasPrefix(bytes.TrimSpace(data), []byte{'{'}) {
		if err := json.Unmarshal(data, &scp.CollectionPage); err != nil {
			return err
		}
	}
	return nil
}

// Implements https://golang.org/pkg/encoding/json/#Unmarshaler
func (sc *StringWithCollection) UnmarshalJSON(data []byte) error {
	if bytes.HasPrefix(bytes.TrimSpace(data), []byte{'"'}) {
		if err := json.Unmarshal(data, &sc.URL); err != nil {
			return err
		}
		if _, err := url.ParseRequestURI(sc.URL); err != nil {
			return err
		}
	} else if bytes.HasPrefix(bytes.TrimSpace(data), []byte{'{'}) {
		if err := json.Unmarshal(data, &sc.Collection); err != nil {
			return err
		}
	}
	return nil
}

// Implements https://golang.org/pkg/encoding/json/#Unmarshaler
func (enp *EndpointsOrString) UnmarshalJSON(data []byte) error {
	if bytes.HasPrefix(bytes.TrimSpace(data), []byte{'"'}) {
		if err := json.Unmarshal(data, &enp.URL); err != nil {
			return err
		}
		if _, err := url.ParseRequestURI(enp.URL); err != nil {
			return err
		}
	} else if bytes.HasPrefix(bytes.TrimSpace(data), []byte{'{'}) {
		decoder := json.NewDecoder(bytes.NewReader(data))
		if err := decoder.Decode(&enp.Endpoints); err != nil {
			return err
		}
	}
	return nil
}

// Implements https://golang.org/pkg/encoding/json/#Marshal
func (ol *ObjectOrLink) MarshalJSON() ([]byte, error) {
	if len(*ol) > 0 {
		if len(*ol) == 1 {
			objl := *ol
			target, err := json.Marshal(objl[0])
			if err != nil {
				return []byte{}, err
			}

			return target, nil
		}

		marshalB := bytes.NewBufferString("[")
		for idx, target := range *ol {
			encoded, err := encodeASType(target)
			if err != nil {
				return []byte{}, err
			}
			marshalB.Write(encoded)
			if idx < len(*ol)-1 {
				_, err := marshalB.WriteString(",")
				if err != nil {
					return []byte{}, err
				}
			}
		}
		if _, err := marshalB.WriteString("]"); err != nil {
			return []byte{}, err
		}
		return marshalB.Bytes(), nil
	}
	return bytes.NewBufferString("[]").Bytes(), nil
}

// Implements https://golang.org/pkg/encoding/json/#Marshal
func (ols *ObjectOrLinkOrString) MarshalJSON() ([]byte, error) {
	var uri []byte
	var target []byte
	var err error

	if len(ols.URL) > 0 && len(ols.Target) > 0 {
		for _, uri := range ols.URL {
			if _, err = url.ParseRequestURI(uri); err != nil {
				return []byte{}, err
			}
		}
		if len(ols.URL) == 1 {
			uri, err = json.Marshal(ols.URL[0])
			if err != nil {
				return []byte{}, err
			}
		} else {
			var uris string
			for idx, link := range ols.URL {
				if idx < len(ols.URL)-1 {
					uris += fmt.Sprintf("\"%s\",\n", link)
				} else {
					uris += fmt.Sprintf("\"%s\"\n", link)
				}
			}
			uri = bytes.NewBufferString(uris).Bytes()
		}
		if len(ols.Target) == 1 {
			target, err = json.Marshal(ols.Target[0])
			if err != nil {
				return []byte{}, err
			}
		} else {
			target, err = json.Marshal(&ols.Target)
			if err != nil {
				return []byte{}, err
			}
		}
		return bytes.NewBufferString(fmt.Sprintf("[%s,%s]", string(uri), string(target))).Bytes(), nil
	} else if len(ols.URL) == 0 && len(ols.Target) > 0 {
		if len(ols.Target) > 1 {
			target, err = json.Marshal(&ols.Target)
			if err != nil {
				return []byte{}, err
			}
			return target, nil
		}
		target, err = json.Marshal(ols.Target[0])
		if err != nil {
			return []byte{}, err
		}
		return target, nil
	} else if len(ols.URL) > 0 && len(ols.Target) == 0 {
		for _, uri := range ols.URL {
			if _, err = url.ParseRequestURI(uri); err != nil {
				return []byte{}, err
			}
		}
		if len(ols.URL) == 1 {
			uri, err = json.Marshal(ols.URL[0])
			if err != nil {
				return []byte{}, err
			}
			return uri, nil
		}
		uri, err = json.Marshal(ols.URL)
		if err != nil {
			return []byte{}, err
		}
		return uri, nil
	}

	if len(ols.Target) == 0 || len(ols.URL) == 0 {
		return bytes.NewBufferString("[]").Bytes(), nil
	}

	return []byte{}, errors.New("unrecognised content, cannot Marshal ObjectOrLinkOrString, use nil for empty value")
}

// Implements https://golang.org/pkg/encoding/json/#Marshal
func (colp *CollectionPage) MarshalJSON() ([]byte, error) {
	encodedBase, err := json.Marshal(&colp.Collection)
	if err != nil {
		return []byte{}, err
	}
	encodedBase = bytes.TrimSuffix(encodedBase, []byte("}"))

	encodedPage, err := json.Marshal(struct {
		Next   *ObjectOrLinkOrString `json:"next,omitempty"`
		PartOf *ObjectOrLinkOrString `json:"partOf,omitempty"`
		Prev   *ObjectOrLinkOrString `json:"prev,omitempty"`
	}{
		Next:   colp.Next,
		PartOf: colp.PartOf,
		Prev:   colp.Prev,
	})
	if err != nil {
		return []byte{}, err
	}
	encodedPage = bytes.TrimPrefix(encodedPage, []byte("{"))

	return bytes.NewBufferString(fmt.Sprintf("%s, %s", encodedBase, encodedPage)).Bytes(), nil
}

// Implements https://golang.org/pkg/encoding/json/#Marshal
func (col *Collection) MarshalJSON() ([]byte, error) {
	encodedBase, err := json.Marshal(ObjectLinker(col).GetObject())
	if err != nil {
		return []byte{}, err
	}
	encodedBase = bytes.TrimSuffix(encodedBase, []byte("}"))

	encodedCollection, err := json.Marshal(struct {
		TotalItems int                   `json:"totalItems,omitempty"`
		Current    *ObjectOrLinkOrString `json:"current,omitempty"`
		First      *ObjectOrLinkOrString `json:"first,omitempty"`
		Last       *ObjectOrLinkOrString `json:"last,omitempty"`
	}{
		TotalItems: col.TotalItems,
		Current:    col.Current,
		First:      col.First,
		Last:       col.Last,
	})
	if err != nil {
		return []byte{}, err
	}
	encodedCollection = bytes.TrimPrefix(encodedCollection, []byte("{"))

	if len(col.Items.URL) > 0 {
		encodedCollection = bytes.TrimSuffix(encodedCollection, []byte("}"))
		encodedBytes, err := json.Marshal(&col.Items.URL)
		if err != nil {
			return []byte{}, err
		}

		return bytes.NewBufferString(fmt.Sprintf("%s, %s, \"items\": \n%s}", encodedBase, encodedCollection, encodedBytes)).Bytes(), nil
	}

	if len(col.Items.Target) > 0 {
		encodedCollection = bytes.TrimSuffix(encodedCollection, []byte("}"))
		encodedBytes, err := json.Marshal(&col.Items.Target)
		if err != nil {
			return []byte{}, err
		}

		return bytes.NewBufferString(fmt.Sprintf("%s, %s, \"items\": [\n%s\n]\n}", encodedBase, encodedCollection, encodedBytes)).Bytes(), nil
	}

	return []byte{}, nil
}

// Implements https://golang.org/pkg/encoding/json/#Marshal
func (ocolp *OrderedCollectionPage) MarshalJSON() ([]byte, error) {
	encodedBase, err := json.Marshal(&ocolp.OrderedCollection)
	if err != nil {
		return []byte{}, err
	}
	encodedBase = bytes.TrimSuffix(encodedBase, []byte("}"))

	encodedPage, err := json.Marshal(struct {
		Next       *ObjectOrLinkOrString `json:"next,omitempty"`
		PartOf     *ObjectOrLinkOrString `json:"partOf,omitempty"`
		Prev       *ObjectOrLinkOrString `json:"prev,omitempty"`
		StartIndex uint                  `json:"startIndex,omitempty"`
	}{
		Next:       ocolp.Next,
		PartOf:     ocolp.PartOf,
		Prev:       ocolp.Prev,
		StartIndex: ocolp.StartIndex,
	})
	if err != nil {
		return []byte{}, err
	}
	encodedPage = bytes.TrimPrefix(encodedPage, []byte("{"))

	return bytes.NewBufferString(fmt.Sprintf("%s, %s", encodedBase, encodedPage)).Bytes(), nil
}

// Implements https://golang.org/pkg/encoding/json/#Marshal
func (oc *OrderedCollection) MarshalJSON() ([]byte, error) {
	encodedBase, err := json.Marshal(ObjectLinker(oc).GetObject())
	if err != nil {
		return []byte{}, err
	}
	encodedBase = bytes.TrimSuffix(encodedBase, []byte("}"))

	encodedOrderedCollection, err := json.Marshal(struct {
		TotalItems int                   `json:"totalItems,omitempty"`
		Current    *ObjectOrLinkOrString `json:"current,omitempty"`
		First      *ObjectOrLinkOrString `json:"first,omitempty"`
		Last       *ObjectOrLinkOrString `json:"last,omitempty"`
	}{
		TotalItems: oc.TotalItems,
		Current:    oc.Current,
		First:      oc.First,
		Last:       oc.Last,
	})
	if err != nil {
		return []byte{}, err
	}
	encodedOrderedCollection = bytes.TrimPrefix(encodedOrderedCollection, []byte("{"))

	if len(oc.OrderedItems.URL) > 0 {
		encodedOrderedCollection = bytes.TrimSuffix(encodedOrderedCollection, []byte("}"))
		encodedBytes, err := json.Marshal(&oc.OrderedItems.URL)
		if err != nil {
			return []byte{}, err
		}

		return bytes.NewBufferString(fmt.Sprintf("%s, %s, \"orderedItems\": \n%s}", encodedBase, encodedOrderedCollection, encodedBytes)).Bytes(), nil
	}

	if len(oc.OrderedItems.Target) > 0 {
		encodedOrderedCollection = bytes.TrimSuffix(encodedOrderedCollection, []byte("}"))
		encodedBytes, err := json.Marshal(&oc.OrderedItems.Target)
		if err != nil {
			return []byte{}, err
		}

		return bytes.NewBufferString(fmt.Sprintf("%s, %s, \"orderedItems\": [\n%s\n]\n}", encodedBase, encodedOrderedCollection, encodedBytes)).Bytes(), nil
	}

	return []byte{}, nil
}

// Implements https://golang.org/pkg/encoding/json/#Marshal
func (ics *Icons) MarshalJSON() ([]byte, error) {
	if ics.URL != "" && len(ics.Icons) > 0 {
		if _, err := url.ParseRequestURI(ics.URL); err != nil {
			return []byte{}, err
		}
		uri, err := json.Marshal(ics.URL)
		if err != nil {
			return []byte{}, err
		}
		var icn []byte
		if len(ics.Icons) > 1 {
			icn, err = json.Marshal(ics.Icons)
		} else {
			icn, err = json.Marshal(ics.Icons[0])
		}
		if err != nil {
			return []byte{}, err
		}
		return bytes.NewBufferString(fmt.Sprintf("[%s,%s]", string(uri), string(icn))).Bytes(), nil
	} else if ics.URL == "" && len(ics.Icons) > 0 {
		if len(ics.Icons) > 1 {
			icn, err := json.Marshal(ics.Icons)
			if err != nil {
				return []byte{}, err
			}
			return icn, nil
		}
		icn, err := json.Marshal(ics.Icons[0])
		if err != nil {
			return []byte{}, err
		}
		return icn, nil
	} else if ics.URL != "" && len(ics.Icons) == 0 {
		if _, err := url.ParseRequestURI(ics.URL); err != nil {
			return []byte{}, err
		}
		uri, err := json.Marshal(ics.URL)
		if err != nil {
			return []byte{}, err
		}
		return uri, nil
	}
	return []byte{}, errors.New("unrecognised content, cannot Marshal icon, use nil for empty value")
}

// Implements https://golang.org/pkg/encoding/json/#Marshal
func (socp *StringWithOrderedCollectionPage) MarshalJSON() ([]byte, error) {
	if len(socp.URL) > 0 {
		jsonb, err := json.Marshal(socp.URL)
		if err != nil {
			return []byte{}, err
		}
		return jsonb, nil
	} else if socp.OrdCollectionPage.PartOf != nil {
		jsonb, err := json.Marshal(socp.OrdCollectionPage)
		if err != nil {
			return []byte{}, err
		}
		return jsonb, nil
	}
	return []byte{}, errors.New("unrecognised content, cannot Marshal StringWithOrderedCollectionPage, use nil for empty value")
}

// Implements https://golang.org/pkg/encoding/json/#Marshal
func (soc *StringWithOrderedCollection) MarshalJSON() ([]byte, error) {
	if len(soc.URL) > 0 {
		jsonb, err := json.Marshal(soc.URL)
		if err != nil {
			return []byte{}, err
		}
		return jsonb, nil
	} else if soc.OrdCollection.TotalItems > 0 {
		jsonb, err := json.Marshal(soc.OrdCollection)
		if err != nil {
			return []byte{}, err
		}
		return jsonb, nil
	}
	return []byte{}, errors.New("unrecognised content, cannot Marshal StringWithOrderedCollection, use nil for empty value")
}

// Implements https://golang.org/pkg/encoding/json/#Marshal
func (scp *StringWithCollectionPage) MarshalJSON() ([]byte, error) {
	if len(scp.URL) > 0 {
		jsonb, err := json.Marshal(scp.URL)
		if err != nil {
			return []byte{}, err
		}
		return jsonb, nil
	} else if scp.CollectionPage.PartOf != nil {
		jsonb, err := json.Marshal(scp.CollectionPage)
		if err != nil {
			return []byte{}, err
		}
		return jsonb, nil
	}
	return []byte{}, errors.New("unrecognised content, cannot Marshal StringWithCollectionPage, use nil for empty value")
}

// Implements https://golang.org/pkg/encoding/json/#Marshal
func (sc *StringWithCollection) MarshalJSON() ([]byte, error) {
	if len(sc.URL) > 0 {
		jsonb, err := json.Marshal(sc.URL)
		if err != nil {
			return []byte{}, err
		}
		return jsonb, nil
	} else if sc.Collection.TotalItems > 0 {
		jsonb, err := json.Marshal(sc.Collection)
		if err != nil {
			return []byte{}, err
		}
		return jsonb, nil
	}
	return []byte{}, errors.New("unrecognised content, cannot Marshal StringWithCollection, use nil for empty value")
}

// Implements https://golang.org/pkg/encoding/json/#Marshal
func (enp *EndpointsOrString) MarshalJSON() ([]byte, error) {
	if len(enp.URL) > 0 {
		jsonb, err := json.Marshal(enp.URL)
		if err != nil {
			return []byte{}, err
		}
		return jsonb, nil
	} else if !((Endpoint{}) == enp.Endpoints) {
		jsonb, err := json.Marshal(enp.Endpoints)
		if err != nil {
			return []byte{}, err
		}
		return jsonb, nil
	}
	return []byte{}, errors.New("unrecognised content, cannot Marshal EndpointsOrString, use nil for empty value")
}

// DecodeJSON is the generic unmarshaller for any valid ActivityStreams 2.0 object
func DecodeJSON[T ActivityStreamer](jsonPayload io.Reader) (T, error) {
	var result T

	err := json.NewDecoder(jsonPayload).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// EncodeJSON is the generic marshaller for any valid ActivityStreams 2.0 object
func EncodeJSON[T ActivityStreamer](toEncode T) ([]byte, error) {
	return json.Marshal(&toEncode)
}

// Implements https://golang.org/pkg/fmt/#Stringer
func (ics *Icons) String() string {
	return fmt.Sprintf("%+v", *ics)
}

// Implements https://golang.org/pkg/fmt/#Stringer
func (ols *ObjectOrLinkOrString) String() string {
	return fmt.Sprintf("%+v", *ols)
}

// Implements https://golang.org/pkg/fmt/#Stringer
func (loc *Location) String() string {
	return fmt.Sprintf("%+v", *loc)
}

// Implements https://golang.org/pkg/fmt/#Stringer
func (oc *OrderedCollection) String() string {
	return fmt.Sprintf("%+v", *oc)
}

// Implements https://golang.org/pkg/fmt/#Stringer
func (soc *StringWithOrderedCollection) String() string {
	return fmt.Sprintf("%+v", *soc)
}

// Implements https://golang.org/pkg/fmt/#Stringer
func (sc *StringWithCollection) String() string {
	return fmt.Sprintf("%+v", *sc)
}

// Implements https://golang.org/pkg/fmt/#Stringer
func (enp *EndpointsOrString) String() string {
	return fmt.Sprintf("%+v", *enp)
}
