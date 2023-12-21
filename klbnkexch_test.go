package clbnkexch

import(
	//"fmt"
	"os"
	"testing"
	"time"
)

func TestRender(t *testing.T) {
	documents := []BankDocumenter{&PPDocument{Num: 1,
		Date: time.Now(),
		Sum: 175000,
		Payer: &Firm{Name: `ООО "Рога и Копыта"`, Inn: "1234567891", Account: "12345678901234567890"},
		PayerBank: &Bank{Name: "Объёббанк ОАО", Place: "г. Москва", Bik :"123456789", Account:"12345678901234567890"},
		Receiver: &Firm{Name: `ИП Иванов А.А.`, Inn: "111122223344", Account: "12345678901234567890"},
		ReceiverBank: &Bank{Name: "Наёббанк ОАО", Place: "г. Москва", Bik: "123456789", Account: "12345678901234567890"},
		PayType: PAY_TYPE_DIG,
		OplType: "01",
		Order: 5,
		PayComment: "За товары, по счету №125 на сумму 175000-00",		
	}}
	doc_types := []DocumentType{DOCUMENT_TYPE_PP}
	
	fl := NewExchFile(doc_types, documents)
	f, err := os.Create("cl_to_bank.txt")
	if err != nil {
		t.Fatalf("os.Create() failed: %v", err)
	}
	b, err := fl.Render()
	if err != nil {
		t.Fatalf("ExchFile.String() failed: %v", err)
	}	
	f.Write(b)	
	f.Close()
	//fmt.Println()
}

