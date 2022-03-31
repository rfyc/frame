package route

type Input struct {
	Request map[string]interface{}
	Get     map[string]interface{}
	Post    map[string]interface{}
	Cookie  map[string]interface{}
	Header  map[string]interface{}
	Server  struct {
		IsGET      bool
		IsPOST     bool
		IsAJAX     bool
		IsHTTP     bool
		IsHTTPS    bool
		IsTCP      bool
		IsUDP      bool
		HostName   string
		Port       string
		Path       string
		Query      string
		Referer    string
		UserAgent  string
		RemoteAddr string
	}
}
