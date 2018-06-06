package main

import (
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "testing"
    "strings"
    "strconv"
    "encoding/json"
    "bytes"
)

// checkFail fails the tests if err is an error (not nil)
func checkFail(t *testing.T, err error) {
    if err != nil {
        t.Fatal(err)
    }
}

// setup recreates the database according to schema.sql.
// This is so that all tests occurs from a clean slate
func setup(t *testing.T) {
    db := ConnectDb()
    defer db.Close()

    sqlData, err := ioutil.ReadFile("schema.sql")
    checkFail(t, err)

    sqlStmts := string(sqlData)
    for _, stmt := range strings.Split(sqlStmts, ";") {
        _, err := db.Exec(stmt)
        checkFail(t, err)
    }

}

// TestTask tests all functionality with regards to tasks
func TestTask(test *testing.T) {
    // Setup the database and create a test server
    setup(test)
    testServer := httptest.NewServer(Handlers())
    defer testServer.Close()

    // Create a new list for the different tasks
    b, err := json.Marshal(ListCreateRequest{Name: "Work"})
    checkFail(test, err)
    res, err := http.Post(testServer.URL + "/list", "application/json", bytes.NewReader(b))
    checkFail(test, err)

    // Read the response and retrieve the id for the newly created list
    listResp := ListCreateResponse{}
    body, err := ioutil.ReadAll(res.Body)
    defer res.Body.Close()
    err = json.Unmarshal(body, &listResp)

    // Get all the tasks from the new list
    res, err = http.Get(testServer.URL + "/list/" + strconv.Itoa(listResp.Id))
    checkFail(test, err)
    if res.StatusCode != 200 {
        test.Errorf("Expected 200 got %d", res.StatusCode)
    }

    body, err = ioutil.ReadAll(res.Body)
    defer res.Body.Close()
    var tasks []Task
    err = json.Unmarshal(body, &tasks)
    checkFail(test, err)

    // There shouldn't be any tasks in the list yet
    if len(tasks) != 0 {
        test.Errorf("Expected len(tasks) == 0 got %d", len(tasks))
    }

    // Test the creation of a task in the list
    b, err = json.Marshal(CreateTaskRequest{Name: "Work", ListId: listResp.Id})
    checkFail(test, err)
    res, err = http.Post(testServer.URL + "/task", "application/json", bytes.NewReader(b))
    checkFail(test, err)
    if res.StatusCode != 200 {
        test.Errorf("Expected 200, got %d")
    }

    // Test the if the task got stored in the list
    res, err = http.Get(testServer.URL + "/list/" + strconv.Itoa(listResp.Id))
    checkFail(test, err)
    if res.StatusCode != 200 {
        test.Errorf("Expected 200 got %d", res.StatusCode)
    }

    // Check so that the list contains 1 task
    body, err = ioutil.ReadAll(res.Body)
    defer res.Body.Close()
    err = json.Unmarshal(body, &tasks)
    checkFail(test, err)
    if len(tasks) != 1 {
        test.Errorf("Expected len(tasks) == 1 got %d", len(tasks))
    }
}

// TestList manages the testing of lists
func TestList(test *testing.T) {
    // Setup a clean database and a new test server
    setup(test)
    testServer := httptest.NewServer(Handlers())
    defer testServer.Close()

    // Retrieve the lists
    res, err := http.Get(testServer.URL + "/list")
    checkFail(test, err)
    body, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    if res.StatusCode != 200 {
        test.Errorf("Expected status code 200 got %d", res.StatusCode)
    }

    // Check that it didn't contain any lists
    var lists []List
    err = json.Unmarshal(body, &lists)
    checkFail(test, err)

    if len(lists) != 0 {
        test.Errorf("Expected [] got %s", lists)
    }

    // Test creation of a list
    b, err := json.Marshal(ListCreateRequest{Name: "Work"})
    checkFail(test, err)
    res, err = http.Post(testServer.URL + "/list", "application/json", bytes.NewReader(b))
    checkFail(test, err)

    if res.StatusCode != 200 {
        test.Errorf("Expected status code 200 got %d", res.StatusCode)
    }

    // Make sure that the list is stored
    res, err = http.Get(testServer.URL + "/list")
    checkFail(test, err)
    body, err = ioutil.ReadAll(res.Body)
    res.Body.Close()

    if res.StatusCode != 200 {
        test.Errorf("Expected status code 200 got %d", res.StatusCode)
    }

    err = json.Unmarshal(body, &lists)
    checkFail(test, err)
    if len(lists) != 1 {
        test.Errorf("Expected [] got %s", lists)
    }
}
