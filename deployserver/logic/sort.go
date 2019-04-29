package logic

import (
	"fmt"
	"math/rand"
	"reflect"
	"tiger/deployserver/pipeline"
	"time"
)

func enum() {

}
// 各符號的參數
const(
	N = 1
	T= 2
	A = 3
	J = 4
	Q = 5
	K = 6
	P1 = 7
	P2 = 8
	P3 = 9
	P4 = 10
	//P5 = 11
	W = 11
	D = 12
)
// 各個符號賠率(要設計資料表)
const(
	threeNormal = 5
	fourNormal = 10
	fiveNormal = 50
	threeP1 = 10
	fourP1 = 20
	fiveP1 = 100
	threeP2 = 25
	fourP2 = 50
	fiveP2 = 250
	threeP3 = 40
	fourP3 = 80
	fiveP3 = 400
	threeP4 = 50
	fourP4 = 100
	fiveP4 = 500
	//threeP5 = 100
	//fourP5 = 200
	//fiveP5 = 1000
)
//func main() {
//	//result = make([]int, len(result))
//	num := make([]int, 15)
//
//	// 設定公司 依照符號的判斷結果 * 賠率 * 每輪組相同的加乘(2個*2，3個*3)
//
//	checkPay(num)
//
//	fmt.Println("n = ", N)
//	//const filename = "small.in"
//	//const n = 64
//	//// 寫文件
//	//file, err := os.Create(filename)
//	//if err != nil {
//	//	panic(err)
//	//}
//	//defer file.Close()
//	//
//	//p := pipeline.RandomSource(n)
//	//writer := bufio.NewWriter(file)
//	//pipeline.WriterSink(writer, p)
//	//// 如果有用bufio的話，要用flush，不然檔案長度不會是我們設定的
//	//writer.Flush()
//	//
//	//// 將檔案讀進來
//	//file, err = os.Open(filename)
//	//if err != nil {
//	//	panic(err)
//	//}
//	//defer file.Close()
//	//
//	//p = pipeline.ReaderSource(bufio.NewReader(file), -1)
//	//count := 0
//	//for v := range p {
//	//	fmt.Println(v)
//	//	count++
//	//	if count >= 100 {
//	//		break
//	//	}
//	//}
//
//}

func checkPay(result []int) int {

	result = make([]int, len(result))

	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	//先隨機計算結果給前端 0 :=沒中，1 := 中彩金，2 :=freeSpin, 3 := bigWin, 4 := jackPort

	//neighbor := [5]bool{false,false,false,false,false}
	//

	//oneM  := 0
	//doubleM := 0
	//threeM := 0
	// 15個滾輪，13個symbol 1-13
	reelNum := 0  // 記錄目前是計算那列滾輪
	//var sliceA [3]int
	sliceA := make([]int, 3)
	sliceB := make([]int, 3)
	sliceC := make([]int, 3)
	sliceD := make([]int, 3)
	sliceE := make([]int, 3)
	//sliceF := make([]int, 3)

	for i:=0 ; i < len(result) ; i++ {
		result[i] = r.Intn(13)+1
	}

	for i, v :=range result {
		if i % 3 == 0 {
			reelNum ++
			fmt.Println("reelNum = ",reelNum)
		}
		// 第一個滾輪的值
		if(reelNum == 1){
			// 加入元素
			sliceA[i%3] = v
			//sliceA = append(sliceA,v)
			//fmt.Println("sliceA = ",sliceA,v)
		}
		// 第二個滾輪的值
		if(reelNum == 2){
			// 加入元素
			sliceB[i%3] = v
			//sliceA = append(sliceA,v)
			//fmt.Println("sliceB = ",sliceB,v)
		}
		// 第三個滾輪的值
		if(reelNum == 3){
			// 加入元素
			sliceC[i%3] = v
			//sliceA = append(sliceA,v)
			//fmt.Println("sliceC = ",sliceC,v)
		}
		// 第四個滾輪的值
		if(reelNum == 3){
			// 加入元素
			sliceD[i%3] = v
			//sliceA = append(sliceA,v)
			//fmt.Println("sliceD = ",sliceD,v)
		}
		// 第五個滾輪的值
		if(reelNum == 3){
			// 加入元素
			sliceE[i%3] = v
			//sliceA = append(sliceA,v)
			//fmt.Println("sliceE = ",sliceE,v)
		}

	}
	fmt.Println("slice = \n",sliceA, "\n",sliceB, "\n",sliceC, "\n",sliceD, "\n",sliceE)
	IntSliceTheSame(sliceA, sliceB, sliceC, sliceD, sliceE)
	//sliceF, &oneM, &doubleM, &threeM =  IntSliceTheSame(sliceA,sliceB)
	//fmt.Println("回傳第一組比較 = ", sliceF, oneM, doubleM, threeM)

	if len(result) != 0 {
		return 100
	} else {
		return 0
	}

}

func IntSliceReflectEqual(a,b []int) bool {
	return reflect.DeepEqual(a, b)
}

func IntSliceTheSame(a,b,c,d,e []int) {

	one := 0
	double := 0
	three := 0

	// 暫存相同的值
	sliceTemp := make([]int, 3)
	// 滾輪組 A
	for i, v :=range  a {
		// 滾輪組 B
		for j, z := range b {
			if v == z || v == W {

				// 判斷是否有三連線(至少要三連線才開始計算paytable)
				for k, y := range c {
					if v == y || v == W {
						//  判斷是否出現第一次
						if one <= 1 {
							one++
							break
						}
						if one >= 1 {
							double++
						}
						if double >= 1 {
							three++
						}

					}
					// 滾輪組 A 及 B 出現相同
					fmt.Println("值 = ",a[i],b[j],c[k])
				}


				// 滾輪組A不能出現 Wild
				if v == W {
					break
				}


				// 將相同的值加入暫存組，只要判斷第一組是否有相同的即可
				sliceTemp[i] = v

			}
		}
	}
	fmt.Println("出現幾次的值", one, double, three)

}

//func IntSliceTheSame(a,b []int) (sliceTempR []int,oneR *int,doubleR *int,threeR *int) {
//
//	one := 0
//	double := 0
//	three := 0
//
//	// 暫存相同的值
//	sliceTemp := make([]int, 3)
//	// 滾輪組 A
//	for i, v :=range  a {
//		// 滾輪組 B
//		for j, z := range b {
//			if v == z || v == W {
//				//  判斷是否出現第一次
//				if one <= 1 {
//
//				}
//				if one >= 1 {
//					double++
//				}
//				if double >= 1 {
//					three++
//				}
//
//				// 滾輪組A不能出現 Wild
//				if v == W {
//					break
//				}
//
//				// 滾輪組 A 及 B 出現相同
//				fmt.Println("值 = ",a[i],b[j])
//				// 將相同的值加入暫存組，只要判斷第一組是否有相同的即可
//				sliceTemp[i] = v
//
//			}
//		}
//	}
//	fmt.Println("出現幾次的值", one, double, three)
//	return sliceTemp, &one, &double, &three
//}

func mergeDemo() {
	// Creates a slice of int
	//a := []int{3,6,2,1,9,10,8}
	//sort.Ints(a)
	//
	//for _ , v := range a {
	//	fmt.Println(v)
	//}

	// 宣告一個 p 作為 channel
	//p := pipeline.ArraySource(3,2,6,7,4)  // 單純排序
	// 使用 InMemSort 包起來
	//p := pipeline.InMemSort(pipeline.ArraySource(3,2,6,7,4))

	// 使用 merge 再包一層
	p := pipeline.Merge(
		pipeline.InMemSort(pipeline.ArraySource(3,2,6,7,4)),
		pipeline.InMemSort(pipeline.ArraySource(7,4,0,3,2,13,8)))


	// 第一種寫法 ， GO 沒有 while ， for 可用來作為go中的 while 來使用
	//for {
	//	if num, ok := <- p ; ok {
	//		fmt.Println(num)
	//	} else {
	//		break
	//	}
	//
	//}
	// 第二種寫法
	for v := range p {
		fmt.Println(v)
	}
}
