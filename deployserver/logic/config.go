package logic

type Symbol int8

const (
	// symAtkins is a wild
	symNine         = Symbol(1)
	symTen          = Symbol(2)
	symJ            = Symbol(3)
	symQ            = Symbol(4)
	symK            = Symbol(5)
	symA            = Symbol(6)
	symLion         = Symbol(7)
	symFish         = Symbol(8)
	symFurnace      = Symbol(9)
	symJewel        = Symbol(10)
	symWild         = Symbol(11)
	symScatter      = Symbol(12)

	freeSpinsAmount     = 10
	freeSpinsMultiplier = 3
)

var paytable = [30]struct {
	win    int16
	n      int8
	symbol Symbol
}{
	{500, 5, symJewel},
	{400, 5, symFurnace},
	{250, 5, symFish},
	{100, 5, symLion},
	{50, 5, symA},
	{50, 5, symK},
	{50, 5, symQ},
	{50, 5, symJ},
	{50, 5, symTen},
	{50, 5, symNine},
	{100, 4, symJewel},
	{80, 4, symFurnace},
	{50, 4, symFish},
	{20, 4, symLion},
	{10, 4, symA},
	{10, 4, symK},
	{10, 4, symQ},
	{10, 4, symJ},
	{10, 4, symTen},
	{10, 4, symNine},
	{50, 3, symJewel},
	{40, 3, symFurnace},
	{25, 3, symFish},
	{10, 3, symLion},
	{5, 3, symA},
	{5, 3, symK},
	{5, 3, symQ},
	{5, 3, symJ},
	{5, 3, symTen},
	{5, 3, symNine},
}

// 未來要 32 要改成 25 ， 且用隨機產生
var ReelStrips = [5][25]Symbol{
	{symNine,  symTen,     symFurnace, symNine, symA,     symJ,     symLion,  symNine,    symTen,  symJ,    symQ,   symK,       symNine, symTen,     symFish, symQ,       symK,   symScatter, symNine, symTen,  symJewel, symNine,    symTen,   symJ,    symQ},
	{symJewel, symWild,    symA,       symNine, symTen,   symFish,  symNine,  symTen,     symJ,    symQ,    symK,   symFurnace, symNine, symFurnace, symJ,    symQ,       symK,   symScatter, symNine, symNine, symTen,   symA,       symK,     symLion, symJ},
	{symK,     symA,       symJ,       symA,    symJewel, symJ,     symJewel, symNine,    symNine, symNine, symTen, symJ,       symQ,    symWild,    symLion, symScatter, symTen, symJ,       symTen,  symWild, symJ,     symFurnace, symNine,  symTen,  symFish},
	{symJ,     symNine,    symJ,       symTen,  symQ,     symJewel, symQ,     symScatter, symA,    symNine, symTen, symJ,       symFish, symQ,       symK,    symWild,    symTen, symLion,    symLion, symTen,  symJ,     symFurnace, symTen,   symA,    symK},
	{symWild,  symScatter, symA,       symTen,  symNine,  symFish,  symJ,     symQ,       symNine, symNine, symTen, symFish,    symK,    symQ,       symWild, symNine,    symTen, symJ,       symK,    symA,    symLion,  symFurnace, symJewel, symJ,    symScatter},
}
//const (
//	// symAtkins is a wild
//	symAtkins       = symbol(1)
//	symSteak        = symbol(2)
//	symHam          = symbol(3)
//	symBuffaloWings = symbol(4)
//	symSausage      = symbol(5)
//	symEggs         = symbol(6)
//	symButter       = symbol(7)
//	symCheese       = symbol(8)
//	symBacon        = symbol(9)
//	symMayonnaise   = symbol(10)
//	symScale        = symbol(11)
//
//	freeSpinsAmount     = 10
//	freeSpinsMultiplier = 3
//)

//var paytable = [34]struct {
//	win    int16
//	n      int8
//	symbol symbol
//}{
//	{5000, 5, symAtkins},
//	{1000, 5, symSteak},
//	{500, 4, symAtkins},
//	{500, 5, symHam},
//	{300, 5, symBuffaloWings},
//	{200, 4, symSteak},
//	{200, 5, symSausage},
//	{200, 5, symEggs},
//	{150, 4, symHam},
//	{100, 4, symBuffaloWings},
//	{100, 5, symButter},
//	{100, 5, symCheese},
//	{75, 4, symSausage},
//	{75, 4, symEggs},
//	{50, 3, symAtkins},
//	{50, 4, symButter},
//	{50, 4, symCheese},
//	{50, 5, symBacon},
//	{50, 5, symMayonnaise},
//	{40, 3, symSteak},
//	{30, 3, symHam},
//	{25, 3, symBuffaloWings},
//	{25, 4, symBacon},
//	{25, 4, symMayonnaise},
//	{20, 3, symSausage},
//	{20, 3, symEggs},
//	{15, 3, symButter},
//	{15, 3, symCheese},
//	{10, 3, symBacon},
//	{10, 3, symMayonnaise},
//	{2, 5, symAtkins},
//	{3, 2, symSteak},
//	{2, 2, symHam},
//	{2, 2, symBuffaloWings},
//}

//var reelStrips = [5][32]symbol{
//	{symScale, symMayonnaise, symHam, symSausage, symBacon, symEggs, symCheese, symMayonnaise, symSausage, symButter, symBuffaloWings, symBacon, symEggs, symMayonnaise, symSteak, symBuffaloWings, symButter, symCheese, symEggs, symAtkins, symBacon, symMayonnaise, symHam, symCheese, symEggs, symScale, symButter, symBacon, symSausage, symBuffaloWings, symSteak, symButter},
//	{symMayonnaise, symBuffaloWings, symSteak, symSausage, symCheese, symMayonnaise, symHam, symButter, symBacon, symSteak, symSausage, symMayonnaise, symHam, symAtkins, symButter, symEggs, symCheese, symBacon, symSausage, symBuffaloWings, symScale, symMayonnaise, symButter, symCheese, symBacon, symEggs, symBuffaloWings, symMayonnaise, symSteak, symHam, symCheese, symBacon},
//	{symHam, symButter, symEggs, symScale, symCheese, symMayonnaise, symButter, symHam, symSausage, symBacon, symSteak, symBuffaloWings, symButter, symMayonnaise, symCheese, symSausage, symEggs, symBacon, symMayonnaise, symBuffaloWings, symHam, symSausage, symBacon, symCheese, symEggs, symAtkins, symBuffaloWings, symBacon, symButter, symCheese, symMayonnaise, symSteak},
//	{symHam, symCheese, symAtkins, symScale, symButter, symBacon, symCheese, symSausage, symSteak, symEggs, symBacon, symMayonnaise, symSausage, symCheese, symButter, symHam, symMayonnaise, symBacon, symBuffaloWings, symSausage, symCheese, symEggs, symButter, symBuffaloWings, symBacon, symMayonnaise, symEggs, symHam, symSausage, symSteak, symMayonnaise, symBacon},
//	{symBacon, symScale, symSteak, symHam, symCheese, symSausage, symButter, symBacon, symBuffaloWings, symCheese, symSausage, symHam, symButter, symSteak, symMayonnaise, symEggs, symSausage, symHam, symAtkins, symButter, symBuffaloWings, symMayonnaise, symEggs, symHam, symBacon, symButter, symSteak, symMayonnaise, symSausage, symEggs, symCheese, symBuffaloWings},
//}

var paylines = [20][5]int8{
	{1, 1, 1, 1, 1},
	{0, 0, 0, 0, 0},
	{2, 2, 2, 2, 2},
	{0, 1, 2, 1, 0},
	{2, 1, 0, 1, 2},
	{1, 0, 0, 0, 1},
	{1, 2, 2, 2, 1},
	{0, 0, 1, 2, 2},
	{2, 2, 1, 0, 0},
	{1, 0, 1, 2, 1},
	{1, 2, 1, 0, 1},
	{0, 1, 1, 1, 0},
	{2, 1, 1, 1, 2},
	{0, 1, 0, 1, 0},
	{2, 1, 2, 1, 2},
	{1, 1, 0, 1, 1},
	{1, 1, 2, 1, 1},
	{0, 0, 2, 0, 0},
	{2, 2, 0, 2, 2},
	{0, 2, 2, 2, 0},
}

var testlines = [3][5]Symbol{
	{10, 2, 2, 10, 12},
	{6, 10, 8, 10, 11},
	{2, 5, 10, 10, 12},
}
