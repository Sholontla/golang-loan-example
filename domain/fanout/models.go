package fanout

import "github.com/google/uuid"

type Worker1 struct {
	ID1          uuid.UUID `json:"id1"`
	Name1        string    `json:"name1"`
	Description1 string    `json:"description1"`
	FloatNum1    float64   `json:"float_num1"`
	CreatedAt1   string    `json:"created_at1"`
}

type Worker2 struct {
	ID2          uuid.UUID `json:"id2"`
	Name2        string    `json:"name2"`
	Description2 string    `json:"description2"`
	FloatNum2    float64   `json:"float_num2"`
	CreatedAt2   string    `json:"created_at2"`
}

type Worker3 struct {
	ID3          uuid.UUID `json:"id3"`
	Name3        string    `json:"name3"`
	Description3 string    `json:"description3"`
	FloatNum3    float64   `json:"float_num3"`
	CreatedAt3   string    `json:"created_at3"`
}

type Worker4 struct {
	ID4          uuid.UUID `json:"id4"`
	Name4        string    `json:"name4"`
	Description4 string    `json:"description4"`
	FloatNum4    float64   `json:"float_num4"`
	CreatedAt4   string    `json:"created_at4"`
}

type Workers struct {
	Name1        string  `json:"name1"`
	Description1 string  `json:"description1"`
	FloatNum1    float64 `json:"float_num1"`
	CreatedAt1   string  `json:"created_at1"`
	Name2        string  `json:"name2"`
	Description2 string  `json:"description2"`
	FloatNum2    float64 `json:"float_num2"`
	CreatedAt2   string  `json:"created_at2"`
	Name3        string  `json:"name3"`
	Description3 string  `json:"description3"`
	FloatNum3    float64 `json:"float_num3"`
	CreatedAt3   string  `json:"created_at3"`
	Name4        string  `json:"name4"`
	Description4 string  `json:"description4"`
	FloatNum4    float64 `json:"float_num4"`
	CreatedAt4   string  `json:"created_at4"`
}

type PassArgs struct {
	Activate string `json:"activate"`
}
