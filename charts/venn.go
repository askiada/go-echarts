package charts

import (
	"bytes"
	"encoding/json"
	"html/template"

	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
)

// Venn represents a venn chart.
type Venn struct {
	Width  int
	Heigth int
	VennChart
}

// Type returns the chart type.
func (Venn) Type() string { return types.ChartVenn }

// NewVenn creates a new venn chart instance.
func NewVenn(width, heigth int) *Venn {
	c := &Venn{
		Width:  width,
		Heigth: heigth,
	}
	c.Initialization.Validate()
	return c
}

// AddSeries adds the new series.
func (c *Venn) AddSeries(name string, data []opts.VennData, options ...SeriesOpts) *Venn {
	series := SingleSeries{Name: name, Type: types.ChartVenn, Data: data}
	c.SingleSeries = series
	return c
}

// Validate validates the given configuration.
func (c *Venn) Validate() {}

// VennConfiguration represents an option set needed by all chart types.
type VennConfiguration struct {
	render.Renderer     `json:"-"`
	opts.Initialization `json:"-"`
	SingleSeries
}

// JSON wraps all the options to a map so that it could be used in the base template
//
// Get data in bytes
// bs, _ : = json.Marshal(bar.JSON())
func (bc *VennConfiguration) JSON() interface{} {
	return bc.json()
}

// JSONNotEscaped works like method JSON, but it returns a marshaled object whose characters will not be escaped in the template
func (bc *VennConfiguration) JSONNotEscaped() template.HTML {
	obj := bc.json()
	buff := bytes.NewBufferString("")
	enc := json.NewEncoder(buff)
	enc.SetEscapeHTML(false)
	enc.Encode(obj)

	return template.HTML(buff.String())
}

func (bc *VennConfiguration) json() interface{} {
	return bc.Data
}

// GetAssets returns the Assets options.
func (bc *VennConfiguration) GetAssets() opts.Assets {
	return opts.Assets{}
}
