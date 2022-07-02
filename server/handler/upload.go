package handler

import (
	"fileserver/model"
	"fileserver/utils"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type FileUploadRequest struct {
	FileName      string
	FileHash      string // unique file
	FilePartHash  string // unique file part
	FileSize      int64
	FileTotalPart int
	FilePartNum   int // 0 mean create request
	FilePWD       string
}

func GetFileUploadRequest(r *http.Request, deal bool) (*FileUploadRequest, *model.ResponseCode) {
	var req = new(FileUploadRequest)
	var ok bool = true
	req.FileName, ok = utils.GetFormString(r, "file_name")
	if !ok {
		return nil, model.ResponseCodeMissingParam
	}
	req.FileHash, ok = utils.GetFormString(r, "file_hash")
	if !ok {
		return nil, model.ResponseCodeMissingParam
	}
	req.FileSize, ok = utils.GetFormInt64(r, "file_size")
	if !ok {
		return nil, model.ResponseCodeMissingParam
	}
	req.FileTotalPart, ok = utils.GetFormInt(r, "file_total_part")
	if !ok {
		return nil, model.ResponseCodeMissingParam
	}
	req.FilePWD, ok = utils.GetFormString(r, "file_pwd")
	if !ok {
		return nil, model.ResponseCodeMissingParam
	}
	if deal {

		req.FilePartHash, ok = utils.GetFormString(r, "file_part_hash")
		if !ok {
			return nil, model.ResponseCodeMissingParam
		}
		// TODO check
		req.FilePartNum, ok = utils.GetFormInt(r, "file_part_num")
		if !ok {
			return nil, model.ResponseCodeMissingParam
		}
	}
	return req, nil
}

var tempCache = make(map[string]*FileUploadRequest)

// CreateUploadTask param:
func CreateUploadTask(r *http.Request, userinfo *model.Userinfo) *model.ResponseCode {
	r.ParseForm()
	_, resCode := GetFileUploadRequest(r, false)
	if resCode != nil {
		return resCode
	}

	// TODO cache fileinfo
	return model.ResponseCodeOK
}

func DealUploadTask(r *http.Request, userinfo *model.Userinfo) *model.ResponseCode {
	r.ParseForm()
	req, resCode := GetFileUploadRequest(r, true)
	if resCode != nil {
		return resCode
	}

	filepath := fmt.Sprintf("%s/temp/%s", userinfo.StoragePath, req.FileHash)
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		os.MkdirAll(filepath, 0766)
	}
	filepath += fmt.Sprintf("/%d", req.FilePartNum)

	fd, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0766)
	if err != nil {
		log.Printf("%s: %v", filepath, err)
		return model.ResponseCodeFailedToAllocateStorageSpace
	}

	buf := make([]byte, 1024*1024)
	for {
		n, err := r.Body.Read(buf)
		log.Printf("got body of %d with %d, %v", req.FilePartNum, n, err)
		if n <= 0 && err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			fd.Close()
			os.Remove(filepath)
			return model.ResponseCodeFailedToWriteIntoFile
		}
		_, err = fd.Write(buf[:n])
		log.Printf("write into file with err: %v", err)
		if err != nil {
			fd.Close()
			os.Remove(filepath)
			return model.ResponseCodeFailedToWriteIntoFile
		}
	}

	fd.Close()
	// TODO cache part of file upload success

	return model.ResponseCodeOK
}

func MergeUploadTask(r *http.Request, userinfo *model.Userinfo) *model.ResponseCode {
	r.ParseForm()
	req, resCode := GetFileUploadRequest(r, false)
	if resCode != nil {
		return resCode
	}

	// TODO get cache of upload task and recheck

	// TODO if recheck failed, need ReUpload

	filepath := fmt.Sprintf("%s/storage", userinfo.StoragePath)
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		os.MkdirAll(filepath, 0766)
	}
	fd, err := os.OpenFile(filepath+"/"+req.FileHash, os.O_CREATE|os.O_WRONLY, 0766)
	if err != nil {
		return model.ResponseCodeFailedToAllocateStorageSpace
	}

	// TODO part check
	var buf = make([]byte, 1024*1024)
	tempPath := fmt.Sprintf("%s/temp/%s/", userinfo.StoragePath, req.FileHash)
	for i := 1; i <= req.FileTotalPart; i++ {
		tmpFd, err := os.Open(fmt.Sprintf("%s%d", tempPath, i))
		if err != nil {
			fd.Close()
			os.Remove(filepath)
			return model.ResponseCodeFailedToWriteIntoFile
		}
		for {
			n, err := tmpFd.Read(buf)
			if n <= 0 && err == io.EOF {
				break
			}
			if err != nil {
				fd.Close()
				os.Remove(filepath)
				return model.ResponseCodeFailedToWriteIntoFile
			}
			// fd.Write(buf[:n])
			_, err = fd.Write(buf[:n])
			if err != nil {
				fd.Close()
				os.Remove(filepath)
				return model.ResponseCodeFailedToWriteIntoFile
			}
		}
	}
	fd.Close()
	return model.ResponseCodeOK
}
