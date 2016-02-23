package controllers

import (
	"github.com/astaxie/beego"
)

// CMS API
type ExamplesController struct {
	beego.Controller
}

func (c *ExamplesController) URLMapping() {
	c.Mapping("StaticBlock", c.StaticBlock)
	c.Mapping("AllBlock", c.AllBlock)
}

// @router /staticblock/:key [get]
func (this *ExamplesController) StaticBlock() {

}

// @router /all/:key [get]
func (this *ExamplesController) AllBlock() {

}

// @router /examples/accordion.html [get]
func (this *ExamplesController) Accordion() {
	this.TplName = "examples/accordion.html"
}

// @router /examples/actionsheet-plus.html [get]
func (this *ExamplesController) Actionsheetplus() {
	this.TplName = "examples/actionsheet-plus.html"
}

// @router /examples/actionsheet.html [get]
func (this *ExamplesController) Actionsheet() {
	this.TplName = "examples/actionsheet.html"
}

// @router /examples/ad.html [get]
func (this *ExamplesController) Ad() {
	this.TplName = "examples/ad.html"
}

// @router /examples/ajax.html [get]
func (this *ExamplesController) Ajax() {
	this.TplName = "examples/ajax.html"
}

// @router /examples/badges.html [get]
func (this *ExamplesController) Badges() {
	this.TplName = "examples/badges.html"
}

// @router /examples/beecloud.html [get]
func (this *ExamplesController) Beecloud() {
	this.TplName = "examples/beecloud.html"
}

// @router /examples/buttons-with-badges.html [get]
func (this *ExamplesController) Buttonswithbadges() {
	this.TplName = "examples/buttons-with-badges.html"
}

// @router /examples/buttons-with-block.html [get]
func (this *ExamplesController) Buttonswithblock() {
	this.TplName = "examples/buttons-with-block.html"
}

// @router /examples/buttons-with-icons.html [get]
func (this *ExamplesController) Buttonswithicons() {
	this.TplName = "examples/buttons-with-icons.html"
}

// @router /examples/buttons.html [get]
func (this *ExamplesController) Buttons() {
	this.TplName = "examples/buttons.html"
}

// @router /examples/checkbox.html [get]
func (this *ExamplesController) Checkbox() {
	this.TplName = "examples/checkbox.html"
}

// @router /examples/clouddb_wilddog.html [get]
func (this *ExamplesController) Clouddbwilddog() {
	this.TplName = "examples/clouddb_wilddog.html"
}

// @router /examples/date.html [get]
func (this *ExamplesController) Date() {
	this.TplName = "examples/date.html"
}

// @router /examples/dialog.html [get]
func (this *ExamplesController) Dialog() {
	this.TplName = "examples/dialog.html"
}

// @router /examples/dtpicker.html [get]
func (this *ExamplesController) Dtpicker() {
	this.TplName = "examples/dtpicker.html"
}

// @router /examples/echarts.html [get]
func (this *ExamplesController) Echarts() {
	this.TplName = "examples/echarts.html"
}

// @router /examples/grid-default.html [get]
func (this *ExamplesController) Griddefault() {
	this.TplName = "examples/grid-default.html"
}

// @router /examples/grid-pagination.html [get]
func (this *ExamplesController) Gridpagination() {
	this.TplName = "examples/grid-pagination.html"
}

// @router /examples/guide.html [get]
func (this *ExamplesController) Guide() {
	this.TplName = "examples/guide.html"
}

// @router /examples/icons.html [get]
func (this *ExamplesController) Icons() {
	this.TplName = "examples/icons.html"
}

// @router /examples/im-chat.html [get]
func (this *ExamplesController) Imchat() {
	this.TplName = "examples/im-chat.html"
}

// @router /examples/imageviewer.html [get]
func (this *ExamplesController) Imageviewer() {
	this.TplName = "examples/imageviewer.html"
}

// @router /examples/indexed-list-select.html [get]
func (this *ExamplesController) Indexedlistselect() {
	this.TplName = "examples/indexed-list-select.html"
}

// @router /examples/indexed-list.html [get]
func (this *ExamplesController) Indexedlist() {
	this.TplName = "examples/indexed-list.html"
}

// @router /examples/info.html [get]
func (this *ExamplesController) Info() {
	this.TplName = "examples/info.html"
}

// @router /examples/input.html [get]
func (this *ExamplesController) Input() {
	this.TplName = "examples/input.html"
}

// @router /examples/lazyload-image.html [get]
func (this *ExamplesController) Lazyloadimage() {
	this.TplName = "examples/lazyload-image.html"
}

// @router /examples/list-triplex-row.html [get]
func (this *ExamplesController) Listtriplexrow() {
	this.TplName = "examples/list-triplex-row.html"
}

// @router /examples/list-with-input.html [get]
func (this *ExamplesController) Listwithinput() {
	this.TplName = "examples/list-with-input.html"
}

// @router /examples/locker-dom.html [get]
func (this *ExamplesController) Lockerdom() {
	this.TplName = "examples/locker-dom.html"
}

// @router /examples/login.html [get]
func (this *ExamplesController) Login() {
	this.TplName = "examples/login.html"
}

// @router /examples/media-list.html [get]
func (this *ExamplesController) Medialist() {
	this.TplName = "examples/media-list.html"
}

// @router /examples/modals.html [get]
func (this *ExamplesController) Modals() {
	this.TplName = "examples/modals.html"
}

// @router /examples/nav.html [get]
func (this *ExamplesController) Nav() {
	this.TplName = "examples/nav.html"
}

// @router /examples/numbox.html [get]
func (this *ExamplesController) Numbox() {
	this.TplName = "examples/numbox.html"
}

// @router /examples/offcanvas-drag-down.html [get]
func (this *ExamplesController) Offcanvasdragdown() {
	this.TplName = "examples/offcanvas-drag-down.html"
}

// @router /examples/offcanvas-drag-left-plus-main.html [get]
func (this *ExamplesController) Offcanvasdragleftplusmain() {
	this.TplName = "examples/offcanvas-drag-left-plus-main.html"
}

// @router /examples/offcanvas-drag-left-plus-menu.html [get]
func (this *ExamplesController) Offcanvasdragleftplusmenu() {
	this.TplName = "examples/offcanvas-drag-left-plus-menu.html"
}

// @router /examples/offcanvas-drag-left.html [get]
func (this *ExamplesController) Offcanvasdragleft() {
	this.TplName = "examples/offcanvas-drag-left.html"
}

// @router /examples/offcanvas-drag-right-plus-main.html [get]
func (this *ExamplesController) Offcanvasdragrightplusmain() {
	this.TplName = "examples/offcanvas-drag-right-plus-main.html"
}

// @router /examples/offcanvas-drag-right-plus-menu.html [get]
func (this *ExamplesController) Offcanvasdragrightplusmenu() {
	this.TplName = "examples/offcanvas-drag-right-plus-menu.html"
}

// @router /examples/offcanvas-drag-right.html [get]
func (this *ExamplesController) Offcanvasdragright() {
	this.TplName = "examples/offcanvas-drag-right.html"
}

// @router /examples/pagination.html [get]
func (this *ExamplesController) Pagination() {
	this.TplName = "examples/pagination.html"
}

// @router /examples/picker.html [get]
func (this *ExamplesController) Picker() {
	this.TplName = "examples/picker.html"
}

// @router /examples/popovers.html [get]
func (this *ExamplesController) Popovers() {
	this.TplName = "examples/popovers.html"
}

// @router /examples/pullrefresh_main.html [get]
func (this *ExamplesController) Pullrefreshmain() {
	this.TplName = "examples/pullrefresh_main.html"
}

// @router /examples/pullrefresh_sub.html [get]
func (this *ExamplesController) Pullrefresh_sub() {
	this.TplName = "examples/pullrefresh_sub.html"
}

// @router /examples/pullrefresh_with_tab.html [get]
func (this *ExamplesController) Pullrefresh_with_tab() {
	this.TplName = "examples/pullrefresh_with_tab.html"
}

// @router /examples/radio.html [get]
func (this *ExamplesController) Radio() {
	this.TplName = "examples/radio.html"
}

// @router /examples/range.html [get]
func (this *ExamplesController) Range() {
	this.TplName = "examples/range.html"
}

// @router /examples/setting.html [get]
func (this *ExamplesController) Setting() {
	this.TplName = "examples/setting.html"
}

// @router /examples/slider-default.html [get]
func (this *ExamplesController) Sliderdefault() {
	this.TplName = "examples/slider-default.html"
}

// @router /examples/slider-table-default.html [get]
func (this *ExamplesController) Slidertabledefault() {
	this.TplName = "examples/slider-table-default.html"
}

// @router /examples/slider-table-pagination.html [get]
func (this *ExamplesController) Slidertablepagination() {
	this.TplName = "examples/slider-table-pagination.html"
}

// @router /examples/slider-with-title.html [get]
func (this *ExamplesController) Sliderwithtitle() {
	this.TplName = "examples/slider-with-title.html"
}

// @router /examples/switches.html [get]
func (this *ExamplesController) Switches() {
	this.TplName = "examples/switches.html"
}

// @router /examples/tab-webview-main.html [get]
func (this *ExamplesController) Tabwebviewmain() {
	this.TplName = "examples/tab-webview-main.html"
}

// @router /examples/tab-webview-subpage-about.html [get]
func (this *ExamplesController) Tabwebviewsubpageabout() {
	this.TplName = "examples/tab-webview-subpage-about.html"
}

// @router /examples/tab-webview-subpage-chat.html [get]
func (this *ExamplesController) Tabwebviewsubpagechat() {
	this.TplName = "examples/tab-webview-subpage-chat.html"
}

// @router /examples/tab-webview-subpage-contact.html [get]
func (this *ExamplesController) Tabwebviewsubpagecontact() {
	this.TplName = "examples/tab-webview-subpage-contact.html"
}

// @router /examples/tab-webview-subpage-setting.html [get]
func (this *ExamplesController) Tabwebviewsubpagesetting() {
	this.TplName = "examples/tab-webview-subpage-setting.html"
}

// @router /examples/tab-with-segmented-control-vertical.html [get]
func (this *ExamplesController) Tabwithsegmentedcontrolvertical() {
	this.TplName = "examples/tab-with-segmented-control-vertical.html"
}

// @router /examples/tab-with-segmented-control.html [get]
func (this *ExamplesController) Tabwithsegmentedcontrol() {
	this.TplName = "examples/tab-with-segmented-control.html"
}

// @router /examples/tab-with-viewpagerindicator.html [get]
func (this *ExamplesController) Tabwithviewpagerindicator() {
	this.TplName = "examples/tab-with-viewpagerindicator.html"
}

// @router /examples/tabbar-labels-only.html [get]
func (this *ExamplesController) Tabbarlabelsonly() {
	this.TplName = "examples/tabbar-labels-only.html"
}

// @router /examples/tabbar-with-submenus.html [get]
func (this *ExamplesController) Tabbarwithsubmenus() {
	this.TplName = "examples/tabbar-with-submenus.html"
}

// @router /examples/tabbar.html [get]
func (this *ExamplesController) Tabbar() {
	this.TplName = "examples/tabbar.html"
}

// @router /examples/tableviews-with-badges.html [get]
func (this *ExamplesController) Tableviewswithbadges() {
	this.TplName = "examples/tableviews-with-badges.html"
}

// @router /examples/tableviews-with-collapses.html [get]
func (this *ExamplesController) Tableviewswithcollapses() {
	this.TplName = "examples/tableviews-with-collapses.html"
}

// @router /examples/tableviews-with-swipe.html [get]
func (this *ExamplesController) Tableviewswithswipe() {
	this.TplName = "examples/tableviews-with-swipe.html"
}

// @router /examples/tableviews.html [get]
func (this *ExamplesController) Tableviews() {
	this.TplName = "examples/tableviews.html"
}

// @router /examples/template.html [get]
func (this *ExamplesController) Template() {
	this.TplName = "examples/template.html"
}

// @router /examples/typography.html [get]
func (this *ExamplesController) Typography() {
	this.TplName = "examples/typography.html"
}
