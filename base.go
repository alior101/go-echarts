package goecharts

// 图形初始化配置项
type InitOpts struct {
	// 生成的 HTML 页面标题
	PageTitle string `default:"Awesome go-echarts"`
	// 画布宽度
	Width string `default:"800px"`
	// 画布高度
	Height string `default:"500px"`
	// 图表 ID，是图表唯一标识
	ChartID string
	// 静态资源 host 地址
	AssetsHost string `default:"http://chenjiandongx.com/go-echarts-assets/assets/"`
	// 图表主题
	Theme string `default:"white"`
}

// 静态资源配置项
type AssetsOpts struct {
	JSAssets  []string
	CSSAssets []string
}

// 初始化静态资源配置项
func (opt *AssetsOpts) initAssetsOpts() {
	opt.JSAssets = []string{echartsJS}
	opt.CSSAssets = []string{bulmaCSS}
}

// 追加 js 资源
func (opt *AssetsOpts) appendJsAssets(asset string) {
	opt.JSAssets = append(opt.JSAssets, asset)
}

// 返回资源列表
func (opt *AssetsOpts) yieldAssets() ([]string, []string) {
	return opt.JSAssets, opt.CSSAssets
}

// 判断 js 资源是否在列表中
func (opt *AssetsOpts) jsIn(jsRef string) bool {
	isIn := false
	for i := 0; i < len(opt.JSAssets); i++ {
		if opt.JSAssets[i] == jsRef {
			isIn = true
			break
		}
	}
	return isIn
}

// 判断 css 资源是否在列表中
func (opt *AssetsOpts) cssIn(cssRef string) bool {
	isIn := false
	for i := 0; i < len(opt.CSSAssets); i++ {
		if opt.CSSAssets[i] == cssRef {
			isIn = true
			break
		}
	}
	return isIn
}

// 校验静态资源配置项，追加 host
func (opt *AssetsOpts) validateAssets(host string) {
	for i := 0; i < len(opt.JSAssets); i++ {
		opt.JSAssets[i] = host + opt.JSAssets[i]
	}
	for j := 0; j < len(opt.CSSAssets); j++ {
		opt.CSSAssets[j] = host + opt.CSSAssets[j]
	}
}

// 为 InitOptions 设置字段默认值
func (opt *InitOpts) setDefault() {
	err := setDefaultValue(opt)
	checkError(err)
}

// 确保 ContainerID 不为空且唯一
func (opt *InitOpts) checkID() {
	if opt.ChartID == "" {
		opt.ChartID = genChartID()
	}
}

// 验证初始化参数，确保图形能够得到正确渲染
func (opt *InitOpts) validateInitOpt() {
	opt.setDefault()
	opt.checkID()
}

// Http 路由
type HttpRouter struct {
	// 路由 URL
	Url string
	// 路由显示文字
	Text string
}

// Http 路由列表
type HttpRouters []HttpRouter

// Len() 用于 template 方法
func (hr HttpRouters) Len() int {
	return len(hr)
}

// 全局颜色配置项
type ColorOpts struct {
	Color []string
}

// 所有图表都拥有的基本配置项
type BaseOpts struct {
	// 图形初始化配置项
	InitOpts
	// 图例组件配置项
	LegendOpts
	// 提示框组件配置项
	TooltipOpts
	// 标题组件配置项
	TitleOpts
	// 静态资源配置项
	AssetsOpts
	// 全局颜色列表
	ColorList []string
	// 追加全局颜色列表
	appendColor []string
	// 路由列表
	HttpRouters
	// 区域缩放组件配置项列表
	DataZoomOptsList
	// 视觉映射组件配置项列表
	VisualMapOptsList
}

// 设置全局颜色
func (opt *BaseOpts) setColor(options ...interface{}) {
	for i := 0; i < len(options); i++ {
		option := options[i]
		switch option.(type) {
		case ColorOpts:
			opt.appendColor = append(opt.appendColor, option.(ColorOpts).Color...)
		}
	}
}

// 初始化全局颜色列表
func (opt *BaseOpts) initSeriesColors() {
	opt.ColorList = []string{
		"#c23531", "#2f4554", "#61a0a8", "#d48265", "#91c7ae", "#749f83",
		"#ca8622", "#bda29a", "#6e7074", "#546570", "#c4ccd3"}
}

// 初始化 BaseOpts
func (opt *BaseOpts) initBaseOpts(routers ...HttpRouter) {
	for i := 0; i < len(routers); i++ {
		opt.HttpRouters = append(opt.HttpRouters, routers[i])
	}
	opt.initSeriesColors()
}

// 插入颜色到颜色列表首部
func (opt *BaseOpts) insertSeriesColors(s []string) {
	// 翻转颜色列表
	tmpCl := reverseSlice(s)
	// 颜色追加至首部
	for i := 0; i < len(tmpCl); i++ {
		opt.ColorList = append(opt.ColorList, "")
		copy(opt.ColorList[1:], opt.ColorList[0:])
		opt.ColorList[0] = tmpCl[i]
	}
}

// 设置 BaseOptions 全局配置项
func (opt *BaseOpts) setBaseGlobalConfig(options ...interface{}) {
	for i := 0; i < len(options); i++ {
		option := options[i]
		switch option.(type) {
		case InitOpts:
			opt.InitOpts = option.(InitOpts)
			if opt.InitOpts.Theme != "" {
				opt.JSAssets = append(opt.JSAssets, "themes/"+opt.Theme+".js")
			}
		case TitleOpts:
			opt.TitleOpts = option.(TitleOpts)
		case LegendOpts:
			opt.LegendOpts = option.(LegendOpts)
		case ColorOpts:
			opt.insertSeriesColors(option.(ColorOpts).Color)
		case DataZoomOpts:
			opt.DataZoomOptsList = append(opt.DataZoomOptsList, option.(DataZoomOpts))
		case VisualMapOpts:
			opt.VisualMapOptsList = append(opt.VisualMapOptsList, option.(VisualMapOpts))
		}
	}
}