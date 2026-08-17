package repository

type PopularResponse struct {
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

type PopularResult struct {
	Meta     Meta              `json:"meta"`
	Response []PopularResponse `json:"response"`
}

type PopularRepository interface {
	GetPopular() PopularResult
}

type popularRepository struct{}

func NewPopularYou() PopularRepository {
	return &popularRepository{}
}

// func (r *popularRepository) GetPopular() PopularResult {
// 	return PopularResult{
// 		Meta: Meta{
// 			Error:      false,
// 			Message:    "",
// 			StatusCode: 200,
// 		},

// 		Response: []PopularResponse{
// 			{
// 				ID:              62729,
// 				Name:            "Диван Арткор Пуфак 290х105 см, бежевый",
// 				Slug:            "divan-artkor-pufak-290h105-sm-bezhevyy",
// 				Rating:          "0.0",
// 				RatingCount:     0,
// 				MinPrice:        "12800.00",
// 				FinalPrice:      "12800.00",
// 				DefaultDuration: 24,
// 				MaxCommission:   37,
// 				MaxConditionID:  25296,
// 				Images: []string{
// 					"https://storage.alifshop.tj/media/images/alifshop/62729/divan-artkor-pufak-290h105-sm-bezhevyy-1784725701437-ttuxWaJO.png",
// 					"https://storage.alifshop.tj/media/images/alifshop/62729/divan-artkor-pufak-290h105-sm-bezhevyy-1784725692887-V4MQUWt6.png",
// 				},
// 				WishlistID: nil,
// 				Discount:   "0.00",
// 				Gifts:      []any{},
// 				Labels: []Label{
// 					{
// 						ID:      "new",
// 						Label:   "Новинка",
// 						Color:   "#ffffff",
// 						BgColor: "#9833FD",
// 					},
// 				},
// 				DiscountPercent: 0,
// 				MonthlyPayment:  "730.66",
// 			},
// 			{
// 				ID:              62729,
// 				Name:            "Диван Арткор Пуфак 290х105 см, бежевый",
// 				Slug:            "divan-artkor-pufak-290h105-sm-bezhevyy",
// 				Rating:          "0.0",
// 				RatingCount:     0,
// 				MinPrice:        "12800.00",
// 				FinalPrice:      "12800.00",
// 				DefaultDuration: 24,
// 				MaxCommission:   37,
// 				MaxConditionID:  25296,
// 				Images: []string{
// 					"https://storage.alifshop.tj/media/images/alifshop/62729/divan-artkor-pufak-290h105-sm-bezhevyy-1784725701437-ttuxWaJO.png",
// 					"https://storage.alifshop.tj/media/images/alifshop/62729/divan-artkor-pufak-290h105-sm-bezhevyy-1784725692887-V4MQUWt6.png",
// 				},
// 				WishlistID: nil,
// 				Discount:   "0.00",
// 				Gifts:      []any{},
// 				Labels: []Label{
// 					{
// 						ID:      "new",
// 						Label:   "Новинка",
// 						Color:   "#ffffff",
// 						BgColor: "#9833FD",
// 					},
// 				},
// 				DiscountPercent: 0,
// 				MonthlyPayment:  "730.66",
// 			},
// 			{
// 				ID:              62729,
// 				Name:            "Диван Арткор Пуфак 290х105 см, бежевый",
// 				Slug:            "divan-artkor-pufak-290h105-sm-bezhevyy",
// 				Rating:          "0.0",
// 				RatingCount:     0,
// 				MinPrice:        "12800.00",
// 				FinalPrice:      "12800.00",
// 				DefaultDuration: 24,
// 				MaxCommission:   37,
// 				MaxConditionID:  25296,
// 				Images: []string{
// 					"https://storage.alifshop.tj/media/images/alifshop/62729/divan-artkor-pufak-290h105-sm-bezhevyy-1784725701437-ttuxWaJO.png",
// 					"https://storage.alifshop.tj/media/images/alifshop/62729/divan-artkor-pufak-290h105-sm-bezhevyy-1784725692887-V4MQUWt6.png",
// 				},
// 				WishlistID: nil,
// 				Discount:   "0.00",
// 				Gifts:      []any{},
// 				Labels: []Label{
// 					{
// 						ID:      "new",
// 						Label:   "Новинка",
// 						Color:   "#ffffff",
// 						BgColor: "#9833FD",
// 					},
// 				},
// 				DiscountPercent: 0,
// 				MonthlyPayment:  "730.66",
// 			}, {
// 				ID:              62729,
// 				Name:            "Диван Арткор Пуфак 290х105 см, бежевый",
// 				Slug:            "divan-artkor-pufak-290h105-sm-bezhevyy",
// 				Rating:          "0.0",
// 				RatingCount:     0,
// 				MinPrice:        "12800.00",
// 				FinalPrice:      "12800.00",
// 				DefaultDuration: 24,
// 				MaxCommission:   37,
// 				MaxConditionID:  25296,
// 				Images: []string{
// 					"https://storage.alifshop.tj/media/images/alifshop/62729/divan-artkor-pufak-290h105-sm-bezhevyy-1784725701437-ttuxWaJO.png",
// 					"https://storage.alifshop.tj/media/images/alifshop/62729/divan-artkor-pufak-290h105-sm-bezhevyy-1784725692887-V4MQUWt6.png",
// 				},
// 				WishlistID: nil,
// 				Discount:   "0.00",
// 				Gifts:      []any{},
// 				Labels: []Label{
// 					{
// 						ID:      "new",
// 						Label:   "Новинка",
// 						Color:   "#ffffff",
// 						BgColor: "#9833FD",
// 					},
// 				},
// 				DiscountPercent: 0,
// 				MonthlyPayment:  "730.66",
// 			}, {
// 				ID:              62729,
// 				Name:            "Диван Арткор Пуфак 290х105 см, бежевый",
// 				Slug:            "divan-artkor-pufak-290h105-sm-bezhevyy",
// 				Rating:          "0.0",
// 				RatingCount:     0,
// 				MinPrice:        "12800.00",
// 				FinalPrice:      "12800.00",
// 				DefaultDuration: 24,
// 				MaxCommission:   37,
// 				MaxConditionID:  25296,
// 				Images: []string{
// 					"https://storage.alifshop.tj/media/images/alifshop/62729/divan-artkor-pufak-290h105-sm-bezhevyy-1784725701437-ttuxWaJO.png",
// 					"https://storage.alifshop.tj/media/images/alifshop/62729/divan-artkor-pufak-290h105-sm-bezhevyy-1784725692887-V4MQUWt6.png",
// 				},
// 				WishlistID: nil,
// 				Discount:   "0.00",
// 				Gifts:      []any{},
// 				Labels: []Label{
// 					{
// 						ID:      "new",
// 						Label:   "Новинка",
// 						Color:   "#ffffff",
// 						BgColor: "#9833FD",
// 					},
// 				},
// 				DiscountPercent: 0,
// 				MonthlyPayment:  "730.66",
// 			}, {
// 				ID:              62729,
// 				Name:            "Диван Арткор Пуфак 290х105 см, бежевый",
// 				Slug:            "divan-artkor-pufak-290h105-sm-bezhevyy",
// 				Rating:          "0.0",
// 				RatingCount:     0,
// 				MinPrice:        "12800.00",
// 				FinalPrice:      "12800.00",
// 				DefaultDuration: 24,
// 				MaxCommission:   37,
// 				MaxConditionID:  25296,
// 				Images: []string{
// 					"https://storage.alifshop.tj/media/images/alifshop/62729/divan-artkor-pufak-290h105-sm-bezhevyy-1784725701437-ttuxWaJO.png",
// 					"https://storage.alifshop.tj/media/images/alifshop/62729/divan-artkor-pufak-290h105-sm-bezhevyy-1784725692887-V4MQUWt6.png",
// 				},
// 				WishlistID: nil,
// 				Discount:   "0.00",
// 				Gifts:      []any{},
// 				Labels: []Label{
// 					{
// 						ID:      "new",
// 						Label:   "Новинка",
// 						Color:   "#ffffff",
// 						BgColor: "#9833FD",
// 					},
// 				},
// 				DiscountPercent: 0,
// 				MonthlyPayment:  "730.66",
// 			}, {
// 				ID:              62729,
// 				Name:            "Диван Арткор Пуфак 290х105 см, бежевый",
// 				Slug:            "divan-artkor-pufak-290h105-sm-bezhevyy",
// 				Rating:          "0.0",
// 				RatingCount:     0,
// 				MinPrice:        "12800.00",
// 				FinalPrice:      "12800.00",
// 				DefaultDuration: 24,
// 				MaxCommission:   37,
// 				MaxConditionID:  25296,
// 				Images: []string{
// 					"https://storage.alifshop.tj/media/images/alifshop/62729/divan-artkor-pufak-290h105-sm-bezhevyy-1784725701437-ttuxWaJO.png",
// 					"https://storage.alifshop.tj/media/images/alifshop/62729/divan-artkor-pufak-290h105-sm-bezhevyy-1784725692887-V4MQUWt6.png",
// 				},
// 				WishlistID: nil,
// 				Discount:   "0.00",
// 				Gifts:      []any{},
// 				Labels: []Label{
// 					{
// 						ID:      "new",
// 						Label:   "Новинка",
// 						Color:   "#ffffff",
// 						BgColor: "#9833FD",
// 					},
// 				},
// 				DiscountPercent: 0,
// 				MonthlyPayment:  "730.66",
// 			},
// 		},
// 	}
// }

func (r *popularRepository) GetPopular() PopularResult {
	return PopularResult{
		Meta: Meta{
			Error:      false,
			Message:    "",
			StatusCode: 200,
		},
		Response: []PopularResponse{},
	}
}
