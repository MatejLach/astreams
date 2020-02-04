/*
Package astreams implements JSON serialization/deserialization support for the ActivityStreams 2.0 vocabulary.
Simply use json.Marshal/json.Unmarshal on AS 2.0 compliant objects to obtain their JSON/Go representation respectively.

Indirect object references such as a link to an outbox collection within an actor profile need to be fetched explicitly and parsed as a separate object.
*/
package astreams
