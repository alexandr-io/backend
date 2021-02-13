package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/gofiber/fiber/v2"
)

type info struct {
	Auth    string
	User    string
	Library string
	Media   string
}

type security struct {
	Auth string
}

func fillTemplateInfo(ctx *fiber.Ctx, folder string) (string, error) {
	url := string(ctx.Request().URI().Scheme()) + "://" + string(ctx.Request().URI().Host())
	infos := info{
		Auth:    url + "/auth",
		User:    url + "/user",
		Library: url + "/library",
		Media:   url + "/media",
	}
	dat, err := ioutil.ReadFile("./" + folder + "/info.yml")
	if err != nil {
		log.Println(err)
		return "", err
	}

	t, err := template.New("info").Parse(string(dat))
	if err != nil {
		log.Println(err)
		return "", err
	}
	var tpl bytes.Buffer
	err = t.Execute(&tpl, infos)
	if err != nil {
		log.Println(err)
		return "", err
	}

	var filePath = "./merged/" + folder + "-info.yml"
	if err := ioutil.WriteFile(filePath, tpl.Bytes(), 0644); err != nil {
		log.Println(err)
		return "", err
	}
	return filePath, nil
}

func fillTemplateSecurity(ctx *fiber.Ctx, folder string) (string, error) {
	if folder == auth {
		return auth + "/security.yml", nil
	}
	url := string(ctx.Request().URI().Scheme()) + "://" + string(ctx.Request().URI().Host())
	sec := security{
		Auth: url + "/auth#section/Authentication",
	}
	if _, err := os.Stat("./" + folder + "/security.yml"); os.IsNotExist(err) {
		return "", fmt.Errorf("NO FILE")
	}
	dat, err := ioutil.ReadFile("./" + folder + "/security.yml")
	if err != nil {
		log.Println(err)
		return "", err
	}

	t, err := template.New("security").Parse(string(dat))
	if err != nil {
		log.Println(err)
		return "", err
	}
	var tpl bytes.Buffer
	err = t.Execute(&tpl, sec)
	if err != nil {
		log.Println(err)
		return "", err
	}

	var filePath = "/tmp/" + folder + "-security.yml"
	if err := ioutil.WriteFile(filePath, tpl.Bytes(), 0644); err != nil {
		log.Println(err)
		return "", err
	}
	return filePath, nil
}
