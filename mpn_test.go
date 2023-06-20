package main

import (
	"bufio"
	"bytes"
	"context"
	"dockerGo/db"
	"dockerGo/helper"
	"dockerGo/model"
	"encoding/json"
	"fmt"
	api2captcha "github.com/2captcha/2captcha-go"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/devices"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/redis/go-redis/v9"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"
)

func TestRedisGet(t *testing.T) {
	rd := db.NewRedis()
	ctx := context.Background()

	rdb := rd.Conn()

	result, err := rdb.Get(ctx, "dana").Result()
	if err == redis.Nil {
		fmt.Println("ga ada")
	} else if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func TestRedisInsert(t *testing.T) {
	rd := db.NewRedis()
	ctx := context.Background()

	rdb := rd.Conn()

	if err := rdb.Set(ctx, "token_ntpn_single", "PHPSESSID=c1u3s19hln2v79uq7f3r4niib1", 0).Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("berhasil simpan")
}
func TestRedisTime(t *testing.T) {

	rd := db.NewRedis()
	ctx := context.Background()

	rdb := rd.Conn()

	result, err := rdb.Get(ctx, "token_ntpn_bulk").Result()

	if err == redis.Nil {
		log.Println(err)
	} else if err != nil {
		log.Println(err)
	} else {
		timestamp, err := strconv.ParseInt(result, 10, 64)
		if err != nil {
			fmt.Println("Gagal mengonversi waktu:", err)
		} else {
			goTime := time.Unix(timestamp, 0)
			fmt.Println("Waktu Redis untuk kunci :", goTime)
		}
	}

}
func TestMPN(t *testing.T) {

	url := "http://10.242.104.66/mpnonline/?page=cGVya3Bwbi5rb25mX2Fka19wYWpha19wZW1kYQ=="
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open("./document.txt")
	defer file.Close()
	part1, errFile1 := writer.CreateFormFile("filecsv", filepath.Base("./document.txt"))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		fmt.Println("disini errFile ", errFile1)
		return
	}
	_ = writer.WriteField("upload", "")
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("Accept-Language", "id-ID,id;q=0.9,en-US;q=0.8,en;q=0.7,zh;q=0.6,ko;q=0.5")
	req.Header.Add("Cache-Control", "max-age=0")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cookie", "PHPSESSID=aqi1lls4dbu53n4pd8iu4klc27")
	req.Header.Add("Origin", "http://10.242.104.66")
	req.Header.Add("Referer", "http://10.242.104.66/mpnonline/?page=cGVya3Bwbi5rb25mX2Fka19wYWpha19wZW1kYQ==")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	//
	//body, err := io.ReadAll(res.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(body))

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal(err)
	}
	rows := make([]model.NTPN, 0)

	doc.Find("#examplex > tbody").Each(func(i int, s *goquery.Selection) {
		//row := new(NTPN)
		//row.KodeBilling = s.Find("#examplex > tbody > tr:nth-child(1) > td:nth-child(2)").Text()
		//row.Ntpn = s.Find("#examplex > tbody > tr:nth-child(6) > td:nth-child(2)").Text()

		for ii := 1; ii <= 30; ii += 5 {
			row := new(model.NTPN)
			str := strconv.Itoa(ii)
			row.Ntpn = s.Find("#examplex > tbody > tr:nth-child(" + str + ") > td:nth-child(2)").Text()
			row.KodeBilling = s.Find("#examplex > tbody > tr:nth-child(" + str + ") > td:nth-child(6)").Text()
			row.Akun = helper.SpaceRemover(s.Find("#examplex > tbody > tr:nth-child(" + str + ") > td:nth-child(8)").Text())
			row.Nilai = helper.SpaceRemover(s.Find("#examplex > tbody > tr:nth-child(" + str + ") > td:nth-child(10)").Text())
			row.Ket = s.Find("#examplex > tbody > tr:nth-child(" + str + ") > td:nth-child(11)").Text()
			rows = append(rows, *row)
		}
		//
		//s.Find("#examplex > tbody > tr:nth-child(1) > td:nth-child(2)").Each(func(j int, td *goquery.Selection) {
		//	// Print the text value of each <td>
		//	fmt.Println(td.Text())
		//})

	})

	bts, err := json.MarshalIndent(rows, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(bts))
}

func TestBaris(t *testing.T) {
	filePath := "document.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0

	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Jumlah baris :", lineCount)
}

type MPNCookie struct {
	Name  string
	Value string
}

func TestGetCookieUsingGolangRod(t *testing.T) {
	u := launcher.New().
		Headless(false).
		MustLaunch()

	page := rod.New().ControlURL(u).MustConnect().MustPage("http://10.242.104.66/mpn").MustWindowFullscreen()
	page.MustEmulate(devices.Device{
		Title:          "iPhone 8",
		Capabilities:   []string{"touch", "mobile"},
		UserAgent:      "Mozilla/5.0 (iPhone; CPU iPhone OS 7_1_2 like Mac OS X)",
		AcceptLanguage: "en",
		Screen: devices.Screen{
			DevicePixelRatio: 2,
			Horizontal: devices.ScreenSize{
				Width:  1920,
				Height: 1080,
			},
		},
	})

	defer page.Close()

	page.MustWaitStable()
	page.MustElement("#header > div > nav > ul > a").MustClick()
	time.Sleep(3 * time.Second)

	page.MustElement("#user").MustInput("kppn155")
	page.MustElement("#login-password").MustInput("Benteng@155")
	bin := page.MustElement("#login > div > div.row.mt-2 > div:nth-child(1) > div > form > div:nth-child(5) > div:nth-child(2) > div > img").MustResource()

	helper.ConvertBinary(bin)

	code := helper.CaptchaSolver("./output.png")
	fmt.Println(code)
	kode := strconv.Itoa(code)
	page.MustElement("#login > div > div.row.mt-2 > div:nth-child(1) > div > form > div:nth-child(5) > input").MustInput(kode)
	time.Sleep(300 * time.Millisecond)
	page.MustElement("#login > div > div.row.mt-2 > div:nth-child(1) > div > form > div.text-center > button").MustClick()
	time.Sleep(2300 * time.Millisecond)

	cook := MPNCookie{}
	cookies := page.MustCookies("http://10.242.104.66/mpnonline/")
	for _, cookie := range cookies {
		cook.Name = cookie.Name
		cook.Value = cookie.Value
	}

	fmt.Println(cook)
}

func TestCaptcha(t *testing.T) {
	client := api2captcha.NewClient("251d4bbc3f50099a8de592680b8d33c2")

	cap := api2captcha.Normal{
		File: "./output.png",
	}

	code, err := client.Solve(cap.ToRequest())
	if err != nil {
		if err == api2captcha.ErrTimeout {
			log.Fatal("Timeout")
		} else if err == api2captcha.ErrApi {
			log.Fatal("API error")
		} else if err == api2captcha.ErrNetwork {
			log.Fatal("Network error")
		} else {
			log.Fatal(err)
		}
	}
	fmt.Println("code " + code)

}
