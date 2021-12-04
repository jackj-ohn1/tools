package myself

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/howeyc/gopass"
)

func Init(token []string) string {
	str := strings.Split(token[0], ":")

	return str[1]
}
func Get() string {
	//获取未加密的Token
	url := "http://pass.muxi-tech.xyz/auth/api/signin"
	method := "POST"

	type first struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var (
		password string
		username string
	)
	fmt.Println("请输入你的用户名：")
	fmt.Scanf("%s", &username)
	fmt.Scanln()
	fmt.Println("请输入密码：")
	while, _ := gopass.GetPasswdMasked()

	password = base64.StdEncoding.EncodeToString(while)

	infor := first{
		Password: password,
		Username: username,
	}

	content, _ := json.Marshal(infor)
	payload := strings.NewReader(string(content))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	re := regexp.MustCompile(`"token":"(.*?)"`)
	token := re.FindAllString(string(body), -1)

	//获取加密后的Token
	tmp := Init(token)

	type second struct {
		Email string `json:"email"`
		Token string `json:"token"`
	}

	url = "http://work.muxi-tech.xyz/api/v1.0/auth/login/"
	method = "POST"

	other := second{
		Email: username,
		Token: tmp,
	}

	content, _ = json.Marshal(other)
	payload = strings.NewReader(string(content))
	req, err = http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)

	}
	req.Header.Add("Content-Type", "application/json")

	res, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	re = regexp.MustCompile(`"token": "(.*?)"`)
	token = re.FindAllString(string(body), -1)

	end := []byte(Init(token))
	return string(end[2 : len(end)-1])
}
