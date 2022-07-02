package handler

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func TestUpload(t *testing.T) {
	fd, err := os.Open("/home/xx/Videos/Webcam/1.webm")
	if err != nil {
		t.Fatal(err)
	}
	fileStat, err := fd.Stat()
	if err != nil {
		t.Fatal(err)
	}
	encoder := md5.New()
	buf := make([]byte, 150*1024*1024)
	io.Copy(encoder, fd)
	totalPart := fileStat.Size() / int64(cap(buf))
	if fileStat.Size()%int64(cap(buf)) > 0 {
		totalPart++
	}
	fileHash := hex.EncodeToString(encoder.Sum(nil))
	var i = 0
	var index = 0
	var n = 0
	for {
		index += n
		n, err = fd.ReadAt(buf, int64(index))
		if n <= 0 && (err == io.EOF) {
			break
		}
		if err != nil && err != io.EOF {
			panic(err)
		}
		i++
		data := url.Values{}
		encoder.Reset()
		encoder.Write(buf[:n])
		partHash := hex.EncodeToString(encoder.Sum(nil))
		data["file_name"] = []string{fd.Name()}
		data["file_hash"] = []string{fileHash}
		data["file_part_hash"] = []string{partHash}
		data["file_size"] = []string{fmt.Sprintf("%d", fileStat.Size())}
		data["file_total_part"] = []string{fmt.Sprintf("%d", totalPart)}
		data["file_part_num"] = []string{fmt.Sprintf("%d", i)}
		data["file_pwd"] = []string{"test"}
		for k, v := range data {
			log.Printf("%s: %d", k, len(v[0]))
		}
		u := "http://127.0.0.1:8080/upload?" + data.Encode()
		log.Printf(u)
		func(d url.Values) {
			res, err := http.Post("http://127.0.0.1:8080/upload?"+data.Encode(), "multipart/form-data", bytes.NewReader(buf[:n]))
			if err != nil {
				t.Log(err)
			}
			b, _ := io.ReadAll(res.Body)
			t.Log(string(b))
		}(data)
		log.Printf("send %d", i)
	}
}

func TestMerge(t *testing.T) {
	fd, err := os.Open("/home/xx/Videos/Webcam/1.webm")
	if err != nil {
		t.Fatal(err)
	}
	fileStat, err := fd.Stat()
	if err != nil {
		t.Fatal(err)
	}
	encoder := md5.New()
	buf := make([]byte, 150*1024*1024)
	io.Copy(encoder, fd)
	totalPart := fileStat.Size() / int64(cap(buf))
	if fileStat.Size()%int64(cap(buf)) > 0 {
		totalPart++
	}
	fileHash := hex.EncodeToString(encoder.Sum(nil))

	var data = url.Values{}
	data["file_name"] = []string{fd.Name()}
	data["file_hash"] = []string{fileHash}
	data["file_size"] = []string{fmt.Sprintf("%d", fileStat.Size())}
	data["file_total_part"] = []string{fmt.Sprintf("%d", totalPart)}
	data["file_pwd"] = []string{"test"}

	res, err := http.PostForm("http://127.0.0.1:8080/merge", data)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := io.ReadAll(res.Body)

	var resStruct = struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}{}
	json.Unmarshal(b, &resStruct)
	if resStruct.Code != 0 {
		t.Fatal(resStruct)
	}

	newFd, err := os.Open("/home/xx/storage/test/storage/" + fileHash)
	if err != nil {
		t.Fatal(err)
	}

	encoder.Reset()
	io.Copy(encoder, newFd)
	newFileHash := hex.EncodeToString(encoder.Sum(nil))
	if fileHash != newFileHash {
		t.Log("hash is different")
	}
	t.Log(fileHash)
	t.Log(newFileHash)
}
