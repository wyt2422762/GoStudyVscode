package webserver

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	uploadDir      = "./upload"   //文件上传目录
	templateDir    = "./template" //模板文件目录
	templateSuffix = ".html"      //模板文件后缀
)

var templates = make(map[string]*template.Template) //全局变量，key为模板文件路径

func server() {
	fmt.Println("服务端启动")
	http.HandleFunc("/toUpload", baseHandler(toUploadHandler))
	http.HandleFunc("/", baseHandler(indexHandler))
	http.HandleFunc("/test", baseHandler(testHandler))
	http.HandleFunc("/upload", baseHandler(uploadHandler))
	err := http.ListenAndServe(":8181", nil)
	if err != nil {
		fmt.Println("启动服务端出错：", err)
	}
}

func init() {
	fmt.Println("init方法")
	//初始化加载模板
	initTemplate(templateDir)
}

func initTemplate(dir string) {
	//遍历template目录下的所有文件
	fileArr, err := ioutil.ReadDir(dir)
	check(err)
	//遍历文件
	for _, fileInfo := range fileArr {
		if fileInfo.IsDir() { //如果是文件夹继续遍历
			initTemplate(dir + "/" + fileInfo.Name())
		} else { //是模板文件存入到全局模板map中，key为文件路径
			key := dir + "/" + fileInfo.Name()
			fmt.Println("加载模板文件：", key)
			t := template.Must(template.ParseFiles(key)) //解析模板，Must方法确保出错会panic异常
			templates[key] = t                           //将模板存入全局map
		}
	}
}

//统一抛出异常方法
func check(err error) {
	if err != nil {
		panic(err) //抛出异常，panic相当于java中的throw
	}
}

//baseHandler
func baseHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//统一异常处理
		defer func() {
			if r := recover(); r != nil { //recover相当于java中的catch
				fmt.Println("捕获的错误：", r)
				data := make(map[string]interface{}) //模板数据
				data["statusCode"] = http.StatusInternalServerError
				data["err"] = r
				renderTemplate(w, "/err.html", data) //渲染模板
			}
		}()
		fn(w, r) //执行具体的handler
	}
}

//testHandler
func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("testHandler")
}

//首页处理方法
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("indexHandler")
	renderTemplate(w, "/index.html", nil) //渲染模板
}

// /toUpload处理方法
func toUploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("toUploadHandler")
	renderTemplate(w, "/upload/upload.html", nil) //渲染模板
}

//上传文件处理方法
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("uploadHandler")
	if r.Method == http.MethodPost {
		f, fHeader, err := r.FormFile("file") //获取上传的文件
		check(err)
		filename := fHeader.Filename
		defer f.Close()
		t, err := ioutil.TempFile(uploadDir, filename)
		check(err)
		defer t.Close()
		_, err = io.Copy(t, f)
		check(err)
	}
	http.Redirect(w, r, "/", http.StatusOK) //重定向到首页
}

//封装的渲染模板的方法
/* func renderTemplate(w http.ResponseWriter, tPath string, data map[string]interface{}) error {
	t, err := template.ParseFiles(templateDir + tPath) //这里可以优化改成预加载
	if err != nil {
		fmt.Println("renderTemplate: 解析模板" + templateDir + tPath + "错误")
		return err
	}
	err = t.Execute(w, data)
	return err
} */
//渲染模板方法2，配合统一异常处理使用
func renderTemplate(w http.ResponseWriter, tPath string, data map[string]interface{}) {
	/* t, err := template.ParseFiles(templateDir + tPath) //这里可以优化改成预加载
	check(err)
	err = t.Execute(w, data)
	check(err) */
	key := templateDir + tPath
	t := templates[key]
	err := t.Execute(w, data)
	check(err)
}
