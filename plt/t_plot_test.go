// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plt

import (
	"path"
	"testing"

	"github.com/cpmech/gosl/chk"
	"github.com/cpmech/gosl/io"
)

func Test_PlotSimpleCurve(tst *testing.T) {

	verbose()
	chk.PrintTitle("PlotSimpleCurve")

	Init(ParamsControl{UsePython: true})
	Plot([]float64{0, 1, 2, 3}, []float64{0, 1, 2, 3})
	Plot([]float64{0, 1, 2, 3}, []float64{0, 1, 4, 9}, ParamsPlot{Marker: "*"})
	Save("PlotSimpleCurve")
	res := io.ReadFile(path.Join(paramsControl.TmpDir, "gosl.py"))
	sol := `### file generated by Gosl #################################################
import numpy as np
import matplotlib.pyplot as plt
import matplotlib.ticker as tck
import matplotlib.patches as pat
import matplotlib.path as pth
import matplotlib.patheffects as pff
import matplotlib.lines as lns
import mpl_toolkits.mplot3d as m3d
NaN = np.NaN
EXTRA_ARTISTS = []
def addToEA(obj):
    if obj!=None: EXTRA_ARTISTS.append(obj)
COLORMAPS = [plt.cm.bwr, plt.cm.RdBu, plt.cm.hsv, plt.cm.jet, plt.cm.terrain, plt.cm.pink, plt.cm.Greys]
def getCmap(idx): return COLORMAPS[idx % len(COLORMAPS)]
x582=np.array([0,1,2,3,],dtype=float)
y582=np.array([0,1,2,3,],dtype=float)
plt.plot(x582,y582)
x678=np.array([0,1,2,3,],dtype=float)
y678=np.array([0,1,4,9,],dtype=float)
plt.plot(x678,y678,marker='*')
plt.savefig(r'/tmp/gosl/plt/PlotSimpleCurve.png', bbox_inches='tight', bbox_extra_artists=EXTRA_ARTISTS)
`
	chk.String(tst, string(res), sol)
}