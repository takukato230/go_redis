package types

type (
	SetRequest struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	SetResponse struct {
		Code string `json:"code"`
	}

	GetRequest struct {
		ID string `json:"id"`
	}

	GetResponse struct {
		Name string `json:"name"`
		Code string `json:"code"`
	}
)
