//usr/bin/env go run $0 "$@"; exit
package main

import (
    "os"
    "fmt"
    "net/http"
    "io/ioutil"
    "github.com/xeipuuv/gojsonschema"
    "consolia-api/utils"
    "strconv"
    "text/tabwriter"
    "errors"
    "path"
    "runtime"
)

func main() {

    if len(os.Args) == 2 {

        token := os.Args[1]

        w := new(tabwriter.Writer)
        w.Init(os.Stdout, 0, 8, 0, '\t', 0)

        validateSchemaOfEndpoint(w, "comics.json",             "GET", "/comic",                       token)
        validateSchemaOfEndpoint(w, "comic.json",              "GET", "/comic/1",                     token)
        validateSchemaOfEndpoint(w, "popular_comics.json",     "GET", "/comic/get_popular",           token)
        validateSchemaOfEndpoint(w, "comic.json",              "GET", "/comic/get_upcoming_comic",    token)
        validateSchemaOfEndpoint(w, "awesome_stats.json",      "GET", "/awesome_stats",               token)
        validateSchemaOfEndpoint(w, "notifications.json",      "GET", "/comic/notifications",         token)
        validateSchemaOfEndpoint(w, "publish_activity.json",   "GET", "/comic/get_publish_activity",  token)
        validateSchemaOfEndpoint(w, "trending_comics.json",    "GET", "/comic_history/get_trending",  token)

        validateSchemaOfEndpoint(w, "error.json",              "GET", "/404040404",                   token)
        validateSchemaOfEndpoint(w, "error.json",              "GET", "/comic/",                      token)
        validateSchemaOfEndpoint(w, "error.json",              "GET", "/comic/666",                   token)
        validateSchemaOfEndpoint(w, "error.json",              "POST", "/comic",                      token)
        validateSchemaOfEndpoint(w, "error.json",              "GET", "/",                            "incorrectTokenTest")

        fmt.Fprintln(w)

        validateStatusCodeOfEndpoint(w, 200, "GET", "/",              token)
        validateStatusCodeOfEndpoint(w, 200, "GET", "/health",        token)

        validateStatusCodeOfEndpoint(w, 200, "GET", "/comic",         token)
        validateStatusCodeOfEndpoint(w, 404, "GET", "/comic/",        token)

        validateStatusCodeOfEndpoint(w, 200, "GET", "/comic/1",       token)
        validateStatusCodeOfEndpoint(w, 404, "GET", "/comic/1/",      token)
        validateStatusCodeOfEndpoint(w, 404, "GET", "/comic/666",     token)

        validateStatusCodeOfEndpoint(w, 404, "GET", "/token",         "incorrectTokenTest")

        validateStatusCodeOfEndpoint(w, 401, "GET", "/",              "incorrectTokenTest")
        validateStatusCodeOfEndpoint(w, 401, "GET", "/health",        "incorrectTokenTest")
        validateStatusCodeOfEndpoint(w, 401, "GET", "/comic",         "incorrectTokenTest")
        validateStatusCodeOfEndpoint(w, 401, "GET", "/404404404",     "incorrectTokenTest")

        w.Flush()
    }
}

func validateSchemaOfEndpoint(w *tabwriter.Writer, schema string, requestMethod string, endpoint string, token string) error {

    // Schema
    // cwd, _ := os.Getwd()
    _, filename, _, _ := runtime.Caller(1)
    cwd := path.Dir(filename)
    schemaLoader := gojsonschema.NewReferenceLoader("file://" + cwd + "/" + schema)

    // Document
    body, _, _ := request(requestMethod, endpoint, token)
    documentLoader := gojsonschema.NewStringLoader(body)

    result, err := gojsonschema.Validate(schemaLoader, documentLoader)

    if err != nil {
        panic(err.Error())
    }

    if result.Valid() {
        fmt.Fprintln(w, "schema\t" + schema + "\t" + requestMethod + " " + endpoint + "\tOK")
        return nil
    } else {
        fmt.Fprintln(w, "schema\t" + schema + "\t" + requestMethod + " " + endpoint + "\tNot valid. Errors:")
        for _, desc := range result.Errors() {
            fmt.Fprintln(w, "- " + desc.Description())
        }

        return errors.New("Not valid")
    }
}

func validateStatusCodeOfEndpoint(w *tabwriter.Writer, statusCode int, requestMethod string, endpoint string, token string) error {

    _, returnedStatusCode, _ := request(requestMethod, endpoint, token)

    if statusCode == returnedStatusCode {
        fmt.Fprintln(w, "status\t" + strconv.Itoa(statusCode) + "\t" + requestMethod + " " + endpoint + " \tOK")
        return nil
    } else {
        fmt.Fprintln(w, "status\t" + strconv.Itoa(statusCode) + "\t" + requestMethod + " " + endpoint + " \tNot valid. Returns " + strconv.Itoa(returnedStatusCode))
        return errors.New("Not valid. Expecting " + strconv.Itoa(statusCode) + ", received " + strconv.Itoa(returnedStatusCode))
    }
}

func request(requestMethod string, endpoint string, token string) (string, int, error) {

    conf, _ := utils.GetConf()

    client := &http.Client{}

    req, err := http.NewRequest(requestMethod, "http://127.0.0.1:" + conf.Port + endpoint, nil)

    req.Header.Add("Authorization", "Bearer " + token)
    resp, err := client.Do(req)

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    return string(body), resp.StatusCode, err
}
