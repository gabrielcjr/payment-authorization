package dto

type TransactionDto struct {
	ID              string  `json:"transaction_id"`
	ClientID        string  `json:"client_id"`
	AuthorizationID string  `json:"authorization_id"`
	Status          string  `json:"status"`
	Value           float32 `json:"value"`
	DeniedAt        string  `json:"denied_at"`
	ApprovedAt      string  `json:"approved_at"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
	DeletedAt       string  `json:"deleted_at"`
}

type TransactionCreateDto struct {
	ID       string  `json:"transaction_id"`
	ClientID string  `json:"client_id"`
	Value    float32 `json:"value"`
}

type TransactionFindDto struct {
	ID string `json:"transaction_id"`
}

type TransactionUpdateDto struct {
	ID    string  `json:"transaction_id"`
	Value float32 `json:"value"`
}

type TransactionDeleteDto struct {
	ID string `json:"transaction_id"`
}

type TransactionFindAllResponseDto struct {
	Transactions []*TransactionDto
}

type TransactionAuthorizeDto struct {
	ID              string  `json:"transaction_id"`
	ClientID        string  `json:"client_id"`
	AuthorizationId string  `json:"authorization_id"`
	Value           float32 `json:"value"`
}
