package utils

import (
	"encoding/json"
	"testing"
	"time"
)

func TestAddKeyPrefixToMap(t *testing.T) {
	var testNumber = 5
	var timeSum time.Duration
	for i := 0; i < testNumber; i++ {
		start := time.Now()
		map1 := make(map[string]interface{})
		map1["name"] = "aa"
		map1["age"] = 11
		bag := make(map[string]interface{})
		bag["pen"] = "pencil"
		pencil := make(map[string]interface{})
		pencil["p1"] = "p1"
		bag["pencil"] = pencil
		map1["bag"] = bag
		t.Log("Source：", map1)
		//map2 := make(map[string]interface{})
		target := make(map[string]interface{})
		AddKeyPrefixToMap(map1, target, "c_", "bag")
		t.Log("Target：", target)
		t.Log("Time：", time.Since(start))
		timeSum += time.Since(start)
	}
	t.Log("ALL_Time：", timeSum/time.Duration(testNumber))
}

func TestTrimKeyPrefixToMap(t *testing.T) {
	var testNumber = 5
	var timeSum time.Duration
	for i := 0; i < testNumber; i++ {
		start := time.Now()
		map1 := make(map[string]interface{})
		map1["c_name"] = "aa"
		map1["c_age"] = 11
		bag := make(map[string]interface{})
		bag["c_pen"] = "pencil"
		pencil := make(map[string]interface{})
		pencil["c_p1"] = "p1"
		bag["c_pencil"] = pencil
		map1["c_bag"] = bag
		t.Log("Source：", map1)
		//map2 := make(map[string]interface{})
		target := make(map[string]interface{})
		TrimKeyPrefixToMap(map1, target, "c_", "bag")
		t.Log("Target：", target)
		t.Log("Time：", time.Since(start))
		timeSum += time.Since(start)
	}
	t.Log("ALL_Time：", timeSum/time.Duration(testNumber))
}

func TestFoFaPro(t *testing.T) {
	source := make(map[string]interface{})
	target := make(map[string]interface{})
	if err := json.Unmarshal([]byte(data), &source); err != nil {
		t.Fatal(err)
	}
	ignores := []string{"lat", "lon"}
	AddKeyPrefixToMap(source, target, "c_", ignores...)

	t.Log(target)
}

var data = `{
    "_index":"fofapro_service_cert",
    "_type":"_doc",
    "_id":"47.95.0.170:80",
    "_score":1,
    "_source":{
        "asn":{
            "as_number":37963,
            "as_organization":"Hangzhou Alibaba Advertising Co.,Ltd."
        },
        "ban_len":224,
        "banner":"",
        "base_protocol":"tcp",
        "category":[
            "其他基础软件"
        ],
        "company":[
            "阿里巴巴集团"
        ],
        "geoip":{
            "city_name":"",
            "continent_code":"AS",
            "country_code2":"CN",
            "country_code3":"CN",
            "country_name":"China",
            "dma_code":null,
            "latitude":34.7725,
            "location":{
                "lat":34.7725,
                "lon":113.7266
            },
            "longitude":113.7266,
            "postal_code":"",
            "real_region_name":"",
            "region_name":"",
            "timezone":"Asia/Shanghai"
        },
        "ip":"47.95.0.170",
        "ipcnet":"47.95.0.0",
        "is_ipv6":false,
        "lastchecktime":"2022-12-11 12:34:53",
        "lastupdatetime":"2022-12-11 12:34:53",
        "parent_category":[
            "基础软件"
        ],
        "port":80,
        "product":[
            "Alibaba-Tengine"
        ],
        "protocol":"http",
        "rule_tags":[
            {
                "category":"Service",
                "cn_category":"其他基础软件",
                "cn_company":"阿里巴巴集团",
                "cn_parent_category":"基础软件",
                "cn_product":"Alibaba-Tengine",
                "company":"Alibaba Group",
                "level":"3",
                "parent_category":"Support System",
                "product":"Tengine",
                "rule_id":"212",
                "softhard":"2"
            }
        ],
        "time":"2022-12-11T04:34:53.12039980Z",
        "v":5
    }
}`

func TestDeleteAbsentKeys(t *testing.T) {
	top := make(map[string]interface{})
	top["one1"] = 11
	top["one2"] = "aa"
	top["one3"] = 1.11
	top["one4"] = []int{1, 2, 3, 4}
	two := make(map[string]interface{})
	two["tow1"] = 22
	two["tow2"] = "bb"
	two["tow3"] = 2.22
	two["tow4"] = []int{5, 6, 7, 8}
	three := make(map[string]interface{})
	three["three1"] = 33
	three["three2"] = "cc"
	three["three3"] = 3.33
	three["three4"] = []int{5, 6, 7, 8}
	four := make([]map[string]interface{}, 0)
	four = append(four, map[string]interface{}{"four1": 1, "four2": 2})
	four = append(four, map[string]interface{}{"four1": 3, "four2": 4})
	three["three5"] = four
	two["tow5"] = three
	top["one5"] = two
	t.Log(top)
	keys := []string{"one1", "one2", "one4", "one5", "tow1", "tow2", "tow4", "tow5", "three1", "three5", "four1"}
	target := DeleteAbsentElementKeys(top, make(map[string]interface{}), keys)
	t.Log(target)
}

func TestDeleteAbsentMapping(t *testing.T) {
	top := make(map[string]interface{})
	top["one1"] = 11
	top["one2"] = "aa"
	top["one3"] = 1.11
	top["one4"] = []int{1, 2, 3, 4}
	two := make(map[string]interface{})
	two["tow1"] = 22
	two["tow2"] = "bb"
	two["tow3"] = 2.22
	two["tow4"] = []int{5, 6, 7, 8}
	three := make(map[string]interface{})
	three["three1"] = 33
	three["three2"] = "cc"
	three["three3"] = 3.33
	three["three4"] = []int{5, 6, 7, 8}
	four := make([]map[string]interface{}, 0)
	four = append(four, map[string]interface{}{"four1": 1, "four2": 2})
	four = append(four, map[string]interface{}{"four1": 3, "four2": 4})
	three["three5"] = four
	two["tow5"] = three
	top["one5"] = two
	t.Log(top)

	mapping := make(map[string]interface{})
	mapping["one1"] = struct{}{}
	mapping["one2"] = struct{}{}
	mapping["one4"] = struct{}{}
	twocp := make(map[string]interface{})
	twocp["tow1"] = struct{}{}
	twocp["tow2"] = struct{}{}
	two["tow4"] = struct{}{}
	threecp := make(map[string]interface{})
	threecp["three1"] = 33
	fourcp := make(map[string]interface{})
	fourcp["four1"] = struct{}{}
	threecp["three5"] = fourcp
	twocp["tow5"] = threecp
	mapping["one5"] = twocp
	t.Log(mapping)

	target := DeleteAbsentElementByMapping(top, make(map[string]interface{}), mapping)
	t.Log(target)
}

func TestDeleteAbsentMappingFofaPro(t *testing.T) {

}
