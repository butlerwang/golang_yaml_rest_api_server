package memdb

var memdatabase = make(map[string]interface{})

func Save(key string, item interface{}) {
	memdatabase[key] = item
}

func Get(key string) (interface{}, bool) {
	if app, ok := memdatabase[key]; ok {
		return app, ok
	} else {
		return nil, false
	}
}

func GetAll() ([]interface{}, bool) {
	if len(memdatabase) == 0 {
		return nil, false
	} else {

		apps := make([]interface{}, 0, len(memdatabase))
		for _, v := range memdatabase {
			apps = append(apps, v)
		}
		return apps, true
	}
}

func Remove(key string) {
	delete(memdatabase, key)
}
