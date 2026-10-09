package repository

// SimplePagination describes pagination block in widget response.
type SimplePagination struct {
	CurrentPage  int  `json:"current_page"`
	PerPage      int  `json:"per_page"`
	HasMorePages bool `json:"has_more_pages"`
	From         int  `json:"from"`
	To           int  `json:"to"`
}

// WidgetLabel is a product label.
type WidgetLabel struct {
	ID      string `json:"id"`
	Label   string `json:"label"`
	Color   string `json:"color"`
	BGColor string `json:"bg_color"`
}

// WidgetProduct is a product item inside a widget.
type WidgetProduct struct {
	ID              int           `json:"id"`
	Name            string        `json:"name"`
	Images          []string      `json:"images"`
	Slug            string        `json:"slug"`
	MinPrice        string        `json:"min_price"`
	Discount        string        `json:"discount"`
	DiscountPercent int           `json:"discount_percent"`
	DefaultDuration int           `json:"default_duration"`
	FinalPrice      string        `json:"final_price"`
	MaxCommission   int           `json:"max_commission"`
	MonthlyPayment  string        `json:"monthly_payment"`
	Gifts           []any         `json:"gifts"`
	Labels          []WidgetLabel `json:"labels"`
	Rating          string        `json:"rating"`
	RatingCount     int           `json:"rating_count"`
	WishlistID      any           `json:"wishlist_id"`
}

// WidgetItem is one widget in the response list.
type WidgetItem struct {
	ID          int             `json:"id"`
	CityID      int             `json:"city_id"`
	Name        string          `json:"name"`
	Title       string          `json:"title"`
	IsHidden    int             `json:"is_hidden_title"`
	WidgetType  string          `json:"widget_type"`
	ViewMode    string          `json:"view_mode"`
	Source      string          `json:"source"`
	WidgetItems []WidgetProduct `json:"widget_items"`
	ViewAllURL  string          `json:"view_all_url"`
	Position    float64         `json:"position"`
	Properties  map[string]any  `json:"properties"`
}

// WidgetsResponse is the wrapper for widget object.
type WidgetsResponse struct {
	SimplePagination SimplePagination `json:"simple_pagination"`
	Items            []WidgetItem     `json:"items"`
}

// WidgetsResult is return structure for widgets endpoint.
type WidgetsResult struct {
	Meta     Meta            `json:"meta"`
	Response WidgetsResponse `json:"response"`
}

type WidgetsRepository interface {
	GetWidgets() WidgetsResult
}

type widgetsRepository struct{}

func NewWidgets() WidgetsRepository {
	return &widgetsRepository{}
}

func (r *widgetsRepository) GetWidgets() WidgetsResult {
	return WidgetsResult{
		Meta: Meta{
			Error:      false,
			Message:    "",
			StatusCode: 200,
		},
		Response: WidgetsResponse{
			SimplePagination: SimplePagination{
				CurrentPage:  1,
				PerPage:      4,
				HasMorePages: true,
				From:         1,
				To:           13,
			},
			Items: []WidgetItem{
				{
					ID:         9004,
					CityID:     1,
					Name:       "Recent_6",
					Title:      "Recent_6",
					IsHidden:   0,
					WidgetType: "product_rec",
					ViewMode:   "web",
					Source:     "main",
					WidgetItems: []WidgetProduct{
						{
							ID:              10006,
							Name:            "Samsung Galaxy A6 5G 12/512 ГБ, серый",
							Images:          []string{"https://storage.alifshop.tj/media/images/alifshop/10006/image-6-1.png", "https://storage.alifshop.tj/media/images/alifshop/10006/image-6-2.png"},
							Slug:            "samsung-galaxy-a6-5g-12-512-gb-seryy",
							MinPrice:        "5506.00",
							Discount:        "0.00",
							DiscountPercent: 0,
							DefaultDuration: 6,
							FinalPrice:      "5506.00",
							MaxCommission:   12,
							MonthlyPayment:  "1036.44",
							Gifts:           []any{},
							Labels: []WidgetLabel{{
								ID:      "new",
								Label:   "Новинка",
								Color:   "#ffffff",
								BGColor: "#9833FD",
							}},
							Rating:      "5.0",
							RatingCount: 6,
							WishlistID:  nil,
						},
						{
							ID:              10007,
							Name:            "Samsung Galaxy A7 5G 12/512 ГБ, серый",
							Images:          []string{"https://storage.alifshop.tj/media/images/alifshop/10007/image-7-1.png", "https://storage.alifshop.tj/media/images/alifshop/10007/image-7-2.png"},
							Slug:            "samsung-galaxy-a7-5g-12-512-gb-seryy",
							MinPrice:        "5507.00",
							Discount:        "0.00",
							DiscountPercent: 0,
							DefaultDuration: 6,
							FinalPrice:      "5507.00",
							MaxCommission:   12,
							MonthlyPayment:  "1038.12",
							Gifts:           []any{},
							Labels: []WidgetLabel{{
								ID:      "new",
								Label:   "Новинка",
								Color:   "#ffffff",
								BGColor: "#9833FD",
							}},
							Rating:      "5.0",
							RatingCount: 7,
							WishlistID:  nil,
						},
						{
							ID:              10008,
							Name:            "Samsung Galaxy A8 5G 12/512 ГБ, серый",
							Images:          []string{"https://storage.alifshop.tj/media/images/alifshop/10008/image-8-1.png", "https://storage.alifshop.tj/media/images/alifshop/10008/image-8-2.png"},
							Slug:            "samsung-galaxy-a8-5g-12-512-gb-seryy",
							MinPrice:        "5508.00",
							Discount:        "0.00",
							DiscountPercent: 0,
							DefaultDuration: 6,
							FinalPrice:      "5508.00",
							MaxCommission:   12,
							MonthlyPayment:  "1039.80",
							Gifts:           []any{},
							Labels: []WidgetLabel{{
								ID:      "original",
								Label:   "Оригинал",
								Color:   "#336BFD",
								BGColor: "#F0F9FF",
							}},
							Rating:      "4.9",
							RatingCount: 8,
							WishlistID:  nil,
						},
						{
							ID:              10009,
							Name:            "Samsung Galaxy A9 5G 12/512 ГБ, серый",
							Images:          []string{"https://storage.alifshop.tj/media/images/alifshop/10009/image-9-1.png", "https://storage.alifshop.tj/media/images/alifshop/10009/image-9-2.png"},
							Slug:            "samsung-galaxy-a9-5g-12-512-gb-seryy",
							MinPrice:        "5509.00",
							Discount:        "0.00",
							DiscountPercent: 0,
							DefaultDuration: 6,
							FinalPrice:      "5509.00",
							MaxCommission:   12,
							MonthlyPayment:  "1041.48",
							Gifts:           []any{},
							Labels: []WidgetLabel{{
								ID:      "new",
								Label:   "Новинка",
								Color:   "#ffffff",
								BGColor: "#9833FD",
							}},
							Rating:      "5.0",
							RatingCount: 9,
							WishlistID:  nil,
						},
						{
							ID:              10010,
							Name:            "Samsung Galaxy A10 5G 12/512 ГБ, серый",
							Images:          []string{"https://storage.alifshop.tj/media/images/alifshop/10010/image-10-1.png", "https://storage.alifshop.tj/media/images/alifshop/10010/image-10-2.png"},
							Slug:            "samsung-galaxy-a10-5g-12-512-gb-seryy",
							MinPrice:        "5510.00",
							Discount:        "0.00",
							DiscountPercent: 0,
							DefaultDuration: 6,
							FinalPrice:      "5510.00",
							MaxCommission:   12,
							MonthlyPayment:  "1043.16",
							Gifts:           []any{},
							Labels: []WidgetLabel{{
								ID:      "new",
								Label:   "Новинка",
								Color:   "#ffffff",
								BGColor: "#9833FD",
							}},
							Rating:      "5.0",
							RatingCount: 10,
							WishlistID:  nil,
						},
						{
							ID:              10011,
							Name:            "Samsung Galaxy A11 5G 12/512 ГБ, серый",
							Images:          []string{"https://storage.alifshop.tj/media/images/alifshop/10011/image-11-1.png", "https://storage.alifshop.tj/media/images/alifshop/10011/image-11-2.png"},
							Slug:            "samsung-galaxy-a11-5g-12-512-gb-seryy",
							MinPrice:        "5511.00",
							Discount:        "0.00",
							DiscountPercent: 0,
							DefaultDuration: 6,
							FinalPrice:      "5511.00",
							MaxCommission:   12,
							MonthlyPayment:  "1044.84",
							Gifts:           []any{},
							Labels: []WidgetLabel{{
								ID:      "original",
								Label:   "Оригинал",
								Color:   "#336BFD",
								BGColor: "#F0F9FF",
							}},
							Rating:      "4.8",
							RatingCount: 11,
							WishlistID:  nil,
						},
					},
					ViewAllURL: "/recent/6",
					Position:   0.004,
					Properties: map[string]any{"rec_type": "recent"},
				},
				{
					ID:         1055,
					CityID:     1,
					Name:       "Онлайн супермаркет Моб",
					Title:      "Вкусные подборки 🧺",
					IsHidden:   0,
					WidgetType: "product_query",
					ViewMode:   "mobile",
					Source:     "main",
					WidgetItems: []WidgetProduct{
						{
							ID:              58036,
							Name:            "Сливки Петел Молочные для взбивания 33 %, 1000 г",
							Images:          []string{"https://storage.alifshop.tj/media/images/alifshop/58036/slivki-petmol-kulinarnye-dlya-vzbivaniya-33-1000-g-1779776410909.png", "https://storage.alifshop.tj/media/images/alifshop/58036/slivki-petmol-kulinarnye-dlya-vzbivaniya-33-1000-g-1779776618882.png"},
							Slug:            "slivki-petmol-kulinarnye-dlya-vzbivaniya-33-1000-g",
							MinPrice:        "73.80",
							Discount:        "0.00",
							DiscountPercent: 0,
							DefaultDuration: 6,
							FinalPrice:      "73.80",
							MaxCommission:   0,
							MonthlyPayment:  "12.30",
							Gifts:           []any{},
							Labels: []WidgetLabel{{
								ID:      "new",
								Label:   "Новинка",
								Color:   "#ffffff",
								BGColor: "#9833FD",
							}},
							Rating:      "5.0",
							RatingCount: 2,
							WishlistID:  nil,
						},
						{
							ID:              43174,
							Name:            "Какао-порошок Mix Fix, 350 г",
							Images:          []string{"https://storage.alifshop.tj/media/images/alifshop/43174/kakao-poroshok-mix-fix-375-g-1785409170135-EuA2CWui.png"},
							Slug:            "kakao-poroshok-mix-fix-350-g",
							MinPrice:        "33.00",
							Discount:        "0.00",
							DiscountPercent: 0,
							DefaultDuration: 6,
							FinalPrice:      "33.00",
							MaxCommission:   0,
							MonthlyPayment:  "5.50",
							Gifts:           []any{},
							Labels:          []WidgetLabel{},
							Rating:          "5.0",
							RatingCount:     2,
							WishlistID:      nil,
						},
					},
					ViewAllURL: "/groceries?cityId=1",
					Position:   0.0032050974701363,
					Properties: map[string]any(nil),
				},
				{
					ID:         1244,
					CityID:     1,
					Name:       "Gigashop.tj",
					Title:      "Техносайд 🔥",
					IsHidden:   0,
					WidgetType: "modified_product",
					ViewMode:   "mobile",
					Source:     "main",
					WidgetItems: []WidgetProduct{
						{
							ID:              67397,
							Name:            "Игра для Sony Playstation 5 FC 27",
							Slug:            "igra-dlya-sony-playstation-5-fc-27",
							Images:          []string{"https://storage.alifshop.tj/media/images/alifshop/67397/igra-dlya-sony-playstation-5-fc-27-1790250690723-MUhEAuIs.png"},
							MinPrice:        "750.00",
							Discount:        "152.00",
							DiscountPercent: 20,
							FinalPrice:      "598.00",
							Gifts:           []any{},
							Labels: []WidgetLabel{{
								ID:      "new",
								Label:   "Новинка",
								Color:   "#ffffff",
								BGColor: "#9833FD",
							}, {
								ID:      "original",
								Label:   "Оригинал",
								Color:   "#336BFD",
								BGColor: "#F0F9FF",
							}},
							DefaultDuration: 12,
							MaxCommission:   20,
							MonthlyPayment:  "59.80",
							Rating:          "5.0",
							RatingCount:     2,
							WishlistID:      nil,
						},
					},
					ViewAllURL: "/shops/gigashoptj?cityId=1&activeTab=products",
					Position:   0.0032051020414136,
					Properties: map[string]any{"background_color": "e3d9db"},
				},
			},
		},
	}
}
