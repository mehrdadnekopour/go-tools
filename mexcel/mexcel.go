package mexcel

import (
	"github.com/360EntSecGroup-Skylar/excelize"
)

// ExcelFile ...
type ExcelFile struct {
	File *excelize.File
	Path string
}

// ExcelColumn ...
type ExcelColumn int

func (ec ExcelColumn) String() string {
	switch ec {
	case ColumnA:
		return "A"
	case ColumnB:
		return "B"
	case ColumnC:
		return "C"
	case ColumnD:
		return "D"
	case ColumnE:
		return "E"
	case ColumnF:
		return "F"
	case ColumnG:
		return "G"
	case ColumnH:
		return "H"
	case ColumnI:
		return "I"
	case ColumnJ:
		return "J"
	case ColumnK:
		return "K"
	case ColumnL:
		return "L"
	case ColumnM:
		return "M"
	case ColumnN:
		return "N"
	case ColumnO:
		return "O"
	case ColumnP:
		return "P"
	case ColumnQ:
		return "Q"
	case ColumnR:
		return "R"
	case ColumnS:
		return "S"
	case ColumnT:
		return "T"
	case ColumnU:
		return "U"
	case ColumnV:
		return "V"
	case ColumnW:
		return "W"
	case ColumnX:
		return "X"
	case ColumnY:
		return "Y"
	case ColumnZ:
		return "Z"
	case ColumnAA:
		return "AA"
	case ColumnAB:
		return "AB"
	case ColumnAC:
		return "AC"
	case ColumnAD:
		return "AD"
	case ColumnAE:
		return "AE"
	case ColumnAF:
		return "AF"
	case ColumnAG:
		return "AG"
	case ColumnAH:
		return "AH"
	case ColumnAI:
		return "AI"
	case ColumnAJ:
		return "AJ"
	case ColumnAK:
		return "AK"
	case ColumnAL:
		return "AL"
	case ColumnAM:
		return "AM"
	case ColumnAN:
		return "AN"
	case ColumnAO:
		return "AO"
	case ColumnAP:
		return "AP"
	case ColumnAQ:
		return "AQ"
	case ColumnAR:
		return "AR"
	case ColumnAS:
		return "AS"
	case ColumnAT:
		return "AT"
	case ColumnAU:
		return "AU"
	case ColumnAV:
		return "AV"
	case ColumnAW:
		return "AW"
	case ColumnAX:
		return "AX"
	case ColumnAY:
		return "AY"
	case ColumnAZ:
		return "AZ"
	default:
		return "-"
	}
}

const (
	// ColumnA ....
	ColumnA ExcelColumn = iota
	// ColumnB ....
	ColumnB ExcelColumn = iota
	// ColumnC ....
	ColumnC ExcelColumn = iota
	// ColumnD ....
	ColumnD ExcelColumn = iota
	// ColumnE ....
	ColumnE ExcelColumn = iota
	// ColumnF ....
	ColumnF ExcelColumn = iota
	// ColumnG ....
	ColumnG ExcelColumn = iota
	// ColumnH ....
	ColumnH ExcelColumn = iota
	// ColumnI ....
	ColumnI ExcelColumn = iota
	// ColumnJ ....
	ColumnJ ExcelColumn = iota
	// ColumnK ....
	ColumnK ExcelColumn = iota
	// ColumnL ....
	ColumnL ExcelColumn = iota
	// ColumnM ....
	ColumnM ExcelColumn = iota
	// ColumnN ....
	ColumnN ExcelColumn = iota
	// ColumnO ....
	ColumnO ExcelColumn = iota
	// ColumnP ....
	ColumnP ExcelColumn = iota
	// ColumnQ ....
	ColumnQ ExcelColumn = iota
	// ColumnR ....
	ColumnR ExcelColumn = iota
	// ColumnS ....
	ColumnS ExcelColumn = iota
	// ColumnT ....
	ColumnT ExcelColumn = iota
	// ColumnU ....
	ColumnU ExcelColumn = iota
	// ColumnV ....
	ColumnV ExcelColumn = iota
	// ColumnW ....
	ColumnW ExcelColumn = iota
	// ColumnX ....
	ColumnX ExcelColumn = iota
	// ColumnY ....
	ColumnY ExcelColumn = iota
	// ColumnZ ....
	ColumnZ ExcelColumn = iota
	// ColumnAA ....
	ColumnAA ExcelColumn = iota
	// ColumnAB ....
	ColumnAB ExcelColumn = iota
	// ColumnAC ....
	ColumnAC ExcelColumn = iota
	// ColumnAD ....
	ColumnAD ExcelColumn = iota
	// ColumnAE ....
	ColumnAE ExcelColumn = iota
	// ColumnAF ....
	ColumnAF ExcelColumn = iota
	// ColumnAG ....
	ColumnAG ExcelColumn = iota
	// ColumnAH ....
	ColumnAH ExcelColumn = iota
	// ColumnAI ....
	ColumnAI ExcelColumn = iota
	// ColumnAJ ....
	ColumnAJ ExcelColumn = iota
	// ColumnAK ....
	ColumnAK ExcelColumn = iota
	// ColumnAL ....
	ColumnAL ExcelColumn = iota
	// ColumnAM ....
	ColumnAM ExcelColumn = iota
	// ColumnAN ....
	ColumnAN ExcelColumn = iota
	// ColumnAO ....
	ColumnAO ExcelColumn = iota
	// ColumnAP ....
	ColumnAP ExcelColumn = iota
	// ColumnAQ ....
	ColumnAQ ExcelColumn = iota
	// ColumnAR ....
	ColumnAR ExcelColumn = iota
	// ColumnAS ....
	ColumnAS ExcelColumn = iota
	// ColumnAT ....
	ColumnAT ExcelColumn = iota
	// ColumnAU ....
	ColumnAU ExcelColumn = iota
	// ColumnAV ....
	ColumnAV ExcelColumn = iota
	// ColumnAW ....
	ColumnAW ExcelColumn = iota
	// ColumnAX ....
	ColumnAX ExcelColumn = iota
	// ColumnAY ....
	ColumnAY ExcelColumn = iota
	// ColumnAZ ....
	ColumnAZ ExcelColumn = iota
)
