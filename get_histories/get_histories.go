package get_histories

import (
	"browser_history/get_users"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

var global_users_slice = get_users.Get_Users()
var global_dbs = []string{}
var global_targets = map[string]string{
	"safari":  "/Library/Safari/History.db",
	"firefox": "/Library/Application Support/Firefox/Profiles/",
	"chrome":  "/Library/Application Support/Google/Chrome/Default/History",
	"edge":    "/Library/Application Support/Microsoft Edge/Default/History",
}
var global_sqlite_dbs = map[string]string{}

func path_exists(_path string) bool {
	if _, err := os.Stat(_path); err == nil {
		return true
	} else {
		return false
	}
}

func get_db(_global_target string, tag string) {
	for i := 0; i < len(global_users_slice); i++ {
		_path := "/Users/" + global_users_slice[i] + _global_target
		if path_exists(_path) {
			// global_dbs = append(global_dbs, _path)

			global_sqlite_dbs[tag] = _path
		}
	}
}

func get_firefox_dbs() {
	for i := 0; i < len(global_users_slice); i++ {
		history_path := "/Users/" + string(global_users_slice[i]) + global_targets["firefox"]
		if path_exists(history_path) {
			current := strconv.Itoa(i)
			read_and_store_dir(history_path, "places.sqlite", "firefox-"+current)
		}
	}
}

func read_and_store_dir(target_dir string, target_object string, tag string) {
	files, err := ioutil.ReadDir(target_dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		target_obj := target_dir + file.Name() + "/" + target_object
		if path_exists(target_obj) {
			// global_dbs = append(global_dbs, target_obj)
			global_sqlite_dbs[tag] = target_obj
		}
	}
}

func Get_Histories() map[string]string {
	get_db(global_targets["safari"], "safari")
	get_db(global_targets["chrome"], "chrome")
	get_db(global_targets["edge"], "edge")
	get_firefox_dbs()
	return global_sqlite_dbs
}
