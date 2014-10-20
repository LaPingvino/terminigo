package models

import (
	data "github.com/komputeko/komputeko-data"
	"github.com/astaxie/beego"
	"encoding/json"
	"os"
	"time"
	"strings"
)

var theTerminaro data.Terminaro
var t chan data.Terminaro
var entryCache map[[2]string]data.Terminaro

func init() {
	t = make(chan data.Terminaro,0)
	terminaro := make(data.Terminaro,0)
	lastModTime := time.Now()
	go func() {
		for {
			file, err := os.Open(beego.AppConfig.String("terminaro"))
			if err != nil {
				beego.Error(err.Error() + ", probably json not correctly invoked, current terminaro conf entry is " + beego.AppConfig.String("terminaro"))
			}
			stat, err := file.Stat()
			if err != nil {
				beego.Error(err.Error())
				continue
			}
			if stat.ModTime() != lastModTime {
				err = json.NewDecoder(file).Decode(&terminaro)
				if err == nil {
					t <- terminaro
					lastModTime = stat.ModTime()
				} else {
					beego.Error(err.Error())
				}
			}
			time.Sleep(1 * time.Minute)
		}
	}()
}

func GetTerminaro() (t2 data.Terminaro) {
	if theTerminaro == nil {
		theTerminaro = make(data.Terminaro,0)
	}
	select {
	case theTerminaro = <-t:
		entryCache = nil
	default:
	}
	t2 = make(data.Terminaro, len(theTerminaro))
	copy(t2, theTerminaro)
	return t2
}

func GetEntries(lang string, target string) (re data.Terminaro) {
	re = entryCache[[2]string{lang,target,}]
	if len(re) != 0 {
		return re
	}
	terminaro := GetTerminaro()
	if entryCache == nil { entryCache = make(map[[2]string]data.Terminaro) }
	if len(terminaro) > 1 {
		targetfound := false
		for _, entry := range terminaro[1:] {
			for _, translation := range entry.Translations {
				if translation.Language == lang || lang == "" {
					for _, word := range translation.Words {
						if len(word.Written) < len(target) {
							continue
						}
						if strings.ToLower(word.Written[:len(target)]) == strings.ToLower(target) {
							targetfound = true
						}
					}
				}
			}
		if targetfound { re = append(re, entry) }
		}
	} else { beego.Error("Empty terminaro") }
	re.SortLang(lang, target)
	entryCache[[2]string{lang,target,}] = re
	return re
}

func GetLangs() map[string]string {
	return map[string]string{"en": "Angla", "eo": "Esperanto", "nl": "Nederlanda", "de": "Germana", "fr": "Franca",}
}

