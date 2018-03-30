package main

import (
	"net/http"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"os"
	"io/ioutil"
	"encoding/json"
	"time"
	"bufio"
	"io"
	"net/url"
	"errors"
)

const (
	// 文本文件
	fileName = "car.txt"
	fileNameOk = "car_ok.txt"
)

var Api = []string{
	"http://apis.map.qq.com/ws/geocoder/v1/?key=HQYBZ-CUHCX-4KZ42-ZQP2N-5KCST-HFBW7&address=", // 腾讯地图API
	"http://restapi.amap.com/v3/geocode/geo?key=fec2ef2f65b72b39ddb73f4a02edee4e&address=", // 高德地图API
}
var currentApiKey = 0
var channel = make(chan *CarInfo)


type CarInfo struct {
	name string
	brand string
	mobile string
	address string
}

type cityInfo struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Result struct{
		AddressComponents struct{
			Province string `json:"province"`
			City string `json:"city"`
			District string `json:"district"`
		} `json:"address_components"`
	} `json:"result"`
}

type GdApiInfo struct {
	Status string `json:"status"`
	Count string `json:"count"`
	Infocode string `json:"infocode"`
	Geocodes []GdApiGeocode `json:"geocodes"`
}

type GdApiGeocode struct {
	Province string `json:"province"`
	City interface{} `json:"city"`
	District interface{} `json:"district"`
}


func main() {
	//go chanHandel(fileName)
	//for i := 1; i <= 1635; i++ {
	//	httpGet(i)
	//}
	readFile()
}

func httpGet(index int) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://dealer.autohome.com.cn/china?countyId=0&brandId=0&seriesId=0&factoryId=0&pageIndex=%d&kindId=1&orderType=0&isSales=0", index), nil)
	if err != nil {
		fmt.Println("request build err", err)
		return
	}
	req.Header.Set("Content-Type", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7")

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("url get err", err)
		return
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		fmt.Println("pageIndex url status err", response.StatusCode)
		return
	}
	h, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		fmt.Println("build html err", err)
		return
	}
	parseHtml(h)
	fmt.Println("第", index, "页完成")
}


func parseHtml(document *goquery.Document) {
	li := document.Find(".list-box .list-item")
	li.Each(func(i int, content *goquery.Selection) {
		info := content.Find("li")
		is4s := info.First().Find(".green").Text()
		if is4s == "4S店" {
			carInfo := &CarInfo{}
			carInfo.name = info.First().Find("a span").Text()
			carInfo.brand = info.Eq(1).Find("span em").Text()
			carInfo.mobile = info.Eq(2).Find("span.tel").Text()
			carInfo.address = info.Eq(3).Find("span.info-addr").Text()
			channel <- carInfo
		}
	})
}

func getCity(address string) (string, error) {
	//index := strings.IndexRune(address, rune('（'))
	//if index > 0 {
	//	address = strings.Replace(address, "（", "", -1)
	//	address = strings.Replace(address, "）", "", -1)
	//}
	client := &http.Client{}
	u := Api[currentApiKey] + url.QueryEscape(address)
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7")
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	if body, err := ioutil.ReadAll(response.Body); err != nil {
		return "", err
	} else {
		if currentApiKey == 0 {
			return TxApiHandle(body, address)
		} else {
			return GdApiHandle(body, address)
		}
	}

}

func TxApiHandle(body []byte, address string) (string, error) {
	city := &cityInfo{}
	if err := json.Unmarshal(body, city); err != nil {
		return "", err
	} else {
		if city.Status == 0 {
			return fmt.Sprintf("%s,%s,%s,%s",
				city.Result.AddressComponents.Province,city.Result.AddressComponents.City, city.Result.AddressComponents.District, address), nil
		} else if city.Status == 121 {
			if currentApiKey < len(Api) {
				currentApiKey ++
				return getCity(address)
			} else {
				return "",errors.New("API上限")
			}
		} else if city.Status == 120 {
			time.Sleep(1 * time.Second)
			return getCity(address)
		} else {
			return address, errors.New("next")
		}
	}
}

func GdApiHandle(body []byte, address string) (string, error) {
	city := &GdApiInfo{}
	if err := json.Unmarshal(body, city); err != nil {
		return "", err
	} else {
		if city.Infocode == "10000" && city.Count != "0" {
			_city := ""
			if v, ok := city.Geocodes[0].City.(string); ok {
				_city = v
			}
			district := ""
			if v, ok := city.Geocodes[0].District.(string); ok {
				district = v
			}
			return fmt.Sprintf("%s,%s,%s,%s",
				city.Geocodes[0].Province, _city, district, strings.TrimSuffix(address, "\n")), nil
		} else if city.Infocode == "10003" {
			if currentApiKey < len(Api) {
				currentApiKey ++
				return getCity(address)
			} else {
				return "", errors.New("API上限")
			}
		} else if city.Infocode == "10004" {
			time.Sleep(1 * time.Second)
			return getCity(address)
		} else {
			return address, errors.New("next")
		}
	}
}

func readFile() {
	go chanHandel(fileNameOk)
	file, err := os.OpenFile(fileName, os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("file open err", err)
	}
	buf := bufio.NewReader(file)
	afterText := ""
	isHandle := true
	i := 0
	for {
		i ++
		line, err := buf.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		lineT := strings.TrimSuffix(line, "\n")
		if isHandle {
			carInfo := &CarInfo{}
			lineS := strings.Split(lineT, ",")
			carInfo.name = lineS[0]
			carInfo.brand = lineS[1]
			carInfo.mobile = lineS[2]
			if address, err := getCity(lineS[3]); err == nil {
				carInfo.address = address
				channel <- carInfo
				fmt.Println("第", i, "条完成")
			} else {
				fmt.Println(err)
				if address == "" {
					isHandle = false
					fmt.Println("错误,执行退出")
				}
				afterText = afterText + line
			}
		} else {
			afterText = afterText + line
		}
	}
	os.Truncate(fileName, 0)
	file.Write([]byte(afterText))
}

func chanHandel(name string)  {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("file open err", err)
	}
	for {
		select {
		case info := <- channel:
			info.mobile = strings.Replace(info.mobile, "-", "", -1)
			s := info.name + "," + info.brand + "," + info.mobile + "," + info.address + "\r\n"
			file.Write([]byte(s))
		}
	}
}



