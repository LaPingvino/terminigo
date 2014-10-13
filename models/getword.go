package models

import (
	data "github.com/komputeko/komputeko-data"
	"github.com/astaxie/beego"
	"encoding/json"
	"os"
	"sync"
	"time"
	"strings"
)

type SafeTerminaro struct {
	T *data.Terminaro
	lastUpdated time.Time
	sync.Mutex
}

var theTerminaro SafeTerminaro
var entryCache map[[2]string]data.Terminaro

func initData() (terminaro *data.Terminaro, err error) {
	var t data.Terminaro
	terminaro = &t
	file, err := os.Open(beego.AppConfig.String("terminaro"))
	if err != nil { panic(err.Error() + ", probably json not correctly invoked, current terminaro conf entry is " + beego.AppConfig.String("terminaro")) }
	err = json.NewDecoder(file).Decode(terminaro)
	return
}

func (t *SafeTerminaro) Get() *SafeTerminaro {
	if time.Now().Unix() > t.lastUpdated.Add(10 * time.Minute).Unix() {
		terminaro, err := initData()
		if err != nil { panic(err.Error()) }
		entryCache = nil
		t.Lock()
		t.T = terminaro
		t.lastUpdated = time.Now()
		t.Unlock()
	}
	return t
}

func GetEntries(lang string, target string) (re data.Terminaro) {
	re = entryCache[[2]string{lang,target,}]
	if len(re) != 0 {
		return re
	}
	theTerminaro.Get().Lock()
	defer theTerminaro.Get().Unlock()
	terminaro := *theTerminaro.Get().T
	if entryCache == nil { entryCache = make(map[[2]string]data.Terminaro) }
	for _, entry := range terminaro[1:] {
		for _, translation := range entry.Translations {
			if translation.Language == lang || lang == "" {
				for _, word := range translation.Words {
					if len(word.Written) < len(target) { continue }
					if strings.ToLower(word.Written[:len(target)]) == strings.ToLower(target) {
						re = append(re, entry)
					}
				}
			}
		}
	}
	re.SortBy(lang, target)
	entryCache[[2]string{lang,target,}] = re
	return
}

