package controllers

import (
	"github.com/astaxie/beego"
	"github.com/komputeko/terminigo/models"
	data "github.com/komputeko/komputeko-data"
	templ "html/template"
	"strings"
)

type MainController struct {
	beego.Controller
}

type ShowPage struct {
	beego.Controller
}

type SearchWord struct {
	beego.Controller
}

func htmlTerminaro (terminaro data.Terminaro, slang, wlang, prefix string) (html templ.HTML) {
	var header string
	html += templ.HTML("<dl>\n")
	for _, entry := range terminaro {
		for _, tr := range entry.Translations {
			if tr.Language == wlang || wlang == "" {
				for _, wr := range tr.Words {
					if len(wr.Written) < len(prefix) { continue }
					if strings.ToLower(wr.Written[:len(prefix)]) != strings.ToLower(prefix) { continue }
					header = wr.Written
				}
			}
		}
		html += templ.HTML("<dt>" + header + "</dt>\n<dd>\n")
		for _, tr := range entry.Translations {
			html += templ.HTML("<div class=\"word\"><div id=\"" + tr.Language + "\" class=\"wordheader\">" + tr.Language + "</div>\n<div class=\"definitions\">")
			for _, word := range tr.Words {
				html += templ.HTML("<div class=\"definition\"><a href=\"/" +
					templ.HTMLEscapeString(slang) + "/" +
					templ.HTMLEscapeString(word.Written) + "/" +
					templ.HTMLEscapeString(tr.Language) + "/\">" +
					templ.HTMLEscapeString(word.Written) + "</a> ")
				for _, sw := range word.Sources {
					html += templ.HTML("<span class=\"source\">" +
						templ.HTMLEscapeString(sw) + "</span>")
				}
	
				html += templ.HTML("</div>")
			}
			html += templ.HTML("</div></div>")
		}
		html += templ.HTML("</dd>")
	}
	html += templ.HTML("</dl>")
	return
}

func (this *MainController) Get() {
	this.Data["Lang"] = this.Ctx.Input.Param(":slang")
	if this.Data["Lang"] == "" { this.Data["Lang"] = "eo" }
	this.TplNames = "index.tpl"
	this.Data["Topbar"] = "Bonvenon ĉe komputeko!"
}

func (this *ShowPage) Get() {
	slang := this.Ctx.Input.Param(":slang")
	if slang == "" { slang = "eo" }
	this.Data["Lang"] = slang
	this.TplNames = "index.tpl"
	lang := this.Ctx.Input.Param(":wlang")
	word := this.Ctx.Input.Param(":word")
	result := models.GetEntries(lang, word)
	this.Data["Pagecontent"] = htmlTerminaro(result, slang, lang, word)
	this.Data["Topbar"] = "Rezultoj por " + word + " en " + lang
}

func (this *SearchWord) Get() {
	slang := this.Ctx.Input.Param(":slang")
	if slang == "" { slang = "eo" }
	this.Data["Lang"] = slang
	this.TplNames = "index.tpl"
	word := this.GetString("vorto")
	result := models.GetEntries("", word)
	this.Data["Pagecontent"] = htmlTerminaro(result, slang, "", word)
	this.Data["Topbar"] = "Rezultoj por " + word
}