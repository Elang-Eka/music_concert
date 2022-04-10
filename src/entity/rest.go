package entity

type HTTPResp struct {
	Message HTTPMessage `json:"message"`
	Meta    Meta        `json:"metadata"`
	Data    interface{} `json:"data,omitempty"`
}

type HTTPMessage struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Meta struct {
	Path       string     `json:"path"`
	StatusCode int        `json:"statusCode"`
	Status     string     `json:"status"`
	Message    string     `json:"message"`
	Timestamp  string     `json:"timestamp"`
	Error      *MetaError `json:"error,omitempty"`
}

type MetaError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
