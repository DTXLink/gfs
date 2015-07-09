package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math"
	"os"
	//_ "reflect"
	"regexp"
	//"runtime"
	"strconv"
	"strings"
)

/**
 * @brief is_md5 Check the string is a md5 style.
 *
 * @param s The string.
 *
 * @return 1 for yes and -1 for no.
 */
func is_md5(str string) bool {
	regular := `^([0-9a-zA-Z]){32}$`
	regx := regexp.MustCompile(regular)
	return regx.MatchString(str)
}

func gen_md5() string {
	h := md5.New()
	return hex.EncodeToString(h.Sum(nil))
}

func encode_base64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

/**
 * @brief get_type It tell you the type of a file.
 *
 * @param filename The name of the file.
 * @param type Save the type string.
 *
 * @return 1 for success and -1 for fail.
 */
func get_type(file_name string) (string, error) {
	i := strings.LastIndex(file_name, ".")
	if i == -1 {
		return "", fmt.Errorf("FileName [%s] Has No '.' in It.", file_name)
	}
	//fmt.Printf("ext index : %d", i)

	ext := file_name[(i + 1):len(file_name)]
	return ext, nil
}

func is_exist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

/**
 * @brief gen_key Generate storage key from md5 and other args.
 *
 * @param key The key string.
 * @param md5 The md5 string.
 * @param argc Count of args.
 * @param ... Args.
 *
 * @return Generate result.
 */
func gen_key(md5 string, args ...interface{}) string {
	s := []string{}
	s = append(s, md5)
	for _, argv := range args {
		switch v := argv.(type) {
		case string:
			s = append(s, v)
		case int:
			s = append(s, strconv.Itoa(v))
		}
	}
	return strings.Join(s, ":")
}

func gen_md5_str(data []byte) string {
	fmt.Println("Begin to Caculate MD5...")
	m := md5.New()
	m.Write(data)
	return hex.EncodeToString(m.Sum(nil))
}

func round(val float64) float64 {
	if val > 0.0 {
		return math.Floor(val + 0.5)
	} else {
		return math.Ceil(val - 0.5)
	}
}
