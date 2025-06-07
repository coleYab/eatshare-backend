package utils

import (
	"fmt"

	"github.com/coleYab/erestourant/internal/dto"
	"github.com/coleYab/erestourant/internal/store"
)

func GetTotalOrders(paylaod dto.CreateOrderDto, st *store.MenuStore) (float64, error) {
	var ans float64 = 0
	for _, u := range paylaod.Items {
		order, err := st.GetMenuByID(u.Id)
		if err != nil {
			return 0, fmt.Errorf("menu is not found with id %v", u.Id)
		}
		ans += float64(u.Qty) * order.Price

		if u.Qty > int(order.Qty) {
			return 0, fmt.Errorf("menu with id %v is out of stock", u.Id)
		}
	}

	return ans, nil
}
