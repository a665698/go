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
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	sync2 "sync"
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

type City struct {
	Id string
	Name string
}

var SqlDb *sql.DB
var ProvinceInfo []City
var CityInfo  []City
var CurrentData = make(map[int]bool)
var ErrData = make(chan string)
var sync = sync2.Mutex{}
var WriteDirectory = make(chan []string)
var MainExit = make(chan bool)
var UpdateBrandId = make(chan []int)

func main() {
	//go chanHandel(fileName)
	//for i := 1; i <= 1635; i++ {
	//	httpGet(i)
	//}
	//readFile()

	WriteData()
	<- MainExit
}

func WriteData()  {

	if err := MysqlInit(); err != nil {
		ErrResult(err)
		return
	}
	//if err := GetSqlCity(); err != nil {
	//	ErrResult(err)
	//	return
	//}
	//
	//if err := GetCurrentDatas(); err != nil {
	//	ErrResult(err)
	//	return
	//}

	go func() {
		defer func() {
			MainExit <- true
		}()
		err := channelHandle()
		if err != nil {
			ErrResult(err)
		}
		return
	}()

	//GetWriteDatas()
	GetBrandId()
}

func CurrentDataHandle(i int) bool {
	sync.Lock()
	defer sync.Unlock()
	if _, ok := CurrentData[i]; ok {
		return true
	}
	CurrentData[i] = true
	return false
}

func MysqlInit() error {
	var err error
	SqlDb, err = sql.Open("mysql", "dev_user:9xujqm@tcp(121.40.52.21:3306)/iwebmall_dev?charset=utf8")
	if err != nil {
		return err
	}
	return nil
}

func GetSqlCity() (err error) {
	rows, err := SqlDb.Query("select id,name,level from imall_citys where level in (1,2)")
	defer rows.Close()
	if err != nil {
		return
	}
	level, id, name := 0, "0", ""
	for rows.Next() {
		rows.Scan(&id, &name, &level)
		if level == 1 {
			ProvinceInfo = append(ProvinceInfo, City{id, name})
		} else if level == 2 {
			CityInfo = append(CityInfo, City{id, name})
		}
	}
	return
}

func GetCurrentDatas() (err error) {
	rows, err := SqlDb.Query("select telphone from imall_uc_directory")
	defer rows.Close()
	if err != nil {
		return
	}
	mobile := 0
	for rows.Next() {
		rows.Scan(&mobile)
		CurrentData[mobile] = true
	}
	return
}

func GetWriteDatas() (error) {
	file, err := os.Open(fileNameOk)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			return err
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		go WriteCurrentDataHandle(line)
	}
	return nil
}

func channelHandle() error {
	errFile, err := os.OpenFile("cat_err.txt", os.O_CREATE|os.O_APPEND, 0666)
	tx, err := SqlDb.Begin()
	if err != nil {
		return err
	}
	//stmt, err := tx.Prepare("insert into imall_uc_directory(company,province,city,district,province_id,city_id,msn,telphone,address,brand) values(?, ?, ?, ?, ?, ?, '4s', ?, ?, ?)")
	stmt, err := tx.Prepare("update imall_uc_directory set brand_id = ? where id = ?")
	defer stmt.Close()
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	for {
		select {
		case info := <- ErrData:
			errFile.WriteString(info + "\n")
		case info := <- WriteDirectory:
			stmt.Exec(info[0], info[3], info[4], info[5], info[7], info[8], info[2], info[6], info[1])
		case brandId := <- UpdateBrandId:
			fmt.Println(brandId)
			stmt.Exec(brandId[1], brandId[0])
		case <- time.After(10 * time.Second):
			tx.Commit()
			return nil
		}
	}
	return nil
}

func GetBrandId() error {
	rows, err := SqlDb.Query("SELECT UD.id,CC.cat_id from imall_uc_directory as UD left JOIN imall_car_category as CC on UD.brand=CC.cat_name where UD.brand_id=0 and CC.parent_id=0")
	defer rows.Close()
	if err != nil {
		return err
	}
	id, carId := 0, 0
	for rows.Next() {
		rows.Scan(&id, &carId)
		UpdateBrandId <- []int{id, carId}
	}
	return nil
}

func WriteCurrentDataHandle(info string) {
	infoArr := strings.Split(info, ",")
	if len(infoArr) == 7 {
		mobile, err := strconv.Atoi(infoArr[2])
		if err != nil {
			ErrData <- info
			return
		}
		if CurrentDataHandle(mobile) {
			ErrData <- info
			return
		}
		infoArr = append(infoArr, GetProvinceId(infoArr[3]), GetCityId(infoArr[4]))
		WriteDirectory <- infoArr
	} else {
		ErrData <- info
	}
}

func GetProvinceId(province string) string {
	province = strings.TrimSpace(province)
	if province == "" {
		return "0"
	}
	for _, v := range ProvinceInfo {
		if strings.Contains(province, v.Name) {
			return v.Id
		}
	}
	return "0"
}

func GetCityId(city string) string {
	city = strings.TrimSpace(city)
	if city == "" {
		return "0"
	}
	for _, v := range CityInfo {
		if strings.Contains(city, v.Name) {
			return v.Id
		}
	}
	return "0"
}



func ErrResult(err error) {
	fmt.Println(err)
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



