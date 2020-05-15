// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plt

import (
	"bytes"

	"github.com/cpmech/gosl/utl"
)

// CurveStyle holds definitions for the style of curves
type CurveStyle struct {
	// lines
	LineColor string  `json:"lineColor"` // color
	LineAlpha float64 `json:"lineAlpha"` // alpha (0, 1]. A<1e-14 => A=1.0
	LineStyle string  `json:"lineStyle"` // style
	LineWidth float64 `json:"lineWidth"` // width

	// markers
	MarkerType      string  `json:"markerType"`      // type
	MarkerImg       string  `json:"markerImg"`       // image filename
	MarkerColor     string  `json:"markerColor"`     // color
	MarkerAlpha     float64 `json:"markerAlpha"`     // alpha (0, 1]
	MarkerSize      float64 `json:"markerSize"`      // size; when using images, set markerSize=0 to use the image width
	MarkerEvery     float64 `json:"markerEvery"`     // mark-every
	MarkerLineColor string  `json:"markerLineColor"` // edge color
	MarkerLineWidth float64 `json:"markerLineWidth"` // edge width
	MarkerLineStyle string  `json:"markerLineStyle"` // edge style
	MarkerIsVoid    bool    `json:"markerIsVoid"`    // void marker (draw edge only)
}

// Curve defines the curve data
type Curve struct {
	Style         CurveStyle `json:"style"`         // line and marker arguments
	Label         string     `json:"label"`         // curve name or connection pair such as 'San Francisco -> Los Angeles'
	X             []float64  `json:"x"`             // x-coordinates
	Y             []float64  `json:"y"`             // y-coordinates
	Z             []float64  `json:"z,omitempty"`   // [optional] z-coordinates
	Kind          string     `json:"kind"`          // e.g. connection, city, fortress, base, mine, ...
	TagFirstPoint bool       `json:"tagFirstPoint"` // tag first point with label

	// for Python only
	FigElevation float64 `json:"figElevation"` // figure elevation (z-index)
	NoClip       bool    `json:"noClip"`       // turn clipping off
}

// Encode encodes Curve into JSON string
func (o *Curve) Encode() []byte {
	buf := new(bytes.Buffer)
	enc := utl.NewEncoder(buf, "json")
	enc.Encode(o)
	return buf.Bytes()
}

// Curves holds a list of curves
type Curves struct {
	List []Curve
}

// DefaultCurveStyle defines the default style
var DefaultCurveStyle = CurveStyle{
	// lines
	LineColor: "#b33434",
	LineAlpha: 0.7,
	LineStyle: "-",
	LineWidth: 3,

	// markers
	MarkerType:      "o",
	MarkerImg:       "",
	MarkerColor:     "#4c4deb",
	MarkerAlpha:     1,
	MarkerSize:      0,
	MarkerEvery:     0,
	MarkerLineColor: "#ffffff",
	MarkerLineWidth: 2,
	MarkerLineStyle: "none",
	MarkerIsVoid:    false,
}
