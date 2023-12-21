package clbnkexch

import (
	"time"
	"strings"
	"errors"
	"fmt"
	
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding"
)

var ErNoDocuments = errors.New("не задан список документов")
var ErNoPayer = errors.New("не задан плательщик документа")
var ErNoPayerBank = errors.New("не задан банк плательщика документа")
var ErNoReceiver = errors.New("не задан получатель документа")
var ErNoReceiverBank = errors.New("не задан банк получателя документа")

const (
	HEADER = "1CClientBankExchange"
	
	DEF_FORMAT_VERSION = "1.03"	
	DEF_SENDER = "Бухгалтерия предприятия, редакция 3.0"
)

type EncodingType int
func (e EncodingType) String() string {
	if e < ENCODING_TYPE_WIN || e > ENCODING_TYPE_DOS {
		return ""
	}
	v := []string{"Windows", "DOS"}
	return v[int(e)]
}
const (
	ENCODING_TYPE_WIN EncodingType = iota
	ENCODING_TYPE_DOS
)

//PayType
type PayType int
func (d PayType) String() string {
	if d < PAY_TYPE_DIG || d > PAY_TYPE_DIG {
		return ""
	}
	v := []string{"Электронно"}
	return v[int(d)]
}
const (
	PAY_TYPE_DIG PayType = iota
)

//DocumentType
type DocumentType int
func (d DocumentType) String() string {
	if d < DOCUMENT_TYPE_PP || d > DOCUMENT_TYPE_PP {
		return ""
	}
	v := []string{"Платежное поручение"}
	return v[int(d)]
}
const (
	DOCUMENT_TYPE_PP DocumentType = iota
)

//
type Firm struct {
	Name string
	Inn string
	Account string
}

type Bank struct {
	Name string
	Place string
	Bik string
	Account string
}

type BankDocumenter interface{
	GetDocumentType() DocumentType
	Render(*strings.Builder) error
	GetPayer() *Firm
}

//
type PPDocument struct {
	Num int
	Date time.Time
	Sum float64
	Payer *Firm
	PayerBank *Bank
	Receiver *Firm
	ReceiverBank *Bank
	PayType	PayType		//вид платежа
	OplType string		//вид оплаты
	Order int
	PayComment string
}
func (d *PPDocument) GetPayer() *Firm {
	return d.Payer
}

func (d *PPDocument) Render(s *strings.Builder) error {
	if d.Payer == nil {
		return ErNoPayer
	}
	if d.PayerBank == nil {
		return ErNoPayerBank
	}
	if d.Receiver == nil {
		return ErNoReceiver
	}
	if d.PayerBank == nil {
		return ErNoReceiverBank
	}
	
	//
	s.WriteString("Номер=")
	s.WriteString(fmt.Sprintf("%d", d.Num))
	s.WriteString("\n")
	
	//
	s.WriteString("Дата=")
	s.WriteString(d.Date.Format("01.02.2006"))
	s.WriteString("\n")

	//
	s.WriteString("Сумма=")
	s.WriteString(fmt.Sprintf("%.2f", d.Sum))
	s.WriteString("\n")

	//
	s.WriteString("ПлательщикСчет=")
	s.WriteString(d.Payer.Account)
	s.WriteString("\n")

	//
	s.WriteString("Плательщик=")
	s.WriteString("ИНН ")
	s.WriteString(d.Payer.Inn)	
	s.WriteString(" ")
	s.WriteString(d.Payer.Name)
	s.WriteString("\n")

	//
	s.WriteString("ПлательщикИНН=")
	s.WriteString(d.Payer.Inn)
	s.WriteString("\n")

	//
	s.WriteString("Плательщик1=")
	s.WriteString(d.Payer.Name)
	s.WriteString("\n")

	//
	s.WriteString("ПлательщикРасчСчет=")
	s.WriteString(d.Payer.Account)
	s.WriteString("\n")

	//
	s.WriteString("ПлательщикБанк1=")
	s.WriteString(d.PayerBank.Name)
	s.WriteString("\n")

	//
	s.WriteString("ПлательщикБанк2=")
	s.WriteString(d.PayerBank.Place)
	s.WriteString("\n")

	//
	s.WriteString("ПлательщикБИК=")
	s.WriteString(d.PayerBank.Bik)
	s.WriteString("\n")

	//
	s.WriteString("ПлательщикКорсчет=")
	s.WriteString(d.PayerBank.Account)
	s.WriteString("\n")

	//
	s.WriteString("ПолучательСчет=")
	s.WriteString(d.Receiver.Account)
	s.WriteString("\n")

	//
	s.WriteString("Получатель=")
	s.WriteString("ИНН ")
	s.WriteString(d.Receiver.Inn)
	s.WriteString(" ")
	s.WriteString(d.Receiver.Name)
	s.WriteString("\n")

	//
	s.WriteString("ПолучательИНН=")
	s.WriteString(d.Receiver.Inn)
	s.WriteString("\n")

	//
	s.WriteString("Получатель1=")
	s.WriteString(d.Receiver.Name)
	s.WriteString("\n")

	//
	s.WriteString("ПолучательРасчСчет=")
	s.WriteString(d.Receiver.Account)
	s.WriteString("\n")

	//
	s.WriteString("ПолучательБанк1=")
	s.WriteString(d.ReceiverBank.Name)
	s.WriteString("\n")

	//
	s.WriteString("ПолучательБанк2=")
	s.WriteString(d.ReceiverBank.Place)
	s.WriteString("\n")

	//
	s.WriteString("ПолучательБИК=")
	s.WriteString(d.ReceiverBank.Bik)
	s.WriteString("\n")

	//
	s.WriteString("ПолучательКорсчет=")
	s.WriteString(d.ReceiverBank.Account)
	s.WriteString("\n")

	//
	s.WriteString("ВидПлатежа=")
	s.WriteString(d.PayType.String())
	s.WriteString("\n")

	//
	s.WriteString("ВидОплаты=")
	s.WriteString(d.OplType)
	s.WriteString("\n")

	//
	s.WriteString("Очередность=")
	s.WriteString(fmt.Sprintf("%d", d.Order))
	s.WriteString("\n")

	//
	s.WriteString("НазначениеПлатежа=")
	s.WriteString(d.PayComment)
	s.WriteString("\n")

	//
	s.WriteString("НазначениеПлатежа1=")
	s.WriteString(d.PayComment)
	s.WriteString("\n")

	return nil
}
func (d *PPDocument) GetDocumentType() DocumentType {
	return DOCUMENT_TYPE_PP
}

type ExchFile struct {
	Version string
	EncodingType EncodingType
	Sender string
	CreateDate time.Time
	DateFrom time.Time
	DateTo time.Time
	DocumentTypes []DocumentType
	Documents []BankDocumenter
}
func NewExchFile(documentTypes []DocumentType, documents []BankDocumenter) *ExchFile{
	return &ExchFile{DocumentTypes: documentTypes,
		Documents: documents,	
		EncodingType: ENCODING_TYPE_WIN,
		Sender: DEF_SENDER,
	}
}

func (f *ExchFile) Render() ([]byte, error) {
	if f.Documents == nil || len(f.Documents) == 0 {
		return []byte{}, ErNoDocuments
	}
	empty_time := time.Time{}
	s := strings.Builder{}
	
	//
	s.WriteString(HEADER)
	s.WriteString("\n")
	
	//
	s.WriteString("ВерсияФормата=")
	if f.Version == "" {
		f.Version = DEF_FORMAT_VERSION
	}
	s.WriteString(f.Version)
	s.WriteString("\n")

	//
	s.WriteString("Кодировка=")
	s.WriteString(f.EncodingType.String())
	s.WriteString("\n")

	//
	s.WriteString("Отправитель=")
	s.WriteString(f.Sender)
	s.WriteString("\n")
	
	//
	s.WriteString("ДатаСоздания=")
	if f.CreateDate == empty_time {
		f.CreateDate = time.Now()
	}
	s.WriteString(f.CreateDate.Format("01.02.2006"))
	s.WriteString("\n")

	//
	s.WriteString("ВремяСоздания=")
	s.WriteString(f.CreateDate.Format("15:04:05"))
	s.WriteString("\n")
	
	//
	s.WriteString("ДатаНачала=")
	if f.DateFrom == empty_time {
		f.DateFrom = f.CreateDate
	}
	s.WriteString(f.DateFrom.Format("01.02.2006"))
	s.WriteString("\n")

	//
	s.WriteString("ДатаКонца=")
	if f.DateTo == empty_time {
		f.DateTo = f.CreateDate
	}
	s.WriteString(f.DateTo.Format("01.02.2006"))
	s.WriteString("\n")
	
	//
	s.WriteString("РасчСчет=")
	if f.Documents[0].GetPayer() == nil {
		return []byte{}, ErNoPayer
	}
	s.WriteString(f.Documents[0].GetPayer().Account)
	s.WriteString("\n")
	
	//
	for _,t := range f.DocumentTypes {
		s.WriteString("Документ=")
		s.WriteString(t.String())
		s.WriteString("\n")
	}
	
	//
	for _, doc := range f.Documents {
		s.WriteString("СекцияДокумент=")		
		s.WriteString(doc.GetDocumentType().String())
		s.WriteString("\n")
		
		if err := doc.Render(&s); err != nil {
			return []byte{}, err
		}
		
		s.WriteString("КонецДокумента")
		s.WriteString("\n")
	}

	//
	s.WriteString("КонецФайла")
	s.WriteString("\n")
	
	//
	var enc *encoding.Encoder
	if f.EncodingType == ENCODING_TYPE_WIN {
		enc = charmap.Windows1251.NewEncoder()
	}else{
		enc = charmap.CodePage866.NewEncoder()
	}
	out, err := enc.String(s.String())
	if err != nil {
		return []byte{}, err
	}
	return []byte(out), nil
}

