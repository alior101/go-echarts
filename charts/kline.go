package charts

type Kline struct {
	RectChart
}

func NewKLine(routers ...HTTPRouter) *Kline {
	chart := new(Kline)
	chart.initBaseOpts(true, routers...)
	chart.initXYOpts()
	return chart
}

func (c *Kline) AddXAxis(xAxis interface{}) *Kline {
	c.xAxisData = xAxis
	return c
}

func (c *Kline) AddYAxis(name string, yAxis interface{}, options ...interface{}) *Kline {
	series := singleSeries{Name: name, Type: "candlestick", Data: yAxis}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}