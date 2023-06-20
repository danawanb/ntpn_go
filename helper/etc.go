package helper

import (
	"bytes"
	api2captcha "github.com/2captcha/2captcha-go"
	"image"
	"image/png"
	"log"
	"os"
	"strconv"
	"strings"
)

func SpaceRemover(s string) string {
	nilaiWithSpace := strings.ReplaceAll(s, " ", "")
	nilai := strings.ReplaceAll(nilaiWithSpace, "\n", "")
	nil := strings.ReplaceAll(nilai, ".", "")
	return nil
}

func CaptchaSolver(path string) int {
	client := api2captcha.NewClient("251d4bbc3f50099a8de592680b8d33c2")

	cap := api2captcha.Normal{
		File: path,
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
	num, err := strconv.Atoi(code)
	if err != nil {
		log.Println(err)
	}
	return num
}

func ConvertBinary(binaryData []byte) {

	file, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := image.Decode(bytes.NewReader(binaryData))
	if err != nil {
		panic(err)
	}

	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}
