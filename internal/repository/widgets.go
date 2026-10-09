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
    "meta": {
        "error": false,
        "message": "",
        "statusCode": 200
    },
    "response": {
        "simple_pagination": {
            "current_page": 1,
            "per_page": 4,
            "has_more_pages": true,
            "from": 1,
            "to": 13
        },
        "items": [
			{
                "id": 1055,
                "city_id": 1,
                "name": "\u041e\u043d\u043b\u0430\u0439\u043d \u0441\u0443\u043f\u0435\u0440\u043c\u0430\u043a\u0435\u0442 \u041c\u043e\u0431",
                "title": "\u0412\u043a\u0443\u0441\u043d\u044b\u0435 \u043f\u043e\u043a\u0443\u043f\u043a\u0438 \ud83e\uddfa",
                "is_hidden_title": 0,
                "widget_type": "product_query",
                "view_mode": "mobile",
                "source": "main",
                "widget_items": [
                    {
                        "id": 58036,
                        "name": "\u0421\u043b\u0438\u0432\u043a\u0438 \u041f\u0435\u0442\u043c\u043e\u043b \u043a\u0443\u043b\u0438\u043d\u0430\u0440\u043d\u044b\u0435 \u0434\u043b\u044f \u0432\u0437\u0431\u0438\u0432\u0430\u043d\u0438\u044f 33 %, 1000 \u0433",
                        "images": [
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/58036\/slivki-petmol-kulinarnye-dlya-vzbivaniya-33-1000-g-1779776410909.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/58036\/slivki-petmol-kulinarnye-dlya-vzbivaniya-33-1000-g-1779776618882.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/58036\/slivki-petmol-kulinarnye-dlya-vzbivaniya-33-1000-g-1779776408728.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/58036\/slivki-petmol-kulinarnye-dlya-vzbivaniya-33-1000-g-1779776406216.png"
                        ],
                        "slug": "slivki-petmol-kulinarnye-dlya-vzbivaniya-33-1000-g",
                        "min_price": "73.80",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "73.80",
                        "max_commission": 0,
                        "monthly_payment": "12.30",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "\u041d\u043e\u0432\u0438\u043d\u043a\u0430",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 2,
                        "wishlist_id": null
                    },
                    {
                        "id": 43174,
                        "name": "\u041a\u0430\u043a\u0430\u043e-\u043f\u043e\u0440\u043e\u0448\u043e\u043a Mix Fix, 350 \u0433",
                        "images": [
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/43174\/kakao-poroshok-mix-fix-375-g-1785409170135-EuA2CWui.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/43174\/kakao-poroshok-mix-fix-375-g-1785409170232-9SztdUfK.png"
                        ],
                        "slug": "kakao-poroshok-mix-fix-350-g",
                        "min_price": "33.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "33.00",
                        "max_commission": 0,
                        "monthly_payment": "5.50",
                        "gifts": [],
                        "labels": [],
                        "rating": "5.0",
                        "rating_count": 2,
                        "wishlist_id": null
                    },
                    {
                        "id": 41317,
                        "name": "\u0421\u0443\u0445\u0430\u0440\u0438\u043a\u0438 Flint \u0441\u043e \u0432\u043a\u0443\u0441\u043e\u043c \"\u0421\u043c\u0435\u0442\u0430\u043d\u0430 \u0438 \u0437\u0435\u043b\u0435\u043d\u044c\", 100 \u0433 ",
                        "images": [
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/41317\/suhariki-flint-so-vkusom-smetana-i-zelen-100-g-1758985271195.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/41317\/suhariki-flint-so-vkusom-smetana-i-zelen-100-g-1758985273370.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/41317\/suhariki-flint-so-vkusom-smetana-i-zelen-100-g-1758985275493.png"
                        ],
                        "slug": "suhariki-flint-so-vkusom-smetana-i-zelen-100-g",
                        "min_price": "6.80",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "6.80",
                        "max_commission": 0,
                        "monthly_payment": "1.13",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "\u041d\u043e\u0432\u0438\u043d\u043a\u0430",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "0.0",
                        "rating_count": 0,
                        "wishlist_id": null
                    },
                    {
                        "id": 63558,
                        "name": "\u042f\u0439\u0446\u0430 \u041f\u043e\u0440\u0441\u0438\u0451\u043d \u0422\u0443\u0445\u043c\u0438 \u043c\u0443\u0440\u0493\u0438 \u0442\u043e\u0437\u0430, 12 \u0448\u0442.",
                        "images": [
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/63558\/yayca-porsien-t-hmi-mur-i-toza-12-sht-1786363722739-zK67wQjI.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/63558\/yayca-porsien-t-hmi-mur-i-toza-12-sht-1786363802396-VaKRnOpk.png"
                        ],
                        "slug": "yayca-porsien-tuhmi-mur-i-toza-12-sht",
                        "min_price": "27.25",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "27.25",
                        "max_commission": 0,
                        "monthly_payment": "4.54",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "\u041d\u043e\u0432\u0438\u043d\u043a\u0430",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 3,
                        "wishlist_id": 143345
                    },
                    {
                        "id": 55594,
                        "name": "\u0414\u0440\u0430\u0436\u0435 \u0441 \u043c\u043e\u043b\u043e\u0447\u043d\u044b\u043c \u0448\u043e\u043a\u043e\u043b\u0430\u0434\u043e\u043c M&M, 45 \u0433   ",
                        "images": [
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/55594\/drazhe-s-molochnym-shokoladom-m-m-45-g-1775817822616.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/55594\/drazhe-s-molochnym-shokoladom-m-m-45-g-1775817821402.jpg",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/55594\/drazhe-s-molochnym-shokoladom-m-m-45-g-1775817820516.jpg",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/55594\/drazhe-s-molochnym-shokoladom-m-m-45-g-1775817819654.jpg"
                        ],
                        "slug": "drazhe-s-molochnym-shokoladom-m-m-45-g",
                        "min_price": "7.90",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "7.90",
                        "max_commission": 0,
                        "monthly_payment": "1.31",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "\u041d\u043e\u0432\u0438\u043d\u043a\u0430",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 1,
                        "wishlist_id": null
                    },
                    {
                        "id": 63483,
                        "name": "\u041a\u0438\u043b\u044c\u043a\u0430 \u043e\u0431\u0436\u0430\u0440\u0435\u043d\u043d\u0430\u044f \u0417\u0430 \u0420\u043e\u0434\u0438\u043d\u0443 \u0432 \u0442\u043e\u043c\u0430\u0442\u043d\u043e\u043c \u0441\u043e\u0443\u0441\u0435, 175 \u0433",
                        "images": [
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/63483\/kilka-obzharennaya-za-rodinu-v-tomatnom-souse-175-g-1786344743215-3EsIMuW2.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/63483\/kilka-obzharennaya-za-rodinu-v-tomatnom-souse-175-g-1786344743273-zMyC7qav.png"
                        ],
                        "slug": "kilka-obzharennaya-za-rodinu-v-tomatnom-souse-175-g",
                        "min_price": "18.80",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "18.80",
                        "max_commission": 0,
                        "monthly_payment": "3.13",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "\u041d\u043e\u0432\u0438\u043d\u043a\u0430",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 1,
                        "wishlist_id": null
                    },
                    {
                        "id": 44269,
                        "name": "\u0421\u0433\u0443\u0449\u0451\u043d\u043a\u0430 \u0441 \u0441\u0430\u0445\u0430\u0440\u043e\u043c \u041c\u0430\u0440\u0438\u041c\u043e\u043b\u043e\u043a\u043e 8.5%, 1450 \u0433 ",
                        "images": [
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/44269\/sgushchenka-s-saharom-marimoloko-8-5-1450-g-1764653863245.png"
                        ],
                        "slug": "sgushchenka-s-saharom-marimoloko-8-5-1450-g",
                        "min_price": "36.40",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "36.40",
                        "max_commission": 0,
                        "monthly_payment": "6.06",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "\u041d\u043e\u0432\u0438\u043d\u043a\u0430",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "0.0",
                        "rating_count": 0,
                        "wishlist_id": null
                    },
                    {
                        "id": 64137,
                        "name": "\u0410\u043d\u0430\u043d\u0430\u0441\u044b \u043a\u043e\u043b\u044c\u0446\u0430\u043c\u0438 \u0432 \u0441\u0438\u0440\u043e\u043f\u0435 \u0411\u0430\u0440\u043a\u043e, 580 \u043c\u043b",
                        "images": [
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/64137\/ananas-kolcami-v-sirope-barko-580-ml-1787222695971-i6EnDz52.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/64137\/ananas-kolcami-v-sirope-barko-580-ml-1787222695704-3mczBJi2.png"
                        ],
                        "slug": "ananasy-kolcami-v-sirope-barko-580-ml",
                        "min_price": "26.20",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "26.20",
                        "max_commission": 0,
                        "monthly_payment": "4.36",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "\u041d\u043e\u0432\u0438\u043d\u043a\u0430",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 1,
                        "wishlist_id": null
                    },
                    {
                        "id": 63477,
                        "name": "\u0421\u043b\u0438\u0432\u043e\u0447\u043d\u043e\u0435 \u043c\u0430\u0441\u043b\u043e Valio FIN 82%, 200 \u0433",
                        "images": [
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/63477\/slivochnoe-maslo-valio-tradicionnoe-82-200-g-1786340994277-y6MGRT3X.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/63477\/slivochnoe-maslo-valio-fin-82-500-g-copy-1786340737973-SbTAXnEI.png"
                        ],
                        "slug": "slivochnoe-maslo-valio-fin-82-200-g",
                        "min_price": "41.70",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "41.70",
                        "max_commission": 0,
                        "monthly_payment": "6.95",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "\u041d\u043e\u0432\u0438\u043d\u043a\u0430",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "4.5",
                        "rating_count": 6,
                        "wishlist_id": null
                    },
                    {
                        "id": 63308,
                        "name": "\u0428\u043e\u043a\u043e\u043b\u0430\u0434 Toblerone \u043c\u043e\u043b\u043e\u0447\u043d\u044b\u0439 \u0441 \u043c\u0435\u0434\u043e\u0432\u043e-\u043c\u0438\u043d\u0434\u0430\u043b\u044c\u043d\u043e\u0439 \u043d\u0443\u0433\u043e\u0439, 100 \u0433",
                        "images": [
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/63308\/shokolad-toblerone-temnyy-s-medovo-mindalnoy-nugoy-100-g-copy-1785848048073-kuv6Hpn6.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/63308\/shokolad-toblerone-temnyy-s-medovo-mindalnoy-nugoy-100-g-copy-1785848048075-LlBoYw72.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/63308\/shokolad-toblerone-temnyy-s-medovo-mindalnoy-nugoy-100-g-copy-1785848048076-Lgayf9Y0.png"
                        ],
                        "slug": "shokolad-toblerone-molochnyy-s-medovo-mindalnoy-nugoy-100-g",
                        "min_price": "20.10",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "20.10",
                        "max_commission": 0,
                        "monthly_payment": "3.35",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "\u041d\u043e\u0432\u0438\u043d\u043a\u0430",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 3,
                        "wishlist_id": null
                    },
                    {
                        "id": 61737,
                        "name": "\u0416\u0435\u0432\u0430\u0442\u0435\u043b\u044c\u043d\u0430\u044f \u0440\u0435\u0437\u0438\u043d\u043a\u0430 Five \u0421\u043b\u0430\u0434\u043a\u0438\u0435 \u044f\u0433\u043e\u0434\u044b \u0431\u0435\u0437 \u0441\u0430\u0445\u0430\u0440\u0430, 31.2\u0433",
                        "images": [
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/61737\/zhevatelnaya-rezinka-five-sladkie-yagody-bez-sahara-31-2g-1783490790094-ZtygkLDG.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/61737\/zhevatelnaya-rezinka-chupa-chups-big-babol-so-vkusom-banana-21-g-copy-1783489674519-ZRWDSmDJ.png"
                        ],
                        "slug": "zhevatelnaya-rezinka-five-sladkie-yagody-bez-sahara-31-2g",
                        "min_price": "6.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "6.00",
                        "max_commission": 0,
                        "monthly_payment": "1.00",
                        "gifts": [],
                        "labels": [],
                        "rating": "5.0",
                        "rating_count": 3,
                        "wishlist_id": null
                    },
                    {
                        "id": 62791,
                        "name": "\u0421\u0430\u0445\u0430\u0440 \u0440\u0430\u0444\u0438\u043d\u0430\u0434 \u0411\u0430\u0440\u0430\u043a\u0430\u0442 \u0431\u044b\u0441\u0442\u0440\u043e\u0440\u0430\u0441\u0442, 250 \u0433",
                        "images": [
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/62791\/sahar-rafinad-barakat-bystrorast-1-kg-copy-1784807739921-OagstIad.png"
                        ],
                        "slug": "sahar-rafinad-barakat-bystrorast-250-g",
                        "min_price": "6.60",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "6.60",
                        "max_commission": 0,
                        "monthly_payment": "1.10",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "\u041d\u043e\u0432\u0438\u043d\u043a\u0430",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "0.0",
                        "rating_count": 0,
                        "wishlist_id": null
                    }
                ],
                "view_all_url": "\/groceries?cityId=1",
                "position": 0.0032050974701363,
                "properties": null
            },
            {
                "id": 1244,
                "city_id": 1,
                "name": "Gigashop.tj",
                "title": "\u0422\u0435\u0445\u043d\u043e\u0441\u043a\u0438\u0434\u043a\u0438 \ud83d\udd25",
                "is_hidden_title": 0,
                "widget_type": "modified_product",
                "view_mode": "mobile",
                "source": "main",
                "widget_items": [
                    {
                        "id": 67397,
                        "name": "\u0418\u0433\u0440\u0430 \u0434\u043b\u044f Sony Playstation 5 FC 27",
                        "slug": "igra-dlya-sony-playstation-5-fc-27",
                        "images": [
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/67397\/igra-dlya-sony-playstation-5-fc-27-1790250690723-MUhEAuIs.png"
                        ],
                        "min_price": "750.00",
                        "discount": "152.00",
                        "discount_percent": 20,
                        "final_price": "598.00",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "\u041d\u043e\u0432\u0438\u043d\u043a\u0430",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            },
                            {
                                "id": "original",
                                "label": "\u041e\u0440\u0438\u0433\u0438\u043d\u0430\u043b",
                                "color": "#336BFD",
                                "bg_color": "#F0F9FF"
                            }
                        ],
                        "default_duration": 12,
                        "max_commission": 20,
                        "monthly_payment": "59.80",
                        "rating": "5.0",
                        "rating_count": 2,
                        "wishlist_id": null
                    },
                    {
                        "id": 68055,
                        "name": "\u0413\u0435\u0439\u043c\u043f\u0430\u0434 Sony PlayStation 5 DualSense Limited Edition Marvels Wolverine, \u0436\u0451\u043b\u0442\u043e-\u0447\u0451\u0440\u043d\u044b\u0439",
                        "slug": "geympad-sony-playstation-5-dualsense-limited-edition-marvels-wolverine-zhelto-chernyy",
                        "images": [
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/68055\/geympad-sony-playstation-5-dualsense-limited-edition-marvels-wolverine-zhelto-chernyy-1790675388539-SuljQ6E8.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/68055\/geympad-sony-playstation-5-dualsense-limited-edition-marvels-wolverine-zhelto-chernyy-1790675388170-CprpF9g3.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/68055\/geympad-sony-playstation-5-dualsense-limited-edition-marvels-wolverine-zhelto-chernyy-1790675388170-4i3WmFKo.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/68055\/geympad-sony-playstation-5-dualsense-limited-edition-marvels-wolverine-zhelto-chernyy-1790675388339-y3gAYaY7.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/68055\/geympad-sony-playstation-5-dualsense-limited-edition-marvels-wolverine-zhelto-chernyy-1790675388242-WlB6I96J.png"
                        ],
                        "min_price": "1200.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "final_price": "1200.00",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "\u041d\u043e\u0432\u0438\u043d\u043a\u0430",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            },
                            {
                                "id": "original",
                                "label": "\u041e\u0440\u0438\u0433\u0438\u043d\u0430\u043b",
                                "color": "#336BFD",
                                "bg_color": "#F0F9FF"
                            }
                        ],
                        "default_duration": 12,
                        "max_commission": 20,
                        "monthly_payment": "120.00",
                        "rating": "0.0",
                        "rating_count": 0,
                        "wishlist_id": null
                    },
                    {
                        "id": 19142,
                        "name": "\u0418\u0433\u0440\u043e\u0432\u0430\u044f \u043f\u0440\u0438\u0441\u0442\u0430\u0432\u043a\u0430 Sony PlayStation 5 Slim, 1000 GB",
                        "slug": "igrovaya-pristavka-sony-playstation-5-slim-1000-gb",
                        "images": [
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/19142\/igrovaya-pristavka-sony-playstation-5-slim-1000-gb-1766996515659.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/19142\/igrovaya-pristavka-sony-playstation-5-slim-1000-gb-1766996516371.png"
                        ],
                        "min_price": "7950.00",
                        "discount": "360.00",
                        "discount_percent": 4,
                        "final_price": "7590.00",
                        "gifts": [
                            271959
                        ],
                        "labels": [
                            {
                                "id": "original",
                                "label": "\u041e\u0440\u0438\u0433\u0438\u043d\u0430\u043b",
                                "color": "#336BFD",
                                "bg_color": "#F0F9FF"
                            }
                        ],
                        "default_duration": 24,
                        "max_commission": 30,
                        "monthly_payment": "411.12",
                        "rating": "5.0",
                        "rating_count": 45,
                        "wishlist_id": null
                    },
                    {
                        "id": 10370,
                        "name": "\u0413\u0435\u0439\u043c\u043f\u0430\u0434 Sony DualSense, \u0431\u0435\u043b\u044b\u0439",
                        "slug": "geympad-sony-dualsense-belyy",
                        "images": [
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/10370\/geympad-sony-dualsense-belyy-1767001387243.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/10370\/geympad-sony-dualsense-belyy-1767001392904.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/10370\/geympad-sony-dualsense-belyy-1767001394133.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/10370\/geympad-sony-dualsense-belyy-1767001388572.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/10370\/geympad-sony-dualsense-belyy-1767001389549.png"
                        ],
                        "min_price": "620.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "final_price": "620.00",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "original",
                                "label": "\u041e\u0440\u0438\u0433\u0438\u043d\u0430\u043b",
                                "color": "#336BFD",
                                "bg_color": "#F0F9FF"
                            }
                        ],
                        "default_duration": 12,
                        "max_commission": 20,
                        "monthly_payment": "62.00",
                        "rating": "5.0",
                        "rating_count": 15,
                        "wishlist_id": null
                    },
                    {
                        "id": 51442,
                        "name": "\u0411\u0435\u0441\u043f\u0440\u043e\u0432\u043e\u0434\u043d\u044b\u0435 \u043d\u0430\u0443\u0448\u043d\u0438\u043a\u0438 Sony Pulse Elite, \u0447\u0451\u0440\u043d\u044b\u0439",
                        "slug": "besprovodnye-naushniki-sony-pulse-elite-chernyy",
                        "images": [
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/51442\/besprovodnye-naushniki-sony-pulse-elite-chernyy-1771309089996.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/51442\/besprovodnye-naushniki-sony-pulse-elite-chernyy-1771309086443.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/51442\/besprovodnye-naushniki-sony-pulse-elite-chernyy-1771309088986.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/51442\/besprovodnye-naushniki-sony-pulse-elite-chernyy-1771309088104.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/51442\/besprovodnye-naushniki-sony-pulse-elite-chernyy-1771309085651.png"
                        ],
                        "min_price": "1850.00",
                        "discount": "150.00",
                        "discount_percent": 8,
                        "final_price": "1700.00",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "\u041d\u043e\u0432\u0438\u043d\u043a\u0430",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            },
                            {
                                "id": "original",
                                "label": "\u041e\u0440\u0438\u0433\u0438\u043d\u0430\u043b",
                                "color": "#336BFD",
                                "bg_color": "#F0F9FF"
                            }
                        ],
                        "default_duration": 12,
                        "max_commission": 20,
                        "monthly_payment": "170.00",
                        "rating": "0.0",
                        "rating_count": 0,
                        "wishlist_id": null
                    },
                    {
                        "id": 42938,
                        "name": "\u0411\u0435\u0441\u043f\u0440\u043e\u0432\u043e\u0434\u043d\u044b\u0435 \u043d\u0430\u0443\u0448\u043d\u0438\u043a\u0438 Sony PlayStation PULSE Explore, \u0431\u0435\u043b\u043e-\u0447\u0451\u0440\u043d\u044b\u0439",
                        "slug": "besprovodnye-naushniki-sony-playstation-pulse-explore-belo-chernyy",
                        "images": [
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/42938\/besprovodnye-naushniki-sony-playstation-pulse-explore-belo-chernyy-1761935093815.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/42938\/besprovodnye-naushniki-sony-playstation-pulse-explore-belo-chernyy-1761935085861.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/42938\/besprovodnye-naushniki-sony-playstation-pulse-explore-belo-chernyy-1761935091775.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/42938\/besprovodnye-naushniki-sony-playstation-pulse-explore-belo-chernyy-1761935088148.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/42938\/besprovodnye-naushniki-sony-playstation-pulse-explore-belo-chernyy-1761935090046.png"
                        ],
                        "min_price": "1500.00",
                        "discount": "202.00",
                        "discount_percent": 13,
                        "final_price": "1298.00",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "\u041d\u043e\u0432\u0438\u043d\u043a\u0430",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            },
                            {
                                "id": "original",
                                "label": "\u041e\u0440\u0438\u0433\u0438\u043d\u0430\u043b",
                                "color": "#336BFD",
                                "bg_color": "#F0F9FF"
                            }
                        ],
                        "default_duration": 12,
                        "max_commission": 20,
                        "monthly_payment": "129.80",
                        "rating": "5.0",
                        "rating_count": 2,
                        "wishlist_id": null
                    },
                    {
                        "id": 22367,
                        "name": "Meta Quest 3, 512\u0413\u0411, \u0431\u0430\u0437\u043e\u0432\u0430\u044f, \u0431\u0435\u043b\u044b\u0439 ",
                        "slug": "meta-quest-3-512gb-bazovaya-belyy",
                        "images": [
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/22367\/meta-quest-3-512gb-bazovaya-belyy-1766995313832.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/22367\/meta-quest-3-512gb-bazovaya-belyy-1766995315128.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/22367\/meta-quest-3-512gb-bazovaya-belyy-1766995316277.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/22367\/meta-quest-3-512gb-bazovaya-belyy-1766995317666.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/22367\/meta-quest-3-512gb-bazovaya-belyy-1766995318724.png"
                        ],
                        "min_price": "8000.00",
                        "discount": "2150.00",
                        "discount_percent": 26,
                        "final_price": "5850.00",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "original",
                                "label": "\u041e\u0440\u0438\u0433\u0438\u043d\u0430\u043b",
                                "color": "#336BFD",
                                "bg_color": "#F0F9FF"
                            }
                        ],
                        "default_duration": 12,
                        "max_commission": 20,
                        "monthly_payment": "585.00",
                        "rating": "5.0",
                        "rating_count": 2,
                        "wishlist_id": null
                    },
                    {
                        "id": 19775,
                        "name": "\u0418\u0433\u0440\u043e\u0432\u0430\u044f \u043f\u0440\u0438\u0441\u0442\u0430\u0432\u043a\u0430 PlayStation Portal, \u0431\u0435\u043b\u044b\u0439",
                        "slug": "igrovaya-pristavka-playstation-portal-belyy",
                        "images": [
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/19775\/igrovaya-pristavka-playstation-portal-belyy-1766999390151.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/19775\/igrovaya-pristavka-playstation-portal-belyy-1766999391373.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/19775\/igrovaya-pristavka-playstation-portal-belyy-1766999392519.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/19775\/igrovaya-pristavka-playstation-portal-belyy-1766999393457.png"
                        ],
                        "min_price": "2800.00",
                        "discount": "201.00",
                        "discount_percent": 7,
                        "final_price": "2599.00",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "original",
                                "label": "\u041e\u0440\u0438\u0433\u0438\u043d\u0430\u043b",
                                "color": "#336BFD",
                                "bg_color": "#F0F9FF"
                            }
                        ],
                        "default_duration": 24,
                        "max_commission": 30,
                        "monthly_payment": "140.77",
                        "rating": "5.0",
                        "rating_count": 1,
                        "wishlist_id": null
                    },
                    {
                        "id": 27061,
                        "name": "\u0418\u0433\u0440\u043e\u0432\u0430\u044f \u043f\u0440\u0438\u0441\u0442\u0430\u0432\u043a\u0430 Sony PlayStation 5 Pro Digital Edition, 2000 GB ",
                        "slug": "igrovaya-pristavka-sony-playstation-5-pro-digital-edition-2000-gb",
                        "images": [
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/27061\/igrovaya-pristavka-sony-playstation-5-pro-digital-edition-2000-gb-1765361375541.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/27061\/igrovaya-pristavka-sony-playstation-5-pro-digital-edition-2000-gb-1765361377198.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/27061\/igrovaya-pristavka-sony-playstation-5-pro-digital-edition-2000-gb-1765361378642.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/27061\/igrovaya-pristavka-sony-playstation-5-pro-digital-edition-2000-gb-1765361380138.png"
                        ],
                        "min_price": "12500.00",
                        "discount": "801.00",
                        "discount_percent": 6,
                        "final_price": "11699.00",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "\u041d\u043e\u0432\u0438\u043d\u043a\u0430",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            },
                            {
                                "id": "original",
                                "label": "\u041e\u0440\u0438\u0433\u0438\u043d\u0430\u043b",
                                "color": "#336BFD",
                                "bg_color": "#F0F9FF"
                            }
                        ],
                        "default_duration": 24,
                        "max_commission": 30,
                        "monthly_payment": "633.69",
                        "rating": "5.0",
                        "rating_count": 20,
                        "wishlist_id": null
                    },
                    {
                        "id": 18100,
                        "name": "Sony PlayStation VR2 \u0431\u0430\u0437\u043e\u0432\u0430\u044f, \u0431\u0435\u043b\u044b\u0439",
                        "slug": "sony-playstation-vr2-bazovaya-belyy",
                        "images": [
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/18100\/sony-playstation-vr2-bazovaya-belyy-1766995724157.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/18100\/sony-playstation-vr2-bazovaya-belyy-1766995725020.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/18100\/sony-playstation-vr2-bazovaya-belyy-1766995726071.png"
                        ],
                        "min_price": "6000.00",
                        "discount": "2350.00",
                        "discount_percent": 39,
                        "final_price": "3650.00",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "original",
                                "label": "\u041e\u0440\u0438\u0433\u0438\u043d\u0430\u043b",
                                "color": "#336BFD",
                                "bg_color": "#F0F9FF"
                            }
                        ],
                        "default_duration": 12,
                        "max_commission": 20,
                        "monthly_payment": "365.00",
                        "rating": "5.0",
                        "rating_count": 2,
                        "wishlist_id": null
                    },
                    {
                        "id": 27136,
                        "name": "Meta Quest 3S, 128 \u0413\u0411, \u0431\u0430\u0437\u043e\u0432\u0430\u044f, \u0431\u0435\u043b\u044b\u0439",
                        "slug": "meta-quest-3s-128-gb-bazovaya-belyy",
                        "images": [
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/27136\/meta-quest-3s-128-gb-bazovaya-belyy-1766995582784.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/27136\/meta-quest-3s-128-gb-bazovaya-belyy-1766995583647.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/27136\/meta-quest-3s-128-gb-bazovaya-belyy-1766995584427.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/27136\/meta-quest-3s-128-gb-bazovaya-belyy-1766995585321.png"
                        ],
                        "min_price": "4198.00",
                        "discount": "199.00",
                        "discount_percent": 4,
                        "final_price": "3999.00",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "original",
                                "label": "\u041e\u0440\u0438\u0433\u0438\u043d\u0430\u043b",
                                "color": "#336BFD",
                                "bg_color": "#F0F9FF"
                            }
                        ],
                        "default_duration": 12,
                        "max_commission": 20,
                        "monthly_payment": "399.90",
                        "rating": "5.0",
                        "rating_count": 1,
                        "wishlist_id": null
                    },
                    {
                        "id": 21460,
                        "name": "\u0418\u0433\u0440\u043e\u0432\u044b\u0435 \u043d\u0430\u0443\u0448\u043d\u0438\u043a\u0438 Sony Pulse Elite, \u0431\u0435\u043b\u044b\u0439",
                        "slug": "igrovye-naushniki-sony-pulse-elite-belyy",
                        "images": [
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/21460\/igrovye-naushniki-sony-pulse-elite-belyy-1709122779860.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/21460\/igrovye-naushniki-sony-pulse-elite-belyy-1709122780493.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/21460\/igrovye-naushniki-sony-pulse-elite-belyy-1709122780476.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/21460\/igrovye-naushniki-sony-pulse-elite-belyy-1709122779078.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/21460\/igrovye-naushniki-sony-pulse-elite-belyy-1709122776185.png"
                        ],
                        "min_price": "2000.00",
                        "discount": "300.00",
                        "discount_percent": 15,
                        "final_price": "1700.00",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "original",
                                "label": "\u041e\u0440\u0438\u0433\u0438\u043d\u0430\u043b",
                                "color": "#336BFD",
                                "bg_color": "#F0F9FF"
                            }
                        ],
                        "default_duration": 12,
                        "max_commission": 20,
                        "monthly_payment": "170.00",
                        "rating": "5.0",
                        "rating_count": 2,
                        "wishlist_id": null
                    },
                    {
                        "id": 63268,
                        "name": "\u0411\u0435\u0441\u043f\u0440\u043e\u0432\u043e\u0434\u043d\u044b\u0435 \u043d\u0430\u0443\u0448\u043d\u0438\u043a\u0438 Sony PlayStation PULSE Explore, \u0447\u0451\u0440\u043d\u044b\u0439",
                        "slug": "besprovodnye-naushniki-sony-playstation-pulse-explore-chernyy",
                        "images": [
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/63268\/besprovodnye-naushniki-sony-playstation-pulse-explore-chernyy-1785821184574-cT8fqLhf.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/63268\/besprovodnye-naushniki-sony-playstation-pulse-explore-chernyy-1785821184175-bAbrMdwP.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/63268\/besprovodnye-naushniki-sony-playstation-pulse-explore-chernyy-1785821184977-PZUdTtip.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/63268\/besprovodnye-naushniki-sony-playstation-pulse-explore-chernyy-1785821184075-mLKjwku9.png",
                            "https:\/\/storage.alifshop.tj\/media\/images\/alifshop\/63268\/besprovodnye-naushniki-sony-playstation-pulse-explore-chernyy-1785821188382-sXCxcFta.png"
                        ],
                        "min_price": "1400.00",
                        "discount": "111.00",
                        "discount_percent": 7,
                        "final_price": "1289.00",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "\u041d\u043e\u0432\u0438\u043d\u043a\u0430",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            },
                            {
                                "id": "original",
                                "label": "\u041e\u0440\u0438\u0433\u0438\u043d\u0430\u043b",
                                "color": "#336BFD",
                                "bg_color": "#F0F9FF"
                            }
                        ],
                        "default_duration": 12,
                        "max_commission": 20,
                        "monthly_payment": "128.90",
                        "rating": "0.0",
                        "rating_count": 0,
                        "wishlist_id": null
                    }
                ],
                "view_all_url": "\/shops\/gigashoptj?cityId=1&activeTab=products",
                "position": 0.0032051020414136,
                "properties": {
                    "background_color": "e3d9db"
                }
            },
            {
                "id": 9004,
                "city_id": 1,
                "name": "Recent_6",
                "title": "Recent_6",
                "is_hidden_title": 0,
                "widget_type": "product_rec",
                "view_mode": "web",
                "source": "main",
                "widget_items": [
                    {
                        "id": 10006,
                        "name": "Samsung Galaxy A6 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/10006/image-6-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/10006/image-6-2.png"
                        ],
                        "slug": "samsung-galaxy-a6-5g-12-512-gb-seryy",
                        "min_price": "5506.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5506.00",
                        "max_commission": 12,
                        "monthly_payment": "1036.44",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 6,
                        "wishlist_id": null
                    },
                    {
                        "id": 10007,
                        "name": "Samsung Galaxy A7 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/10007/image-7-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/10007/image-7-2.png"
                        ],
                        "slug": "samsung-galaxy-a7-5g-12-512-gb-seryy",
                        "min_price": "5507.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5507.00",
                        "max_commission": 12,
                        "monthly_payment": "1038.12",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 7,
                        "wishlist_id": null
                    },
                    {
                        "id": 10008,
                        "name": "Samsung Galaxy A8 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/10008/image-8-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/10008/image-8-2.png"
                        ],
                        "slug": "samsung-galaxy-a8-5g-12-512-gb-seryy",
                        "min_price": "5508.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5508.00",
                        "max_commission": 12,
                        "monthly_payment": "1039.80",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "original",
                                "label": "Оригинал",
                                "color": "#336BFD",
                                "bg_color": "#F0F9FF"
                            }
                        ],
                        "rating": "4.9",
                        "rating_count": 8,
                        "wishlist_id": null
                    },
                    {
                        "id": 10009,
                        "name": "Samsung Galaxy A9 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/10009/image-9-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/10009/image-9-2.png"
                        ],
                        "slug": "samsung-galaxy-a9-5g-12-512-gb-seryy",
                        "min_price": "5509.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5509.00",
                        "max_commission": 12,
                        "monthly_payment": "1041.48",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 9,
                        "wishlist_id": null
                    },
                    {
                        "id": 10010,
                        "name": "Samsung Galaxy A10 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/10010/image-10-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/10010/image-10-2.png"
                        ],
                        "slug": "samsung-galaxy-a10-5g-12-512-gb-seryy",
                        "min_price": "5510.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5510.00",
                        "max_commission": 12,
                        "monthly_payment": "1043.16",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 10,
                        "wishlist_id": null
                    },
                    {
                        "id": 10011,
                        "name": "Samsung Galaxy A11 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/10011/image-11-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/10011/image-11-2.png"
                        ],
                        "slug": "samsung-galaxy-a11-5g-12-512-gb-seryy",
                        "min_price": "5511.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5511.00",
                        "max_commission": 12,
                        "monthly_payment": "1044.84",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "original",
                                "label": "Оригинал",
                                "color": "#336BFD",
                                "bg_color": "#F0F9FF"
                            }
                        ],
                        "rating": "4.8",
                        "rating_count": 11,
                        "wishlist_id": null
                    }
                ],
                "view_all_url": "/recent/6",
                "position": 0.004,
                "properties": {
                    "rec_type": "recent"
                }
            },
            {
                "id": 9014,
                "city_id": 1,
                "name": "Buy_again_6",
                "title": "Buy_again_6",
                "is_hidden_title": 0,
                "widget_type": "product_rec",
                "view_mode": "web",
                "source": "main",
                "widget_items": [
                    {
                        "id": 20006,
                        "name": "Samsung Galaxy B6 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/20006/image-6-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/20006/image-6-2.png"
                        ],
                        "slug": "samsung-galaxy-b6-5g-12-512-gb-seryy",
                        "min_price": "5606.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5606.00",
                        "max_commission": 12,
                        "monthly_payment": "1056.60",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 6,
                        "wishlist_id": null
                    },
                    {
                        "id": 20007,
                        "name": "Samsung Galaxy B7 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/20007/image-7-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/20007/image-7-2.png"
                        ],
                        "slug": "samsung-galaxy-b7-5g-12-512-gb-seryy",
                        "min_price": "5607.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5607.00",
                        "max_commission": 12,
                        "monthly_payment": "1058.28",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 7,
                        "wishlist_id": null
                    },
                    {
                        "id": 20008,
                        "name": "Samsung Galaxy B8 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/20008/image-8-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/20008/image-8-2.png"
                        ],
                        "slug": "samsung-galaxy-b8-5g-12-512-gb-seryy",
                        "min_price": "5608.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5608.00",
                        "max_commission": 12,
                        "monthly_payment": "1059.96",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "original",
                                "label": "Оригинал",
                                "color": "#336BFD",
                                "bg_color": "#F0F9FF"
                            }
                        ],
                        "rating": "4.9",
                        "rating_count": 8,
                        "wishlist_id": null
                    },
                    {
                        "id": 20009,
                        "name": "Samsung Galaxy B9 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/20009/image-9-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/20009/image-9-2.png"
                        ],
                        "slug": "samsung-galaxy-b9-5g-12-512-gb-seryy",
                        "min_price": "5609.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5609.00",
                        "max_commission": 12,
                        "monthly_payment": "1061.64",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 9,
                        "wishlist_id": null
                    },
                    {
                        "id": 20010,
                        "name": "Samsung Galaxy B10 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/20010/image-10-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/20010/image-10-2.png"
                        ],
                        "slug": "samsung-galaxy-b10-5g-12-512-gb-seryy",
                        "min_price": "5610.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5610.00",
                        "max_commission": 12,
                        "monthly_payment": "1063.32",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 10,
                        "wishlist_id": null
                    },
                    {
                        "id": 20011,
                        "name": "Samsung Galaxy B11 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/20011/image-11-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/20011/image-11-2.png"
                        ],
                        "slug": "samsung-galaxy-b11-5g-12-512-gb-seryy",
                        "min_price": "5611.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5611.00",
                        "max_commission": 12,
                        "monthly_payment": "1065.00",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "original",
                                "label": "Оригинал",
                                "color": "#336BFD",
                                "bg_color": "#F0F9FF"
                            }
                        ],
                        "rating": "4.8",
                        "rating_count": 11,
                        "wishlist_id": null
                    }
                ],
                "view_all_url": "/buy_again/6",
                "position": 0.014,
                "properties": {
                    "rec_type": "buy_again"
                }
            },
            {
                "id": 9031,
                "city_id": 1,
                "name": "Popular_",
                "title": "Popular_",
                "is_hidden_title": 0,
                "widget_type": "product_rec",
                "view_mode": "web",
                "source": "main",
                "view_all_url": "/popular/all",
                "position": 0.031,
                "properties": {
                    "rec_type": "popular"
                }
            },
            {
                "id": 9032,
                "city_id": 1,
                "name": "Popular_1",
                "title": "Popular_1",
                "is_hidden_title": 0,
                "widget_type": "product_rec",
                "view_mode": "web",
                "source": "main",
                "widget_items": [
                    {
                        "id": 40001,
                        "name": "Samsung Galaxy D1 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40001/image-1-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40001/image-1-2.png"
                        ],
                        "slug": "samsung-galaxy-d1-5g-12-512-gb-seryy",
                        "min_price": "5801.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5801.00",
                        "max_commission": 12,
                        "monthly_payment": "1086.84",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 1,
                        "wishlist_id": null
                    }
                ],
                "view_all_url": "/popular/1",
                "position": 0.032,
                "properties": {
                    "rec_type": "popular"
                }
            },
            {
                "id": 9033,
                "city_id": 1,
                "name": "Popular_4",
                "title": "Popular_4",
                "is_hidden_title": 0,
                "widget_type": "product_rec",
                "view_mode": "web",
                "source": "main",
                "widget_items": [
                    {
                        "id": 40002,
                        "name": "Samsung Galaxy D2 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40002/image-2-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40002/image-2-2.png"
                        ],
                        "slug": "samsung-galaxy-d2-5g-12-512-gb-seryy",
                        "min_price": "5802.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5802.00",
                        "max_commission": 12,
                        "monthly_payment": "1088.52",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "4.9",
                        "rating_count": 2,
                        "wishlist_id": null
                    },
                    {
                        "id": 40003,
                        "name": "Samsung Galaxy D3 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40003/image-3-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40003/image-3-2.png"
                        ],
                        "slug": "samsung-galaxy-d3-5g-12-512-gb-seryy",
                        "min_price": "5803.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5803.00",
                        "max_commission": 12,
                        "monthly_payment": "1090.20",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "original",
                                "label": "Оригинал",
                                "color": "#336BFD",
                                "bg_color": "#F0F9FF"
                            }
                        ],
                        "rating": "4.8",
                        "rating_count": 3,
                        "wishlist_id": null
                    },
                    {
                        "id": 40004,
                        "name": "Samsung Galaxy D4 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40004/image-4-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40004/image-4-2.png"
                        ],
                        "slug": "samsung-galaxy-d4-5g-12-512-gb-seryy",
                        "min_price": "5804.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5804.00",
                        "max_commission": 12,
                        "monthly_payment": "1091.88",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 4,
                        "wishlist_id": null
                    },
                    {
                        "id": 40005,
                        "name": "Samsung Galaxy D5 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40005/image-5-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40005/image-5-2.png"
                        ],
                        "slug": "samsung-galaxy-d5-5g-12-512-gb-seryy",
                        "min_price": "5805.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5805.00",
                        "max_commission": 12,
                        "monthly_payment": "1093.56",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 5,
                        "wishlist_id": null
                    }
                ],
                "view_all_url": "/popular/4",
                "position": 0.033,
                "properties": {
                    "rec_type": "popular"
                }
            },
            {
                "id": 9034,
                "city_id": 1,
                "name": "Popular_6",
                "title": "Popular_6",
                "is_hidden_title": 0,
                "widget_type": "product_rec",
                "view_mode": "web",
                "source": "main",
                "widget_items": [
                    {
                        "id": 40006,
                        "name": "Samsung Galaxy D6 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40006/image-6-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40006/image-6-2.png"
                        ],
                        "slug": "samsung-galaxy-d6-5g-12-512-gb-seryy",
                        "min_price": "5806.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5806.00",
                        "max_commission": 12,
                        "monthly_payment": "1095.24",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 6,
                        "wishlist_id": null
                    },
                    {
                        "id": 40007,
                        "name": "Samsung Galaxy D7 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40007/image-7-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40007/image-7-2.png"
                        ],
                        "slug": "samsung-galaxy-d7-5g-12-512-gb-seryy",
                        "min_price": "5807.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5807.00",
                        "max_commission": 12,
                        "monthly_payment": "1096.92",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 7,
                        "wishlist_id": null
                    },
                    {
                        "id": 40008,
                        "name": "Samsung Galaxy D8 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40008/image-8-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40008/image-8-2.png"
                        ],
                        "slug": "samsung-galaxy-d8-5g-12-512-gb-seryy",
                        "min_price": "5808.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5808.00",
                        "max_commission": 12,
                        "monthly_payment": "1098.60",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 8,
                        "wishlist_id": null
                    },
                    {
                        "id": 40009,
                        "name": "Samsung Galaxy D9 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40009/image-9-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40009/image-9-2.png"
                        ],
                        "slug": "samsung-galaxy-d9-5g-12-512-gb-seryy",
                        "min_price": "5809.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5809.00",
                        "max_commission": 12,
                        "monthly_payment": "1100.28",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 9,
                        "wishlist_id": null
                    },
                    {
                        "id": 40010,
                        "name": "Samsung Galaxy D10 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40010/image-10-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40010/image-10-2.png"
                        ],
                        "slug": "samsung-galaxy-d10-5g-12-512-gb-seryy",
                        "min_price": "5810.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5810.00",
                        "max_commission": 12,
                        "monthly_payment": "1101.96",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 10,
                        "wishlist_id": null
                    },
                    {
                        "id": 40011,
                        "name": "Samsung Galaxy D11 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40011/image-11-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40011/image-11-2.png"
                        ],
                        "slug": "samsung-galaxy-d11-5g-12-512-gb-seryy",
                        "min_price": "5811.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5811.00",
                        "max_commission": 12,
                        "monthly_payment": "1103.64",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 11,
                        "wishlist_id": null
                    }
                ],
                "view_all_url": "/popular/6",
                "position": 0.034,
                "properties": {
                    "rec_type": "popular"
                }
            },
            {
                "id": 9035,
                "city_id": 1,
                "name": "Popular_8",
                "title": "Popular_8",
                "is_hidden_title": 0,
                "widget_type": "product_rec",
                "view_mode": "web",
                "source": "main",
                "widget_items": [
                    {
                        "id": 40012,
                        "name": "Samsung Galaxy D12 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40012/image-12-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40012/image-12-2.png"
                        ],
                        "slug": "samsung-galaxy-d12-5g-12-512-gb-seryy",
                        "min_price": "5812.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5812.00",
                        "max_commission": 12,
                        "monthly_payment": "1105.32",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 12,
                        "wishlist_id": null
                    },
                    {
                        "id": 40013,
                        "name": "Samsung Galaxy D13 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40013/image-13-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40013/image-13-2.png"
                        ],
                        "slug": "samsung-galaxy-d13-5g-12-512-gb-seryy",
                        "min_price": "5813.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5813.00",
                        "max_commission": 12,
                        "monthly_payment": "1107.00",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 13,
                        "wishlist_id": null
                    },
                    {
                        "id": 40014,
                        "name": "Samsung Galaxy D14 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40014/image-14-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40014/image-14-2.png"
                        ],
                        "slug": "samsung-galaxy-d14-5g-12-512-gb-seryy",
                        "min_price": "5814.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5814.00",
                        "max_commission": 12,
                        "monthly_payment": "1108.68",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 14,
                        "wishlist_id": null
                    },
                    {
                        "id": 40015,
                        "name": "Samsung Galaxy D15 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40015/image-15-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40015/image-15-2.png"
                        ],
                        "slug": "samsung-galaxy-d15-5g-12-512-gb-seryy",
                        "min_price": "5815.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5815.00",
                        "max_commission": 12,
                        "monthly_payment": "1110.36",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 15,
                        "wishlist_id": null
                    },
                    {
                        "id": 40016,
                        "name": "Samsung Galaxy D16 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40016/image-16-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40016/image-16-2.png"
                        ],
                        "slug": "samsung-galaxy-d16-5g-12-512-gb-seryy",
                        "min_price": "5816.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5816.00",
                        "max_commission": 12,
                        "monthly_payment": "1112.04",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 16,
                        "wishlist_id": null
                    },
                    {
                        "id": 40017,
                        "name": "Samsung Galaxy D17 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40017/image-17-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40017/image-17-2.png"
                        ],
                        "slug": "samsung-galaxy-d17-5g-12-512-gb-seryy",
                        "min_price": "5817.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5817.00",
                        "max_commission": 12,
                        "monthly_payment": "1113.72",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 17,
                        "wishlist_id": null
                    },
                    {
                        "id": 40018,
                        "name": "Samsung Galaxy D18 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40018/image-18-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40018/image-18-2.png"
                        ],
                        "slug": "samsung-galaxy-d18-5g-12-512-gb-seryy",
                        "min_price": "5818.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5818.00",
                        "max_commission": 12,
                        "monthly_payment": "1115.40",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 18,
                        "wishlist_id": null
                    },
                    {
                        "id": 40019,
                        "name": "Samsung Galaxy D19 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40019/image-19-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40019/image-19-2.png"
                        ],
                        "slug": "samsung-galaxy-d19-5g-12-512-gb-seryy",
                        "min_price": "5819.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5819.00",
                        "max_commission": 12,
                        "monthly_payment": "1117.08",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 19,
                        "wishlist_id": null
                    }
                ],
                "view_all_url": "/popular/8",
                "position": 0.035,
                "properties": {
                    "rec_type": "popular"
                }
            },
            {
                "id": 9036,
                "city_id": 1,
                "name": "Popular_12",
                "title": "Popular_12",
                "is_hidden_title": 0,
                "widget_type": "product_rec",
                "view_mode": "web",
                "source": "main",
                "widget_items": [
                    {
                        "id": 40020,
                        "name": "Samsung Galaxy D20 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40020/image-20-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40020/image-20-2.png"
                        ],
                        "slug": "samsung-galaxy-d20-5g-12-512-gb-seryy",
                        "min_price": "5820.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5820.00",
                        "max_commission": 12,
                        "monthly_payment": "1118.76",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 20,
                        "wishlist_id": null
                    },
                    {
                        "id": 40021,
                        "name": "Samsung Galaxy D21 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40021/image-21-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40021/image-21-2.png"
                        ],
                        "slug": "samsung-galaxy-d21-5g-12-512-gb-seryy",
                        "min_price": "5821.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5821.00",
                        "max_commission": 12,
                        "monthly_payment": "1120.44",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 21,
                        "wishlist_id": null
                    },
                    {
                        "id": 40022,
                        "name": "Samsung Galaxy D22 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40022/image-22-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40022/image-22-2.png"
                        ],
                        "slug": "samsung-galaxy-d22-5g-12-512-gb-seryy",
                        "min_price": "5822.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5822.00",
                        "max_commission": 12,
                        "monthly_payment": "1122.12",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 22,
                        "wishlist_id": null
                    },
                    {
                        "id": 40023,
                        "name": "Samsung Galaxy D23 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40023/image-23-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40023/image-23-2.png"
                        ],
                        "slug": "samsung-galaxy-d23-5g-12-512-gb-seryy",
                        "min_price": "5823.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5823.00",
                        "max_commission": 12,
                        "monthly_payment": "1123.80",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 23,
                        "wishlist_id": null
                    },
                    {
                        "id": 40024,
                        "name": "Samsung Galaxy D24 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40024/image-24-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40024/image-24-2.png"
                        ],
                        "slug": "samsung-galaxy-d24-5g-12-512-gb-seryy",
                        "min_price": "5824.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5824.00",
                        "max_commission": 12,
                        "monthly_payment": "1125.48",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 24,
                        "wishlist_id": null
                    },
                    {
                        "id": 40025,
                        "name": "Samsung Galaxy D25 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40025/image-25-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40025/image-25-2.png"
                        ],
                        "slug": "samsung-galaxy-d25-5g-12-512-gb-seryy",
                        "min_price": "5825.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5825.00",
                        "max_commission": 12,
                        "monthly_payment": "1127.16",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 25,
                        "wishlist_id": null
                    },
                    {
                        "id": 40026,
                        "name": "Samsung Galaxy D26 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40026/image-26-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40026/image-26-2.png"
                        ],
                        "slug": "samsung-galaxy-d26-5g-12-512-gb-seryy",
                        "min_price": "5826.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5826.00",
                        "max_commission": 12,
                        "monthly_payment": "1128.84",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 26,
                        "wishlist_id": null
                    },
                    {
                        "id": 40027,
                        "name": "Samsung Galaxy D27 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40027/image-27-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40027/image-27-2.png"
                        ],
                        "slug": "samsung-galaxy-d27-5g-12-512-gb-seryy",
                        "min_price": "5827.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5827.00",
                        "max_commission": 12,
                        "monthly_payment": "1130.52",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 27,
                        "wishlist_id": null
                    },
                    {
                        "id": 40028,
                        "name": "Samsung Galaxy D28 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40028/image-28-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40028/image-28-2.png"
                        ],
                        "slug": "samsung-galaxy-d28-5g-12-512-gb-seryy",
                        "min_price": "5828.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5828.00",
                        "max_commission": 12,
                        "monthly_payment": "1132.20",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 28,
                        "wishlist_id": null
                    },
                    {
                        "id": 40029,
                        "name": "Samsung Galaxy D29 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40029/image-29-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40029/image-29-2.png"
                        ],
                        "slug": "samsung-galaxy-d29-5g-12-512-gb-seryy",
                        "min_price": "5829.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5829.00",
                        "max_commission": 12,
                        "monthly_payment": "1133.88",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 29,
                        "wishlist_id": null
                    },
                    {
                        "id": 40030,
                        "name": "Samsung Galaxy D30 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40030/image-30-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40030/image-30-2.png"
                        ],
                        "slug": "samsung-galaxy-d30-5g-12-512-gb-seryy",
                        "min_price": "5830.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5830.00",
                        "max_commission": 12,
                        "monthly_payment": "1135.56",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 30,
                        "wishlist_id": null
                    },
                    {
                        "id": 40031,
                        "name": "Samsung Galaxy D31 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40031/image-31-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40031/image-31-2.png"
                        ],
                        "slug": "samsung-galaxy-d31-5g-12-512-gb-seryy",
                        "min_price": "5831.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5831.00",
                        "max_commission": 12,
                        "monthly_payment": "1137.24",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 31,
                        "wishlist_id": null
                    }
                ],
                "view_all_url": "/popular/12",
                "position": 0.036,
                "properties": {
                    "rec_type": "popular"
                }
            },
            {
                "id": 9037,
                "city_id": 1,
                "name": "Popular_15",
                "title": "Popular_15",
                "is_hidden_title": 0,
                "widget_type": "product_rec",
                "view_mode": "web",
                "source": "main",
                "widget_items": [
                    {
                        "id": 40032,
                        "name": "Samsung Galaxy D32 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40032/image-32-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40032/image-32-2.png"
                        ],
                        "slug": "samsung-galaxy-d32-5g-12-512-gb-seryy",
                        "min_price": "5832.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5832.00",
                        "max_commission": 12,
                        "monthly_payment": "1138.92",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 32,
                        "wishlist_id": null
                    },
                    {
                        "id": 40033,
                        "name": "Samsung Galaxy D33 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40033/image-33-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40033/image-33-2.png"
                        ],
                        "slug": "samsung-galaxy-d33-5g-12-512-gb-seryy",
                        "min_price": "5833.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5833.00",
                        "max_commission": 12,
                        "monthly_payment": "1140.60",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 33,
                        "wishlist_id": null
                    },
                    {
                        "id": 40034,
                        "name": "Samsung Galaxy D34 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40034/image-34-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40034/image-34-2.png"
                        ],
                        "slug": "samsung-galaxy-d34-5g-12-512-gb-seryy",
                        "min_price": "5834.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5834.00",
                        "max_commission": 12,
                        "monthly_payment": "1142.28",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 34,
                        "wishlist_id": null
                    },
                    {
                        "id": 40035,
                        "name": "Samsung Galaxy D35 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40035/image-35-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40035/image-35-2.png"
                        ],
                        "slug": "samsung-galaxy-d35-5g-12-512-gb-seryy",
                        "min_price": "5835.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5835.00",
                        "max_commission": 12,
                        "monthly_payment": "1143.96",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 35,
                        "wishlist_id": null
                    },
                    {
                        "id": 40036,
                        "name": "Samsung Galaxy D36 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40036/image-36-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40036/image-36-2.png"
                        ],
                        "slug": "samsung-galaxy-d36-5g-12-512-gb-seryy",
                        "min_price": "5836.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5836.00",
                        "max_commission": 12,
                        "monthly_payment": "1145.64",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 36,
                        "wishlist_id": null
                    },
                    {
                        "id": 40037,
                        "name": "Samsung Galaxy D37 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40037/image-37-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40037/image-37-2.png"
                        ],
                        "slug": "samsung-galaxy-d37-5g-12-512-gb-seryy",
                        "min_price": "5837.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5837.00",
                        "max_commission": 12,
                        "monthly_payment": "1147.32",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 37,
                        "wishlist_id": null
                    },
                    {
                        "id": 40038,
                        "name": "Samsung Galaxy D38 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40038/image-38-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40038/image-38-2.png"
                        ],
                        "slug": "samsung-galaxy-d38-5g-12-512-gb-seryy",
                        "min_price": "5838.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5838.00",
                        "max_commission": 12,
                        "monthly_payment": "1149.00",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 38,
                        "wishlist_id": null
                    },
                    {
                        "id": 40039,
                        "name": "Samsung Galaxy D39 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40039/image-39-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40039/image-39-2.png"
                        ],
                        "slug": "samsung-galaxy-d39-5g-12-512-gb-seryy",
                        "min_price": "5839.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5839.00",
                        "max_commission": 12,
                        "monthly_payment": "1150.68",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 39,
                        "wishlist_id": null
                    },
                    {
                        "id": 40040,
                        "name": "Samsung Galaxy D40 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40040/image-40-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40040/image-40-2.png"
                        ],
                        "slug": "samsung-galaxy-d40-5g-12-512-gb-seryy",
                        "min_price": "5840.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5840.00",
                        "max_commission": 12,
                        "monthly_payment": "1152.36",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 40,
                        "wishlist_id": null
                    },
                    {
                        "id": 40041,
                        "name": "Samsung Galaxy D41 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40041/image-41-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40041/image-41-2.png"
                        ],
                        "slug": "samsung-galaxy-d41-5g-12-512-gb-seryy",
                        "min_price": "5841.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5841.00",
                        "max_commission": 12,
                        "monthly_payment": "1154.04",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 41,
                        "wishlist_id": null
                    },
                    {
                        "id": 40042,
                        "name": "Samsung Galaxy D42 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40042/image-42-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40042/image-42-2.png"
                        ],
                        "slug": "samsung-galaxy-d42-5g-12-512-gb-seryy",
                        "min_price": "5842.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5842.00",
                        "max_commission": 12,
                        "monthly_payment": "1155.72",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 42,
                        "wishlist_id": null
                    },
                    {
                        "id": 40043,
                        "name": "Samsung Galaxy D43 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40043/image-43-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40043/image-43-2.png"
                        ],
                        "slug": "samsung-galaxy-d43-5g-12-512-gb-seryy",
                        "min_price": "5843.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5843.00",
                        "max_commission": 12,
                        "monthly_payment": "1157.40",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 43,
                        "wishlist_id": null
                    },
                    {
                        "id": 40044,
                        "name": "Samsung Galaxy D44 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40044/image-44-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40044/image-44-2.png"
                        ],
                        "slug": "samsung-galaxy-d44-5g-12-512-gb-seryy",
                        "min_price": "5844.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5844.00",
                        "max_commission": 12,
                        "monthly_payment": "1159.08",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 44,
                        "wishlist_id": null
                    },
                    {
                        "id": 40045,
                        "name": "Samsung Galaxy D45 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40045/image-45-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40045/image-45-2.png"
                        ],
                        "slug": "samsung-galaxy-d45-5g-12-512-gb-seryy",
                        "min_price": "5845.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5845.00",
                        "max_commission": 12,
                        "monthly_payment": "1160.76",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 45,
                        "wishlist_id": null
                    },
                    {
                        "id": 40046,
                        "name": "Samsung Galaxy D46 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40046/image-46-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40046/image-46-2.png"
                        ],
                        "slug": "samsung-galaxy-d46-5g-12-512-gb-seryy",
                        "min_price": "5846.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5846.00",
                        "max_commission": 12,
                        "monthly_payment": "1162.44",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 46,
                        "wishlist_id": null
                    }
                ],
                "view_all_url": "/popular/15",
                "position": 0.037,
                "properties": {
                    "rec_type": "popular"
                }
            },
            {
                "id": 9038,
                "city_id": 1,
                "name": "Popular_18",
                "title": "Popular_18",
                "is_hidden_title": 0,
                "widget_type": "product_rec",
                "view_mode": "web",
                "source": "main",
                "widget_items": [
                    {
                        "id": 40047,
                        "name": "Samsung Galaxy D47 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40047/image-47-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40047/image-47-2.png"
                        ],
                        "slug": "samsung-galaxy-d47-5g-12-512-gb-seryy",
                        "min_price": "5847.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5847.00",
                        "max_commission": 12,
                        "monthly_payment": "1164.12",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 47,
                        "wishlist_id": null
                    },
                    {
                        "id": 40048,
                        "name": "Samsung Galaxy D48 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40048/image-48-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40048/image-48-2.png"
                        ],
                        "slug": "samsung-galaxy-d48-5g-12-512-gb-seryy",
                        "min_price": "5848.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5848.00",
                        "max_commission": 12,
                        "monthly_payment": "1165.80",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 48,
                        "wishlist_id": null
                    },
                    {
                        "id": 40049,
                        "name": "Samsung Galaxy D49 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40049/image-49-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40049/image-49-2.png"
                        ],
                        "slug": "samsung-galaxy-d49-5g-12-512-gb-seryy",
                        "min_price": "5849.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5849.00",
                        "max_commission": 12,
                        "monthly_payment": "1167.48",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 49,
                        "wishlist_id": null
                    },
                    {
                        "id": 40050,
                        "name": "Samsung Galaxy D50 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40050/image-50-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40050/image-50-2.png"
                        ],
                        "slug": "samsung-galaxy-d50-5g-12-512-gb-seryy",
                        "min_price": "5850.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5850.00",
                        "max_commission": 12,
                        "monthly_payment": "1169.16",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 50,
                        "wishlist_id": null
                    },
                    {
                        "id": 40051,
                        "name": "Samsung Galaxy D51 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40051/image-51-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40051/image-51-2.png"
                        ],
                        "slug": "samsung-galaxy-d51-5g-12-512-gb-seryy",
                        "min_price": "5851.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5851.00",
                        "max_commission": 12,
                        "monthly_payment": "1170.84",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 51,
                        "wishlist_id": null
                    },
                    {
                        "id": 40052,
                        "name": "Samsung Galaxy D52 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40052/image-52-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40052/image-52-2.png"
                        ],
                        "slug": "samsung-galaxy-d52-5g-12-512-gb-seryy",
                        "min_price": "5852.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5852.00",
                        "max_commission": 12,
                        "monthly_payment": "1172.52",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 52,
                        "wishlist_id": null
                    },
                    {
                        "id": 40053,
                        "name": "Samsung Galaxy D53 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40053/image-53-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40053/image-53-2.png"
                        ],
                        "slug": "samsung-galaxy-d53-5g-12-512-gb-seryy",
                        "min_price": "5853.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5853.00",
                        "max_commission": 12,
                        "monthly_payment": "1174.20",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 53,
                        "wishlist_id": null
                    },
                    {
                        "id": 40054,
                        "name": "Samsung Galaxy D54 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40054/image-54-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40054/image-54-2.png"
                        ],
                        "slug": "samsung-galaxy-d54-5g-12-512-gb-seryy",
                        "min_price": "5854.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5854.00",
                        "max_commission": 12,
                        "monthly_payment": "1175.88",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 54,
                        "wishlist_id": null
                    },
                    {
                        "id": 40055,
                        "name": "Samsung Galaxy D55 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40055/image-55-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40055/image-55-2.png"
                        ],
                        "slug": "samsung-galaxy-d55-5g-12-512-gb-seryy",
                        "min_price": "5855.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5855.00",
                        "max_commission": 12,
                        "monthly_payment": "1177.56",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 55,
                        "wishlist_id": null
                    },
                    {
                        "id": 40056,
                        "name": "Samsung Galaxy D56 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40056/image-56-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40056/image-56-2.png"
                        ],
                        "slug": "samsung-galaxy-d56-5g-12-512-gb-seryy",
                        "min_price": "5856.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5856.00",
                        "max_commission": 12,
                        "monthly_payment": "1179.24",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 56,
                        "wishlist_id": null
                    },
                    {
                        "id": 40057,
                        "name": "Samsung Galaxy D57 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40057/image-57-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40057/image-57-2.png"
                        ],
                        "slug": "samsung-galaxy-d57-5g-12-512-gb-seryy",
                        "min_price": "5857.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5857.00",
                        "max_commission": 12,
                        "monthly_payment": "1180.92",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 57,
                        "wishlist_id": null
                    },
                    {
                        "id": 40058,
                        "name": "Samsung Galaxy D58 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40058/image-58-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40058/image-58-2.png"
                        ],
                        "slug": "samsung-galaxy-d58-5g-12-512-gb-seryy",
                        "min_price": "5858.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5858.00",
                        "max_commission": 12,
                        "monthly_payment": "1182.60",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 58,
                        "wishlist_id": null
                    },
                    {
                        "id": 40059,
                        "name": "Samsung Galaxy D59 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40059/image-59-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40059/image-59-2.png"
                        ],
                        "slug": "samsung-galaxy-d59-5g-12-512-gb-seryy",
                        "min_price": "5859.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5859.00",
                        "max_commission": 12,
                        "monthly_payment": "1184.28",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 59,
                        "wishlist_id": null
                    },
                    {
                        "id": 40060,
                        "name": "Samsung Galaxy D60 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40060/image-60-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40060/image-60-2.png"
                        ],
                        "slug": "samsung-galaxy-d60-5g-12-512-gb-seryy",
                        "min_price": "5860.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5860.00",
                        "max_commission": 12,
                        "monthly_payment": "1185.96",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 60,
                        "wishlist_id": null
                    },
                    {
                        "id": 40061,
                        "name": "Samsung Galaxy D61 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40061/image-61-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40061/image-61-2.png"
                        ],
                        "slug": "samsung-galaxy-d61-5g-12-512-gb-seryy",
                        "min_price": "5861.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5861.00",
                        "max_commission": 12,
                        "monthly_payment": "1187.64",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 61,
                        "wishlist_id": null
                    },
                    {
                        "id": 40062,
                        "name": "Samsung Galaxy D62 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40062/image-62-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40062/image-62-2.png"
                        ],
                        "slug": "samsung-galaxy-d62-5g-12-512-gb-seryy",
                        "min_price": "5862.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5862.00",
                        "max_commission": 12,
                        "monthly_payment": "1189.32",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 62,
                        "wishlist_id": null
                    },
                    {
                        "id": 40063,
                        "name": "Samsung Galaxy D63 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40063/image-63-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40063/image-63-2.png"
                        ],
                        "slug": "samsung-galaxy-d63-5g-12-512-gb-seryy",
                        "min_price": "5863.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5863.00",
                        "max_commission": 12,
                        "monthly_payment": "1191.00",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 63,
                        "wishlist_id": null
                    },
                    {
                        "id": 40064,
                        "name": "Samsung Galaxy D64 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40064/image-64-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40064/image-64-2.png"
                        ],
                        "slug": "samsung-galaxy-d64-5g-12-512-gb-seryy",
                        "min_price": "5864.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5864.00",
                        "max_commission": 12,
                        "monthly_payment": "1192.68",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 64,
                        "wishlist_id": null
                    }
                ],
                "view_all_url": "/popular/18",
                "position": 0.038,
                "properties": {
                    "rec_type": "popular"
                }
            },
            {
                "id": 9039,
                "city_id": 1,
                "name": "Popular_20",
                "title": "Popular_20",
                "is_hidden_title": 0,
                "widget_type": "product_rec",
                "view_mode": "web",
                "source": "main",
                "widget_items": [
                    {
                        "id": 40065,
                        "name": "Samsung Galaxy D65 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40065/image-65-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40065/image-65-2.png"
                        ],
                        "slug": "samsung-galaxy-d65-5g-12-512-gb-seryy",
                        "min_price": "5865.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5865.00",
                        "max_commission": 12,
                        "monthly_payment": "1194.36",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 65,
                        "wishlist_id": null
                    },
                    {
                        "id": 40066,
                        "name": "Samsung Galaxy D66 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40066/image-66-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40066/image-66-2.png"
                        ],
                        "slug": "samsung-galaxy-d66-5g-12-512-gb-seryy",
                        "min_price": "5866.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5866.00",
                        "max_commission": 12,
                        "monthly_payment": "1196.04",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 66,
                        "wishlist_id": null
                    },
                    {
                        "id": 40067,
                        "name": "Samsung Galaxy D67 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40067/image-67-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40067/image-67-2.png"
                        ],
                        "slug": "samsung-galaxy-d67-5g-12-512-gb-seryy",
                        "min_price": "5867.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5867.00",
                        "max_commission": 12,
                        "monthly_payment": "1197.72",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 67,
                        "wishlist_id": null
                    },
                    {
                        "id": 40068,
                        "name": "Samsung Galaxy D68 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40068/image-68-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40068/image-68-2.png"
                        ],
                        "slug": "samsung-galaxy-d68-5g-12-512-gb-seryy",
                        "min_price": "5868.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5868.00",
                        "max_commission": 12,
                        "monthly_payment": "1199.40",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 68,
                        "wishlist_id": null
                    },
                    {
                        "id": 40069,
                        "name": "Samsung Galaxy D69 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40069/image-69-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40069/image-69-2.png"
                        ],
                        "slug": "samsung-galaxy-d69-5g-12-512-gb-seryy",
                        "min_price": "5869.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5869.00",
                        "max_commission": 12,
                        "monthly_payment": "1201.08",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 69,
                        "wishlist_id": null
                    },
                    {
                        "id": 40070,
                        "name": "Samsung Galaxy D70 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40070/image-70-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40070/image-70-2.png"
                        ],
                        "slug": "samsung-galaxy-d70-5g-12-512-gb-seryy",
                        "min_price": "5870.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5870.00",
                        "max_commission": 12,
                        "monthly_payment": "1202.76",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 70,
                        "wishlist_id": null
                    },
                    {
                        "id": 40071,
                        "name": "Samsung Galaxy D71 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40071/image-71-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40071/image-71-2.png"
                        ],
                        "slug": "samsung-galaxy-d71-5g-12-512-gb-seryy",
                        "min_price": "5871.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5871.00",
                        "max_commission": 12,
                        "monthly_payment": "1204.44",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 71,
                        "wishlist_id": null
                    },
                    {
                        "id": 40072,
                        "name": "Samsung Galaxy D72 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40072/image-72-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40072/image-72-2.png"
                        ],
                        "slug": "samsung-galaxy-d72-5g-12-512-gb-seryy",
                        "min_price": "5872.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5872.00",
                        "max_commission": 12,
                        "monthly_payment": "1206.12",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 72,
                        "wishlist_id": null
                    },
                    {
                        "id": 40073,
                        "name": "Samsung Galaxy D73 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40073/image-73-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40073/image-73-2.png"
                        ],
                        "slug": "samsung-galaxy-d73-5g-12-512-gb-seryy",
                        "min_price": "5873.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5873.00",
                        "max_commission": 12,
                        "monthly_payment": "1207.80",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 73,
                        "wishlist_id": null
                    },
                    {
                        "id": 40074,
                        "name": "Samsung Galaxy D74 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40074/image-74-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40074/image-74-2.png"
                        ],
                        "slug": "samsung-galaxy-d74-5g-12-512-gb-seryy",
                        "min_price": "5874.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5874.00",
                        "max_commission": 12,
                        "monthly_payment": "1209.48",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 74,
                        "wishlist_id": null
                    },
                    {
                        "id": 40075,
                        "name": "Samsung Galaxy D75 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40075/image-75-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40075/image-75-2.png"
                        ],
                        "slug": "samsung-galaxy-d75-5g-12-512-gb-seryy",
                        "min_price": "5875.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5875.00",
                        "max_commission": 12,
                        "monthly_payment": "1211.16",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 75,
                        "wishlist_id": null
                    },
                    {
                        "id": 40076,
                        "name": "Samsung Galaxy D76 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40076/image-76-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40076/image-76-2.png"
                        ],
                        "slug": "samsung-galaxy-d76-5g-12-512-gb-seryy",
                        "min_price": "5876.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5876.00",
                        "max_commission": 12,
                        "monthly_payment": "1212.84",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 76,
                        "wishlist_id": null
                    },
                    {
                        "id": 40077,
                        "name": "Samsung Galaxy D77 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40077/image-77-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40077/image-77-2.png"
                        ],
                        "slug": "samsung-galaxy-d77-5g-12-512-gb-seryy",
                        "min_price": "5877.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5877.00",
                        "max_commission": 12,
                        "monthly_payment": "1214.52",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 77,
                        "wishlist_id": null
                    },
                    {
                        "id": 40078,
                        "name": "Samsung Galaxy D78 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40078/image-78-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40078/image-78-2.png"
                        ],
                        "slug": "samsung-galaxy-d78-5g-12-512-gb-seryy",
                        "min_price": "5878.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5878.00",
                        "max_commission": 12,
                        "monthly_payment": "1216.20",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 78,
                        "wishlist_id": null
                    },
                    {
                        "id": 40079,
                        "name": "Samsung Galaxy D79 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40079/image-79-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40079/image-79-2.png"
                        ],
                        "slug": "samsung-galaxy-d79-5g-12-512-gb-seryy",
                        "min_price": "5879.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5879.00",
                        "max_commission": 12,
                        "monthly_payment": "1217.88",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 79,
                        "wishlist_id": null
                    },
                    {
                        "id": 40080,
                        "name": "Samsung Galaxy D80 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40080/image-80-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40080/image-80-2.png"
                        ],
                        "slug": "samsung-galaxy-d80-5g-12-512-gb-seryy",
                        "min_price": "5880.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5880.00",
                        "max_commission": 12,
                        "monthly_payment": "1219.56",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 80,
                        "wishlist_id": null
                    },
                    {
                        "id": 40081,
                        "name": "Samsung Galaxy D81 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40081/image-81-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40081/image-81-2.png"
                        ],
                        "slug": "samsung-galaxy-d81-5g-12-512-gb-seryy",
                        "min_price": "5881.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5881.00",
                        "max_commission": 12,
                        "monthly_payment": "1221.24",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 81,
                        "wishlist_id": null
                    },
                    {
                        "id": 40082,
                        "name": "Samsung Galaxy D82 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40082/image-82-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40082/image-82-2.png"
                        ],
                        "slug": "samsung-galaxy-d82-5g-12-512-gb-seryy",
                        "min_price": "5882.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5882.00",
                        "max_commission": 12,
                        "monthly_payment": "1222.92",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 82,
                        "wishlist_id": null
                    },
                    {
                        "id": 40083,
                        "name": "Samsung Galaxy D83 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40083/image-83-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40083/image-83-2.png"
                        ],
                        "slug": "samsung-galaxy-d83-5g-12-512-gb-seryy",
                        "min_price": "5883.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5883.00",
                        "max_commission": 12,
                        "monthly_payment": "1224.60",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 83,
                        "wishlist_id": null
                    },
                    {
                        "id": 40084,
                        "name": "Samsung Galaxy D84 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40084/image-84-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40084/image-84-2.png"
                        ],
                        "slug": "samsung-galaxy-d84-5g-12-512-gb-seryy",
                        "min_price": "5884.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5884.00",
                        "max_commission": 12,
                        "monthly_payment": "1226.28",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 84,
                        "wishlist_id": null
                    }
                ],
                "view_all_url": "/popular/20",
                "position": 0.039,
                "properties": {
                    "rec_type": "popular"
                }
            },
            {
                "id": 9040,
                "city_id": 1,
                "name": "Popular_30",
                "title": "Popular_30",
                "is_hidden_title": 0,
                "widget_type": "product_rec",
                "view_mode": "web",
                "source": "main",
                "widget_items": [
                    {
                        "id": 40085,
                        "name": "Samsung Galaxy D85 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40085/image-85-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40085/image-85-2.png"
                        ],
                        "slug": "samsung-galaxy-d85-5g-12-512-gb-seryy",
                        "min_price": "5885.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5885.00",
                        "max_commission": 12,
                        "monthly_payment": "1227.96",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 85,
                        "wishlist_id": null
                    },
                    {
                        "id": 40086,
                        "name": "Samsung Galaxy D86 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40086/image-86-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40086/image-86-2.png"
                        ],
                        "slug": "samsung-galaxy-d86-5g-12-512-gb-seryy",
                        "min_price": "5886.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5886.00",
                        "max_commission": 12,
                        "monthly_payment": "1229.64",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 86,
                        "wishlist_id": null
                    },
                    {
                        "id": 40087,
                        "name": "Samsung Galaxy D87 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40087/image-87-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40087/image-87-2.png"
                        ],
                        "slug": "samsung-galaxy-d87-5g-12-512-gb-seryy",
                        "min_price": "5887.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5887.00",
                        "max_commission": 12,
                        "monthly_payment": "1231.32",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 87,
                        "wishlist_id": null
                    },
                    {
                        "id": 40088,
                        "name": "Samsung Galaxy D88 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40088/image-88-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40088/image-88-2.png"
                        ],
                        "slug": "samsung-galaxy-d88-5g-12-512-gb-seryy",
                        "min_price": "5888.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5888.00",
                        "max_commission": 12,
                        "monthly_payment": "1233.00",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 88,
                        "wishlist_id": null
                    },
                    {
                        "id": 40089,
                        "name": "Samsung Galaxy D89 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40089/image-89-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40089/image-89-2.png"
                        ],
                        "slug": "samsung-galaxy-d89-5g-12-512-gb-seryy",
                        "min_price": "5889.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5889.00",
                        "max_commission": 12,
                        "monthly_payment": "1234.68",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 89,
                        "wishlist_id": null
                    },
                    {
                        "id": 40090,
                        "name": "Samsung Galaxy D90 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40090/image-90-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40090/image-90-2.png"
                        ],
                        "slug": "samsung-galaxy-d90-5g-12-512-gb-seryy",
                        "min_price": "5890.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5890.00",
                        "max_commission": 12,
                        "monthly_payment": "1236.36",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 90,
                        "wishlist_id": null
                    },
                    {
                        "id": 40091,
                        "name": "Samsung Galaxy D91 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40091/image-91-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40091/image-91-2.png"
                        ],
                        "slug": "samsung-galaxy-d91-5g-12-512-gb-seryy",
                        "min_price": "5891.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5891.00",
                        "max_commission": 12,
                        "monthly_payment": "1238.04",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 91,
                        "wishlist_id": null
                    },
                    {
                        "id": 40092,
                        "name": "Samsung Galaxy D92 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40092/image-92-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40092/image-92-2.png"
                        ],
                        "slug": "samsung-galaxy-d92-5g-12-512-gb-seryy",
                        "min_price": "5892.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5892.00",
                        "max_commission": 12,
                        "monthly_payment": "1239.72",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 92,
                        "wishlist_id": null
                    },
                    {
                        "id": 40093,
                        "name": "Samsung Galaxy D93 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40093/image-93-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40093/image-93-2.png"
                        ],
                        "slug": "samsung-galaxy-d93-5g-12-512-gb-seryy",
                        "min_price": "5893.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5893.00",
                        "max_commission": 12,
                        "monthly_payment": "1241.40",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 93,
                        "wishlist_id": null
                    },
                    {
                        "id": 40094,
                        "name": "Samsung Galaxy D94 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40094/image-94-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40094/image-94-2.png"
                        ],
                        "slug": "samsung-galaxy-d94-5g-12-512-gb-seryy",
                        "min_price": "5894.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5894.00",
                        "max_commission": 12,
                        "monthly_payment": "1243.08",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 94,
                        "wishlist_id": null
                    },
                    {
                        "id": 40095,
                        "name": "Samsung Galaxy D95 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40095/image-95-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40095/image-95-2.png"
                        ],
                        "slug": "samsung-galaxy-d95-5g-12-512-gb-seryy",
                        "min_price": "5895.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5895.00",
                        "max_commission": 12,
                        "monthly_payment": "1244.76",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 95,
                        "wishlist_id": null
                    },
                    {
                        "id": 40096,
                        "name": "Samsung Galaxy D96 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40096/image-96-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40096/image-96-2.png"
                        ],
                        "slug": "samsung-galaxy-d96-5g-12-512-gb-seryy",
                        "min_price": "5896.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5896.00",
                        "max_commission": 12,
                        "monthly_payment": "1246.44",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 96,
                        "wishlist_id": null
                    },
                    {
                        "id": 40097,
                        "name": "Samsung Galaxy D97 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40097/image-97-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40097/image-97-2.png"
                        ],
                        "slug": "samsung-galaxy-d97-5g-12-512-gb-seryy",
                        "min_price": "5897.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5897.00",
                        "max_commission": 12,
                        "monthly_payment": "1248.12",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 97,
                        "wishlist_id": null
                    },
                    {
                        "id": 40098,
                        "name": "Samsung Galaxy D98 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40098/image-98-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40098/image-98-2.png"
                        ],
                        "slug": "samsung-galaxy-d98-5g-12-512-gb-seryy",
                        "min_price": "5898.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5898.00",
                        "max_commission": 12,
                        "monthly_payment": "1249.80",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 98,
                        "wishlist_id": null
                    },
                    {
                        "id": 40099,
                        "name": "Samsung Galaxy D99 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40099/image-99-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40099/image-99-2.png"
                        ],
                        "slug": "samsung-galaxy-d99-5g-12-512-gb-seryy",
                        "min_price": "5899.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5899.00",
                        "max_commission": 12,
                        "monthly_payment": "1251.48",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 99,
                        "wishlist_id": null
                    },
                    {
                        "id": 40100,
                        "name": "Samsung Galaxy D100 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40100/image-100-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40100/image-100-2.png"
                        ],
                        "slug": "samsung-galaxy-d100-5g-12-512-gb-seryy",
                        "min_price": "5900.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5900.00",
                        "max_commission": 12,
                        "monthly_payment": "1253.16",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 100,
                        "wishlist_id": null
                    },
                    {
                        "id": 40101,
                        "name": "Samsung Galaxy D101 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40101/image-101-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40101/image-101-2.png"
                        ],
                        "slug": "samsung-galaxy-d101-5g-12-512-gb-seryy",
                        "min_price": "5901.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5901.00",
                        "max_commission": 12,
                        "monthly_payment": "1254.84",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 101,
                        "wishlist_id": null
                    },
                    {
                        "id": 40102,
                        "name": "Samsung Galaxy D102 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40102/image-102-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40102/image-102-2.png"
                        ],
                        "slug": "samsung-galaxy-d102-5g-12-512-gb-seryy",
                        "min_price": "5902.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5902.00",
                        "max_commission": 12,
                        "monthly_payment": "1256.52",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 102,
                        "wishlist_id": null
                    },
                    {
                        "id": 40103,
                        "name": "Samsung Galaxy D103 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40103/image-103-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40103/image-103-2.png"
                        ],
                        "slug": "samsung-galaxy-d103-5g-12-512-gb-seryy",
                        "min_price": "5903.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5903.00",
                        "max_commission": 12,
                        "monthly_payment": "1258.20",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 103,
                        "wishlist_id": null
                    },
                    {
                        "id": 40104,
                        "name": "Samsung Galaxy D104 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40104/image-104-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40104/image-104-2.png"
                        ],
                        "slug": "samsung-galaxy-d104-5g-12-512-gb-seryy",
                        "min_price": "5904.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5904.00",
                        "max_commission": 12,
                        "monthly_payment": "1259.88",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 104,
                        "wishlist_id": null
                    },
                    {
                        "id": 40105,
                        "name": "Samsung Galaxy D105 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40105/image-105-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40105/image-105-2.png"
                        ],
                        "slug": "samsung-galaxy-d105-5g-12-512-gb-seryy",
                        "min_price": "5905.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5905.00",
                        "max_commission": 12,
                        "monthly_payment": "1261.56",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 105,
                        "wishlist_id": null
                    },
                    {
                        "id": 40106,
                        "name": "Samsung Galaxy D106 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40106/image-106-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40106/image-106-2.png"
                        ],
                        "slug": "samsung-galaxy-d106-5g-12-512-gb-seryy",
                        "min_price": "5906.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5906.00",
                        "max_commission": 12,
                        "monthly_payment": "1263.24",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 106,
                        "wishlist_id": null
                    },
                    {
                        "id": 40107,
                        "name": "Samsung Galaxy D107 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40107/image-107-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40107/image-107-2.png"
                        ],
                        "slug": "samsung-galaxy-d107-5g-12-512-gb-seryy",
                        "min_price": "5907.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5907.00",
                        "max_commission": 12,
                        "monthly_payment": "1264.92",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 107,
                        "wishlist_id": null
                    },
                    {
                        "id": 40108,
                        "name": "Samsung Galaxy D108 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40108/image-108-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40108/image-108-2.png"
                        ],
                        "slug": "samsung-galaxy-d108-5g-12-512-gb-seryy",
                        "min_price": "5908.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5908.00",
                        "max_commission": 12,
                        "monthly_payment": "1266.60",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 108,
                        "wishlist_id": null
                    },
                    {
                        "id": 40109,
                        "name": "Samsung Galaxy D109 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40109/image-109-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40109/image-109-2.png"
                        ],
                        "slug": "samsung-galaxy-d109-5g-12-512-gb-seryy",
                        "min_price": "5909.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5909.00",
                        "max_commission": 12,
                        "monthly_payment": "1268.28",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 109,
                        "wishlist_id": null
                    },
                    {
                        "id": 40110,
                        "name": "Samsung Galaxy D110 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40110/image-110-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40110/image-110-2.png"
                        ],
                        "slug": "samsung-galaxy-d110-5g-12-512-gb-seryy",
                        "min_price": "5910.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5910.00",
                        "max_commission": 12,
                        "monthly_payment": "1269.96",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 110,
                        "wishlist_id": null
                    },
                    {
                        "id": 40111,
                        "name": "Samsung Galaxy D111 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40111/image-111-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40111/image-111-2.png"
                        ],
                        "slug": "samsung-galaxy-d111-5g-12-512-gb-seryy",
                        "min_price": "5911.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5911.00",
                        "max_commission": 12,
                        "monthly_payment": "1271.64",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 111,
                        "wishlist_id": null
                    },
                    {
                        "id": 40112,
                        "name": "Samsung Galaxy D112 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40112/image-112-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40112/image-112-2.png"
                        ],
                        "slug": "samsung-galaxy-d112-5g-12-512-gb-seryy",
                        "min_price": "5912.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5912.00",
                        "max_commission": 12,
                        "monthly_payment": "1273.32",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 112,
                        "wishlist_id": null
                    },
                    {
                        "id": 40113,
                        "name": "Samsung Galaxy D113 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40113/image-113-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40113/image-113-2.png"
                        ],
                        "slug": "samsung-galaxy-d113-5g-12-512-gb-seryy",
                        "min_price": "5913.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5913.00",
                        "max_commission": 12,
                        "monthly_payment": "1275.00",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 113,
                        "wishlist_id": null
                    },
                    {
                        "id": 40114,
                        "name": "Samsung Galaxy D114 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40114/image-114-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40114/image-114-2.png"
                        ],
                        "slug": "samsung-galaxy-d114-5g-12-512-gb-seryy",
                        "min_price": "5914.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5914.00",
                        "max_commission": 12,
                        "monthly_payment": "1276.68",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 114,
                        "wishlist_id": null
                    }
                ],
                "view_all_url": "/popular/30",
                "position": 0.04,
                "properties": {
                    "rec_type": "popular"
                }
            },
            {
                "id": 9041,
                "city_id": 1,
                "name": "Popular_35",
                "title": "Popular_35",
                "is_hidden_title": 0,
                "widget_type": "product_rec",
                "view_mode": "web",
                "source": "main",
                "widget_items": [
                    {
                        "id": 40115,
                        "name": "Samsung Galaxy D115 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40115/image-115-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40115/image-115-2.png"
                        ],
                        "slug": "samsung-galaxy-d115-5g-12-512-gb-seryy",
                        "min_price": "5915.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5915.00",
                        "max_commission": 12,
                        "monthly_payment": "1278.36",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 115,
                        "wishlist_id": null
                    },
                    {
                        "id": 40116,
                        "name": "Samsung Galaxy D116 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40116/image-116-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40116/image-116-2.png"
                        ],
                        "slug": "samsung-galaxy-d116-5g-12-512-gb-seryy",
                        "min_price": "5916.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5916.00",
                        "max_commission": 12,
                        "monthly_payment": "1280.04",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 116,
                        "wishlist_id": null
                    },
                    {
                        "id": 40117,
                        "name": "Samsung Galaxy D117 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40117/image-117-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40117/image-117-2.png"
                        ],
                        "slug": "samsung-galaxy-d117-5g-12-512-gb-seryy",
                        "min_price": "5917.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5917.00",
                        "max_commission": 12,
                        "monthly_payment": "1281.72",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 117,
                        "wishlist_id": null
                    },
                    {
                        "id": 40118,
                        "name": "Samsung Galaxy D118 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40118/image-118-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40118/image-118-2.png"
                        ],
                        "slug": "samsung-galaxy-d118-5g-12-512-gb-seryy",
                        "min_price": "5918.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5918.00",
                        "max_commission": 12,
                        "monthly_payment": "1283.40",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 118,
                        "wishlist_id": null
                    },
                    {
                        "id": 40119,
                        "name": "Samsung Galaxy D119 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40119/image-119-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40119/image-119-2.png"
                        ],
                        "slug": "samsung-galaxy-d119-5g-12-512-gb-seryy",
                        "min_price": "5919.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5919.00",
                        "max_commission": 12,
                        "monthly_payment": "1285.08",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 119,
                        "wishlist_id": null
                    },
                    {
                        "id": 40120,
                        "name": "Samsung Galaxy D120 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40120/image-120-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40120/image-120-2.png"
                        ],
                        "slug": "samsung-galaxy-d120-5g-12-512-gb-seryy",
                        "min_price": "5920.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5920.00",
                        "max_commission": 12,
                        "monthly_payment": "1286.76",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 120,
                        "wishlist_id": null
                    },
                    {
                        "id": 40121,
                        "name": "Samsung Galaxy D121 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40121/image-121-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40121/image-121-2.png"
                        ],
                        "slug": "samsung-galaxy-d121-5g-12-512-gb-seryy",
                        "min_price": "5921.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5921.00",
                        "max_commission": 12,
                        "monthly_payment": "1288.44",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 121,
                        "wishlist_id": null
                    },
                    {
                        "id": 40122,
                        "name": "Samsung Galaxy D122 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40122/image-122-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40122/image-122-2.png"
                        ],
                        "slug": "samsung-galaxy-d122-5g-12-512-gb-seryy",
                        "min_price": "5922.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5922.00",
                        "max_commission": 12,
                        "monthly_payment": "1290.12",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 122,
                        "wishlist_id": null
                    },
                    {
                        "id": 40123,
                        "name": "Samsung Galaxy D123 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40123/image-123-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40123/image-123-2.png"
                        ],
                        "slug": "samsung-galaxy-d123-5g-12-512-gb-seryy",
                        "min_price": "5923.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5923.00",
                        "max_commission": 12,
                        "monthly_payment": "1291.80",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 123,
                        "wishlist_id": null
                    },
                    {
                        "id": 40124,
                        "name": "Samsung Galaxy D124 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40124/image-124-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40124/image-124-2.png"
                        ],
                        "slug": "samsung-galaxy-d124-5g-12-512-gb-seryy",
                        "min_price": "5924.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5924.00",
                        "max_commission": 12,
                        "monthly_payment": "1293.48",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 124,
                        "wishlist_id": null
                    },
                    {
                        "id": 40125,
                        "name": "Samsung Galaxy D125 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40125/image-125-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40125/image-125-2.png"
                        ],
                        "slug": "samsung-galaxy-d125-5g-12-512-gb-seryy",
                        "min_price": "5925.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5925.00",
                        "max_commission": 12,
                        "monthly_payment": "1295.16",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 125,
                        "wishlist_id": null
                    },
                    {
                        "id": 40126,
                        "name": "Samsung Galaxy D126 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40126/image-126-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40126/image-126-2.png"
                        ],
                        "slug": "samsung-galaxy-d126-5g-12-512-gb-seryy",
                        "min_price": "5926.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5926.00",
                        "max_commission": 12,
                        "monthly_payment": "1296.84",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 126,
                        "wishlist_id": null
                    },
                    {
                        "id": 40127,
                        "name": "Samsung Galaxy D127 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40127/image-127-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40127/image-127-2.png"
                        ],
                        "slug": "samsung-galaxy-d127-5g-12-512-gb-seryy",
                        "min_price": "5927.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5927.00",
                        "max_commission": 12,
                        "monthly_payment": "1298.52",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 127,
                        "wishlist_id": null
                    },
                    {
                        "id": 40128,
                        "name": "Samsung Galaxy D128 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40128/image-128-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40128/image-128-2.png"
                        ],
                        "slug": "samsung-galaxy-d128-5g-12-512-gb-seryy",
                        "min_price": "5928.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5928.00",
                        "max_commission": 12,
                        "monthly_payment": "1300.20",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 128,
                        "wishlist_id": null
                    },
                    {
                        "id": 40129,
                        "name": "Samsung Galaxy D129 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40129/image-129-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40129/image-129-2.png"
                        ],
                        "slug": "samsung-galaxy-d129-5g-12-512-gb-seryy",
                        "min_price": "5929.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5929.00",
                        "max_commission": 12,
                        "monthly_payment": "1301.88",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 129,
                        "wishlist_id": null
                    },
                    {
                        "id": 40130,
                        "name": "Samsung Galaxy D130 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40130/image-130-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40130/image-130-2.png"
                        ],
                        "slug": "samsung-galaxy-d130-5g-12-512-gb-seryy",
                        "min_price": "5930.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5930.00",
                        "max_commission": 12,
                        "monthly_payment": "1303.56",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 130,
                        "wishlist_id": null
                    },
                    {
                        "id": 40131,
                        "name": "Samsung Galaxy D131 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40131/image-131-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40131/image-131-2.png"
                        ],
                        "slug": "samsung-galaxy-d131-5g-12-512-gb-seryy",
                        "min_price": "5931.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5931.00",
                        "max_commission": 12,
                        "monthly_payment": "1305.24",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 131,
                        "wishlist_id": null
                    },
                    {
                        "id": 40132,
                        "name": "Samsung Galaxy D132 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40132/image-132-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40132/image-132-2.png"
                        ],
                        "slug": "samsung-galaxy-d132-5g-12-512-gb-seryy",
                        "min_price": "5932.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5932.00",
                        "max_commission": 12,
                        "monthly_payment": "1306.92",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 132,
                        "wishlist_id": null
                    },
                    {
                        "id": 40133,
                        "name": "Samsung Galaxy D133 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40133/image-133-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40133/image-133-2.png"
                        ],
                        "slug": "samsung-galaxy-d133-5g-12-512-gb-seryy",
                        "min_price": "5933.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5933.00",
                        "max_commission": 12,
                        "monthly_payment": "1308.60",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 133,
                        "wishlist_id": null
                    },
                    {
                        "id": 40134,
                        "name": "Samsung Galaxy D134 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40134/image-134-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40134/image-134-2.png"
                        ],
                        "slug": "samsung-galaxy-d134-5g-12-512-gb-seryy",
                        "min_price": "5934.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5934.00",
                        "max_commission": 12,
                        "monthly_payment": "1310.28",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 134,
                        "wishlist_id": null
                    },
                    {
                        "id": 40135,
                        "name": "Samsung Galaxy D135 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40135/image-135-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40135/image-135-2.png"
                        ],
                        "slug": "samsung-galaxy-d135-5g-12-512-gb-seryy",
                        "min_price": "5935.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5935.00",
                        "max_commission": 12,
                        "monthly_payment": "1311.96",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 135,
                        "wishlist_id": null
                    },
                    {
                        "id": 40136,
                        "name": "Samsung Galaxy D136 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40136/image-136-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40136/image-136-2.png"
                        ],
                        "slug": "samsung-galaxy-d136-5g-12-512-gb-seryy",
                        "min_price": "5936.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5936.00",
                        "max_commission": 12,
                        "monthly_payment": "1313.64",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 136,
                        "wishlist_id": null
                    },
                    {
                        "id": 40137,
                        "name": "Samsung Galaxy D137 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40137/image-137-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40137/image-137-2.png"
                        ],
                        "slug": "samsung-galaxy-d137-5g-12-512-gb-seryy",
                        "min_price": "5937.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5937.00",
                        "max_commission": 12,
                        "monthly_payment": "1315.32",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 137,
                        "wishlist_id": null
                    },
                    {
                        "id": 40138,
                        "name": "Samsung Galaxy D138 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40138/image-138-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40138/image-138-2.png"
                        ],
                        "slug": "samsung-galaxy-d138-5g-12-512-gb-seryy",
                        "min_price": "5938.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5938.00",
                        "max_commission": 12,
                        "monthly_payment": "1317.00",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 138,
                        "wishlist_id": null
                    },
                    {
                        "id": 40139,
                        "name": "Samsung Galaxy D139 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40139/image-139-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40139/image-139-2.png"
                        ],
                        "slug": "samsung-galaxy-d139-5g-12-512-gb-seryy",
                        "min_price": "5939.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5939.00",
                        "max_commission": 12,
                        "monthly_payment": "1318.68",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 139,
                        "wishlist_id": null
                    },
                    {
                        "id": 40140,
                        "name": "Samsung Galaxy D140 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40140/image-140-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40140/image-140-2.png"
                        ],
                        "slug": "samsung-galaxy-d140-5g-12-512-gb-seryy",
                        "min_price": "5940.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5940.00",
                        "max_commission": 12,
                        "monthly_payment": "1320.36",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 140,
                        "wishlist_id": null
                    },
                    {
                        "id": 40141,
                        "name": "Samsung Galaxy D141 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40141/image-141-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40141/image-141-2.png"
                        ],
                        "slug": "samsung-galaxy-d141-5g-12-512-gb-seryy",
                        "min_price": "5941.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5941.00",
                        "max_commission": 12,
                        "monthly_payment": "1322.04",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 141,
                        "wishlist_id": null
                    },
                    {
                        "id": 40142,
                        "name": "Samsung Galaxy D142 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40142/image-142-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40142/image-142-2.png"
                        ],
                        "slug": "samsung-galaxy-d142-5g-12-512-gb-seryy",
                        "min_price": "5942.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5942.00",
                        "max_commission": 12,
                        "monthly_payment": "1323.72",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 142,
                        "wishlist_id": null
                    },
                    {
                        "id": 40143,
                        "name": "Samsung Galaxy D143 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40143/image-143-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40143/image-143-2.png"
                        ],
                        "slug": "samsung-galaxy-d143-5g-12-512-gb-seryy",
                        "min_price": "5943.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5943.00",
                        "max_commission": 12,
                        "monthly_payment": "1325.40",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 143,
                        "wishlist_id": null
                    },
                    {
                        "id": 40144,
                        "name": "Samsung Galaxy D144 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40144/image-144-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40144/image-144-2.png"
                        ],
                        "slug": "samsung-galaxy-d144-5g-12-512-gb-seryy",
                        "min_price": "5944.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5944.00",
                        "max_commission": 12,
                        "monthly_payment": "1327.08",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 144,
                        "wishlist_id": null
                    },
                    {
                        "id": 40145,
                        "name": "Samsung Galaxy D145 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40145/image-145-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40145/image-145-2.png"
                        ],
                        "slug": "samsung-galaxy-d145-5g-12-512-gb-seryy",
                        "min_price": "5945.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5945.00",
                        "max_commission": 12,
                        "monthly_payment": "1328.76",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 145,
                        "wishlist_id": null
                    },
                    {
                        "id": 40146,
                        "name": "Samsung Galaxy D146 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40146/image-146-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40146/image-146-2.png"
                        ],
                        "slug": "samsung-galaxy-d146-5g-12-512-gb-seryy",
                        "min_price": "5946.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5946.00",
                        "max_commission": 12,
                        "monthly_payment": "1330.44",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 146,
                        "wishlist_id": null
                    },
                    {
                        "id": 40147,
                        "name": "Samsung Galaxy D147 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40147/image-147-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40147/image-147-2.png"
                        ],
                        "slug": "samsung-galaxy-d147-5g-12-512-gb-seryy",
                        "min_price": "5947.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5947.00",
                        "max_commission": 12,
                        "monthly_payment": "1332.12",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 147,
                        "wishlist_id": null
                    },
                    {
                        "id": 40148,
                        "name": "Samsung Galaxy D148 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40148/image-148-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40148/image-148-2.png"
                        ],
                        "slug": "samsung-galaxy-d148-5g-12-512-gb-seryy",
                        "min_price": "5948.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5948.00",
                        "max_commission": 12,
                        "monthly_payment": "1333.80",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 148,
                        "wishlist_id": null
                    },
                    {
                        "id": 40149,
                        "name": "Samsung Galaxy D149 5G 12/512 ГБ, серый",
                        "images": [
                            "https://storage.alifshop.tj/media/images/alifshop/40149/image-149-1.png",
                            "https://storage.alifshop.tj/media/images/alifshop/40149/image-149-2.png"
                        ],
                        "slug": "samsung-galaxy-d149-5g-12-512-gb-seryy",
                        "min_price": "5949.00",
                        "discount": "0.00",
                        "discount_percent": 0,
                        "default_duration": 6,
                        "final_price": "5949.00",
                        "max_commission": 12,
                        "monthly_payment": "1335.48",
                        "gifts": [],
                        "labels": [
                            {
                                "id": "new",
                                "label": "Новинка",
                                "color": "#ffffff",
                                "bg_color": "#9833FD"
                            }
                        ],
                        "rating": "5.0",
                        "rating_count": 149,
                        "wishlist_id": null
                    }
                ],
                "view_all_url": "/popular/35",
                "position": 0.041,
                "properties": {
                    "rec_type": "popular"
                }
            }
        ]
    }
}
		},
	}
}
