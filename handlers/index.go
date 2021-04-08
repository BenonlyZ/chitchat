package handlers

import (
	"net/http"

	"chitchat/models"
)

// 论坛首页路由处理器方法
func Index(w http.ResponseWriter, r *http.Request) {
	/* files := []string{"views/layout.html", "views/navbar.html", "views/index.html"}
	templates := template.Must(template.ParseFiles(files...))
	threads, err := models.Threads()
	if err == nil {
		templates.ExecuteTemplate(w, "layout", threads)
	}
	*/
	threads, err := models.Threads()
	if err == nil {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, threads, "layout", "navbar", "index") //还没登录时调用
		} else {
			generateHTML(w, threads, "layout", "auth.navbar", "index") //有账户登录时调用
		}
	}
}

func Err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err := session(writer, request)
	if err != nil {
		generateHTML(writer, vals.Get("msg"), "layout", "navbar", "error")
	} else {
		generateHTML(writer, vals.Get("msg"), "layout", "auth.navbar", "error")
	}
}
