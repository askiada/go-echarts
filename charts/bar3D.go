package charts

import "github.com/go-echarts/go-echarts/types"

// Bar3D represents a 3D bar chart.
type Bar3D struct {
	Chart3D
}

func (Bar3D) Type() string { return types.ChartBar3D }

// NewBar3D creates a new 3D bar chart.
func NewBar3D() *Bar3D {
	chart := &Bar3D{}
	chart.initBaseConfiguration()
	chart.initChart3D()
	return chart
}

// AddXYAxis adds both the X axis and the Y axis.
func (c *Bar3D) AddXYAxis(xAxis, yAxis interface{}) *Bar3D {
	c.xData = xAxis
	c.yData = yAxis
	return c
}

// AddZAxis adds the Z axis.
func (c *Bar3D) AddZAxis(name string, zAxis interface{}, opts ...SeriesOpts) *Bar3D {
	c.addZAxis(types.ChartBar3D, name, zAxis, opts...)
	return c
}
