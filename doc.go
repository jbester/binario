//  Package binaryio provides a wrapper to encoding/binary to provide a concise API.
//
//  Numbers are translated by reading and writing fixed-size values. A fixed-size value
//  is either a fixed-size arithmetic type (bool, int8, uint8, int16, ...) or an array
//  only fixed-size values.
//
//  This package is intended for decoding and encoding existing binary formats and
//  is not written as a high performance serializer deserializer.
package binaryio
