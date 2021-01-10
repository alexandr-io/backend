package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/miracl/conflate"
)

func getAllFilesToMerge(directory string) ([]string, error) {
	var matches []string
	err := filepath.Walk(directory, func(path string, f os.FileInfo, err error) error {
		if filepath.Base(path) == "swagger.yml" ||
			filepath.Base(path) == "info.yml" ||
			filepath.Base(path) == "security.yml" {
			return nil
		}
		if filepath.Ext(path) == ".yml" || filepath.Ext(path) == ".yaml" {
			matches = append(matches, path)
		}
		return nil
	})
	return matches, err
}

func mergeDocsFiles(ctx *fiber.Ctx, folder string) error {
	thisDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return err
	}

	matches, err := getAllFilesToMerge(path.Join(thisDir, folder))
	if err != nil {
		log.Println(err)
		return err
	}

	infoPath, err := fillTemplateInfo(ctx, folder)
	if err != nil {
		return err
	}
	securityPath, err := fillTemplateSecurity(ctx, folder)
	if err != nil {
		if err.Error() != "NO FILE" {
			return err
		}
	} else {
		matches = append(matches, securityPath)
	}
	matches = append(matches, infoPath)

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

	if err := ioutil.WriteFile(path.Join(path.Join(thisDir, "merged"), folder+".yml"), yaml, 0644); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
