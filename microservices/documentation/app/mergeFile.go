package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/miracl/conflate"
)

func getAllFilesToMerge(thisDir string) ([]string, error) {
	var matches []string
	err := filepath.Walk(thisDir, func(path string, f os.FileInfo, err error) error {
		if filepath.Base(path) == "swagger.yml" {
			return nil
		}
		if filepath.Ext(path) == ".yml" || filepath.Ext(path) == ".yaml" {
			matches = append(matches, path)
		}
		return nil
	})
	return matches, err
}

func mergeDocsFiles(ctx *fiber.Ctx) error {
	thisDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return err
	}

	matches, err := getAllFilesToMerge(thisDir)
	if err != nil {
		log.Println(err)
		return err
	}

	c := conflate.New()
	if err := c.AddFiles(matches...); err != nil {
		log.Println(err)
		return err
	}
	yaml, err := c.MarshalYAML()
	if err != nil {
		log.Println(err)
		return err
	}

	if err := ioutil.WriteFile(thisDir+"/swagger.yml", yaml, 0644); err != nil {
		log.Println(err)
		return err
	}
	return ctx.Next()
}