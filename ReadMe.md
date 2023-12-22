# Формирование платежных поручений для банка.
Реализует формат выгрузки [1CClientBankExchange](https://v8.1c.ru/tekhnologii/obmen-dannymi-i-integratsiya/standarty-i-formaty/standart-obmena-s-sistemami-klient-banka/formaty-obmena/) платежных документов в банк.

### Как использовать:
```go
	import (
		"os"
		"time"
		
		"github.com/dronm/clbnkexch"
	)
	//список документов
	documents := []clbnkexch.BankDocumenter{&clbnkexch.PPDocument{Num: 1,
		Date: time.Now(),
		Sum: 175000,
		Payer: &clbnkexch.Firm{Name: `ООО "Рога и Копыта"`, Inn: "1234567891", Account: "12345678901234567890"},
		PayerBank: &clbnkexch.Bank{Name: "Банк ОАО", Place: "г. Москва", Bik :"123456789", Account:"12345678901234567890"},
		Receiver: &clbnkexch.Firm{Name: `ИП Иванов А.А.`, Inn: "111122223344", Account: "12345678901234567890"},
		ReceiverBank: &clbnkexch.Bank{Name: "Банк2 ОАО", Place: "г. Москва", Bik: "123456789", Account: "12345678901234567890"},
		PayType: clbnkexch.PAY_TYPE_DIG,
		OplType: "01",
		Order: 5,
		PayComment: "За товары, по счету №125 на сумму 175000-00",		
	}}
	//Типы документов: п/п
	doc_types := []clbnkexch.DocumentType{clbnkexch.DOCUMENT_TYPE_PP}
	
	//объект выгрузки
	fl := clbnkexch.NewExchFile(doc_types, documents)
	
	//сохраним в файл
	f, err := os.Create("cl_to_bank.txt")
	if err != nil {
		panic("os.Create() failed: %v", err)
	}
	defer f.Close()
	
	//вывод в файл
	b, err := fl.Render()
	if err != nil {
		panic("ExchFile.String() failed: %v", err)
	}	
	f.Write(b)	
```
	


