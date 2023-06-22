package handler

import (
	"context"
	"dockerGo/db"
	"dockerGo/helper"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"log"
)

func InsertToken(c *fiber.Ctx) error {

	tt := c.FormValue("token")

	rd := db.NewRedis()
	ctx := context.Background()

	rdb := rd.Conn()

	if err := rdb.Set(ctx, "token_ntpn_bulk", tt, 0).Err(); err != nil {
		return c.Render("statusinsert", fiber.Map{
			"status": "Gagal",
		})
	}

	return c.Render("statusinsert", fiber.Map{
		"status": "Berhasil",
	})

}

func InsertNTPN(c *fiber.Ctx) error {
	single := c.Query("single")
	bulk := c.Query("bulk")

	rd := db.NewRedis()
	ctx := context.Background()

	rdb := rd.Conn()

	if single != "" {
		if err := rdb.Set(ctx, "token_ntpn_single", single, 0).Err(); err != nil {
			return helper.ErrorResponse(c, fiber.StatusInternalServerError, "gagal simpan")
		}

		return helper.SuccessResponse(c, fiber.StatusOK, "Berhasil simpan")
	}

	if bulk != "" {
		if err := rdb.Set(ctx, "token_ntpn_bulk", bulk, 0).Err(); err != nil {
			return helper.ErrorResponse(c, fiber.StatusInternalServerError, "gagal simpan")
		}

		return helper.SuccessResponse(c, fiber.StatusOK, "Berhasil simpan")

	}
	return helper.ErrorResponse(c, fiber.StatusBadRequest, "Tidak ada data tokennya")
}

func GetToken(key string) (string, error) {
	rd := db.NewRedis()
	ctx := context.Background()

	rdb := rd.Conn()

	result, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", errors.New("key tidak ada")
	} else if err != nil {
		return "", err
	} else {
		return result, nil
	}
}

func InsertTokenFromMPN(c *fiber.Ctx) error {

	rd := db.NewRedis()
	ctx := context.Background()

	rdb := rd.Conn()

	tt := helper.GetTokenFromMPN()
	if err := rdb.Set(ctx, "token_ntpn_single", tt, 0).Err(); err != nil {
		return err
	}

	if err := rdb.Set(ctx, "token_ntpn_bulk", tt, 0).Err(); err != nil {
		return err
	}
	log.Println("Berhasil insert token", tt)
	return helper.SuccessResponse(c, fiber.StatusOK, "Berhasil Update Token")
}

func InsertTokenFromNPNCron() {

	rd := db.NewRedis()
	ctx := context.Background()

	rdb := rd.Conn()

	tt := helper.GetTokenFromMPN()
	if err := rdb.Set(ctx, "token_ntpn_single", tt, 0).Err(); err != nil {
		log.Println(err)
		return
	}

	if err := rdb.Set(ctx, "token_ntpn_bulk", tt, 0).Err(); err != nil {
		log.Println(err)
		return
	}

	log.Println("Berhasil insert token", tt)

}
