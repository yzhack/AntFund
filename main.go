package main

import (
	"AntFund/easyjson"
	"AntFund/fetcher"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/cookiejar"
	"os"
	"regexp"
	"strconv"
	"time"
)

var jsonRe = regexp.MustCompile(`{(.*?)};`)

type InitInfo struct {
	Fundname         string // 基金名称
	Csrf             string // csrftoken
	ProudctId        string // 产品ID
	NetValue         string // 净值
	NetValueDate     string // 净值日期
	DayOfGrowth      string // 前一日涨幅
	ForecastNetValue string // 估算净值
	ForecastGrowth   string // 估算涨幅
}

func main() {
	fmt.Println("基金名称 | 基金代码 | 单位净值 | 日涨幅 | 单位净值日期 | 估算净值 | 估算涨幅")
	fileName := "./conf/conf.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		code := scanner.Text()
		fetcher.Cookie = nil
		fetcher.Cookiejar, _ = cookiejar.New(nil)
		var info InitInfo
		info = Init(code)

		info, err := DatePost(info)
		if err != nil {
			fmt.Println(err)
		}
		dayOfGrowth := floatToStr(info.DayOfGrowth, 2, "") + "%"
		forecastNetValue := floatToStr(info.ForecastNetValue, 4, "")
		forecastGrowth := floatToStr(info.ForecastGrowth, 2, "a") + "%"
		fmt.Printf("%s | %s | %s | %s | %s | %s | %s\n", info.Fundname, code, info.NetValue, dayOfGrowth, info.NetValueDate, forecastNetValue, forecastGrowth)

	}

	fmt.Scanf("%s")

}

func Init(code string) InitInfo {
	url := "http://www.fund123.cn/matiaria?fundCode=" + code
	bytedata, err := fetcher.Fetch(url, "GET", nil)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(bytedata))
	//metches:=jsonRe.FindAll(bytedata,-1)
	//fmt.Println(string(metches[0][25:len(metches[0])-10]))
	metches := jsonRe.FindAllString(string(bytedata), 1)
	//fmt.Println()
	var datas easyjson.JsonInfo
	json.Unmarshal([]byte(metches[0][:len(metches[0])-1]), &datas)
	//fmt.Println(datas)
	info := InitInfo{}
	info.Fundname = datas.MaterialInfo.FundBrief.FundNameAbbr
	info.Csrf = datas.Csrf
	info.ProudctId = datas.MaterialInfo.ProductID
	info.NetValue = datas.MaterialInfo.TitleInfo.NetValue
	info.NetValueDate = datas.MaterialInfo.TitleInfo.NetValueDate
	info.DayOfGrowth = datas.MaterialInfo.TitleInfo.DayOfGrowth

	return info
}

func DatePost(info InitInfo) (InitInfo, error) {

	//fmt.Println(info)
	url := "http://www.fund123.cn/api/fund/queryFundEstimateIntraday?_csrf=" + info.Csrf
	data := make(map[string]interface{})
	now := time.Now()
	startTime := now.Format("2006-01-02")
	endtTime := now.AddDate(0, 0, 1).Format("2006-01-02")
	data["startTime"] = startTime
	data["endTime"] = endtTime
	data["limit"] = 200
	data["productId"] = info.ProudctId
	data["format"] = "true"
	data["source"] = "WEALTHBFFWEB"
	//fmt.Println(data)
	datas, _ := json.Marshal(data)
	//fmt.Println(string(datas))
	byte_data, err := fetcher.Fetch(url, "POST", bytes.NewBuffer(datas))
	if err != nil {
		fmt.Println(err)
		return InitInfo{}, err
	}
	//fmt.Println(string(byte_data))
	var result easyjson.Datas
	err = json.Unmarshal(byte_data, &result)
	if err != nil {
		fmt.Println(err)
		return InitInfo{}, err
	}
	//fmt.Println("开始显示当前基金图")
	value := result.List[len(result.List)-1]
	info.ForecastGrowth = value.ForecastGrowth
	info.ForecastNetValue = value.ForecastNetValue
	//for _,v:=range result.List{
	//	fmt.Println(v.BizSeq,v.Time,v.ForecastNetValue,v.ForecastGrowth)
	//}
	//fmt.Println(value)
	return info, nil

}

// float字符串转换成你想要的float字符串
func floatToStr(str string, prec int, zhuangtai string) string {
	strToflost, _ := strconv.ParseFloat(str, 64)
	if zhuangtai != "" {
		strToflost = strToflost * 100
	}
	floststr := strconv.FormatFloat(strToflost, 'f', prec, 64)
	return floststr
}
