package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	prices   = map[string]int{"A1": 2, "A2": 3, "A3": 4, "A4": 5, "A5": 8, "A6": 6}
	moneySeq = []int{1, 2, 5, 10}
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			return
		}
		cmds := strings.Split(string(line), ";")
		var machine *SellMachine

		for _, cmd := range cmds {
			if len(cmd) == 0 || cmd == "\n" {
				continue
			}
			info := strings.Split(cmd, " ")

			switch info[0] {
			case "r":
				if len(info) != 3 {
					return
				}
				machine, err = InitSellMachine(info[1], info[2])
				if err != nil {
					return
				}

			case "p":
				b, _ := strconv.Atoi(info[1])
				machine.UpdateBalance(b)
			case "b":
				machine.Sell(info[1])
			case "c":
				machine.Cancel()
			case "q":
				t, _ := strconv.Atoi(info[1])
				machine.Search(t)
			default:
				fmt.Printf("E010:Parameter error")
			}
		}
	}
}

type goods struct {
	Name  string
	Price int
	Num   int
}

type money struct {
	Currency int
	Num      int
}

type GoodsSlince []*goods
type MoneyScline []*money
type SellMachine struct {
	GoodsInfo GoodsSlince
	MoneyInfo MoneyScline
	Balance   int
}

func InitSellMachine(initGoods, initMoney string) (*SellMachine, error) {
	goodsInfo, err := InitGoods(initGoods)
	if err != nil {
		return nil, err
	}
	moneyInfo, err := InitMoney(initMoney)
	if err != nil {
		return nil, err
	}

	fmt.Printf("S001:Initialization is successful\n")
	return &SellMachine{
		GoodsInfo: goodsInfo,
		MoneyInfo: moneyInfo,
	}, nil
}

func (m *SellMachine) UpdateBalance(b int) {
	if isCurrencyValueErr(b) {
		fmt.Printf("%s\n", "E002:Denomination error")
	} else if m.isTooBigger(b) {
		fmt.Printf("%s\n", "E003:Change is not enough, pay fail")
	} else if b > 10 {
		fmt.Printf("%s\n", "E004:Pay the balance is beyond the scope biggest")
	} else if m.outOfSell() {
		fmt.Printf("%s\n", "E005:All the goods sold out")
	} else {
		m.Balance += b
		m.MoneyInfo.Add(b, 1)
		fmt.Printf("S002:Pay success,balance=%d\n", m.Balance)
	}
}

func isCurrencyValueErr(b int) bool {
	err := true
	for _, m := range moneySeq {
		if b == m {
			err = false
		}
	}
	return err
}

func (m *SellMachine) isTooBigger(b int) bool {
	value := 0
	for _, val := range m.MoneyInfo {
		if val.Currency == 1 || val.Currency == 2 {
			value += val.Currency * val.Num
		}
	}
	return value < b
}

func (m *SellMachine) outOfSell() bool {
	for _, g := range m.GoodsInfo {
		if g.Num > 0 {
			return false
		}
	}
	return true
}

func (m *SellMachine) Cancel() {
	if m.Balance <= 0 {
		fmt.Printf("E009:Work failure")
	} else {
		m.doCancel()
	}
}

func (m *SellMachine) doCancel() {
	refund := map[int]int{}
	for i := len(moneySeq) - 1; i >= 0; {
		if m.Balance >= moneySeq[i] && m.MoneyInfo[i].Num > 0 {
			refund[moneySeq[i]]++
			m.MoneyInfo[i].Num--
			m.Balance -= moneySeq[i]
		} else {
			i--
		}
	}
	for _, s := range moneySeq {
		fmt.Printf("%d yuan coin number=%d\n", s, refund[s])
	}
}

func (m *SellMachine) Search(t int) {
	if t == 0 {
		m.searchGoods()
	} else if t == 1 {
		m.searchMoney()
	} else {
		fmt.Printf("E010:Parameter error\n")
	}
}

func (m *SellMachine) searchMoney() {
	for _, s := range m.MoneyInfo {
		fmt.Printf("%d yuan coin number=%d\n", s.Currency, s.Num)
	}
}

func (m *SellMachine) searchGoods() {
	tmpGoodsInfo := m.GoodsInfo.Copy()
	sort.Sort(tmpGoodsInfo)
	for _, g := range tmpGoodsInfo {
		fmt.Printf("%s %d %d\n", g.Name, g.Price, g.Num)
	}
}

func (m *SellMachine) Sell(name string) {
	if m.goodsNotExist(name) {
		fmt.Printf("E006:Goods does not exist\n")
	} else if m.goodsOutOfSell(name) {
		fmt.Printf("E007:The goods sold out\n")
	} else if m.Balance < prices[name] {
		fmt.Printf("E008:Lack of balance\n")
	} else {
		m.doSell(name)
		fmt.Printf("S003:Buy success,balance=%d\n", m.Balance)
	}
}

func (m *SellMachine) goodsNotExist(name string) bool {
	for _, g := range m.GoodsInfo {
		if g.Name == name {
			return false
		}
	}
	return true
}
func (m *SellMachine) goodsOutOfSell(name string) bool {
	for _, g := range m.GoodsInfo {
		if g.Name == name && g.Num > 0 {
			return false
		}
	}
	return true
}
func (m *SellMachine) doSell(name string) {
	for _, g := range m.GoodsInfo {
		if g.Name == name && g.Num > 0 {
			g.Num--
			m.Balance -= g.Price
			break
		}
	}
}

func InitGoods(initCmd string) (GoodsSlince, error) {
	nums := strings.Split(initCmd, "-")
	if len(nums) != len(prices) {
		return nil, fmt.Errorf("Init with error goods cmd: %s", initCmd)
	}

	arrs := make(GoodsSlince, len(prices))
	for i := 0; i < len(prices); i++ {
		num, err := strconv.Atoi(nums[i])
		if err != nil {
			return nil, err
		}
		arrs[i] = &goods{
			Name:  fmt.Sprintf("A%d", i+1),
			Price: prices[fmt.Sprintf("A%d", i+1)],
			Num:   num,
		}
	}
	return arrs, nil
}

func InitMoney(initCmd string) (MoneyScline, error) {
	nums := strings.Split(initCmd, "-")
	if len(nums) != len(moneySeq) {
		return nil, fmt.Errorf("Init with error money cmd: %s", initCmd)
	}

	arrs := make(MoneyScline, len(moneySeq))
	for i := 0; i < len(moneySeq); i++ {
		num, err := strconv.Atoi(nums[i])
		if err != nil {
			return nil, err
		}
		arrs[i] = &money{
			Currency: moneySeq[i],
			Num:      num,
		}
	}
	return arrs, nil
}

func (g GoodsSlince) Copy() GoodsSlince {
	tmpGoodsInfo := make(GoodsSlince, len(g))
	for i, g := range g {
		tmpGoodsInfo[i] = &goods{
			Name:  g.Name,
			Price: g.Price,
			Num:   g.Num,
		}
	}
	return tmpGoodsInfo
}

func (g GoodsSlince) Len() int {
	return len(g)
}
func (g GoodsSlince) Less(i, j int) bool {
	if g[i].Num < g[j].Num {
		return true
	} else if g[i].Num == g[j].Num {
		return g[i].Name < g[j].Name
	} else {
		return false
	}
}

func (g GoodsSlince) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

func (g GoodsSlince) Print() {
	for _, t := range g {
		fmt.Printf("%+v", *t)
	}
	fmt.Println()
}

func (g MoneyScline) Print() {
	for _, t := range g {
		fmt.Printf("%+v", *t)
	}
	fmt.Println()
}

func (g MoneyScline) Add(currency, num int) {
	for _, m := range g {
		if m.Currency == currency {
			m.Num += num
			break
		}
	}
}
