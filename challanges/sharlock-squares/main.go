package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

// NOTE: https://www.hackerrank.com/challenges/sherlock-and-squares/problem?isFullScreen=true
func main() {
	// challenge()

	studies()
}

func studies() {
	var (
		scenarios = [][]int32{
			{1, 25},
			{465868129, 988379794},
			{181510012, 293922871},
			{395151610, 407596310},
			{481403421, 520201871},
			{309804254, 776824625},
			{304742289, 566848910},
			{267554756, 828997506},
			{387224568, 926504395},
			{244571677, 871603971},
			{444567315, 582147918},
			{334350264, 342469009},
			{400474096, 410940967},
			{488907300, 943628573},
			// {26441071, 801576263},
			// {182001128, 267557939},
			// {115732998, 974318256},
			// {192538332, 862327048},
			// {45429427, 307805497},
			// {358658006, 842644090},
			// {92930998, 431601473},
			// {231163983, 893672132},
			// {275221547, 298953978},
			// {351237326, 981399371},
			// {484598992, 985428966},
			// {103405553, 529324202},
			// {37393469, 768655346},
			// {30179914, 482808626},
			// {208821550, 538302223},
			// {154654533, 791652309},
			// {68424067, 854065374},
			// {246956110, 517538724},
			// {51395253, 876949418},
			// {57778758, 368742600},
			// {227566632, 606529208},
			// {360551779, 824396068},
			// {396042598, 543511438},
			// {411041425, 473345854},
			// {469310505, 774761014},
			// {90386202, 342472887},
			// {79110819, 503433812},
			// {444288332, 918848597},
			// {280603787, 548642983},
			// {127990618, 324129686},
			// {479256115, 753819852},
			// {253766533, 449479844},
			// {332623760, 979143500},
			// {122448725, 384124626},
			// {101854341, 491806437},
			// {324275294, 529171976},
			// {378385813, 569430598},
			// {152484569, 322626345},
			// {397853336, 772523199},
			// {126282101, 235322141},
			// {327610187, 686112903},
			// {353962429, 655163953},
			// {391009046, 400936407},
			// {45488493, 815834125},
			// {394001259, 642294484},
			// {111624270, 253455511},
			// {220990588, 255719882},
			// {10907106, 950118387},
			// {177210402, 555916808},
			// {316351449, 741082827},
			// {425556581, 856116768},
			// {421507248, 586064719},
			// {93888922, 642637722},
			// {99843590, 473754317},
			// {8213634, 764585393},
			// {488097822, 840278753},
			// {435342015, 842508745},
			// {349406673, 484900634},
			// {280010853, 370443835},
			// {400888260, 691956893},
			// {421748251, 737067324},
			// {305905572, 386773097},
			// {435191289, 887843065},
			// {260846393, 731941575},
			// {445631476, 519474101},
			// {59915010, 237909218},
			// {306805801, 381681659},
			// {324432467, 859919996},
			// {279649964, 564769173},
			// {118888227, 313212102},
			// {302008280, 466227315},
			// {321239164, 872789513},
			// {207077522, 456470830},
			// {242838944, 464459846},
			// {43481253, 51762643},
			// {208198294, 523145398},
			// {47696526, 364130319},
			// {179957324, 441551904},
			// {50523163, 83566712},
			// {478585845, 765208335},
			// {75756080, 429669067},
			// {475218138, 630450638},
			// {55559381, 939697803},
			// {240414829, 549868781},
			// {420846535, 916723099},
			// {329104468, 908085028},
			// {150336667, 828238028},
		}
	)

	wg := &sync.WaitGroup{}

	for _, scenario := range scenarios {
		wg.Add(1)

		go executeScenario(scenario, wg)
	}

	wg.Wait()
}

func executeScenario(scenario []int32, wg *sync.WaitGroup) {
	defer wg.Done()

	a, b := scenario[0], scenario[1]

	fmt.Printf("checking squares between %v and %v...\n", a, b)
	fmt.Printf("-- [RESULT] there are %v squares between %v and %v --\n", squaresWithMath(a, b), a, b)
}

func challenge() {
	concurrentStart := time.Now()

	var (
		a, b             int32 = 4, 25
		resultConcurrent int32 = squaresConcurrent(a, b)
	)

	concurrentEnd := time.Now()

	fmt.Printf("how many squares in: from %v to %v are %v squares\n", a, b, resultConcurrent)

	fmt.Printf("took %v microseconds to complete in concurrent mode\n", concurrentEnd.Sub(concurrentStart))

	fmt.Println("---- x ---- x ---- x ----")

	// ---- NON CONCURRENT ---- //

	nonConcurrentStart := time.Now()

	var resultNoConcurrent int32 = squares(a, b)

	nonConcurrentEnd := time.Now()

	fmt.Printf("how many squares in: from %v to %v are %v squares\n", a, b, resultNoConcurrent)

	fmt.Printf("took %v microseconds to complete in non concurrent mode\n", nonConcurrentEnd.Sub(nonConcurrentStart))
}

func squares(a int32, b int32) int32 {
	squaresTracker := int32(0)

	for number := a; number <= b; number++ {
		ok := isNumberSquarePerfect(number)

		if ok {
			squaresTracker++
		}
	}

	return squaresTracker
}

func squaresWithMath(a, b int32) int32 {
	startRoot := math.Sqrt(float64(a))
	endRoot := math.Sqrt(float64(b))

	startInt := int32(math.Floor(startRoot))
	endInt := int32(math.Floor(endRoot))

	return (endInt - startInt + 1)
}

func squaresConcurrent(a int32, b int32) int32 {
	squaresTracker := int32(0)

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for number := a; number <= b; number++ {
		wg.Add(1)

		go checkSquareConcurrent(number, &squaresTracker, &wg, &mu)
	}

	wg.Wait()

	return squaresTracker
}

func checkSquareConcurrent(number int32, squaresTracker *int32, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	ok := isNumberSquarePerfect(number)

	if ok {
		defer mu.Unlock()

		mu.Lock()
		*squaresTracker++
	}
}

func isNumberSquarePerfect(number int32) bool {
	// sqrtN, ok := cacheSquares[number]

	// if ok {
	//   return sqrtN == math.Floor(sqrtN)
	// }

	root := int32(math.Sqrt(float64(number)))
	// cacheSquares[number] = root

	return root*root == number
}

var cacheSquares = make(map[int32]float64)
