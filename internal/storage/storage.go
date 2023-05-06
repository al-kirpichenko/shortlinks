package storage

//var urls = make(map[string]string)
//
//func GetURL(id string) (string, error) {
//	result, ok := urls[id]
//	if ok {
//		return result, nil
//	}
//	return "", errors.New("id not found")
//}
//
//func SetURL(URL, id string) {
//	urls[id] = URL
//}

type Storage struct {
	Urls map[string]string
}
