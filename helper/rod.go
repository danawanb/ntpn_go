package helper

import (
	"dockerGo/model"
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/devices"
	"github.com/go-rod/rod/lib/launcher"
	"log"
	"strconv"
	"time"
)

func GetTokenFromMPN() string {
	path, _ := launcher.LookPath()

	u := launcher.New().Bin(path).MustLaunch()
	page := rod.New().ControlURL(u).MustConnect().MustPage("http://10.242.104.66/mpn")
	//u := launcher.New().Set("disable-web-security").
	//	Set("disable-setuid-sandbox").
	//	Set("no-sandbox").
	//	Set("no-first-run", "true").
	//	Set("disable-gpu").
	//	Headless(true).
	//	MustLaunch()
	//
	//page := rod.New().ControlURL(u).MustConnect().MustPage("http://10.242.104.66/mpn").MustWindowFullscreen()
	page.MustEmulate(devices.Device{
		Title:          "iPhone 8",
		Capabilities:   []string{"touch", "mobile"},
		UserAgent:      "Mozilla/5.0 (iPhone; CPU iPhone OS 7_1_2 like Mac OS X)",
		AcceptLanguage: "en",
		Screen: devices.Screen{
			Horizontal: devices.ScreenSize{
				Width:  1920,
				Height: 1080,
			},
		},
	})

	defer page.Close()

	time.Sleep(2 * time.Second)
	page.MustElement("#header > div > nav > ul > a").MustClick()
	time.Sleep(2 * time.Second)
	log.Println("masuk mpn")
	page.MustElement("#user").MustInput("kppn155")
	page.MustElement("#login-password").MustInput("Benteng@155")
	bin := page.MustElement("#login > div > div.row.mt-2 > div:nth-child(1) > div > form > div:nth-child(5) > div:nth-child(2) > div > img").MustResource()

	ConvertBinary(bin)

	code := CaptchaSolver("./output.png")
	fmt.Println(code)
	kode := strconv.Itoa(code)
	page.MustElement("#login > div > div.row.mt-2 > div:nth-child(1) > div > form > div:nth-child(5) > input").MustInput(kode)
	time.Sleep(300 * time.Millisecond)
	page.MustElement("#login > div > div.row.mt-2 > div:nth-child(1) > div > form > div.text-center > button").MustClick()
	time.Sleep(2300 * time.Millisecond)
	log.Println("berhasil login")
	cook := model.MPNCookie{}
	cookies := page.MustCookies("http://10.242.104.66/mpnonline/")
	for _, cookie := range cookies {
		cook.Name = cookie.Name
		cook.Value = cookie.Value
	}

	cok := cook.Name + "=" + cook.Value
	return cok
}
