package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path"
	"strings"
	"time"

	"github.com/gobuffalo/packr/v2"
)

func main() {
	u, err := user.Current()
	if err != nil {
		log.Fatalf("ERROR user.Current: %v", err)
	}
	vimrcFile := fmt.Sprintf("%s%c%s", u.HomeDir, os.PathSeparator, ".vimrc")
	vimDir := fmt.Sprintf("%s%c%s", u.HomeDir, os.PathSeparator, ".vim")
	now := time.Now()
	vimrcBackupFile := fmt.Sprintf("%s.%d", vimrcFile, now.Unix())
	vimBackupDir := fmt.Sprintf("%s.%d", vimDir, now.Unix())
	if _, err := os.Stat(vimrcFile); err != nil {
		if os.IsNotExist(err) {
			// pass
		} else {
			log.Fatalf("ERROR os.Stat: %v (%s)", err, vimrcFile)
		}
	} else {
		if err := os.Rename(vimrcFile, vimrcBackupFile); err != nil {
			log.Fatalf("ERROR os.Rename: %v (%s)", err, vimrcBackupFile)
		}
	}
	if _, err := os.Stat(vimDir); err != nil {
		if os.IsNotExist(err) {
			// pass
		} else {
			log.Fatalf("ERROR os.Stat: %v (%s)", err, vimDir)
		}
	} else {
		if err := os.Rename(vimDir, vimBackupDir); err != nil {
			log.Fatalf("ERROR os.Rename: %v (%s)", err, vimBackupDir)
		}
	}

	box := packr.New("assets", "assets")
	list := box.List()
	for _, file := range list {
		destFile := ""
		prefix := fmt.Sprintf("%s%c", "vim", os.PathSeparator)
		if strings.HasPrefix(file, prefix) {
			destFile = fmt.Sprintf("%s%c%s", vimDir, os.PathSeparator, strings.TrimPrefix(file, prefix))
		} else if file == "vimrc" {
			destFile = vimrcFile
		} else {
			continue
		}
		dir := path.Dir(destFile)
		os.MkdirAll(dir, 0755)
		fileBytes, err := box.Find(file)
		if err != nil {
			log.Fatalf("ERROR packr.Box.Find: %v (%s)", err, file)
		}
		if err := ioutil.WriteFile(destFile, fileBytes, 0644); err != nil {
			log.Fatalf("ERROR ioutil.WriteFile: %v (%s)", err, destFile)
		}
	}
}
