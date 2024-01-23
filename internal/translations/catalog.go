// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package translations

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type dictionary struct {
	index []uint32
	data  string
}

func (d *dictionary) Lookup(key string) (data string, ok bool) {
	p, ok := messageKeyToIndex[key]
	if !ok {
		return "", false
	}
	start, end := d.index[p], d.index[p+1]
	if start == end {
		return "", false
	}
	return d.data[start:end], true
}

func init() {
	dict := map[string]catalog.Dictionary{
		"de": &dictionary{index: deIndex, data: deData},
		"en": &dictionary{index: enIndex, data: enData},
		"fr": &dictionary{index: frIndex, data: frData},
	}
	fallback := language.MustParse("en")
	cat, err := catalog.NewFromMap(dict, catalog.Fallback(fallback))
	if err != nil {
		panic(err)
	}
	message.DefaultCatalog = cat
}

var messageKeyToIndex = map[string]int{
	"About":                   0,
	"App":                     3,
	"App Registration":        6,
	"App Settings":            9,
	"Auto-discovered Servers": 13,
	"Fyne":                    4,
	"Fyne Settings":           8,
	"MQTT Password":           19,
	"MQTT Server":             17,
	"MQTT User":               18,
	"Manual Server Entry":     15,
	"Please restart the agent to use changed settings.": 11,
	"Quit":     5,
	"Save":     10,
	"Sensors":  1,
	"Settings": 2,
	"To register the agent, please enter the relevant details for your Home Assistant\nserver (if not auto-detected) and long-lived access token.": 7,
	"Token":              12,
	"Use Custom Server?": 14,
	"Use MQTT?":          16,
}

var deIndex = []uint32{ // 21 elements
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000,
} // Size: 108 bytes

const deData string = ""

var enIndex = []uint32{ // 21 elements
	0x00000000, 0x00000006, 0x0000000e, 0x00000017,
	0x0000001b, 0x00000020, 0x00000025, 0x00000036,
	0x000000c2, 0x000000d0, 0x000000dd, 0x000000e2,
	0x00000114, 0x0000011a, 0x00000132, 0x00000145,
	0x00000159, 0x00000163, 0x0000016f, 0x00000179,
	0x00000187,
} // Size: 108 bytes

const enData string = "" + // Size: 391 bytes
	"\x02About\x02Sensors\x02Settings\x02App\x02Fyne\x02Quit\x02App Registrat" +
	"ion\x02To register the agent, please enter the relevant details for your" +
	" Home Assistant\x0aserver (if not auto-detected) and long-lived access t" +
	"oken.\x02Fyne Settings\x02App Settings\x02Save\x02Please restart the age" +
	"nt to use changed settings.\x02Token\x02Auto-discovered Servers\x02Use C" +
	"ustom Server?\x02Manual Server Entry\x02Use MQTT?\x02MQTT Server\x02MQTT" +
	" User\x02MQTT Password"

var frIndex = []uint32{ // 21 elements
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000,
} // Size: 108 bytes

const frData string = ""

// Total table size 715 bytes (0KiB); checksum: 1D84AA00
