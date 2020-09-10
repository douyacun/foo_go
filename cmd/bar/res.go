package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)

type InAppBidding struct {
	Key int `json:"key"`
}

func main() {
	configFile := flag.String("config", "", "sdk json配置文件")
	url := flag.String("url", "", "sdk json url")
	key := flag.String("key", "", "媒体密钥(aes密钥)")
	flag.Parse()
	if *configFile == "" && *url == "" {
		log.Fatalf("请指定配置文件或配置文件地址: --config <filepath> or --url <http>")
	}
	if *key == "" {
		log.Fatalf("请指定媒体密钥：--key <key>")
	}
	var data []byte
	if *configFile != "" {
		f, err := os.Open(*configFile)
		defer f.Close()
		if err != nil {
			log.Fatalf("err: %s", err.Error())
		}
		data, err = ioutil.ReadAll(f)
		if err != nil {
			log.Fatalf("err: %s", err.Error())
		}
	} else {
		resp, err := http.Get(*url)
		if err != nil {
			log.Fatalf("http err: %s", err)
		}
		defer resp.Body.Close()
		data, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("http body read err: %s", err.Error())
		}
	}
	res, err := AesDecrypt(data, []byte(*key))
	if err != nil {
		log.Fatalf("err: %s", err.Error())
	}
	inAppBidding := InAppBidding{}
	if err := json.Unmarshal(res, &inAppBidding); err != nil {
		log.Fatalf("err: %s", err.Error())
	}
	if err := ioutil.WriteFile(path.Join(os.Getenv("HOME"), "Desktop", fmt.Sprintf("%d.json", inAppBidding.Key%1000000000)), res, 0666); err != nil {
		log.Fatalf("err: %s", err.Error())
	}
}

//@brief:AES解密
func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//AES分组长度为128位，所以blockSize=16，单位字节
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) //初始向量的长度必须等于块block的长度16字节
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
