package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["medexam/controllers:CMSController"] = append(beego.GlobalControllerRouter["medexam/controllers:CMSController"],
		beego.ControllerComments{
			"StaticBlock",
			`/staticblock/:key`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:CMSController"] = append(beego.GlobalControllerRouter["medexam/controllers:CMSController"],
		beego.ControllerComments{
			"AllBlock",
			`/all/:key`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:CMSController"] = append(beego.GlobalControllerRouter["medexam/controllers:CMSController"],
		beego.ControllerComments{
			"Info",
			`/info.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:CMSController"] = append(beego.GlobalControllerRouter["medexam/controllers:CMSController"],
		beego.ControllerComments{
			"List",
			`/list.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:CMSController"] = append(beego.GlobalControllerRouter["medexam/controllers:CMSController"],
		beego.ControllerComments{
			"Gitpull",
			`/gitpull`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:CMSController"] = append(beego.GlobalControllerRouter["medexam/controllers:CMSController"],
		beego.ControllerComments{
			"Shell",
			`/shell`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"StaticBlock",
			`/staticblock/:key`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"AllBlock",
			`/all/:key`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Accordion",
			`/examples/accordion.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Actionsheetplus",
			`/examples/actionsheet-plus.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Actionsheet",
			`/examples/actionsheet.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Ad",
			`/examples/ad.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Ajax",
			`/examples/ajax.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Badges",
			`/examples/badges.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Beecloud",
			`/examples/beecloud.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Buttonswithbadges",
			`/examples/buttons-with-badges.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Buttonswithblock",
			`/examples/buttons-with-block.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Buttonswithicons",
			`/examples/buttons-with-icons.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Buttons",
			`/examples/buttons.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Checkbox",
			`/examples/checkbox.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Clouddbwilddog",
			`/examples/clouddb_wilddog.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Date",
			`/examples/date.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Dialog",
			`/examples/dialog.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Dtpicker",
			`/examples/dtpicker.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Echarts",
			`/examples/echarts.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Griddefault",
			`/examples/grid-default.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Gridpagination",
			`/examples/grid-pagination.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Guide",
			`/examples/guide.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Icons",
			`/examples/icons.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Imchat",
			`/examples/im-chat.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Imageviewer",
			`/examples/imageviewer.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Indexedlistselect",
			`/examples/indexed-list-select.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Indexedlist",
			`/examples/indexed-list.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Info",
			`/examples/info.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Input",
			`/examples/input.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Lazyloadimage",
			`/examples/lazyload-image.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Listtriplexrow",
			`/examples/list-triplex-row.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Listwithinput",
			`/examples/list-with-input.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Lockerdom",
			`/examples/locker-dom.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Login",
			`/examples/login.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Medialist",
			`/examples/media-list.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Modals",
			`/examples/modals.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Nav",
			`/examples/nav.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Numbox",
			`/examples/numbox.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Offcanvasdragdown",
			`/examples/offcanvas-drag-down.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Offcanvasdragleftplusmain",
			`/examples/offcanvas-drag-left-plus-main.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Offcanvasdragleftplusmenu",
			`/examples/offcanvas-drag-left-plus-menu.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Offcanvasdragleft",
			`/examples/offcanvas-drag-left.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Offcanvasdragrightplusmain",
			`/examples/offcanvas-drag-right-plus-main.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Offcanvasdragrightplusmenu",
			`/examples/offcanvas-drag-right-plus-menu.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Offcanvasdragright",
			`/examples/offcanvas-drag-right.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Pagination",
			`/examples/pagination.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Picker",
			`/examples/picker.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Popovers",
			`/examples/popovers.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Pullrefreshmain",
			`/examples/pullrefresh_main.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Pullrefresh_sub",
			`/examples/pullrefresh_sub.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Pullrefresh_with_tab",
			`/examples/pullrefresh_with_tab.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Radio",
			`/examples/radio.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Range",
			`/examples/range.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Setting",
			`/examples/setting.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Sliderdefault",
			`/examples/slider-default.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Slidertabledefault",
			`/examples/slider-table-default.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Slidertablepagination",
			`/examples/slider-table-pagination.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Sliderwithtitle",
			`/examples/slider-with-title.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Switches",
			`/examples/switches.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Tabwebviewmain",
			`/examples/tab-webview-main.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Tabwebviewsubpageabout",
			`/examples/tab-webview-subpage-about.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Tabwebviewsubpagechat",
			`/examples/tab-webview-subpage-chat.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Tabwebviewsubpagecontact",
			`/examples/tab-webview-subpage-contact.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Tabwebviewsubpagesetting",
			`/examples/tab-webview-subpage-setting.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Tabwithsegmentedcontrolvertical",
			`/examples/tab-with-segmented-control-vertical.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Tabwithsegmentedcontrol",
			`/examples/tab-with-segmented-control.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Tabwithviewpagerindicator",
			`/examples/tab-with-viewpagerindicator.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Tabbarlabelsonly",
			`/examples/tabbar-labels-only.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Tabbarwithsubmenus",
			`/examples/tabbar-with-submenus.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Tabbar",
			`/examples/tabbar.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Tableviewswithbadges",
			`/examples/tableviews-with-badges.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Tableviewswithcollapses",
			`/examples/tableviews-with-collapses.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Tableviewswithswipe",
			`/examples/tableviews-with-swipe.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Tableviews",
			`/examples/tableviews.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Template",
			`/examples/template.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:ExamplesController"] = append(beego.GlobalControllerRouter["medexam/controllers:ExamplesController"],
		beego.ControllerComments{
			"Typography",
			`/examples/typography.html`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["medexam/controllers:MainController"] = append(beego.GlobalControllerRouter["medexam/controllers:MainController"],
		beego.ControllerComments{
			"StaticBlock",
			`/staticblock/:key`,
			[]string{"get"},
			nil})

}
