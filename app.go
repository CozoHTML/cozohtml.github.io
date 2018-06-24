package main

import (
    "fmt"
    "html/template"
    "io"
    "log"
    "net/http"
    "os"
)

var buf []byte

func upload(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    if r.Method == "POST" {
        file, handle, err := r.FormFile("file")
        checkErr(err)
        f, err := os.OpenFile("./"+handle.Filename, os.O_WRONLY|os.O_CREATE, 0666)
        io.Copy(f, file)
        checkErr(err)
        defer f.Close()
        defer file.Close()
        fmt.Println("upload success")
    }
}

func checkErr(err error) {
    if err != nil {
        err.Error()
    }
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    var page = '
<html>
    <head></head>
    <body>
        <form action="upload" method="post" enctype="multipart/form-data">
            <input type="file" name="file" value="" /> 
            <input type="submit" name="submit" />
        </form>
    </body>
</html>
    '
    fmt.Fprintln(w, page)
}

func main() {
    http.HandleFunc("/", IndexHandler)
    http.HandleFunc("/upload", upload)
    err := http.ListenAndServe(":8888", nil)
    if err != nil {
        log.Fatal("listenAndServe: ", err)
    }
}
