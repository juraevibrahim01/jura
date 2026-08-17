package repository

type TestResponse struct {
	ID              int      `json:"id"`
	Name            string   `json:"name"`
	Slug            string   `json:"slug"`
	Rating          string   `json:"rating"`
	RatingCount     int      `json:"rating_count"`
	MinPrice        string   `json:"min_price"`
	FinalPrice      string   `json:"final_price"`
	DefaultDuration int      `json:"default_duration"`
	MaxCommission   int      `json:"max_commission"`
	MaxConditionID  int      `json:"max_condition_id"`
	Images          []string `json:"images"`
	WishlistID      *int     `json:"wishlist_id"`
	Discount        string   `json:"discount"`
	Gifts           []any    `json:"gifts"`
	Labels          []Label  `json:"labels"`
	DiscountPercent int      `json:"discount_percent"`
	MonthlyPayment  string   `json:"monthly_payment"`
}

type Label struct {
	ID      string `json:"id"`
	Label   string `json:"label"`
	Color   string `json:"color"`
	BgColor string `json:"bg_color"`
}

type Meta struct {
	Error      bool   `json:"error"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

type TestResult struct {
	Meta     Meta           `json:"meta"`
	Response []TestResponse `json:"response"`
}

type TestRepository interface {
	GetTest() TestResult
	GetTestNull() TestResult
}

type testRepository struct{}

func NewForYou() TestRepository {
	return &testRepository{}
}

func (r *testRepository) GetTest() TestResult {
	return TestResult{
		Meta: Meta{
			Error:      false,
			Message:    "",
			StatusCode: 200,
		},

		Response: []TestResponse{
			{
				ID:              62729,
				Name:            "Диван Артkор Пуфак 290х105 см, бежевый",
				Slug:            "divan-artkor-pufak-290h105-sm-bezhevyy",
				Rating:          "0.0",
				RatingCount:     0,
				MinPrice:        "12800.00",
				FinalPrice:      "12800.00",
				DefaultDuration: 24,
				MaxCommission:   37,
				MaxConditionID:  25296,
				Images: []string{
					"https://storage.alifshop.tj/media/images/alifshop/62729/divan-artkor-pufak-290h105-sm-bezhevyy-1784725701437-ttuxWaJO.png",
					"https://storage.alifshop.tj/media/images/alifshop/62729/divan-artkor-pufak-290h105-sm-bezhevyy-1784725692887-V4MQUWt6.png",
				},
				WishlistID: nil,
				Discount:   "0.00",
				Gifts:      []any{},
				Labels: []Label{
					{
						ID:      "new",
						Label:   "Новинка",
						Color:   "#ffffff",
						BgColor: "#9833FD",
					},
				},
				DiscountPercent: 0,
				MonthlyPayment:  "730.66",
			}, {
				ID:              62730,
				Name:            "Диван Комфорт 250х100 см, серый",
				Slug:            "divan-komfort-250h100-sm-seryy",
				Rating:          "4.5",
				RatingCount:     18,
				MinPrice:        "15000.00",
				FinalPrice:      "13500.00",
				DefaultDuration: 24,
				MaxCommission:   35,
				MaxConditionID:  25297,
				Images: []string{
					"https://storage.alifshop.tj/media/images/alifshop/62730/divan-komfort-250h100-sm-seryy.png",
					"https://storage.alifshop.tj/media/images/alifshop/62730/divan-komfort-250h100-sm-seryy-2.png",
				},
				WishlistID: nil,
				Discount:   "1500.00",
				Gifts:      []any{},
				Labels: []Label{
					{
						ID:      "sale",
						Label:   "Скидка",
						Color:   "#ffffff",
						BgColor: "#FF3B30",
					},
				},
				DiscountPercent: 10,
				MonthlyPayment:  "770.12",
			}, {
				ID:              62731,
				Name:            "Кресло Комфорт 85х90 см, зеленый",
				Slug:            "kreslo-komfort-85h90-sm-zelenyy",
				Rating:          "4.8",
				RatingCount:     32,
				MinPrice:        "5500.00",
				FinalPrice:      "5500.00",
				DefaultDuration: 12,
				MaxCommission:   30,
				MaxConditionID:  25298,
				Images: []string{
					"https://storage.alifshop.tj/media/images/alifshop/62731/kreslo-komfort-85h90-sm-zelenyy.png",
					"https://storage.alifshop.tj/media/images/alifshop/62731/kreslo-komfort-85h90-sm-zelenyy-2.png",
				},
				WishlistID: nil,
				Discount:   "0.00",
				Gifts:      []any{},
				Labels: []Label{
					{
						ID:      "new",
						Label:   "Новинка",
						Color:   "#ffffff",
						BgColor: "#9833FD",
					},
				},
				DiscountPercent: 0,
				MonthlyPayment:  "495.00",
			}, {
				ID:              62730,
				Name:            "Диван Комфорт 250х100 см, серый",
				Slug:            "divan-komfort-250h100-sm-seryy",
				Rating:          "4.5",
				RatingCount:     18,
				MinPrice:        "15000.00",
				FinalPrice:      "13500.00",
				DefaultDuration: 24,
				MaxCommission:   35,
				MaxConditionID:  25297,
				Images: []string{
					"https://storage.alifshop.tj/media/images/alifshop/62730/divan-komfort-250h100-sm-seryy.png",
					"https://storage.alifshop.tj/media/images/alifshop/62730/divan-komfort-250h100-sm-seryy-2.png",
				},
				WishlistID: nil,
				Discount:   "1500.00",
				Gifts:      []any{},
				Labels: []Label{
					{
						ID:      "sale",
						Label:   "Скидка",
						Color:   "#ffffff",
						BgColor: "#FF3B30",
					},
				},
				DiscountPercent: 10,
				MonthlyPayment:  "770.12",
			},
		},
	}
}

func (r *testRepository) GetTestNull() TestResult {
	return TestResult{
		Meta: Meta{
			Error:      false,
			Message:    "",
			StatusCode: 200,
		},

		Response: nil,
	}
}
