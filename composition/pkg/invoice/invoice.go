package invoice

import (
	"github.com/AJRDRGZ/composition/pkg/customer"
	"github.com/AJRDRGZ/composition/pkg/invoiceitem"
)

// Invoice is the struct of a invoice
type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer
	items   invoiceitem.Items
}

// New returs a new Invoice
func New(country, city string, client customer.Customer, items invoiceitem.Items) Invoice {
	return Invoice{
		country: country,
		city:    city,
		client:  client,
		items:   items,
	}
}

// SetTotal is the setter of Invoice.total
func (i *Invoice) SetTotal() {
	i.total = i.items.Total()
}
