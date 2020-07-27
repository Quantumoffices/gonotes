package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gocolly/colly/v2"
	"log"
	"strconv"
	"strings"
)

//证监会行业市盈率

type ZhjhHyShyl struct {
	Hydm string `json:"行业代码"`

	Hymc string `json:"行业名称"`

	Zxsj *float64 `json:"最新数据"`

	Gpjs int `json:"股票家数"`

	Ksjs int `json:"亏损家数"`

	Jygy *float64 `json:"近一个月"`

	Jsgy *float64 `json:"近三个月"`

	Jlgy *float64 `json:"近六个月"`

	Jyn *float64 `json:"近一年"`

	Zhy []*ZhjhHyShyl `json:"细分行业"`
}

func main() {

	var err error
	c := colly.NewCollector()
	//设置浏览器
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36 Edge/16.16299"

	//zjhHyShyl := make([]*ZhjhHyShyl, 0)
	//citys := map[string][]string{}
	//c.OnRequest(func(r *colly.Request) {
	//	fmt.Printf("%+v\r\n%+v\r\n", *r, *(r.Headers))
	//})

	//province := ""
	// Find and visit all links
	provinces := make(map[string]string)
	c.OnHTML("div.menufix>p:first-child>a", func(e *colly.HTMLElement) {
		//fmt.Println(e.Attr("href"))
		if e.Text != "全部" {
			provinces[e.Text] = e.Attr("href")
			fmt.Println(e.Text, e.Attr("href"))
			oneItem(e.Text, e.Attr("href"))
		}
	})

	//抓取完成后
	c.OnScraped(func(req *colly.Response) {
		for k, v := range provinces {
			fmt.Println(k, v)
			req.Request.Visit(v)
		}
	})

	err = c.Visit("http://college.gaokao.com/schlist/")

	if err != nil {

		log.Fatal(err)

	}
	oneItem("北京", "http://college.gaokao.com/schlist/a1/")
}

func oneItem(name string, url string) {
	var err error

	c := colly.NewCollector()

	//设置浏览器
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36 Edge/16.16299"

	//zjhHyShyl := make([]*ZhjhHyShyl, 0)
	//citys := map[string][]string{}
	//c.OnRequest(func(r *colly.Request) {
	//	fmt.Printf("%+v\r\n%+v\r\n", *r, *(r.Headers))
	//})

	//province := ""
	// Find and visit all links
	//provinces := make(map[string]string)
	//c.OnHTML("div.menufix>p:first-child>a", func(e *colly.HTMLElement) {
	//	//fmt.Println(e.Attr("href"))
	//	if e.Text != "全部" {
	//		provinces[e.Text] = e.Attr("href")
	//	}
	//})
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet1")
	f.SetCellValue("Sheet1", "A1", "学校名称")
	f.SetCellValue("Sheet1", "B1", "高校性质")
	f.SetCellValue("Sheet1", "C1", "学校网址")
	f.SetCellValue("Sheet1", "D1", "高校隶属")
	f.SetCellValue("Sheet1", "E1", "高校类型")
	f.SetCellValue("Sheet1", "F1", "院校特色")
	f.SetCellValue("Sheet1", "G1", "logo")
	//for i := 0; i < len(array); i++ {
	//	indexStr := strconv.Itoa(i + 2)
	//	// Set value of a cell.
	//	f.SetCellValue("Sheet1", "A"+indexStr, array[i])
	//	//f.SetCellValue("Sheet1", "B"+indexStr, provices[i].Name)
	//	//time.Sleep(time.Millisecond * 500)
	//}
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)

	rowIndex := 1
	c.OnHTML("div.scores_List>dl", func(e *colly.HTMLElement) {
		rowIndex++
		indexStr := strconv.Itoa(rowIndex)
		e.ForEach("dt>strong", func(i int, element *colly.HTMLElement) {
			sufix := "学院"
			if strings.HasSuffix(element.Text, "大学") {
				sufix = ""
			}
			//fmt.Println("学校名称：", element.Text+sufix)
			f.SetCellValue("Sheet1", "A"+indexStr, element.Text+sufix)
		})
		//校标
		e.ForEach("dt>a>img", func(i int, element *colly.HTMLElement) {
			//fmt.Println("ico：", element.Attr("src"))
			f.SetCellValue("Sheet1", "G"+indexStr, element.Attr("src"))
		})

		e.ForEach("dd>ul", func(i int, element *colly.HTMLElement) {
			e.ForEach("li", func(i int, element *colly.HTMLElement) {
				if strings.HasPrefix(element.Text, "高校性质") {
					f.SetCellValue("Sheet1", "B"+indexStr, strings.TrimSpace(strings.Split(element.Text, "：")[1]))
				}
				if strings.HasPrefix(element.Text, "学校网址") {
					f.SetCellValue("Sheet1", "C"+indexStr, strings.TrimSpace(strings.Split(element.Text, "：")[1]))
				}
				//f.SetCellValue("Sheet1", "D"+indexStr, "高校隶属")
				if strings.HasPrefix(element.Text, "高校隶属") {
					f.SetCellValue("Sheet1", "D"+indexStr, strings.TrimSpace(strings.Split(element.Text, "：")[1]))
				}
				//f.SetCellValue("Sheet1", "E"+indexStr, "高校类型")
				if strings.HasPrefix(element.Text, "高校类型") {
					f.SetCellValue("Sheet1", "E"+indexStr, strings.TrimSpace(strings.Split(element.Text, "：")[1]))
				}
				//f.SetCellValue("Sheet1", "F"+indexStr, "院校特色")
				if strings.HasPrefix(element.Text, "院校特色") {
					fmt.Println(strings.Split(element.Text, "："))
					strs := strings.Split(element.Text, "：")
					fmt.Println(strs[1])
					f.SetCellValue("Sheet1", "F"+indexStr, strings.TrimSpace(strings.Split(element.Text, "：")[1]))
				}
				//if strings.HasPrefix(element.Text, "学校网址") {
				//	f.SetCellValue("Sheet1", "C"+indexStr, strings.TrimSpace(strings.Split(element.Text, ":")[1]))
				//}
			})
		})
	})

	c.OnHTML(".fany>li:nth-child(3)>a", func(eli *colly.HTMLElement) {
		url := eli.Attr("href")
		if len(url) > 5 {
			eli.Request.Visit(url)
		}
		fmt.Println(eli.Text, eli.Attr("href"))
	})

	//抓取完成后
	c.OnScraped(func(req *colly.Response) {
		// Save xlsx file by the given path.
		if err := f.SaveAs(fmt.Sprintf("./colly/school/%s.xlsx", name)); err != nil {
			fmt.Println(err)
		}
	})
	err = c.Visit(url)
	if err != nil {
		panic(err)
	}
}

//爬取省份
func scrapeProvince() {

	var err error

	c := colly.NewCollector()

	//设置浏览器
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36 Edge/16.16299"

	//zjhHyShyl := make([]*ZhjhHyShyl, 0)
	citys := map[string][]string{}
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("%+v\r\n%+v\r\n", *r, *(r.Headers))
	})

	province := ""
	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		//e.Request.Visit(e.Attr("href"))
		if len(e.Attr("name")) > 0 {
			fmt.Println(e.Attr("name"))
			province = e.Text
			return
		}
		citys[province] = append(citys[province], e.Text)
	})

	c.OnScraped(func(_ *colly.Response) {

		//count := 0
		for k, array := range citys {
			//count++
			//if count > 1000 {
			//	break
			//}
			if len(k) > 1 {
				f := excelize.NewFile()
				// Create a new sheet.
				index := f.NewSheet("Sheet1")
				f.SetCellValue("Sheet1", "A1", "城市名")
				for i := 0; i < len(array); i++ {
					indexStr := strconv.Itoa(i + 2)
					// Set value of a cell.
					f.SetCellValue("Sheet1", "A"+indexStr, array[i])
					//f.SetCellValue("Sheet1", "B"+indexStr, provices[i].Name)
					//time.Sleep(time.Millisecond * 500)
				}
				// Set active sheet of the workbook.
				f.SetActiveSheet(index)
				// Save xlsx file by the given path.
				if err := f.SaveAs(fmt.Sprintf("./colly/%s.xlsx", k)); err != nil {
					fmt.Println(err)
				}
			}

		}

	})

	err = c.Visit("http://www.maps7.com/china_province.php")

	if err != nil {

		log.Fatal(err)

	}
}
