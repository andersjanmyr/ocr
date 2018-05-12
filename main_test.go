package main

import (
	"reflect"
	"testing"
)

type fixture struct {
	input    []string
	expected map[string]string
}

func TestParseAnnotations(t *testing.T) {
	fixtures := []fixture{
		fixture{
			input: []string{"amazon.com", "Final", "Details", "for", "Order", "#", "111_3750649-8093818", "rint", "this", "pa", "r", "your", "records", "Order", "Placed:", "November", "11,", "2017", "Amazon.com", "order", "number:", "111-3750649-8093818", "Order", "Total:", "$43.99", "Shipped", "on", "November", "12,", "2017", "Price", "$40.36", "Items", "Ordered", "1", "of:", "Kubernetes:", "Up", "and", "Running:", "Dive", "into", "the", "Future", "of", "Infrastructure,", "Hightower,", "Kelsey", "Sold", "by:", "Amazon.com", "Services,", "Inc.", "Condition:", "New", "Shipping", "Address:", "Anders", "Janmyr", "592", "LOMA", "VERDE", "AVE", "PALO", "ALTO,", "CA", "94306-3032", "United", "States", "Item(s)", "Subtotal:", "$40.36", "Shipping", "&", "Handling:", "$0.00", "Total", "before", "tax:", "$40.36", "Sales", "Tax:", "$3.63", "Total", "for", "This", "Shipment:", "$43.99", "Shipping", "Speed:", "Two-Day", "Shipping", "Payment", "information", "Payment", "Method:", "MasterCard", "|", "Last", "digits:", "5931", "Item(s)", "Subtotal:", "$40.36", "Shipping", "&", "Handling:", "$0.00", "Billing", "address", "Anders", "Janmyr", "592", "LOMA", "VERDE", "AVE", "PALO", "ALTO,", "CA", "94306-3032", "United", "States", "Total", "before", "tax:", "$40.36", "Estimated", "tax", "to", "be", "collected:", "$3.63", "Grand", "Total:", "$43.99", "Credit", "Card", "transactions", "MasterCard", "ending", "in", "5931:", "November", "12,", "2017:", "$43.99", "To", "view", "the", "status", "of", "your", "order,", "return", "to", "Order", "Summary.", "nditions", "of", "Use", "l", "Privacy", "Notice", "©", "1996-2018", ",", "Amazon.com", ",", "Inc.", "or", "its", "affiliates"},
			expected: map[string]string{
				"Total": "$43.99",
				"Tax":   "$3.69",
			},
		},
	}
	for _, f := range fixtures {
		output := ParseAnnotations(f.input)
		if !reflect.DeepEqual(output, f.expected) {
			t.Error(output, f.expected)
		}
	}
}
