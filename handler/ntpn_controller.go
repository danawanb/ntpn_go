package handler

import (
	"bufio"
	"bytes"
	"dockerGo/helper"
	"dockerGo/model"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber/v2"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func GetNTPN(c *fiber.Ctx) error {

	ntpn := c.Params("ntpn")
	url := "http://10.242.104.66/mpnonline/pages/dash/cari_detil.kanwil_action.php"
	method := "POST"
	//F9B563IF5FOGFC2F
	payload := strings.NewReader(`masukan=NTPN&data=` + ntpn)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	cookie_token, err := GetToken("token_ntpn_single")
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "tidak dapat mengambil token")
	}
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Cookie", cookie_token)
	req.Header.Add("Origin", "http://10.242.104.66")
	req.Header.Add("Referer", "http://10.242.104.66/mpnonline/index.php?page=Y2FyaV9kZXRpbC5rYW53aWw=")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Mobile Safari/537.36 Edg/114.0.1823.41")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Tidak dapat melakukan request ke MPN Online")
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal(err)
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Tidak membaca data dari MPN Online")
	}
	row := model.NTPN{}

	doc.Find("#examplex > tbody > tr").Each(func(i int, s *goquery.Selection) {
		//row := new(NTPN)
		row.KodeBilling = s.Find("#examplex > tbody > tr > td:nth-child(3)").Text()
		row.Ntpn = s.Find("#examplex > tbody > tr > td:nth-child(4)").Text()
		nilaibefore := s.Find("#examplex > tbody > tr > td:nth-child(7)").Text()
		row.Nilai = helper.SpaceRemover(nilaibefore)

		akunbefore := s.Find("#examplex > tbody > tr > td:nth-child(9)").Text()
		row.Akun = helper.SpaceRemover(akunbefore)
		row.Ket = s.Find("#examplex > tbody > tr > td:nth-child(11)").Text()

		//rows = append(rows, *row)
	})

	return helper.SuccessResponse(c, fiber.StatusOK, row)
}

func BulkNTPN(c *fiber.Ctx) error {
	docu, err := c.FormFile("document")
	if err != nil {
		return err
	}
	c.SaveFile(docu, fmt.Sprintf("./%s", "document.txt"))

	url := "http://10.242.104.66/mpnonline/?page=cGVya3Bwbi5rb25mX2Fka19wYWpha19wZW1kYQ=="
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open("./document.txt")
	defer file.Close()

	file2, err := os.Open("./document.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()
	scanner := bufio.NewScanner(file2)
	lineCount := 0

	for scanner.Scan() {
		lineCount++
	}
	log.Println(lineCount)

	//if err := scanner.Err(); err != nil {
	//	log.Fatal(err)
	//}

	part1, errFile1 := writer.CreateFormFile("filecsv", filepath.Base("./document.txt"))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		fmt.Println("disini errFile ", errFile1)
		return nil
	}
	_ = writer.WriteField("upload", "")
	err = writer.Close()

	if err != nil {
		fmt.Println(err)
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return err
	}

	cookie_token, err := GetToken("token_ntpn_bulk")
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, err.Error()+" tidak dapat mengambil token")
	}

	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("Accept-Language", "id-ID,id;q=0.9,en-US;q=0.8,en;q=0.7,zh;q=0.6,ko;q=0.5")
	req.Header.Add("Cache-Control", "max-age=0")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cookie", cookie_token)
	req.Header.Add("Origin", "http://10.242.104.66")
	req.Header.Add("Referer", "http://10.242.104.66/mpnonline/?page=cGVya3Bwbi5rb25mX2Fka19wYWpha19wZW1kYQ==")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal(err)
	}
	rows := make([]model.NTPN, 0)
	//log.Println(lineCount * 5)
	doc.Find("#examplex > tbody").Each(func(i int, s *goquery.Selection) {
		//row := new(NTPN)
		//row.KodeBilling = s.Find("#examplex > tbody > tr:nth-child(1) > td:nth-child(2)").Text()
		//row.Ntpn = s.Find("#examplex > tbody > tr:nth-child(6) > td:nth-child(2)").Text()

		for ii := 1; ii <= lineCount*5; ii += 5 {
			row := new(model.NTPN)
			str := strconv.Itoa(ii)
			row.Ntpn = s.Find("#examplex > tbody > tr:nth-child(" + str + ") > td:nth-child(2)").Text()
			row.KodeBilling = s.Find("#examplex > tbody > tr:nth-child(" + str + ") > td:nth-child(6)").Text()
			row.Akun = helper.SpaceRemover(s.Find("#examplex > tbody > tr:nth-child(" + str + ") > td:nth-child(8)").Text())
			row.Nilai = helper.SpaceRemover(s.Find("#examplex > tbody > tr:nth-child(" + str + ") > td:nth-child(10)").Text())
			row.Ket = s.Find("#examplex > tbody > tr:nth-child(" + str + ") > td:nth-child(11)").Text()
			rows = append(rows, *row)
		}

	})

	return helper.SuccessResponse(c, fiber.StatusOK, rows)
}

func RefreshTokenUsingGetRequest() {
	url := "http://10.242.104.66/mpnonline/?page=cGVya3Bwbi5rb25mX2Fka19wYWpha19wZW1kYQ=="
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	cookie_token, err := GetToken("token_ntpn_bulk")
	if err != nil {
		log.Println(err.Error() + "tidak bisa mengambil token")
	}

	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("Accept-Language", "id-ID,id;q=0.9,en-US;q=0.8,en;q=0.7,zh;q=0.6,ko;q=0.5")
	req.Header.Add("Cache-Control", "max-age=0")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cookie", cookie_token)
	req.Header.Add("Referer", "http://10.242.104.66/mpnonline/")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	log.Println("refresh tokent")
}

func GetTokenUsingGetRequest() bool {
	url := "http://10.242.104.66/mpnonline/?page=cGVya3Bwbi5rb25mX2Fka19wYWpha19wZW1kYQ=="
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return false
	}

	cookie_token, err := GetToken("token_ntpn_bulk")
	if err != nil {
		log.Println(err.Error() + "tidak bisa mengambil token")
	}

	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("Accept-Language", "id-ID,id;q=0.9,en-US;q=0.8,en;q=0.7,zh;q=0.6,ko;q=0.5")
	req.Header.Add("Cache-Control", "max-age=0")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cookie", cookie_token)
	req.Header.Add("Referer", "http://10.242.104.66/mpnonline/")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer res.Body.Close()

	return true
}
