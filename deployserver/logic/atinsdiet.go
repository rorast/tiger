package logic

import (
	"errors"
	"fmt"
	"reflect"
)

const (
	spinTypeMain = "main"
	spinTypeFree = "free"
)

var (
	// ErrNonPositiveBet is an error value returned by Machine.Spin method when bet is non positive.
	ErrNonPositiveBet = errors.New("bet is not positive")
	// ErrBadNLines is an error value returned by Machine.Spin method when lines isn't in (0,20]).
	ErrBadNLines = errors.New("number of lines must be in (0,20]")
	// 增加 waygame 的錯誤訊息
	ErrBadWays = errors.New("number of ways must be in (3,4,5)")

	Stops [3][5]int8
	Symbols [3][5]Symbol
	rateA int8
	symNumA  [5]int8
	rateB int8
	symNumB  [5]int8
	rateC int8
	symNumC  [5]int8
	Scatter  int8
	WinTotal      int64
)

// Machine is an Atkins Diet Slot Machine implementation.
// https://wizardofodds.com/games/slots/atkins-diet/
type Machine struct {
	rand *randomizer
}

// SpinStats represents statistics. Used mostly for simulation purpose.
type SpinStats struct {
	LinePays int64
	Scatter  int64
	Bonus    int64
	Total    int64
}

// SpinResult is a spin result structure.
type SpinResult struct {
	Type  string  `json:"type"`
	Total int64   `json:"total"`
	Stops [5]int8 `json:"stops"`
}

// New creates atkins diet Machine.
func New() *Machine {
	return &Machine{
		rand: newRandomizer(),
	}
}

// Spin does a spin with bet per line 'bet' and number of lines 'nline', returns its total win and spins.
func (m *Machine) Spin(bet, lines int) (SpinStats, []SpinResult, error) {
	var st SpinStats

	if bet <= 0 {
		return st, nil, ErrNonPositiveBet
	}

	if lines < 0 {
		return st, nil, ErrBadNLines
	}
	if len(paylines) < lines {
		return st, nil, ErrBadNLines
	}

	st, sr := m.spin(bet, lines)
	return st, sr, nil
}


// spin does spin loop in case there're free spins.
func (m *Machine) spin(bet, lines int) (SpinStats, []SpinResult) {
	// main spin
	//stops := m.reelStops()
	Stops = m.reelStops()
	fmt.Println("spin-stops = ",Stops)
	Symbols = mapSymbolsToStops(Stops)
	//Symbols := testlines

	fmt.Println("spin-symbols = ", Symbols)
	baseLineWin := caclculateLinesWin(Symbols, bet, lines)
	fmt.Println("spin-baseLineWin = ", baseLineWin)
	// scatterMultiplier := int64(0)
	scatterMultiplier := determineScattersMultiplier(Symbols)
	freeSpins := 0
	var scattersWin int64
	if 0 < scatterMultiplier {
		scattersWin = int64(bet) * int64(lines) * scatterMultiplier
		freeSpins = freeSpinsAmount
	}

	stats := SpinStats{
		LinePays: baseLineWin,
		Scatter:  scattersWin,
		Total:    baseLineWin + scattersWin,
	}
	spins := []SpinResult{
		{
			Type:  spinTypeMain,
			Total: stats.Total,
			Stops: Stops[1],
		},
	}

	// free spins
	for i := 0; i < freeSpins; i++ {
		Stops = m.reelStops()
		Symbols = mapSymbolsToStops(Stops)
		baseLineWin = caclculateLinesWin(Symbols, bet, lines) * freeSpinsMultiplier

		scattersWin = 0
		scatterMultiplier = determineScattersMultiplier(Symbols)
		if 0 < scatterMultiplier {
			scattersWin = int64(bet) * int64(lines) * scatterMultiplier * freeSpinsMultiplier
			freeSpins += freeSpinsAmount
		}

		spinWin := baseLineWin + scattersWin
		spins = append(spins, SpinResult{
			Type:  spinTypeFree,
			Total: spinWin,
			Stops: Stops[1],
		})

		stats.Bonus += spinWin
		stats.Total += spinWin
	}
	return stats, spins
}

// reelStops randomly chooses and returns reel stops.
func (m *Machine) reelStops() [3][5]int8 {
	var stops [3][5]int8

	// randomly choose reel stops for the top row
	for i := range stops[0] {
		stops[0][i] = int8(m.rand.Intn(25))
	}

	// populate two other rows respectfully
	for i := 1; i < 3; i++ {
		for j := range stops[i] {
			if stops[i-1][j] == 24 {
				stops[i][j] = 0
				continue
			}
			stops[i][j] = stops[i-1][j] + 1
		}
	}
	fmt.Println("stops values = ",stops)
	return stops
}

// mapSymbolsToStops maps symbols to reel stops.
func mapSymbolsToStops(stops [3][5]int8) [3][5]Symbol {
	var symbols [3][5]Symbol
	for reelID := 0; reelID < 5; reelID++ {
		for rowID := 0; rowID < 3; rowID++ {
			stop := stops[rowID][reelID]
			symbols[rowID][reelID] = ReelStrips[reelID][stop]
		}
	}
	return symbols
}


func Duplicate(a interface{}) (ret []interface{}) {
	va := reflect.ValueOf(a)
	for i := 0 ; i < va.Len() ; i++ {
		if i > 0 && reflect.DeepEqual(va.Index(i-1).Interface(), va.Index(i).Interface()) {
			continue
		}
		ret = append(ret, va.Index(i).Interface())
	}
	return ret
}

// caclculateLinesWin calculates win amount for symbols with bet and lines.
func caclculateLinesWin(symbols [3][5]Symbol, bet, lines int) int64 {
	var (
		win  int64
		line [5]Symbol
		//rate [6]int8  // testing
		//rate int8
		//symLen [5]int8  // 記錄 5 條滾輪中每條滾輪相同的數量
		//sym  symbol
		//scatter int8
		symNumA [5]int8  // 記錄第一個符號出現
		symNumB [5]int8  // 記錄第一個符號出現
		symNumC [5]int8  // 記錄第一個符號出現
	)
	//rate = 0 // testing
	//scatter = 0
	for i := 0; i < lines; i++ {
		payline := paylines[i][:]
		//fmt.Println("cacu...-payline = ",payline)
		for reelID := 0; reelID < 5; reelID++ {
			line[reelID] = symbols[payline[reelID]][reelID]
			//fmt.Println("cacu...-line[reelID] = ",line[reelID])
		}
		//win += int64(bet) * int64(winForLine(line))
	}

	// 20190427 Add
	rateTempA := checkSymbolLine(0,symbols)
	fmt.Println("rateTemp = ",rateTempA)
	// 算數量
	if rateTempA >= 3 {
		symNumA = checkSymNum(symbols[0][0],symbols)
		fmt.Println("reelA每行的數量",symNumA)
		// 計算 win
		winTempA := int64(winForSymbol(symbols[0][0],rateTempA))
		for i := range symNumA {
			if symNumA[i] != 0 {
				winTempA = winTempA * int64(symNumA[i])
			}
		}
		win = winTempA
	}
	if symbols[1][0] != symbols[0][0] {
		rateTempB := checkSymbolLine(1,symbols)
		fmt.Println("rateTemp = ",rateTempB)
		// 算數量
		if rateTempB >= 3 {
			symNumB = checkSymNum(symbols[1][0],symbols)
			fmt.Println("reelB每行的數量",symNumB)
		}
	}
	if symbols[1][0] != symbols[0][0] || symbols[1][0] != symbols[2][0]{
		rateTempC := checkSymbolLine(2,symbols)
		fmt.Println("rateTemp = ",rateTempC)
		// 算數量
		if rateTempC >= 3 {
			symNumC = checkSymNum(symbols[2][0],symbols)
			fmt.Println("reelC每行的數量",symNumC)
		}
	}
	// 計算 scatter
	Scatter = checkScatter(12,symbols)
	fmt.Println("scatter的數量",Scatter)
	// 計算 win 一注是 88 所以bet要先除以 88
	//win += int64(bet)/88 * int64(winForSymbol(symbols[0][0],rateTempA))
	//for i := range symNumA {
	//	if symNumA[i] != 0 {
	//		win += win * int64(symNumA[i])
	//	}
	//}

	fmt.Println("reelA 的 WIN = ",win)
	WinTotal = win
	return win
}

// 算 symbol 連線數
func checkSymbolLine(reelID int8,symbols [3][5]Symbol) int8 {
	rate := int8(0)
	if symbols[0][1] == 11 || symbols[reelID][0] == symbols[0][1] ||
		symbols[1][1] == 11 || symbols[reelID][0] == symbols[1][1] ||
		symbols[2][1] == 11 || symbols[reelID][0] == symbols[2][1] {
		rate = 2
		// 2連線後判斷3連線
		if symbols[0][2] == 11 || symbols[reelID][0] == symbols[0][2] ||
			symbols[1][2] == 11 || symbols[reelID][0] == symbols[1][2] ||
			symbols[2][2] == 11 || symbols[reelID][0] == symbols[2][2] {
			rate = 3
			// 3連線後判斷4連線
			if symbols[0][3] == 11 || symbols[reelID][0] == symbols[0][3] ||
				symbols[1][3] == 11 || symbols[reelID][0] == symbols[1][3] ||
				symbols[2][3] == 11 || symbols[reelID][0] == symbols[2][3] {
				rate = 4
				// 4連線後判斷5連線
				if symbols[0][4] == 11 || symbols[reelID][0] == symbols[0][4] ||
					symbols[1][4] == 11 || symbols[reelID][0] == symbols[1][4] ||
					symbols[2][4] == 11 || symbols[reelID][0] == symbols[2][4] {
					rate = 5
				}
			}
		}

	}
	return rate
}

func checkSymNum(sym Symbol,symbols [3][5]Symbol) [5]int8 {
	var num [5]int8
	for i  :=range symbols {
		for j :=range symbols[i] {
			//fmt.Printf("算數相同的符號數量，第 %d 列，第 %d 行，符號 %d，[i][j]值是 %d \n",i,j,symbol,symbols[i][j])
			if sym == symbols[i][j] || symbols[i][j] == 11{
				num[j]++
			}
		}
	}
	return num
}

func checkScatter(sym Symbol,symbols [3][5]Symbol) int8 {
	var scatter int8
	for i  :=range symbols {
		for j :=range symbols[i] {
			//fmt.Printf("算數相同的符號數量，第 %d 列，第 %d 行，符號 %d，[i][j]值是 %d \n",i,j,symbol,symbols[i][j])
			if symbols[i][j] == sym{
				scatter++
			}
		}
	}
	return scatter
}

func caclculateLine(value Symbol,symbols [3][5]Symbol,n int) int8 {
	rate := int8(0)
	if value == symbols[n][1] {
		if value == symbols[n][2] {
			if value == symbols[n][3] {
				if value == symbols[n][4] {
					rate = 5
				}
				rate = 4
			}
			rate = 3
		} else {

		}
	} else {
		if value == symbols[n+1][1] {
			if value == symbols[n+1][2] {
				if value == symbols[n+1][3] {
					if value == symbols[n+1][4] {
						rate = 5
					}
					rate = 4
				}
				rate = 3
			}
		}
	}
	fmt.Println("第一列的倍率 = ", rate)
	return rate
}

func winForSymbol(sym Symbol,n int8) int16{
	numPays := len(paytable)
	for payID := 0 ; payID < numPays ; payID++ {
		if paytable[payID].symbol == sym && paytable[payID].n == n {
			fmt.Println("winForSymbol 的 win = ",paytable[payID].win)
			return paytable[payID].win
		}
	}
	return 0
}

// winForLine retuns a single win for the line.
func winForLine(line [5]Symbol) int16 {
	numPays := len(paytable)
	for payID := 0; payID < numPays; payID++ {
		if checkLine(line, paytable[payID].n, paytable[payID].symbol) {
			return paytable[payID].win
		}
	}
	return 0
}

// checkLine check whether the line has n sym in a row. 這裡的 row 是橫列
func checkLine(line [5]Symbol, n int8, sym Symbol) bool {
	for i := int8(0); i < n; i++ {
		// symAtkins is wild so it is allowed to be in the line
		//if line[i] != sym && line[i] != symAtkins {
		// 判斷是否為 wild
		if line[i] != sym && line[i] != symWild {
			return false
		}
	}
	return true
}

// determineScattersMultiplier returns scatter multiplier.
func determineScattersMultiplier(symbols [3][5]Symbol) int64 {
	var n int64
	for _, row := range symbols {
		for _, s := range row {
			//if s == symScale {
			//判斷是否為 scatter
			if s == symScatter {
				n++
			}
		}
	}
	// 依據 scatter 的數量決定 freespin 的次數
	switch n {
	case 3:
		//return 5
		return 7
	case 4:
		//return 25
		return 10
	case 5:
		//return 100
		return 15
	default:
		return 0
	}
}